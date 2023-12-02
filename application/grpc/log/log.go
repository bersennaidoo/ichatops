package log

import (
	"context"
	"fmt"
	"log"
	"runtime"
	"sync"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// Needed for log to configure himself
const (
	Ldate = 1 << iota
	Ltime
	Lmicroseconds
	Llongfile
	Lshortfile
	LUTC
	LstdFlags = Ldate | Ltime
)

// Logger provides access to the standard library's default logger.
var Logger *log.Logger = log.Default()

// setup the standard logger with flags.
var std = &logger{flag: LstdFlags}

// pool provides a pool of Event objects to keep our allocations to a minimum.
var pool = &eventPool{
	buf: make(chan *Event, 100),
	pool: sync.Pool{
		New: func() interface{} {
			return &Event{}
		},
	},
}

// eventPool uses a set amount of Event objects and a sync.Pool for overflow.
type eventPool struct {
	buf  chan *Event
	pool sync.Pool
}

// log uses get to get a event from the buf channel or directly from the pool.
func (e *eventPool) get() *Event {
	select {
	case ev := <-e.buf:
		return ev
	default:
	}
	return e.pool.Get().(*Event)
}

// log uses put to put an event to a buf channel or directly to the pool.
func (e *eventPool) put(ev *Event) {
	ev.reset()
	select {
	case e.buf <- ev:
	default:
	}
	e.pool.Put(ev)
}

// Event represents a named event that occurs with attributes key/values pairs.
type Event struct {
	name  string
	attrs []attribute.KeyValue
}

// NewEvent returns a new Event.
func NewEvent(name string) *Event {
	ev := pool.get()
	ev.name = name
	return ev
}

// Resets Event
func (e *Event) reset() {
	e.name = ""
	e.attrs = e.attrs[0:0]
}

// Add adds attributes to an Event.
func (e *Event) Add(k string, i interface{}) {
	if e.name == "" {
		return
	}
	switch v := i.(type) {
	case bool:
		e.attrs = append(e.attrs, attribute.Bool(k, v))
	case []bool:
		e.attrs = append(e.attrs, attribute.BoolSlice(k, v))
	case float64:
		e.attrs = append(e.attrs, attribute.Float64(k, v))
	case []float64:
		e.attrs = append(e.attrs, attribute.Float64Slice(k, v))
	case int:
		e.attrs = append(e.attrs, attribute.Int(k, v))
	case []int:
		e.attrs = append(e.attrs, attribute.IntSlice(k, v))
	case int64:
		e.attrs = append(e.attrs, attribute.Int64(k, v))
	case []int64:
		e.attrs = append(e.attrs, attribute.Int64Slice(k, v))
	case string:
		e.attrs = append(e.attrs, attribute.String(k, v))
	case []string:
		e.attrs = append(e.attrs, attribute.StringSlice(k, v))
	case time.Duration:
		e.attrs = append(e.attrs, attribute.String(k, v.String()))
	default:
		log.Printf("bug: event.Add(): receiveing %T which is not supported", v)
	}
}

// Done renders the Event to the span in the Context.
func (e *Event) Done(ctx context.Context) {
	defer pool.put(e)

	if e.name == "" {
		return
	}
	span := trace.SpanFromContext(ctx)
	if e.attrs == nil {
		return
	}
	span.AddEvent(e.name, trace.WithAttributes(e.attrs...))
}

// Println acts like log.Println() except we log to the OTEL span in the Context.
func Println(ctx context.Context, v ...interface{}) {
	span := trace.SpanFromContext(ctx)
	if !span.IsRecording() {
		return
	}
	std.output(span, 2, fmt.Sprintln(v...))
}

// Printf acts like log.Printf() except we log to the OTEL span in the Context.
func Printf(ctx context.Context, format string, v ...interface{}) {
	span := trace.SpanFromContext(ctx)
	if !span.IsRecording() {
		return
	}
	std.output(span, 2, fmt.Sprintf(format, v...))
}

// SetFlags sets the output flags for the standard logger.
func SetFlags(flag int) {
	std.flag = flag
}

// logger is an implementation of log.Logger that writes to a Span.
type logger struct {
	mu   sync.Mutex
	flag int    // properties
	buf  []byte // for accumulating text to write
}

// Cheap integer to fixed-width decimal ASCII. Give a negative width to avoid zero-padding.
func itoa(buf *[]byte, i int, wid int) {
	// Assemble decimal in reverse order.
	var b [20]byte
	bp := len(b) - 1
	for i >= 10 || wid > 1 {
		wid--
		q := i / 10
		b[bp] = byte('0' + i - q*10)
		bp--
		i = q
	}
	// i < 10
	b[bp] = byte('0' + i)
	*buf = append(*buf, b[bp:]...)
}

func (l *logger) output(span trace.Span, calldepth int, s string) error {
	now := time.Now() // get this early
	var file string
	var line int
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.flag&(Lshortfile|Llongfile) != 0 {
		// Release lock while getting caller info - it's expensive.
		l.mu.Unlock()
		var ok bool
		_, file, line, ok = runtime.Caller(calldepth)
		if !ok {
			file = "???"
			line = 0
		}
		l.mu.Lock()
	}
	l.buf = l.buf[:0]
	l.formatHeader(&l.buf, now, file, line)
	l.buf = append(l.buf, s...)
	if len(s) == 0 || s[len(s)-1] != '\n' {
		l.buf = append(l.buf, '\n')
	}
	span.AddEvent(string(l.buf), trace.WithAttributes(attribute.Bool("log", true)))
	return nil
}

// formatHeader writes log header to buf in following order:
//   - date and/or time (if corresponding flags are provided),
//   - file and line number (if corresponding flags are provided),
func (l *logger) formatHeader(buf *[]byte, t time.Time, file string, line int) {
	if l.flag&(Ldate|Ltime|Lmicroseconds) != 0 {
		if l.flag&LUTC != 0 {
			t = t.UTC()
		}
		if l.flag&Ldate != 0 {
			year, month, day := t.Date()
			itoa(buf, year, 4)
			*buf = append(*buf, '/')
			itoa(buf, int(month), 2)
			*buf = append(*buf, '/')
			itoa(buf, day, 2)
			*buf = append(*buf, ' ')
		}
		if l.flag&(Ltime|Lmicroseconds) != 0 {
			hour, min, sec := t.Clock()
			itoa(buf, hour, 2)
			*buf = append(*buf, ':')
			itoa(buf, min, 2)
			*buf = append(*buf, ':')
			itoa(buf, sec, 2)
			if l.flag&Lmicroseconds != 0 {
				*buf = append(*buf, '.')
				itoa(buf, t.Nanosecond()/1e3, 6)
			}
			*buf = append(*buf, ' ')
		}
	}
	if l.flag&(Lshortfile|Llongfile) != 0 {
		if l.flag&Lshortfile != 0 {
			short := file
			for i := len(file) - 1; i > 0; i-- {
				if file[i] == '/' {
					short = file[i+1:]
					break
				}
			}
			file = short
		}
		*buf = append(*buf, file...)
		*buf = append(*buf, ':')
		itoa(buf, line, -1)
		*buf = append(*buf, ": "...)
	}
}

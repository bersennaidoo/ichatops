// Copyright (c) 2019 The Jaeger Authors.
// Copyright (c) 2018 Uber Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: jaeger/sampling.proto

package jaeger

import (
	_ "github.com/gogo/protobuf/gogoproto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// See description of the SamplingStrategyResponse.strategyType field.
type SamplingStrategyType int32

const (
	SamplingStrategyType_PROBABILISTIC SamplingStrategyType = 0
	SamplingStrategyType_RATE_LIMITING SamplingStrategyType = 1
)

// Enum value maps for SamplingStrategyType.
var (
	SamplingStrategyType_name = map[int32]string{
		0: "PROBABILISTIC",
		1: "RATE_LIMITING",
	}
	SamplingStrategyType_value = map[string]int32{
		"PROBABILISTIC": 0,
		"RATE_LIMITING": 1,
	}
)

func (x SamplingStrategyType) Enum() *SamplingStrategyType {
	p := new(SamplingStrategyType)
	*p = x
	return p
}

func (x SamplingStrategyType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SamplingStrategyType) Descriptor() protoreflect.EnumDescriptor {
	return file_jaeger_sampling_proto_enumTypes[0].Descriptor()
}

func (SamplingStrategyType) Type() protoreflect.EnumType {
	return &file_jaeger_sampling_proto_enumTypes[0]
}

func (x SamplingStrategyType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SamplingStrategyType.Descriptor instead.
func (SamplingStrategyType) EnumDescriptor() ([]byte, []int) {
	return file_jaeger_sampling_proto_rawDescGZIP(), []int{0}
}

// ProbabilisticSamplingStrategy samples traces with a fixed probability.
type ProbabilisticSamplingStrategy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// samplingRate is the sampling probability in the range [0.0, 1.0].
	SamplingRate float64 `protobuf:"fixed64,1,opt,name=samplingRate,proto3" json:"samplingRate,omitempty"`
}

func (x *ProbabilisticSamplingStrategy) Reset() {
	*x = ProbabilisticSamplingStrategy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jaeger_sampling_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProbabilisticSamplingStrategy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProbabilisticSamplingStrategy) ProtoMessage() {}

func (x *ProbabilisticSamplingStrategy) ProtoReflect() protoreflect.Message {
	mi := &file_jaeger_sampling_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProbabilisticSamplingStrategy.ProtoReflect.Descriptor instead.
func (*ProbabilisticSamplingStrategy) Descriptor() ([]byte, []int) {
	return file_jaeger_sampling_proto_rawDescGZIP(), []int{0}
}

func (x *ProbabilisticSamplingStrategy) GetSamplingRate() float64 {
	if x != nil {
		return x.SamplingRate
	}
	return 0
}

// RateLimitingSamplingStrategy samples a fixed number of traces per time interval.
// The typical implementations use the leaky bucket algorithm.
type RateLimitingSamplingStrategy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// TODO this field type should be changed to double, to support rates like 1 per minute.
	MaxTracesPerSecond int32 `protobuf:"varint,1,opt,name=maxTracesPerSecond,proto3" json:"maxTracesPerSecond,omitempty"`
}

func (x *RateLimitingSamplingStrategy) Reset() {
	*x = RateLimitingSamplingStrategy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jaeger_sampling_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RateLimitingSamplingStrategy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RateLimitingSamplingStrategy) ProtoMessage() {}

func (x *RateLimitingSamplingStrategy) ProtoReflect() protoreflect.Message {
	mi := &file_jaeger_sampling_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RateLimitingSamplingStrategy.ProtoReflect.Descriptor instead.
func (*RateLimitingSamplingStrategy) Descriptor() ([]byte, []int) {
	return file_jaeger_sampling_proto_rawDescGZIP(), []int{1}
}

func (x *RateLimitingSamplingStrategy) GetMaxTracesPerSecond() int32 {
	if x != nil {
		return x.MaxTracesPerSecond
	}
	return 0
}

// OperationSamplingStrategy is a sampling strategy for a given operation
// (aka endpoint, span name). Only probabilistic sampling is currently supported.
type OperationSamplingStrategy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Operation             string                         `protobuf:"bytes,1,opt,name=operation,proto3" json:"operation,omitempty"`
	ProbabilisticSampling *ProbabilisticSamplingStrategy `protobuf:"bytes,2,opt,name=probabilisticSampling,proto3" json:"probabilisticSampling,omitempty"`
}

func (x *OperationSamplingStrategy) Reset() {
	*x = OperationSamplingStrategy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jaeger_sampling_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperationSamplingStrategy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperationSamplingStrategy) ProtoMessage() {}

func (x *OperationSamplingStrategy) ProtoReflect() protoreflect.Message {
	mi := &file_jaeger_sampling_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperationSamplingStrategy.ProtoReflect.Descriptor instead.
func (*OperationSamplingStrategy) Descriptor() ([]byte, []int) {
	return file_jaeger_sampling_proto_rawDescGZIP(), []int{2}
}

func (x *OperationSamplingStrategy) GetOperation() string {
	if x != nil {
		return x.Operation
	}
	return ""
}

func (x *OperationSamplingStrategy) GetProbabilisticSampling() *ProbabilisticSamplingStrategy {
	if x != nil {
		return x.ProbabilisticSampling
	}
	return nil
}

// PerOperationSamplingStrategies is a combination of strategies for different endpoints
// as well as some service-wide defaults. It is particularly useful for services whose
// endpoints receive vastly different traffic, so that any single rate of sampling would
// result in either too much data for some endpoints or almost no data for other endpoints.
type PerOperationSamplingStrategies struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// defaultSamplingProbability is the sampling probability for spans that do not match
	// any of the perOperationStrategies.
	DefaultSamplingProbability float64 `protobuf:"fixed64,1,opt,name=defaultSamplingProbability,proto3" json:"defaultSamplingProbability,omitempty"`
	// defaultLowerBoundTracesPerSecond defines a lower-bound rate limit used to ensure that
	// there is some minimal amount of traces sampled for an endpoint that might otherwise
	// be never sampled via probabilistic strategies. The limit is local to a service instance,
	// so if a service is deployed with many (N) instances, the effective minimum rate of sampling
	// will be N times higher. This setting applies to ALL operations, whether or not they match
	// one of the perOperationStrategies.
	DefaultLowerBoundTracesPerSecond float64 `protobuf:"fixed64,2,opt,name=defaultLowerBoundTracesPerSecond,proto3" json:"defaultLowerBoundTracesPerSecond,omitempty"`
	// perOperationStrategies describes sampling strategiesf for individual operations within
	// a given service.
	PerOperationStrategies []*OperationSamplingStrategy `protobuf:"bytes,3,rep,name=perOperationStrategies,proto3" json:"perOperationStrategies,omitempty"`
	// defaultUpperBoundTracesPerSecond defines an upper bound rate limit.
	// However, almost no Jaeger SDKs support this parameter.
	DefaultUpperBoundTracesPerSecond float64 `protobuf:"fixed64,4,opt,name=defaultUpperBoundTracesPerSecond,proto3" json:"defaultUpperBoundTracesPerSecond,omitempty"`
}

func (x *PerOperationSamplingStrategies) Reset() {
	*x = PerOperationSamplingStrategies{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jaeger_sampling_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PerOperationSamplingStrategies) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PerOperationSamplingStrategies) ProtoMessage() {}

func (x *PerOperationSamplingStrategies) ProtoReflect() protoreflect.Message {
	mi := &file_jaeger_sampling_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PerOperationSamplingStrategies.ProtoReflect.Descriptor instead.
func (*PerOperationSamplingStrategies) Descriptor() ([]byte, []int) {
	return file_jaeger_sampling_proto_rawDescGZIP(), []int{3}
}

func (x *PerOperationSamplingStrategies) GetDefaultSamplingProbability() float64 {
	if x != nil {
		return x.DefaultSamplingProbability
	}
	return 0
}

func (x *PerOperationSamplingStrategies) GetDefaultLowerBoundTracesPerSecond() float64 {
	if x != nil {
		return x.DefaultLowerBoundTracesPerSecond
	}
	return 0
}

func (x *PerOperationSamplingStrategies) GetPerOperationStrategies() []*OperationSamplingStrategy {
	if x != nil {
		return x.PerOperationStrategies
	}
	return nil
}

func (x *PerOperationSamplingStrategies) GetDefaultUpperBoundTracesPerSecond() float64 {
	if x != nil {
		return x.DefaultUpperBoundTracesPerSecond
	}
	return 0
}

// SamplingStrategyResponse contains an overall sampling strategy for a given service.
// This type should be treated as a union where only one of the strategy field is present.
type SamplingStrategyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Legacy field that was meant to indicate which one of the strategy fields
	// below is present. This enum was not extended when per-operation strategy
	// was introduced, because extending enum has backwards compatiblity issues.
	// The recommended approach for consumers is to ignore this field and instead
	// checks the other fields being not null (starting with operationSampling).
	// For producers, it is recommended to set this field correctly for probabilistic
	// and rate-limiting strategies, but if per-operation strategy is returned,
	// the enum can be set to 0 (probabilistic).
	StrategyType          SamplingStrategyType            `protobuf:"varint,1,opt,name=strategyType,proto3,enum=jaeger.api_v2.SamplingStrategyType" json:"strategyType,omitempty"`
	ProbabilisticSampling *ProbabilisticSamplingStrategy  `protobuf:"bytes,2,opt,name=probabilisticSampling,proto3" json:"probabilisticSampling,omitempty"`
	RateLimitingSampling  *RateLimitingSamplingStrategy   `protobuf:"bytes,3,opt,name=rateLimitingSampling,proto3" json:"rateLimitingSampling,omitempty"`
	OperationSampling     *PerOperationSamplingStrategies `protobuf:"bytes,4,opt,name=operationSampling,proto3" json:"operationSampling,omitempty"`
}

func (x *SamplingStrategyResponse) Reset() {
	*x = SamplingStrategyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jaeger_sampling_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SamplingStrategyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SamplingStrategyResponse) ProtoMessage() {}

func (x *SamplingStrategyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_jaeger_sampling_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SamplingStrategyResponse.ProtoReflect.Descriptor instead.
func (*SamplingStrategyResponse) Descriptor() ([]byte, []int) {
	return file_jaeger_sampling_proto_rawDescGZIP(), []int{4}
}

func (x *SamplingStrategyResponse) GetStrategyType() SamplingStrategyType {
	if x != nil {
		return x.StrategyType
	}
	return SamplingStrategyType_PROBABILISTIC
}

func (x *SamplingStrategyResponse) GetProbabilisticSampling() *ProbabilisticSamplingStrategy {
	if x != nil {
		return x.ProbabilisticSampling
	}
	return nil
}

func (x *SamplingStrategyResponse) GetRateLimitingSampling() *RateLimitingSamplingStrategy {
	if x != nil {
		return x.RateLimitingSampling
	}
	return nil
}

func (x *SamplingStrategyResponse) GetOperationSampling() *PerOperationSamplingStrategies {
	if x != nil {
		return x.OperationSampling
	}
	return nil
}

// SamplingStrategyParameters defines request parameters for remote sampler.
type SamplingStrategyParameters struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// serviceName is a required argument.
	ServiceName string `protobuf:"bytes,1,opt,name=serviceName,proto3" json:"serviceName,omitempty"`
}

func (x *SamplingStrategyParameters) Reset() {
	*x = SamplingStrategyParameters{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jaeger_sampling_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SamplingStrategyParameters) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SamplingStrategyParameters) ProtoMessage() {}

func (x *SamplingStrategyParameters) ProtoReflect() protoreflect.Message {
	mi := &file_jaeger_sampling_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SamplingStrategyParameters.ProtoReflect.Descriptor instead.
func (*SamplingStrategyParameters) Descriptor() ([]byte, []int) {
	return file_jaeger_sampling_proto_rawDescGZIP(), []int{5}
}

func (x *SamplingStrategyParameters) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

var File_jaeger_sampling_proto protoreflect.FileDescriptor

var file_jaeger_sampling_proto_rawDesc = []byte{
	0x0a, 0x15, 0x6a, 0x61, 0x65, 0x67, 0x65, 0x72, 0x2f, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x69, 0x6e,
	0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6a, 0x61, 0x65, 0x67, 0x65, 0x72, 0x2e,
	0x61, 0x70, 0x69, 0x5f, 0x76, 0x32, 0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x43, 0x0a, 0x1d, 0x50, 0x72,
	0x6f, 0x62, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x63, 0x53, 0x61, 0x6d, 0x70, 0x6c,
	0x69, 0x6e, 0x67, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x12, 0x22, 0x0a, 0x0c, 0x73,
	0x61, 0x6d, 0x70, 0x6c, 0x69, 0x6e, 0x67, 0x52, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x0c, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x69, 0x6e, 0x67, 0x52, 0x61, 0x74, 0x65, 0x22,
	0x4e, 0x0a, 0x1c, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x6e, 0x67, 0x53,
	0x61, 0x6d, 0x70, 0x6c, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x12,
	0x2e, 0x0a, 0x12, 0x6d, 0x61, 0x78, 0x54, 0x72, 0x61, 0x63, 0x65, 0x73, 0x50, 0x65, 0x72, 0x53,
	0x65, 0x63, 0x6f, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x12, 0x6d, 0x61, 0x78,
	0x54, 0x72, 0x61, 0x63, 0x65, 0x73, 0x50, 0x65, 0x72, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x22,
	0x9d, 0x01, 0x0a, 0x19, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x61, 0x6d,
	0x70, 0x6c, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x12, 0x1c, 0x0a,
	0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x62, 0x0a, 0x15, 0x70,
	0x72, 0x6f, 0x62, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x63, 0x53, 0x61, 0x6d, 0x70,
	0x6c, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x6a, 0x61, 0x65,
	0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x5f, 0x76, 0x32, 0x2e, 0x50, 0x72, 0x6f, 0x62, 0x61,
	0x62, 0x69, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x63, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x69, 0x6e, 0x67,
	0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x52, 0x15, 0x70, 0x72, 0x6f, 0x62, 0x61, 0x62,
	0x69, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x63, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x69, 0x6e, 0x67, 0x22,
	0xda, 0x02, 0x0a, 0x1e, 0x50, 0x65, 0x72, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x61, 0x6d, 0x70, 0x6c, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x69,
	0x65, 0x73, 0x12, 0x3e, 0x0a, 0x1a, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x53, 0x61, 0x6d,
	0x70, 0x6c, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x62, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x1a, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x53,
	0x61, 0x6d, 0x70, 0x6c, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x62, 0x61, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x12, 0x4a, 0x0a, 0x20, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x4c, 0x6f, 0x77,
	0x65, 0x72, 0x42, 0x6f, 0x75, 0x6e, 0x64, 0x54, 0x72, 0x61, 0x63, 0x65, 0x73, 0x50, 0x65, 0x72,
	0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x20, 0x64, 0x65,
	0x66, 0x61, 0x75, 0x6c, 0x74, 0x4c, 0x6f, 0x77, 0x65, 0x72, 0x42, 0x6f, 0x75, 0x6e, 0x64, 0x54,
	0x72, 0x61, 0x63, 0x65, 0x73, 0x50, 0x65, 0x72, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x12, 0x60,
	0x0a, 0x16, 0x70, 0x65, 0x72, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74,
	0x72, 0x61, 0x74, 0x65, 0x67, 0x69, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28,
	0x2e, 0x6a, 0x61, 0x65, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x5f, 0x76, 0x32, 0x2e, 0x4f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x69, 0x6e, 0x67,
	0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x52, 0x16, 0x70, 0x65, 0x72, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x69, 0x65, 0x73,
	0x12, 0x4a, 0x0a, 0x20, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x55, 0x70, 0x70, 0x65, 0x72,
	0x42, 0x6f, 0x75, 0x6e, 0x64, 0x54, 0x72, 0x61, 0x63, 0x65, 0x73, 0x50, 0x65, 0x72, 0x53, 0x65,
	0x63, 0x6f, 0x6e, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x20, 0x64, 0x65, 0x66, 0x61,
	0x75, 0x6c, 0x74, 0x55, 0x70, 0x70, 0x65, 0x72, 0x42, 0x6f, 0x75, 0x6e, 0x64, 0x54, 0x72, 0x61,
	0x63, 0x65, 0x73, 0x50, 0x65, 0x72, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x22, 0x85, 0x03, 0x0a,
	0x18, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x0c, 0x73, 0x74, 0x72,
	0x61, 0x74, 0x65, 0x67, 0x79, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x23, 0x2e, 0x6a, 0x61, 0x65, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x5f, 0x76, 0x32, 0x2e,
	0x53, 0x61, 0x6d, 0x70, 0x6c, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x0c, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x62, 0x0a, 0x15, 0x70, 0x72, 0x6f, 0x62, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x73,
	0x74, 0x69, 0x63, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2c, 0x2e, 0x6a, 0x61, 0x65, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x5f, 0x76,
	0x32, 0x2e, 0x50, 0x72, 0x6f, 0x62, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x63, 0x53,
	0x61, 0x6d, 0x70, 0x6c, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x52,
	0x15, 0x70, 0x72, 0x6f, 0x62, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x63, 0x53, 0x61,
	0x6d, 0x70, 0x6c, 0x69, 0x6e, 0x67, 0x12, 0x5f, 0x0a, 0x14, 0x72, 0x61, 0x74, 0x65, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x69, 0x6e, 0x67, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x6a, 0x61, 0x65, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70,
	0x69, 0x5f, 0x76, 0x32, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x6e,
	0x67, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67,
	0x79, 0x52, 0x14, 0x72, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x6e, 0x67, 0x53,
	0x61, 0x6d, 0x70, 0x6c, 0x69, 0x6e, 0x67, 0x12, 0x5b, 0x0a, 0x11, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x69, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x6a, 0x61, 0x65, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x5f,
	0x76, 0x32, 0x2e, 0x50, 0x65, 0x72, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x61, 0x6d, 0x70, 0x6c, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x69, 0x65,
	0x73, 0x52, 0x11, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x61, 0x6d, 0x70,
	0x6c, 0x69, 0x6e, 0x67, 0x22, 0x3e, 0x0a, 0x1a, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x69, 0x6e, 0x67,
	0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65,
	0x72, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x2a, 0x3c, 0x0a, 0x14, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x69, 0x6e, 0x67,
	0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x11, 0x0a, 0x0d,
	0x50, 0x52, 0x4f, 0x42, 0x41, 0x42, 0x49, 0x4c, 0x49, 0x53, 0x54, 0x49, 0x43, 0x10, 0x00, 0x12,
	0x11, 0x0a, 0x0d, 0x52, 0x41, 0x54, 0x45, 0x5f, 0x4c, 0x49, 0x4d, 0x49, 0x54, 0x49, 0x4e, 0x47,
	0x10, 0x01, 0x32, 0xa2, 0x01, 0x0a, 0x0f, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x69, 0x6e, 0x67, 0x4d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12, 0x8e, 0x01, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x53, 0x61,
	0x6d, 0x70, 0x6c, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x12, 0x29,
	0x2e, 0x6a, 0x61, 0x65, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x5f, 0x76, 0x32, 0x2e, 0x53,
	0x61, 0x6d, 0x70, 0x6c, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x1a, 0x27, 0x2e, 0x6a, 0x61, 0x65, 0x67,
	0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x5f, 0x76, 0x32, 0x2e, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x69,
	0x6e, 0x67, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x3a, 0x01, 0x2a, 0x22, 0x18, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x69, 0x6e, 0x67, 0x53,
	0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x42, 0x64, 0xc8, 0xe2, 0x1e, 0x01, 0xd0, 0xe2, 0x1e,
	0x01, 0xe0, 0xe2, 0x1e, 0x01, 0x0a, 0x17, 0x69, 0x6f, 0x2e, 0x6a, 0x61, 0x65, 0x67, 0x65, 0x72,
	0x74, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x5f, 0x76, 0x32, 0x5a, 0x3d,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x65, 0x72, 0x73, 0x65,
	0x6e, 0x6e, 0x61, 0x69, 0x64, 0x6f, 0x6f, 0x2f, 0x69, 0x63, 0x68, 0x61, 0x74, 0x6f, 0x70, 0x73,
	0x2f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6f, 0x70, 0x73,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6a, 0x61, 0x65, 0x67, 0x65, 0x72, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_jaeger_sampling_proto_rawDescOnce sync.Once
	file_jaeger_sampling_proto_rawDescData = file_jaeger_sampling_proto_rawDesc
)

func file_jaeger_sampling_proto_rawDescGZIP() []byte {
	file_jaeger_sampling_proto_rawDescOnce.Do(func() {
		file_jaeger_sampling_proto_rawDescData = protoimpl.X.CompressGZIP(file_jaeger_sampling_proto_rawDescData)
	})
	return file_jaeger_sampling_proto_rawDescData
}

var file_jaeger_sampling_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_jaeger_sampling_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_jaeger_sampling_proto_goTypes = []interface{}{
	(SamplingStrategyType)(0),              // 0: jaeger.api_v2.SamplingStrategyType
	(*ProbabilisticSamplingStrategy)(nil),  // 1: jaeger.api_v2.ProbabilisticSamplingStrategy
	(*RateLimitingSamplingStrategy)(nil),   // 2: jaeger.api_v2.RateLimitingSamplingStrategy
	(*OperationSamplingStrategy)(nil),      // 3: jaeger.api_v2.OperationSamplingStrategy
	(*PerOperationSamplingStrategies)(nil), // 4: jaeger.api_v2.PerOperationSamplingStrategies
	(*SamplingStrategyResponse)(nil),       // 5: jaeger.api_v2.SamplingStrategyResponse
	(*SamplingStrategyParameters)(nil),     // 6: jaeger.api_v2.SamplingStrategyParameters
}
var file_jaeger_sampling_proto_depIdxs = []int32{
	1, // 0: jaeger.api_v2.OperationSamplingStrategy.probabilisticSampling:type_name -> jaeger.api_v2.ProbabilisticSamplingStrategy
	3, // 1: jaeger.api_v2.PerOperationSamplingStrategies.perOperationStrategies:type_name -> jaeger.api_v2.OperationSamplingStrategy
	0, // 2: jaeger.api_v2.SamplingStrategyResponse.strategyType:type_name -> jaeger.api_v2.SamplingStrategyType
	1, // 3: jaeger.api_v2.SamplingStrategyResponse.probabilisticSampling:type_name -> jaeger.api_v2.ProbabilisticSamplingStrategy
	2, // 4: jaeger.api_v2.SamplingStrategyResponse.rateLimitingSampling:type_name -> jaeger.api_v2.RateLimitingSamplingStrategy
	4, // 5: jaeger.api_v2.SamplingStrategyResponse.operationSampling:type_name -> jaeger.api_v2.PerOperationSamplingStrategies
	6, // 6: jaeger.api_v2.SamplingManager.GetSamplingStrategy:input_type -> jaeger.api_v2.SamplingStrategyParameters
	5, // 7: jaeger.api_v2.SamplingManager.GetSamplingStrategy:output_type -> jaeger.api_v2.SamplingStrategyResponse
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_jaeger_sampling_proto_init() }
func file_jaeger_sampling_proto_init() {
	if File_jaeger_sampling_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_jaeger_sampling_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProbabilisticSamplingStrategy); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_jaeger_sampling_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RateLimitingSamplingStrategy); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_jaeger_sampling_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperationSamplingStrategy); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_jaeger_sampling_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PerOperationSamplingStrategies); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_jaeger_sampling_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SamplingStrategyResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_jaeger_sampling_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SamplingStrategyParameters); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_jaeger_sampling_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_jaeger_sampling_proto_goTypes,
		DependencyIndexes: file_jaeger_sampling_proto_depIdxs,
		EnumInfos:         file_jaeger_sampling_proto_enumTypes,
		MessageInfos:      file_jaeger_sampling_proto_msgTypes,
	}.Build()
	File_jaeger_sampling_proto = out.File
	file_jaeger_sampling_proto_rawDesc = nil
	file_jaeger_sampling_proto_goTypes = nil
	file_jaeger_sampling_proto_depIdxs = nil
}

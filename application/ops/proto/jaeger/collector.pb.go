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
// source: jaeger/collector.proto

package jaeger

import (
	model "github.com/bersennaidoo/ichatops/application/ops/proto/jaeger/model"
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

type PostSpansRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Batch *model.Batch `protobuf:"bytes,1,opt,name=batch,proto3" json:"batch,omitempty"`
}

func (x *PostSpansRequest) Reset() {
	*x = PostSpansRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jaeger_collector_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostSpansRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostSpansRequest) ProtoMessage() {}

func (x *PostSpansRequest) ProtoReflect() protoreflect.Message {
	mi := &file_jaeger_collector_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostSpansRequest.ProtoReflect.Descriptor instead.
func (*PostSpansRequest) Descriptor() ([]byte, []int) {
	return file_jaeger_collector_proto_rawDescGZIP(), []int{0}
}

func (x *PostSpansRequest) GetBatch() *model.Batch {
	if x != nil {
		return x.Batch
	}
	return nil
}

type PostSpansResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PostSpansResponse) Reset() {
	*x = PostSpansResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jaeger_collector_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostSpansResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostSpansResponse) ProtoMessage() {}

func (x *PostSpansResponse) ProtoReflect() protoreflect.Message {
	mi := &file_jaeger_collector_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostSpansResponse.ProtoReflect.Descriptor instead.
func (*PostSpansResponse) Descriptor() ([]byte, []int) {
	return file_jaeger_collector_proto_rawDescGZIP(), []int{1}
}

var File_jaeger_collector_proto protoreflect.FileDescriptor

var file_jaeger_collector_proto_rawDesc = []byte{
	0x0a, 0x16, 0x6a, 0x61, 0x65, 0x67, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6a, 0x61, 0x65, 0x67, 0x65, 0x72,
	0x2e, 0x61, 0x70, 0x69, 0x5f, 0x76, 0x32, 0x1a, 0x18, 0x6a, 0x61, 0x65, 0x67, 0x65, 0x72, 0x2f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x44, 0x0a, 0x10, 0x50, 0x6f, 0x73, 0x74, 0x53, 0x70, 0x61,
	0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x05, 0x62, 0x61, 0x74,
	0x63, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6a, 0x61, 0x65, 0x67, 0x65,
	0x72, 0x2e, 0x61, 0x70, 0x69, 0x5f, 0x76, 0x32, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x42, 0x04,
	0xc8, 0xde, 0x1f, 0x00, 0x52, 0x05, 0x62, 0x61, 0x74, 0x63, 0x68, 0x22, 0x13, 0x0a, 0x11, 0x50,
	0x6f, 0x73, 0x74, 0x53, 0x70, 0x61, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x32, 0x7c, 0x0a, 0x10, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x68, 0x0a, 0x09, 0x50, 0x6f, 0x73, 0x74, 0x53, 0x70, 0x61, 0x6e,
	0x73, 0x12, 0x1f, 0x2e, 0x6a, 0x61, 0x65, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x5f, 0x76,
	0x32, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x53, 0x70, 0x61, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x20, 0x2e, 0x6a, 0x61, 0x65, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x5f,
	0x76, 0x32, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x53, 0x70, 0x61, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x3a, 0x01, 0x2a, 0x22,
	0x0d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x73, 0x70, 0x61, 0x6e, 0x73, 0x42, 0x64,
	0xc8, 0xe2, 0x1e, 0x01, 0xd0, 0xe2, 0x1e, 0x01, 0xe0, 0xe2, 0x1e, 0x01, 0x0a, 0x17, 0x69, 0x6f,
	0x2e, 0x6a, 0x61, 0x65, 0x67, 0x65, 0x72, 0x74, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x2e, 0x61,
	0x70, 0x69, 0x5f, 0x76, 0x32, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x62, 0x65, 0x72, 0x73, 0x65, 0x6e, 0x6e, 0x61, 0x69, 0x64, 0x6f, 0x6f, 0x2f, 0x69,
	0x63, 0x68, 0x61, 0x74, 0x6f, 0x70, 0x73, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x6f, 0x70, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6a, 0x61,
	0x65, 0x67, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_jaeger_collector_proto_rawDescOnce sync.Once
	file_jaeger_collector_proto_rawDescData = file_jaeger_collector_proto_rawDesc
)

func file_jaeger_collector_proto_rawDescGZIP() []byte {
	file_jaeger_collector_proto_rawDescOnce.Do(func() {
		file_jaeger_collector_proto_rawDescData = protoimpl.X.CompressGZIP(file_jaeger_collector_proto_rawDescData)
	})
	return file_jaeger_collector_proto_rawDescData
}

var file_jaeger_collector_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_jaeger_collector_proto_goTypes = []interface{}{
	(*PostSpansRequest)(nil),  // 0: jaeger.api_v2.PostSpansRequest
	(*PostSpansResponse)(nil), // 1: jaeger.api_v2.PostSpansResponse
	(*model.Batch)(nil),       // 2: jaeger.api_v2.Batch
}
var file_jaeger_collector_proto_depIdxs = []int32{
	2, // 0: jaeger.api_v2.PostSpansRequest.batch:type_name -> jaeger.api_v2.Batch
	0, // 1: jaeger.api_v2.CollectorService.PostSpans:input_type -> jaeger.api_v2.PostSpansRequest
	1, // 2: jaeger.api_v2.CollectorService.PostSpans:output_type -> jaeger.api_v2.PostSpansResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_jaeger_collector_proto_init() }
func file_jaeger_collector_proto_init() {
	if File_jaeger_collector_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_jaeger_collector_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostSpansRequest); i {
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
		file_jaeger_collector_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostSpansResponse); i {
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
			RawDescriptor: file_jaeger_collector_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_jaeger_collector_proto_goTypes,
		DependencyIndexes: file_jaeger_collector_proto_depIdxs,
		MessageInfos:      file_jaeger_collector_proto_msgTypes,
	}.Build()
	File_jaeger_collector_proto = out.File
	file_jaeger_collector_proto_rawDesc = nil
	file_jaeger_collector_proto_goTypes = nil
	file_jaeger_collector_proto_depIdxs = nil
}
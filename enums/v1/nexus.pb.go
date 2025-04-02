// The MIT License
//
// Copyright (c) 2025 Temporal Technologies Inc. All rights reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by protoc-gen-go. DO NOT EDIT.
// plugins:
// 	protoc-gen-go
// 	protoc
// source: temporal/api/enums/v1/nexus.proto

package enums

import (
	reflect "reflect"
	"strconv"
	sync "sync"
	unsafe "unsafe"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// NexusHandlerErrorRetryBehavior allows nexus handlers to explicity set the retry behavior of a HandlerError. If not
// specified, retry behavior is determined from the error type. For example internal errors are not retryable by default
// unless specified otherwise.
type NexusHandlerErrorRetryBehavior int32

const (
	NEXUS_HANDLER_ERROR_RETRY_BEHAVIOR_UNSPECIFIED NexusHandlerErrorRetryBehavior = 0
	// A handler error is explicitly marked as retryable.
	NEXUS_HANDLER_ERROR_RETRY_BEHAVIOR_RETRYABLE NexusHandlerErrorRetryBehavior = 1
	// A handler error is explicitly marked as non-retryable.
	NEXUS_HANDLER_ERROR_RETRY_BEHAVIOR_NON_RETRYABLE NexusHandlerErrorRetryBehavior = 2
)

// Enum value maps for NexusHandlerErrorRetryBehavior.
var (
	NexusHandlerErrorRetryBehavior_name = map[int32]string{
		0: "NEXUS_HANDLER_ERROR_RETRY_BEHAVIOR_UNSPECIFIED",
		1: "NEXUS_HANDLER_ERROR_RETRY_BEHAVIOR_RETRYABLE",
		2: "NEXUS_HANDLER_ERROR_RETRY_BEHAVIOR_NON_RETRYABLE",
	}
	NexusHandlerErrorRetryBehavior_value = map[string]int32{
		"NEXUS_HANDLER_ERROR_RETRY_BEHAVIOR_UNSPECIFIED":   0,
		"NEXUS_HANDLER_ERROR_RETRY_BEHAVIOR_RETRYABLE":     1,
		"NEXUS_HANDLER_ERROR_RETRY_BEHAVIOR_NON_RETRYABLE": 2,
	}
)

func (x NexusHandlerErrorRetryBehavior) Enum() *NexusHandlerErrorRetryBehavior {
	p := new(NexusHandlerErrorRetryBehavior)
	*p = x
	return p
}

func (x NexusHandlerErrorRetryBehavior) String() string {
	switch x {
	case NEXUS_HANDLER_ERROR_RETRY_BEHAVIOR_UNSPECIFIED:
		return "Unspecified"
	case NEXUS_HANDLER_ERROR_RETRY_BEHAVIOR_RETRYABLE:
		return "Retryable"
	case NEXUS_HANDLER_ERROR_RETRY_BEHAVIOR_NON_RETRYABLE:
		return "NonRetryable"
	default:
		return strconv.Itoa(int(x))
	}

}

func (NexusHandlerErrorRetryBehavior) Descriptor() protoreflect.EnumDescriptor {
	return file_temporal_api_enums_v1_nexus_proto_enumTypes[0].Descriptor()
}

func (NexusHandlerErrorRetryBehavior) Type() protoreflect.EnumType {
	return &file_temporal_api_enums_v1_nexus_proto_enumTypes[0]
}

func (x NexusHandlerErrorRetryBehavior) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NexusHandlerErrorRetryBehavior.Descriptor instead.
func (NexusHandlerErrorRetryBehavior) EnumDescriptor() ([]byte, []int) {
	return file_temporal_api_enums_v1_nexus_proto_rawDescGZIP(), []int{0}
}

var File_temporal_api_enums_v1_nexus_proto protoreflect.FileDescriptor

const file_temporal_api_enums_v1_nexus_proto_rawDesc = "" +
	"\n" +
	"!temporal/api/enums/v1/nexus.proto\x12\x15temporal.api.enums.v1*\xbc\x01\n" +
	"\x1eNexusHandlerErrorRetryBehavior\x122\n" +
	".NEXUS_HANDLER_ERROR_RETRY_BEHAVIOR_UNSPECIFIED\x10\x00\x120\n" +
	",NEXUS_HANDLER_ERROR_RETRY_BEHAVIOR_RETRYABLE\x10\x01\x124\n" +
	"0NEXUS_HANDLER_ERROR_RETRY_BEHAVIOR_NON_RETRYABLE\x10\x02B\x82\x01\n" +
	"\x18io.temporal.api.enums.v1B\n" +
	"NexusProtoP\x01Z!go.temporal.io/api/enums/v1;enums\xaa\x02\x17Temporalio.Api.Enums.V1\xea\x02\x1aTemporalio::Api::Enums::V1b\x06proto3"

var (
	file_temporal_api_enums_v1_nexus_proto_rawDescOnce sync.Once
	file_temporal_api_enums_v1_nexus_proto_rawDescData []byte
)

func file_temporal_api_enums_v1_nexus_proto_rawDescGZIP() []byte {
	file_temporal_api_enums_v1_nexus_proto_rawDescOnce.Do(func() {
		file_temporal_api_enums_v1_nexus_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_temporal_api_enums_v1_nexus_proto_rawDesc), len(file_temporal_api_enums_v1_nexus_proto_rawDesc)))
	})
	return file_temporal_api_enums_v1_nexus_proto_rawDescData
}

var file_temporal_api_enums_v1_nexus_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_temporal_api_enums_v1_nexus_proto_goTypes = []any{
	(NexusHandlerErrorRetryBehavior)(0), // 0: temporal.api.enums.v1.NexusHandlerErrorRetryBehavior
}
var file_temporal_api_enums_v1_nexus_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_temporal_api_enums_v1_nexus_proto_init() }
func file_temporal_api_enums_v1_nexus_proto_init() {
	if File_temporal_api_enums_v1_nexus_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_temporal_api_enums_v1_nexus_proto_rawDesc), len(file_temporal_api_enums_v1_nexus_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_temporal_api_enums_v1_nexus_proto_goTypes,
		DependencyIndexes: file_temporal_api_enums_v1_nexus_proto_depIdxs,
		EnumInfos:         file_temporal_api_enums_v1_nexus_proto_enumTypes,
	}.Build()
	File_temporal_api_enums_v1_nexus_proto = out.File
	file_temporal_api_enums_v1_nexus_proto_goTypes = nil
	file_temporal_api_enums_v1_nexus_proto_depIdxs = nil
}

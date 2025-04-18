// The MIT License
//
// Copyright (c) 2022 Temporal Technologies Inc.  All rights reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
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
// source: temporal/api/cloud/resource/v1/message.proto

package resource

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

type ResourceState int32

const (
	RESOURCE_STATE_UNSPECIFIED       ResourceState = 0
	RESOURCE_STATE_ACTIVATING        ResourceState = 1  // The resource is being activated.
	RESOURCE_STATE_ACTIVATION_FAILED ResourceState = 2  // The resource failed to activate. This is an error state. Reach out to support for remediation.
	RESOURCE_STATE_ACTIVE            ResourceState = 3  // The resource is active and ready to use.
	RESOURCE_STATE_UPDATING          ResourceState = 4  // The resource is being updated.
	RESOURCE_STATE_UPDATE_FAILED     ResourceState = 5  // The resource failed to update. This is an error state. Reach out to support for remediation.
	RESOURCE_STATE_DELETING          ResourceState = 6  // The resource is being deleted.
	RESOURCE_STATE_DELETE_FAILED     ResourceState = 7  // The resource failed to delete. This is an error state. Reach out to support for remediation.
	RESOURCE_STATE_DELETED           ResourceState = 8  // The resource has been deleted.
	RESOURCE_STATE_SUSPENDED         ResourceState = 9  // The resource is suspended and not available for use. Reach out to support for remediation.
	RESOURCE_STATE_EXPIRED           ResourceState = 10 // The resource has expired and is no longer available for use.
)

// Enum value maps for ResourceState.
var (
	ResourceState_name = map[int32]string{
		0:  "RESOURCE_STATE_UNSPECIFIED",
		1:  "RESOURCE_STATE_ACTIVATING",
		2:  "RESOURCE_STATE_ACTIVATION_FAILED",
		3:  "RESOURCE_STATE_ACTIVE",
		4:  "RESOURCE_STATE_UPDATING",
		5:  "RESOURCE_STATE_UPDATE_FAILED",
		6:  "RESOURCE_STATE_DELETING",
		7:  "RESOURCE_STATE_DELETE_FAILED",
		8:  "RESOURCE_STATE_DELETED",
		9:  "RESOURCE_STATE_SUSPENDED",
		10: "RESOURCE_STATE_EXPIRED",
	}
	ResourceState_value = map[string]int32{
		"RESOURCE_STATE_UNSPECIFIED":       0,
		"RESOURCE_STATE_ACTIVATING":        1,
		"RESOURCE_STATE_ACTIVATION_FAILED": 2,
		"RESOURCE_STATE_ACTIVE":            3,
		"RESOURCE_STATE_UPDATING":          4,
		"RESOURCE_STATE_UPDATE_FAILED":     5,
		"RESOURCE_STATE_DELETING":          6,
		"RESOURCE_STATE_DELETE_FAILED":     7,
		"RESOURCE_STATE_DELETED":           8,
		"RESOURCE_STATE_SUSPENDED":         9,
		"RESOURCE_STATE_EXPIRED":           10,
	}
)

func (x ResourceState) Enum() *ResourceState {
	p := new(ResourceState)
	*p = x
	return p
}

func (x ResourceState) String() string {
	switch x {
	case RESOURCE_STATE_UNSPECIFIED:
		return "Unspecified"
	case RESOURCE_STATE_ACTIVATING:
		return "Activating"
	case RESOURCE_STATE_ACTIVATION_FAILED:
		return "ActivationFailed"
	case RESOURCE_STATE_ACTIVE:
		return "Active"
	case RESOURCE_STATE_UPDATING:
		return "Updating"
	case RESOURCE_STATE_UPDATE_FAILED:
		return "UpdateFailed"
	case RESOURCE_STATE_DELETING:
		return "Deleting"
	case RESOURCE_STATE_DELETE_FAILED:
		return "DeleteFailed"
	case RESOURCE_STATE_DELETED:
		return "Deleted"
	case RESOURCE_STATE_SUSPENDED:

		// Deprecated: Use ResourceState.Descriptor instead.
		return "Suspended"
	case RESOURCE_STATE_EXPIRED:
		return "Expired"
	default:
		return strconv.Itoa(int(x))
	}

}

func (ResourceState) Descriptor() protoreflect.EnumDescriptor {
	return file_temporal_api_cloud_resource_v1_message_proto_enumTypes[0].Descriptor()
}

func (ResourceState) Type() protoreflect.EnumType {
	return &file_temporal_api_cloud_resource_v1_message_proto_enumTypes[0]
}

func (x ResourceState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

func (ResourceState) EnumDescriptor() ([]byte, []int) {
	return file_temporal_api_cloud_resource_v1_message_proto_rawDescGZIP(), []int{0}
}

var File_temporal_api_cloud_resource_v1_message_proto protoreflect.FileDescriptor

const file_temporal_api_cloud_resource_v1_message_proto_rawDesc = "" +
	"\n" +
	",temporal/api/cloud/resource/v1/message.proto\x12\x1etemporal.api.cloud.resource.v1*\xe3\x02\n" +
	"\rResourceState\x12\x1e\n" +
	"\x1aRESOURCE_STATE_UNSPECIFIED\x10\x00\x12\x1d\n" +
	"\x19RESOURCE_STATE_ACTIVATING\x10\x01\x12$\n" +
	" RESOURCE_STATE_ACTIVATION_FAILED\x10\x02\x12\x19\n" +
	"\x15RESOURCE_STATE_ACTIVE\x10\x03\x12\x1b\n" +
	"\x17RESOURCE_STATE_UPDATING\x10\x04\x12 \n" +
	"\x1cRESOURCE_STATE_UPDATE_FAILED\x10\x05\x12\x1b\n" +
	"\x17RESOURCE_STATE_DELETING\x10\x06\x12 \n" +
	"\x1cRESOURCE_STATE_DELETE_FAILED\x10\a\x12\x1a\n" +
	"\x16RESOURCE_STATE_DELETED\x10\b\x12\x1c\n" +
	"\x18RESOURCE_STATE_SUSPENDED\x10\t\x12\x1a\n" +
	"\x16RESOURCE_STATE_EXPIRED\x10\n" +
	"B\xac\x01\n" +
	"!io.temporal.api.cloud.resource.v1B\fMessageProtoP\x01Z-go.temporal.io/api/cloud/resource/v1;resource\xaa\x02 Temporalio.Api.Cloud.Resource.V1\xea\x02$Temporalio::Api::Cloud::Resource::V1b\x06proto3"

var (
	file_temporal_api_cloud_resource_v1_message_proto_rawDescOnce sync.Once
	file_temporal_api_cloud_resource_v1_message_proto_rawDescData []byte
)

func file_temporal_api_cloud_resource_v1_message_proto_rawDescGZIP() []byte {
	file_temporal_api_cloud_resource_v1_message_proto_rawDescOnce.Do(func() {
		file_temporal_api_cloud_resource_v1_message_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_temporal_api_cloud_resource_v1_message_proto_rawDesc), len(file_temporal_api_cloud_resource_v1_message_proto_rawDesc)))
	})
	return file_temporal_api_cloud_resource_v1_message_proto_rawDescData
}

var file_temporal_api_cloud_resource_v1_message_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_temporal_api_cloud_resource_v1_message_proto_goTypes = []any{
	(ResourceState)(0), // 0: temporal.api.cloud.resource.v1.ResourceState
}
var file_temporal_api_cloud_resource_v1_message_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_temporal_api_cloud_resource_v1_message_proto_init() }
func file_temporal_api_cloud_resource_v1_message_proto_init() {
	if File_temporal_api_cloud_resource_v1_message_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_temporal_api_cloud_resource_v1_message_proto_rawDesc), len(file_temporal_api_cloud_resource_v1_message_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_temporal_api_cloud_resource_v1_message_proto_goTypes,
		DependencyIndexes: file_temporal_api_cloud_resource_v1_message_proto_depIdxs,
		EnumInfos:         file_temporal_api_cloud_resource_v1_message_proto_enumTypes,
	}.Build()
	File_temporal_api_cloud_resource_v1_message_proto = out.File
	file_temporal_api_cloud_resource_v1_message_proto_goTypes = nil
	file_temporal_api_cloud_resource_v1_message_proto_depIdxs = nil
}

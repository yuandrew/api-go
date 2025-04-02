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
// source: temporal/api/cloud/operation/v1/message.proto

package operation

import (
	reflect "reflect"
	"strconv"
	sync "sync"
	unsafe "unsafe"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AsyncOperation_State int32

const (
	AsyncOperation_STATE_UNSPECIFIED AsyncOperation_State = 0
	AsyncOperation_STATE_PENDING     AsyncOperation_State = 1
	AsyncOperation_STATE_IN_PROGRESS AsyncOperation_State = // The operation is pending.
	2
	AsyncOperation_STATE_FAILED AsyncOperation_State = // The operation is in progress.
	3
	AsyncOperation_STATE_CANCELLED AsyncOperation_State = // The operation failed, check failure_reason for more details.
	4
	AsyncOperation_STATE_FULFILLED AsyncOperation_State = // The operation was cancelled.
	5                              // The operation was fulfilled.
)

// Enum value maps for AsyncOperation_State.
var (
	AsyncOperation_State_name = map[int32]string{
		0: "STATE_UNSPECIFIED",
		1: "STATE_PENDING",
		2: "STATE_IN_PROGRESS",
		3: "STATE_FAILED",
		4: "STATE_CANCELLED",
		5: "STATE_FULFILLED",
	}
	AsyncOperation_State_value = map[string]int32{
		"STATE_UNSPECIFIED": 0,
		"STATE_PENDING":     1,
		"STATE_IN_PROGRESS": 2,
		"STATE_FAILED":      3,
		"STATE_CANCELLED":   4,
		"STATE_FULFILLED":   5,
	}
)

func (x AsyncOperation_State) Enum() *AsyncOperation_State {
	p := new(AsyncOperation_State)
	*p = x
	return p
}

func (x AsyncOperation_State) String() string {
	switch x {
	case AsyncOperation_STATE_UNSPECIFIED:
		return "AsyncOperationStateUnspecified"
	case AsyncOperation_STATE_PENDING:
		return "AsyncOperationStatePending"
	case AsyncOperation_STATE_IN_PROGRESS:
		return "AsyncOperationStateInProgress"
	case AsyncOperation_STATE_FAILED:
		return "AsyncOperationStateFailed"
	case AsyncOperation_STATE_CANCELLED:
		return "AsyncOperationStateCancelled"
	case AsyncOperation_STATE_FULFILLED:
		return "AsyncOperationStateFulfilled"
	default:
		return strconv.Itoa(int(x))
	}

}

func (AsyncOperation_State) Descriptor() protoreflect.EnumDescriptor {
	return file_temporal_api_cloud_operation_v1_message_proto_enumTypes[0].Descriptor()
}

func (AsyncOperation_State) Type() protoreflect.EnumType {
	return &file_temporal_api_cloud_operation_v1_message_proto_enumTypes[0]
}

func (x AsyncOperation_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AsyncOperation_State.Descriptor instead.
func (AsyncOperation_State) EnumDescriptor() ([]byte, []int) {
	return file_temporal_api_cloud_operation_v1_message_proto_rawDescGZIP(), []int{0, 0}
}

type AsyncOperation struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The operation id.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The current state of this operation.
	// Possible values are: pending, in_progress, failed, cancelled, fulfilled.
	// Deprecated: Not supported after 2024-10-01-00 api version. Use state instead.
	// temporal:versioning:max_version=2024-10-01-00
	//
	// Deprecated: Marked as deprecated in temporal/api/cloud/operation/v1/message.proto.
	StateDeprecated string `protobuf:"bytes,2,opt,name=state_deprecated,json=stateDeprecated,proto3" json:"state_deprecated,omitempty"`
	// The current state of this operation.
	// temporal:versioning:min_version=2024-10-01-00
	// temporal:enums:replaces=state_deprecated
	State AsyncOperation_State `protobuf:"varint,9,opt,name=state,proto3,enum=temporal.api.cloud.operation.v1.AsyncOperation_State" json:"state,omitempty"`
	// The recommended duration to check back for an update in the operation's state.
	CheckDuration *durationpb.Duration `protobuf:"bytes,3,opt,name=check_duration,json=checkDuration,proto3" json:"check_duration,omitempty"`
	// The type of operation being performed.
	OperationType string `protobuf:"bytes,4,opt,name=operation_type,json=operationType,proto3" json:"operation_type,omitempty"`
	// The input to the operation being performed.
	//
	// (-- api-linter: core::0146::any=disabled --)
	OperationInput *anypb.Any `protobuf:"bytes,5,opt,name=operation_input,json=operationInput,proto3" json:"operation_input,omitempty"`
	// If the operation failed, the reason for the failure.
	FailureReason string `protobuf:"bytes,6,opt,name=failure_reason,json=failureReason,proto3" json:"failure_reason,omitempty"`
	// The date and time when the operation initiated.
	StartedTime *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=started_time,json=startedTime,proto3" json:"started_time,omitempty"`
	// The date and time when the operation completed.
	FinishedTime  *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=finished_time,json=finishedTime,proto3" json:"finished_time,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AsyncOperation) Reset() {
	*x = AsyncOperation{}
	mi := &file_temporal_api_cloud_operation_v1_message_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AsyncOperation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AsyncOperation) ProtoMessage() {}

func (x *AsyncOperation) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_cloud_operation_v1_message_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AsyncOperation.ProtoReflect.Descriptor instead.
func (*AsyncOperation) Descriptor() ([]byte, []int) {
	return file_temporal_api_cloud_operation_v1_message_proto_rawDescGZIP(), []int{0}
}

func (x *AsyncOperation) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// Deprecated: Marked as deprecated in temporal/api/cloud/operation/v1/message.proto.
func (x *AsyncOperation) GetStateDeprecated() string {
	if x != nil {
		return x.StateDeprecated
	}
	return ""
}

func (x *AsyncOperation) GetState() AsyncOperation_State {
	if x != nil {
		return x.State
	}
	return AsyncOperation_STATE_UNSPECIFIED
}

func (x *AsyncOperation) GetCheckDuration() *durationpb.Duration {
	if x != nil {
		return x.CheckDuration
	}
	return nil
}

func (x *AsyncOperation) GetOperationType() string {
	if x != nil {
		return x.OperationType
	}
	return ""
}

func (x *AsyncOperation) GetOperationInput() *anypb.Any {
	if x != nil {
		return x.OperationInput
	}
	return nil
}

func (x *AsyncOperation) GetFailureReason() string {
	if x != nil {
		return x.FailureReason
	}
	return ""
}

func (x *AsyncOperation) GetStartedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartedTime
	}
	return nil
}

func (x *AsyncOperation) GetFinishedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.FinishedTime
	}
	return nil
}

var File_temporal_api_cloud_operation_v1_message_proto protoreflect.FileDescriptor

const file_temporal_api_cloud_operation_v1_message_proto_rawDesc = "" +
	"\n" +
	"-temporal/api/cloud/operation/v1/message.proto\x12\x1ftemporal.api.cloud.operation.v1\x1a\x1egoogle/protobuf/duration.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x19google/protobuf/any.proto\"\xf2\x04\n" +
	"\x0eAsyncOperation\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12-\n" +
	"\x10state_deprecated\x18\x02 \x01(\tB\x02\x18\x01R\x0fstateDeprecated\x12K\n" +
	"\x05state\x18\t \x01(\x0e25.temporal.api.cloud.operation.v1.AsyncOperation.StateR\x05state\x12@\n" +
	"\x0echeck_duration\x18\x03 \x01(\v2\x19.google.protobuf.DurationR\rcheckDuration\x12%\n" +
	"\x0eoperation_type\x18\x04 \x01(\tR\roperationType\x12=\n" +
	"\x0foperation_input\x18\x05 \x01(\v2\x14.google.protobuf.AnyR\x0eoperationInput\x12%\n" +
	"\x0efailure_reason\x18\x06 \x01(\tR\rfailureReason\x12=\n" +
	"\fstarted_time\x18\a \x01(\v2\x1a.google.protobuf.TimestampR\vstartedTime\x12?\n" +
	"\rfinished_time\x18\b \x01(\v2\x1a.google.protobuf.TimestampR\ffinishedTime\"\x84\x01\n" +
	"\x05State\x12\x15\n" +
	"\x11STATE_UNSPECIFIED\x10\x00\x12\x11\n" +
	"\rSTATE_PENDING\x10\x01\x12\x15\n" +
	"\x11STATE_IN_PROGRESS\x10\x02\x12\x10\n" +
	"\fSTATE_FAILED\x10\x03\x12\x13\n" +
	"\x0fSTATE_CANCELLED\x10\x04\x12\x13\n" +
	"\x0fSTATE_FULFILLED\x10\x05B\xb1\x01\n" +
	"\"io.temporal.api.cloud.operation.v1B\fMessageProtoP\x01Z/go.temporal.io/api/cloud/operation/v1;operation\xaa\x02!Temporalio.Api.Cloud.Operation.V1\xea\x02%Temporalio::Api::Cloud::Operation::V1b\x06proto3"

var (
	file_temporal_api_cloud_operation_v1_message_proto_rawDescOnce sync.Once
	file_temporal_api_cloud_operation_v1_message_proto_rawDescData []byte
)

func file_temporal_api_cloud_operation_v1_message_proto_rawDescGZIP() []byte {
	file_temporal_api_cloud_operation_v1_message_proto_rawDescOnce.Do(func() {
		file_temporal_api_cloud_operation_v1_message_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_temporal_api_cloud_operation_v1_message_proto_rawDesc), len(file_temporal_api_cloud_operation_v1_message_proto_rawDesc)))
	})
	return file_temporal_api_cloud_operation_v1_message_proto_rawDescData
}

var file_temporal_api_cloud_operation_v1_message_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_temporal_api_cloud_operation_v1_message_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_temporal_api_cloud_operation_v1_message_proto_goTypes = []any{
	(AsyncOperation_State)(0),     // 0: temporal.api.cloud.operation.v1.AsyncOperation.State
	(*AsyncOperation)(nil),        // 1: temporal.api.cloud.operation.v1.AsyncOperation
	(*durationpb.Duration)(nil),   // 2: google.protobuf.Duration
	(*anypb.Any)(nil),             // 3: google.protobuf.Any
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_temporal_api_cloud_operation_v1_message_proto_depIdxs = []int32{
	0, // 0: temporal.api.cloud.operation.v1.AsyncOperation.state:type_name -> temporal.api.cloud.operation.v1.AsyncOperation.State
	2, // 1: temporal.api.cloud.operation.v1.AsyncOperation.check_duration:type_name -> google.protobuf.Duration
	3, // 2: temporal.api.cloud.operation.v1.AsyncOperation.operation_input:type_name -> google.protobuf.Any
	4, // 3: temporal.api.cloud.operation.v1.AsyncOperation.started_time:type_name -> google.protobuf.Timestamp
	4, // 4: temporal.api.cloud.operation.v1.AsyncOperation.finished_time:type_name -> google.protobuf.Timestamp
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_temporal_api_cloud_operation_v1_message_proto_init() }
func file_temporal_api_cloud_operation_v1_message_proto_init() {
	if File_temporal_api_cloud_operation_v1_message_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_temporal_api_cloud_operation_v1_message_proto_rawDesc), len(file_temporal_api_cloud_operation_v1_message_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_temporal_api_cloud_operation_v1_message_proto_goTypes,
		DependencyIndexes: file_temporal_api_cloud_operation_v1_message_proto_depIdxs,
		EnumInfos:         file_temporal_api_cloud_operation_v1_message_proto_enumTypes,
		MessageInfos:      file_temporal_api_cloud_operation_v1_message_proto_msgTypes,
	}.Build()
	File_temporal_api_cloud_operation_v1_message_proto = out.File
	file_temporal_api_cloud_operation_v1_message_proto_goTypes = nil
	file_temporal_api_cloud_operation_v1_message_proto_depIdxs = nil
}

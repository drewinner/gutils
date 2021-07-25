// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.7.0
// source: proto/rpc/task.proto

package rpc

import (
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

type TaskReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                //任务id
	LogId      int32  `protobuf:"varint,2,opt,name=logId,proto3" json:"logId,omitempty"`          // 日志id
	JobHandler string `protobuf:"bytes,3,opt,name=jobHandler,proto3" json:"jobHandler,omitempty"` //任务标识、找到对应的handler
	Params     string `protobuf:"bytes,4,opt,name=params,proto3" json:"params,omitempty"`         //任务参数
}

func (x *TaskReq) Reset() {
	*x = TaskReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rpc_task_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskReq) ProtoMessage() {}

func (x *TaskReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rpc_task_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskReq.ProtoReflect.Descriptor instead.
func (*TaskReq) Descriptor() ([]byte, []int) {
	return file_proto_rpc_task_proto_rawDescGZIP(), []int{0}
}

func (x *TaskReq) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TaskReq) GetLogId() int32 {
	if x != nil {
		return x.LogId
	}
	return 0
}

func (x *TaskReq) GetJobHandler() string {
	if x != nil {
		return x.JobHandler
	}
	return ""
}

func (x *TaskReq) GetParams() string {
	if x != nil {
		return x.Params
	}
	return ""
}

type TaskResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                      //任务id
	LogId         int32  `protobuf:"varint,2,opt,name=logId,proto3" json:"logId,omitempty"`                // 日志id
	Status        int32  `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`              //执行状态
	ExecStartTime string `protobuf:"bytes,4,opt,name=execStartTime,proto3" json:"execStartTime,omitempty"` //执行器开始执行时间
	ExecEndTime   string `protobuf:"bytes,5,opt,name=execEndTime,proto3" json:"execEndTime,omitempty"`     //执行器执行完时间
	LogMsg        string `protobuf:"bytes,6,opt,name=logMsg,proto3" json:"logMsg,omitempty"`               //日志信息
}

func (x *TaskResp) Reset() {
	*x = TaskResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rpc_task_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskResp) ProtoMessage() {}

func (x *TaskResp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rpc_task_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskResp.ProtoReflect.Descriptor instead.
func (*TaskResp) Descriptor() ([]byte, []int) {
	return file_proto_rpc_task_proto_rawDescGZIP(), []int{1}
}

func (x *TaskResp) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TaskResp) GetLogId() int32 {
	if x != nil {
		return x.LogId
	}
	return 0
}

func (x *TaskResp) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *TaskResp) GetExecStartTime() string {
	if x != nil {
		return x.ExecStartTime
	}
	return ""
}

func (x *TaskResp) GetExecEndTime() string {
	if x != nil {
		return x.ExecEndTime
	}
	return ""
}

func (x *TaskResp) GetLogMsg() string {
	if x != nil {
		return x.LogMsg
	}
	return ""
}

var File_proto_rpc_task_proto protoreflect.FileDescriptor

var file_proto_rpc_task_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x74, 0x61, 0x73, 0x6b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x67, 0x0a, 0x07, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65,
	0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x6c, 0x6f, 0x67, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x6a, 0x6f, 0x62, 0x48, 0x61,
	0x6e, 0x64, 0x6c, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6a, 0x6f, 0x62,
	0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22,
	0xa8, 0x01, 0x0a, 0x08, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x6c, 0x6f, 0x67, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x6f, 0x67,
	0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x65, 0x78,
	0x65, 0x63, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x65, 0x78, 0x65, 0x63, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x65, 0x78, 0x65, 0x63, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x78, 0x65, 0x63, 0x45, 0x6e, 0x64, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x6f, 0x67, 0x4d, 0x73, 0x67, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x6c, 0x6f, 0x67, 0x4d, 0x73, 0x67, 0x32, 0x29, 0x0a, 0x0b, 0x54, 0x61,
	0x73, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x03, 0x52, 0x75, 0x6e,
	0x12, 0x08, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x1a, 0x09, 0x2e, 0x54, 0x61, 0x73,
	0x6b, 0x52, 0x65, 0x73, 0x70, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x72, 0x70, 0x63, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_rpc_task_proto_rawDescOnce sync.Once
	file_proto_rpc_task_proto_rawDescData = file_proto_rpc_task_proto_rawDesc
)

func file_proto_rpc_task_proto_rawDescGZIP() []byte {
	file_proto_rpc_task_proto_rawDescOnce.Do(func() {
		file_proto_rpc_task_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_rpc_task_proto_rawDescData)
	})
	return file_proto_rpc_task_proto_rawDescData
}

var file_proto_rpc_task_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_rpc_task_proto_goTypes = []interface{}{
	(*TaskReq)(nil),  // 0: TaskReq
	(*TaskResp)(nil), // 1: TaskResp
}
var file_proto_rpc_task_proto_depIdxs = []int32{
	0, // 0: TaskService.Run:input_type -> TaskReq
	1, // 1: TaskService.Run:output_type -> TaskResp
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_rpc_task_proto_init() }
func file_proto_rpc_task_proto_init() {
	if File_proto_rpc_task_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_rpc_task_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskReq); i {
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
		file_proto_rpc_task_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskResp); i {
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
			RawDescriptor: file_proto_rpc_task_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_rpc_task_proto_goTypes,
		DependencyIndexes: file_proto_rpc_task_proto_depIdxs,
		MessageInfos:      file_proto_rpc_task_proto_msgTypes,
	}.Build()
	File_proto_rpc_task_proto = out.File
	file_proto_rpc_task_proto_rawDesc = nil
	file_proto_rpc_task_proto_goTypes = nil
	file_proto_rpc_task_proto_depIdxs = nil
}

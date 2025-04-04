// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: task.proto

package api

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateTaskRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OwnerId       int32                  `protobuf:"varint,1,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	Title         string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Deadline      *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=deadline,proto3" json:"deadline,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateTaskRequest) Reset() {
	*x = CreateTaskRequest{}
	mi := &file_task_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTaskRequest) ProtoMessage() {}

func (x *CreateTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_task_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTaskRequest.ProtoReflect.Descriptor instead.
func (*CreateTaskRequest) Descriptor() ([]byte, []int) {
	return file_task_proto_rawDescGZIP(), []int{0}
}

func (x *CreateTaskRequest) GetOwnerId() int32 {
	if x != nil {
		return x.OwnerId
	}
	return 0
}

func (x *CreateTaskRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateTaskRequest) GetDeadline() *timestamppb.Timestamp {
	if x != nil {
		return x.Deadline
	}
	return nil
}

type CreateTaskResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateTaskResponse) Reset() {
	*x = CreateTaskResponse{}
	mi := &file_task_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTaskResponse) ProtoMessage() {}

func (x *CreateTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_task_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTaskResponse.ProtoReflect.Descriptor instead.
func (*CreateTaskResponse) Descriptor() ([]byte, []int) {
	return file_task_proto_rawDescGZIP(), []int{1}
}

func (x *CreateTaskResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GetTasksRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OwnerId       int32                  `protobuf:"varint,1,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetTasksRequest) Reset() {
	*x = GetTasksRequest{}
	mi := &file_task_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTasksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTasksRequest) ProtoMessage() {}

func (x *GetTasksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_task_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTasksRequest.ProtoReflect.Descriptor instead.
func (*GetTasksRequest) Descriptor() ([]byte, []int) {
	return file_task_proto_rawDescGZIP(), []int{2}
}

func (x *GetTasksRequest) GetOwnerId() int32 {
	if x != nil {
		return x.OwnerId
	}
	return 0
}

type GetTasksResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Tasks         []*Task                `protobuf:"bytes,1,rep,name=tasks,proto3" json:"tasks,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetTasksResponse) Reset() {
	*x = GetTasksResponse{}
	mi := &file_task_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTasksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTasksResponse) ProtoMessage() {}

func (x *GetTasksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_task_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTasksResponse.ProtoReflect.Descriptor instead.
func (*GetTasksResponse) Descriptor() ([]byte, []int) {
	return file_task_proto_rawDescGZIP(), []int{3}
}

func (x *GetTasksResponse) GetTasks() []*Task {
	if x != nil {
		return x.Tasks
	}
	return nil
}

func (x *GetTasksResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Task struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	OwnerId       int32                  `protobuf:"varint,2,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	Title         string                 `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Deadline      *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=deadline,proto3" json:"deadline,omitempty"`
	Status        string                 `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Task) Reset() {
	*x = Task{}
	mi := &file_task_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Task) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task) ProtoMessage() {}

func (x *Task) ProtoReflect() protoreflect.Message {
	mi := &file_task_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Task.ProtoReflect.Descriptor instead.
func (*Task) Descriptor() ([]byte, []int) {
	return file_task_proto_rawDescGZIP(), []int{4}
}

func (x *Task) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Task) GetOwnerId() int32 {
	if x != nil {
		return x.OwnerId
	}
	return 0
}

func (x *Task) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Task) GetDeadline() *timestamppb.Timestamp {
	if x != nil {
		return x.Deadline
	}
	return nil
}

func (x *Task) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type CompleteTaskRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	OwnerId       int32                  `protobuf:"varint,2,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CompleteTaskRequest) Reset() {
	*x = CompleteTaskRequest{}
	mi := &file_task_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CompleteTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompleteTaskRequest) ProtoMessage() {}

func (x *CompleteTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_task_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompleteTaskRequest.ProtoReflect.Descriptor instead.
func (*CompleteTaskRequest) Descriptor() ([]byte, []int) {
	return file_task_proto_rawDescGZIP(), []int{5}
}

func (x *CompleteTaskRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CompleteTaskRequest) GetOwnerId() int32 {
	if x != nil {
		return x.OwnerId
	}
	return 0
}

type CompleteTaskResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CompleteTaskResponse) Reset() {
	*x = CompleteTaskResponse{}
	mi := &file_task_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CompleteTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompleteTaskResponse) ProtoMessage() {}

func (x *CompleteTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_task_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompleteTaskResponse.ProtoReflect.Descriptor instead.
func (*CompleteTaskResponse) Descriptor() ([]byte, []int) {
	return file_task_proto_rawDescGZIP(), []int{6}
}

func (x *CompleteTaskResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type DeleteTaskRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	OwnerId       int32                  `protobuf:"varint,2,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteTaskRequest) Reset() {
	*x = DeleteTaskRequest{}
	mi := &file_task_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTaskRequest) ProtoMessage() {}

func (x *DeleteTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_task_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTaskRequest.ProtoReflect.Descriptor instead.
func (*DeleteTaskRequest) Descriptor() ([]byte, []int) {
	return file_task_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteTaskRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DeleteTaskRequest) GetOwnerId() int32 {
	if x != nil {
		return x.OwnerId
	}
	return 0
}

type DeleteTaskResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteTaskResponse) Reset() {
	*x = DeleteTaskResponse{}
	mi := &file_task_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTaskResponse) ProtoMessage() {}

func (x *DeleteTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_task_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTaskResponse.ProtoReflect.Descriptor instead.
func (*DeleteTaskResponse) Descriptor() ([]byte, []int) {
	return file_task_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteTaskResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_task_proto protoreflect.FileDescriptor

const file_task_proto_rawDesc = "" +
	"\n" +
	"\n" +
	"task.proto\x12\x03api\x1a\x1fgoogle/protobuf/timestamp.proto\"|\n" +
	"\x11CreateTaskRequest\x12\x19\n" +
	"\bowner_id\x18\x01 \x01(\x05R\aownerId\x12\x14\n" +
	"\x05title\x18\x02 \x01(\tR\x05title\x126\n" +
	"\bdeadline\x18\x03 \x01(\v2\x1a.google.protobuf.TimestampR\bdeadline\".\n" +
	"\x12CreateTaskResponse\x12\x18\n" +
	"\amessage\x18\x01 \x01(\tR\amessage\",\n" +
	"\x0fGetTasksRequest\x12\x19\n" +
	"\bowner_id\x18\x01 \x01(\x05R\aownerId\"M\n" +
	"\x10GetTasksResponse\x12\x1f\n" +
	"\x05tasks\x18\x01 \x03(\v2\t.api.TaskR\x05tasks\x12\x18\n" +
	"\amessage\x18\x02 \x01(\tR\amessage\"\x97\x01\n" +
	"\x04Task\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\x12\x19\n" +
	"\bowner_id\x18\x02 \x01(\x05R\aownerId\x12\x14\n" +
	"\x05title\x18\x03 \x01(\tR\x05title\x126\n" +
	"\bdeadline\x18\x04 \x01(\v2\x1a.google.protobuf.TimestampR\bdeadline\x12\x16\n" +
	"\x06status\x18\x05 \x01(\tR\x06status\"@\n" +
	"\x13CompleteTaskRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\x12\x19\n" +
	"\bowner_id\x18\x02 \x01(\x05R\aownerId\"0\n" +
	"\x14CompleteTaskResponse\x12\x18\n" +
	"\amessage\x18\x01 \x01(\tR\amessage\">\n" +
	"\x11DeleteTaskRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\x12\x19\n" +
	"\bowner_id\x18\x02 \x01(\x05R\aownerId\".\n" +
	"\x12DeleteTaskResponse\x12\x18\n" +
	"\amessage\x18\x01 \x01(\tR\amessage2\x91\x02\n" +
	"\vTaskService\x12?\n" +
	"\n" +
	"CreateTask\x12\x16.api.CreateTaskRequest\x1a\x17.api.CreateTaskResponse\"\x00\x129\n" +
	"\bGetTasks\x12\x14.api.GetTasksRequest\x1a\x15.api.GetTasksResponse\"\x00\x12E\n" +
	"\fCompleteTask\x12\x18.api.CompleteTaskRequest\x1a\x19.api.CompleteTaskResponse\"\x00\x12?\n" +
	"\n" +
	"DeleteTask\x12\x16.api.DeleteTaskRequest\x1a\x17.api.DeleteTaskResponse\"\x00B6Z4github.com/kianyari/microservice-practice/common/apib\x06proto3"

var (
	file_task_proto_rawDescOnce sync.Once
	file_task_proto_rawDescData []byte
)

func file_task_proto_rawDescGZIP() []byte {
	file_task_proto_rawDescOnce.Do(func() {
		file_task_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_task_proto_rawDesc), len(file_task_proto_rawDesc)))
	})
	return file_task_proto_rawDescData
}

var file_task_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_task_proto_goTypes = []any{
	(*CreateTaskRequest)(nil),     // 0: api.CreateTaskRequest
	(*CreateTaskResponse)(nil),    // 1: api.CreateTaskResponse
	(*GetTasksRequest)(nil),       // 2: api.GetTasksRequest
	(*GetTasksResponse)(nil),      // 3: api.GetTasksResponse
	(*Task)(nil),                  // 4: api.Task
	(*CompleteTaskRequest)(nil),   // 5: api.CompleteTaskRequest
	(*CompleteTaskResponse)(nil),  // 6: api.CompleteTaskResponse
	(*DeleteTaskRequest)(nil),     // 7: api.DeleteTaskRequest
	(*DeleteTaskResponse)(nil),    // 8: api.DeleteTaskResponse
	(*timestamppb.Timestamp)(nil), // 9: google.protobuf.Timestamp
}
var file_task_proto_depIdxs = []int32{
	9, // 0: api.CreateTaskRequest.deadline:type_name -> google.protobuf.Timestamp
	4, // 1: api.GetTasksResponse.tasks:type_name -> api.Task
	9, // 2: api.Task.deadline:type_name -> google.protobuf.Timestamp
	0, // 3: api.TaskService.CreateTask:input_type -> api.CreateTaskRequest
	2, // 4: api.TaskService.GetTasks:input_type -> api.GetTasksRequest
	5, // 5: api.TaskService.CompleteTask:input_type -> api.CompleteTaskRequest
	7, // 6: api.TaskService.DeleteTask:input_type -> api.DeleteTaskRequest
	1, // 7: api.TaskService.CreateTask:output_type -> api.CreateTaskResponse
	3, // 8: api.TaskService.GetTasks:output_type -> api.GetTasksResponse
	6, // 9: api.TaskService.CompleteTask:output_type -> api.CompleteTaskResponse
	8, // 10: api.TaskService.DeleteTask:output_type -> api.DeleteTaskResponse
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_task_proto_init() }
func file_task_proto_init() {
	if File_task_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_task_proto_rawDesc), len(file_task_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_task_proto_goTypes,
		DependencyIndexes: file_task_proto_depIdxs,
		MessageInfos:      file_task_proto_msgTypes,
	}.Build()
	File_task_proto = out.File
	file_task_proto_goTypes = nil
	file_task_proto_depIdxs = nil
}

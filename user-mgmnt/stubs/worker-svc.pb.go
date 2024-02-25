// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.3
// source: proto/svc/worker-svc.proto

package stubs

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

type WorkerManagementRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	WorkerData *WorkerData `protobuf:"bytes,2,opt,name=workerData,proto3" json:"workerData,omitempty"`
}

func (x *WorkerManagementRequest) Reset() {
	*x = WorkerManagementRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_svc_worker_svc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkerManagementRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkerManagementRequest) ProtoMessage() {}

func (x *WorkerManagementRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_svc_worker_svc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkerManagementRequest.ProtoReflect.Descriptor instead.
func (*WorkerManagementRequest) Descriptor() ([]byte, []int) {
	return file_proto_svc_worker_svc_proto_rawDescGZIP(), []int{0}
}

func (x *WorkerManagementRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *WorkerManagementRequest) GetWorkerData() *WorkerData {
	if x != nil {
		return x.WorkerData
	}
	return nil
}

type WorkerManagementResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string        `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Status     string        `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	WorkerData []*WorkerData `protobuf:"bytes,3,rep,name=workerData,proto3" json:"workerData,omitempty"`
}

func (x *WorkerManagementResponse) Reset() {
	*x = WorkerManagementResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_svc_worker_svc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkerManagementResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkerManagementResponse) ProtoMessage() {}

func (x *WorkerManagementResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_svc_worker_svc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkerManagementResponse.ProtoReflect.Descriptor instead.
func (*WorkerManagementResponse) Descriptor() ([]byte, []int) {
	return file_proto_svc_worker_svc_proto_rawDescGZIP(), []int{1}
}

func (x *WorkerManagementResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *WorkerManagementResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *WorkerManagementResponse) GetWorkerData() []*WorkerData {
	if x != nil {
		return x.WorkerData
	}
	return nil
}

var File_proto_svc_worker_svc_proto protoreflect.FileDescriptor

var file_proto_svc_worker_svc_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x76, 0x63, 0x2f, 0x77, 0x6f, 0x72, 0x6b,
	0x65, 0x72, 0x2d, 0x73, 0x76, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x73, 0x76,
	0x63, 0x1a, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f,
	0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2d, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x5d, 0x0a, 0x17, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x4d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x32, 0x0a, 0x0a,
	0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61,
	0x22, 0x76, 0x0a, 0x18, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x32, 0x0a, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x44, 0x61,
	0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x73, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0a, 0x77, 0x6f,
	0x72, 0x6b, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x32, 0xb8, 0x03, 0x0a, 0x10, 0x57, 0x6f, 0x72,
	0x6b, 0x65, 0x72, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x54, 0x0a,
	0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1c, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x65,
	0x72, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x4d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x54, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x57, 0x6f, 0x72,
	0x6b, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1c, 0x2e, 0x73, 0x76, 0x63,
	0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x57,
	0x6f, 0x72, 0x6b, 0x65, 0x72, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x55, 0x0a, 0x14, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x73, 0x12, 0x1c, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x4d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1d, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x4d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x4e, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x42, 0x79, 0x49,
	0x64, 0x12, 0x1c, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x4d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1d, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x4d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x51, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72,
	0x42, 0x79, 0x49, 0x64, 0x12, 0x1c, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x65,
	0x72, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x4d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x08, 0x5a, 0x06, 0x2f, 0x73, 0x74, 0x75, 0x62, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_svc_worker_svc_proto_rawDescOnce sync.Once
	file_proto_svc_worker_svc_proto_rawDescData = file_proto_svc_worker_svc_proto_rawDesc
)

func file_proto_svc_worker_svc_proto_rawDescGZIP() []byte {
	file_proto_svc_worker_svc_proto_rawDescOnce.Do(func() {
		file_proto_svc_worker_svc_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_svc_worker_svc_proto_rawDescData)
	})
	return file_proto_svc_worker_svc_proto_rawDescData
}

var file_proto_svc_worker_svc_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_svc_worker_svc_proto_goTypes = []interface{}{
	(*WorkerManagementRequest)(nil),  // 0: svc.WorkerManagementRequest
	(*WorkerManagementResponse)(nil), // 1: svc.WorkerManagementResponse
	(*WorkerData)(nil),               // 2: models.WorkerData
}
var file_proto_svc_worker_svc_proto_depIdxs = []int32{
	2, // 0: svc.WorkerManagementRequest.workerData:type_name -> models.WorkerData
	2, // 1: svc.WorkerManagementResponse.workerData:type_name -> models.WorkerData
	0, // 2: svc.WorkerManagement.CreateWorkerAccount:input_type -> svc.WorkerManagementRequest
	0, // 3: svc.WorkerManagement.DeleteWorkerAccount:input_type -> svc.WorkerManagementRequest
	0, // 4: svc.WorkerManagement.GetAllWorkerAccounts:input_type -> svc.WorkerManagementRequest
	0, // 5: svc.WorkerManagement.GetWorkerById:input_type -> svc.WorkerManagementRequest
	0, // 6: svc.WorkerManagement.UpdateWorkerById:input_type -> svc.WorkerManagementRequest
	1, // 7: svc.WorkerManagement.CreateWorkerAccount:output_type -> svc.WorkerManagementResponse
	1, // 8: svc.WorkerManagement.DeleteWorkerAccount:output_type -> svc.WorkerManagementResponse
	1, // 9: svc.WorkerManagement.GetAllWorkerAccounts:output_type -> svc.WorkerManagementResponse
	1, // 10: svc.WorkerManagement.GetWorkerById:output_type -> svc.WorkerManagementResponse
	1, // 11: svc.WorkerManagement.UpdateWorkerById:output_type -> svc.WorkerManagementResponse
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_svc_worker_svc_proto_init() }
func file_proto_svc_worker_svc_proto_init() {
	if File_proto_svc_worker_svc_proto != nil {
		return
	}
	file_proto_models_worker_info_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_svc_worker_svc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkerManagementRequest); i {
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
		file_proto_svc_worker_svc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkerManagementResponse); i {
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
			RawDescriptor: file_proto_svc_worker_svc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_svc_worker_svc_proto_goTypes,
		DependencyIndexes: file_proto_svc_worker_svc_proto_depIdxs,
		MessageInfos:      file_proto_svc_worker_svc_proto_msgTypes,
	}.Build()
	File_proto_svc_worker_svc_proto = out.File
	file_proto_svc_worker_svc_proto_rawDesc = nil
	file_proto_svc_worker_svc_proto_goTypes = nil
	file_proto_svc_worker_svc_proto_depIdxs = nil
}

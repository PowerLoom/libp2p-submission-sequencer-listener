// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.2
// source: proto/submission.proto

package pkgs

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

// Request structure as defined in your Solidity contract
type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SlotId      uint64 `protobuf:"varint,1,opt,name=slotId,proto3" json:"slotId,omitempty"`
	Deadline    uint64 `protobuf:"varint,2,opt,name=deadline,proto3" json:"deadline,omitempty"`
	SnapshotCid string `protobuf:"bytes,3,opt,name=snapshotCid,proto3" json:"snapshotCid,omitempty"`
	EpochId     uint64 `protobuf:"varint,4,opt,name=epochId,proto3" json:"epochId,omitempty"`
	ProjectId   string `protobuf:"bytes,5,opt,name=projectId,proto3" json:"projectId,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_submission_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_proto_submission_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_proto_submission_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetSlotId() uint64 {
	if x != nil {
		return x.SlotId
	}
	return 0
}

func (x *Request) GetDeadline() uint64 {
	if x != nil {
		return x.Deadline
	}
	return 0
}

func (x *Request) GetSnapshotCid() string {
	if x != nil {
		return x.SnapshotCid
	}
	return ""
}

func (x *Request) GetEpochId() uint64 {
	if x != nil {
		return x.EpochId
	}
	return 0
}

func (x *Request) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

type SnapshotSubmission struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request    *Request `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
	Signature  string   `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	Header     string   `protobuf:"bytes,3,opt,name=header,proto3" json:"header,omitempty"`
	DataMarket string   `protobuf:"bytes,4,opt,name=dataMarket,proto3" json:"dataMarket,omitempty"`
}

func (x *SnapshotSubmission) Reset() {
	*x = SnapshotSubmission{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_submission_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SnapshotSubmission) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SnapshotSubmission) ProtoMessage() {}

func (x *SnapshotSubmission) ProtoReflect() protoreflect.Message {
	mi := &file_proto_submission_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SnapshotSubmission.ProtoReflect.Descriptor instead.
func (*SnapshotSubmission) Descriptor() ([]byte, []int) {
	return file_proto_submission_proto_rawDescGZIP(), []int{1}
}

func (x *SnapshotSubmission) GetRequest() *Request {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *SnapshotSubmission) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

func (x *SnapshotSubmission) GetHeader() string {
	if x != nil {
		return x.Header
	}
	return ""
}

func (x *SnapshotSubmission) GetDataMarket() string {
	if x != nil {
		return x.DataMarket
	}
	return ""
}

type SubmissionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"` // Response message
}

func (x *SubmissionResponse) Reset() {
	*x = SubmissionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_submission_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmissionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmissionResponse) ProtoMessage() {}

func (x *SubmissionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_submission_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmissionResponse.ProtoReflect.Descriptor instead.
func (*SubmissionResponse) Descriptor() ([]byte, []int) {
	return file_proto_submission_proto_rawDescGZIP(), []int{2}
}

func (x *SubmissionResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_proto_submission_proto protoreflect.FileDescriptor

var file_proto_submission_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x22, 0x97, 0x01, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x6c, 0x6f, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x06, 0x73, 0x6c, 0x6f, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x61, 0x64,
	0x6c, 0x69, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x64, 0x65, 0x61, 0x64,
	0x6c, 0x69, 0x6e, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74,
	0x43, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x6e, 0x61, 0x70, 0x73,
	0x68, 0x6f, 0x74, 0x43, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x49,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x49, 0x64,
	0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x22, 0x99,
	0x01, 0x0a, 0x12, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x53, 0x75, 0x62, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2d, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x07, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x61,
	0x74, 0x61, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x64, 0x61, 0x74, 0x61, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x22, 0x2e, 0x0a, 0x12, 0x53, 0x75,
	0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xbc, 0x01, 0x0a, 0x0a, 0x53,
	0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x5c, 0x0a, 0x18, 0x53, 0x75, 0x62,
	0x6d, 0x69, 0x74, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x53, 0x69, 0x6d, 0x75, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x2e, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x2e, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x53, 0x75, 0x62, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x1a, 0x1e, 0x2e, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x12, 0x50, 0x0a, 0x0e, 0x53, 0x75, 0x62, 0x6d, 0x69,
	0x74, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x12, 0x1e, 0x2e, 0x73, 0x75, 0x62, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x53,
	0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x1a, 0x1e, 0x2e, 0x73, 0x75, 0x62, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x4c, 0x6f, 0x6f,
	0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74,
	0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_submission_proto_rawDescOnce sync.Once
	file_proto_submission_proto_rawDescData = file_proto_submission_proto_rawDesc
)

func file_proto_submission_proto_rawDescGZIP() []byte {
	file_proto_submission_proto_rawDescOnce.Do(func() {
		file_proto_submission_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_submission_proto_rawDescData)
	})
	return file_proto_submission_proto_rawDescData
}

var file_proto_submission_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_submission_proto_goTypes = []any{
	(*Request)(nil),            // 0: submission.Request
	(*SnapshotSubmission)(nil), // 1: submission.SnapshotSubmission
	(*SubmissionResponse)(nil), // 2: submission.SubmissionResponse
}
var file_proto_submission_proto_depIdxs = []int32{
	0, // 0: submission.SnapshotSubmission.request:type_name -> submission.Request
	1, // 1: submission.Submission.SubmitSnapshotSimulation:input_type -> submission.SnapshotSubmission
	1, // 2: submission.Submission.SubmitSnapshot:input_type -> submission.SnapshotSubmission
	2, // 3: submission.Submission.SubmitSnapshotSimulation:output_type -> submission.SubmissionResponse
	2, // 4: submission.Submission.SubmitSnapshot:output_type -> submission.SubmissionResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_submission_proto_init() }
func file_proto_submission_proto_init() {
	if File_proto_submission_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_submission_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Request); i {
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
		file_proto_submission_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*SnapshotSubmission); i {
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
		file_proto_submission_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*SubmissionResponse); i {
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
			RawDescriptor: file_proto_submission_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_submission_proto_goTypes,
		DependencyIndexes: file_proto_submission_proto_depIdxs,
		MessageInfos:      file_proto_submission_proto_msgTypes,
	}.Build()
	File_proto_submission_proto = out.File
	file_proto_submission_proto_rawDesc = nil
	file_proto_submission_proto_goTypes = nil
	file_proto_submission_proto_depIdxs = nil
}

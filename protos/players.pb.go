// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: players.proto

package protos

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

type GetPlayerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *GetPlayerRequest) Reset() {
	*x = GetPlayerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_players_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPlayerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlayerRequest) ProtoMessage() {}

func (x *GetPlayerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_players_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPlayerRequest.ProtoReflect.Descriptor instead.
func (*GetPlayerRequest) Descriptor() ([]byte, []int) {
	return file_players_proto_rawDescGZIP(), []int{0}
}

func (x *GetPlayerRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetPlayerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	TeamId string `protobuf:"bytes,3,opt,name=TeamId,proto3" json:"TeamId,omitempty"`
}

func (x *GetPlayerResponse) Reset() {
	*x = GetPlayerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_players_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPlayerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlayerResponse) ProtoMessage() {}

func (x *GetPlayerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_players_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPlayerResponse.ProtoReflect.Descriptor instead.
func (*GetPlayerResponse) Descriptor() ([]byte, []int) {
	return file_players_proto_rawDescGZIP(), []int{1}
}

func (x *GetPlayerResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetPlayerResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetPlayerResponse) GetTeamId() string {
	if x != nil {
		return x.TeamId
	}
	return ""
}

var File_players_proto protoreflect.FileDescriptor

var file_players_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x22, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x49, 0x64, 0x22, 0x4f, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x54, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x54, 0x65,
	0x61, 0x6d, 0x49, 0x64, 0x32, 0x36, 0x0a, 0x06, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x2c,
	0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x11, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0a, 0x5a, 0x08,
	0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_players_proto_rawDescOnce sync.Once
	file_players_proto_rawDescData = file_players_proto_rawDesc
)

func file_players_proto_rawDescGZIP() []byte {
	file_players_proto_rawDescOnce.Do(func() {
		file_players_proto_rawDescData = protoimpl.X.CompressGZIP(file_players_proto_rawDescData)
	})
	return file_players_proto_rawDescData
}

var file_players_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_players_proto_goTypes = []interface{}{
	(*GetPlayerRequest)(nil),  // 0: GetPlayerRequest
	(*GetPlayerResponse)(nil), // 1: GetPlayerResponse
}
var file_players_proto_depIdxs = []int32{
	0, // 0: Player.Get:input_type -> GetPlayerRequest
	1, // 1: Player.Get:output_type -> GetPlayerResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_players_proto_init() }
func file_players_proto_init() {
	if File_players_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_players_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPlayerRequest); i {
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
		file_players_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPlayerResponse); i {
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
			RawDescriptor: file_players_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_players_proto_goTypes,
		DependencyIndexes: file_players_proto_depIdxs,
		MessageInfos:      file_players_proto_msgTypes,
	}.Build()
	File_players_proto = out.File
	file_players_proto_rawDesc = nil
	file_players_proto_goTypes = nil
	file_players_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        (unknown)
// source: geo/v1/service.proto

package geov1

import (
	v1 "github.com/xabsvoid/tracker/pkg/api/type/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type LocateByUUIDRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Uuid          string                 `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LocateByUUIDRequest) Reset() {
	*x = LocateByUUIDRequest{}
	mi := &file_geo_v1_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LocateByUUIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocateByUUIDRequest) ProtoMessage() {}

func (x *LocateByUUIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_geo_v1_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocateByUUIDRequest.ProtoReflect.Descriptor instead.
func (*LocateByUUIDRequest) Descriptor() ([]byte, []int) {
	return file_geo_v1_service_proto_rawDescGZIP(), []int{0}
}

func (x *LocateByUUIDRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type LocateByUUIDResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Latlng        *v1.LatLng             `protobuf:"bytes,1,opt,name=latlng,proto3" json:"latlng,omitempty"`
	Deadline      *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=deadline,proto3" json:"deadline,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LocateByUUIDResponse) Reset() {
	*x = LocateByUUIDResponse{}
	mi := &file_geo_v1_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LocateByUUIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocateByUUIDResponse) ProtoMessage() {}

func (x *LocateByUUIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_geo_v1_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocateByUUIDResponse.ProtoReflect.Descriptor instead.
func (*LocateByUUIDResponse) Descriptor() ([]byte, []int) {
	return file_geo_v1_service_proto_rawDescGZIP(), []int{1}
}

func (x *LocateByUUIDResponse) GetLatlng() *v1.LatLng {
	if x != nil {
		return x.Latlng
	}
	return nil
}

func (x *LocateByUUIDResponse) GetDeadline() *timestamppb.Timestamp {
	if x != nil {
		return x.Deadline
	}
	return nil
}

type TrackRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Uuid          string                 `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Latlng        *v1.LatLng             `protobuf:"bytes,2,opt,name=latlng,proto3" json:"latlng,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TrackRequest) Reset() {
	*x = TrackRequest{}
	mi := &file_geo_v1_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TrackRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrackRequest) ProtoMessage() {}

func (x *TrackRequest) ProtoReflect() protoreflect.Message {
	mi := &file_geo_v1_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrackRequest.ProtoReflect.Descriptor instead.
func (*TrackRequest) Descriptor() ([]byte, []int) {
	return file_geo_v1_service_proto_rawDescGZIP(), []int{2}
}

func (x *TrackRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *TrackRequest) GetLatlng() *v1.LatLng {
	if x != nil {
		return x.Latlng
	}
	return nil
}

type TrackResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TrackResponse) Reset() {
	*x = TrackResponse{}
	mi := &file_geo_v1_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TrackResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrackResponse) ProtoMessage() {}

func (x *TrackResponse) ProtoReflect() protoreflect.Message {
	mi := &file_geo_v1_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrackResponse.ProtoReflect.Descriptor instead.
func (*TrackResponse) Descriptor() ([]byte, []int) {
	return file_geo_v1_service_proto_rawDescGZIP(), []int{3}
}

var File_geo_v1_service_proto protoreflect.FileDescriptor

var file_geo_v1_service_proto_rawDesc = []byte{
	0x0a, 0x14, 0x67, 0x65, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x67, 0x65, 0x6f, 0x2e, 0x76, 0x31, 0x1a, 0x14,
	0x74, 0x79, 0x70, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x61, 0x74, 0x6c, 0x6e, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x29, 0x0a, 0x13, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x42,
	0x79, 0x55, 0x55, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64,
	0x22, 0x77, 0x0a, 0x14, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x42, 0x79, 0x55, 0x55, 0x49, 0x44,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x06, 0x6c, 0x61, 0x74, 0x6c,
	0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x61, 0x74, 0x4c, 0x6e, 0x67, 0x52, 0x06, 0x6c, 0x61, 0x74, 0x6c, 0x6e,
	0x67, 0x12, 0x36, 0x0a, 0x08, 0x64, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x08, 0x64, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x22, 0x4b, 0x0a, 0x0c, 0x54, 0x72, 0x61,
	0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x27, 0x0a,
	0x06, 0x6c, 0x61, 0x74, 0x6c, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x61, 0x74, 0x4c, 0x6e, 0x67, 0x52, 0x06,
	0x6c, 0x61, 0x74, 0x6c, 0x6e, 0x67, 0x22, 0x0f, 0x0a, 0x0d, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x91, 0x01, 0x0a, 0x0a, 0x47, 0x65, 0x6f, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4b, 0x0a, 0x0c, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x65,
	0x42, 0x79, 0x55, 0x55, 0x49, 0x44, 0x12, 0x1b, 0x2e, 0x67, 0x65, 0x6f, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x42, 0x79, 0x55, 0x55, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x67, 0x65, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x65, 0x42, 0x79, 0x55, 0x55, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x05, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x12, 0x14, 0x2e, 0x67,
	0x65, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x15, 0x2e, 0x67, 0x65, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x63,
	0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x85, 0x01, 0x0a, 0x0a,
	0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x65, 0x6f, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x61, 0x62, 0x73, 0x76, 0x6f, 0x69, 0x64, 0x2f,
	0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x67, 0x65, 0x6f, 0x2f, 0x76, 0x31, 0x3b, 0x67, 0x65, 0x6f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x47,
	0x58, 0x58, 0xaa, 0x02, 0x06, 0x47, 0x65, 0x6f, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x06, 0x47, 0x65,
	0x6f, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x12, 0x47, 0x65, 0x6f, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x07, 0x47, 0x65, 0x6f, 0x3a,
	0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_geo_v1_service_proto_rawDescOnce sync.Once
	file_geo_v1_service_proto_rawDescData = file_geo_v1_service_proto_rawDesc
)

func file_geo_v1_service_proto_rawDescGZIP() []byte {
	file_geo_v1_service_proto_rawDescOnce.Do(func() {
		file_geo_v1_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_geo_v1_service_proto_rawDescData)
	})
	return file_geo_v1_service_proto_rawDescData
}

var file_geo_v1_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_geo_v1_service_proto_goTypes = []any{
	(*LocateByUUIDRequest)(nil),   // 0: geo.v1.LocateByUUIDRequest
	(*LocateByUUIDResponse)(nil),  // 1: geo.v1.LocateByUUIDResponse
	(*TrackRequest)(nil),          // 2: geo.v1.TrackRequest
	(*TrackResponse)(nil),         // 3: geo.v1.TrackResponse
	(*v1.LatLng)(nil),             // 4: type.v1.LatLng
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
}
var file_geo_v1_service_proto_depIdxs = []int32{
	4, // 0: geo.v1.LocateByUUIDResponse.latlng:type_name -> type.v1.LatLng
	5, // 1: geo.v1.LocateByUUIDResponse.deadline:type_name -> google.protobuf.Timestamp
	4, // 2: geo.v1.TrackRequest.latlng:type_name -> type.v1.LatLng
	0, // 3: geo.v1.GeoService.LocateByUUID:input_type -> geo.v1.LocateByUUIDRequest
	2, // 4: geo.v1.GeoService.Track:input_type -> geo.v1.TrackRequest
	1, // 5: geo.v1.GeoService.LocateByUUID:output_type -> geo.v1.LocateByUUIDResponse
	3, // 6: geo.v1.GeoService.Track:output_type -> geo.v1.TrackResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_geo_v1_service_proto_init() }
func file_geo_v1_service_proto_init() {
	if File_geo_v1_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_geo_v1_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_geo_v1_service_proto_goTypes,
		DependencyIndexes: file_geo_v1_service_proto_depIdxs,
		MessageInfos:      file_geo_v1_service_proto_msgTypes,
	}.Build()
	File_geo_v1_service_proto = out.File
	file_geo_v1_service_proto_rawDesc = nil
	file_geo_v1_service_proto_goTypes = nil
	file_geo_v1_service_proto_depIdxs = nil
}

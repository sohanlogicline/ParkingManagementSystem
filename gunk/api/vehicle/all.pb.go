// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        (unknown)
// source: parking/gunk/api/vehicle/all.proto

package vehicle

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Vehicle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          string                 `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	VehicleNo   string                 `protobuf:"bytes,2,opt,name=VehicleNo,proto3" json:"VehicleNo,omitempty"`
	Entry       *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=Entry,proto3" json:"Entry,omitempty"`
	Exit        *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=Exit,proto3" json:"Exit,omitempty"`
	ParkedHours int32                  `protobuf:"varint,5,opt,name=ParkedHours,proto3" json:"ParkedHours,omitempty"`
	Fee         int32                  `protobuf:"varint,6,opt,name=Fee,proto3" json:"Fee,omitempty"`
	SlotID      string                 `protobuf:"bytes,7,opt,name=SlotID,proto3" json:"SlotID,omitempty"`
	CreatedAt   *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt   *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
}

func (x *Vehicle) Reset() {
	*x = Vehicle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_parking_gunk_api_vehicle_all_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Vehicle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Vehicle) ProtoMessage() {}

func (x *Vehicle) ProtoReflect() protoreflect.Message {
	mi := &file_parking_gunk_api_vehicle_all_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Vehicle.ProtoReflect.Descriptor instead.
func (*Vehicle) Descriptor() ([]byte, []int) {
	return file_parking_gunk_api_vehicle_all_proto_rawDescGZIP(), []int{0}
}

func (x *Vehicle) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Vehicle) GetVehicleNo() string {
	if x != nil {
		return x.VehicleNo
	}
	return ""
}

func (x *Vehicle) GetEntry() *timestamppb.Timestamp {
	if x != nil {
		return x.Entry
	}
	return nil
}

func (x *Vehicle) GetExit() *timestamppb.Timestamp {
	if x != nil {
		return x.Exit
	}
	return nil
}

func (x *Vehicle) GetParkedHours() int32 {
	if x != nil {
		return x.ParkedHours
	}
	return 0
}

func (x *Vehicle) GetFee() int32 {
	if x != nil {
		return x.Fee
	}
	return 0
}

func (x *Vehicle) GetSlotID() string {
	if x != nil {
		return x.SlotID
	}
	return ""
}

func (x *Vehicle) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Vehicle) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type VehicleData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vehicle *Vehicle `protobuf:"bytes,1,opt,name=Vehicle,proto3" json:"Vehicle,omitempty"`
}

func (x *VehicleData) Reset() {
	*x = VehicleData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_parking_gunk_api_vehicle_all_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VehicleData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VehicleData) ProtoMessage() {}

func (x *VehicleData) ProtoReflect() protoreflect.Message {
	mi := &file_parking_gunk_api_vehicle_all_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VehicleData.ProtoReflect.Descriptor instead.
func (*VehicleData) Descriptor() ([]byte, []int) {
	return file_parking_gunk_api_vehicle_all_proto_rawDescGZIP(), []int{1}
}

func (x *VehicleData) GetVehicle() *Vehicle {
	if x != nil {
		return x.Vehicle
	}
	return nil
}

type ParkVehicleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VehicleNo string `protobuf:"bytes,1,opt,name=VehicleNo,proto3" json:"VehicleNo,omitempty"`
	LotID     string `protobuf:"bytes,2,opt,name=LotID,proto3" json:"LotID,omitempty"`
}

func (x *ParkVehicleRequest) Reset() {
	*x = ParkVehicleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_parking_gunk_api_vehicle_all_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParkVehicleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParkVehicleRequest) ProtoMessage() {}

func (x *ParkVehicleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_parking_gunk_api_vehicle_all_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParkVehicleRequest.ProtoReflect.Descriptor instead.
func (*ParkVehicleRequest) Descriptor() ([]byte, []int) {
	return file_parking_gunk_api_vehicle_all_proto_rawDescGZIP(), []int{2}
}

func (x *ParkVehicleRequest) GetVehicleNo() string {
	if x != nil {
		return x.VehicleNo
	}
	return ""
}

func (x *ParkVehicleRequest) GetLotID() string {
	if x != nil {
		return x.LotID
	}
	return ""
}

// ParkVehicleResponse ...
type ParkVehicleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *VehicleData `protobuf:"bytes,1,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (x *ParkVehicleResponse) Reset() {
	*x = ParkVehicleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_parking_gunk_api_vehicle_all_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParkVehicleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParkVehicleResponse) ProtoMessage() {}

func (x *ParkVehicleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_parking_gunk_api_vehicle_all_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParkVehicleResponse.ProtoReflect.Descriptor instead.
func (*ParkVehicleResponse) Descriptor() ([]byte, []int) {
	return file_parking_gunk_api_vehicle_all_proto_rawDescGZIP(), []int{3}
}

func (x *ParkVehicleResponse) GetData() *VehicleData {
	if x != nil {
		return x.Data
	}
	return nil
}

type UnParkVehicleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID        string `protobuf:"bytes,1,opt,name=ID,json=ParkingID,proto3" json:"ParkingID,omitempty"`
	VehicleNo string `protobuf:"bytes,2,opt,name=VehicleNo,proto3" json:"VehicleNo,omitempty"`
}

func (x *UnParkVehicleRequest) Reset() {
	*x = UnParkVehicleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_parking_gunk_api_vehicle_all_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnParkVehicleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnParkVehicleRequest) ProtoMessage() {}

func (x *UnParkVehicleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_parking_gunk_api_vehicle_all_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnParkVehicleRequest.ProtoReflect.Descriptor instead.
func (*UnParkVehicleRequest) Descriptor() ([]byte, []int) {
	return file_parking_gunk_api_vehicle_all_proto_rawDescGZIP(), []int{4}
}

func (x *UnParkVehicleRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *UnParkVehicleRequest) GetVehicleNo() string {
	if x != nil {
		return x.VehicleNo
	}
	return ""
}

// UnParkVehicleResponse ...
type UnParkVehicleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *VehicleData `protobuf:"bytes,1,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (x *UnParkVehicleResponse) Reset() {
	*x = UnParkVehicleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_parking_gunk_api_vehicle_all_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnParkVehicleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnParkVehicleResponse) ProtoMessage() {}

func (x *UnParkVehicleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_parking_gunk_api_vehicle_all_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnParkVehicleResponse.ProtoReflect.Descriptor instead.
func (*UnParkVehicleResponse) Descriptor() ([]byte, []int) {
	return file_parking_gunk_api_vehicle_all_proto_rawDescGZIP(), []int{5}
}

func (x *UnParkVehicleResponse) GetData() *VehicleData {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_parking_gunk_api_vehicle_all_proto protoreflect.FileDescriptor

var file_parking_gunk_api_vehicle_all_proto_rawDesc = []byte{
	0x0a, 0x22, 0x70, 0x61, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x67, 0x75, 0x6e, 0x6b, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x2f, 0x61, 0x6c, 0x6c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61,
	0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfa, 0x03, 0x0a,
	0x07, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x1f, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x0f, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00,
	0x78, 0x00, 0x80, 0x01, 0x00, 0x52, 0x02, 0x49, 0x44, 0x12, 0x2d, 0x0a, 0x09, 0x56, 0x65, 0x68,
	0x69, 0x63, 0x6c, 0x65, 0x4e, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0f, 0x08, 0x00,
	0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x78, 0x00, 0x80, 0x01, 0x00, 0x52, 0x09, 0x56,
	0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x4e, 0x6f, 0x12, 0x41, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x42, 0x0f, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x78,
	0x00, 0x80, 0x01, 0x00, 0x52, 0x05, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x3f, 0x0a, 0x04, 0x45,
	0x78, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x0f, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50,
	0x00, 0x78, 0x00, 0x80, 0x01, 0x00, 0x52, 0x04, 0x45, 0x78, 0x69, 0x74, 0x12, 0x31, 0x0a, 0x0b,
	0x50, 0x61, 0x72, 0x6b, 0x65, 0x64, 0x48, 0x6f, 0x75, 0x72, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x05, 0x42, 0x0f, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x78, 0x00, 0x80,
	0x01, 0x00, 0x52, 0x0b, 0x50, 0x61, 0x72, 0x6b, 0x65, 0x64, 0x48, 0x6f, 0x75, 0x72, 0x73, 0x12,
	0x21, 0x0a, 0x03, 0x46, 0x65, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0f, 0x08, 0x00,
	0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x78, 0x00, 0x80, 0x01, 0x00, 0x52, 0x03, 0x46,
	0x65, 0x65, 0x12, 0x27, 0x0a, 0x06, 0x53, 0x6c, 0x6f, 0x74, 0x49, 0x44, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x0f, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x78, 0x00,
	0x80, 0x01, 0x00, 0x52, 0x06, 0x53, 0x6c, 0x6f, 0x74, 0x49, 0x44, 0x12, 0x49, 0x0a, 0x09, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x0f, 0x08, 0x00, 0x18, 0x00,
	0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x78, 0x00, 0x80, 0x01, 0x00, 0x52, 0x09, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x49, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x0f, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50,
	0x00, 0x78, 0x00, 0x80, 0x01, 0x00, 0x52, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x3a, 0x06, 0x08, 0x00, 0x10, 0x00, 0x18, 0x00, 0x22, 0x52, 0x0a, 0x0b, 0x56, 0x65, 0x68,
	0x69, 0x63, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x3b, 0x0a, 0x07, 0x56, 0x65, 0x68, 0x69,
	0x63, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x76, 0x65, 0x68, 0x69,
	0x63, 0x6c, 0x65, 0x2e, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x42, 0x0f, 0x08, 0x00, 0x18,
	0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x78, 0x00, 0x80, 0x01, 0x00, 0x52, 0x07, 0x56, 0x65,
	0x68, 0x69, 0x63, 0x6c, 0x65, 0x3a, 0x06, 0x08, 0x00, 0x10, 0x00, 0x18, 0x00, 0x22, 0x8b, 0x01,
	0x0a, 0x12, 0x50, 0x61, 0x72, 0x6b, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x09, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x4e,
	0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0f, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30,
	0x00, 0x50, 0x00, 0x78, 0x00, 0x80, 0x01, 0x00, 0x52, 0x09, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c,
	0x65, 0x4e, 0x6f, 0x12, 0x25, 0x0a, 0x05, 0x4c, 0x6f, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x0f, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x78, 0x00,
	0x80, 0x01, 0x00, 0x52, 0x05, 0x4c, 0x6f, 0x74, 0x49, 0x44, 0x3a, 0x1f, 0x08, 0x00, 0x10, 0x00,
	0x18, 0x00, 0x92, 0x41, 0x16, 0x0a, 0x14, 0xd2, 0x01, 0x09, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c,
	0x65, 0x4e, 0x6f, 0xd2, 0x01, 0x05, 0x4c, 0x6f, 0x74, 0x49, 0x44, 0x22, 0x58, 0x0a, 0x13, 0x50,
	0x61, 0x72, 0x6b, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x39, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x56, 0x65, 0x68, 0x69, 0x63,
	0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x42, 0x0f, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00,
	0x50, 0x00, 0x78, 0x00, 0x80, 0x01, 0x00, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x3a, 0x06, 0x08,
	0x00, 0x10, 0x00, 0x18, 0x00, 0x22, 0x75, 0x0a, 0x14, 0x55, 0x6e, 0x50, 0x61, 0x72, 0x6b, 0x56,
	0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a,
	0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0f, 0x08, 0x00, 0x18, 0x00, 0x28,
	0x00, 0x30, 0x00, 0x50, 0x00, 0x78, 0x00, 0x80, 0x01, 0x00, 0x52, 0x09, 0x50, 0x61, 0x72, 0x6b,
	0x69, 0x6e, 0x67, 0x49, 0x44, 0x12, 0x2d, 0x0a, 0x09, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65,
	0x4e, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0f, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00,
	0x30, 0x00, 0x50, 0x00, 0x78, 0x00, 0x80, 0x01, 0x00, 0x52, 0x09, 0x56, 0x65, 0x68, 0x69, 0x63,
	0x6c, 0x65, 0x4e, 0x6f, 0x3a, 0x06, 0x08, 0x00, 0x10, 0x00, 0x18, 0x00, 0x22, 0x5a, 0x0a, 0x15,
	0x55, 0x6e, 0x50, 0x61, 0x72, 0x6b, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x56, 0x65,
	0x68, 0x69, 0x63, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x42, 0x0f, 0x08, 0x00, 0x18, 0x00, 0x28,
	0x00, 0x30, 0x00, 0x50, 0x00, 0x78, 0x00, 0x80, 0x01, 0x00, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61,
	0x3a, 0x06, 0x08, 0x00, 0x10, 0x00, 0x18, 0x00, 0x32, 0x93, 0x0b, 0x0a, 0x0e, 0x56, 0x65, 0x68,
	0x69, 0x63, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xb3, 0x05, 0x0a, 0x0b,
	0x50, 0x61, 0x72, 0x6b, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x1b, 0x2e, 0x76, 0x65,
	0x68, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x50, 0x61, 0x72, 0x6b, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x76, 0x65, 0x68, 0x69, 0x63,
	0x6c, 0x65, 0x2e, 0x50, 0x61, 0x72, 0x6b, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xe4, 0x04, 0x88, 0x02, 0x00, 0x90, 0x02, 0x00, 0x92,
	0x41, 0xca, 0x04, 0x0a, 0x0c, 0x50, 0x61, 0x72, 0x6b, 0x20, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c,
	0x65, 0x12, 0x0c, 0x50, 0x61, 0x72, 0x6b, 0x20, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x1a,
	0x14, 0x50, 0x61, 0x72, 0x6b, 0x20, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x20, 0x72, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x2e, 0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x4a, 0x55, 0x0a, 0x03, 0x32, 0x30, 0x30, 0x12, 0x4e,
	0x0a, 0x1e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x20, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x65, 0x64, 0x20, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x6c, 0x79, 0x2e,
	0x12, 0x2c, 0x0a, 0x2a, 0x1a, 0x28, 0x23, 0x2f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x50, 0x61, 0x72, 0x6b, 0x56,
	0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4a, 0x7e,
	0x0a, 0x03, 0x34, 0x30, 0x30, 0x12, 0x77, 0x0a, 0x30, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x65,
	0x64, 0x20, 0x77, 0x68, 0x65, 0x6e, 0x20, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x20, 0x69,
	0x73, 0x20, 0x69, 0x6e, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x63, 0x74, 0x20, 0x6f, 0x72, 0x20, 0x6d,
	0x61, 0x6c, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x64, 0x2e, 0x12, 0x43, 0x0a, 0x41, 0x4a, 0x3f, 0x7b,
	0x20, 0x22, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x3a, 0x20, 0x34, 0x30, 0x30, 0x2c, 0x20, 0x22, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x3a, 0x20, 0x22, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x20, 0x69, 0x73, 0x20, 0x69, 0x6e, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x63, 0x74, 0x20, 0x6f,
	0x72, 0x20, 0x6d, 0x61, 0x6c, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x64, 0x22, 0x20, 0x7d, 0x4a, 0x6f,
	0x0a, 0x03, 0x34, 0x30, 0x31, 0x12, 0x68, 0x0a, 0x34, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x65,
	0x64, 0x20, 0x77, 0x68, 0x65, 0x6e, 0x20, 0x6e, 0x6f, 0x74, 0x20, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x7a, 0x65, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x70, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d,
	0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x12, 0x30, 0x0a,
	0x2e, 0x4a, 0x2c, 0x7b, 0x20, 0x22, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x3a, 0x20, 0x34, 0x30, 0x31,
	0x2c, 0x20, 0x22, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x3a, 0x20, 0x22, 0x6e, 0x6f,
	0x74, 0x20, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x22, 0x20, 0x7d, 0x4a,
	0x5b, 0x0a, 0x03, 0x34, 0x30, 0x34, 0x12, 0x54, 0x0a, 0x18, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e,
	0x65, 0x64, 0x20, 0x77, 0x68, 0x65, 0x6e, 0x20, 0x6e, 0x6f, 0x74, 0x20, 0x66, 0x6f, 0x75, 0x6e,
	0x64, 0x2e, 0x12, 0x38, 0x0a, 0x36, 0x4a, 0x34, 0x7b, 0x20, 0x22, 0x63, 0x6f, 0x64, 0x65, 0x22,
	0x3a, 0x20, 0x34, 0x30, 0x34, 0x2c, 0x20, 0x22, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x3a, 0x20, 0x22, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x70, 0x61, 0x72,
	0x6b, 0x20, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x22, 0x20, 0x7d, 0x4a, 0x5f, 0x0a, 0x03,
	0x35, 0x30, 0x30, 0x12, 0x58, 0x0a, 0x24, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x65, 0x64, 0x20,
	0x77, 0x68, 0x65, 0x6e, 0x20, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x20, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x20, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x12, 0x30, 0x0a, 0x2e, 0x4a,
	0x2c, 0x7b, 0x20, 0x22, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x3a, 0x20, 0x35, 0x30, 0x30, 0x2c, 0x20,
	0x22, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x3a, 0x20, 0x22, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x20, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x20, 0x7d, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x0a, 0x3a, 0x01, 0x2a, 0x22, 0x05, 0x2f, 0x70, 0x61, 0x72, 0x6b, 0x28, 0x00, 0x30,
	0x00, 0x12, 0xc5, 0x05, 0x0a, 0x0d, 0x55, 0x6e, 0x50, 0x61, 0x72, 0x6b, 0x56, 0x65, 0x68, 0x69,
	0x63, 0x6c, 0x65, 0x12, 0x1d, 0x2e, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x55, 0x6e,
	0x50, 0x61, 0x72, 0x6b, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x55, 0x6e, 0x50,
	0x61, 0x72, 0x6b, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0xf0, 0x04, 0x88, 0x02, 0x00, 0x90, 0x02, 0x00, 0x92, 0x41, 0xd4, 0x04, 0x0a,
	0x0e, 0x55, 0x6e, 0x50, 0x61, 0x72, 0x6b, 0x20, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x12,
	0x0e, 0x55, 0x6e, 0x50, 0x61, 0x72, 0x6b, 0x20, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x1a,
	0x16, 0x55, 0x6e, 0x50, 0x61, 0x72, 0x6b, 0x20, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x20,
	0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x4a, 0x57, 0x0a, 0x03, 0x32, 0x30, 0x30,
	0x12, 0x50, 0x0a, 0x1e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x20, 0x65, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x65, 0x64, 0x20, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x6c,
	0x79, 0x2e, 0x12, 0x2e, 0x0a, 0x2c, 0x1a, 0x2a, 0x23, 0x2f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x55, 0x6e, 0x50,
	0x61, 0x72, 0x6b, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x4a, 0x7e, 0x0a, 0x03, 0x34, 0x30, 0x30, 0x12, 0x77, 0x0a, 0x30, 0x52, 0x65, 0x74,
	0x75, 0x72, 0x6e, 0x65, 0x64, 0x20, 0x77, 0x68, 0x65, 0x6e, 0x20, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x20, 0x69, 0x73, 0x20, 0x69, 0x6e, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x63, 0x74, 0x20,
	0x6f, 0x72, 0x20, 0x6d, 0x61, 0x6c, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x64, 0x2e, 0x12, 0x43, 0x0a,
	0x41, 0x4a, 0x3f, 0x7b, 0x20, 0x22, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x3a, 0x20, 0x34, 0x30, 0x30,
	0x2c, 0x20, 0x22, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x3a, 0x20, 0x22, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x20, 0x69, 0x73, 0x20, 0x69, 0x6e, 0x63, 0x6f, 0x72, 0x72, 0x65,
	0x63, 0x74, 0x20, 0x6f, 0x72, 0x20, 0x6d, 0x61, 0x6c, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x64, 0x22,
	0x20, 0x7d, 0x4a, 0x6f, 0x0a, 0x03, 0x34, 0x30, 0x31, 0x12, 0x68, 0x0a, 0x34, 0x52, 0x65, 0x74,
	0x75, 0x72, 0x6e, 0x65, 0x64, 0x20, 0x77, 0x68, 0x65, 0x6e, 0x20, 0x6e, 0x6f, 0x74, 0x20, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x70, 0x65, 0x72,
	0x66, 0x6f, 0x72, 0x6d, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x12, 0x30, 0x0a, 0x2e, 0x4a, 0x2c, 0x7b, 0x20, 0x22, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x3a,
	0x20, 0x34, 0x30, 0x31, 0x2c, 0x20, 0x22, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x3a,
	0x20, 0x22, 0x6e, 0x6f, 0x74, 0x20, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64,
	0x22, 0x20, 0x7d, 0x4a, 0x5d, 0x0a, 0x03, 0x34, 0x30, 0x34, 0x12, 0x56, 0x0a, 0x18, 0x52, 0x65,
	0x74, 0x75, 0x72, 0x6e, 0x65, 0x64, 0x20, 0x77, 0x68, 0x65, 0x6e, 0x20, 0x6e, 0x6f, 0x74, 0x20,
	0x66, 0x6f, 0x75, 0x6e, 0x64, 0x2e, 0x12, 0x3a, 0x0a, 0x38, 0x4a, 0x36, 0x7b, 0x20, 0x22, 0x63,
	0x6f, 0x64, 0x65, 0x22, 0x3a, 0x20, 0x34, 0x30, 0x34, 0x2c, 0x20, 0x22, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x3a, 0x20, 0x22, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x20, 0x74, 0x6f,
	0x20, 0x75, 0x6e, 0x70, 0x61, 0x72, 0x6b, 0x20, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x22,
	0x20, 0x7d, 0x4a, 0x5f, 0x0a, 0x03, 0x35, 0x30, 0x30, 0x12, 0x58, 0x0a, 0x24, 0x52, 0x65, 0x74,
	0x75, 0x72, 0x6e, 0x65, 0x64, 0x20, 0x77, 0x68, 0x65, 0x6e, 0x20, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x20, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x20, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x2e, 0x12, 0x30, 0x0a, 0x2e, 0x4a, 0x2c, 0x7b, 0x20, 0x22, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x3a,
	0x20, 0x35, 0x30, 0x30, 0x2c, 0x20, 0x22, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x3a,
	0x20, 0x22, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x20, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x22, 0x20, 0x7d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x3a, 0x01, 0x2a, 0x22, 0x07, 0x2f, 0x75,
	0x6e, 0x70, 0x61, 0x72, 0x6b, 0x28, 0x00, 0x30, 0x00, 0x1a, 0x03, 0x88, 0x02, 0x00, 0x42, 0x3b,
	0x48, 0x01, 0x50, 0x00, 0x5a, 0x20, 0x70, 0x61, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x67, 0x75,
	0x6e, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x3b, 0x76,
	0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x80, 0x01, 0x00, 0x88, 0x01, 0x00, 0x90, 0x01, 0x00, 0xb8,
	0x01, 0x00, 0xd8, 0x01, 0x00, 0xf8, 0x01, 0x01, 0xd0, 0x02, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_parking_gunk_api_vehicle_all_proto_rawDescOnce sync.Once
	file_parking_gunk_api_vehicle_all_proto_rawDescData = file_parking_gunk_api_vehicle_all_proto_rawDesc
)

func file_parking_gunk_api_vehicle_all_proto_rawDescGZIP() []byte {
	file_parking_gunk_api_vehicle_all_proto_rawDescOnce.Do(func() {
		file_parking_gunk_api_vehicle_all_proto_rawDescData = protoimpl.X.CompressGZIP(file_parking_gunk_api_vehicle_all_proto_rawDescData)
	})
	return file_parking_gunk_api_vehicle_all_proto_rawDescData
}

var (
	file_parking_gunk_api_vehicle_all_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
	file_parking_gunk_api_vehicle_all_proto_goTypes  = []interface{}{
		(*Vehicle)(nil),               // 0: vehicle.Vehicle
		(*VehicleData)(nil),           // 1: vehicle.VehicleData
		(*ParkVehicleRequest)(nil),    // 2: vehicle.ParkVehicleRequest
		(*ParkVehicleResponse)(nil),   // 3: vehicle.ParkVehicleResponse
		(*UnParkVehicleRequest)(nil),  // 4: vehicle.UnParkVehicleRequest
		(*UnParkVehicleResponse)(nil), // 5: vehicle.UnParkVehicleResponse
		(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
	}
)

var file_parking_gunk_api_vehicle_all_proto_depIdxs = []int32{
	6, // 0: vehicle.Vehicle.Entry:type_name -> google.protobuf.Timestamp
	6, // 1: vehicle.Vehicle.Exit:type_name -> google.protobuf.Timestamp
	6, // 2: vehicle.Vehicle.CreatedAt:type_name -> google.protobuf.Timestamp
	6, // 3: vehicle.Vehicle.UpdatedAt:type_name -> google.protobuf.Timestamp
	0, // 4: vehicle.VehicleData.Vehicle:type_name -> vehicle.Vehicle
	1, // 5: vehicle.ParkVehicleResponse.Data:type_name -> vehicle.VehicleData
	1, // 6: vehicle.UnParkVehicleResponse.Data:type_name -> vehicle.VehicleData
	2, // 7: vehicle.VehicleService.ParkVehicle:input_type -> vehicle.ParkVehicleRequest
	4, // 8: vehicle.VehicleService.UnParkVehicle:input_type -> vehicle.UnParkVehicleRequest
	3, // 9: vehicle.VehicleService.ParkVehicle:output_type -> vehicle.ParkVehicleResponse
	5, // 10: vehicle.VehicleService.UnParkVehicle:output_type -> vehicle.UnParkVehicleResponse
	9, // [9:11] is the sub-list for method output_type
	7, // [7:9] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_parking_gunk_api_vehicle_all_proto_init() }
func file_parking_gunk_api_vehicle_all_proto_init() {
	if File_parking_gunk_api_vehicle_all_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_parking_gunk_api_vehicle_all_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Vehicle); i {
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
		file_parking_gunk_api_vehicle_all_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VehicleData); i {
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
		file_parking_gunk_api_vehicle_all_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParkVehicleRequest); i {
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
		file_parking_gunk_api_vehicle_all_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParkVehicleResponse); i {
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
		file_parking_gunk_api_vehicle_all_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnParkVehicleRequest); i {
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
		file_parking_gunk_api_vehicle_all_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnParkVehicleResponse); i {
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
			RawDescriptor: file_parking_gunk_api_vehicle_all_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_parking_gunk_api_vehicle_all_proto_goTypes,
		DependencyIndexes: file_parking_gunk_api_vehicle_all_proto_depIdxs,
		MessageInfos:      file_parking_gunk_api_vehicle_all_proto_msgTypes,
	}.Build()
	File_parking_gunk_api_vehicle_all_proto = out.File
	file_parking_gunk_api_vehicle_all_proto_rawDesc = nil
	file_parking_gunk_api_vehicle_all_proto_goTypes = nil
	file_parking_gunk_api_vehicle_all_proto_depIdxs = nil
}

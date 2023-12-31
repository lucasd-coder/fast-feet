// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: request/get_all_order_request.proto

package pb

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

type GetAllOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	StartDate     string   `protobuf:"bytes,2,opt,name=startDate,proto3" json:"startDate,omitempty"`
	EndDate       string   `protobuf:"bytes,3,opt,name=endDate,proto3" json:"endDate,omitempty"`
	Product       *Product `protobuf:"bytes,4,opt,name=product,proto3" json:"product,omitempty"`
	Addresses     *Address `protobuf:"bytes,5,opt,name=addresses,proto3" json:"addresses,omitempty"`
	CreatedAt     string   `protobuf:"bytes,6,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt     string   `protobuf:"bytes,7,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	DeliverymanId string   `protobuf:"bytes,8,opt,name=deliverymanId,proto3" json:"deliverymanId,omitempty"`
	CanceledAt    string   `protobuf:"bytes,9,opt,name=canceledAt,proto3" json:"canceledAt,omitempty"`
	Limit         int64    `protobuf:"varint,10,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset        int64    `protobuf:"varint,11,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *GetAllOrderRequest) Reset() {
	*x = GetAllOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_request_get_all_order_request_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllOrderRequest) ProtoMessage() {}

func (x *GetAllOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_request_get_all_order_request_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllOrderRequest.ProtoReflect.Descriptor instead.
func (*GetAllOrderRequest) Descriptor() ([]byte, []int) {
	return file_request_get_all_order_request_proto_rawDescGZIP(), []int{0}
}

func (x *GetAllOrderRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetAllOrderRequest) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *GetAllOrderRequest) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

func (x *GetAllOrderRequest) GetProduct() *Product {
	if x != nil {
		return x.Product
	}
	return nil
}

func (x *GetAllOrderRequest) GetAddresses() *Address {
	if x != nil {
		return x.Addresses
	}
	return nil
}

func (x *GetAllOrderRequest) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *GetAllOrderRequest) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *GetAllOrderRequest) GetDeliverymanId() string {
	if x != nil {
		return x.DeliverymanId
	}
	return ""
}

func (x *GetAllOrderRequest) GetCanceledAt() string {
	if x != nil {
		return x.CanceledAt
	}
	return ""
}

func (x *GetAllOrderRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetAllOrderRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

var File_request_get_all_order_request_proto protoreflect.FileDescriptor

var file_request_get_all_order_request_proto_rawDesc = []byte{
	0x0a, 0x23, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2f, 0x67, 0x65, 0x74, 0x5f, 0x61, 0x6c,
	0x6c, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x11, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xde, 0x02, 0x0a,
	0x12, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x25, 0x0a, 0x07, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70,
	0x62, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x12, 0x29, 0x0a, 0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x52, 0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x12, 0x1c, 0x0a,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x64, 0x65, 0x6c,
	0x69, 0x76, 0x65, 0x72, 0x79, 0x6d, 0x61, 0x6e, 0x49, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x6d, 0x61, 0x6e, 0x49, 0x64, 0x12,
	0x1e, 0x0a, 0x0a, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x65, 0x64, 0x41, 0x74, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x42, 0x0a, 0x5a,
	0x08, 0x2e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_request_get_all_order_request_proto_rawDescOnce sync.Once
	file_request_get_all_order_request_proto_rawDescData = file_request_get_all_order_request_proto_rawDesc
)

func file_request_get_all_order_request_proto_rawDescGZIP() []byte {
	file_request_get_all_order_request_proto_rawDescOnce.Do(func() {
		file_request_get_all_order_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_request_get_all_order_request_proto_rawDescData)
	})
	return file_request_get_all_order_request_proto_rawDescData
}

var file_request_get_all_order_request_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_request_get_all_order_request_proto_goTypes = []interface{}{
	(*GetAllOrderRequest)(nil), // 0: pb.GetAllOrderRequest
	(*Product)(nil),            // 1: pb.Product
	(*Address)(nil),            // 2: pb.Address
}
var file_request_get_all_order_request_proto_depIdxs = []int32{
	1, // 0: pb.GetAllOrderRequest.product:type_name -> pb.Product
	2, // 1: pb.GetAllOrderRequest.addresses:type_name -> pb.Address
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_request_get_all_order_request_proto_init() }
func file_request_get_all_order_request_proto_init() {
	if File_request_get_all_order_request_proto != nil {
		return
	}
	file_model_order_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_request_get_all_order_request_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllOrderRequest); i {
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
			RawDescriptor: file_request_get_all_order_request_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_request_get_all_order_request_proto_goTypes,
		DependencyIndexes: file_request_get_all_order_request_proto_depIdxs,
		MessageInfos:      file_request_get_all_order_request_proto_msgTypes,
	}.Build()
	File_request_get_all_order_request_proto = out.File
	file_request_get_all_order_request_proto_rawDesc = nil
	file_request_get_all_order_request_proto_goTypes = nil
	file_request_get_all_order_request_proto_depIdxs = nil
}

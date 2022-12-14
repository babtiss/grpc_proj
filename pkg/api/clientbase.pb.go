// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.7.1
// source: clientbase.proto

package api

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

type Client struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Client) Reset() {
	*x = Client{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clientbase_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Client) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Client) ProtoMessage() {}

func (x *Client) ProtoReflect() protoreflect.Message {
	mi := &file_clientbase_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Client.ProtoReflect.Descriptor instead.
func (*Client) Descriptor() ([]byte, []int) {
	return file_clientbase_proto_rawDescGZIP(), []int{0}
}

func (x *Client) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type AddClientRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Client *Client `protobuf:"bytes,1,opt,name=client,proto3" json:"client,omitempty"`
}

func (x *AddClientRequest) Reset() {
	*x = AddClientRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clientbase_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddClientRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddClientRequest) ProtoMessage() {}

func (x *AddClientRequest) ProtoReflect() protoreflect.Message {
	mi := &file_clientbase_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddClientRequest.ProtoReflect.Descriptor instead.
func (*AddClientRequest) Descriptor() ([]byte, []int) {
	return file_clientbase_proto_rawDescGZIP(), []int{1}
}

func (x *AddClientRequest) GetClient() *Client {
	if x != nil {
		return x.Client
	}
	return nil
}

type AddClientResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AddClientResponse) Reset() {
	*x = AddClientResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clientbase_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddClientResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddClientResponse) ProtoMessage() {}

func (x *AddClientResponse) ProtoReflect() protoreflect.Message {
	mi := &file_clientbase_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddClientResponse.ProtoReflect.Descriptor instead.
func (*AddClientResponse) Descriptor() ([]byte, []int) {
	return file_clientbase_proto_rawDescGZIP(), []int{2}
}

type DeleteClientRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Client *Client `protobuf:"bytes,1,opt,name=client,proto3" json:"client,omitempty"`
}

func (x *DeleteClientRequest) Reset() {
	*x = DeleteClientRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clientbase_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteClientRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteClientRequest) ProtoMessage() {}

func (x *DeleteClientRequest) ProtoReflect() protoreflect.Message {
	mi := &file_clientbase_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteClientRequest.ProtoReflect.Descriptor instead.
func (*DeleteClientRequest) Descriptor() ([]byte, []int) {
	return file_clientbase_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteClientRequest) GetClient() *Client {
	if x != nil {
		return x.Client
	}
	return nil
}

type DeleteClientResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteClientResponse) Reset() {
	*x = DeleteClientResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clientbase_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteClientResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteClientResponse) ProtoMessage() {}

func (x *DeleteClientResponse) ProtoReflect() protoreflect.Message {
	mi := &file_clientbase_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteClientResponse.ProtoReflect.Descriptor instead.
func (*DeleteClientResponse) Descriptor() ([]byte, []int) {
	return file_clientbase_proto_rawDescGZIP(), []int{4}
}

type GetClientsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetClientsRequest) Reset() {
	*x = GetClientsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clientbase_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClientsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClientsRequest) ProtoMessage() {}

func (x *GetClientsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_clientbase_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClientsRequest.ProtoReflect.Descriptor instead.
func (*GetClientsRequest) Descriptor() ([]byte, []int) {
	return file_clientbase_proto_rawDescGZIP(), []int{5}
}

type GetClientsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Clients []*Client `protobuf:"bytes,1,rep,name=clients,proto3" json:"clients,omitempty"`
}

func (x *GetClientsResponse) Reset() {
	*x = GetClientsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_clientbase_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClientsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClientsResponse) ProtoMessage() {}

func (x *GetClientsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_clientbase_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClientsResponse.ProtoReflect.Descriptor instead.
func (*GetClientsResponse) Descriptor() ([]byte, []int) {
	return file_clientbase_proto_rawDescGZIP(), []int{6}
}

func (x *GetClientsResponse) GetClients() []*Client {
	if x != nil {
		return x.Clients
	}
	return nil
}

var File_clientbase_proto protoreflect.FileDescriptor

var file_clientbase_proto_rawDesc = []byte{
	0x0a, 0x10, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x62, 0x61, 0x73, 0x65, 0x22, 0x1c,
	0x0a, 0x06, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3e, 0x0a, 0x10,
	0x41, 0x64, 0x64, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x2a, 0x0a, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x22, 0x13, 0x0a, 0x11,
	0x41, 0x64, 0x64, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x41, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x06, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x22, 0x16, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x13, 0x0a, 0x11,
	0x47, 0x65, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x42, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x07, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x73, 0x32, 0xc8, 0x02, 0x0a, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x62, 0x61, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x1c, 0x2e, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x06, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x12, 0x1f, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x62, 0x61, 0x73,
	0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x0a, 0x47, 0x65, 0x74,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x1d, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x62,
	0x61, 0x73, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x56, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x46, 0x72, 0x6f, 0x6d, 0x52, 0x65, 0x64, 0x69, 0x73, 0x12,
	0x1d, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x47, 0x65, 0x74,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e,
	0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_clientbase_proto_rawDescOnce sync.Once
	file_clientbase_proto_rawDescData = file_clientbase_proto_rawDesc
)

func file_clientbase_proto_rawDescGZIP() []byte {
	file_clientbase_proto_rawDescOnce.Do(func() {
		file_clientbase_proto_rawDescData = protoimpl.X.CompressGZIP(file_clientbase_proto_rawDescData)
	})
	return file_clientbase_proto_rawDescData
}

var file_clientbase_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_clientbase_proto_goTypes = []interface{}{
	(*Client)(nil),               // 0: clientbase.Client
	(*AddClientRequest)(nil),     // 1: clientbase.AddClientRequest
	(*AddClientResponse)(nil),    // 2: clientbase.AddClientResponse
	(*DeleteClientRequest)(nil),  // 3: clientbase.DeleteClientRequest
	(*DeleteClientResponse)(nil), // 4: clientbase.DeleteClientResponse
	(*GetClientsRequest)(nil),    // 5: clientbase.GetClientsRequest
	(*GetClientsResponse)(nil),   // 6: clientbase.GetClientsResponse
}
var file_clientbase_proto_depIdxs = []int32{
	0, // 0: clientbase.AddClientRequest.client:type_name -> clientbase.Client
	0, // 1: clientbase.DeleteClientRequest.client:type_name -> clientbase.Client
	0, // 2: clientbase.GetClientsResponse.clients:type_name -> clientbase.Client
	1, // 3: clientbase.clientbase.Add:input_type -> clientbase.AddClientRequest
	3, // 4: clientbase.clientbase.Delete:input_type -> clientbase.DeleteClientRequest
	5, // 5: clientbase.clientbase.GetClients:input_type -> clientbase.GetClientsRequest
	5, // 6: clientbase.clientbase.GetClientsFromRedis:input_type -> clientbase.GetClientsRequest
	2, // 7: clientbase.clientbase.Add:output_type -> clientbase.AddClientResponse
	4, // 8: clientbase.clientbase.Delete:output_type -> clientbase.DeleteClientResponse
	6, // 9: clientbase.clientbase.GetClients:output_type -> clientbase.GetClientsResponse
	6, // 10: clientbase.clientbase.GetClientsFromRedis:output_type -> clientbase.GetClientsResponse
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_clientbase_proto_init() }
func file_clientbase_proto_init() {
	if File_clientbase_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_clientbase_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Client); i {
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
		file_clientbase_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddClientRequest); i {
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
		file_clientbase_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddClientResponse); i {
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
		file_clientbase_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteClientRequest); i {
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
		file_clientbase_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteClientResponse); i {
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
		file_clientbase_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetClientsRequest); i {
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
		file_clientbase_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetClientsResponse); i {
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
			RawDescriptor: file_clientbase_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_clientbase_proto_goTypes,
		DependencyIndexes: file_clientbase_proto_depIdxs,
		MessageInfos:      file_clientbase_proto_msgTypes,
	}.Build()
	File_clientbase_proto = out.File
	file_clientbase_proto_rawDesc = nil
	file_clientbase_proto_goTypes = nil
	file_clientbase_proto_depIdxs = nil
}

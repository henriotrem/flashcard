// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: card/v1/card.proto

package card

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Card struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Label    string `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
	Type     string `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Content  string `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	Created  int64  `protobuf:"varint,5,opt,name=created,proto3" json:"created,omitempty"`
	Modified int64  `protobuf:"varint,6,opt,name=modified,proto3" json:"modified,omitempty"`
}

func (x *Card) Reset() {
	*x = Card{}
	if protoimpl.UnsafeEnabled {
		mi := &file_card_v1_card_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Card) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Card) ProtoMessage() {}

func (x *Card) ProtoReflect() protoreflect.Message {
	mi := &file_card_v1_card_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Card.ProtoReflect.Descriptor instead.
func (*Card) Descriptor() ([]byte, []int) {
	return file_card_v1_card_proto_rawDescGZIP(), []int{0}
}

func (x *Card) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Card) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *Card) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Card) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Card) GetCreated() int64 {
	if x != nil {
		return x.Created
	}
	return 0
}

func (x *Card) GetModified() int64 {
	if x != nil {
		return x.Modified
	}
	return 0
}

type AddCardRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Card *Card `protobuf:"bytes,1,opt,name=card,proto3" json:"card,omitempty"`
}

func (x *AddCardRequest) Reset() {
	*x = AddCardRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_card_v1_card_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddCardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddCardRequest) ProtoMessage() {}

func (x *AddCardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_card_v1_card_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddCardRequest.ProtoReflect.Descriptor instead.
func (*AddCardRequest) Descriptor() ([]byte, []int) {
	return file_card_v1_card_proto_rawDescGZIP(), []int{1}
}

func (x *AddCardRequest) GetCard() *Card {
	if x != nil {
		return x.Card
	}
	return nil
}

type AddCardResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Card *Card `protobuf:"bytes,1,opt,name=card,proto3" json:"card,omitempty"`
}

func (x *AddCardResponse) Reset() {
	*x = AddCardResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_card_v1_card_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddCardResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddCardResponse) ProtoMessage() {}

func (x *AddCardResponse) ProtoReflect() protoreflect.Message {
	mi := &file_card_v1_card_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddCardResponse.ProtoReflect.Descriptor instead.
func (*AddCardResponse) Descriptor() ([]byte, []int) {
	return file_card_v1_card_proto_rawDescGZIP(), []int{2}
}

func (x *AddCardResponse) GetCard() *Card {
	if x != nil {
		return x.Card
	}
	return nil
}

type GetCardRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetCardRequest) Reset() {
	*x = GetCardRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_card_v1_card_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCardRequest) ProtoMessage() {}

func (x *GetCardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_card_v1_card_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCardRequest.ProtoReflect.Descriptor instead.
func (*GetCardRequest) Descriptor() ([]byte, []int) {
	return file_card_v1_card_proto_rawDescGZIP(), []int{3}
}

func (x *GetCardRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetCardResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Card *Card `protobuf:"bytes,1,opt,name=card,proto3" json:"card,omitempty"`
}

func (x *GetCardResponse) Reset() {
	*x = GetCardResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_card_v1_card_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCardResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCardResponse) ProtoMessage() {}

func (x *GetCardResponse) ProtoReflect() protoreflect.Message {
	mi := &file_card_v1_card_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCardResponse.ProtoReflect.Descriptor instead.
func (*GetCardResponse) Descriptor() ([]byte, []int) {
	return file_card_v1_card_proto_rawDescGZIP(), []int{4}
}

func (x *GetCardResponse) GetCard() *Card {
	if x != nil {
		return x.Card
	}
	return nil
}

type DeleteCardRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteCardRequest) Reset() {
	*x = DeleteCardRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_card_v1_card_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCardRequest) ProtoMessage() {}

func (x *DeleteCardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_card_v1_card_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCardRequest.ProtoReflect.Descriptor instead.
func (*DeleteCardRequest) Descriptor() ([]byte, []int) {
	return file_card_v1_card_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteCardRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteCardResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *DeleteCardResponse) Reset() {
	*x = DeleteCardResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_card_v1_card_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCardResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCardResponse) ProtoMessage() {}

func (x *DeleteCardResponse) ProtoReflect() protoreflect.Message {
	mi := &file_card_v1_card_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCardResponse.ProtoReflect.Descriptor instead.
func (*DeleteCardResponse) Descriptor() ([]byte, []int) {
	return file_card_v1_card_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteCardResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_card_v1_card_proto protoreflect.FileDescriptor

var file_card_v1_card_proto_rawDesc = []byte{
	0x0a, 0x12, 0x63, 0x61, 0x72, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x72, 0x64, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x90, 0x01, 0x0a, 0x04, 0x43, 0x61, 0x72, 0x64, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6d,
	0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6d,
	0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x22, 0x2b, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x43, 0x61,
	0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x04, 0x63, 0x61, 0x72,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x52, 0x04,
	0x63, 0x61, 0x72, 0x64, 0x22, 0x2c, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x43, 0x61, 0x72, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x04, 0x63, 0x61, 0x72, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x52, 0x04, 0x63, 0x61,
	0x72, 0x64, 0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x2c, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x04, 0x63, 0x61, 0x72, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x52, 0x04, 0x63, 0x61,
	0x72, 0x64, 0x22, 0x23, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x72, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2c, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x43, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0xa0, 0x01, 0x0a, 0x0b, 0x43, 0x61, 0x72, 0x64, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2c, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x43, 0x61, 0x72, 0x64,
	0x12, 0x0f, 0x2e, 0x41, 0x64, 0x64, 0x43, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x10, 0x2e, 0x41, 0x64, 0x64, 0x43, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x64, 0x12, 0x0f,
	0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x10, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x35, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x72, 0x64, 0x12,
	0x12, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x72, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x11, 0x5a, 0x0f, 0x2e, 0x2f, 0x63, 0x61,
	0x72, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x3b, 0x63, 0x61, 0x72, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_card_v1_card_proto_rawDescOnce sync.Once
	file_card_v1_card_proto_rawDescData = file_card_v1_card_proto_rawDesc
)

func file_card_v1_card_proto_rawDescGZIP() []byte {
	file_card_v1_card_proto_rawDescOnce.Do(func() {
		file_card_v1_card_proto_rawDescData = protoimpl.X.CompressGZIP(file_card_v1_card_proto_rawDescData)
	})
	return file_card_v1_card_proto_rawDescData
}

var file_card_v1_card_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_card_v1_card_proto_goTypes = []interface{}{
	(*Card)(nil),               // 0: Card
	(*AddCardRequest)(nil),     // 1: AddCardRequest
	(*AddCardResponse)(nil),    // 2: AddCardResponse
	(*GetCardRequest)(nil),     // 3: GetCardRequest
	(*GetCardResponse)(nil),    // 4: GetCardResponse
	(*DeleteCardRequest)(nil),  // 5: DeleteCardRequest
	(*DeleteCardResponse)(nil), // 6: DeleteCardResponse
}
var file_card_v1_card_proto_depIdxs = []int32{
	0, // 0: AddCardRequest.card:type_name -> Card
	0, // 1: AddCardResponse.card:type_name -> Card
	0, // 2: GetCardResponse.card:type_name -> Card
	1, // 3: CardService.AddCard:input_type -> AddCardRequest
	3, // 4: CardService.GetCard:input_type -> GetCardRequest
	5, // 5: CardService.DeleteCard:input_type -> DeleteCardRequest
	2, // 6: CardService.AddCard:output_type -> AddCardResponse
	4, // 7: CardService.GetCard:output_type -> GetCardResponse
	6, // 8: CardService.DeleteCard:output_type -> DeleteCardResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_card_v1_card_proto_init() }
func file_card_v1_card_proto_init() {
	if File_card_v1_card_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_card_v1_card_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Card); i {
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
		file_card_v1_card_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddCardRequest); i {
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
		file_card_v1_card_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddCardResponse); i {
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
		file_card_v1_card_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCardRequest); i {
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
		file_card_v1_card_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCardResponse); i {
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
		file_card_v1_card_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCardRequest); i {
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
		file_card_v1_card_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCardResponse); i {
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
			RawDescriptor: file_card_v1_card_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_card_v1_card_proto_goTypes,
		DependencyIndexes: file_card_v1_card_proto_depIdxs,
		MessageInfos:      file_card_v1_card_proto_msgTypes,
	}.Build()
	File_card_v1_card_proto = out.File
	file_card_v1_card_proto_rawDesc = nil
	file_card_v1_card_proto_goTypes = nil
	file_card_v1_card_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CardServiceClient is the client API for CardService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CardServiceClient interface {
	AddCard(ctx context.Context, in *AddCardRequest, opts ...grpc.CallOption) (*AddCardResponse, error)
	GetCard(ctx context.Context, in *GetCardRequest, opts ...grpc.CallOption) (*GetCardResponse, error)
	DeleteCard(ctx context.Context, in *DeleteCardRequest, opts ...grpc.CallOption) (*DeleteCardResponse, error)
}

type cardServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCardServiceClient(cc grpc.ClientConnInterface) CardServiceClient {
	return &cardServiceClient{cc}
}

func (c *cardServiceClient) AddCard(ctx context.Context, in *AddCardRequest, opts ...grpc.CallOption) (*AddCardResponse, error) {
	out := new(AddCardResponse)
	err := c.cc.Invoke(ctx, "/CardService/AddCard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cardServiceClient) GetCard(ctx context.Context, in *GetCardRequest, opts ...grpc.CallOption) (*GetCardResponse, error) {
	out := new(GetCardResponse)
	err := c.cc.Invoke(ctx, "/CardService/GetCard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cardServiceClient) DeleteCard(ctx context.Context, in *DeleteCardRequest, opts ...grpc.CallOption) (*DeleteCardResponse, error) {
	out := new(DeleteCardResponse)
	err := c.cc.Invoke(ctx, "/CardService/DeleteCard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CardServiceServer is the server API for CardService service.
type CardServiceServer interface {
	AddCard(context.Context, *AddCardRequest) (*AddCardResponse, error)
	GetCard(context.Context, *GetCardRequest) (*GetCardResponse, error)
	DeleteCard(context.Context, *DeleteCardRequest) (*DeleteCardResponse, error)
}

// UnimplementedCardServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCardServiceServer struct {
}

func (*UnimplementedCardServiceServer) AddCard(context.Context, *AddCardRequest) (*AddCardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCard not implemented")
}
func (*UnimplementedCardServiceServer) GetCard(context.Context, *GetCardRequest) (*GetCardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCard not implemented")
}
func (*UnimplementedCardServiceServer) DeleteCard(context.Context, *DeleteCardRequest) (*DeleteCardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCard not implemented")
}

func RegisterCardServiceServer(s *grpc.Server, srv CardServiceServer) {
	s.RegisterService(&_CardService_serviceDesc, srv)
}

func _CardService_AddCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CardServiceServer).AddCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CardService/AddCard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CardServiceServer).AddCard(ctx, req.(*AddCardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CardService_GetCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CardServiceServer).GetCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CardService/GetCard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CardServiceServer).GetCard(ctx, req.(*GetCardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CardService_DeleteCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CardServiceServer).DeleteCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CardService/DeleteCard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CardServiceServer).DeleteCard(ctx, req.(*DeleteCardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CardService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "CardService",
	HandlerType: (*CardServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddCard",
			Handler:    _CardService_AddCard_Handler,
		},
		{
			MethodName: "GetCard",
			Handler:    _CardService_GetCard_Handler,
		},
		{
			MethodName: "DeleteCard",
			Handler:    _CardService_DeleteCard_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "card/v1/card.proto",
}

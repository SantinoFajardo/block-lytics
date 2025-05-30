// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: accounts/accounts.proto

package accountspb

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

type GetAccountByAdressRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Address       string                 `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAccountByAdressRequest) Reset() {
	*x = GetAccountByAdressRequest{}
	mi := &file_accounts_accounts_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAccountByAdressRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccountByAdressRequest) ProtoMessage() {}

func (x *GetAccountByAdressRequest) ProtoReflect() protoreflect.Message {
	mi := &file_accounts_accounts_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccountByAdressRequest.ProtoReflect.Descriptor instead.
func (*GetAccountByAdressRequest) Descriptor() ([]byte, []int) {
	return file_accounts_accounts_proto_rawDescGZIP(), []int{0}
}

func (x *GetAccountByAdressRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type GetAccountByAdressResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Address       string                 `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Blockchain    string                 `protobuf:"bytes,3,opt,name=blockchain,proto3" json:"blockchain,omitempty"`
	UserId        int64                  `protobuf:"varint,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAccountByAdressResponse) Reset() {
	*x = GetAccountByAdressResponse{}
	mi := &file_accounts_accounts_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAccountByAdressResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccountByAdressResponse) ProtoMessage() {}

func (x *GetAccountByAdressResponse) ProtoReflect() protoreflect.Message {
	mi := &file_accounts_accounts_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccountByAdressResponse.ProtoReflect.Descriptor instead.
func (*GetAccountByAdressResponse) Descriptor() ([]byte, []int) {
	return file_accounts_accounts_proto_rawDescGZIP(), []int{1}
}

func (x *GetAccountByAdressResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetAccountByAdressResponse) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *GetAccountByAdressResponse) GetBlockchain() string {
	if x != nil {
		return x.Blockchain
	}
	return ""
}

func (x *GetAccountByAdressResponse) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *GetAccountByAdressResponse) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *GetAccountByAdressResponse) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type CreateAccountRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Address       string                 `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Blockchain    string                 `protobuf:"bytes,2,opt,name=blockchain,proto3" json:"blockchain,omitempty"`
	UserId        int64                  `protobuf:"varint,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateAccountRequest) Reset() {
	*x = CreateAccountRequest{}
	mi := &file_accounts_accounts_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAccountRequest) ProtoMessage() {}

func (x *CreateAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_accounts_accounts_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAccountRequest.ProtoReflect.Descriptor instead.
func (*CreateAccountRequest) Descriptor() ([]byte, []int) {
	return file_accounts_accounts_proto_rawDescGZIP(), []int{2}
}

func (x *CreateAccountRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *CreateAccountRequest) GetBlockchain() string {
	if x != nil {
		return x.Blockchain
	}
	return ""
}

func (x *CreateAccountRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type CreateAccountResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`          // UUID generado o identificador del registro
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"` // Confirmación u observación
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateAccountResponse) Reset() {
	*x = CreateAccountResponse{}
	mi := &file_accounts_accounts_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateAccountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAccountResponse) ProtoMessage() {}

func (x *CreateAccountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_accounts_accounts_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAccountResponse.ProtoReflect.Descriptor instead.
func (*CreateAccountResponse) Descriptor() ([]byte, []int) {
	return file_accounts_accounts_proto_rawDescGZIP(), []int{3}
}

func (x *CreateAccountResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CreateAccountResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_accounts_accounts_proto protoreflect.FileDescriptor

const file_accounts_accounts_proto_rawDesc = "" +
	"\n" +
	"\x17accounts/accounts.proto\x12\baccounts\x1a\x1fgoogle/protobuf/timestamp.proto\"5\n" +
	"\x19GetAccountByAdressRequest\x12\x18\n" +
	"\aaddress\x18\x01 \x01(\tR\aaddress\"\xf5\x01\n" +
	"\x1aGetAccountByAdressResponse\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x03R\x02id\x12\x18\n" +
	"\aaddress\x18\x02 \x01(\tR\aaddress\x12\x1e\n" +
	"\n" +
	"blockchain\x18\x03 \x01(\tR\n" +
	"blockchain\x12\x17\n" +
	"\auser_id\x18\x04 \x01(\x03R\x06userId\x129\n" +
	"\n" +
	"created_at\x18\x05 \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\x129\n" +
	"\n" +
	"updated_at\x18\x06 \x01(\v2\x1a.google.protobuf.TimestampR\tupdatedAt\"i\n" +
	"\x14CreateAccountRequest\x12\x18\n" +
	"\aaddress\x18\x01 \x01(\tR\aaddress\x12\x1e\n" +
	"\n" +
	"blockchain\x18\x02 \x01(\tR\n" +
	"blockchain\x12\x17\n" +
	"\auser_id\x18\x03 \x01(\x03R\x06userId\"A\n" +
	"\x15CreateAccountResponse\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x03R\x02id\x12\x18\n" +
	"\amessage\x18\x02 \x01(\tR\amessage2\xc4\x01\n" +
	"\x0fAccountsService\x12P\n" +
	"\rCreateAccount\x12\x1e.accounts.CreateAccountRequest\x1a\x1f.accounts.CreateAccountResponse\x12_\n" +
	"\x12GetAccountByAdress\x12#.accounts.GetAccountByAdressRequest\x1a$.accounts.GetAccountByAdressResponseB3Z1github.com/block-lytics/proto/accounts;accountspbb\x06proto3"

var (
	file_accounts_accounts_proto_rawDescOnce sync.Once
	file_accounts_accounts_proto_rawDescData []byte
)

func file_accounts_accounts_proto_rawDescGZIP() []byte {
	file_accounts_accounts_proto_rawDescOnce.Do(func() {
		file_accounts_accounts_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_accounts_accounts_proto_rawDesc), len(file_accounts_accounts_proto_rawDesc)))
	})
	return file_accounts_accounts_proto_rawDescData
}

var file_accounts_accounts_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_accounts_accounts_proto_goTypes = []any{
	(*GetAccountByAdressRequest)(nil),  // 0: accounts.GetAccountByAdressRequest
	(*GetAccountByAdressResponse)(nil), // 1: accounts.GetAccountByAdressResponse
	(*CreateAccountRequest)(nil),       // 2: accounts.CreateAccountRequest
	(*CreateAccountResponse)(nil),      // 3: accounts.CreateAccountResponse
	(*timestamppb.Timestamp)(nil),      // 4: google.protobuf.Timestamp
}
var file_accounts_accounts_proto_depIdxs = []int32{
	4, // 0: accounts.GetAccountByAdressResponse.created_at:type_name -> google.protobuf.Timestamp
	4, // 1: accounts.GetAccountByAdressResponse.updated_at:type_name -> google.protobuf.Timestamp
	2, // 2: accounts.AccountsService.CreateAccount:input_type -> accounts.CreateAccountRequest
	0, // 3: accounts.AccountsService.GetAccountByAdress:input_type -> accounts.GetAccountByAdressRequest
	3, // 4: accounts.AccountsService.CreateAccount:output_type -> accounts.CreateAccountResponse
	1, // 5: accounts.AccountsService.GetAccountByAdress:output_type -> accounts.GetAccountByAdressResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_accounts_accounts_proto_init() }
func file_accounts_accounts_proto_init() {
	if File_accounts_accounts_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_accounts_accounts_proto_rawDesc), len(file_accounts_accounts_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_accounts_accounts_proto_goTypes,
		DependencyIndexes: file_accounts_accounts_proto_depIdxs,
		MessageInfos:      file_accounts_accounts_proto_msgTypes,
	}.Build()
	File_accounts_accounts_proto = out.File
	file_accounts_accounts_proto_goTypes = nil
	file_accounts_accounts_proto_depIdxs = nil
}

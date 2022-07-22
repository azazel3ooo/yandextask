// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.3
// source: grpc/grpc.proto

package grpc

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

type GetterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetterRequest) Reset() {
	*x = GetterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_grpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetterRequest) ProtoMessage() {}

func (x *GetterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_grpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetterRequest.ProtoReflect.Descriptor instead.
func (*GetterRequest) Descriptor() ([]byte, []int) {
	return file_grpc_grpc_proto_rawDescGZIP(), []int{0}
}

func (x *GetterRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url   string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Error string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *GetterResponse) Reset() {
	*x = GetterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_grpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetterResponse) ProtoMessage() {}

func (x *GetterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_grpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetterResponse.ProtoReflect.Descriptor instead.
func (*GetterResponse) Descriptor() ([]byte, []int) {
	return file_grpc_grpc_proto_rawDescGZIP(), []int{1}
}

func (x *GetterResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *GetterResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

// для setter и jsonSetter одинаковые структуры
type SetterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *SetterRequest) Reset() {
	*x = SetterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_grpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetterRequest) ProtoMessage() {}

func (x *SetterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_grpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetterRequest.ProtoReflect.Descriptor instead.
func (*SetterRequest) Descriptor() ([]byte, []int) {
	return file_grpc_grpc_proto_rawDescGZIP(), []int{2}
}

func (x *SetterRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type SetterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	Error  string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *SetterResponse) Reset() {
	*x = SetterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_grpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetterResponse) ProtoMessage() {}

func (x *SetterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_grpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetterResponse.ProtoReflect.Descriptor instead.
func (*SetterResponse) Descriptor() ([]byte, []int) {
	return file_grpc_grpc_proto_rawDescGZIP(), []int{3}
}

func (x *SetterResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

func (x *SetterResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type UserUrlsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *UserUrlsRequest) Reset() {
	*x = UserUrlsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_grpc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserUrlsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserUrlsRequest) ProtoMessage() {}

func (x *UserUrlsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_grpc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserUrlsRequest.ProtoReflect.Descriptor instead.
func (*UserUrlsRequest) Descriptor() ([]byte, []int) {
	return file_grpc_grpc_proto_rawDescGZIP(), []int{4}
}

func (x *UserUrlsRequest) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

type UserUrlsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result []*UserUrlsResponse_UrlsPair `protobuf:"bytes,1,rep,name=result,proto3" json:"result,omitempty"`
	Error  string                       `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *UserUrlsResponse) Reset() {
	*x = UserUrlsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_grpc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserUrlsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserUrlsResponse) ProtoMessage() {}

func (x *UserUrlsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_grpc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserUrlsResponse.ProtoReflect.Descriptor instead.
func (*UserUrlsResponse) Descriptor() ([]byte, []int) {
	return file_grpc_grpc_proto_rawDescGZIP(), []int{5}
}

func (x *UserUrlsResponse) GetResult() []*UserUrlsResponse_UrlsPair {
	if x != nil {
		return x.Result
	}
	return nil
}

func (x *UserUrlsResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type SetManyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Set []*SetManyRequest_CustomIDSetReq `protobuf:"bytes,1,rep,name=set,proto3" json:"set,omitempty"`
}

func (x *SetManyRequest) Reset() {
	*x = SetManyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_grpc_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetManyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetManyRequest) ProtoMessage() {}

func (x *SetManyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_grpc_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetManyRequest.ProtoReflect.Descriptor instead.
func (*SetManyRequest) Descriptor() ([]byte, []int) {
	return file_grpc_grpc_proto_rawDescGZIP(), []int{6}
}

func (x *SetManyRequest) GetSet() []*SetManyRequest_CustomIDSetReq {
	if x != nil {
		return x.Set
	}
	return nil
}

type SetManyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Set   []*SetManyResponse_CustomIDSetResp `protobuf:"bytes,1,rep,name=set,proto3" json:"set,omitempty"`
	Error string                             `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *SetManyResponse) Reset() {
	*x = SetManyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_grpc_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetManyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetManyResponse) ProtoMessage() {}

func (x *SetManyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_grpc_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetManyResponse.ProtoReflect.Descriptor instead.
func (*SetManyResponse) Descriptor() ([]byte, []int) {
	return file_grpc_grpc_proto_rawDescGZIP(), []int{7}
}

func (x *SetManyResponse) GetSet() []*SetManyResponse_CustomIDSetResp {
	if x != nil {
		return x.Set
	}
	return nil
}

func (x *SetManyResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type AsyncDeleteReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Array []string `protobuf:"bytes,1,rep,name=array,proto3" json:"array,omitempty"`
}

func (x *AsyncDeleteReq) Reset() {
	*x = AsyncDeleteReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_grpc_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AsyncDeleteReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AsyncDeleteReq) ProtoMessage() {}

func (x *AsyncDeleteReq) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_grpc_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AsyncDeleteReq.ProtoReflect.Descriptor instead.
func (*AsyncDeleteReq) Descriptor() ([]byte, []int) {
	return file_grpc_grpc_proto_rawDescGZIP(), []int{8}
}

func (x *AsyncDeleteReq) GetArray() []string {
	if x != nil {
		return x.Array
	}
	return nil
}

type AsyncDeleteResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Error  string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *AsyncDeleteResp) Reset() {
	*x = AsyncDeleteResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_grpc_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AsyncDeleteResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AsyncDeleteResp) ProtoMessage() {}

func (x *AsyncDeleteResp) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_grpc_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AsyncDeleteResp.ProtoReflect.Descriptor instead.
func (*AsyncDeleteResp) Descriptor() ([]byte, []int) {
	return file_grpc_grpc_proto_rawDescGZIP(), []int{9}
}

func (x *AsyncDeleteResp) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *AsyncDeleteResp) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type GetStatReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ping string `protobuf:"bytes,1,opt,name=ping,proto3" json:"ping,omitempty"`
}

func (x *GetStatReq) Reset() {
	*x = GetStatReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_grpc_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStatReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStatReq) ProtoMessage() {}

func (x *GetStatReq) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_grpc_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStatReq.ProtoReflect.Descriptor instead.
func (*GetStatReq) Descriptor() ([]byte, []int) {
	return file_grpc_grpc_proto_rawDescGZIP(), []int{10}
}

func (x *GetStatReq) GetPing() string {
	if x != nil {
		return x.Ping
	}
	return ""
}

type GetStatResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users int32  `protobuf:"varint,1,opt,name=users,proto3" json:"users,omitempty"`
	Urls  int32  `protobuf:"varint,2,opt,name=urls,proto3" json:"urls,omitempty"`
	Error string `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *GetStatResp) Reset() {
	*x = GetStatResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_grpc_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStatResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStatResp) ProtoMessage() {}

func (x *GetStatResp) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_grpc_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStatResp.ProtoReflect.Descriptor instead.
func (*GetStatResp) Descriptor() ([]byte, []int) {
	return file_grpc_grpc_proto_rawDescGZIP(), []int{11}
}

func (x *GetStatResp) GetUsers() int32 {
	if x != nil {
		return x.Users
	}
	return 0
}

func (x *GetStatResp) GetUrls() int32 {
	if x != nil {
		return x.Urls
	}
	return 0
}

func (x *GetStatResp) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type UserUrlsResponse_UrlsPair struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Short    string `protobuf:"bytes,1,opt,name=short,proto3" json:"short,omitempty"`
	Original string `protobuf:"bytes,2,opt,name=original,proto3" json:"original,omitempty"`
}

func (x *UserUrlsResponse_UrlsPair) Reset() {
	*x = UserUrlsResponse_UrlsPair{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_grpc_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserUrlsResponse_UrlsPair) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserUrlsResponse_UrlsPair) ProtoMessage() {}

func (x *UserUrlsResponse_UrlsPair) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_grpc_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserUrlsResponse_UrlsPair.ProtoReflect.Descriptor instead.
func (*UserUrlsResponse_UrlsPair) Descriptor() ([]byte, []int) {
	return file_grpc_grpc_proto_rawDescGZIP(), []int{5, 0}
}

func (x *UserUrlsResponse_UrlsPair) GetShort() string {
	if x != nil {
		return x.Short
	}
	return ""
}

func (x *UserUrlsResponse_UrlsPair) GetOriginal() string {
	if x != nil {
		return x.Original
	}
	return ""
}

type SetManyRequest_CustomIDSetReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Original string `protobuf:"bytes,2,opt,name=original,proto3" json:"original,omitempty"`
}

func (x *SetManyRequest_CustomIDSetReq) Reset() {
	*x = SetManyRequest_CustomIDSetReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_grpc_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetManyRequest_CustomIDSetReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetManyRequest_CustomIDSetReq) ProtoMessage() {}

func (x *SetManyRequest_CustomIDSetReq) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_grpc_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetManyRequest_CustomIDSetReq.ProtoReflect.Descriptor instead.
func (*SetManyRequest_CustomIDSetReq) Descriptor() ([]byte, []int) {
	return file_grpc_grpc_proto_rawDescGZIP(), []int{6, 0}
}

func (x *SetManyRequest_CustomIDSetReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SetManyRequest_CustomIDSetReq) GetOriginal() string {
	if x != nil {
		return x.Original
	}
	return ""
}

type SetManyResponse_CustomIDSetResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Original string `protobuf:"bytes,2,opt,name=original,proto3" json:"original,omitempty"`
	Short    string `protobuf:"bytes,3,opt,name=short,proto3" json:"short,omitempty"`
}

func (x *SetManyResponse_CustomIDSetResp) Reset() {
	*x = SetManyResponse_CustomIDSetResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_grpc_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetManyResponse_CustomIDSetResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetManyResponse_CustomIDSetResp) ProtoMessage() {}

func (x *SetManyResponse_CustomIDSetResp) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_grpc_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetManyResponse_CustomIDSetResp.ProtoReflect.Descriptor instead.
func (*SetManyResponse_CustomIDSetResp) Descriptor() ([]byte, []int) {
	return file_grpc_grpc_proto_rawDescGZIP(), []int{7, 0}
}

func (x *SetManyResponse_CustomIDSetResp) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SetManyResponse_CustomIDSetResp) GetOriginal() string {
	if x != nil {
		return x.Original
	}
	return ""
}

func (x *SetManyResponse_CustomIDSetResp) GetShort() string {
	if x != nil {
		return x.Short
	}
	return ""
}

var File_grpc_grpc_proto protoreflect.FileDescriptor

var file_grpc_grpc_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x67, 0x72, 0x70, 0x63, 0x22, 0x1f, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x74, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x38, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x74,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x22, 0x21, 0x0a, 0x0d, 0x53, 0x65, 0x74, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x3e, 0x0a, 0x0e, 0x53, 0x65, 0x74, 0x74, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x25, 0x0a, 0x0f, 0x55, 0x73, 0x65, 0x72, 0x55, 0x72, 0x6c,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x9f, 0x01, 0x0a,
	0x10, 0x55, 0x73, 0x65, 0x72, 0x55, 0x72, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x37, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x55, 0x72, 0x6c,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x55, 0x72, 0x6c, 0x73, 0x50, 0x61,
	0x69, 0x72, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x1a, 0x3c, 0x0a, 0x08, 0x55, 0x72, 0x6c, 0x73, 0x50, 0x61, 0x69, 0x72, 0x12, 0x14, 0x0a, 0x05,
	0x73, 0x68, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x68, 0x6f,
	0x72, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x22, 0x85,
	0x01, 0x0a, 0x0e, 0x53, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x35, 0x0a, 0x03, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x49, 0x44, 0x53, 0x65, 0x74,
	0x52, 0x65, 0x71, 0x52, 0x03, 0x73, 0x65, 0x74, 0x1a, 0x3c, 0x0a, 0x0e, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x49, 0x44, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x72,
	0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x72,
	0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x22, 0xb5, 0x01, 0x0a, 0x0f, 0x53, 0x65, 0x74, 0x4d, 0x61,
	0x6e, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x03, 0x73, 0x65,
	0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x53,
	0x65, 0x74, 0x4d, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x49, 0x44, 0x53, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x52, 0x03,
	0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x1a, 0x53, 0x0a, 0x0f, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x49, 0x44, 0x53, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x68, 0x6f, 0x72,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x22, 0x26,
	0x0a, 0x0e, 0x41, 0x73, 0x79, 0x6e, 0x63, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x12, 0x14, 0x0a, 0x05, 0x61, 0x72, 0x72, 0x61, 0x79, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x05, 0x61, 0x72, 0x72, 0x61, 0x79, 0x22, 0x3f, 0x0a, 0x0f, 0x41, 0x73, 0x79, 0x6e, 0x63, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x20, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x69, 0x6e, 0x67, 0x22, 0x4d, 0x0a, 0x0b, 0x47, 0x65, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x75, 0x72, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x75, 0x72,
	0x6c, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0xc9, 0x02, 0x0a, 0x04, 0x55, 0x72, 0x6c,
	0x73, 0x12, 0x30, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x13, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x47, 0x65, 0x74, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x03, 0x53, 0x65, 0x74, 0x12, 0x13, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x53, 0x65, 0x74, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x14, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x65, 0x74, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x55, 0x72, 0x6c,
	0x73, 0x12, 0x15, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x55, 0x72, 0x6c,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x55, 0x72, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x36, 0x0a, 0x07, 0x53, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x79, 0x12, 0x14, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x53, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x15, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x0b, 0x41, 0x73, 0x79, 0x6e,
	0x63, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x41,
	0x73, 0x79, 0x6e, 0x63, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x73, 0x79, 0x6e, 0x63, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x2e, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x12,
	0x10, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x52, 0x65,
	0x71, 0x1a, 0x11, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x42, 0x0f, 0x5a, 0x0d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_grpc_proto_rawDescOnce sync.Once
	file_grpc_grpc_proto_rawDescData = file_grpc_grpc_proto_rawDesc
)

func file_grpc_grpc_proto_rawDescGZIP() []byte {
	file_grpc_grpc_proto_rawDescOnce.Do(func() {
		file_grpc_grpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_grpc_proto_rawDescData)
	})
	return file_grpc_grpc_proto_rawDescData
}

var file_grpc_grpc_proto_msgTypes = make([]protoimpl.MessageInfo, 15)
var file_grpc_grpc_proto_goTypes = []interface{}{
	(*GetterRequest)(nil),                   // 0: grpc.GetterRequest
	(*GetterResponse)(nil),                  // 1: grpc.GetterResponse
	(*SetterRequest)(nil),                   // 2: grpc.SetterRequest
	(*SetterResponse)(nil),                  // 3: grpc.SetterResponse
	(*UserUrlsRequest)(nil),                 // 4: grpc.UserUrlsRequest
	(*UserUrlsResponse)(nil),                // 5: grpc.UserUrlsResponse
	(*SetManyRequest)(nil),                  // 6: grpc.SetManyRequest
	(*SetManyResponse)(nil),                 // 7: grpc.SetManyResponse
	(*AsyncDeleteReq)(nil),                  // 8: grpc.AsyncDeleteReq
	(*AsyncDeleteResp)(nil),                 // 9: grpc.AsyncDeleteResp
	(*GetStatReq)(nil),                      // 10: grpc.GetStatReq
	(*GetStatResp)(nil),                     // 11: grpc.GetStatResp
	(*UserUrlsResponse_UrlsPair)(nil),       // 12: grpc.UserUrlsResponse.UrlsPair
	(*SetManyRequest_CustomIDSetReq)(nil),   // 13: grpc.SetManyRequest.CustomIDSetReq
	(*SetManyResponse_CustomIDSetResp)(nil), // 14: grpc.SetManyResponse.CustomIDSetResp
}
var file_grpc_grpc_proto_depIdxs = []int32{
	12, // 0: grpc.UserUrlsResponse.result:type_name -> grpc.UserUrlsResponse.UrlsPair
	13, // 1: grpc.SetManyRequest.set:type_name -> grpc.SetManyRequest.CustomIDSetReq
	14, // 2: grpc.SetManyResponse.set:type_name -> grpc.SetManyResponse.CustomIDSetResp
	0,  // 3: grpc.Urls.Get:input_type -> grpc.GetterRequest
	2,  // 4: grpc.Urls.Set:input_type -> grpc.SetterRequest
	4,  // 5: grpc.Urls.UserUrls:input_type -> grpc.UserUrlsRequest
	6,  // 6: grpc.Urls.SetMany:input_type -> grpc.SetManyRequest
	8,  // 7: grpc.Urls.AsyncDelete:input_type -> grpc.AsyncDeleteReq
	10, // 8: grpc.Urls.GetStat:input_type -> grpc.GetStatReq
	1,  // 9: grpc.Urls.Get:output_type -> grpc.GetterResponse
	3,  // 10: grpc.Urls.Set:output_type -> grpc.SetterResponse
	5,  // 11: grpc.Urls.UserUrls:output_type -> grpc.UserUrlsResponse
	7,  // 12: grpc.Urls.SetMany:output_type -> grpc.SetManyResponse
	9,  // 13: grpc.Urls.AsyncDelete:output_type -> grpc.AsyncDeleteResp
	11, // 14: grpc.Urls.GetStat:output_type -> grpc.GetStatResp
	9,  // [9:15] is the sub-list for method output_type
	3,  // [3:9] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_grpc_grpc_proto_init() }
func file_grpc_grpc_proto_init() {
	if File_grpc_grpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_grpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetterRequest); i {
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
		file_grpc_grpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetterResponse); i {
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
		file_grpc_grpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetterRequest); i {
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
		file_grpc_grpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetterResponse); i {
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
		file_grpc_grpc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserUrlsRequest); i {
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
		file_grpc_grpc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserUrlsResponse); i {
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
		file_grpc_grpc_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetManyRequest); i {
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
		file_grpc_grpc_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetManyResponse); i {
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
		file_grpc_grpc_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AsyncDeleteReq); i {
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
		file_grpc_grpc_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AsyncDeleteResp); i {
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
		file_grpc_grpc_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStatReq); i {
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
		file_grpc_grpc_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStatResp); i {
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
		file_grpc_grpc_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserUrlsResponse_UrlsPair); i {
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
		file_grpc_grpc_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetManyRequest_CustomIDSetReq); i {
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
		file_grpc_grpc_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetManyResponse_CustomIDSetResp); i {
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
			RawDescriptor: file_grpc_grpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   15,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_grpc_proto_goTypes,
		DependencyIndexes: file_grpc_grpc_proto_depIdxs,
		MessageInfos:      file_grpc_grpc_proto_msgTypes,
	}.Build()
	File_grpc_grpc_proto = out.File
	file_grpc_grpc_proto_rawDesc = nil
	file_grpc_grpc_proto_goTypes = nil
	file_grpc_grpc_proto_depIdxs = nil
}

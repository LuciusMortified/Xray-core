// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v5.29.2
// source: transport/internet/headers/http/config.proto

package http

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

type Header struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// "Accept", "Cookie", etc
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Each entry must be valid in one piece. Random entry will be chosen if
	// multiple entries present.
	Value         []string `protobuf:"bytes,2,rep,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Header) Reset() {
	*x = Header{}
	mi := &file_transport_internet_headers_http_config_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Header) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Header) ProtoMessage() {}

func (x *Header) ProtoReflect() protoreflect.Message {
	mi := &file_transport_internet_headers_http_config_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Header.ProtoReflect.Descriptor instead.
func (*Header) Descriptor() ([]byte, []int) {
	return file_transport_internet_headers_http_config_proto_rawDescGZIP(), []int{0}
}

func (x *Header) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Header) GetValue() []string {
	if x != nil {
		return x.Value
	}
	return nil
}

// HTTP version. Default value "1.1".
type Version struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Value         string                 `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Version) Reset() {
	*x = Version{}
	mi := &file_transport_internet_headers_http_config_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Version) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Version) ProtoMessage() {}

func (x *Version) ProtoReflect() protoreflect.Message {
	mi := &file_transport_internet_headers_http_config_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Version.ProtoReflect.Descriptor instead.
func (*Version) Descriptor() ([]byte, []int) {
	return file_transport_internet_headers_http_config_proto_rawDescGZIP(), []int{1}
}

func (x *Version) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

// HTTP method. Default value "GET".
type Method struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Value         string                 `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Method) Reset() {
	*x = Method{}
	mi := &file_transport_internet_headers_http_config_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Method) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Method) ProtoMessage() {}

func (x *Method) ProtoReflect() protoreflect.Message {
	mi := &file_transport_internet_headers_http_config_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Method.ProtoReflect.Descriptor instead.
func (*Method) Descriptor() ([]byte, []int) {
	return file_transport_internet_headers_http_config_proto_rawDescGZIP(), []int{2}
}

func (x *Method) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type RequestConfig struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Full HTTP version like "1.1".
	Version *Version `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	// GET, POST, CONNECT etc
	Method *Method `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
	// URI like "/login.php"
	Uri           []string  `protobuf:"bytes,3,rep,name=uri,proto3" json:"uri,omitempty"`
	Header        []*Header `protobuf:"bytes,4,rep,name=header,proto3" json:"header,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RequestConfig) Reset() {
	*x = RequestConfig{}
	mi := &file_transport_internet_headers_http_config_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RequestConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestConfig) ProtoMessage() {}

func (x *RequestConfig) ProtoReflect() protoreflect.Message {
	mi := &file_transport_internet_headers_http_config_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestConfig.ProtoReflect.Descriptor instead.
func (*RequestConfig) Descriptor() ([]byte, []int) {
	return file_transport_internet_headers_http_config_proto_rawDescGZIP(), []int{3}
}

func (x *RequestConfig) GetVersion() *Version {
	if x != nil {
		return x.Version
	}
	return nil
}

func (x *RequestConfig) GetMethod() *Method {
	if x != nil {
		return x.Method
	}
	return nil
}

func (x *RequestConfig) GetUri() []string {
	if x != nil {
		return x.Uri
	}
	return nil
}

func (x *RequestConfig) GetHeader() []*Header {
	if x != nil {
		return x.Header
	}
	return nil
}

type Status struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Status code. Default "200".
	Code string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	// Statue reason. Default "OK".
	Reason        string `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Status) Reset() {
	*x = Status{}
	mi := &file_transport_internet_headers_http_config_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_transport_internet_headers_http_config_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_transport_internet_headers_http_config_proto_rawDescGZIP(), []int{4}
}

func (x *Status) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Status) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

type ResponseConfig struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Version       *Version               `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	Status        *Status                `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	Header        []*Header              `protobuf:"bytes,3,rep,name=header,proto3" json:"header,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ResponseConfig) Reset() {
	*x = ResponseConfig{}
	mi := &file_transport_internet_headers_http_config_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResponseConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseConfig) ProtoMessage() {}

func (x *ResponseConfig) ProtoReflect() protoreflect.Message {
	mi := &file_transport_internet_headers_http_config_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseConfig.ProtoReflect.Descriptor instead.
func (*ResponseConfig) Descriptor() ([]byte, []int) {
	return file_transport_internet_headers_http_config_proto_rawDescGZIP(), []int{5}
}

func (x *ResponseConfig) GetVersion() *Version {
	if x != nil {
		return x.Version
	}
	return nil
}

func (x *ResponseConfig) GetStatus() *Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *ResponseConfig) GetHeader() []*Header {
	if x != nil {
		return x.Header
	}
	return nil
}

type Config struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Settings for authenticating requests. If not set, client side will not send
	// authentication header, and server side will bypass authentication.
	Request *RequestConfig `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
	// Settings for authenticating responses. If not set, client side will bypass
	// authentication, and server side will not send authentication header.
	Response      *ResponseConfig `protobuf:"bytes,2,opt,name=response,proto3" json:"response,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Config) Reset() {
	*x = Config{}
	mi := &file_transport_internet_headers_http_config_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_transport_internet_headers_http_config_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_transport_internet_headers_http_config_proto_rawDescGZIP(), []int{6}
}

func (x *Config) GetRequest() *RequestConfig {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *Config) GetResponse() *ResponseConfig {
	if x != nil {
		return x.Response
	}
	return nil
}

var File_transport_internet_headers_http_config_proto protoreflect.FileDescriptor

var file_transport_internet_headers_http_config_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x65, 0x74, 0x2f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x2f, 0x68, 0x74, 0x74,
	0x70, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x24,
	0x78, 0x72, 0x61, 0x79, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x2e,
	0x68, 0x74, 0x74, 0x70, 0x22, 0x32, 0x0a, 0x06, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x1f, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x1e, 0x0a, 0x06, 0x4d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xf6, 0x01, 0x0a, 0x0d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x47, 0x0a, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x78,
	0x72, 0x61, 0x79, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x68,
	0x74, 0x74, 0x70, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x44, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x78, 0x72, 0x61, 0x79, 0x2e, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x2e, 0x68,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x4d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72,
	0x69, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x69, 0x12, 0x44, 0x0a, 0x06,
	0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x78,
	0x72, 0x61, 0x79, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x68,
	0x74, 0x74, 0x70, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x22, 0x34, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x22, 0xe5, 0x01, 0x0a, 0x0e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x47, 0x0a, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x78,
	0x72, 0x61, 0x79, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x68,
	0x74, 0x74, 0x70, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x44, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x78, 0x72, 0x61, 0x79, 0x2e, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x2e, 0x68,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x44, 0x0a, 0x06, 0x68, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x78, 0x72, 0x61,
	0x79, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x65, 0x74, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74,
	0x70, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x22, 0xa9, 0x01, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x4d, 0x0a, 0x07, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x78,
	0x72, 0x61, 0x79, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x68,
	0x74, 0x74, 0x70, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x50, 0x0a, 0x08, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x78,
	0x72, 0x61, 0x79, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x2e, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x68,
	0x74, 0x74, 0x70, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x8e, 0x01, 0x0a,
	0x28, 0x63, 0x6f, 0x6d, 0x2e, 0x78, 0x72, 0x61, 0x79, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70,
	0x6f, 0x72, 0x74, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x2e, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x50, 0x01, 0x5a, 0x39, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x74, 0x6c, 0x73, 0x2f, 0x78, 0x72, 0x61,
	0x79, 0x2d, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74,
	0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x2f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0xaa, 0x02, 0x24, 0x58, 0x72, 0x61, 0x79, 0x2e, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x74,
	0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_transport_internet_headers_http_config_proto_rawDescOnce sync.Once
	file_transport_internet_headers_http_config_proto_rawDescData = file_transport_internet_headers_http_config_proto_rawDesc
)

func file_transport_internet_headers_http_config_proto_rawDescGZIP() []byte {
	file_transport_internet_headers_http_config_proto_rawDescOnce.Do(func() {
		file_transport_internet_headers_http_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_transport_internet_headers_http_config_proto_rawDescData)
	})
	return file_transport_internet_headers_http_config_proto_rawDescData
}

var file_transport_internet_headers_http_config_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_transport_internet_headers_http_config_proto_goTypes = []any{
	(*Header)(nil),         // 0: xray.transport.internet.headers.http.Header
	(*Version)(nil),        // 1: xray.transport.internet.headers.http.Version
	(*Method)(nil),         // 2: xray.transport.internet.headers.http.Method
	(*RequestConfig)(nil),  // 3: xray.transport.internet.headers.http.RequestConfig
	(*Status)(nil),         // 4: xray.transport.internet.headers.http.Status
	(*ResponseConfig)(nil), // 5: xray.transport.internet.headers.http.ResponseConfig
	(*Config)(nil),         // 6: xray.transport.internet.headers.http.Config
}
var file_transport_internet_headers_http_config_proto_depIdxs = []int32{
	1, // 0: xray.transport.internet.headers.http.RequestConfig.version:type_name -> xray.transport.internet.headers.http.Version
	2, // 1: xray.transport.internet.headers.http.RequestConfig.method:type_name -> xray.transport.internet.headers.http.Method
	0, // 2: xray.transport.internet.headers.http.RequestConfig.header:type_name -> xray.transport.internet.headers.http.Header
	1, // 3: xray.transport.internet.headers.http.ResponseConfig.version:type_name -> xray.transport.internet.headers.http.Version
	4, // 4: xray.transport.internet.headers.http.ResponseConfig.status:type_name -> xray.transport.internet.headers.http.Status
	0, // 5: xray.transport.internet.headers.http.ResponseConfig.header:type_name -> xray.transport.internet.headers.http.Header
	3, // 6: xray.transport.internet.headers.http.Config.request:type_name -> xray.transport.internet.headers.http.RequestConfig
	5, // 7: xray.transport.internet.headers.http.Config.response:type_name -> xray.transport.internet.headers.http.ResponseConfig
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_transport_internet_headers_http_config_proto_init() }
func file_transport_internet_headers_http_config_proto_init() {
	if File_transport_internet_headers_http_config_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_transport_internet_headers_http_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_transport_internet_headers_http_config_proto_goTypes,
		DependencyIndexes: file_transport_internet_headers_http_config_proto_depIdxs,
		MessageInfos:      file_transport_internet_headers_http_config_proto_msgTypes,
	}.Build()
	File_transport_internet_headers_http_config_proto = out.File
	file_transport_internet_headers_http_config_proto_rawDesc = nil
	file_transport_internet_headers_http_config_proto_goTypes = nil
	file_transport_internet_headers_http_config_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: client/controller_forward.proto

package client

import (
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

// Forward REST request
type ForwardRequestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Method    string                 `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"` // GET, POST, PUT, DELETE
	Path      string                 `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`     // /api/v1/clients
	Headers   map[string]string      `protobuf:"bytes,3,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Body      []byte                 `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`                         // Request body
	ClientId  string                 `protobuf:"bytes,5,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"` // Target client ID
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *ForwardRequestRequest) Reset() {
	*x = ForwardRequestRequest{}
	mi := &file_client_controller_forward_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ForwardRequestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ForwardRequestRequest) ProtoMessage() {}

func (x *ForwardRequestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_client_controller_forward_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ForwardRequestRequest.ProtoReflect.Descriptor instead.
func (*ForwardRequestRequest) Descriptor() ([]byte, []int) {
	return file_client_controller_forward_proto_rawDescGZIP(), []int{0}
}

func (x *ForwardRequestRequest) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *ForwardRequestRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *ForwardRequestRequest) GetHeaders() map[string]string {
	if x != nil {
		return x.Headers
	}
	return nil
}

func (x *ForwardRequestRequest) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *ForwardRequestRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *ForwardRequestRequest) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

type ForwardRequestResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success    bool              `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Error      string            `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	StatusCode int32             `protobuf:"varint,3,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"` // HTTP status code
	Headers    map[string]string `protobuf:"bytes,4,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Body       []byte            `protobuf:"bytes,5,opt,name=body,proto3" json:"body,omitempty"` // Response body
}

func (x *ForwardRequestResponse) Reset() {
	*x = ForwardRequestResponse{}
	mi := &file_client_controller_forward_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ForwardRequestResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ForwardRequestResponse) ProtoMessage() {}

func (x *ForwardRequestResponse) ProtoReflect() protoreflect.Message {
	mi := &file_client_controller_forward_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ForwardRequestResponse.ProtoReflect.Descriptor instead.
func (*ForwardRequestResponse) Descriptor() ([]byte, []int) {
	return file_client_controller_forward_proto_rawDescGZIP(), []int{1}
}

func (x *ForwardRequestResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *ForwardRequestResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *ForwardRequestResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *ForwardRequestResponse) GetHeaders() map[string]string {
	if x != nil {
		return x.Headers
	}
	return nil
}

func (x *ForwardRequestResponse) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

// Health check
type ForwardHealthCheckRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Service string `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
}

func (x *ForwardHealthCheckRequest) Reset() {
	*x = ForwardHealthCheckRequest{}
	mi := &file_client_controller_forward_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ForwardHealthCheckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ForwardHealthCheckRequest) ProtoMessage() {}

func (x *ForwardHealthCheckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_client_controller_forward_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ForwardHealthCheckRequest.ProtoReflect.Descriptor instead.
func (*ForwardHealthCheckRequest) Descriptor() ([]byte, []int) {
	return file_client_controller_forward_proto_rawDescGZIP(), []int{2}
}

func (x *ForwardHealthCheckRequest) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

type ForwardHealthCheckResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Healthy   bool                   `protobuf:"varint,1,opt,name=healthy,proto3" json:"healthy,omitempty"`
	Message   string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *ForwardHealthCheckResponse) Reset() {
	*x = ForwardHealthCheckResponse{}
	mi := &file_client_controller_forward_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ForwardHealthCheckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ForwardHealthCheckResponse) ProtoMessage() {}

func (x *ForwardHealthCheckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_client_controller_forward_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ForwardHealthCheckResponse.ProtoReflect.Descriptor instead.
func (*ForwardHealthCheckResponse) Descriptor() ([]byte, []int) {
	return file_client_controller_forward_proto_rawDescGZIP(), []int{3}
}

func (x *ForwardHealthCheckResponse) GetHealthy() bool {
	if x != nil {
		return x.Healthy
	}
	return false
}

func (x *ForwardHealthCheckResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ForwardHealthCheckResponse) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

var File_client_controller_forward_proto protoreflect.FileDescriptor

var file_client_controller_forward_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x6c, 0x65, 0x72, 0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x12, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x66, 0x6f,
	0x72, 0x77, 0x61, 0x72, 0x64, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbc, 0x02, 0x0a, 0x15, 0x46, 0x6f, 0x72, 0x77, 0x61,
	0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x50, 0x0a, 0x07,
	0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x36, 0x2e,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61,
	0x72, 0x64, 0x2e, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x62, 0x6f,
	0x64, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x1a, 0x3a, 0x0a, 0x0c, 0x48, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x8c, 0x02, 0x0a, 0x16, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x51, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x37, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x5f,
	0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x48,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x1a, 0x3a, 0x0a, 0x0c, 0x48, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x22, 0x35, 0x0a, 0x19, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x48,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x8a, 0x01, 0x0a, 0x1a,
	0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x68, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x68, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x38,
	0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x32, 0xf1, 0x01, 0x0a, 0x18, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x67, 0x0a, 0x0e, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x46, 0x6f, 0x72,
	0x77, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x5f,
	0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6c,
	0x0a, 0x0b, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x2d, 0x2e,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61,
	0x72, 0x64, 0x2e, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72,
	0x64, 0x2e, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x30, 0x5a, 0x2e,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x4e, 0x61, 0x74, 0x69, 0x76, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x2f, 0x65, 0x6c, 0x63, 0x68,
	0x69, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_client_controller_forward_proto_rawDescOnce sync.Once
	file_client_controller_forward_proto_rawDescData = file_client_controller_forward_proto_rawDesc
)

func file_client_controller_forward_proto_rawDescGZIP() []byte {
	file_client_controller_forward_proto_rawDescOnce.Do(func() {
		file_client_controller_forward_proto_rawDescData = protoimpl.X.CompressGZIP(file_client_controller_forward_proto_rawDescData)
	})
	return file_client_controller_forward_proto_rawDescData
}

var file_client_controller_forward_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_client_controller_forward_proto_goTypes = []any{
	(*ForwardRequestRequest)(nil),      // 0: controller_forward.ForwardRequestRequest
	(*ForwardRequestResponse)(nil),     // 1: controller_forward.ForwardRequestResponse
	(*ForwardHealthCheckRequest)(nil),  // 2: controller_forward.ForwardHealthCheckRequest
	(*ForwardHealthCheckResponse)(nil), // 3: controller_forward.ForwardHealthCheckResponse
	nil,                                // 4: controller_forward.ForwardRequestRequest.HeadersEntry
	nil,                                // 5: controller_forward.ForwardRequestResponse.HeadersEntry
	(*timestamppb.Timestamp)(nil),      // 6: google.protobuf.Timestamp
}
var file_client_controller_forward_proto_depIdxs = []int32{
	4, // 0: controller_forward.ForwardRequestRequest.headers:type_name -> controller_forward.ForwardRequestRequest.HeadersEntry
	6, // 1: controller_forward.ForwardRequestRequest.timestamp:type_name -> google.protobuf.Timestamp
	5, // 2: controller_forward.ForwardRequestResponse.headers:type_name -> controller_forward.ForwardRequestResponse.HeadersEntry
	6, // 3: controller_forward.ForwardHealthCheckResponse.timestamp:type_name -> google.protobuf.Timestamp
	0, // 4: controller_forward.ControllerForwardService.ForwardRequest:input_type -> controller_forward.ForwardRequestRequest
	2, // 5: controller_forward.ControllerForwardService.HealthCheck:input_type -> controller_forward.ForwardHealthCheckRequest
	1, // 6: controller_forward.ControllerForwardService.ForwardRequest:output_type -> controller_forward.ForwardRequestResponse
	3, // 7: controller_forward.ControllerForwardService.HealthCheck:output_type -> controller_forward.ForwardHealthCheckResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_client_controller_forward_proto_init() }
func file_client_controller_forward_proto_init() {
	if File_client_controller_forward_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_client_controller_forward_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_client_controller_forward_proto_goTypes,
		DependencyIndexes: file_client_controller_forward_proto_depIdxs,
		MessageInfos:      file_client_controller_forward_proto_msgTypes,
	}.Build()
	File_client_controller_forward_proto = out.File
	file_client_controller_forward_proto_rawDesc = nil
	file_client_controller_forward_proto_goTypes = nil
	file_client_controller_forward_proto_depIdxs = nil
}

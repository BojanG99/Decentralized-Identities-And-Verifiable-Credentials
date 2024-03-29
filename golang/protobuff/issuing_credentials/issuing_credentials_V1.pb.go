// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: issuing_credentials/issuing_credentials_V1.proto

package verifiable_credentials

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

type ServerMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Message:
	//
	//	*ServerMessage_Sh
	//	*ServerMessage_Sp
	//	*ServerMessage_Vc
	//	*ServerMessage_Err
	Message isServerMessage_Message `protobuf_oneof:"message"`
}

func (x *ServerMessage) Reset() {
	*x = ServerMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_issuing_credentials_issuing_credentials_V1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerMessage) ProtoMessage() {}

func (x *ServerMessage) ProtoReflect() protoreflect.Message {
	mi := &file_issuing_credentials_issuing_credentials_V1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerMessage.ProtoReflect.Descriptor instead.
func (*ServerMessage) Descriptor() ([]byte, []int) {
	return file_issuing_credentials_issuing_credentials_V1_proto_rawDescGZIP(), []int{0}
}

func (m *ServerMessage) GetMessage() isServerMessage_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (x *ServerMessage) GetSh() *ServerHello {
	if x, ok := x.GetMessage().(*ServerMessage_Sh); ok {
		return x.Sh
	}
	return nil
}

func (x *ServerMessage) GetSp() *ServerClaimsProposal {
	if x, ok := x.GetMessage().(*ServerMessage_Sp); ok {
		return x.Sp
	}
	return nil
}

func (x *ServerMessage) GetVc() *VerifiableCredential {
	if x, ok := x.GetMessage().(*ServerMessage_Vc); ok {
		return x.Vc
	}
	return nil
}

func (x *ServerMessage) GetErr() *Error {
	if x, ok := x.GetMessage().(*ServerMessage_Err); ok {
		return x.Err
	}
	return nil
}

type isServerMessage_Message interface {
	isServerMessage_Message()
}

type ServerMessage_Sh struct {
	Sh *ServerHello `protobuf:"bytes,1,opt,name=sh,proto3,oneof"`
}

type ServerMessage_Sp struct {
	Sp *ServerClaimsProposal `protobuf:"bytes,2,opt,name=sp,proto3,oneof"`
}

type ServerMessage_Vc struct {
	Vc *VerifiableCredential `protobuf:"bytes,3,opt,name=vc,proto3,oneof"`
}

type ServerMessage_Err struct {
	Err *Error `protobuf:"bytes,4,opt,name=err,proto3,oneof"`
}

func (*ServerMessage_Sh) isServerMessage_Message() {}

func (*ServerMessage_Sp) isServerMessage_Message() {}

func (*ServerMessage_Vc) isServerMessage_Message() {}

func (*ServerMessage_Err) isServerMessage_Message() {}

type ClientMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Message:
	//
	//	*ClientMessage_Ch
	//	*ClientMessage_Ca
	//	*ClientMessage_Cr
	//	*ClientMessage_Err
	Message isClientMessage_Message `protobuf_oneof:"message"`
}

func (x *ClientMessage) Reset() {
	*x = ClientMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_issuing_credentials_issuing_credentials_V1_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientMessage) ProtoMessage() {}

func (x *ClientMessage) ProtoReflect() protoreflect.Message {
	mi := &file_issuing_credentials_issuing_credentials_V1_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientMessage.ProtoReflect.Descriptor instead.
func (*ClientMessage) Descriptor() ([]byte, []int) {
	return file_issuing_credentials_issuing_credentials_V1_proto_rawDescGZIP(), []int{1}
}

func (m *ClientMessage) GetMessage() isClientMessage_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (x *ClientMessage) GetCh() *ClientHello {
	if x, ok := x.GetMessage().(*ClientMessage_Ch); ok {
		return x.Ch
	}
	return nil
}

func (x *ClientMessage) GetCa() *ClientAuthentication {
	if x, ok := x.GetMessage().(*ClientMessage_Ca); ok {
		return x.Ca
	}
	return nil
}

func (x *ClientMessage) GetCr() *ClientResponse {
	if x, ok := x.GetMessage().(*ClientMessage_Cr); ok {
		return x.Cr
	}
	return nil
}

func (x *ClientMessage) GetErr() *Error {
	if x, ok := x.GetMessage().(*ClientMessage_Err); ok {
		return x.Err
	}
	return nil
}

type isClientMessage_Message interface {
	isClientMessage_Message()
}

type ClientMessage_Ch struct {
	Ch *ClientHello `protobuf:"bytes,1,opt,name=ch,proto3,oneof"`
}

type ClientMessage_Ca struct {
	Ca *ClientAuthentication `protobuf:"bytes,2,opt,name=ca,proto3,oneof"`
}

type ClientMessage_Cr struct {
	Cr *ClientResponse `protobuf:"bytes,3,opt,name=cr,proto3,oneof"`
}

type ClientMessage_Err struct {
	Err *Error `protobuf:"bytes,4,opt,name=err,proto3,oneof"`
}

func (*ClientMessage_Ch) isClientMessage_Message() {}

func (*ClientMessage_Ca) isClientMessage_Message() {}

func (*ClientMessage_Cr) isClientMessage_Message() {}

func (*ClientMessage_Err) isClientMessage_Message() {}

type ClientHello struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DidString       string `protobuf:"bytes,1,opt,name=didString,proto3" json:"didString,omitempty"`
	ServerChallange string `protobuf:"bytes,2,opt,name=server_challange,json=serverChallange,proto3" json:"server_challange,omitempty"`
	RpcKey          string `protobuf:"bytes,3,opt,name=rpc_key,json=rpcKey,proto3" json:"rpc_key,omitempty"`
}

func (x *ClientHello) Reset() {
	*x = ClientHello{}
	if protoimpl.UnsafeEnabled {
		mi := &file_issuing_credentials_issuing_credentials_V1_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientHello) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientHello) ProtoMessage() {}

func (x *ClientHello) ProtoReflect() protoreflect.Message {
	mi := &file_issuing_credentials_issuing_credentials_V1_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientHello.ProtoReflect.Descriptor instead.
func (*ClientHello) Descriptor() ([]byte, []int) {
	return file_issuing_credentials_issuing_credentials_V1_proto_rawDescGZIP(), []int{2}
}

func (x *ClientHello) GetDidString() string {
	if x != nil {
		return x.DidString
	}
	return ""
}

func (x *ClientHello) GetServerChallange() string {
	if x != nil {
		return x.ServerChallange
	}
	return ""
}

func (x *ClientHello) GetRpcKey() string {
	if x != nil {
		return x.RpcKey
	}
	return ""
}

type ServerHello struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DidString          string `protobuf:"bytes,1,opt,name=didString,proto3" json:"didString,omitempty"`
	ClientChallange    string `protobuf:"bytes,2,opt,name=client_challange,json=clientChallange,proto3" json:"client_challange,omitempty"`
	CompletedChallange string `protobuf:"bytes,3,opt,name=completedChallange,proto3" json:"completedChallange,omitempty"`
	KeyUrl             string `protobuf:"bytes,4,opt,name=keyUrl,proto3" json:"keyUrl,omitempty"`
}

func (x *ServerHello) Reset() {
	*x = ServerHello{}
	if protoimpl.UnsafeEnabled {
		mi := &file_issuing_credentials_issuing_credentials_V1_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerHello) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerHello) ProtoMessage() {}

func (x *ServerHello) ProtoReflect() protoreflect.Message {
	mi := &file_issuing_credentials_issuing_credentials_V1_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerHello.ProtoReflect.Descriptor instead.
func (*ServerHello) Descriptor() ([]byte, []int) {
	return file_issuing_credentials_issuing_credentials_V1_proto_rawDescGZIP(), []int{3}
}

func (x *ServerHello) GetDidString() string {
	if x != nil {
		return x.DidString
	}
	return ""
}

func (x *ServerHello) GetClientChallange() string {
	if x != nil {
		return x.ClientChallange
	}
	return ""
}

func (x *ServerHello) GetCompletedChallange() string {
	if x != nil {
		return x.CompletedChallange
	}
	return ""
}

func (x *ServerHello) GetKeyUrl() string {
	if x != nil {
		return x.KeyUrl
	}
	return ""
}

type ClientAuthentication struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CompletedChallange string `protobuf:"bytes,1,opt,name=completed_challange,json=completedChallange,proto3" json:"completed_challange,omitempty"`
	KeyUrl             string `protobuf:"bytes,2,opt,name=key_url,json=keyUrl,proto3" json:"key_url,omitempty"`
	Pin                string `protobuf:"bytes,3,opt,name=pin,proto3" json:"pin,omitempty"`
}

func (x *ClientAuthentication) Reset() {
	*x = ClientAuthentication{}
	if protoimpl.UnsafeEnabled {
		mi := &file_issuing_credentials_issuing_credentials_V1_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientAuthentication) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientAuthentication) ProtoMessage() {}

func (x *ClientAuthentication) ProtoReflect() protoreflect.Message {
	mi := &file_issuing_credentials_issuing_credentials_V1_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientAuthentication.ProtoReflect.Descriptor instead.
func (*ClientAuthentication) Descriptor() ([]byte, []int) {
	return file_issuing_credentials_issuing_credentials_V1_proto_rawDescGZIP(), []int{4}
}

func (x *ClientAuthentication) GetCompletedChallange() string {
	if x != nil {
		return x.CompletedChallange
	}
	return ""
}

func (x *ClientAuthentication) GetKeyUrl() string {
	if x != nil {
		return x.KeyUrl
	}
	return ""
}

func (x *ClientAuthentication) GetPin() string {
	if x != nil {
		return x.Pin
	}
	return ""
}

type ServerClaimsProposal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type        []string          `protobuf:"bytes,1,rep,name=type,proto3" json:"type,omitempty"`
	Description string            `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Claims      map[string]string `protobuf:"bytes,3,rep,name=claims,proto3" json:"claims,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ServerClaimsProposal) Reset() {
	*x = ServerClaimsProposal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_issuing_credentials_issuing_credentials_V1_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerClaimsProposal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerClaimsProposal) ProtoMessage() {}

func (x *ServerClaimsProposal) ProtoReflect() protoreflect.Message {
	mi := &file_issuing_credentials_issuing_credentials_V1_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerClaimsProposal.ProtoReflect.Descriptor instead.
func (*ServerClaimsProposal) Descriptor() ([]byte, []int) {
	return file_issuing_credentials_issuing_credentials_V1_proto_rawDescGZIP(), []int{5}
}

func (x *ServerClaimsProposal) GetType() []string {
	if x != nil {
		return x.Type
	}
	return nil
}

func (x *ServerClaimsProposal) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ServerClaimsProposal) GetClaims() map[string]string {
	if x != nil {
		return x.Claims
	}
	return nil
}

type ClientResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AcceptCredentials bool `protobuf:"varint,1,opt,name=acceptCredentials,proto3" json:"acceptCredentials,omitempty"`
}

func (x *ClientResponse) Reset() {
	*x = ClientResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_issuing_credentials_issuing_credentials_V1_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientResponse) ProtoMessage() {}

func (x *ClientResponse) ProtoReflect() protoreflect.Message {
	mi := &file_issuing_credentials_issuing_credentials_V1_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientResponse.ProtoReflect.Descriptor instead.
func (*ClientResponse) Descriptor() ([]byte, []int) {
	return file_issuing_credentials_issuing_credentials_V1_proto_rawDescGZIP(), []int{6}
}

func (x *ClientResponse) GetAcceptCredentials() bool {
	if x != nil {
		return x.AcceptCredentials
	}
	return false
}

type VerifiableCredential struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type         []string          `protobuf:"bytes,2,rep,name=type,proto3" json:"type,omitempty"`
	Issuer       string            `protobuf:"bytes,3,opt,name=issuer,proto3" json:"issuer,omitempty"`
	Subject      string            `protobuf:"bytes,4,opt,name=subject,proto3" json:"subject,omitempty"`
	IssuanceDate string            `protobuf:"bytes,5,opt,name=issuanceDate,proto3" json:"issuanceDate,omitempty"`
	Description  string            `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	Claims       map[string]string `protobuf:"bytes,7,rep,name=claims,proto3" json:"claims,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Proof        map[string]string `protobuf:"bytes,8,rep,name=proof,proto3" json:"proof,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *VerifiableCredential) Reset() {
	*x = VerifiableCredential{}
	if protoimpl.UnsafeEnabled {
		mi := &file_issuing_credentials_issuing_credentials_V1_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifiableCredential) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifiableCredential) ProtoMessage() {}

func (x *VerifiableCredential) ProtoReflect() protoreflect.Message {
	mi := &file_issuing_credentials_issuing_credentials_V1_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifiableCredential.ProtoReflect.Descriptor instead.
func (*VerifiableCredential) Descriptor() ([]byte, []int) {
	return file_issuing_credentials_issuing_credentials_V1_proto_rawDescGZIP(), []int{7}
}

func (x *VerifiableCredential) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *VerifiableCredential) GetType() []string {
	if x != nil {
		return x.Type
	}
	return nil
}

func (x *VerifiableCredential) GetIssuer() string {
	if x != nil {
		return x.Issuer
	}
	return ""
}

func (x *VerifiableCredential) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *VerifiableCredential) GetIssuanceDate() string {
	if x != nil {
		return x.IssuanceDate
	}
	return ""
}

func (x *VerifiableCredential) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *VerifiableCredential) GetClaims() map[string]string {
	if x != nil {
		return x.Claims
	}
	return nil
}

func (x *VerifiableCredential) GetProof() map[string]string {
	if x != nil {
		return x.Proof
	}
	return nil
}

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrorMessage string `protobuf:"bytes,1,opt,name=errorMessage,proto3" json:"errorMessage,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_issuing_credentials_issuing_credentials_V1_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_issuing_credentials_issuing_credentials_V1_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_issuing_credentials_issuing_credentials_V1_proto_rawDescGZIP(), []int{8}
}

func (x *Error) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

var File_issuing_credentials_issuing_credentials_V1_proto protoreflect.FileDescriptor

var file_issuing_credentials_issuing_credentials_V1_proto_rawDesc = []byte{
	0x0a, 0x30, 0x69, 0x73, 0x73, 0x75, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x61, 0x6c, 0x73, 0x2f, 0x69, 0x73, 0x73, 0x75, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x72,
	0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x5f, 0x56, 0x31, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1e, 0x69, 0x73, 0x73, 0x75, 0x69, 0x6e, 0x67, 0x5f, 0x76, 0x65, 0x72, 0x69,
	0x66, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61,
	0x6c, 0x73, 0x22, 0xa4, 0x02, 0x0a, 0x0d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x3d, 0x0a, 0x02, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2b, 0x2e, 0x69, 0x73, 0x73, 0x75, 0x69, 0x6e, 0x67, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66,
	0x69, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c,
	0x73, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x48, 0x00, 0x52,
	0x02, 0x73, 0x68, 0x12, 0x46, 0x0a, 0x02, 0x73, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x34, 0x2e, 0x69, 0x73, 0x73, 0x75, 0x69, 0x6e, 0x67, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69,
	0x61, 0x62, 0x6c, 0x65, 0x5f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73,
	0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x73, 0x50, 0x72, 0x6f,
	0x70, 0x6f, 0x73, 0x61, 0x6c, 0x48, 0x00, 0x52, 0x02, 0x73, 0x70, 0x12, 0x46, 0x0a, 0x02, 0x76,
	0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x69, 0x73, 0x73, 0x75, 0x69, 0x6e,
	0x67, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x63, 0x72, 0x65,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x61,
	0x62, 0x6c, 0x65, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x48, 0x00, 0x52,
	0x02, 0x76, 0x63, 0x12, 0x39, 0x0a, 0x03, 0x65, 0x72, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x25, 0x2e, 0x69, 0x73, 0x73, 0x75, 0x69, 0x6e, 0x67, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66,
	0x69, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c,
	0x73, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x48, 0x00, 0x52, 0x03, 0x65, 0x72, 0x72, 0x42, 0x09,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x9e, 0x02, 0x0a, 0x0d, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x3d, 0x0a, 0x02, 0x63,
	0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x69, 0x73, 0x73, 0x75, 0x69, 0x6e,
	0x67, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x63, 0x72, 0x65,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x48,
	0x65, 0x6c, 0x6c, 0x6f, 0x48, 0x00, 0x52, 0x02, 0x63, 0x68, 0x12, 0x46, 0x0a, 0x02, 0x63, 0x61,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x69, 0x73, 0x73, 0x75, 0x69, 0x6e, 0x67,
	0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x63, 0x72, 0x65, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x41, 0x75,
	0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x02,
	0x63, 0x61, 0x12, 0x40, 0x0a, 0x02, 0x63, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e,
	0x2e, 0x69, 0x73, 0x73, 0x75, 0x69, 0x6e, 0x67, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x61,
	0x62, 0x6c, 0x65, 0x5f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x2e,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00,
	0x52, 0x02, 0x63, 0x72, 0x12, 0x39, 0x0a, 0x03, 0x65, 0x72, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x25, 0x2e, 0x69, 0x73, 0x73, 0x75, 0x69, 0x6e, 0x67, 0x5f, 0x76, 0x65, 0x72, 0x69,
	0x66, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61,
	0x6c, 0x73, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x48, 0x00, 0x52, 0x03, 0x65, 0x72, 0x72, 0x42,
	0x09, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x6f, 0x0a, 0x0b, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x69, 0x64,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x69,
	0x64, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x29, 0x0a, 0x10, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x5f, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x61, 0x6e,
	0x67, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x70, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x70, 0x63, 0x4b, 0x65, 0x79, 0x22, 0x9e, 0x01, 0x0a, 0x0b,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x1c, 0x0a, 0x09, 0x64,
	0x69, 0x64, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x64, 0x69, 0x64, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x29, 0x0a, 0x10, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x5f, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x68, 0x61, 0x6c, 0x6c,
	0x61, 0x6e, 0x67, 0x65, 0x12, 0x2e, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65,
	0x64, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x12, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x43, 0x68, 0x61, 0x6c, 0x6c,
	0x61, 0x6e, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6b, 0x65, 0x79, 0x55, 0x72, 0x6c, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6b, 0x65, 0x79, 0x55, 0x72, 0x6c, 0x22, 0x72, 0x0a, 0x14,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2f, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65,
	0x64, 0x5f, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x12, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x43, 0x68, 0x61, 0x6c,
	0x6c, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x6b, 0x65, 0x79, 0x5f, 0x75, 0x72, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6b, 0x65, 0x79, 0x55, 0x72, 0x6c, 0x12, 0x10,
	0x0a, 0x03, 0x70, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x70, 0x69, 0x6e,
	0x22, 0xe1, 0x01, 0x0a, 0x14, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6c, 0x61, 0x69, 0x6d,
	0x73, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x58, 0x0a, 0x06, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x40, 0x2e, 0x69, 0x73, 0x73, 0x75, 0x69, 0x6e, 0x67, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69,
	0x61, 0x62, 0x6c, 0x65, 0x5f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73,
	0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x73, 0x50, 0x72, 0x6f,
	0x70, 0x6f, 0x73, 0x61, 0x6c, 0x2e, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x06, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x43, 0x6c, 0x61,
	0x69, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x22, 0x3e, 0x0a, 0x0e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x11, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74,
	0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x11, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x61, 0x6c, 0x73, 0x22, 0xd8, 0x03, 0x0a, 0x14, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x61,
	0x62, 0x6c, 0x65, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x69, 0x73, 0x73, 0x75, 0x61, 0x6e, 0x63, 0x65, 0x44,
	0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x73, 0x73, 0x75, 0x61,
	0x6e, 0x63, 0x65, 0x44, 0x61, 0x74, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x58, 0x0a, 0x06, 0x63, 0x6c, 0x61,
	0x69, 0x6d, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x40, 0x2e, 0x69, 0x73, 0x73, 0x75,
	0x69, 0x6e, 0x67, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x63,
	0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66,
	0x69, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x2e,
	0x43, 0x6c, 0x61, 0x69, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x63, 0x6c, 0x61,
	0x69, 0x6d, 0x73, 0x12, 0x55, 0x0a, 0x05, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x18, 0x08, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x69, 0x73, 0x73, 0x75, 0x69, 0x6e, 0x67, 0x5f, 0x76, 0x65, 0x72,
	0x69, 0x66, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x61, 0x6c, 0x73, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x72,
	0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x2e, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x05, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x1a, 0x39, 0x0a, 0x0b, 0x43, 0x6c,
	0x61, 0x69, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x38, 0x0a, 0x0a, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22,
	0x2b, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x22, 0x0a, 0x0c, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x92, 0x01, 0x0a,
	0x1b, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x72, 0x65, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x73, 0x0a, 0x0f,
	0x49, 0x73, 0x73, 0x75, 0x65, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x12,
	0x2d, 0x2e, 0x69, 0x73, 0x73, 0x75, 0x69, 0x6e, 0x67, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69,
	0x61, 0x62, 0x6c, 0x65, 0x5f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73,
	0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x2d,
	0x2e, 0x69, 0x73, 0x73, 0x75, 0x69, 0x6e, 0x67, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x61,
	0x62, 0x6c, 0x65, 0x5f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x2e,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x28, 0x01, 0x30,
	0x01, 0x42, 0x6a, 0x5a, 0x68, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x42, 0x6f, 0x6a, 0x61, 0x6e, 0x47, 0x39, 0x39, 0x2f, 0x44, 0x65, 0x63, 0x65, 0x6e, 0x74, 0x72,
	0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x2d, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x73,
	0x65, 0x2d, 0x41, 0x6e, 0x64, 0x2d, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x61, 0x62, 0x6c, 0x65,
	0x2d, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x66, 0x3b, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x61, 0x62, 0x6c,
	0x65, 0x5f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_issuing_credentials_issuing_credentials_V1_proto_rawDescOnce sync.Once
	file_issuing_credentials_issuing_credentials_V1_proto_rawDescData = file_issuing_credentials_issuing_credentials_V1_proto_rawDesc
)

func file_issuing_credentials_issuing_credentials_V1_proto_rawDescGZIP() []byte {
	file_issuing_credentials_issuing_credentials_V1_proto_rawDescOnce.Do(func() {
		file_issuing_credentials_issuing_credentials_V1_proto_rawDescData = protoimpl.X.CompressGZIP(file_issuing_credentials_issuing_credentials_V1_proto_rawDescData)
	})
	return file_issuing_credentials_issuing_credentials_V1_proto_rawDescData
}

var file_issuing_credentials_issuing_credentials_V1_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_issuing_credentials_issuing_credentials_V1_proto_goTypes = []interface{}{
	(*ServerMessage)(nil),        // 0: issuing_verifiable_credentials.ServerMessage
	(*ClientMessage)(nil),        // 1: issuing_verifiable_credentials.ClientMessage
	(*ClientHello)(nil),          // 2: issuing_verifiable_credentials.ClientHello
	(*ServerHello)(nil),          // 3: issuing_verifiable_credentials.ServerHello
	(*ClientAuthentication)(nil), // 4: issuing_verifiable_credentials.ClientAuthentication
	(*ServerClaimsProposal)(nil), // 5: issuing_verifiable_credentials.ServerClaimsProposal
	(*ClientResponse)(nil),       // 6: issuing_verifiable_credentials.ClientResponse
	(*VerifiableCredential)(nil), // 7: issuing_verifiable_credentials.VerifiableCredential
	(*Error)(nil),                // 8: issuing_verifiable_credentials.Error
	nil,                          // 9: issuing_verifiable_credentials.ServerClaimsProposal.ClaimsEntry
	nil,                          // 10: issuing_verifiable_credentials.VerifiableCredential.ClaimsEntry
	nil,                          // 11: issuing_verifiable_credentials.VerifiableCredential.ProofEntry
}
var file_issuing_credentials_issuing_credentials_V1_proto_depIdxs = []int32{
	3,  // 0: issuing_verifiable_credentials.ServerMessage.sh:type_name -> issuing_verifiable_credentials.ServerHello
	5,  // 1: issuing_verifiable_credentials.ServerMessage.sp:type_name -> issuing_verifiable_credentials.ServerClaimsProposal
	7,  // 2: issuing_verifiable_credentials.ServerMessage.vc:type_name -> issuing_verifiable_credentials.VerifiableCredential
	8,  // 3: issuing_verifiable_credentials.ServerMessage.err:type_name -> issuing_verifiable_credentials.Error
	2,  // 4: issuing_verifiable_credentials.ClientMessage.ch:type_name -> issuing_verifiable_credentials.ClientHello
	4,  // 5: issuing_verifiable_credentials.ClientMessage.ca:type_name -> issuing_verifiable_credentials.ClientAuthentication
	6,  // 6: issuing_verifiable_credentials.ClientMessage.cr:type_name -> issuing_verifiable_credentials.ClientResponse
	8,  // 7: issuing_verifiable_credentials.ClientMessage.err:type_name -> issuing_verifiable_credentials.Error
	9,  // 8: issuing_verifiable_credentials.ServerClaimsProposal.claims:type_name -> issuing_verifiable_credentials.ServerClaimsProposal.ClaimsEntry
	10, // 9: issuing_verifiable_credentials.VerifiableCredential.claims:type_name -> issuing_verifiable_credentials.VerifiableCredential.ClaimsEntry
	11, // 10: issuing_verifiable_credentials.VerifiableCredential.proof:type_name -> issuing_verifiable_credentials.VerifiableCredential.ProofEntry
	1,  // 11: issuing_verifiable_credentials.VerifiableCredentialService.IssueCredential:input_type -> issuing_verifiable_credentials.ClientMessage
	0,  // 12: issuing_verifiable_credentials.VerifiableCredentialService.IssueCredential:output_type -> issuing_verifiable_credentials.ServerMessage
	12, // [12:13] is the sub-list for method output_type
	11, // [11:12] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_issuing_credentials_issuing_credentials_V1_proto_init() }
func file_issuing_credentials_issuing_credentials_V1_proto_init() {
	if File_issuing_credentials_issuing_credentials_V1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_issuing_credentials_issuing_credentials_V1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerMessage); i {
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
		file_issuing_credentials_issuing_credentials_V1_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientMessage); i {
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
		file_issuing_credentials_issuing_credentials_V1_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientHello); i {
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
		file_issuing_credentials_issuing_credentials_V1_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerHello); i {
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
		file_issuing_credentials_issuing_credentials_V1_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientAuthentication); i {
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
		file_issuing_credentials_issuing_credentials_V1_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerClaimsProposal); i {
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
		file_issuing_credentials_issuing_credentials_V1_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientResponse); i {
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
		file_issuing_credentials_issuing_credentials_V1_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifiableCredential); i {
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
		file_issuing_credentials_issuing_credentials_V1_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
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
	file_issuing_credentials_issuing_credentials_V1_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*ServerMessage_Sh)(nil),
		(*ServerMessage_Sp)(nil),
		(*ServerMessage_Vc)(nil),
		(*ServerMessage_Err)(nil),
	}
	file_issuing_credentials_issuing_credentials_V1_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*ClientMessage_Ch)(nil),
		(*ClientMessage_Ca)(nil),
		(*ClientMessage_Cr)(nil),
		(*ClientMessage_Err)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_issuing_credentials_issuing_credentials_V1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_issuing_credentials_issuing_credentials_V1_proto_goTypes,
		DependencyIndexes: file_issuing_credentials_issuing_credentials_V1_proto_depIdxs,
		MessageInfos:      file_issuing_credentials_issuing_credentials_V1_proto_msgTypes,
	}.Build()
	File_issuing_credentials_issuing_credentials_V1_proto = out.File
	file_issuing_credentials_issuing_credentials_V1_proto_rawDesc = nil
	file_issuing_credentials_issuing_credentials_V1_proto_goTypes = nil
	file_issuing_credentials_issuing_credentials_V1_proto_depIdxs = nil
}

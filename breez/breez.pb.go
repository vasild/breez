// Code generated by protoc-gen-go. DO NOT EDIT.
// source: breez.proto

/*
Package breez is a generated protocol buffer package.

It is generated from these files:
	breez.proto

It has these top-level messages:
	OpenChannelRequest
	OpenChannelReply
	UpdateChannelPolicyRequest
	UpdateChannelPolicyReply
	AddFundInitRequest
	AddFundInitReply
	AddFundStatusRequest
	AddFundStatusReply
	RemoveFundRequest
	RemoveFundReply
	RedeemRemovedFundsRequest
	RedeemRemovedFundsReply
	GetSwapPaymentRequest
	GetSwapPaymentReply
	MempoolRegisterRequest
	MempoolRegisterReply
	PingRequest
	PingReply
*/
package breez

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type OpenChannelRequest struct {
	PubKey string `protobuf:"bytes,1,opt,name=pubKey" json:"pubKey,omitempty"`
}

func (m *OpenChannelRequest) Reset()                    { *m = OpenChannelRequest{} }
func (m *OpenChannelRequest) String() string            { return proto.CompactTextString(m) }
func (*OpenChannelRequest) ProtoMessage()               {}
func (*OpenChannelRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *OpenChannelRequest) GetPubKey() string {
	if m != nil {
		return m.PubKey
	}
	return ""
}

type OpenChannelReply struct {
}

func (m *OpenChannelReply) Reset()                    { *m = OpenChannelReply{} }
func (m *OpenChannelReply) String() string            { return proto.CompactTextString(m) }
func (*OpenChannelReply) ProtoMessage()               {}
func (*OpenChannelReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type UpdateChannelPolicyRequest struct {
	PubKey string `protobuf:"bytes,1,opt,name=pubKey" json:"pubKey,omitempty"`
}

func (m *UpdateChannelPolicyRequest) Reset()                    { *m = UpdateChannelPolicyRequest{} }
func (m *UpdateChannelPolicyRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateChannelPolicyRequest) ProtoMessage()               {}
func (*UpdateChannelPolicyRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *UpdateChannelPolicyRequest) GetPubKey() string {
	if m != nil {
		return m.PubKey
	}
	return ""
}

type UpdateChannelPolicyReply struct {
}

func (m *UpdateChannelPolicyReply) Reset()                    { *m = UpdateChannelPolicyReply{} }
func (m *UpdateChannelPolicyReply) String() string            { return proto.CompactTextString(m) }
func (*UpdateChannelPolicyReply) ProtoMessage()               {}
func (*UpdateChannelPolicyReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type AddFundInitRequest struct {
	NodeID            string `protobuf:"bytes,1,opt,name=nodeID" json:"nodeID,omitempty"`
	NotificationToken string `protobuf:"bytes,2,opt,name=notificationToken" json:"notificationToken,omitempty"`
	Pubkey            []byte `protobuf:"bytes,3,opt,name=pubkey,proto3" json:"pubkey,omitempty"`
	Hash              []byte `protobuf:"bytes,4,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (m *AddFundInitRequest) Reset()                    { *m = AddFundInitRequest{} }
func (m *AddFundInitRequest) String() string            { return proto.CompactTextString(m) }
func (*AddFundInitRequest) ProtoMessage()               {}
func (*AddFundInitRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *AddFundInitRequest) GetNodeID() string {
	if m != nil {
		return m.NodeID
	}
	return ""
}

func (m *AddFundInitRequest) GetNotificationToken() string {
	if m != nil {
		return m.NotificationToken
	}
	return ""
}

func (m *AddFundInitRequest) GetPubkey() []byte {
	if m != nil {
		return m.Pubkey
	}
	return nil
}

func (m *AddFundInitRequest) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

type AddFundInitReply struct {
	Address           string `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
	Pubkey            []byte `protobuf:"bytes,2,opt,name=pubkey,proto3" json:"pubkey,omitempty"`
	LockHeight        int64  `protobuf:"varint,3,opt,name=lockHeight" json:"lockHeight,omitempty"`
	MaxAllowedDeposit int64  `protobuf:"varint,4,opt,name=maxAllowedDeposit" json:"maxAllowedDeposit,omitempty"`
	ErrorMessage      string `protobuf:"bytes,5,opt,name=errorMessage" json:"errorMessage,omitempty"`
}

func (m *AddFundInitReply) Reset()                    { *m = AddFundInitReply{} }
func (m *AddFundInitReply) String() string            { return proto.CompactTextString(m) }
func (*AddFundInitReply) ProtoMessage()               {}
func (*AddFundInitReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *AddFundInitReply) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *AddFundInitReply) GetPubkey() []byte {
	if m != nil {
		return m.Pubkey
	}
	return nil
}

func (m *AddFundInitReply) GetLockHeight() int64 {
	if m != nil {
		return m.LockHeight
	}
	return 0
}

func (m *AddFundInitReply) GetMaxAllowedDeposit() int64 {
	if m != nil {
		return m.MaxAllowedDeposit
	}
	return 0
}

func (m *AddFundInitReply) GetErrorMessage() string {
	if m != nil {
		return m.ErrorMessage
	}
	return ""
}

type AddFundStatusRequest struct {
	Addresses         []string `protobuf:"bytes,1,rep,name=addresses" json:"addresses,omitempty"`
	NotificationToken string   `protobuf:"bytes,2,opt,name=notificationToken" json:"notificationToken,omitempty"`
}

func (m *AddFundStatusRequest) Reset()                    { *m = AddFundStatusRequest{} }
func (m *AddFundStatusRequest) String() string            { return proto.CompactTextString(m) }
func (*AddFundStatusRequest) ProtoMessage()               {}
func (*AddFundStatusRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *AddFundStatusRequest) GetAddresses() []string {
	if m != nil {
		return m.Addresses
	}
	return nil
}

func (m *AddFundStatusRequest) GetNotificationToken() string {
	if m != nil {
		return m.NotificationToken
	}
	return ""
}

type AddFundStatusReply struct {
	Statuses map[string]*AddFundStatusReply_AddressStatus `protobuf:"bytes,1,rep,name=statuses" json:"statuses,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *AddFundStatusReply) Reset()                    { *m = AddFundStatusReply{} }
func (m *AddFundStatusReply) String() string            { return proto.CompactTextString(m) }
func (*AddFundStatusReply) ProtoMessage()               {}
func (*AddFundStatusReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *AddFundStatusReply) GetStatuses() map[string]*AddFundStatusReply_AddressStatus {
	if m != nil {
		return m.Statuses
	}
	return nil
}

type AddFundStatusReply_AddressStatus struct {
	Tx        string `protobuf:"bytes,1,opt,name=tx" json:"tx,omitempty"`
	Amount    int64  `protobuf:"varint,2,opt,name=amount" json:"amount,omitempty"`
	Confirmed bool   `protobuf:"varint,3,opt,name=confirmed" json:"confirmed,omitempty"`
	BlockHash string `protobuf:"bytes,4,opt,name=blockHash" json:"blockHash,omitempty"`
}

func (m *AddFundStatusReply_AddressStatus) Reset()         { *m = AddFundStatusReply_AddressStatus{} }
func (m *AddFundStatusReply_AddressStatus) String() string { return proto.CompactTextString(m) }
func (*AddFundStatusReply_AddressStatus) ProtoMessage()    {}
func (*AddFundStatusReply_AddressStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{7, 0}
}

func (m *AddFundStatusReply_AddressStatus) GetTx() string {
	if m != nil {
		return m.Tx
	}
	return ""
}

func (m *AddFundStatusReply_AddressStatus) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *AddFundStatusReply_AddressStatus) GetConfirmed() bool {
	if m != nil {
		return m.Confirmed
	}
	return false
}

func (m *AddFundStatusReply_AddressStatus) GetBlockHash() string {
	if m != nil {
		return m.BlockHash
	}
	return ""
}

type RemoveFundRequest struct {
	Address string `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
	Amount  int64  `protobuf:"varint,2,opt,name=amount" json:"amount,omitempty"`
}

func (m *RemoveFundRequest) Reset()                    { *m = RemoveFundRequest{} }
func (m *RemoveFundRequest) String() string            { return proto.CompactTextString(m) }
func (*RemoveFundRequest) ProtoMessage()               {}
func (*RemoveFundRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *RemoveFundRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *RemoveFundRequest) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type RemoveFundReply struct {
	PaymentRequest string `protobuf:"bytes,1,opt,name=paymentRequest" json:"paymentRequest,omitempty"`
	ErrorMessage   string `protobuf:"bytes,2,opt,name=errorMessage" json:"errorMessage,omitempty"`
}

func (m *RemoveFundReply) Reset()                    { *m = RemoveFundReply{} }
func (m *RemoveFundReply) String() string            { return proto.CompactTextString(m) }
func (*RemoveFundReply) ProtoMessage()               {}
func (*RemoveFundReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *RemoveFundReply) GetPaymentRequest() string {
	if m != nil {
		return m.PaymentRequest
	}
	return ""
}

func (m *RemoveFundReply) GetErrorMessage() string {
	if m != nil {
		return m.ErrorMessage
	}
	return ""
}

type RedeemRemovedFundsRequest struct {
	Paymenthash string `protobuf:"bytes,1,opt,name=paymenthash" json:"paymenthash,omitempty"`
}

func (m *RedeemRemovedFundsRequest) Reset()                    { *m = RedeemRemovedFundsRequest{} }
func (m *RedeemRemovedFundsRequest) String() string            { return proto.CompactTextString(m) }
func (*RedeemRemovedFundsRequest) ProtoMessage()               {}
func (*RedeemRemovedFundsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *RedeemRemovedFundsRequest) GetPaymenthash() string {
	if m != nil {
		return m.Paymenthash
	}
	return ""
}

type RedeemRemovedFundsReply struct {
	Txid string `protobuf:"bytes,1,opt,name=txid" json:"txid,omitempty"`
}

func (m *RedeemRemovedFundsReply) Reset()                    { *m = RedeemRemovedFundsReply{} }
func (m *RedeemRemovedFundsReply) String() string            { return proto.CompactTextString(m) }
func (*RedeemRemovedFundsReply) ProtoMessage()               {}
func (*RedeemRemovedFundsReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *RedeemRemovedFundsReply) GetTxid() string {
	if m != nil {
		return m.Txid
	}
	return ""
}

type GetSwapPaymentRequest struct {
	PaymentRequest string `protobuf:"bytes,1,opt,name=paymentRequest" json:"paymentRequest,omitempty"`
}

func (m *GetSwapPaymentRequest) Reset()                    { *m = GetSwapPaymentRequest{} }
func (m *GetSwapPaymentRequest) String() string            { return proto.CompactTextString(m) }
func (*GetSwapPaymentRequest) ProtoMessage()               {}
func (*GetSwapPaymentRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *GetSwapPaymentRequest) GetPaymentRequest() string {
	if m != nil {
		return m.PaymentRequest
	}
	return ""
}

type GetSwapPaymentReply struct {
	PaymentError string `protobuf:"bytes,1,opt,name=paymentError" json:"paymentError,omitempty"`
}

func (m *GetSwapPaymentReply) Reset()                    { *m = GetSwapPaymentReply{} }
func (m *GetSwapPaymentReply) String() string            { return proto.CompactTextString(m) }
func (*GetSwapPaymentReply) ProtoMessage()               {}
func (*GetSwapPaymentReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *GetSwapPaymentReply) GetPaymentError() string {
	if m != nil {
		return m.PaymentError
	}
	return ""
}

type MempoolRegisterRequest struct {
	ClientID  string   `protobuf:"bytes,1,opt,name=clientID" json:"clientID,omitempty"`
	Addresses []string `protobuf:"bytes,2,rep,name=addresses" json:"addresses,omitempty"`
}

func (m *MempoolRegisterRequest) Reset()                    { *m = MempoolRegisterRequest{} }
func (m *MempoolRegisterRequest) String() string            { return proto.CompactTextString(m) }
func (*MempoolRegisterRequest) ProtoMessage()               {}
func (*MempoolRegisterRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *MempoolRegisterRequest) GetClientID() string {
	if m != nil {
		return m.ClientID
	}
	return ""
}

func (m *MempoolRegisterRequest) GetAddresses() []string {
	if m != nil {
		return m.Addresses
	}
	return nil
}

type MempoolRegisterReply struct {
	TXS []*MempoolRegisterReply_Transaction `protobuf:"bytes,1,rep,name=TXS" json:"TXS,omitempty"`
}

func (m *MempoolRegisterReply) Reset()                    { *m = MempoolRegisterReply{} }
func (m *MempoolRegisterReply) String() string            { return proto.CompactTextString(m) }
func (*MempoolRegisterReply) ProtoMessage()               {}
func (*MempoolRegisterReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *MempoolRegisterReply) GetTXS() []*MempoolRegisterReply_Transaction {
	if m != nil {
		return m.TXS
	}
	return nil
}

type MempoolRegisterReply_Transaction struct {
	TX      string  `protobuf:"bytes,1,opt,name=TX" json:"TX,omitempty"`
	Address string  `protobuf:"bytes,2,opt,name=Address" json:"Address,omitempty"`
	Value   float64 `protobuf:"fixed64,3,opt,name=Value" json:"Value,omitempty"`
}

func (m *MempoolRegisterReply_Transaction) Reset()         { *m = MempoolRegisterReply_Transaction{} }
func (m *MempoolRegisterReply_Transaction) String() string { return proto.CompactTextString(m) }
func (*MempoolRegisterReply_Transaction) ProtoMessage()    {}
func (*MempoolRegisterReply_Transaction) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{15, 0}
}

func (m *MempoolRegisterReply_Transaction) GetTX() string {
	if m != nil {
		return m.TX
	}
	return ""
}

func (m *MempoolRegisterReply_Transaction) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *MempoolRegisterReply_Transaction) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type PingRequest struct {
}

func (m *PingRequest) Reset()                    { *m = PingRequest{} }
func (m *PingRequest) String() string            { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()               {}
func (*PingRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

type PingReply struct {
	Version string `protobuf:"bytes,1,opt,name=version" json:"version,omitempty"`
}

func (m *PingReply) Reset()                    { *m = PingReply{} }
func (m *PingReply) String() string            { return proto.CompactTextString(m) }
func (*PingReply) ProtoMessage()               {}
func (*PingReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{17} }

func (m *PingReply) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func init() {
	proto.RegisterType((*OpenChannelRequest)(nil), "breez.OpenChannelRequest")
	proto.RegisterType((*OpenChannelReply)(nil), "breez.OpenChannelReply")
	proto.RegisterType((*UpdateChannelPolicyRequest)(nil), "breez.UpdateChannelPolicyRequest")
	proto.RegisterType((*UpdateChannelPolicyReply)(nil), "breez.UpdateChannelPolicyReply")
	proto.RegisterType((*AddFundInitRequest)(nil), "breez.AddFundInitRequest")
	proto.RegisterType((*AddFundInitReply)(nil), "breez.AddFundInitReply")
	proto.RegisterType((*AddFundStatusRequest)(nil), "breez.AddFundStatusRequest")
	proto.RegisterType((*AddFundStatusReply)(nil), "breez.AddFundStatusReply")
	proto.RegisterType((*AddFundStatusReply_AddressStatus)(nil), "breez.AddFundStatusReply.AddressStatus")
	proto.RegisterType((*RemoveFundRequest)(nil), "breez.RemoveFundRequest")
	proto.RegisterType((*RemoveFundReply)(nil), "breez.RemoveFundReply")
	proto.RegisterType((*RedeemRemovedFundsRequest)(nil), "breez.RedeemRemovedFundsRequest")
	proto.RegisterType((*RedeemRemovedFundsReply)(nil), "breez.RedeemRemovedFundsReply")
	proto.RegisterType((*GetSwapPaymentRequest)(nil), "breez.GetSwapPaymentRequest")
	proto.RegisterType((*GetSwapPaymentReply)(nil), "breez.GetSwapPaymentReply")
	proto.RegisterType((*MempoolRegisterRequest)(nil), "breez.MempoolRegisterRequest")
	proto.RegisterType((*MempoolRegisterReply)(nil), "breez.MempoolRegisterReply")
	proto.RegisterType((*MempoolRegisterReply_Transaction)(nil), "breez.MempoolRegisterReply.Transaction")
	proto.RegisterType((*PingRequest)(nil), "breez.PingRequest")
	proto.RegisterType((*PingReply)(nil), "breez.PingReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Information service

type InformationClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error)
}

type informationClient struct {
	cc *grpc.ClientConn
}

func NewInformationClient(cc *grpc.ClientConn) InformationClient {
	return &informationClient{cc}
}

func (c *informationClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error) {
	out := new(PingReply)
	err := grpc.Invoke(ctx, "/breez.Information/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Information service

type InformationServer interface {
	Ping(context.Context, *PingRequest) (*PingReply, error)
}

func RegisterInformationServer(s *grpc.Server, srv InformationServer) {
	s.RegisterService(&_Information_serviceDesc, srv)
}

func _Information_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InformationServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/breez.Information/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InformationServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Information_serviceDesc = grpc.ServiceDesc{
	ServiceName: "breez.Information",
	HandlerType: (*InformationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Information_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "breez.proto",
}

// Client API for MempoolNotifier service

type MempoolNotifierClient interface {
	MempoolRegister(ctx context.Context, in *MempoolRegisterRequest, opts ...grpc.CallOption) (*MempoolRegisterReply, error)
}

type mempoolNotifierClient struct {
	cc *grpc.ClientConn
}

func NewMempoolNotifierClient(cc *grpc.ClientConn) MempoolNotifierClient {
	return &mempoolNotifierClient{cc}
}

func (c *mempoolNotifierClient) MempoolRegister(ctx context.Context, in *MempoolRegisterRequest, opts ...grpc.CallOption) (*MempoolRegisterReply, error) {
	out := new(MempoolRegisterReply)
	err := grpc.Invoke(ctx, "/breez.MempoolNotifier/MempoolRegister", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MempoolNotifier service

type MempoolNotifierServer interface {
	MempoolRegister(context.Context, *MempoolRegisterRequest) (*MempoolRegisterReply, error)
}

func RegisterMempoolNotifierServer(s *grpc.Server, srv MempoolNotifierServer) {
	s.RegisterService(&_MempoolNotifier_serviceDesc, srv)
}

func _MempoolNotifier_MempoolRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MempoolRegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MempoolNotifierServer).MempoolRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/breez.MempoolNotifier/MempoolRegister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MempoolNotifierServer).MempoolRegister(ctx, req.(*MempoolRegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MempoolNotifier_serviceDesc = grpc.ServiceDesc{
	ServiceName: "breez.MempoolNotifier",
	HandlerType: (*MempoolNotifierServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MempoolRegister",
			Handler:    _MempoolNotifier_MempoolRegister_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "breez.proto",
}

// Client API for FundManager service

type FundManagerClient interface {
	OpenChannel(ctx context.Context, in *OpenChannelRequest, opts ...grpc.CallOption) (*OpenChannelReply, error)
	UpdateChannelPolicy(ctx context.Context, in *UpdateChannelPolicyRequest, opts ...grpc.CallOption) (*UpdateChannelPolicyReply, error)
	AddFundInit(ctx context.Context, in *AddFundInitRequest, opts ...grpc.CallOption) (*AddFundInitReply, error)
	AddFundStatus(ctx context.Context, in *AddFundStatusRequest, opts ...grpc.CallOption) (*AddFundStatusReply, error)
	RemoveFund(ctx context.Context, in *RemoveFundRequest, opts ...grpc.CallOption) (*RemoveFundReply, error)
	RedeemRemovedFunds(ctx context.Context, in *RedeemRemovedFundsRequest, opts ...grpc.CallOption) (*RedeemRemovedFundsReply, error)
	GetSwapPayment(ctx context.Context, in *GetSwapPaymentRequest, opts ...grpc.CallOption) (*GetSwapPaymentReply, error)
}

type fundManagerClient struct {
	cc *grpc.ClientConn
}

func NewFundManagerClient(cc *grpc.ClientConn) FundManagerClient {
	return &fundManagerClient{cc}
}

func (c *fundManagerClient) OpenChannel(ctx context.Context, in *OpenChannelRequest, opts ...grpc.CallOption) (*OpenChannelReply, error) {
	out := new(OpenChannelReply)
	err := grpc.Invoke(ctx, "/breez.FundManager/OpenChannel", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fundManagerClient) UpdateChannelPolicy(ctx context.Context, in *UpdateChannelPolicyRequest, opts ...grpc.CallOption) (*UpdateChannelPolicyReply, error) {
	out := new(UpdateChannelPolicyReply)
	err := grpc.Invoke(ctx, "/breez.FundManager/UpdateChannelPolicy", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fundManagerClient) AddFundInit(ctx context.Context, in *AddFundInitRequest, opts ...grpc.CallOption) (*AddFundInitReply, error) {
	out := new(AddFundInitReply)
	err := grpc.Invoke(ctx, "/breez.FundManager/AddFundInit", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fundManagerClient) AddFundStatus(ctx context.Context, in *AddFundStatusRequest, opts ...grpc.CallOption) (*AddFundStatusReply, error) {
	out := new(AddFundStatusReply)
	err := grpc.Invoke(ctx, "/breez.FundManager/AddFundStatus", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fundManagerClient) RemoveFund(ctx context.Context, in *RemoveFundRequest, opts ...grpc.CallOption) (*RemoveFundReply, error) {
	out := new(RemoveFundReply)
	err := grpc.Invoke(ctx, "/breez.FundManager/RemoveFund", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fundManagerClient) RedeemRemovedFunds(ctx context.Context, in *RedeemRemovedFundsRequest, opts ...grpc.CallOption) (*RedeemRemovedFundsReply, error) {
	out := new(RedeemRemovedFundsReply)
	err := grpc.Invoke(ctx, "/breez.FundManager/RedeemRemovedFunds", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fundManagerClient) GetSwapPayment(ctx context.Context, in *GetSwapPaymentRequest, opts ...grpc.CallOption) (*GetSwapPaymentReply, error) {
	out := new(GetSwapPaymentReply)
	err := grpc.Invoke(ctx, "/breez.FundManager/GetSwapPayment", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for FundManager service

type FundManagerServer interface {
	OpenChannel(context.Context, *OpenChannelRequest) (*OpenChannelReply, error)
	UpdateChannelPolicy(context.Context, *UpdateChannelPolicyRequest) (*UpdateChannelPolicyReply, error)
	AddFundInit(context.Context, *AddFundInitRequest) (*AddFundInitReply, error)
	AddFundStatus(context.Context, *AddFundStatusRequest) (*AddFundStatusReply, error)
	RemoveFund(context.Context, *RemoveFundRequest) (*RemoveFundReply, error)
	RedeemRemovedFunds(context.Context, *RedeemRemovedFundsRequest) (*RedeemRemovedFundsReply, error)
	GetSwapPayment(context.Context, *GetSwapPaymentRequest) (*GetSwapPaymentReply, error)
}

func RegisterFundManagerServer(s *grpc.Server, srv FundManagerServer) {
	s.RegisterService(&_FundManager_serviceDesc, srv)
}

func _FundManager_OpenChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OpenChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FundManagerServer).OpenChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/breez.FundManager/OpenChannel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FundManagerServer).OpenChannel(ctx, req.(*OpenChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FundManager_UpdateChannelPolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateChannelPolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FundManagerServer).UpdateChannelPolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/breez.FundManager/UpdateChannelPolicy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FundManagerServer).UpdateChannelPolicy(ctx, req.(*UpdateChannelPolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FundManager_AddFundInit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddFundInitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FundManagerServer).AddFundInit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/breez.FundManager/AddFundInit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FundManagerServer).AddFundInit(ctx, req.(*AddFundInitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FundManager_AddFundStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddFundStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FundManagerServer).AddFundStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/breez.FundManager/AddFundStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FundManagerServer).AddFundStatus(ctx, req.(*AddFundStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FundManager_RemoveFund_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveFundRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FundManagerServer).RemoveFund(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/breez.FundManager/RemoveFund",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FundManagerServer).RemoveFund(ctx, req.(*RemoveFundRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FundManager_RedeemRemovedFunds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RedeemRemovedFundsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FundManagerServer).RedeemRemovedFunds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/breez.FundManager/RedeemRemovedFunds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FundManagerServer).RedeemRemovedFunds(ctx, req.(*RedeemRemovedFundsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FundManager_GetSwapPayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSwapPaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FundManagerServer).GetSwapPayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/breez.FundManager/GetSwapPayment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FundManagerServer).GetSwapPayment(ctx, req.(*GetSwapPaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FundManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "breez.FundManager",
	HandlerType: (*FundManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OpenChannel",
			Handler:    _FundManager_OpenChannel_Handler,
		},
		{
			MethodName: "UpdateChannelPolicy",
			Handler:    _FundManager_UpdateChannelPolicy_Handler,
		},
		{
			MethodName: "AddFundInit",
			Handler:    _FundManager_AddFundInit_Handler,
		},
		{
			MethodName: "AddFundStatus",
			Handler:    _FundManager_AddFundStatus_Handler,
		},
		{
			MethodName: "RemoveFund",
			Handler:    _FundManager_RemoveFund_Handler,
		},
		{
			MethodName: "RedeemRemovedFunds",
			Handler:    _FundManager_RedeemRemovedFunds_Handler,
		},
		{
			MethodName: "GetSwapPayment",
			Handler:    _FundManager_GetSwapPayment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "breez.proto",
}

func init() { proto.RegisterFile("breez.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 900 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x56, 0x51, 0x73, 0xdb, 0x44,
	0x10, 0x8e, 0x6c, 0xa7, 0xc4, 0xab, 0x24, 0x4d, 0xaf, 0x69, 0xea, 0xa8, 0xa5, 0x84, 0x9b, 0x29,
	0xe4, 0xa1, 0xf8, 0x21, 0xf0, 0x40, 0x99, 0xc9, 0x40, 0x92, 0x06, 0xc8, 0x40, 0xa8, 0xe7, 0x62,
	0x98, 0xcc, 0x30, 0x3c, 0xc8, 0xd6, 0xc6, 0xd1, 0x44, 0xbe, 0x13, 0xa7, 0x73, 0x6a, 0xf3, 0x03,
	0xf8, 0x1d, 0xbc, 0xf2, 0xce, 0xf0, 0xfb, 0x98, 0x3b, 0x9d, 0x64, 0xc9, 0x92, 0x0d, 0xbc, 0x69,
	0xf7, 0x76, 0xbf, 0xfd, 0x76, 0x6f, 0x77, 0x4f, 0xe0, 0x0e, 0x24, 0xe2, 0x6f, 0xdd, 0x58, 0x0a,
	0x25, 0xc8, 0xba, 0x11, 0xe8, 0x2b, 0x20, 0x6f, 0x63, 0xe4, 0x67, 0xb7, 0x3e, 0xe7, 0x18, 0x31,
	0xfc, 0x75, 0x82, 0x89, 0x22, 0x7b, 0xf0, 0x20, 0x9e, 0x0c, 0xbe, 0xc3, 0x59, 0xc7, 0x39, 0x70,
	0x0e, 0xdb, 0xcc, 0x4a, 0x94, 0xc0, 0x4e, 0xc9, 0x3a, 0x8e, 0x66, 0xf4, 0x33, 0xf0, 0x7e, 0x8c,
	0x03, 0x5f, 0xa1, 0xd5, 0xf6, 0x44, 0x14, 0x0e, 0x67, 0xff, 0x86, 0xe4, 0x41, 0xa7, 0xd6, 0x4b,
	0x23, 0xfe, 0xee, 0x00, 0x39, 0x09, 0x82, 0xaf, 0x27, 0x3c, 0xb8, 0xe0, 0xa1, 0x2a, 0x40, 0x71,
	0x11, 0xe0, 0xc5, 0x9b, 0x0c, 0x2a, 0x95, 0xc8, 0x2b, 0x78, 0xc4, 0x85, 0x0a, 0x6f, 0xc2, 0xa1,
	0xaf, 0x42, 0xc1, 0xfb, 0xe2, 0x0e, 0x79, 0xa7, 0x61, 0x4c, 0xaa, 0x07, 0x96, 0xd0, 0x1d, 0xce,
	0x3a, 0xcd, 0x03, 0xe7, 0x70, 0x93, 0x59, 0x89, 0x10, 0x68, 0xdd, 0xfa, 0xc9, 0x6d, 0xa7, 0x65,
	0xb4, 0xe6, 0x9b, 0xfe, 0xed, 0xc0, 0x4e, 0x89, 0x48, 0x1c, 0xcd, 0x48, 0x07, 0xde, 0xf3, 0x83,
	0x40, 0x62, 0x92, 0x58, 0x1e, 0x99, 0x58, 0x80, 0x6e, 0x94, 0xa0, 0x5f, 0x00, 0x44, 0x62, 0x78,
	0xf7, 0x2d, 0x86, 0xa3, 0x5b, 0x65, 0xc2, 0x36, 0x59, 0x41, 0xa3, 0x13, 0x18, 0xfb, 0xd3, 0x93,
	0x28, 0x12, 0xef, 0x30, 0x78, 0x83, 0xb1, 0x48, 0x42, 0x65, 0x78, 0x34, 0x59, 0xf5, 0x80, 0x50,
	0xd8, 0x44, 0x29, 0x85, 0xbc, 0xc4, 0x24, 0xf1, 0x47, 0xd8, 0x59, 0x37, 0x24, 0x4a, 0x3a, 0x3a,
	0x80, 0x5d, 0xcb, 0xfb, 0x4a, 0xf9, 0x6a, 0x92, 0x64, 0x25, 0x7c, 0x0e, 0x6d, 0x4b, 0x16, 0x35,
	0xfb, 0xe6, 0x61, 0x9b, 0xcd, 0x15, 0xff, 0xaf, 0x90, 0xf4, 0xaf, 0x46, 0x7e, 0x4b, 0x59, 0x10,
	0x5d, 0x9e, 0x33, 0xd8, 0x48, 0x8c, 0x68, 0x23, 0xb8, 0x47, 0x1f, 0x77, 0xd3, 0xbe, 0xab, 0x1a,
	0x77, 0xaf, 0xac, 0xe5, 0x39, 0x57, 0x72, 0xc6, 0x72, 0x47, 0x2f, 0x81, 0xad, 0x93, 0x94, 0x56,
	0x6a, 0x41, 0xb6, 0xa1, 0xa1, 0xa6, 0xb6, 0xde, 0x0d, 0x35, 0xd5, 0xa5, 0xf6, 0xc7, 0x62, 0xc2,
	0x95, 0xe1, 0xd7, 0x64, 0x56, 0xd2, 0x09, 0x0e, 0x05, 0xbf, 0x09, 0xe5, 0x18, 0x03, 0x53, 0xe9,
	0x0d, 0x36, 0x57, 0xe8, 0xd3, 0x81, 0xa9, 0x7b, 0x76, 0xd1, 0x6d, 0x36, 0x57, 0x78, 0x01, 0x6c,
	0x95, 0xf8, 0x90, 0x1d, 0x68, 0xde, 0xe5, 0x8d, 0xab, 0x3f, 0xc9, 0x31, 0xac, 0xdf, 0xfb, 0xd1,
	0x04, 0x4d, 0xd4, 0x95, 0x99, 0x95, 0xe8, 0xb3, 0xd4, 0xeb, 0x8b, 0xc6, 0xe7, 0x0e, 0x3d, 0x87,
	0x47, 0x0c, 0xc7, 0xe2, 0x1e, 0xb5, 0x47, 0x76, 0x2f, 0x2b, 0x7b, 0xaa, 0x2e, 0x51, 0xfa, 0x0b,
	0x3c, 0x2c, 0xc2, 0xe8, 0xca, 0x7f, 0x04, 0xdb, 0xb1, 0x3f, 0x1b, 0x23, 0xcf, 0x26, 0xc6, 0x62,
	0x2d, 0x68, 0x2b, 0x0d, 0xd4, 0xa8, 0x69, 0xa0, 0x63, 0xd8, 0x67, 0x18, 0x20, 0x8e, 0xd3, 0x20,
	0x26, 0xbd, 0xbc, 0x8b, 0x0e, 0xc0, 0xb5, 0x90, 0x66, 0x62, 0xd2, 0x28, 0x45, 0x15, 0xfd, 0x04,
	0x9e, 0xd6, 0xb9, 0x6b, 0x96, 0x04, 0x5a, 0x6a, 0x1a, 0x06, 0xd6, 0xcb, 0x7c, 0xd3, 0x2f, 0xe1,
	0xc9, 0x37, 0xa8, 0xae, 0xde, 0xf9, 0x71, 0xaf, 0x4c, 0xf5, 0x3f, 0xa6, 0x44, 0x5f, 0xc3, 0xe3,
	0x45, 0x00, 0x1d, 0x8b, 0xc2, 0xa6, 0x35, 0x3c, 0xd7, 0xc9, 0x59, 0xe7, 0x92, 0x8e, 0x32, 0xd8,
	0xbb, 0xc4, 0x71, 0x2c, 0x44, 0xc4, 0x70, 0x14, 0x26, 0x0a, 0x65, 0x16, 0xdc, 0x83, 0x8d, 0x61,
	0x14, 0x22, 0x57, 0xf9, 0xc6, 0xc9, 0xe5, 0xf2, 0x20, 0x35, 0x16, 0x06, 0x89, 0xfe, 0xe1, 0xc0,
	0x6e, 0x05, 0x54, 0x13, 0x7a, 0x0d, 0xcd, 0xfe, 0xf5, 0xd5, 0xc2, 0x5c, 0xd4, 0x59, 0x76, 0xfb,
	0xd2, 0xe7, 0x89, 0x3f, 0xd4, 0xe3, 0xc6, 0xb4, 0x8f, 0x77, 0x09, 0x6e, 0x41, 0xa7, 0x07, 0xa2,
	0x7f, 0x9d, 0x0d, 0x44, 0xff, 0x5a, 0x77, 0x90, 0x6d, 0x39, 0x7b, 0x9f, 0x99, 0x48, 0x76, 0x61,
	0xfd, 0x27, 0xd3, 0xb3, 0x7a, 0x1c, 0x1c, 0x96, 0x0a, 0x74, 0x0b, 0xdc, 0x5e, 0xc8, 0x47, 0x59,
	0x01, 0x5f, 0x42, 0x3b, 0x15, 0xed, 0x86, 0xbb, 0x47, 0x99, 0x84, 0x82, 0x67, 0xdd, 0x68, 0xc5,
	0xa3, 0x63, 0x70, 0x2f, 0xf8, 0x8d, 0x90, 0x63, 0xb3, 0x07, 0x48, 0x17, 0x5a, 0xda, 0x8b, 0x10,
	0x9b, 0x49, 0x01, 0xd1, 0xdb, 0x29, 0xe9, 0xf4, 0x5a, 0x5f, 0x3b, 0x1a, 0xc0, 0x43, 0x9b, 0xec,
	0x0f, 0x66, 0x9d, 0xa0, 0x24, 0x6f, 0x73, 0x55, 0x96, 0x3f, 0x79, 0x7f, 0x59, 0x5d, 0x52, 0xe0,
	0x67, 0x2b, 0xca, 0x46, 0xd7, 0x8e, 0xfe, 0x6c, 0x81, 0xab, 0xdb, 0xed, 0xd2, 0xe7, 0xfe, 0x08,
	0x25, 0x39, 0x03, 0xb7, 0xf0, 0x64, 0x91, 0x7d, 0xeb, 0x5d, 0x7d, 0xf4, 0xbc, 0xa7, 0x75, 0x47,
	0x06, 0x94, 0xfc, 0x0c, 0x8f, 0x6b, 0x5e, 0x2b, 0xf2, 0xa1, 0xf5, 0x58, 0xfe, 0xfe, 0x79, 0x1f,
	0xac, 0x32, 0x49, 0xc1, 0xcf, 0xc0, 0x2d, 0x3c, 0x32, 0x39, 0xc3, 0xea, 0x0b, 0x98, 0x33, 0x5c,
	0x7c, 0x93, 0xe8, 0x1a, 0xb9, 0x30, 0x1b, 0x73, 0xbe, 0x85, 0xc8, 0xb3, 0xfa, 0xdd, 0x94, 0x02,
	0xed, 0x2f, 0x5d, 0x5c, 0x74, 0x8d, 0x7c, 0x05, 0x30, 0x5f, 0x2d, 0xa4, 0x63, 0x4d, 0x2b, 0x4b,
	0xcb, 0xdb, 0xab, 0x39, 0x49, 0x11, 0xae, 0x81, 0x54, 0xc7, 0x9f, 0x1c, 0xe4, 0xf6, 0x4b, 0x16,
	0x8b, 0xf7, 0x62, 0x85, 0x45, 0x8a, 0xfc, 0x3d, 0x6c, 0x97, 0x07, 0x9d, 0x3c, 0xb7, 0x3e, 0xb5,
	0x0b, 0xc4, 0xf3, 0x96, 0x9c, 0x1a, 0xb4, 0xd3, 0x97, 0xf0, 0x24, 0x14, 0xdd, 0x91, 0x8c, 0x87,
	0xd6, 0x2c, 0x41, 0x79, 0x1f, 0x0e, 0xf1, 0x14, 0x4e, 0xb5, 0xd8, 0xd3, 0x3f, 0x4a, 0x3d, 0x67,
	0xf0, 0xc0, 0xfc, 0x31, 0x7d, 0xfa, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x29, 0xa7, 0xbc, 0x88,
	0x40, 0x09, 0x00, 0x00,
}

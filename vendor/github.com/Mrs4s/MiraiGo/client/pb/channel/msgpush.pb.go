// Code generated by protoc-gen-golite. DO NOT EDIT.
// source: pb/channel/msgpush.proto

package channel

type FocusInfo struct {
	ChannelIdList []uint64 `protobuf:"varint,1,rep"`
}

func (x *FocusInfo) GetChannelIdList() []uint64 {
	if x != nil {
		return x.ChannelIdList
	}
	return nil
}

type MsgOnlinePush struct {
	Msgs         []*ChannelMsgContent `protobuf:"bytes,1,rep"`
	GeneralFlag  *uint32              `protobuf:"varint,2,opt"`
	NeedResp     *uint32              `protobuf:"varint,3,opt"`
	ServerBuf    []byte               `protobuf:"bytes,4,opt"`
	CompressFlag *uint32              `protobuf:"varint,5,opt"`
	CompressMsg  []byte               `protobuf:"bytes,6,opt"`
	FocusInfo    *FocusInfo           `protobuf:"bytes,7,opt"`
	HugeFlag     *uint32              `protobuf:"varint,8,opt"`
}

func (x *MsgOnlinePush) GetMsgs() []*ChannelMsgContent {
	if x != nil {
		return x.Msgs
	}
	return nil
}

func (x *MsgOnlinePush) GetGeneralFlag() uint32 {
	if x != nil && x.GeneralFlag != nil {
		return *x.GeneralFlag
	}
	return 0
}

func (x *MsgOnlinePush) GetNeedResp() uint32 {
	if x != nil && x.NeedResp != nil {
		return *x.NeedResp
	}
	return 0
}

func (x *MsgOnlinePush) GetServerBuf() []byte {
	if x != nil {
		return x.ServerBuf
	}
	return nil
}

func (x *MsgOnlinePush) GetCompressFlag() uint32 {
	if x != nil && x.CompressFlag != nil {
		return *x.CompressFlag
	}
	return 0
}

func (x *MsgOnlinePush) GetCompressMsg() []byte {
	if x != nil {
		return x.CompressMsg
	}
	return nil
}

func (x *MsgOnlinePush) GetFocusInfo() *FocusInfo {
	if x != nil {
		return x.FocusInfo
	}
	return nil
}

func (x *MsgOnlinePush) GetHugeFlag() uint32 {
	if x != nil && x.HugeFlag != nil {
		return *x.HugeFlag
	}
	return 0
}

type MsgPushResp struct {
	ServerBuf []byte `protobuf:"bytes,1,opt"`
}

func (x *MsgPushResp) GetServerBuf() []byte {
	if x != nil {
		return x.ServerBuf
	}
	return nil
}

type PressMsg struct {
	Msgs []*ChannelMsgContent `protobuf:"bytes,1,rep"`
}

func (x *PressMsg) GetMsgs() []*ChannelMsgContent {
	if x != nil {
		return x.Msgs
	}
	return nil
}

type ServerBuf struct {
	SvrIp   *uint32 `protobuf:"varint,1,opt"`
	SvrPort *uint32 `protobuf:"varint,2,opt"`
	EchoKey []byte  `protobuf:"bytes,3,opt"`
}

func (x *ServerBuf) GetSvrIp() uint32 {
	if x != nil && x.SvrIp != nil {
		return *x.SvrIp
	}
	return 0
}

func (x *ServerBuf) GetSvrPort() uint32 {
	if x != nil && x.SvrPort != nil {
		return *x.SvrPort
	}
	return 0
}

func (x *ServerBuf) GetEchoKey() []byte {
	if x != nil {
		return x.EchoKey
	}
	return nil
}
// Code generated by protoc-gen-golite. DO NOT EDIT.
// source: group0x857.proto

package notify

type NotifyMsgBody struct {
	OptMsgGrayTips    *AIOGrayTipsInfo       `protobuf:"bytes,5,opt"`
	OptMsgRedTips     *RedGrayTipsInfo       `protobuf:"bytes,9,opt"`
	OptMsgRecall      *MessageRecallReminder `protobuf:"bytes,11,opt"`
	OptGeneralGrayTip *GeneralGrayTipInfo    `protobuf:"bytes,26,opt"`
	QqGroupDigestMsg  *QQGroupDigestMsg      `protobuf:"bytes,33,opt"`
	ServiceType       int32                  `protobuf:"varint,13,opt"`
}

func (x *NotifyMsgBody) GetOptMsgGrayTips() *AIOGrayTipsInfo {
	if x != nil {
		return x.OptMsgGrayTips
	}
	return nil
}

func (x *NotifyMsgBody) GetOptMsgRedTips() *RedGrayTipsInfo {
	if x != nil {
		return x.OptMsgRedTips
	}
	return nil
}

func (x *NotifyMsgBody) GetOptMsgRecall() *MessageRecallReminder {
	if x != nil {
		return x.OptMsgRecall
	}
	return nil
}

func (x *NotifyMsgBody) GetOptGeneralGrayTip() *GeneralGrayTipInfo {
	if x != nil {
		return x.OptGeneralGrayTip
	}
	return nil
}

func (x *NotifyMsgBody) GetQqGroupDigestMsg() *QQGroupDigestMsg {
	if x != nil {
		return x.QqGroupDigestMsg
	}
	return nil
}

func (x *NotifyMsgBody) GetServiceType() int32 {
	if x != nil {
		return x.ServiceType
	}
	return 0
}

type AIOGrayTipsInfo struct {
	ShowLatest     uint32 `protobuf:"varint,1,opt"`
	Content        []byte `protobuf:"bytes,2,opt"`
	Remind         uint32 `protobuf:"varint,3,opt"`
	Brief          []byte `protobuf:"bytes,4,opt"`
	ReceiverUin    uint64 `protobuf:"varint,5,opt"`
	ReliaoAdminOpt uint32 `protobuf:"varint,6,opt"`
}

func (x *AIOGrayTipsInfo) GetShowLatest() uint32 {
	if x != nil {
		return x.ShowLatest
	}
	return 0
}

func (x *AIOGrayTipsInfo) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *AIOGrayTipsInfo) GetRemind() uint32 {
	if x != nil {
		return x.Remind
	}
	return 0
}

func (x *AIOGrayTipsInfo) GetBrief() []byte {
	if x != nil {
		return x.Brief
	}
	return nil
}

func (x *AIOGrayTipsInfo) GetReceiverUin() uint64 {
	if x != nil {
		return x.ReceiverUin
	}
	return 0
}

func (x *AIOGrayTipsInfo) GetReliaoAdminOpt() uint32 {
	if x != nil {
		return x.ReliaoAdminOpt
	}
	return 0
}

type GeneralGrayTipInfo struct {
	BusiType      uint64        `protobuf:"varint,1,opt"`
	BusiId        uint64        `protobuf:"varint,2,opt"`
	CtrlFlag      uint32        `protobuf:"varint,3,opt"`
	C2CType       uint32        `protobuf:"varint,4,opt"`
	ServiceType   uint32        `protobuf:"varint,5,opt"`
	TemplId       uint64        `protobuf:"varint,6,opt"`
	MsgTemplParam []*TemplParam `protobuf:"bytes,7,rep"`
	Content       string        `protobuf:"bytes,8,opt"`
}

func (x *GeneralGrayTipInfo) GetBusiType() uint64 {
	if x != nil {
		return x.BusiType
	}
	return 0
}

func (x *GeneralGrayTipInfo) GetBusiId() uint64 {
	if x != nil {
		return x.BusiId
	}
	return 0
}

func (x *GeneralGrayTipInfo) GetCtrlFlag() uint32 {
	if x != nil {
		return x.CtrlFlag
	}
	return 0
}

func (x *GeneralGrayTipInfo) GetC2CType() uint32 {
	if x != nil {
		return x.C2CType
	}
	return 0
}

func (x *GeneralGrayTipInfo) GetServiceType() uint32 {
	if x != nil {
		return x.ServiceType
	}
	return 0
}

func (x *GeneralGrayTipInfo) GetTemplId() uint64 {
	if x != nil {
		return x.TemplId
	}
	return 0
}

func (x *GeneralGrayTipInfo) GetMsgTemplParam() []*TemplParam {
	if x != nil {
		return x.MsgTemplParam
	}
	return nil
}

func (x *GeneralGrayTipInfo) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type TemplParam struct {
	Name  string `protobuf:"bytes,1,opt"`
	Value string `protobuf:"bytes,2,opt"`
}

func (x *TemplParam) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TemplParam) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type MessageRecallReminder struct {
	Uin             int64                  `protobuf:"varint,1,opt"`
	Nickname        []byte                 `protobuf:"bytes,2,opt"`
	RecalledMsgList []*RecalledMessageMeta `protobuf:"bytes,3,rep"`
	ReminderContent []byte                 `protobuf:"bytes,4,opt"`
	Userdef         []byte                 `protobuf:"bytes,5,opt"`
	GroupType       int32                  `protobuf:"varint,6,opt"`
	OpType          int32                  `protobuf:"varint,7,opt"`
}

func (x *MessageRecallReminder) GetUin() int64 {
	if x != nil {
		return x.Uin
	}
	return 0
}

func (x *MessageRecallReminder) GetNickname() []byte {
	if x != nil {
		return x.Nickname
	}
	return nil
}

func (x *MessageRecallReminder) GetRecalledMsgList() []*RecalledMessageMeta {
	if x != nil {
		return x.RecalledMsgList
	}
	return nil
}

func (x *MessageRecallReminder) GetReminderContent() []byte {
	if x != nil {
		return x.ReminderContent
	}
	return nil
}

func (x *MessageRecallReminder) GetUserdef() []byte {
	if x != nil {
		return x.Userdef
	}
	return nil
}

func (x *MessageRecallReminder) GetGroupType() int32 {
	if x != nil {
		return x.GroupType
	}
	return 0
}

func (x *MessageRecallReminder) GetOpType() int32 {
	if x != nil {
		return x.OpType
	}
	return 0
}

type RecalledMessageMeta struct {
	Seq       int32 `protobuf:"varint,1,opt"`
	Time      int32 `protobuf:"varint,2,opt"`
	MsgRandom int32 `protobuf:"varint,3,opt"`
	MsgType   int32 `protobuf:"varint,4,opt"`
	MsgFlag   int32 `protobuf:"varint,5,opt"`
	AuthorUin int64 `protobuf:"varint,6,opt"`
}

func (x *RecalledMessageMeta) GetSeq() int32 {
	if x != nil {
		return x.Seq
	}
	return 0
}

func (x *RecalledMessageMeta) GetTime() int32 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *RecalledMessageMeta) GetMsgRandom() int32 {
	if x != nil {
		return x.MsgRandom
	}
	return 0
}

func (x *RecalledMessageMeta) GetMsgType() int32 {
	if x != nil {
		return x.MsgType
	}
	return 0
}

func (x *RecalledMessageMeta) GetMsgFlag() int32 {
	if x != nil {
		return x.MsgFlag
	}
	return 0
}

func (x *RecalledMessageMeta) GetAuthorUin() int64 {
	if x != nil {
		return x.AuthorUin
	}
	return 0
}

type RedGrayTipsInfo struct {
	ShowLatest          uint32 `protobuf:"varint,1,opt"`
	SenderUin           uint64 `protobuf:"varint,2,opt"`
	ReceiverUin         uint64 `protobuf:"varint,3,opt"`
	SenderRichContent   string `protobuf:"bytes,4,opt"`
	ReceiverRichContent string `protobuf:"bytes,5,opt"`
	AuthKey             []byte `protobuf:"bytes,6,opt"`
	MsgType             int32  `protobuf:"zigzag32,7,opt"`
	LuckyFlag           uint32 `protobuf:"varint,8,opt"`
	HideFlag            uint32 `protobuf:"varint,9,opt"`
	LuckyUin            uint64 `protobuf:"varint,12,opt"`
}

func (x *RedGrayTipsInfo) GetShowLatest() uint32 {
	if x != nil {
		return x.ShowLatest
	}
	return 0
}

func (x *RedGrayTipsInfo) GetSenderUin() uint64 {
	if x != nil {
		return x.SenderUin
	}
	return 0
}

func (x *RedGrayTipsInfo) GetReceiverUin() uint64 {
	if x != nil {
		return x.ReceiverUin
	}
	return 0
}

func (x *RedGrayTipsInfo) GetSenderRichContent() string {
	if x != nil {
		return x.SenderRichContent
	}
	return ""
}

func (x *RedGrayTipsInfo) GetReceiverRichContent() string {
	if x != nil {
		return x.ReceiverRichContent
	}
	return ""
}

func (x *RedGrayTipsInfo) GetAuthKey() []byte {
	if x != nil {
		return x.AuthKey
	}
	return nil
}

func (x *RedGrayTipsInfo) GetMsgType() int32 {
	if x != nil {
		return x.MsgType
	}
	return 0
}

func (x *RedGrayTipsInfo) GetLuckyFlag() uint32 {
	if x != nil {
		return x.LuckyFlag
	}
	return 0
}

func (x *RedGrayTipsInfo) GetHideFlag() uint32 {
	if x != nil {
		return x.HideFlag
	}
	return 0
}

func (x *RedGrayTipsInfo) GetLuckyUin() uint64 {
	if x != nil {
		return x.LuckyUin
	}
	return 0
}

type QQGroupDigestMsg struct {
	GroupCode     uint64 `protobuf:"varint,1,opt"`
	Seq           uint32 `protobuf:"varint,2,opt"`
	Random        uint32 `protobuf:"varint,3,opt"`
	OpType        int32  `protobuf:"varint,4,opt"`
	Sender        uint64 `protobuf:"varint,5,opt"`
	DigestOper    uint64 `protobuf:"varint,6,opt"`
	OpTime        uint32 `protobuf:"varint,7,opt"`
	LastestMsgSeq uint32 `protobuf:"varint,8,opt"`
	OperNick      []byte `protobuf:"bytes,9,opt"`
	SenderNick    []byte `protobuf:"bytes,10,opt"`
	ExtInfo       int32  `protobuf:"varint,11,opt"`
}

func (x *QQGroupDigestMsg) GetGroupCode() uint64 {
	if x != nil {
		return x.GroupCode
	}
	return 0
}

func (x *QQGroupDigestMsg) GetSeq() uint32 {
	if x != nil {
		return x.Seq
	}
	return 0
}

func (x *QQGroupDigestMsg) GetRandom() uint32 {
	if x != nil {
		return x.Random
	}
	return 0
}

func (x *QQGroupDigestMsg) GetOpType() int32 {
	if x != nil {
		return x.OpType
	}
	return 0
}

func (x *QQGroupDigestMsg) GetSender() uint64 {
	if x != nil {
		return x.Sender
	}
	return 0
}

func (x *QQGroupDigestMsg) GetDigestOper() uint64 {
	if x != nil {
		return x.DigestOper
	}
	return 0
}

func (x *QQGroupDigestMsg) GetOpTime() uint32 {
	if x != nil {
		return x.OpTime
	}
	return 0
}

func (x *QQGroupDigestMsg) GetLastestMsgSeq() uint32 {
	if x != nil {
		return x.LastestMsgSeq
	}
	return 0
}

func (x *QQGroupDigestMsg) GetOperNick() []byte {
	if x != nil {
		return x.OperNick
	}
	return nil
}

func (x *QQGroupDigestMsg) GetSenderNick() []byte {
	if x != nil {
		return x.SenderNick
	}
	return nil
}

func (x *QQGroupDigestMsg) GetExtInfo() int32 {
	if x != nil {
		return x.ExtInfo
	}
	return 0
}
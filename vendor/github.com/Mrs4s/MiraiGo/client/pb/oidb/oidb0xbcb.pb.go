// Code generated by protoc-gen-golite. DO NOT EDIT.
// source: oidb0xbcb.proto

package oidb

type CheckUrlReq struct {
	Url         []string `protobuf:"bytes,1,rep"`
	Refer       *string  `protobuf:"bytes,2,opt"`
	Plateform   *string  `protobuf:"bytes,3,opt"`
	QqPfTo      *string  `protobuf:"bytes,4,opt"`
	Type        *uint32  `protobuf:"varint,5,opt"`
	From        *uint32  `protobuf:"varint,6,opt"`
	Chatid      *uint64  `protobuf:"varint,7,opt"`
	ServiceType *uint64  `protobuf:"varint,8,opt"`
	SendUin     *uint64  `protobuf:"varint,9,opt"`
	ReqType     *string  `protobuf:"bytes,10,opt"`
	OriginalUrl *string  `protobuf:"bytes,11,opt"`
	IsArk       *bool    `protobuf:"varint,12,opt"`
	ArkName     *string  `protobuf:"bytes,13,opt"`
	IsFinish    *bool    `protobuf:"varint,14,opt"`
	SrcUrls     []string `protobuf:"bytes,15,rep"`
	SrcPlatform *uint32  `protobuf:"varint,16,opt"`
	Qua         *string  `protobuf:"bytes,17,opt"`
}

func (x *CheckUrlReq) GetUrl() []string {
	if x != nil {
		return x.Url
	}
	return nil
}

func (x *CheckUrlReq) GetRefer() string {
	if x != nil && x.Refer != nil {
		return *x.Refer
	}
	return ""
}

func (x *CheckUrlReq) GetPlateform() string {
	if x != nil && x.Plateform != nil {
		return *x.Plateform
	}
	return ""
}

func (x *CheckUrlReq) GetQqPfTo() string {
	if x != nil && x.QqPfTo != nil {
		return *x.QqPfTo
	}
	return ""
}

func (x *CheckUrlReq) GetType() uint32 {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return 0
}

func (x *CheckUrlReq) GetFrom() uint32 {
	if x != nil && x.From != nil {
		return *x.From
	}
	return 0
}

func (x *CheckUrlReq) GetChatid() uint64 {
	if x != nil && x.Chatid != nil {
		return *x.Chatid
	}
	return 0
}

func (x *CheckUrlReq) GetServiceType() uint64 {
	if x != nil && x.ServiceType != nil {
		return *x.ServiceType
	}
	return 0
}

func (x *CheckUrlReq) GetSendUin() uint64 {
	if x != nil && x.SendUin != nil {
		return *x.SendUin
	}
	return 0
}

func (x *CheckUrlReq) GetReqType() string {
	if x != nil && x.ReqType != nil {
		return *x.ReqType
	}
	return ""
}

func (x *CheckUrlReq) GetOriginalUrl() string {
	if x != nil && x.OriginalUrl != nil {
		return *x.OriginalUrl
	}
	return ""
}

func (x *CheckUrlReq) GetIsArk() bool {
	if x != nil && x.IsArk != nil {
		return *x.IsArk
	}
	return false
}

func (x *CheckUrlReq) GetArkName() string {
	if x != nil && x.ArkName != nil {
		return *x.ArkName
	}
	return ""
}

func (x *CheckUrlReq) GetIsFinish() bool {
	if x != nil && x.IsFinish != nil {
		return *x.IsFinish
	}
	return false
}

func (x *CheckUrlReq) GetSrcUrls() []string {
	if x != nil {
		return x.SrcUrls
	}
	return nil
}

func (x *CheckUrlReq) GetSrcPlatform() uint32 {
	if x != nil && x.SrcPlatform != nil {
		return *x.SrcPlatform
	}
	return 0
}

func (x *CheckUrlReq) GetQua() string {
	if x != nil && x.Qua != nil {
		return *x.Qua
	}
	return ""
}

type CheckUrlReqItem struct {
	Url         *string `protobuf:"bytes,1,opt"`
	Refer       *string `protobuf:"bytes,2,opt"`
	Plateform   *string `protobuf:"bytes,3,opt"`
	QqPfTo      *string `protobuf:"bytes,4,opt"`
	Type        *uint32 `protobuf:"varint,5,opt"`
	From        *uint32 `protobuf:"varint,6,opt"`
	Chatid      *uint64 `protobuf:"varint,7,opt"`
	ServiceType *uint64 `protobuf:"varint,8,opt"`
	SendUin     *uint64 `protobuf:"varint,9,opt"`
	ReqType     *string `protobuf:"bytes,10,opt"`
}

func (x *CheckUrlReqItem) GetUrl() string {
	if x != nil && x.Url != nil {
		return *x.Url
	}
	return ""
}

func (x *CheckUrlReqItem) GetRefer() string {
	if x != nil && x.Refer != nil {
		return *x.Refer
	}
	return ""
}

func (x *CheckUrlReqItem) GetPlateform() string {
	if x != nil && x.Plateform != nil {
		return *x.Plateform
	}
	return ""
}

func (x *CheckUrlReqItem) GetQqPfTo() string {
	if x != nil && x.QqPfTo != nil {
		return *x.QqPfTo
	}
	return ""
}

func (x *CheckUrlReqItem) GetType() uint32 {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return 0
}

func (x *CheckUrlReqItem) GetFrom() uint32 {
	if x != nil && x.From != nil {
		return *x.From
	}
	return 0
}

func (x *CheckUrlReqItem) GetChatid() uint64 {
	if x != nil && x.Chatid != nil {
		return *x.Chatid
	}
	return 0
}

func (x *CheckUrlReqItem) GetServiceType() uint64 {
	if x != nil && x.ServiceType != nil {
		return *x.ServiceType
	}
	return 0
}

func (x *CheckUrlReqItem) GetSendUin() uint64 {
	if x != nil && x.SendUin != nil {
		return *x.SendUin
	}
	return 0
}

func (x *CheckUrlReqItem) GetReqType() string {
	if x != nil && x.ReqType != nil {
		return *x.ReqType
	}
	return ""
}

type CheckUrlRsp struct {
	Results         []*UrlCheckResult `protobuf:"bytes,1,rep"`
	NextReqDuration *uint32           `protobuf:"varint,2,opt"`
}

func (x *CheckUrlRsp) GetResults() []*UrlCheckResult {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *CheckUrlRsp) GetNextReqDuration() uint32 {
	if x != nil && x.NextReqDuration != nil {
		return *x.NextReqDuration
	}
	return 0
}

type DBCBReqBody struct {
	NotUseCache *int32       `protobuf:"varint,9,opt"`
	CheckUrlReq *CheckUrlReq `protobuf:"bytes,10,opt"`
}

func (x *DBCBReqBody) GetNotUseCache() int32 {
	if x != nil && x.NotUseCache != nil {
		return *x.NotUseCache
	}
	return 0
}

func (x *DBCBReqBody) GetCheckUrlReq() *CheckUrlReq {
	if x != nil {
		return x.CheckUrlReq
	}
	return nil
}

type DBCBRspBody struct {
	Wording     *string      `protobuf:"bytes,1,opt"`
	CheckUrlRsp *CheckUrlRsp `protobuf:"bytes,10,opt"`
}

func (x *DBCBRspBody) GetWording() string {
	if x != nil && x.Wording != nil {
		return *x.Wording
	}
	return ""
}

func (x *DBCBRspBody) GetCheckUrlRsp() *CheckUrlRsp {
	if x != nil {
		return x.CheckUrlRsp
	}
	return nil
}

type UrlCheckResult struct {
	Url          *string `protobuf:"bytes,1,opt"`
	Result       *uint32 `protobuf:"varint,2,opt"`
	JumpResult   *uint32 `protobuf:"varint,3,opt"`
	JumpUrl      *string `protobuf:"bytes,4,opt"`
	Level        *uint32 `protobuf:"varint,5,opt"`
	SubLevel     *uint32 `protobuf:"varint,6,opt"`
	Umrtype      *uint32 `protobuf:"varint,7,opt"`
	RetFrom      *uint32 `protobuf:"varint,8,opt"`
	OperationBit *uint64 `protobuf:"varint,9,opt"`
}

func (x *UrlCheckResult) GetUrl() string {
	if x != nil && x.Url != nil {
		return *x.Url
	}
	return ""
}

func (x *UrlCheckResult) GetResult() uint32 {
	if x != nil && x.Result != nil {
		return *x.Result
	}
	return 0
}

func (x *UrlCheckResult) GetJumpResult() uint32 {
	if x != nil && x.JumpResult != nil {
		return *x.JumpResult
	}
	return 0
}

func (x *UrlCheckResult) GetJumpUrl() string {
	if x != nil && x.JumpUrl != nil {
		return *x.JumpUrl
	}
	return ""
}

func (x *UrlCheckResult) GetLevel() uint32 {
	if x != nil && x.Level != nil {
		return *x.Level
	}
	return 0
}

func (x *UrlCheckResult) GetSubLevel() uint32 {
	if x != nil && x.SubLevel != nil {
		return *x.SubLevel
	}
	return 0
}

func (x *UrlCheckResult) GetUmrtype() uint32 {
	if x != nil && x.Umrtype != nil {
		return *x.Umrtype
	}
	return 0
}

func (x *UrlCheckResult) GetRetFrom() uint32 {
	if x != nil && x.RetFrom != nil {
		return *x.RetFrom
	}
	return 0
}

func (x *UrlCheckResult) GetOperationBit() uint64 {
	if x != nil && x.OperationBit != nil {
		return *x.OperationBit
	}
	return 0
}
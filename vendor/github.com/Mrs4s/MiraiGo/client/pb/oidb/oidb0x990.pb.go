// Code generated by protoc-gen-golite. DO NOT EDIT.
// source: oidb0x990.proto

package oidb

type TranslateReqBody struct {
	// TranslateReq translate_req = 1;
	BatchTranslateReq *BatchTranslateReq `protobuf:"bytes,2,opt"`
}

func (x *TranslateReqBody) GetBatchTranslateReq() *BatchTranslateReq {
	if x != nil {
		return x.BatchTranslateReq
	}
	return nil
}

type TranslateRspBody struct {
	// TranslateRsp translate_rsp = 1;
	BatchTranslateRsp *BatchTranslateRsp `protobuf:"bytes,2,opt"`
}

func (x *TranslateRspBody) GetBatchTranslateRsp() *BatchTranslateRsp {
	if x != nil {
		return x.BatchTranslateRsp
	}
	return nil
}

type BatchTranslateReq struct {
	SrcLanguage string   `protobuf:"bytes,1,opt"`
	DstLanguage string   `protobuf:"bytes,2,opt"`
	SrcTextList []string `protobuf:"bytes,3,rep"`
}

func (x *BatchTranslateReq) GetSrcLanguage() string {
	if x != nil {
		return x.SrcLanguage
	}
	return ""
}

func (x *BatchTranslateReq) GetDstLanguage() string {
	if x != nil {
		return x.DstLanguage
	}
	return ""
}

func (x *BatchTranslateReq) GetSrcTextList() []string {
	if x != nil {
		return x.SrcTextList
	}
	return nil
}

type BatchTranslateRsp struct {
	ErrorCode   int32    `protobuf:"varint,1,opt"`
	ErrorMsg    []byte   `protobuf:"bytes,2,opt"`
	SrcLanguage string   `protobuf:"bytes,3,opt"`
	DstLanguage string   `protobuf:"bytes,4,opt"`
	SrcTextList []string `protobuf:"bytes,5,rep"`
	DstTextList []string `protobuf:"bytes,6,rep"`
}

func (x *BatchTranslateRsp) GetErrorCode() int32 {
	if x != nil {
		return x.ErrorCode
	}
	return 0
}

func (x *BatchTranslateRsp) GetErrorMsg() []byte {
	if x != nil {
		return x.ErrorMsg
	}
	return nil
}

func (x *BatchTranslateRsp) GetSrcLanguage() string {
	if x != nil {
		return x.SrcLanguage
	}
	return ""
}

func (x *BatchTranslateRsp) GetDstLanguage() string {
	if x != nil {
		return x.DstLanguage
	}
	return ""
}

func (x *BatchTranslateRsp) GetSrcTextList() []string {
	if x != nil {
		return x.SrcTextList
	}
	return nil
}

func (x *BatchTranslateRsp) GetDstTextList() []string {
	if x != nil {
		return x.DstTextList
	}
	return nil
}
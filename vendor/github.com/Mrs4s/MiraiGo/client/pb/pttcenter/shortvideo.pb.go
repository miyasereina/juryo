// Code generated by protoc-gen-golite. DO NOT EDIT.
// source: shortvideo.proto

package pttcenter

type ShortVideoReqBody struct {
	Cmd                      int32                     `protobuf:"varint,1,opt"`
	Seq                      int32                     `protobuf:"varint,2,opt"`
	PttShortVideoUploadReq   *ShortVideoUploadReq      `protobuf:"bytes,3,opt"`
	PttShortVideoDownloadReq *ShortVideoDownloadReq    `protobuf:"bytes,4,opt"`
	ExtensionReq             []*ShortVideoExtensionReq `protobuf:"bytes,100,rep"`
}

func (x *ShortVideoReqBody) GetCmd() int32 {
	if x != nil {
		return x.Cmd
	}
	return 0
}

func (x *ShortVideoReqBody) GetSeq() int32 {
	if x != nil {
		return x.Seq
	}
	return 0
}

func (x *ShortVideoReqBody) GetPttShortVideoUploadReq() *ShortVideoUploadReq {
	if x != nil {
		return x.PttShortVideoUploadReq
	}
	return nil
}

func (x *ShortVideoReqBody) GetPttShortVideoDownloadReq() *ShortVideoDownloadReq {
	if x != nil {
		return x.PttShortVideoDownloadReq
	}
	return nil
}

func (x *ShortVideoReqBody) GetExtensionReq() []*ShortVideoExtensionReq {
	if x != nil {
		return x.ExtensionReq
	}
	return nil
}

type ShortVideoRspBody struct {
	Cmd                      int32                  `protobuf:"varint,1,opt"`
	Seq                      int32                  `protobuf:"varint,2,opt"`
	PttShortVideoUploadRsp   *ShortVideoUploadRsp   `protobuf:"bytes,3,opt"`
	PttShortVideoDownloadRsp *ShortVideoDownloadRsp `protobuf:"bytes,4,opt"`
}

func (x *ShortVideoRspBody) GetCmd() int32 {
	if x != nil {
		return x.Cmd
	}
	return 0
}

func (x *ShortVideoRspBody) GetSeq() int32 {
	if x != nil {
		return x.Seq
	}
	return 0
}

func (x *ShortVideoRspBody) GetPttShortVideoUploadRsp() *ShortVideoUploadRsp {
	if x != nil {
		return x.PttShortVideoUploadRsp
	}
	return nil
}

func (x *ShortVideoRspBody) GetPttShortVideoDownloadRsp() *ShortVideoDownloadRsp {
	if x != nil {
		return x.PttShortVideoDownloadRsp
	}
	return nil
}

type ShortVideoUploadReq struct {
	FromUin          int64               `protobuf:"varint,1,opt"`
	ToUin            int64               `protobuf:"varint,2,opt"`
	ChatType         int32               `protobuf:"varint,3,opt"`
	ClientType       int32               `protobuf:"varint,4,opt"`
	Info             *ShortVideoFileInfo `protobuf:"bytes,5,opt"`
	GroupCode        int64               `protobuf:"varint,6,opt"`
	AgentType        int32               `protobuf:"varint,7,opt"`
	BusinessType     int32               `protobuf:"varint,8,opt"`
	SupportLargeSize int32               `protobuf:"varint,20,opt"`
}

func (x *ShortVideoUploadReq) GetFromUin() int64 {
	if x != nil {
		return x.FromUin
	}
	return 0
}

func (x *ShortVideoUploadReq) GetToUin() int64 {
	if x != nil {
		return x.ToUin
	}
	return 0
}

func (x *ShortVideoUploadReq) GetChatType() int32 {
	if x != nil {
		return x.ChatType
	}
	return 0
}

func (x *ShortVideoUploadReq) GetClientType() int32 {
	if x != nil {
		return x.ClientType
	}
	return 0
}

func (x *ShortVideoUploadReq) GetInfo() *ShortVideoFileInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *ShortVideoUploadReq) GetGroupCode() int64 {
	if x != nil {
		return x.GroupCode
	}
	return 0
}

func (x *ShortVideoUploadReq) GetAgentType() int32 {
	if x != nil {
		return x.AgentType
	}
	return 0
}

func (x *ShortVideoUploadReq) GetBusinessType() int32 {
	if x != nil {
		return x.BusinessType
	}
	return 0
}

func (x *ShortVideoUploadReq) GetSupportLargeSize() int32 {
	if x != nil {
		return x.SupportLargeSize
	}
	return 0
}

type ShortVideoDownloadReq struct {
	FromUin      int64  `protobuf:"varint,1,opt"`
	ToUin        int64  `protobuf:"varint,2,opt"`
	ChatType     int32  `protobuf:"varint,3,opt"`
	ClientType   int32  `protobuf:"varint,4,opt"`
	FileId       string `protobuf:"bytes,5,opt"`
	GroupCode    int64  `protobuf:"varint,6,opt"`
	AgentType    int32  `protobuf:"varint,7,opt"`
	FileMd5      []byte `protobuf:"bytes,8,opt"`
	BusinessType int32  `protobuf:"varint,9,opt"`
	FileType     int32  `protobuf:"varint,10,opt"`
	DownType     int32  `protobuf:"varint,11,opt"`
	SceneType    int32  `protobuf:"varint,12,opt"`
}

func (x *ShortVideoDownloadReq) GetFromUin() int64 {
	if x != nil {
		return x.FromUin
	}
	return 0
}

func (x *ShortVideoDownloadReq) GetToUin() int64 {
	if x != nil {
		return x.ToUin
	}
	return 0
}

func (x *ShortVideoDownloadReq) GetChatType() int32 {
	if x != nil {
		return x.ChatType
	}
	return 0
}

func (x *ShortVideoDownloadReq) GetClientType() int32 {
	if x != nil {
		return x.ClientType
	}
	return 0
}

func (x *ShortVideoDownloadReq) GetFileId() string {
	if x != nil {
		return x.FileId
	}
	return ""
}

func (x *ShortVideoDownloadReq) GetGroupCode() int64 {
	if x != nil {
		return x.GroupCode
	}
	return 0
}

func (x *ShortVideoDownloadReq) GetAgentType() int32 {
	if x != nil {
		return x.AgentType
	}
	return 0
}

func (x *ShortVideoDownloadReq) GetFileMd5() []byte {
	if x != nil {
		return x.FileMd5
	}
	return nil
}

func (x *ShortVideoDownloadReq) GetBusinessType() int32 {
	if x != nil {
		return x.BusinessType
	}
	return 0
}

func (x *ShortVideoDownloadReq) GetFileType() int32 {
	if x != nil {
		return x.FileType
	}
	return 0
}

func (x *ShortVideoDownloadReq) GetDownType() int32 {
	if x != nil {
		return x.DownType
	}
	return 0
}

func (x *ShortVideoDownloadReq) GetSceneType() int32 {
	if x != nil {
		return x.SceneType
	}
	return 0
}

type ShortVideoDownloadRsp struct {
	RetCode           int32               `protobuf:"varint,1,opt"`
	RetMsg            string              `protobuf:"bytes,2,opt"`
	SameAreaOutAddr   []*ShortVideoIpList `protobuf:"bytes,3,rep"`
	DiffAreaOutAddr   []*ShortVideoIpList `protobuf:"bytes,4,rep"`
	DownloadKey       []byte              `protobuf:"bytes,5,opt"`
	FileMd5           []byte              `protobuf:"bytes,6,opt"`
	SameAreaInnerAddr []*ShortVideoIpList `protobuf:"bytes,7,rep"`
	DiffAreaInnerAddr []*ShortVideoIpList `protobuf:"bytes,8,rep"`
	DownloadAddr      *ShortVideoAddr     `protobuf:"bytes,9,opt"`
	EncryptKey        []byte              `protobuf:"bytes,10,opt"`
}

func (x *ShortVideoDownloadRsp) GetRetCode() int32 {
	if x != nil {
		return x.RetCode
	}
	return 0
}

func (x *ShortVideoDownloadRsp) GetRetMsg() string {
	if x != nil {
		return x.RetMsg
	}
	return ""
}

func (x *ShortVideoDownloadRsp) GetSameAreaOutAddr() []*ShortVideoIpList {
	if x != nil {
		return x.SameAreaOutAddr
	}
	return nil
}

func (x *ShortVideoDownloadRsp) GetDiffAreaOutAddr() []*ShortVideoIpList {
	if x != nil {
		return x.DiffAreaOutAddr
	}
	return nil
}

func (x *ShortVideoDownloadRsp) GetDownloadKey() []byte {
	if x != nil {
		return x.DownloadKey
	}
	return nil
}

func (x *ShortVideoDownloadRsp) GetFileMd5() []byte {
	if x != nil {
		return x.FileMd5
	}
	return nil
}

func (x *ShortVideoDownloadRsp) GetSameAreaInnerAddr() []*ShortVideoIpList {
	if x != nil {
		return x.SameAreaInnerAddr
	}
	return nil
}

func (x *ShortVideoDownloadRsp) GetDiffAreaInnerAddr() []*ShortVideoIpList {
	if x != nil {
		return x.DiffAreaInnerAddr
	}
	return nil
}

func (x *ShortVideoDownloadRsp) GetDownloadAddr() *ShortVideoAddr {
	if x != nil {
		return x.DownloadAddr
	}
	return nil
}

func (x *ShortVideoDownloadRsp) GetEncryptKey() []byte {
	if x != nil {
		return x.EncryptKey
	}
	return nil
}

type ShortVideoUploadRsp struct {
	RetCode           int32               `protobuf:"varint,1,opt"`
	RetMsg            string              `protobuf:"bytes,2,opt"`
	SameAreaOutAddr   []*ShortVideoIpList `protobuf:"bytes,3,rep"`
	DiffAreaOutAddr   []*ShortVideoIpList `protobuf:"bytes,4,rep"`
	FileId            string              `protobuf:"bytes,5,opt"`
	UKey              []byte              `protobuf:"bytes,6,opt"`
	FileExists        int32               `protobuf:"varint,7,opt"`
	SameAreaInnerAddr []*ShortVideoIpList `protobuf:"bytes,8,rep"`
	DiffAreaInnerAddr []*ShortVideoIpList `protobuf:"bytes,9,rep"`
	DataHole          []*DataHole         `protobuf:"bytes,10,rep"`
}

func (x *ShortVideoUploadRsp) GetRetCode() int32 {
	if x != nil {
		return x.RetCode
	}
	return 0
}

func (x *ShortVideoUploadRsp) GetRetMsg() string {
	if x != nil {
		return x.RetMsg
	}
	return ""
}

func (x *ShortVideoUploadRsp) GetSameAreaOutAddr() []*ShortVideoIpList {
	if x != nil {
		return x.SameAreaOutAddr
	}
	return nil
}

func (x *ShortVideoUploadRsp) GetDiffAreaOutAddr() []*ShortVideoIpList {
	if x != nil {
		return x.DiffAreaOutAddr
	}
	return nil
}

func (x *ShortVideoUploadRsp) GetFileId() string {
	if x != nil {
		return x.FileId
	}
	return ""
}

func (x *ShortVideoUploadRsp) GetUKey() []byte {
	if x != nil {
		return x.UKey
	}
	return nil
}

func (x *ShortVideoUploadRsp) GetFileExists() int32 {
	if x != nil {
		return x.FileExists
	}
	return 0
}

func (x *ShortVideoUploadRsp) GetSameAreaInnerAddr() []*ShortVideoIpList {
	if x != nil {
		return x.SameAreaInnerAddr
	}
	return nil
}

func (x *ShortVideoUploadRsp) GetDiffAreaInnerAddr() []*ShortVideoIpList {
	if x != nil {
		return x.DiffAreaInnerAddr
	}
	return nil
}

func (x *ShortVideoUploadRsp) GetDataHole() []*DataHole {
	if x != nil {
		return x.DataHole
	}
	return nil
}

type ShortVideoFileInfo struct {
	FileName      string `protobuf:"bytes,1,opt"`
	FileMd5       []byte `protobuf:"bytes,2,opt"`
	ThumbFileMd5  []byte `protobuf:"bytes,3,opt"`
	FileSize      int64  `protobuf:"varint,4,opt"`
	FileResLength int32  `protobuf:"varint,5,opt"`
	FileResWidth  int32  `protobuf:"varint,6,opt"`
	FileFormat    int32  `protobuf:"varint,7,opt"`
	FileTime      int32  `protobuf:"varint,8,opt"`
	ThumbFileSize int64  `protobuf:"varint,9,opt"`
}

func (x *ShortVideoFileInfo) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *ShortVideoFileInfo) GetFileMd5() []byte {
	if x != nil {
		return x.FileMd5
	}
	return nil
}

func (x *ShortVideoFileInfo) GetThumbFileMd5() []byte {
	if x != nil {
		return x.ThumbFileMd5
	}
	return nil
}

func (x *ShortVideoFileInfo) GetFileSize() int64 {
	if x != nil {
		return x.FileSize
	}
	return 0
}

func (x *ShortVideoFileInfo) GetFileResLength() int32 {
	if x != nil {
		return x.FileResLength
	}
	return 0
}

func (x *ShortVideoFileInfo) GetFileResWidth() int32 {
	if x != nil {
		return x.FileResWidth
	}
	return 0
}

func (x *ShortVideoFileInfo) GetFileFormat() int32 {
	if x != nil {
		return x.FileFormat
	}
	return 0
}

func (x *ShortVideoFileInfo) GetFileTime() int32 {
	if x != nil {
		return x.FileTime
	}
	return 0
}

func (x *ShortVideoFileInfo) GetThumbFileSize() int64 {
	if x != nil {
		return x.ThumbFileSize
	}
	return 0
}

type DataHole struct {
	Begin int64 `protobuf:"varint,1,opt"`
	End   int64 `protobuf:"varint,2,opt"`
}

func (x *DataHole) GetBegin() int64 {
	if x != nil {
		return x.Begin
	}
	return 0
}

func (x *DataHole) GetEnd() int64 {
	if x != nil {
		return x.End
	}
	return 0
}

type ShortVideoIpList struct {
	Ip   int32 `protobuf:"varint,1,opt"`
	Port int32 `protobuf:"varint,2,opt"`
}

func (x *ShortVideoIpList) GetIp() int32 {
	if x != nil {
		return x.Ip
	}
	return 0
}

func (x *ShortVideoIpList) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

type ShortVideoAddr struct {
	Host    []string `protobuf:"bytes,10,rep"`
	UrlArgs string   `protobuf:"bytes,11,opt"` //repeated string domain = 13;
}

func (x *ShortVideoAddr) GetHost() []string {
	if x != nil {
		return x.Host
	}
	return nil
}

func (x *ShortVideoAddr) GetUrlArgs() string {
	if x != nil {
		return x.UrlArgs
	}
	return ""
}

type ShortVideoExtensionReq struct {
	SubBusiType int32 `protobuf:"varint,1,opt"`
	UserCnt     int32 `protobuf:"varint,2,opt"`
}

func (x *ShortVideoExtensionReq) GetSubBusiType() int32 {
	if x != nil {
		return x.SubBusiType
	}
	return 0
}

func (x *ShortVideoExtensionReq) GetUserCnt() int32 {
	if x != nil {
		return x.UserCnt
	}
	return 0
}
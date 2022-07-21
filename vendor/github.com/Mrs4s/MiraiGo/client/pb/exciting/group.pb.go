// Code generated by protoc-gen-golite. DO NOT EDIT.
// source: group.proto

package exciting

type GroupFileUploadExt struct {
	Unknown1 *int32                `protobuf:"varint,1,opt"`
	Unknown2 *int32                `protobuf:"varint,2,opt"`
	Entry    *GroupFileUploadEntry `protobuf:"bytes,100,opt"`
	Unknown3 *int32                `protobuf:"varint,3,opt"`
}

func (x *GroupFileUploadExt) GetUnknown1() int32 {
	if x != nil && x.Unknown1 != nil {
		return *x.Unknown1
	}
	return 0
}

func (x *GroupFileUploadExt) GetUnknown2() int32 {
	if x != nil && x.Unknown2 != nil {
		return *x.Unknown2
	}
	return 0
}

func (x *GroupFileUploadExt) GetEntry() *GroupFileUploadEntry {
	if x != nil {
		return x.Entry
	}
	return nil
}

func (x *GroupFileUploadExt) GetUnknown3() int32 {
	if x != nil && x.Unknown3 != nil {
		return *x.Unknown3
	}
	return 0
}

type GroupFileUploadEntry struct {
	BusiBuff     *ExcitingBusiInfo     `protobuf:"bytes,100,opt"`
	FileEntry    *ExcitingFileEntry    `protobuf:"bytes,200,opt"`
	ClientInfo   *ExcitingClientInfo   `protobuf:"bytes,300,opt"`
	FileNameInfo *ExcitingFileNameInfo `protobuf:"bytes,400,opt"`
	Host         *ExcitingHostConfig   `protobuf:"bytes,500,opt"`
}

func (x *GroupFileUploadEntry) GetBusiBuff() *ExcitingBusiInfo {
	if x != nil {
		return x.BusiBuff
	}
	return nil
}

func (x *GroupFileUploadEntry) GetFileEntry() *ExcitingFileEntry {
	if x != nil {
		return x.FileEntry
	}
	return nil
}

func (x *GroupFileUploadEntry) GetClientInfo() *ExcitingClientInfo {
	if x != nil {
		return x.ClientInfo
	}
	return nil
}

func (x *GroupFileUploadEntry) GetFileNameInfo() *ExcitingFileNameInfo {
	if x != nil {
		return x.FileNameInfo
	}
	return nil
}

func (x *GroupFileUploadEntry) GetHost() *ExcitingHostConfig {
	if x != nil {
		return x.Host
	}
	return nil
}

type ExcitingBusiInfo struct {
	BusId       *int32 `protobuf:"varint,1,opt"`
	SenderUin   *int64 `protobuf:"varint,100,opt"`
	ReceiverUin *int64 `protobuf:"varint,200,opt"` // probable
	GroupCode   *int64 `protobuf:"varint,400,opt"` // probable
}

func (x *ExcitingBusiInfo) GetBusId() int32 {
	if x != nil && x.BusId != nil {
		return *x.BusId
	}
	return 0
}

func (x *ExcitingBusiInfo) GetSenderUin() int64 {
	if x != nil && x.SenderUin != nil {
		return *x.SenderUin
	}
	return 0
}

func (x *ExcitingBusiInfo) GetReceiverUin() int64 {
	if x != nil && x.ReceiverUin != nil {
		return *x.ReceiverUin
	}
	return 0
}

func (x *ExcitingBusiInfo) GetGroupCode() int64 {
	if x != nil && x.GroupCode != nil {
		return *x.GroupCode
	}
	return 0
}

type ExcitingFileEntry struct {
	FileSize  *int64 `protobuf:"varint,100,opt"`
	Md5       []byte `protobuf:"bytes,200,opt"`
	Sha1      []byte `protobuf:"bytes,300,opt"`
	FileId    []byte `protobuf:"bytes,600,opt"`
	UploadKey []byte `protobuf:"bytes,700,opt"`
}

func (x *ExcitingFileEntry) GetFileSize() int64 {
	if x != nil && x.FileSize != nil {
		return *x.FileSize
	}
	return 0
}

func (x *ExcitingFileEntry) GetMd5() []byte {
	if x != nil {
		return x.Md5
	}
	return nil
}

func (x *ExcitingFileEntry) GetSha1() []byte {
	if x != nil {
		return x.Sha1
	}
	return nil
}

func (x *ExcitingFileEntry) GetFileId() []byte {
	if x != nil {
		return x.FileId
	}
	return nil
}

func (x *ExcitingFileEntry) GetUploadKey() []byte {
	if x != nil {
		return x.UploadKey
	}
	return nil
}

type ExcitingClientInfo struct {
	ClientType   *int32  `protobuf:"varint,100,opt"` // probable
	AppId        *string `protobuf:"bytes,200,opt"`
	TerminalType *int32  `protobuf:"varint,300,opt"` // probable
	ClientVer    *string `protobuf:"bytes,400,opt"`
	Unknown      *int32  `protobuf:"varint,600,opt"`
}

func (x *ExcitingClientInfo) GetClientType() int32 {
	if x != nil && x.ClientType != nil {
		return *x.ClientType
	}
	return 0
}

func (x *ExcitingClientInfo) GetAppId() string {
	if x != nil && x.AppId != nil {
		return *x.AppId
	}
	return ""
}

func (x *ExcitingClientInfo) GetTerminalType() int32 {
	if x != nil && x.TerminalType != nil {
		return *x.TerminalType
	}
	return 0
}

func (x *ExcitingClientInfo) GetClientVer() string {
	if x != nil && x.ClientVer != nil {
		return *x.ClientVer
	}
	return ""
}

func (x *ExcitingClientInfo) GetUnknown() int32 {
	if x != nil && x.Unknown != nil {
		return *x.Unknown
	}
	return 0
}

type ExcitingFileNameInfo struct {
	FileName *string `protobuf:"bytes,100,opt"`
}

func (x *ExcitingFileNameInfo) GetFileName() string {
	if x != nil && x.FileName != nil {
		return *x.FileName
	}
	return ""
}

type ExcitingHostConfig struct {
	Hosts []*ExcitingHostInfo `protobuf:"bytes,200,rep"`
}

func (x *ExcitingHostConfig) GetHosts() []*ExcitingHostInfo {
	if x != nil {
		return x.Hosts
	}
	return nil
}

type ExcitingHostInfo struct {
	Url  *ExcitingUrlInfo `protobuf:"bytes,1,opt"`
	Port *int32           `protobuf:"varint,2,opt"`
}

func (x *ExcitingHostInfo) GetUrl() *ExcitingUrlInfo {
	if x != nil {
		return x.Url
	}
	return nil
}

func (x *ExcitingHostInfo) GetPort() int32 {
	if x != nil && x.Port != nil {
		return *x.Port
	}
	return 0
}

type ExcitingUrlInfo struct {
	Unknown *int32  `protobuf:"varint,1,opt"` // not https?
	Host    *string `protobuf:"bytes,2,opt"`
}

func (x *ExcitingUrlInfo) GetUnknown() int32 {
	if x != nil && x.Unknown != nil {
		return *x.Unknown
	}
	return 0
}

func (x *ExcitingUrlInfo) GetHost() string {
	if x != nil && x.Host != nil {
		return *x.Host
	}
	return ""
}

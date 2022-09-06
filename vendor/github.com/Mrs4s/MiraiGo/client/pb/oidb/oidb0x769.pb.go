// Code generated by protoc-gen-golite. DO NOT EDIT.
// source: oidb0x769.proto

package oidb

type CPU struct {
	Model     *string `protobuf:"bytes,1,opt"`
	Cores     *uint32 `protobuf:"varint,2,opt"`
	Frequency *uint32 `protobuf:"varint,3,opt"`
}

func (x *CPU) GetModel() string {
	if x != nil && x.Model != nil {
		return *x.Model
	}
	return ""
}

func (x *CPU) GetCores() uint32 {
	if x != nil && x.Cores != nil {
		return *x.Cores
	}
	return 0
}

func (x *CPU) GetFrequency() uint32 {
	if x != nil && x.Frequency != nil {
		return *x.Frequency
	}
	return 0
}

type Camera struct {
	Primary   *uint64 `protobuf:"varint,1,opt"`
	Secondary *uint64 `protobuf:"varint,2,opt"`
	Flash     *bool   `protobuf:"varint,3,opt"`
}

func (x *Camera) GetPrimary() uint64 {
	if x != nil && x.Primary != nil {
		return *x.Primary
	}
	return 0
}

func (x *Camera) GetSecondary() uint64 {
	if x != nil && x.Secondary != nil {
		return *x.Secondary
	}
	return 0
}

func (x *Camera) GetFlash() bool {
	if x != nil && x.Flash != nil {
		return *x.Flash
	}
	return false
}

type D769ConfigSeq struct {
	Type    *uint32 `protobuf:"varint,1,opt"`
	Version *uint32 `protobuf:"varint,2,opt"`
}

func (x *D769ConfigSeq) GetType() uint32 {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return 0
}

func (x *D769ConfigSeq) GetVersion() uint32 {
	if x != nil && x.Version != nil {
		return *x.Version
	}
	return 0
}

type Content struct {
	TaskId   *uint32 `protobuf:"varint,1,opt"`
	Compress *uint32 `protobuf:"varint,2,opt"`
	Content  []byte  `protobuf:"bytes,10,opt"`
}

func (x *Content) GetTaskId() uint32 {
	if x != nil && x.TaskId != nil {
		return *x.TaskId
	}
	return 0
}

func (x *Content) GetCompress() uint32 {
	if x != nil && x.Compress != nil {
		return *x.Compress
	}
	return 0
}

func (x *Content) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

type D769DeviceInfo struct {
	Brand   *string   `protobuf:"bytes,1,opt"`
	Model   *string   `protobuf:"bytes,2,opt"`
	Os      *C41219OS `protobuf:"bytes,3,opt"`
	Cpu     *CPU      `protobuf:"bytes,4,opt"`
	Memory  *Memory   `protobuf:"bytes,5,opt"`
	Storage *Storage  `protobuf:"bytes,6,opt"`
	Screen  *Screen   `protobuf:"bytes,7,opt"`
	Camera  *Camera   `protobuf:"bytes,8,opt"`
}

func (x *D769DeviceInfo) GetBrand() string {
	if x != nil && x.Brand != nil {
		return *x.Brand
	}
	return ""
}

func (x *D769DeviceInfo) GetModel() string {
	if x != nil && x.Model != nil {
		return *x.Model
	}
	return ""
}

func (x *D769DeviceInfo) GetOs() *C41219OS {
	if x != nil {
		return x.Os
	}
	return nil
}

func (x *D769DeviceInfo) GetCpu() *CPU {
	if x != nil {
		return x.Cpu
	}
	return nil
}

func (x *D769DeviceInfo) GetMemory() *Memory {
	if x != nil {
		return x.Memory
	}
	return nil
}

func (x *D769DeviceInfo) GetStorage() *Storage {
	if x != nil {
		return x.Storage
	}
	return nil
}

func (x *D769DeviceInfo) GetScreen() *Screen {
	if x != nil {
		return x.Screen
	}
	return nil
}

func (x *D769DeviceInfo) GetCamera() *Camera {
	if x != nil {
		return x.Camera
	}
	return nil
}

type Memory struct {
	Total   *uint64 `protobuf:"varint,1,opt"`
	Process *uint64 `protobuf:"varint,2,opt"`
}

func (x *Memory) GetTotal() uint64 {
	if x != nil && x.Total != nil {
		return *x.Total
	}
	return 0
}

func (x *Memory) GetProcess() uint64 {
	if x != nil && x.Process != nil {
		return *x.Process
	}
	return 0
}

type C41219OS struct {
	Type    *uint32 `protobuf:"varint,1,opt"`
	Version *string `protobuf:"bytes,2,opt"`
	Sdk     *string `protobuf:"bytes,3,opt"`
	Kernel  *string `protobuf:"bytes,4,opt"`
	Rom     *string `protobuf:"bytes,5,opt"`
}

func (x *C41219OS) GetType() uint32 {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return 0
}

func (x *C41219OS) GetVersion() string {
	if x != nil && x.Version != nil {
		return *x.Version
	}
	return ""
}

func (x *C41219OS) GetSdk() string {
	if x != nil && x.Sdk != nil {
		return *x.Sdk
	}
	return ""
}

func (x *C41219OS) GetKernel() string {
	if x != nil && x.Kernel != nil {
		return *x.Kernel
	}
	return ""
}

func (x *C41219OS) GetRom() string {
	if x != nil && x.Rom != nil {
		return *x.Rom
	}
	return ""
}

type QueryUinPackageUsageReq struct {
	Type        *uint32 `protobuf:"varint,1,opt"`
	UinFileSize *uint64 `protobuf:"varint,2,opt"`
}

func (x *QueryUinPackageUsageReq) GetType() uint32 {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return 0
}

func (x *QueryUinPackageUsageReq) GetUinFileSize() uint64 {
	if x != nil && x.UinFileSize != nil {
		return *x.UinFileSize
	}
	return 0
}

type QueryUinPackageUsageRsp struct {
	Status             *uint32               `protobuf:"varint,1,opt"`
	LeftUinNum         *uint64               `protobuf:"varint,2,opt"`
	MaxUinNum          *uint64               `protobuf:"varint,3,opt"`
	Proportion         *uint32               `protobuf:"varint,4,opt"`
	UinPackageUsedList []*UinPackageUsedInfo `protobuf:"bytes,10,rep"`
}

func (x *QueryUinPackageUsageRsp) GetStatus() uint32 {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return 0
}

func (x *QueryUinPackageUsageRsp) GetLeftUinNum() uint64 {
	if x != nil && x.LeftUinNum != nil {
		return *x.LeftUinNum
	}
	return 0
}

func (x *QueryUinPackageUsageRsp) GetMaxUinNum() uint64 {
	if x != nil && x.MaxUinNum != nil {
		return *x.MaxUinNum
	}
	return 0
}

func (x *QueryUinPackageUsageRsp) GetProportion() uint32 {
	if x != nil && x.Proportion != nil {
		return *x.Proportion
	}
	return 0
}

func (x *QueryUinPackageUsageRsp) GetUinPackageUsedList() []*UinPackageUsedInfo {
	if x != nil {
		return x.UinPackageUsedList
	}
	return nil
}

type D769ReqBody struct {
	ConfigList              []*D769ConfigSeq         `protobuf:"bytes,1,rep"`
	DeviceInfo              *D769DeviceInfo          `protobuf:"bytes,2,opt"`
	Info                    *string                  `protobuf:"bytes,3,opt"`
	Province                *string                  `protobuf:"bytes,4,opt"`
	City                    *string                  `protobuf:"bytes,5,opt"`
	ReqDebugMsg             *int32                   `protobuf:"varint,6,opt"`
	QueryUinPackageUsageReq *QueryUinPackageUsageReq `protobuf:"bytes,101,opt"`
}

func (x *D769ReqBody) GetConfigList() []*D769ConfigSeq {
	if x != nil {
		return x.ConfigList
	}
	return nil
}

func (x *D769ReqBody) GetDeviceInfo() *D769DeviceInfo {
	if x != nil {
		return x.DeviceInfo
	}
	return nil
}

func (x *D769ReqBody) GetInfo() string {
	if x != nil && x.Info != nil {
		return *x.Info
	}
	return ""
}

func (x *D769ReqBody) GetProvince() string {
	if x != nil && x.Province != nil {
		return *x.Province
	}
	return ""
}

func (x *D769ReqBody) GetCity() string {
	if x != nil && x.City != nil {
		return *x.City
	}
	return ""
}

func (x *D769ReqBody) GetReqDebugMsg() int32 {
	if x != nil && x.ReqDebugMsg != nil {
		return *x.ReqDebugMsg
	}
	return 0
}

func (x *D769ReqBody) GetQueryUinPackageUsageReq() *QueryUinPackageUsageReq {
	if x != nil {
		return x.QueryUinPackageUsageReq
	}
	return nil
}

type D769RspBody struct {
	Result                  *uint32                  `protobuf:"varint,1,opt"`
	ConfigList              []*D769ConfigSeq         `protobuf:"bytes,2,rep"`
	QueryUinPackageUsageRsp *QueryUinPackageUsageRsp `protobuf:"bytes,101,opt"`
}

func (x *D769RspBody) GetResult() uint32 {
	if x != nil && x.Result != nil {
		return *x.Result
	}
	return 0
}

func (x *D769RspBody) GetConfigList() []*D769ConfigSeq {
	if x != nil {
		return x.ConfigList
	}
	return nil
}

func (x *D769RspBody) GetQueryUinPackageUsageRsp() *QueryUinPackageUsageRsp {
	if x != nil {
		return x.QueryUinPackageUsageRsp
	}
	return nil
}

type Screen struct {
	Model      *string `protobuf:"bytes,1,opt"`
	Width      *uint32 `protobuf:"varint,2,opt"`
	Height     *uint32 `protobuf:"varint,3,opt"`
	Dpi        *uint32 `protobuf:"varint,4,opt"`
	MultiTouch *bool   `protobuf:"varint,5,opt"`
}

func (x *Screen) GetModel() string {
	if x != nil && x.Model != nil {
		return *x.Model
	}
	return ""
}

func (x *Screen) GetWidth() uint32 {
	if x != nil && x.Width != nil {
		return *x.Width
	}
	return 0
}

func (x *Screen) GetHeight() uint32 {
	if x != nil && x.Height != nil {
		return *x.Height
	}
	return 0
}

func (x *Screen) GetDpi() uint32 {
	if x != nil && x.Dpi != nil {
		return *x.Dpi
	}
	return 0
}

func (x *Screen) GetMultiTouch() bool {
	if x != nil && x.MultiTouch != nil {
		return *x.MultiTouch
	}
	return false
}

type Storage struct {
	Builtin  *uint64 `protobuf:"varint,1,opt"`
	External *uint64 `protobuf:"varint,2,opt"`
}

func (x *Storage) GetBuiltin() uint64 {
	if x != nil && x.Builtin != nil {
		return *x.Builtin
	}
	return 0
}

func (x *Storage) GetExternal() uint64 {
	if x != nil && x.External != nil {
		return *x.External
	}
	return 0
}

type UinPackageUsedInfo struct {
	RuleId *uint32 `protobuf:"varint,1,opt"`
	Author *string `protobuf:"bytes,2,opt"`
	Url    *string `protobuf:"bytes,3,opt"`
	UinNum *uint64 `protobuf:"varint,4,opt"`
}

func (x *UinPackageUsedInfo) GetRuleId() uint32 {
	if x != nil && x.RuleId != nil {
		return *x.RuleId
	}
	return 0
}

func (x *UinPackageUsedInfo) GetAuthor() string {
	if x != nil && x.Author != nil {
		return *x.Author
	}
	return ""
}

func (x *UinPackageUsedInfo) GetUrl() string {
	if x != nil && x.Url != nil {
		return *x.Url
	}
	return ""
}

func (x *UinPackageUsedInfo) GetUinNum() uint64 {
	if x != nil && x.UinNum != nil {
		return *x.UinNum
	}
	return 0
}
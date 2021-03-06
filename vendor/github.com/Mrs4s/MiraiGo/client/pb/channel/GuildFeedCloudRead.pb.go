// Code generated by protoc-gen-golite. DO NOT EDIT.
// source: pb/channel/GuildFeedCloudRead.proto

package channel

type GetNoticesReq struct {
	ExtInfo    *StCommonExt `protobuf:"bytes,1,opt"`
	PageNum    *uint32      `protobuf:"varint,2,opt"`
	AttachInfo *string      `protobuf:"bytes,3,opt"`
}

func (x *GetNoticesReq) GetExtInfo() *StCommonExt {
	if x != nil {
		return x.ExtInfo
	}
	return nil
}

func (x *GetNoticesReq) GetPageNum() uint32 {
	if x != nil && x.PageNum != nil {
		return *x.PageNum
	}
	return 0
}

func (x *GetNoticesReq) GetAttachInfo() string {
	if x != nil && x.AttachInfo != nil {
		return *x.AttachInfo
	}
	return ""
}

type GetNoticesRsp struct {
	ExtInfo    *StCommonExt `protobuf:"bytes,1,opt"`
	Notices    []*StNotice  `protobuf:"bytes,2,rep"`
	TotalNum   *uint32      `protobuf:"varint,3,opt"`
	IsFinish   *bool        `protobuf:"varint,4,opt"`
	AttachInfo *string      `protobuf:"bytes,5,opt"`
}

func (x *GetNoticesRsp) GetExtInfo() *StCommonExt {
	if x != nil {
		return x.ExtInfo
	}
	return nil
}

func (x *GetNoticesRsp) GetNotices() []*StNotice {
	if x != nil {
		return x.Notices
	}
	return nil
}

func (x *GetNoticesRsp) GetTotalNum() uint32 {
	if x != nil && x.TotalNum != nil {
		return *x.TotalNum
	}
	return 0
}

func (x *GetNoticesRsp) GetIsFinish() bool {
	if x != nil && x.IsFinish != nil {
		return *x.IsFinish
	}
	return false
}

func (x *GetNoticesRsp) GetAttachInfo() string {
	if x != nil && x.AttachInfo != nil {
		return *x.AttachInfo
	}
	return ""
}

type NeedInsertCommentInfo struct {
	CommentID *string `protobuf:"bytes,1,opt"`
}

func (x *NeedInsertCommentInfo) GetCommentID() string {
	if x != nil && x.CommentID != nil {
		return *x.CommentID
	}
	return ""
}

type RefreshToast struct {
	Text *string `protobuf:"bytes,1,opt"`
}

func (x *RefreshToast) GetText() string {
	if x != nil && x.Text != nil {
		return *x.Text
	}
	return ""
}

type StGetChannelFeedsReq struct {
	ExtInfo       *StCommonExt   `protobuf:"bytes,1,opt"`
	Count         *uint32        `protobuf:"varint,2,opt"`
	From          *uint32        `protobuf:"varint,3,opt"`
	ChannelSign   *StChannelSign `protobuf:"bytes,4,opt"`
	FeedAttchInfo *string        `protobuf:"bytes,5,opt"`
}

func (x *StGetChannelFeedsReq) GetExtInfo() *StCommonExt {
	if x != nil {
		return x.ExtInfo
	}
	return nil
}

func (x *StGetChannelFeedsReq) GetCount() uint32 {
	if x != nil && x.Count != nil {
		return *x.Count
	}
	return 0
}

func (x *StGetChannelFeedsReq) GetFrom() uint32 {
	if x != nil && x.From != nil {
		return *x.From
	}
	return 0
}

func (x *StGetChannelFeedsReq) GetChannelSign() *StChannelSign {
	if x != nil {
		return x.ChannelSign
	}
	return nil
}

func (x *StGetChannelFeedsReq) GetFeedAttchInfo() string {
	if x != nil && x.FeedAttchInfo != nil {
		return *x.FeedAttchInfo
	}
	return ""
}

type StGetChannelFeedsRsp struct {
	ExtInfo       *StCommonExt  `protobuf:"bytes,1,opt"`
	VecFeed       []*StFeed     `protobuf:"bytes,2,rep"`
	IsFinish      *uint32       `protobuf:"varint,3,opt"`
	User          *StUser       `protobuf:"bytes,4,opt"`
	FeedAttchInfo *string       `protobuf:"bytes,5,opt"`
	RefreshToast  *RefreshToast `protobuf:"bytes,6,opt"`
}

func (x *StGetChannelFeedsRsp) GetExtInfo() *StCommonExt {
	if x != nil {
		return x.ExtInfo
	}
	return nil
}

func (x *StGetChannelFeedsRsp) GetVecFeed() []*StFeed {
	if x != nil {
		return x.VecFeed
	}
	return nil
}

func (x *StGetChannelFeedsRsp) GetIsFinish() uint32 {
	if x != nil && x.IsFinish != nil {
		return *x.IsFinish
	}
	return 0
}

func (x *StGetChannelFeedsRsp) GetUser() *StUser {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *StGetChannelFeedsRsp) GetFeedAttchInfo() string {
	if x != nil && x.FeedAttchInfo != nil {
		return *x.FeedAttchInfo
	}
	return ""
}

func (x *StGetChannelFeedsRsp) GetRefreshToast() *RefreshToast {
	if x != nil {
		return x.RefreshToast
	}
	return nil
}

type StGetChannelShareFeedReq struct {
	ExtInfo          *StCommonExt        `protobuf:"bytes,1,opt"`
	From             *uint32             `protobuf:"varint,2,opt"`
	ChannelShareInfo *StChannelShareInfo `protobuf:"bytes,3,opt"`
}

func (x *StGetChannelShareFeedReq) GetExtInfo() *StCommonExt {
	if x != nil {
		return x.ExtInfo
	}
	return nil
}

func (x *StGetChannelShareFeedReq) GetFrom() uint32 {
	if x != nil && x.From != nil {
		return *x.From
	}
	return 0
}

func (x *StGetChannelShareFeedReq) GetChannelShareInfo() *StChannelShareInfo {
	if x != nil {
		return x.ChannelShareInfo
	}
	return nil
}

type StGetChannelShareFeedRsp struct {
	ExtInfo *StCommonExt `protobuf:"bytes,1,opt"`
	Feed    *StFeed      `protobuf:"bytes,2,opt"`
}

func (x *StGetChannelShareFeedRsp) GetExtInfo() *StCommonExt {
	if x != nil {
		return x.ExtInfo
	}
	return nil
}

func (x *StGetChannelShareFeedRsp) GetFeed() *StFeed {
	if x != nil {
		return x.Feed
	}
	return nil
}

type StGetFeedCommentsReq struct {
	ExtInfo     *StCommonExt `protobuf:"bytes,1,opt"`
	UserId      *string      `protobuf:"bytes,2,opt"`
	FeedId      *string      `protobuf:"bytes,3,opt"`
	ListNum     *uint32      `protobuf:"varint,4,opt"`
	From        *uint32      `protobuf:"varint,5,opt"`
	AttchInfo   *string      `protobuf:"bytes,6,opt"`
	EntrySchema *string      `protobuf:"bytes,7,opt"`
}

func (x *StGetFeedCommentsReq) GetExtInfo() *StCommonExt {
	if x != nil {
		return x.ExtInfo
	}
	return nil
}

func (x *StGetFeedCommentsReq) GetUserId() string {
	if x != nil && x.UserId != nil {
		return *x.UserId
	}
	return ""
}

func (x *StGetFeedCommentsReq) GetFeedId() string {
	if x != nil && x.FeedId != nil {
		return *x.FeedId
	}
	return ""
}

func (x *StGetFeedCommentsReq) GetListNum() uint32 {
	if x != nil && x.ListNum != nil {
		return *x.ListNum
	}
	return 0
}

func (x *StGetFeedCommentsReq) GetFrom() uint32 {
	if x != nil && x.From != nil {
		return *x.From
	}
	return 0
}

func (x *StGetFeedCommentsReq) GetAttchInfo() string {
	if x != nil && x.AttchInfo != nil {
		return *x.AttchInfo
	}
	return ""
}

func (x *StGetFeedCommentsReq) GetEntrySchema() string {
	if x != nil && x.EntrySchema != nil {
		return *x.EntrySchema
	}
	return ""
}

type StGetFeedCommentsRsp struct {
	ExtInfo    *StCommonExt `protobuf:"bytes,1,opt"`
	VecComment []*StComment `protobuf:"bytes,2,rep"`
	TotalNum   *uint32      `protobuf:"varint,3,opt"`
	IsFinish   *uint32      `protobuf:"varint,4,opt"`
	AttchInfo  *string      `protobuf:"bytes,5,opt"`
}

func (x *StGetFeedCommentsRsp) GetExtInfo() *StCommonExt {
	if x != nil {
		return x.ExtInfo
	}
	return nil
}

func (x *StGetFeedCommentsRsp) GetVecComment() []*StComment {
	if x != nil {
		return x.VecComment
	}
	return nil
}

func (x *StGetFeedCommentsRsp) GetTotalNum() uint32 {
	if x != nil && x.TotalNum != nil {
		return *x.TotalNum
	}
	return 0
}

func (x *StGetFeedCommentsRsp) GetIsFinish() uint32 {
	if x != nil && x.IsFinish != nil {
		return *x.IsFinish
	}
	return 0
}

func (x *StGetFeedCommentsRsp) GetAttchInfo() string {
	if x != nil && x.AttchInfo != nil {
		return *x.AttchInfo
	}
	return ""
}

type StGetFeedDetailReq struct {
	ExtInfo     *StCommonExt   `protobuf:"bytes,1,opt"`
	From        *uint32        `protobuf:"varint,2,opt"`
	UserId      *string        `protobuf:"bytes,3,opt"`
	FeedId      *string        `protobuf:"bytes,4,opt"`
	CreateTime  *uint64        `protobuf:"varint,5,opt"`
	DetailType  *uint32        `protobuf:"varint,6,opt"`
	ChannelSign *StChannelSign `protobuf:"bytes,7,opt"`
}

func (x *StGetFeedDetailReq) GetExtInfo() *StCommonExt {
	if x != nil {
		return x.ExtInfo
	}
	return nil
}

func (x *StGetFeedDetailReq) GetFrom() uint32 {
	if x != nil && x.From != nil {
		return *x.From
	}
	return 0
}

func (x *StGetFeedDetailReq) GetUserId() string {
	if x != nil && x.UserId != nil {
		return *x.UserId
	}
	return ""
}

func (x *StGetFeedDetailReq) GetFeedId() string {
	if x != nil && x.FeedId != nil {
		return *x.FeedId
	}
	return ""
}

func (x *StGetFeedDetailReq) GetCreateTime() uint64 {
	if x != nil && x.CreateTime != nil {
		return *x.CreateTime
	}
	return 0
}

func (x *StGetFeedDetailReq) GetDetailType() uint32 {
	if x != nil && x.DetailType != nil {
		return *x.DetailType
	}
	return 0
}

func (x *StGetFeedDetailReq) GetChannelSign() *StChannelSign {
	if x != nil {
		return x.ChannelSign
	}
	return nil
}

type StGetFeedDetailRsp struct {
	ExtInfo   *StCommonExt `protobuf:"bytes,1,opt"`
	Feed      *StFeed      `protobuf:"bytes,2,opt"`
	LoginUser *StUser      `protobuf:"bytes,3,opt"`
}

func (x *StGetFeedDetailRsp) GetExtInfo() *StCommonExt {
	if x != nil {
		return x.ExtInfo
	}
	return nil
}

func (x *StGetFeedDetailRsp) GetFeed() *StFeed {
	if x != nil {
		return x.Feed
	}
	return nil
}

func (x *StGetFeedDetailRsp) GetLoginUser() *StUser {
	if x != nil {
		return x.LoginUser
	}
	return nil
}

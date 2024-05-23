// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.27.0--rc3
// source: follower/follower.proto

package follower

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UserFollowingDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     int32 `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	FollowerId int32 `protobuf:"varint,2,opt,name=FollowerId,proto3" json:"FollowerId,omitempty"`
}

func (x *UserFollowingDto) Reset() {
	*x = UserFollowingDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_follower_follower_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserFollowingDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserFollowingDto) ProtoMessage() {}

func (x *UserFollowingDto) ProtoReflect() protoreflect.Message {
	mi := &file_follower_follower_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserFollowingDto.ProtoReflect.Descriptor instead.
func (*UserFollowingDto) Descriptor() ([]byte, []int) {
	return file_follower_follower_proto_rawDescGZIP(), []int{0}
}

func (x *UserFollowingDto) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UserFollowingDto) GetFollowerId() int32 {
	if x != nil {
		return x.FollowerId
	}
	return 0
}

type NeoFollowerDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId            string `protobuf:"bytes,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	Username          string `protobuf:"bytes,2,opt,name=Username,proto3" json:"Username,omitempty"`
	FollowingUserId   string `protobuf:"bytes,3,opt,name=FollowingUserId,proto3" json:"FollowingUserId,omitempty"`
	FollowingUsername string `protobuf:"bytes,4,opt,name=FollowingUsername,proto3" json:"FollowingUsername,omitempty"`
}

func (x *NeoFollowerDto) Reset() {
	*x = NeoFollowerDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_follower_follower_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NeoFollowerDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NeoFollowerDto) ProtoMessage() {}

func (x *NeoFollowerDto) ProtoReflect() protoreflect.Message {
	mi := &file_follower_follower_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NeoFollowerDto.ProtoReflect.Descriptor instead.
func (*NeoFollowerDto) Descriptor() ([]byte, []int) {
	return file_follower_follower_proto_rawDescGZIP(), []int{1}
}

func (x *NeoFollowerDto) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *NeoFollowerDto) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *NeoFollowerDto) GetFollowingUserId() string {
	if x != nil {
		return x.FollowingUserId
	}
	return ""
}

func (x *NeoFollowerDto) GetFollowingUsername() string {
	if x != nil {
		return x.FollowingUsername
	}
	return ""
}

type NeoUserDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int32  `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=Username,proto3" json:"Username,omitempty"`
}

func (x *NeoUserDto) Reset() {
	*x = NeoUserDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_follower_follower_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NeoUserDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NeoUserDto) ProtoMessage() {}

func (x *NeoUserDto) ProtoReflect() protoreflect.Message {
	mi := &file_follower_follower_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NeoUserDto.ProtoReflect.Descriptor instead.
func (*NeoUserDto) Descriptor() ([]byte, []int) {
	return file_follower_follower_proto_rawDescGZIP(), []int{2}
}

func (x *NeoUserDto) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *NeoUserDto) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type BlogPostDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             int32                  `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	AuthorId       int32                  `protobuf:"varint,2,opt,name=AuthorId,proto3" json:"AuthorId,omitempty"`
	TourId         int32                  `protobuf:"varint,3,opt,name=TourId,proto3" json:"TourId,omitempty"`
	AuthorUsername string                 `protobuf:"bytes,4,opt,name=AuthorUsername,proto3" json:"AuthorUsername,omitempty"`
	Title          string                 `protobuf:"bytes,5,opt,name=Title,proto3" json:"Title,omitempty"`
	Description    string                 `protobuf:"bytes,6,opt,name=Description,proto3" json:"Description,omitempty"`
	CreationDate   *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=CreationDate,proto3" json:"CreationDate,omitempty"`
	ImageURLs      []string               `protobuf:"bytes,8,rep,name=ImageURLs,proto3" json:"ImageURLs,omitempty"`
	Comments       []*BlogPostCommentDto  `protobuf:"bytes,9,rep,name=Comments,proto3" json:"Comments,omitempty"`
	Ratings        []*BlogPostRatingDto   `protobuf:"bytes,10,rep,name=Ratings,proto3" json:"Ratings,omitempty"`
	Status         string                 `protobuf:"bytes,11,opt,name=Status,proto3" json:"Status,omitempty"`
}

func (x *BlogPostDto) Reset() {
	*x = BlogPostDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_follower_follower_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlogPostDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlogPostDto) ProtoMessage() {}

func (x *BlogPostDto) ProtoReflect() protoreflect.Message {
	mi := &file_follower_follower_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlogPostDto.ProtoReflect.Descriptor instead.
func (*BlogPostDto) Descriptor() ([]byte, []int) {
	return file_follower_follower_proto_rawDescGZIP(), []int{3}
}

func (x *BlogPostDto) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *BlogPostDto) GetAuthorId() int32 {
	if x != nil {
		return x.AuthorId
	}
	return 0
}

func (x *BlogPostDto) GetTourId() int32 {
	if x != nil {
		return x.TourId
	}
	return 0
}

func (x *BlogPostDto) GetAuthorUsername() string {
	if x != nil {
		return x.AuthorUsername
	}
	return ""
}

func (x *BlogPostDto) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *BlogPostDto) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *BlogPostDto) GetCreationDate() *timestamppb.Timestamp {
	if x != nil {
		return x.CreationDate
	}
	return nil
}

func (x *BlogPostDto) GetImageURLs() []string {
	if x != nil {
		return x.ImageURLs
	}
	return nil
}

func (x *BlogPostDto) GetComments() []*BlogPostCommentDto {
	if x != nil {
		return x.Comments
	}
	return nil
}

func (x *BlogPostDto) GetRatings() []*BlogPostRatingDto {
	if x != nil {
		return x.Ratings
	}
	return nil
}

func (x *BlogPostDto) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type BlogPostCommentDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text            string                 `protobuf:"bytes,1,opt,name=Text,proto3" json:"Text,omitempty"`
	UserId          int32                  `protobuf:"varint,2,opt,name=UserId,proto3" json:"UserId,omitempty"`
	Username        string                 `protobuf:"bytes,3,opt,name=Username,proto3" json:"Username,omitempty"`
	CreationTime    *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=CreationTime,proto3" json:"CreationTime,omitempty"`
	LastUpdatedTime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=LastUpdatedTime,proto3" json:"LastUpdatedTime,omitempty"`
}

func (x *BlogPostCommentDto) Reset() {
	*x = BlogPostCommentDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_follower_follower_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlogPostCommentDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlogPostCommentDto) ProtoMessage() {}

func (x *BlogPostCommentDto) ProtoReflect() protoreflect.Message {
	mi := &file_follower_follower_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlogPostCommentDto.ProtoReflect.Descriptor instead.
func (*BlogPostCommentDto) Descriptor() ([]byte, []int) {
	return file_follower_follower_proto_rawDescGZIP(), []int{4}
}

func (x *BlogPostCommentDto) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *BlogPostCommentDto) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *BlogPostCommentDto) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *BlogPostCommentDto) GetCreationTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreationTime
	}
	return nil
}

func (x *BlogPostCommentDto) GetLastUpdatedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.LastUpdatedTime
	}
	return nil
}

type BlogPostRatingDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsPositive   bool                   `protobuf:"varint,1,opt,name=IsPositive,proto3" json:"IsPositive,omitempty"`
	CreationTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=CreationTime,proto3" json:"CreationTime,omitempty"`
	UserId       int32                  `protobuf:"varint,3,opt,name=UserId,proto3" json:"UserId,omitempty"`
}

func (x *BlogPostRatingDto) Reset() {
	*x = BlogPostRatingDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_follower_follower_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlogPostRatingDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlogPostRatingDto) ProtoMessage() {}

func (x *BlogPostRatingDto) ProtoReflect() protoreflect.Message {
	mi := &file_follower_follower_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlogPostRatingDto.ProtoReflect.Descriptor instead.
func (*BlogPostRatingDto) Descriptor() ([]byte, []int) {
	return file_follower_follower_proto_rawDescGZIP(), []int{5}
}

func (x *BlogPostRatingDto) GetIsPositive() bool {
	if x != nil {
		return x.IsPositive
	}
	return false
}

func (x *BlogPostRatingDto) GetCreationTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreationTime
	}
	return nil
}

func (x *BlogPostRatingDto) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type Id struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *Id) Reset() {
	*x = Id{}
	if protoimpl.UnsafeEnabled {
		mi := &file_follower_follower_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Id) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Id) ProtoMessage() {}

func (x *Id) ProtoReflect() protoreflect.Message {
	mi := &file_follower_follower_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Id.ProtoReflect.Descriptor instead.
func (*Id) Descriptor() ([]byte, []int) {
	return file_follower_follower_proto_rawDescGZIP(), []int{6}
}

func (x *Id) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ListBlogPostDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResponseList []*BlogPostDto `protobuf:"bytes,1,rep,name=ResponseList,proto3" json:"ResponseList,omitempty"`
}

func (x *ListBlogPostDto) Reset() {
	*x = ListBlogPostDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_follower_follower_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBlogPostDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBlogPostDto) ProtoMessage() {}

func (x *ListBlogPostDto) ProtoReflect() protoreflect.Message {
	mi := &file_follower_follower_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBlogPostDto.ProtoReflect.Descriptor instead.
func (*ListBlogPostDto) Descriptor() ([]byte, []int) {
	return file_follower_follower_proto_rawDescGZIP(), []int{7}
}

func (x *ListBlogPostDto) GetResponseList() []*BlogPostDto {
	if x != nil {
		return x.ResponseList
	}
	return nil
}

type ListNeoUserDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResponseList []*NeoUserDto `protobuf:"bytes,1,rep,name=ResponseList,proto3" json:"ResponseList,omitempty"`
}

func (x *ListNeoUserDto) Reset() {
	*x = ListNeoUserDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_follower_follower_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListNeoUserDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListNeoUserDto) ProtoMessage() {}

func (x *ListNeoUserDto) ProtoReflect() protoreflect.Message {
	mi := &file_follower_follower_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListNeoUserDto.ProtoReflect.Descriptor instead.
func (*ListNeoUserDto) Descriptor() ([]byte, []int) {
	return file_follower_follower_proto_rawDescGZIP(), []int{8}
}

func (x *ListNeoUserDto) GetResponseList() []*NeoUserDto {
	if x != nil {
		return x.ResponseList
	}
	return nil
}

var File_follower_follower_proto protoreflect.FileDescriptor

var file_follower_follower_proto_rawDesc = []byte{
	0x0a, 0x17, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x2f, 0x66, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4a, 0x0a, 0x10, 0x55, 0x73,
	0x65, 0x72, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x44, 0x74, 0x6f, 0x12, 0x16,
	0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x46, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x65, 0x72, 0x49, 0x64, 0x22, 0x9c, 0x01, 0x0a, 0x0e, 0x4e, 0x65, 0x6f, 0x46, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x44, 0x74, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a,
	0x0f, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e,
	0x67, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x11, 0x46, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x69, 0x6e, 0x67, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x11, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x55, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x38, 0x0a, 0x0a, 0x4e, 0x65, 0x6f, 0x55, 0x73, 0x65, 0x72,
	0x44, 0x74, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x86, 0x03, 0x0a, 0x0b, 0x42, 0x6c, 0x6f, 0x67, 0x50, 0x6f, 0x73, 0x74, 0x44, 0x74, 0x6f, 0x12,
	0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x54,
	0x6f, 0x75, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x54, 0x6f, 0x75,
	0x72, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x55, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x41, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x54,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x3e, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44,
	0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44,
	0x61, 0x74, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x52, 0x4c, 0x73,
	0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x52, 0x4c,
	0x73, 0x12, 0x2f, 0x0a, 0x08, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x09, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x42, 0x6c, 0x6f, 0x67, 0x50, 0x6f, 0x73, 0x74, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x74, 0x6f, 0x52, 0x08, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x12, 0x2c, 0x0a, 0x07, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x0a, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x42, 0x6c, 0x6f, 0x67, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x61,
	0x74, 0x69, 0x6e, 0x67, 0x44, 0x74, 0x6f, 0x52, 0x07, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xe2, 0x01, 0x0a, 0x12, 0x42, 0x6c, 0x6f,
	0x67, 0x50, 0x6f, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x74, 0x6f, 0x12,
	0x12, 0x0a, 0x04, 0x54, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54,
	0x65, 0x78, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x55,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3e, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x44, 0x0a, 0x0f, 0x4c, 0x61, 0x73, 0x74, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0f, 0x4c, 0x61,
	0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x8b, 0x01,
	0x0a, 0x11, 0x42, 0x6c, 0x6f, 0x67, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67,
	0x44, 0x74, 0x6f, 0x12, 0x1e, 0x0a, 0x0a, 0x49, 0x73, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x76,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x49, 0x73, 0x50, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x76, 0x65, 0x12, 0x3e, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x14, 0x0a, 0x02, 0x69,
	0x64, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49,
	0x64, 0x22, 0x43, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x67, 0x50, 0x6f, 0x73,
	0x74, 0x44, 0x74, 0x6f, 0x12, 0x30, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x42, 0x6c, 0x6f,
	0x67, 0x50, 0x6f, 0x73, 0x74, 0x44, 0x74, 0x6f, 0x52, 0x0c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x41, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x65,
	0x6f, 0x55, 0x73, 0x65, 0x72, 0x44, 0x74, 0x6f, 0x12, 0x2f, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b,
	0x2e, 0x4e, 0x65, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x44, 0x74, 0x6f, 0x52, 0x0c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x32, 0xd9, 0x01, 0x0a, 0x08, 0x46, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x12, 0x3a, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4e, 0x65, 0x77, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x12, 0x11, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x44, 0x74, 0x6f, 0x1a,
	0x0f, 0x2e, 0x4e, 0x65, 0x6f, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x44, 0x74, 0x6f,
	0x22, 0x00, 0x12, 0x30, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x03, 0x2e, 0x69,
	0x64, 0x1a, 0x0f, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x65, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x44,
	0x74, 0x6f, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x69, 0x6e, 0x67, 0x73, 0x57, 0x69, 0x74, 0x68, 0x42, 0x6c, 0x6f, 0x67, 0x73, 0x12, 0x03,
	0x2e, 0x69, 0x64, 0x1a, 0x10, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x67, 0x50, 0x6f,
	0x73, 0x74, 0x44, 0x74, 0x6f, 0x22, 0x00, 0x12, 0x2c, 0x0a, 0x12, 0x46, 0x69, 0x6e, 0x64, 0x55,
	0x73, 0x65, 0x72, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x03, 0x2e,
	0x69, 0x64, 0x1a, 0x0f, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x65, 0x6f, 0x55, 0x73, 0x65, 0x72,
	0x44, 0x74, 0x6f, 0x22, 0x00, 0x42, 0x10, 0x5a, 0x0e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_follower_follower_proto_rawDescOnce sync.Once
	file_follower_follower_proto_rawDescData = file_follower_follower_proto_rawDesc
)

func file_follower_follower_proto_rawDescGZIP() []byte {
	file_follower_follower_proto_rawDescOnce.Do(func() {
		file_follower_follower_proto_rawDescData = protoimpl.X.CompressGZIP(file_follower_follower_proto_rawDescData)
	})
	return file_follower_follower_proto_rawDescData
}

var file_follower_follower_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_follower_follower_proto_goTypes = []interface{}{
	(*UserFollowingDto)(nil),      // 0: UserFollowingDto
	(*NeoFollowerDto)(nil),        // 1: NeoFollowerDto
	(*NeoUserDto)(nil),            // 2: NeoUserDto
	(*BlogPostDto)(nil),           // 3: BlogPostDto
	(*BlogPostCommentDto)(nil),    // 4: BlogPostCommentDto
	(*BlogPostRatingDto)(nil),     // 5: BlogPostRatingDto
	(*Id)(nil),                    // 6: id
	(*ListBlogPostDto)(nil),       // 7: ListBlogPostDto
	(*ListNeoUserDto)(nil),        // 8: ListNeoUserDto
	(*timestamppb.Timestamp)(nil), // 9: google.protobuf.Timestamp
}
var file_follower_follower_proto_depIdxs = []int32{
	9,  // 0: BlogPostDto.CreationDate:type_name -> google.protobuf.Timestamp
	4,  // 1: BlogPostDto.Comments:type_name -> BlogPostCommentDto
	5,  // 2: BlogPostDto.Ratings:type_name -> BlogPostRatingDto
	9,  // 3: BlogPostCommentDto.CreationTime:type_name -> google.protobuf.Timestamp
	9,  // 4: BlogPostCommentDto.LastUpdatedTime:type_name -> google.protobuf.Timestamp
	9,  // 5: BlogPostRatingDto.CreationTime:type_name -> google.protobuf.Timestamp
	3,  // 6: ListBlogPostDto.ResponseList:type_name -> BlogPostDto
	2,  // 7: ListNeoUserDto.ResponseList:type_name -> NeoUserDto
	0,  // 8: Follower.CreateNewFollowing:input_type -> UserFollowingDto
	6,  // 9: Follower.GetUserRecommendations:input_type -> id
	6,  // 10: Follower.GetFollowingsWithBlogs:input_type -> id
	6,  // 11: Follower.FindUserFollowings:input_type -> id
	1,  // 12: Follower.CreateNewFollowing:output_type -> NeoFollowerDto
	8,  // 13: Follower.GetUserRecommendations:output_type -> ListNeoUserDto
	7,  // 14: Follower.GetFollowingsWithBlogs:output_type -> ListBlogPostDto
	8,  // 15: Follower.FindUserFollowings:output_type -> ListNeoUserDto
	12, // [12:16] is the sub-list for method output_type
	8,  // [8:12] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_follower_follower_proto_init() }
func file_follower_follower_proto_init() {
	if File_follower_follower_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_follower_follower_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserFollowingDto); i {
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
		file_follower_follower_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NeoFollowerDto); i {
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
		file_follower_follower_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NeoUserDto); i {
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
		file_follower_follower_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlogPostDto); i {
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
		file_follower_follower_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlogPostCommentDto); i {
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
		file_follower_follower_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlogPostRatingDto); i {
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
		file_follower_follower_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Id); i {
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
		file_follower_follower_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBlogPostDto); i {
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
		file_follower_follower_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListNeoUserDto); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_follower_follower_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_follower_follower_proto_goTypes,
		DependencyIndexes: file_follower_follower_proto_depIdxs,
		MessageInfos:      file_follower_follower_proto_msgTypes,
	}.Build()
	File_follower_follower_proto = out.File
	file_follower_follower_proto_rawDesc = nil
	file_follower_follower_proto_goTypes = nil
	file_follower_follower_proto_depIdxs = nil
}
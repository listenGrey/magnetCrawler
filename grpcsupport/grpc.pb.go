// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.0--rc1
// source: grpcsupport/grpc.proto

package grpcsupport

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Tag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content string `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	Url     string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *Tag) Reset() {
	*x = Tag{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpcsupport_grpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tag) ProtoMessage() {}

func (x *Tag) ProtoReflect() protoreflect.Message {
	mi := &file_grpcsupport_grpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tag.ProtoReflect.Descriptor instead.
func (*Tag) Descriptor() ([]byte, []int) {
	return file_grpcsupport_grpc_proto_rawDescGZIP(), []int{0}
}

func (x *Tag) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Tag) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type Magnet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content string `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	Url     string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Data    string `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Magnet) Reset() {
	*x = Magnet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpcsupport_grpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Magnet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Magnet) ProtoMessage() {}

func (x *Magnet) ProtoReflect() protoreflect.Message {
	mi := &file_grpcsupport_grpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Magnet.ProtoReflect.Descriptor instead.
func (*Magnet) Descriptor() ([]byte, []int) {
	return file_grpcsupport_grpc_proto_rawDescGZIP(), []int{1}
}

func (x *Magnet) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Magnet) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Magnet) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type Profile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url         string    `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Id          string    `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Code        string    `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	Title       string    `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	Cover       string    `protobuf:"bytes,5,opt,name=cover,proto3" json:"cover,omitempty"`
	ReleaseData string    `protobuf:"bytes,6,opt,name=releaseData,proto3" json:"releaseData,omitempty"`
	Duration    string    `protobuf:"bytes,7,opt,name=duration,proto3" json:"duration,omitempty"`
	Director    string    `protobuf:"bytes,8,opt,name=director,proto3" json:"director,omitempty"`
	Company     *Tag      `protobuf:"bytes,9,opt,name=company,proto3" json:"company,omitempty"`
	Series      *Tag      `protobuf:"bytes,10,opt,name=series,proto3" json:"series,omitempty"`
	Stars       string    `protobuf:"bytes,11,opt,name=stars,proto3" json:"stars,omitempty"`
	StarsPerson string    `protobuf:"bytes,12,opt,name=starsPerson,proto3" json:"starsPerson,omitempty"`
	Type        []*Tag    `protobuf:"bytes,13,rep,name=type,proto3" json:"type,omitempty"`
	Actor       []*Tag    `protobuf:"bytes,14,rep,name=actor,proto3" json:"actor,omitempty"`
	FetchTime   string    `protobuf:"bytes,15,opt,name=fetchTime,proto3" json:"fetchTime,omitempty"`
	IntroPics   []string  `protobuf:"bytes,16,rep,name=introPics,proto3" json:"introPics,omitempty"`
	Magnets     []*Magnet `protobuf:"bytes,17,rep,name=magnets,proto3" json:"magnets,omitempty"`
}

func (x *Profile) Reset() {
	*x = Profile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpcsupport_grpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Profile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Profile) ProtoMessage() {}

func (x *Profile) ProtoReflect() protoreflect.Message {
	mi := &file_grpcsupport_grpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Profile.ProtoReflect.Descriptor instead.
func (*Profile) Descriptor() ([]byte, []int) {
	return file_grpcsupport_grpc_proto_rawDescGZIP(), []int{2}
}

func (x *Profile) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Profile) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Profile) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Profile) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Profile) GetCover() string {
	if x != nil {
		return x.Cover
	}
	return ""
}

func (x *Profile) GetReleaseData() string {
	if x != nil {
		return x.ReleaseData
	}
	return ""
}

func (x *Profile) GetDuration() string {
	if x != nil {
		return x.Duration
	}
	return ""
}

func (x *Profile) GetDirector() string {
	if x != nil {
		return x.Director
	}
	return ""
}

func (x *Profile) GetCompany() *Tag {
	if x != nil {
		return x.Company
	}
	return nil
}

func (x *Profile) GetSeries() *Tag {
	if x != nil {
		return x.Series
	}
	return nil
}

func (x *Profile) GetStars() string {
	if x != nil {
		return x.Stars
	}
	return ""
}

func (x *Profile) GetStarsPerson() string {
	if x != nil {
		return x.StarsPerson
	}
	return ""
}

func (x *Profile) GetType() []*Tag {
	if x != nil {
		return x.Type
	}
	return nil
}

func (x *Profile) GetActor() []*Tag {
	if x != nil {
		return x.Actor
	}
	return nil
}

func (x *Profile) GetFetchTime() string {
	if x != nil {
		return x.FetchTime
	}
	return ""
}

func (x *Profile) GetIntroPics() []string {
	if x != nil {
		return x.IntroPics
	}
	return nil
}

func (x *Profile) GetMagnets() []*Magnet {
	if x != nil {
		return x.Magnets
	}
	return nil
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpcsupport_grpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_grpcsupport_grpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_grpcsupport_grpc_proto_rawDescGZIP(), []int{3}
}

func (x *Response) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_grpcsupport_grpc_proto protoreflect.FileDescriptor

var file_grpcsupport_grpc_proto_rawDesc = []byte{
	0x0a, 0x16, 0x67, 0x72, 0x70, 0x63, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x6a, 0x61, 0x76, 0x43, 0x72, 0x61,
	0x77, 0x6c, 0x65, 0x72, 0x22, 0x31, 0x0a, 0x03, 0x54, 0x61, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x48, 0x0a, 0x06, 0x4d, 0x61, 0x67, 0x6e, 0x65,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x12, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x22, 0x87, 0x04, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x12,
	0x20, 0x0a, 0x0b, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a,
	0x08, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x29, 0x0a, 0x07, 0x63, 0x6f, 0x6d,
	0x70, 0x61, 0x6e, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6a, 0x61, 0x76,
	0x43, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x2e, 0x54, 0x61, 0x67, 0x52, 0x07, 0x63, 0x6f, 0x6d,
	0x70, 0x61, 0x6e, 0x79, 0x12, 0x27, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x69, 0x65, 0x73, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6a, 0x61, 0x76, 0x43, 0x72, 0x61, 0x77, 0x6c, 0x65,
	0x72, 0x2e, 0x54, 0x61, 0x67, 0x52, 0x06, 0x73, 0x65, 0x72, 0x69, 0x65, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x73, 0x74, 0x61, 0x72, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74,
	0x61, 0x72, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x72, 0x73, 0x50, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x74, 0x61, 0x72, 0x73, 0x50,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0d, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6a, 0x61, 0x76, 0x43, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72,
	0x2e, 0x54, 0x61, 0x67, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x25, 0x0a, 0x05, 0x61, 0x63,
	0x74, 0x6f, 0x72, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6a, 0x61, 0x76, 0x43,
	0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x2e, 0x54, 0x61, 0x67, 0x52, 0x05, 0x61, 0x63, 0x74, 0x6f,
	0x72, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x65, 0x74, 0x63, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x65, 0x74, 0x63, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x69, 0x6e, 0x74, 0x72, 0x6f, 0x50, 0x69, 0x63, 0x73, 0x18, 0x10, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x09, 0x69, 0x6e, 0x74, 0x72, 0x6f, 0x50, 0x69, 0x63, 0x73, 0x12, 0x2c, 0x0a,
	0x07, 0x6d, 0x61, 0x67, 0x6e, 0x65, 0x74, 0x73, 0x18, 0x11, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x6a, 0x61, 0x76, 0x43, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x2e, 0x4d, 0x61, 0x67, 0x6e,
	0x65, 0x74, 0x52, 0x07, 0x6d, 0x61, 0x67, 0x6e, 0x65, 0x74, 0x73, 0x22, 0x24, 0x0a, 0x08, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x32, 0x42, 0x0a, 0x0b, 0x53, 0x61, 0x76, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x33, 0x0a, 0x04, 0x53, 0x61, 0x76, 0x65, 0x12, 0x13, 0x2e, 0x6a, 0x61, 0x76, 0x43, 0x72,
	0x61, 0x77, 0x6c, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x1a, 0x14, 0x2e,
	0x6a, 0x61, 0x76, 0x43, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x18, 0x5a, 0x16, 0x6a, 0x61, 0x76, 0x43, 0x72, 0x61, 0x77,
	0x6c, 0x65, 0x72, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpcsupport_grpc_proto_rawDescOnce sync.Once
	file_grpcsupport_grpc_proto_rawDescData = file_grpcsupport_grpc_proto_rawDesc
)

func file_grpcsupport_grpc_proto_rawDescGZIP() []byte {
	file_grpcsupport_grpc_proto_rawDescOnce.Do(func() {
		file_grpcsupport_grpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpcsupport_grpc_proto_rawDescData)
	})
	return file_grpcsupport_grpc_proto_rawDescData
}

var file_grpcsupport_grpc_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_grpcsupport_grpc_proto_goTypes = []interface{}{
	(*Tag)(nil),      // 0: javCrawler.Tag
	(*Magnet)(nil),   // 1: javCrawler.Magnet
	(*Profile)(nil),  // 2: javCrawler.Profile
	(*Response)(nil), // 3: javCrawler.Response
}
var file_grpcsupport_grpc_proto_depIdxs = []int32{
	0, // 0: javCrawler.Profile.company:type_name -> javCrawler.Tag
	0, // 1: javCrawler.Profile.series:type_name -> javCrawler.Tag
	0, // 2: javCrawler.Profile.type:type_name -> javCrawler.Tag
	0, // 3: javCrawler.Profile.actor:type_name -> javCrawler.Tag
	1, // 4: javCrawler.Profile.magnets:type_name -> javCrawler.Magnet
	2, // 5: javCrawler.SaveService.Save:input_type -> javCrawler.Profile
	3, // 6: javCrawler.SaveService.Save:output_type -> javCrawler.Response
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_grpcsupport_grpc_proto_init() }
func file_grpcsupport_grpc_proto_init() {
	if File_grpcsupport_grpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpcsupport_grpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tag); i {
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
		file_grpcsupport_grpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Magnet); i {
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
		file_grpcsupport_grpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Profile); i {
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
		file_grpcsupport_grpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_grpcsupport_grpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpcsupport_grpc_proto_goTypes,
		DependencyIndexes: file_grpcsupport_grpc_proto_depIdxs,
		MessageInfos:      file_grpcsupport_grpc_proto_msgTypes,
	}.Build()
	File_grpcsupport_grpc_proto = out.File
	file_grpcsupport_grpc_proto_rawDesc = nil
	file_grpcsupport_grpc_proto_goTypes = nil
	file_grpcsupport_grpc_proto_depIdxs = nil
}

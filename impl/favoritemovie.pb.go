// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: favoritemovie.proto

package go_favoritemovie

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

type MovieSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Searchword string `protobuf:"bytes,1,opt,name=searchword,proto3" json:"searchword,omitempty"`
	Pagination int32  `protobuf:"varint,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *MovieSpec) Reset() {
	*x = MovieSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_favoritemovie_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MovieSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MovieSpec) ProtoMessage() {}

func (x *MovieSpec) ProtoReflect() protoreflect.Message {
	mi := &file_favoritemovie_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MovieSpec.ProtoReflect.Descriptor instead.
func (*MovieSpec) Descriptor() ([]byte, []int) {
	return file_favoritemovie_proto_rawDescGZIP(), []int{0}
}

func (x *MovieSpec) GetSearchword() string {
	if x != nil {
		return x.Searchword
	}
	return ""
}

func (x *MovieSpec) GetPagination() int32 {
	if x != nil {
		return x.Pagination
	}
	return 0
}

type MovieDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title  string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Year   string `protobuf:"bytes,2,opt,name=year,proto3" json:"year,omitempty"`
	ImdbId string `protobuf:"bytes,3,opt,name=imdbId,proto3" json:"imdbId,omitempty"`
	Type   string `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	Poster string `protobuf:"bytes,5,opt,name=poster,proto3" json:"poster,omitempty"`
}

func (x *MovieDetail) Reset() {
	*x = MovieDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_favoritemovie_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MovieDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MovieDetail) ProtoMessage() {}

func (x *MovieDetail) ProtoReflect() protoreflect.Message {
	mi := &file_favoritemovie_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MovieDetail.ProtoReflect.Descriptor instead.
func (*MovieDetail) Descriptor() ([]byte, []int) {
	return file_favoritemovie_proto_rawDescGZIP(), []int{1}
}

func (x *MovieDetail) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *MovieDetail) GetYear() string {
	if x != nil {
		return x.Year
	}
	return ""
}

func (x *MovieDetail) GetImdbId() string {
	if x != nil {
		return x.ImdbId
	}
	return ""
}

func (x *MovieDetail) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *MovieDetail) GetPoster() string {
	if x != nil {
		return x.Poster
	}
	return ""
}

type GetMovieSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdMovie string `protobuf:"bytes,1,opt,name=idMovie,proto3" json:"idMovie,omitempty"`
}

func (x *GetMovieSpec) Reset() {
	*x = GetMovieSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_favoritemovie_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMovieSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMovieSpec) ProtoMessage() {}

func (x *GetMovieSpec) ProtoReflect() protoreflect.Message {
	mi := &file_favoritemovie_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMovieSpec.ProtoReflect.Descriptor instead.
func (*GetMovieSpec) Descriptor() ([]byte, []int) {
	return file_favoritemovie_proto_rawDescGZIP(), []int{2}
}

func (x *GetMovieSpec) GetIdMovie() string {
	if x != nil {
		return x.IdMovie
	}
	return ""
}

var File_favoritemovie_proto protoreflect.FileDescriptor

var file_favoritemovie_proto_rawDesc = []byte{
	0x0a, 0x13, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x6d,
	0x6f, 0x76, 0x69, 0x65, 0x22, 0x4b, 0x0a, 0x09, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x53, 0x70, 0x65,
	0x63, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x77, 0x6f, 0x72,
	0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x7b, 0x0a, 0x0b, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x79, 0x65, 0x61, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x6d,
	0x64, 0x62, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x6d, 0x64, 0x62,
	0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x65, 0x72,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x22, 0x28,
	0x0a, 0x0c, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x53, 0x70, 0x65, 0x63, 0x12, 0x18,
	0x0a, 0x07, 0x69, 0x64, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x69, 0x64, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x32, 0xa7, 0x01, 0x0a, 0x15, 0x46, 0x61, 0x76,
	0x6f, 0x72, 0x69, 0x74, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x65, 0x72, 0x12, 0x47, 0x0a, 0x0b, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4d, 0x6f, 0x76, 0x69,
	0x65, 0x12, 0x18, 0x2e, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x6d, 0x6f, 0x76, 0x69,
	0x65, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x53, 0x70, 0x65, 0x63, 0x1a, 0x1a, 0x2e, 0x66, 0x61,
	0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2e, 0x4d, 0x6f, 0x76, 0x69,
	0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x22, 0x00, 0x30, 0x01, 0x12, 0x45, 0x0a, 0x08, 0x47,
	0x65, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x12, 0x1b, 0x2e, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69,
	0x74, 0x65, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65,
	0x53, 0x70, 0x65, 0x63, 0x1a, 0x1a, 0x2e, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x6d,
	0x6f, 0x76, 0x69, 0x65, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x22, 0x00, 0x42, 0x2c, 0x5a, 0x2a, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67,
	0x6f, 0x2d, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x3b,
	0x67, 0x6f, 0x5f, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x6d, 0x6f, 0x76, 0x69, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_favoritemovie_proto_rawDescOnce sync.Once
	file_favoritemovie_proto_rawDescData = file_favoritemovie_proto_rawDesc
)

func file_favoritemovie_proto_rawDescGZIP() []byte {
	file_favoritemovie_proto_rawDescOnce.Do(func() {
		file_favoritemovie_proto_rawDescData = protoimpl.X.CompressGZIP(file_favoritemovie_proto_rawDescData)
	})
	return file_favoritemovie_proto_rawDescData
}

var file_favoritemovie_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_favoritemovie_proto_goTypes = []interface{}{
	(*MovieSpec)(nil),    // 0: favoritemovie.MovieSpec
	(*MovieDetail)(nil),  // 1: favoritemovie.MovieDetail
	(*GetMovieSpec)(nil), // 2: favoritemovie.GetMovieSpec
}
var file_favoritemovie_proto_depIdxs = []int32{
	0, // 0: favoritemovie.FavoriteMovieSearcher.SearchMovie:input_type -> favoritemovie.MovieSpec
	2, // 1: favoritemovie.FavoriteMovieSearcher.GetMovie:input_type -> favoritemovie.GetMovieSpec
	1, // 2: favoritemovie.FavoriteMovieSearcher.SearchMovie:output_type -> favoritemovie.MovieDetail
	1, // 3: favoritemovie.FavoriteMovieSearcher.GetMovie:output_type -> favoritemovie.MovieDetail
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_favoritemovie_proto_init() }
func file_favoritemovie_proto_init() {
	if File_favoritemovie_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_favoritemovie_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MovieSpec); i {
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
		file_favoritemovie_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MovieDetail); i {
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
		file_favoritemovie_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMovieSpec); i {
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
			RawDescriptor: file_favoritemovie_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_favoritemovie_proto_goTypes,
		DependencyIndexes: file_favoritemovie_proto_depIdxs,
		MessageInfos:      file_favoritemovie_proto_msgTypes,
	}.Build()
	File_favoritemovie_proto = out.File
	file_favoritemovie_proto_rawDesc = nil
	file_favoritemovie_proto_goTypes = nil
	file_favoritemovie_proto_depIdxs = nil
}

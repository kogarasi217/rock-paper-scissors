// protoのバージョンです。

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: proto/rock-paper-scissors.proto

// メッセージ型などの名前の衝突を避けるためにパッケージ名を指定します。

package pb

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

// enumでグー、チョキ、パーを定義。
type HandShapes int32

const (
	HandShapes_HAND_SHAPES_UNKNOWN HandShapes = 0
	HandShapes_ROCK                HandShapes = 1
	HandShapes_PAPER               HandShapes = 2
	HandShapes_SCISSORS            HandShapes = 3
)

// Enum value maps for HandShapes.
var (
	HandShapes_name = map[int32]string{
		0: "HAND_SHAPES_UNKNOWN",
		1: "ROCK",
		2: "PAPER",
		3: "SCISSORS",
	}
	HandShapes_value = map[string]int32{
		"HAND_SHAPES_UNKNOWN": 0,
		"ROCK":                1,
		"PAPER":               2,
		"SCISSORS":            3,
	}
)

func (x HandShapes) Enum() *HandShapes {
	p := new(HandShapes)
	*p = x
	return p
}

func (x HandShapes) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HandShapes) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_rock_paper_scissors_proto_enumTypes[0].Descriptor()
}

func (HandShapes) Type() protoreflect.EnumType {
	return &file_proto_rock_paper_scissors_proto_enumTypes[0]
}

func (x HandShapes) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HandShapes.Descriptor instead.
func (HandShapes) EnumDescriptor() ([]byte, []int) {
	return file_proto_rock_paper_scissors_proto_rawDescGZIP(), []int{0}
}

// enumで勝敗とあいこを定義
type Result int32

const (
	Result_RESULT_UNKNOWN Result = 0
	Result_WIN            Result = 1
	Result_LOSE           Result = 2
	Result_DRAW           Result = 3
)

// Enum value maps for Result.
var (
	Result_name = map[int32]string{
		0: "RESULT_UNKNOWN",
		1: "WIN",
		2: "LOSE",
		3: "DRAW",
	}
	Result_value = map[string]int32{
		"RESULT_UNKNOWN": 0,
		"WIN":            1,
		"LOSE":           2,
		"DRAW":           3,
	}
)

func (x Result) Enum() *Result {
	p := new(Result)
	*p = x
	return p
}

func (x Result) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Result) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_rock_paper_scissors_proto_enumTypes[1].Descriptor()
}

func (Result) Type() protoreflect.EnumType {
	return &file_proto_rock_paper_scissors_proto_enumTypes[1]
}

func (x Result) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Result.Descriptor instead.
func (Result) EnumDescriptor() ([]byte, []int) {
	return file_proto_rock_paper_scissors_proto_rawDescGZIP(), []int{1}
}

// 対戦結果のメッセージ型です。
type MatchResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	YourHandShapes     HandShapes             `protobuf:"varint,1,opt,name=yourHandShapes,proto3,enum=game.HandShapes" json:"yourHandShapes,omitempty"`
	OpponentHandShapes HandShapes             `protobuf:"varint,2,opt,name=opponentHandShapes,proto3,enum=game.HandShapes" json:"opponentHandShapes,omitempty"`
	Result             Result                 `protobuf:"varint,3,opt,name=result,proto3,enum=game.Result" json:"result,omitempty"`
	CreateTime         *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
}

func (x *MatchResult) Reset() {
	*x = MatchResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rock_paper_scissors_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchResult) ProtoMessage() {}

func (x *MatchResult) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rock_paper_scissors_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchResult.ProtoReflect.Descriptor instead.
func (*MatchResult) Descriptor() ([]byte, []int) {
	return file_proto_rock_paper_scissors_proto_rawDescGZIP(), []int{0}
}

func (x *MatchResult) GetYourHandShapes() HandShapes {
	if x != nil {
		return x.YourHandShapes
	}
	return HandShapes_HAND_SHAPES_UNKNOWN
}

func (x *MatchResult) GetOpponentHandShapes() HandShapes {
	if x != nil {
		return x.OpponentHandShapes
	}
	return HandShapes_HAND_SHAPES_UNKNOWN
}

func (x *MatchResult) GetResult() Result {
	if x != nil {
		return x.Result
	}
	return Result_RESULT_UNKNOWN
}

func (x *MatchResult) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

// 今までの試合数、勝敗と対戦結果の履歴を持つメッセージ型です。
type Report struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NumberOfGames int32 `protobuf:"varint,1,opt,name=numberOfGames,proto3" json:"numberOfGames,omitempty"`
	NumberOfWins  int32 `protobuf:"varint,2,opt,name=numberOfWins,proto3" json:"numberOfWins,omitempty"`
	// `repeated`を付けることで配列を表現できます。
	MatchResults []*MatchResult `protobuf:"bytes,3,rep,name=matchResults,proto3" json:"matchResults,omitempty"`
}

func (x *Report) Reset() {
	*x = Report{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rock_paper_scissors_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Report) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Report) ProtoMessage() {}

func (x *Report) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rock_paper_scissors_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Report.ProtoReflect.Descriptor instead.
func (*Report) Descriptor() ([]byte, []int) {
	return file_proto_rock_paper_scissors_proto_rawDescGZIP(), []int{1}
}

func (x *Report) GetNumberOfGames() int32 {
	if x != nil {
		return x.NumberOfGames
	}
	return 0
}

func (x *Report) GetNumberOfWins() int32 {
	if x != nil {
		return x.NumberOfWins
	}
	return 0
}

func (x *Report) GetMatchResults() []*MatchResult {
	if x != nil {
		return x.MatchResults
	}
	return nil
}

// PlayGameメソッドのリクエスト用のメッセージ型
type PlayRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HandShapes HandShapes `protobuf:"varint,1,opt,name=handShapes,proto3,enum=game.HandShapes" json:"handShapes,omitempty"`
}

func (x *PlayRequest) Reset() {
	*x = PlayRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rock_paper_scissors_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayRequest) ProtoMessage() {}

func (x *PlayRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rock_paper_scissors_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayRequest.ProtoReflect.Descriptor instead.
func (*PlayRequest) Descriptor() ([]byte, []int) {
	return file_proto_rock_paper_scissors_proto_rawDescGZIP(), []int{2}
}

func (x *PlayRequest) GetHandShapes() HandShapes {
	if x != nil {
		return x.HandShapes
	}
	return HandShapes_HAND_SHAPES_UNKNOWN
}

// PlayGameメソッドのレスポンス用のメッセージ型
type PlayResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MatchResult *MatchResult `protobuf:"bytes,1,opt,name=matchResult,proto3" json:"matchResult,omitempty"`
}

func (x *PlayResponse) Reset() {
	*x = PlayResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rock_paper_scissors_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayResponse) ProtoMessage() {}

func (x *PlayResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rock_paper_scissors_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayResponse.ProtoReflect.Descriptor instead.
func (*PlayResponse) Descriptor() ([]byte, []int) {
	return file_proto_rock_paper_scissors_proto_rawDescGZIP(), []int{3}
}

func (x *PlayResponse) GetMatchResult() *MatchResult {
	if x != nil {
		return x.MatchResult
	}
	return nil
}

// ReportMatchResultsメソッドのリクエスト用のメッセージ型
type ReportRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ReportRequest) Reset() {
	*x = ReportRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rock_paper_scissors_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportRequest) ProtoMessage() {}

func (x *ReportRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rock_paper_scissors_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportRequest.ProtoReflect.Descriptor instead.
func (*ReportRequest) Descriptor() ([]byte, []int) {
	return file_proto_rock_paper_scissors_proto_rawDescGZIP(), []int{4}
}

// ReportMatchResultsメソッドのレスポンス用のメッセージ型
type ReportResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Report *Report `protobuf:"bytes,1,opt,name=report,proto3" json:"report,omitempty"`
}

func (x *ReportResponse) Reset() {
	*x = ReportResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rock_paper_scissors_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportResponse) ProtoMessage() {}

func (x *ReportResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rock_paper_scissors_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportResponse.ProtoReflect.Descriptor instead.
func (*ReportResponse) Descriptor() ([]byte, []int) {
	return file_proto_rock_paper_scissors_proto_rawDescGZIP(), []int{5}
}

func (x *ReportResponse) GetReport() *Report {
	if x != nil {
		return x.Report
	}
	return nil
}

var File_proto_rock_paper_scissors_proto protoreflect.FileDescriptor

var file_proto_rock_paper_scissors_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x6f, 0x63, 0x6b, 0x2d, 0x70, 0x61, 0x70,
	0x65, 0x72, 0x2d, 0x73, 0x63, 0x69, 0x73, 0x73, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x67, 0x61, 0x6d, 0x65, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xec, 0x01, 0x0a, 0x0b, 0x4d, 0x61, 0x74,
	0x63, 0x68, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x38, 0x0a, 0x0e, 0x79, 0x6f, 0x75, 0x72,
	0x48, 0x61, 0x6e, 0x64, 0x53, 0x68, 0x61, 0x70, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x10, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x48, 0x61, 0x6e, 0x64, 0x53, 0x68, 0x61, 0x70,
	0x65, 0x73, 0x52, 0x0e, 0x79, 0x6f, 0x75, 0x72, 0x48, 0x61, 0x6e, 0x64, 0x53, 0x68, 0x61, 0x70,
	0x65, 0x73, 0x12, 0x40, 0x0a, 0x12, 0x6f, 0x70, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x48, 0x61,
	0x6e, 0x64, 0x53, 0x68, 0x61, 0x70, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10,
	0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x48, 0x61, 0x6e, 0x64, 0x53, 0x68, 0x61, 0x70, 0x65, 0x73,
	0x52, 0x12, 0x6f, 0x70, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x48, 0x61, 0x6e, 0x64, 0x53, 0x68,
	0x61, 0x70, 0x65, 0x73, 0x12, 0x24, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x89, 0x01, 0x0a, 0x06, 0x52, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x47, 0x61,
	0x6d, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x4f, 0x66, 0x47, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x4f, 0x66, 0x57, 0x69, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x57, 0x69, 0x6e, 0x73, 0x12, 0x35, 0x0a, 0x0c,
	0x6d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x0c, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x73, 0x22, 0x3f, 0x0a, 0x0b, 0x50, 0x6c, 0x61, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x30, 0x0a, 0x0a, 0x68, 0x61, 0x6e, 0x64, 0x53, 0x68, 0x61, 0x70, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x48, 0x61,
	0x6e, 0x64, 0x53, 0x68, 0x61, 0x70, 0x65, 0x73, 0x52, 0x0a, 0x68, 0x61, 0x6e, 0x64, 0x53, 0x68,
	0x61, 0x70, 0x65, 0x73, 0x22, 0x43, 0x0a, 0x0c, 0x50, 0x6c, 0x61, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x0b, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x67, 0x61, 0x6d, 0x65,
	0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x0b, 0x6d, 0x61,
	0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x0f, 0x0a, 0x0d, 0x52, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x36, 0x0a, 0x0e, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x06,
	0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x67,
	0x61, 0x6d, 0x65, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x06, 0x72, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x2a, 0x48, 0x0a, 0x0a, 0x48, 0x61, 0x6e, 0x64, 0x53, 0x68, 0x61, 0x70, 0x65, 0x73,
	0x12, 0x17, 0x0a, 0x13, 0x48, 0x41, 0x4e, 0x44, 0x5f, 0x53, 0x48, 0x41, 0x50, 0x45, 0x53, 0x5f,
	0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x52, 0x4f, 0x43,
	0x4b, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x50, 0x41, 0x50, 0x45, 0x52, 0x10, 0x02, 0x12, 0x0c,
	0x0a, 0x08, 0x53, 0x43, 0x49, 0x53, 0x53, 0x4f, 0x52, 0x53, 0x10, 0x03, 0x2a, 0x39, 0x0a, 0x06,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x12, 0x0a, 0x0e, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54,
	0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x57, 0x49,
	0x4e, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x4c, 0x4f, 0x53, 0x45, 0x10, 0x02, 0x12, 0x08, 0x0a,
	0x04, 0x44, 0x52, 0x41, 0x57, 0x10, 0x03, 0x32, 0x92, 0x01, 0x0a, 0x18, 0x52, 0x6f, 0x63, 0x6b,
	0x50, 0x61, 0x70, 0x65, 0x72, 0x53, 0x63, 0x69, 0x73, 0x73, 0x6f, 0x72, 0x73, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x33, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x79, 0x47, 0x61, 0x6d, 0x65,
	0x12, 0x11, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x12, 0x52, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12,
	0x13, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x52, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x05, 0x5a, 0x03,
	0x70, 0x62, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_rock_paper_scissors_proto_rawDescOnce sync.Once
	file_proto_rock_paper_scissors_proto_rawDescData = file_proto_rock_paper_scissors_proto_rawDesc
)

func file_proto_rock_paper_scissors_proto_rawDescGZIP() []byte {
	file_proto_rock_paper_scissors_proto_rawDescOnce.Do(func() {
		file_proto_rock_paper_scissors_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_rock_paper_scissors_proto_rawDescData)
	})
	return file_proto_rock_paper_scissors_proto_rawDescData
}

var file_proto_rock_paper_scissors_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_proto_rock_paper_scissors_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_rock_paper_scissors_proto_goTypes = []interface{}{
	(HandShapes)(0),               // 0: game.HandShapes
	(Result)(0),                   // 1: game.Result
	(*MatchResult)(nil),           // 2: game.MatchResult
	(*Report)(nil),                // 3: game.Report
	(*PlayRequest)(nil),           // 4: game.PlayRequest
	(*PlayResponse)(nil),          // 5: game.PlayResponse
	(*ReportRequest)(nil),         // 6: game.ReportRequest
	(*ReportResponse)(nil),        // 7: game.ReportResponse
	(*timestamppb.Timestamp)(nil), // 8: google.protobuf.Timestamp
}
var file_proto_rock_paper_scissors_proto_depIdxs = []int32{
	0,  // 0: game.MatchResult.yourHandShapes:type_name -> game.HandShapes
	0,  // 1: game.MatchResult.opponentHandShapes:type_name -> game.HandShapes
	1,  // 2: game.MatchResult.result:type_name -> game.Result
	8,  // 3: game.MatchResult.create_time:type_name -> google.protobuf.Timestamp
	2,  // 4: game.Report.matchResults:type_name -> game.MatchResult
	0,  // 5: game.PlayRequest.handShapes:type_name -> game.HandShapes
	2,  // 6: game.PlayResponse.matchResult:type_name -> game.MatchResult
	3,  // 7: game.ReportResponse.report:type_name -> game.Report
	4,  // 8: game.RockPaperScissorsService.PlayGame:input_type -> game.PlayRequest
	6,  // 9: game.RockPaperScissorsService.ReportMatchResults:input_type -> game.ReportRequest
	5,  // 10: game.RockPaperScissorsService.PlayGame:output_type -> game.PlayResponse
	7,  // 11: game.RockPaperScissorsService.ReportMatchResults:output_type -> game.ReportResponse
	10, // [10:12] is the sub-list for method output_type
	8,  // [8:10] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_proto_rock_paper_scissors_proto_init() }
func file_proto_rock_paper_scissors_proto_init() {
	if File_proto_rock_paper_scissors_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_rock_paper_scissors_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatchResult); i {
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
		file_proto_rock_paper_scissors_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Report); i {
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
		file_proto_rock_paper_scissors_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayRequest); i {
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
		file_proto_rock_paper_scissors_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayResponse); i {
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
		file_proto_rock_paper_scissors_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportRequest); i {
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
		file_proto_rock_paper_scissors_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportResponse); i {
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
			RawDescriptor: file_proto_rock_paper_scissors_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_rock_paper_scissors_proto_goTypes,
		DependencyIndexes: file_proto_rock_paper_scissors_proto_depIdxs,
		EnumInfos:         file_proto_rock_paper_scissors_proto_enumTypes,
		MessageInfos:      file_proto_rock_paper_scissors_proto_msgTypes,
	}.Build()
	File_proto_rock_paper_scissors_proto = out.File
	file_proto_rock_paper_scissors_proto_rawDesc = nil
	file_proto_rock_paper_scissors_proto_goTypes = nil
	file_proto_rock_paper_scissors_proto_depIdxs = nil
}

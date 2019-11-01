// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tinkoff/cloud/tts/v1/tts.proto

package v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type SsmlVoiceGender int32

const (
	SsmlVoiceGender_SSML_VOICE_GENDER_UNSPECIFIED SsmlVoiceGender = 0
	SsmlVoiceGender_MALE                          SsmlVoiceGender = 1
	SsmlVoiceGender_FEMALE                        SsmlVoiceGender = 2
	SsmlVoiceGender_NEUTRAL                       SsmlVoiceGender = 3
)

var SsmlVoiceGender_name = map[int32]string{
	0: "SSML_VOICE_GENDER_UNSPECIFIED",
	1: "MALE",
	2: "FEMALE",
	3: "NEUTRAL",
}

var SsmlVoiceGender_value = map[string]int32{
	"SSML_VOICE_GENDER_UNSPECIFIED": 0,
	"MALE":                          1,
	"FEMALE":                        2,
	"NEUTRAL":                       3,
}

func (x SsmlVoiceGender) String() string {
	return proto.EnumName(SsmlVoiceGender_name, int32(x))
}

func (SsmlVoiceGender) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_dce28e7d7749d71f, []int{0}
}

type AudioEncoding int32

const (
	AudioEncoding_ENCODING_UNSPECIFIED AudioEncoding = 0
	AudioEncoding_LINEAR16             AudioEncoding = 1
	AudioEncoding_MULAW                AudioEncoding = 3
	AudioEncoding_ALAW                 AudioEncoding = 8
	AudioEncoding_LINEAR32F            AudioEncoding = 9
	AudioEncoding_RAW_OPUS             AudioEncoding = 11
)

var AudioEncoding_name = map[int32]string{
	0:  "ENCODING_UNSPECIFIED",
	1:  "LINEAR16",
	3:  "MULAW",
	8:  "ALAW",
	9:  "LINEAR32F",
	11: "RAW_OPUS",
}

var AudioEncoding_value = map[string]int32{
	"ENCODING_UNSPECIFIED": 0,
	"LINEAR16":             1,
	"MULAW":                3,
	"ALAW":                 8,
	"LINEAR32F":            9,
	"RAW_OPUS":             11,
}

func (x AudioEncoding) String() string {
	return proto.EnumName(AudioEncoding_name, int32(x))
}

func (AudioEncoding) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_dce28e7d7749d71f, []int{1}
}

type Voice struct {
	LanguageCodes          []string        `protobuf:"bytes,1,rep,name=language_codes,json=languageCodes,proto3" json:"language_codes,omitempty"`
	Name                   string          `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	SsmlGender             SsmlVoiceGender `protobuf:"varint,3,opt,name=ssml_gender,json=ssmlGender,proto3,enum=tinkoff.cloud.tts.v1.SsmlVoiceGender" json:"ssml_gender,omitempty"`
	NaturalSampleRateHertz int32           `protobuf:"varint,4,opt,name=natural_sample_rate_hertz,json=naturalSampleRateHertz,proto3" json:"natural_sample_rate_hertz,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}        `json:"-"`
	XXX_unrecognized       []byte          `json:"-"`
	XXX_sizecache          int32           `json:"-"`
}

func (m *Voice) Reset()         { *m = Voice{} }
func (m *Voice) String() string { return proto.CompactTextString(m) }
func (*Voice) ProtoMessage()    {}
func (*Voice) Descriptor() ([]byte, []int) {
	return fileDescriptor_dce28e7d7749d71f, []int{0}
}

func (m *Voice) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Voice.Unmarshal(m, b)
}
func (m *Voice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Voice.Marshal(b, m, deterministic)
}
func (m *Voice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Voice.Merge(m, src)
}
func (m *Voice) XXX_Size() int {
	return xxx_messageInfo_Voice.Size(m)
}
func (m *Voice) XXX_DiscardUnknown() {
	xxx_messageInfo_Voice.DiscardUnknown(m)
}

var xxx_messageInfo_Voice proto.InternalMessageInfo

func (m *Voice) GetLanguageCodes() []string {
	if m != nil {
		return m.LanguageCodes
	}
	return nil
}

func (m *Voice) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Voice) GetSsmlGender() SsmlVoiceGender {
	if m != nil {
		return m.SsmlGender
	}
	return SsmlVoiceGender_SSML_VOICE_GENDER_UNSPECIFIED
}

func (m *Voice) GetNaturalSampleRateHertz() int32 {
	if m != nil {
		return m.NaturalSampleRateHertz
	}
	return 0
}

type ListVoicesRequest struct {
	LanguageCode         string   `protobuf:"bytes,1,opt,name=language_code,json=languageCode,proto3" json:"language_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListVoicesRequest) Reset()         { *m = ListVoicesRequest{} }
func (m *ListVoicesRequest) String() string { return proto.CompactTextString(m) }
func (*ListVoicesRequest) ProtoMessage()    {}
func (*ListVoicesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dce28e7d7749d71f, []int{1}
}

func (m *ListVoicesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListVoicesRequest.Unmarshal(m, b)
}
func (m *ListVoicesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListVoicesRequest.Marshal(b, m, deterministic)
}
func (m *ListVoicesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListVoicesRequest.Merge(m, src)
}
func (m *ListVoicesRequest) XXX_Size() int {
	return xxx_messageInfo_ListVoicesRequest.Size(m)
}
func (m *ListVoicesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListVoicesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListVoicesRequest proto.InternalMessageInfo

func (m *ListVoicesRequest) GetLanguageCode() string {
	if m != nil {
		return m.LanguageCode
	}
	return ""
}

type ListVoicesResponses struct {
	Voices               []*Voice `protobuf:"bytes,1,rep,name=voices,proto3" json:"voices,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListVoicesResponses) Reset()         { *m = ListVoicesResponses{} }
func (m *ListVoicesResponses) String() string { return proto.CompactTextString(m) }
func (*ListVoicesResponses) ProtoMessage()    {}
func (*ListVoicesResponses) Descriptor() ([]byte, []int) {
	return fileDescriptor_dce28e7d7749d71f, []int{2}
}

func (m *ListVoicesResponses) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListVoicesResponses.Unmarshal(m, b)
}
func (m *ListVoicesResponses) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListVoicesResponses.Marshal(b, m, deterministic)
}
func (m *ListVoicesResponses) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListVoicesResponses.Merge(m, src)
}
func (m *ListVoicesResponses) XXX_Size() int {
	return xxx_messageInfo_ListVoicesResponses.Size(m)
}
func (m *ListVoicesResponses) XXX_DiscardUnknown() {
	xxx_messageInfo_ListVoicesResponses.DiscardUnknown(m)
}

var xxx_messageInfo_ListVoicesResponses proto.InternalMessageInfo

func (m *ListVoicesResponses) GetVoices() []*Voice {
	if m != nil {
		return m.Voices
	}
	return nil
}

type SynthesisInput struct {
	Text                 string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SynthesisInput) Reset()         { *m = SynthesisInput{} }
func (m *SynthesisInput) String() string { return proto.CompactTextString(m) }
func (*SynthesisInput) ProtoMessage()    {}
func (*SynthesisInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_dce28e7d7749d71f, []int{3}
}

func (m *SynthesisInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SynthesisInput.Unmarshal(m, b)
}
func (m *SynthesisInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SynthesisInput.Marshal(b, m, deterministic)
}
func (m *SynthesisInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SynthesisInput.Merge(m, src)
}
func (m *SynthesisInput) XXX_Size() int {
	return xxx_messageInfo_SynthesisInput.Size(m)
}
func (m *SynthesisInput) XXX_DiscardUnknown() {
	xxx_messageInfo_SynthesisInput.DiscardUnknown(m)
}

var xxx_messageInfo_SynthesisInput proto.InternalMessageInfo

func (m *SynthesisInput) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type VoiceSelectionParams struct {
	LanguageCode         string          `protobuf:"bytes,1,opt,name=language_code,json=languageCode,proto3" json:"language_code,omitempty"`
	Name                 string          `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	SsmlGender           SsmlVoiceGender `protobuf:"varint,3,opt,name=ssml_gender,json=ssmlGender,proto3,enum=tinkoff.cloud.tts.v1.SsmlVoiceGender" json:"ssml_gender,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *VoiceSelectionParams) Reset()         { *m = VoiceSelectionParams{} }
func (m *VoiceSelectionParams) String() string { return proto.CompactTextString(m) }
func (*VoiceSelectionParams) ProtoMessage()    {}
func (*VoiceSelectionParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_dce28e7d7749d71f, []int{4}
}

func (m *VoiceSelectionParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VoiceSelectionParams.Unmarshal(m, b)
}
func (m *VoiceSelectionParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VoiceSelectionParams.Marshal(b, m, deterministic)
}
func (m *VoiceSelectionParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VoiceSelectionParams.Merge(m, src)
}
func (m *VoiceSelectionParams) XXX_Size() int {
	return xxx_messageInfo_VoiceSelectionParams.Size(m)
}
func (m *VoiceSelectionParams) XXX_DiscardUnknown() {
	xxx_messageInfo_VoiceSelectionParams.DiscardUnknown(m)
}

var xxx_messageInfo_VoiceSelectionParams proto.InternalMessageInfo

func (m *VoiceSelectionParams) GetLanguageCode() string {
	if m != nil {
		return m.LanguageCode
	}
	return ""
}

func (m *VoiceSelectionParams) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *VoiceSelectionParams) GetSsmlGender() SsmlVoiceGender {
	if m != nil {
		return m.SsmlGender
	}
	return SsmlVoiceGender_SSML_VOICE_GENDER_UNSPECIFIED
}

type AudioConfig struct {
	AudioEncoding        AudioEncoding `protobuf:"varint,1,opt,name=audio_encoding,json=audioEncoding,proto3,enum=tinkoff.cloud.tts.v1.AudioEncoding" json:"audio_encoding,omitempty"`
	SpeakingRate         float64       `protobuf:"fixed64,2,opt,name=speaking_rate,json=speakingRate,proto3" json:"speaking_rate,omitempty"`
	SampleRateHertz      int32         `protobuf:"varint,5,opt,name=sample_rate_hertz,json=sampleRateHertz,proto3" json:"sample_rate_hertz,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *AudioConfig) Reset()         { *m = AudioConfig{} }
func (m *AudioConfig) String() string { return proto.CompactTextString(m) }
func (*AudioConfig) ProtoMessage()    {}
func (*AudioConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_dce28e7d7749d71f, []int{5}
}

func (m *AudioConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AudioConfig.Unmarshal(m, b)
}
func (m *AudioConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AudioConfig.Marshal(b, m, deterministic)
}
func (m *AudioConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AudioConfig.Merge(m, src)
}
func (m *AudioConfig) XXX_Size() int {
	return xxx_messageInfo_AudioConfig.Size(m)
}
func (m *AudioConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_AudioConfig.DiscardUnknown(m)
}

var xxx_messageInfo_AudioConfig proto.InternalMessageInfo

func (m *AudioConfig) GetAudioEncoding() AudioEncoding {
	if m != nil {
		return m.AudioEncoding
	}
	return AudioEncoding_ENCODING_UNSPECIFIED
}

func (m *AudioConfig) GetSpeakingRate() float64 {
	if m != nil {
		return m.SpeakingRate
	}
	return 0
}

func (m *AudioConfig) GetSampleRateHertz() int32 {
	if m != nil {
		return m.SampleRateHertz
	}
	return 0
}

type SynthesizeSpeechRequest struct {
	Input                *SynthesisInput       `protobuf:"bytes,1,opt,name=input,proto3" json:"input,omitempty"`
	Voice                *VoiceSelectionParams `protobuf:"bytes,2,opt,name=voice,proto3" json:"voice,omitempty"`
	AudioConfig          *AudioConfig          `protobuf:"bytes,3,opt,name=audio_config,json=audioConfig,proto3" json:"audio_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *SynthesizeSpeechRequest) Reset()         { *m = SynthesizeSpeechRequest{} }
func (m *SynthesizeSpeechRequest) String() string { return proto.CompactTextString(m) }
func (*SynthesizeSpeechRequest) ProtoMessage()    {}
func (*SynthesizeSpeechRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dce28e7d7749d71f, []int{6}
}

func (m *SynthesizeSpeechRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SynthesizeSpeechRequest.Unmarshal(m, b)
}
func (m *SynthesizeSpeechRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SynthesizeSpeechRequest.Marshal(b, m, deterministic)
}
func (m *SynthesizeSpeechRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SynthesizeSpeechRequest.Merge(m, src)
}
func (m *SynthesizeSpeechRequest) XXX_Size() int {
	return xxx_messageInfo_SynthesizeSpeechRequest.Size(m)
}
func (m *SynthesizeSpeechRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SynthesizeSpeechRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SynthesizeSpeechRequest proto.InternalMessageInfo

func (m *SynthesizeSpeechRequest) GetInput() *SynthesisInput {
	if m != nil {
		return m.Input
	}
	return nil
}

func (m *SynthesizeSpeechRequest) GetVoice() *VoiceSelectionParams {
	if m != nil {
		return m.Voice
	}
	return nil
}

func (m *SynthesizeSpeechRequest) GetAudioConfig() *AudioConfig {
	if m != nil {
		return m.AudioConfig
	}
	return nil
}

type SynthesizeSpeechResponse struct {
	AudioContent         []byte   `protobuf:"bytes,1,opt,name=audio_content,json=audioContent,proto3" json:"audio_content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SynthesizeSpeechResponse) Reset()         { *m = SynthesizeSpeechResponse{} }
func (m *SynthesizeSpeechResponse) String() string { return proto.CompactTextString(m) }
func (*SynthesizeSpeechResponse) ProtoMessage()    {}
func (*SynthesizeSpeechResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dce28e7d7749d71f, []int{7}
}

func (m *SynthesizeSpeechResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SynthesizeSpeechResponse.Unmarshal(m, b)
}
func (m *SynthesizeSpeechResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SynthesizeSpeechResponse.Marshal(b, m, deterministic)
}
func (m *SynthesizeSpeechResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SynthesizeSpeechResponse.Merge(m, src)
}
func (m *SynthesizeSpeechResponse) XXX_Size() int {
	return xxx_messageInfo_SynthesizeSpeechResponse.Size(m)
}
func (m *SynthesizeSpeechResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SynthesizeSpeechResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SynthesizeSpeechResponse proto.InternalMessageInfo

func (m *SynthesizeSpeechResponse) GetAudioContent() []byte {
	if m != nil {
		return m.AudioContent
	}
	return nil
}

type StreamingSynthesizeSpeechResponse struct {
	AudioChunk           []byte   `protobuf:"bytes,1,opt,name=audio_chunk,json=audioChunk,proto3" json:"audio_chunk,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamingSynthesizeSpeechResponse) Reset()         { *m = StreamingSynthesizeSpeechResponse{} }
func (m *StreamingSynthesizeSpeechResponse) String() string { return proto.CompactTextString(m) }
func (*StreamingSynthesizeSpeechResponse) ProtoMessage()    {}
func (*StreamingSynthesizeSpeechResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dce28e7d7749d71f, []int{8}
}

func (m *StreamingSynthesizeSpeechResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamingSynthesizeSpeechResponse.Unmarshal(m, b)
}
func (m *StreamingSynthesizeSpeechResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamingSynthesizeSpeechResponse.Marshal(b, m, deterministic)
}
func (m *StreamingSynthesizeSpeechResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamingSynthesizeSpeechResponse.Merge(m, src)
}
func (m *StreamingSynthesizeSpeechResponse) XXX_Size() int {
	return xxx_messageInfo_StreamingSynthesizeSpeechResponse.Size(m)
}
func (m *StreamingSynthesizeSpeechResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamingSynthesizeSpeechResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StreamingSynthesizeSpeechResponse proto.InternalMessageInfo

func (m *StreamingSynthesizeSpeechResponse) GetAudioChunk() []byte {
	if m != nil {
		return m.AudioChunk
	}
	return nil
}

func init() {
	proto.RegisterEnum("tinkoff.cloud.tts.v1.SsmlVoiceGender", SsmlVoiceGender_name, SsmlVoiceGender_value)
	proto.RegisterEnum("tinkoff.cloud.tts.v1.AudioEncoding", AudioEncoding_name, AudioEncoding_value)
	proto.RegisterType((*Voice)(nil), "tinkoff.cloud.tts.v1.Voice")
	proto.RegisterType((*ListVoicesRequest)(nil), "tinkoff.cloud.tts.v1.ListVoicesRequest")
	proto.RegisterType((*ListVoicesResponses)(nil), "tinkoff.cloud.tts.v1.ListVoicesResponses")
	proto.RegisterType((*SynthesisInput)(nil), "tinkoff.cloud.tts.v1.SynthesisInput")
	proto.RegisterType((*VoiceSelectionParams)(nil), "tinkoff.cloud.tts.v1.VoiceSelectionParams")
	proto.RegisterType((*AudioConfig)(nil), "tinkoff.cloud.tts.v1.AudioConfig")
	proto.RegisterType((*SynthesizeSpeechRequest)(nil), "tinkoff.cloud.tts.v1.SynthesizeSpeechRequest")
	proto.RegisterType((*SynthesizeSpeechResponse)(nil), "tinkoff.cloud.tts.v1.SynthesizeSpeechResponse")
	proto.RegisterType((*StreamingSynthesizeSpeechResponse)(nil), "tinkoff.cloud.tts.v1.StreamingSynthesizeSpeechResponse")
}

func init() { proto.RegisterFile("tinkoff/cloud/tts/v1/tts.proto", fileDescriptor_dce28e7d7749d71f) }

var fileDescriptor_dce28e7d7749d71f = []byte{
	// 973 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0xcd, 0x6f, 0xe3, 0x44,
	0x14, 0xc7, 0x8d, 0x9d, 0x75, 0x9f, 0xd3, 0xac, 0x77, 0x5a, 0x76, 0x43, 0x97, 0x85, 0xd6, 0xcb,
	0x8a, 0x62, 0x69, 0xe3, 0x6d, 0x2a, 0xf1, 0xd1, 0x0b, 0xb8, 0x89, 0x9b, 0xa6, 0xca, 0x47, 0x19,
	0xa7, 0x2d, 0x70, 0xb1, 0xdc, 0x64, 0xea, 0x58, 0x71, 0xec, 0x90, 0x99, 0x54, 0xdd, 0x95, 0x10,
	0x12, 0x47, 0xae, 0x1c, 0x38, 0xf0, 0xf7, 0x70, 0xe4, 0xc4, 0x19, 0x89, 0x03, 0x7f, 0x08, 0x9a,
	0xb1, 0xb3, 0xfd, 0x4a, 0xbb, 0xcb, 0x81, 0x93, 0xc7, 0x6f, 0xde, 0x7b, 0xf3, 0x7e, 0xef, 0xf7,
	0x9b, 0x37, 0xf0, 0x01, 0x0b, 0xe3, 0x61, 0x72, 0x7a, 0x6a, 0xf5, 0xa2, 0x64, 0xda, 0xb7, 0x18,
	0xa3, 0xd6, 0xd9, 0x26, 0xff, 0x94, 0xc7, 0x93, 0x84, 0x25, 0x68, 0x25, 0xdb, 0x2f, 0x8b, 0xfd,
	0x32, 0xdf, 0x38, 0xdb, 0x5c, 0x7d, 0x3f, 0x48, 0x92, 0x20, 0x22, 0x96, 0x3f, 0x0e, 0x2d, 0x3f,
	0x8e, 0x13, 0xe6, 0xb3, 0x30, 0x89, 0xb3, 0x18, 0xe3, 0x77, 0x09, 0x94, 0xa3, 0x24, 0xec, 0x11,
	0xf4, 0x0c, 0x8a, 0x91, 0x1f, 0x07, 0x53, 0x3f, 0x20, 0x5e, 0x2f, 0xe9, 0x13, 0x5a, 0x92, 0xd6,
	0x72, 0x1b, 0x8b, 0x78, 0x69, 0x66, 0xad, 0x72, 0x23, 0x42, 0x20, 0xc7, 0xfe, 0x88, 0x94, 0x16,
	0xd6, 0xa4, 0x8d, 0x45, 0x2c, 0xd6, 0x68, 0x17, 0x34, 0x4a, 0x47, 0x91, 0x17, 0x90, 0xb8, 0x4f,
	0x26, 0xa5, 0xdc, 0x9a, 0xb4, 0x51, 0xac, 0x3c, 0x2b, 0xcf, 0x2b, 0xa7, 0xec, 0xd2, 0x51, 0x24,
	0x0e, 0xac, 0x0b, 0x67, 0x0c, 0x3c, 0x32, 0x5d, 0xa3, 0x2f, 0xe0, 0xbd, 0xd8, 0x67, 0xd3, 0x89,
	0x1f, 0x79, 0xd4, 0x1f, 0x8d, 0x23, 0xe2, 0x4d, 0x7c, 0x46, 0xbc, 0x01, 0x99, 0xb0, 0x57, 0x25,
	0x79, 0x4d, 0xda, 0x50, 0xf0, 0xc3, 0xcc, 0xc1, 0x15, 0xfb, 0xd8, 0x67, 0x64, 0x8f, 0xef, 0x1a,
	0x9f, 0xc3, 0x83, 0x66, 0x48, 0x99, 0xc8, 0x4c, 0x31, 0xf9, 0x7e, 0x4a, 0x28, 0x43, 0x4f, 0x61,
	0xe9, 0x0a, 0xa4, 0x92, 0x24, 0x8a, 0x2e, 0x5c, 0x46, 0x64, 0xec, 0xc3, 0xf2, 0xe5, 0x48, 0x3a,
	0x4e, 0x62, 0x4a, 0x28, 0xda, 0x82, 0xfc, 0x99, 0x30, 0x89, 0x36, 0x68, 0x95, 0xc7, 0xf3, 0xe1,
	0x88, 0x30, 0x9c, 0xb9, 0x1a, 0x2f, 0xa0, 0xe8, 0xbe, 0x8c, 0xd9, 0x80, 0xd0, 0x90, 0x36, 0xe2,
	0xf1, 0x94, 0xf1, 0x76, 0x31, 0x72, 0xce, 0xb2, 0x93, 0xc5, 0x7a, 0x5f, 0x56, 0x17, 0xf4, 0x1c,
	0x96, 0x39, 0x70, 0xe3, 0x57, 0x09, 0x56, 0x44, 0x0e, 0x97, 0x44, 0xa4, 0xc7, 0x99, 0x39, 0xf0,
	0x27, 0xfe, 0x88, 0xbe, 0x55, 0xed, 0xff, 0x27, 0x19, 0xc6, 0x1f, 0x12, 0x68, 0xf6, 0xb4, 0x1f,
	0x26, 0xd5, 0x24, 0x3e, 0x0d, 0x03, 0xb4, 0x0f, 0x45, 0x9f, 0xff, 0x7a, 0x24, 0xee, 0x25, 0xfd,
	0x30, 0x0e, 0x44, 0x45, 0xc5, 0xca, 0xd3, 0xf9, 0xa9, 0x45, 0xa8, 0x93, 0xb9, 0xe2, 0x25, 0xff,
	0xf2, 0x2f, 0x07, 0x47, 0xc7, 0xc4, 0x1f, 0x86, 0x71, 0x20, 0x28, 0x16, 0x00, 0x24, 0x5c, 0x98,
	0x19, 0x39, 0xaf, 0xc8, 0x84, 0x07, 0x37, 0x55, 0xa0, 0x08, 0x15, 0xdc, 0xa7, 0x57, 0xe9, 0xdf,
	0x97, 0xd5, 0x9c, 0x2e, 0xef, 0xcb, 0xaa, 0xac, 0x2b, 0x58, 0x19, 0x87, 0xac, 0x37, 0xc0, 0xc5,
	0xb3, 0x24, 0x9a, 0x8e, 0x88, 0x17, 0xf8, 0x61, 0xec, 0xf5, 0x4f, 0x8c, 0xbf, 0x24, 0x78, 0x34,
	0x23, 0xe7, 0x15, 0x71, 0xc7, 0x84, 0xf4, 0x06, 0x33, 0xa1, 0x6c, 0x83, 0x12, 0x72, 0xba, 0x04,
	0x24, 0xad, 0xf2, 0xd1, 0x2d, 0xdd, 0xba, 0x42, 0x2d, 0x4e, 0x43, 0xd0, 0x57, 0xa0, 0x08, 0xf6,
	0x05, 0x06, 0xad, 0x62, 0xde, 0xa1, 0x93, 0x6b, 0x1c, 0xe3, 0x34, 0x10, 0xd5, 0xa0, 0x90, 0x76,
	0xb6, 0x27, 0x3a, 0x2d, 0x28, 0xd3, 0x2a, 0xeb, 0x77, 0xf4, 0x35, 0xa5, 0x04, 0x6b, 0xfe, 0xc5,
	0x8f, 0xf1, 0x25, 0x94, 0x6e, 0xc2, 0x4b, 0xd5, 0xcc, 0xfb, 0xfd, 0xfa, 0x04, 0x46, 0xe2, 0x14,
	0x67, 0x01, 0x17, 0x66, 0xf1, 0xdc, 0x66, 0xd4, 0x60, 0xdd, 0x65, 0x13, 0xe2, 0x8f, 0xc2, 0x38,
	0xb8, 0x35, 0xd3, 0x87, 0xa0, 0x65, 0x99, 0x06, 0xd3, 0x78, 0x98, 0xe5, 0x81, 0x34, 0x0f, 0xb7,
	0x98, 0xc7, 0x70, 0xff, 0x9a, 0xaa, 0xd0, 0x3a, 0x3c, 0x71, 0xdd, 0x56, 0xd3, 0x3b, 0xea, 0x34,
	0xaa, 0x8e, 0x57, 0x77, 0xda, 0x35, 0x07, 0x7b, 0x87, 0x6d, 0xf7, 0xc0, 0xa9, 0x36, 0x76, 0x1b,
	0x4e, 0x4d, 0x7f, 0x07, 0xa9, 0x20, 0xb7, 0xec, 0xa6, 0xa3, 0x4b, 0x08, 0x20, 0xbf, 0xeb, 0x88,
	0xf5, 0x02, 0xd2, 0xe0, 0x5e, 0xdb, 0x39, 0xec, 0x62, 0xbb, 0xa9, 0xe7, 0xcc, 0xbf, 0x25, 0x58,
	0xba, 0x22, 0x2a, 0x54, 0x82, 0x15, 0xa7, 0x5d, 0xed, 0xd4, 0x1a, 0xed, 0xfa, 0xb5, 0x74, 0x05,
	0x50, 0x9b, 0x8d, 0xb6, 0x63, 0xe3, 0xcd, 0x4f, 0x75, 0x09, 0x2d, 0x82, 0xd2, 0x3a, 0x6c, 0xda,
	0xc7, 0x7a, 0x8e, 0x9f, 0x63, 0xf3, 0x95, 0x8a, 0x96, 0x60, 0x31, 0x75, 0xd9, 0xaa, 0xec, 0xea,
	0x8b, 0x3c, 0x02, 0xdb, 0xc7, 0x5e, 0xe7, 0xe0, 0xd0, 0xd5, 0x35, 0x83, 0xdf, 0xd0, 0x05, 0x83,
	0xcb, 0x49, 0x36, 0x64, 0x55, 0xd1, 0x15, 0x43, 0x56, 0xf3, 0x7a, 0xde, 0x90, 0xd5, 0x7b, 0xfa,
	0x3d, 0x43, 0x56, 0x41, 0x07, 0x43, 0x56, 0x0b, 0x7a, 0xc1, 0x94, 0x77, 0x9b, 0x76, 0xd5, 0xcc,
	0xd9, 0x2d, 0x6c, 0xe6, 0xed, 0x16, 0xf6, 0x8e, 0x77, 0x4c, 0xb5, 0x53, 0xaf, 0x8b, 0x44, 0xe6,
	0x43, 0xf7, 0xc0, 0x71, 0xbe, 0xf1, 0x8e, 0x1b, 0xdd, 0x3d, 0x6f, 0xcf, 0xb1, 0x39, 0xec, 0x9d,
	0x6f, 0xbb, 0x8e, 0x09, 0xdc, 0xe3, 0xa8, 0x83, 0x77, 0x1a, 0xae, 0x09, 0xad, 0x03, 0xa7, 0xee,
	0xd9, 0x87, 0xb5, 0x46, 0xa7, 0xf2, 0x5b, 0x0e, 0x0a, 0x5d, 0x72, 0xce, 0xba, 0x49, 0xda, 0x75,
	0xf4, 0x03, 0xc0, 0xc5, 0x6c, 0x42, 0x1f, 0xcf, 0x57, 0xc4, 0x8d, 0xb9, 0xb7, 0xfa, 0xc9, 0x9b,
	0x1d, 0xb3, 0x31, 0x67, 0x3c, 0xfe, 0xe9, 0xcf, 0x7f, 0x7e, 0x59, 0x78, 0x17, 0x2d, 0x67, 0x4f,
	0xc9, 0x76, 0x14, 0x52, 0xe6, 0xa5, 0xe3, 0x0c, 0xfd, 0x2c, 0x01, 0x5c, 0x28, 0x01, 0x3d, 0xbf,
	0xfb, 0x5a, 0x5c, 0xbb, 0x54, 0xab, 0xe5, 0xb7, 0x75, 0x4f, 0x6b, 0x31, 0x9e, 0x88, 0x52, 0x1e,
	0x19, 0x68, 0x56, 0x0a, 0x7d, 0xed, 0xb9, 0x2d, 0x99, 0xe8, 0x47, 0x58, 0x9e, 0x23, 0xcf, 0xff,
	0x5a, 0xd4, 0x67, 0xb7, 0xb8, 0xbf, 0x49, 0xf8, 0x2f, 0xa4, 0x1d, 0xf7, 0xbb, 0xaf, 0x83, 0x90,
	0x0d, 0xa6, 0x27, 0xe5, 0x5e, 0x32, 0xb2, 0xba, 0x69, 0x9a, 0xea, 0x84, 0xf4, 0x43, 0xe6, 0xbe,
	0xa4, 0x8c, 0x8c, 0xa8, 0x25, 0xfa, 0x36, 0x0c, 0xd9, 0x73, 0x72, 0x2e, 0x06, 0x14, 0xb5, 0x82,
	0x84, 0x0f, 0x6c, 0x6b, 0x3c, 0x0c, 0xac, 0x79, 0xaf, 0xf7, 0x49, 0x5e, 0x3c, 0xc3, 0x5b, 0xff,
	0x06, 0x00, 0x00, 0xff, 0xff, 0x1b, 0x43, 0x67, 0xd6, 0xdc, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TextToSpeechClient is the client API for TextToSpeech service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TextToSpeechClient interface {
	ListVoices(ctx context.Context, in *ListVoicesRequest, opts ...grpc.CallOption) (*ListVoicesResponses, error)
	Synthesize(ctx context.Context, in *SynthesizeSpeechRequest, opts ...grpc.CallOption) (*SynthesizeSpeechResponse, error)
	StreamingSynthesize(ctx context.Context, in *SynthesizeSpeechRequest, opts ...grpc.CallOption) (TextToSpeech_StreamingSynthesizeClient, error)
}

type textToSpeechClient struct {
	cc *grpc.ClientConn
}

func NewTextToSpeechClient(cc *grpc.ClientConn) TextToSpeechClient {
	return &textToSpeechClient{cc}
}

func (c *textToSpeechClient) ListVoices(ctx context.Context, in *ListVoicesRequest, opts ...grpc.CallOption) (*ListVoicesResponses, error) {
	out := new(ListVoicesResponses)
	err := c.cc.Invoke(ctx, "/tinkoff.cloud.tts.v1.TextToSpeech/ListVoices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *textToSpeechClient) Synthesize(ctx context.Context, in *SynthesizeSpeechRequest, opts ...grpc.CallOption) (*SynthesizeSpeechResponse, error) {
	out := new(SynthesizeSpeechResponse)
	err := c.cc.Invoke(ctx, "/tinkoff.cloud.tts.v1.TextToSpeech/Synthesize", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *textToSpeechClient) StreamingSynthesize(ctx context.Context, in *SynthesizeSpeechRequest, opts ...grpc.CallOption) (TextToSpeech_StreamingSynthesizeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TextToSpeech_serviceDesc.Streams[0], "/tinkoff.cloud.tts.v1.TextToSpeech/StreamingSynthesize", opts...)
	if err != nil {
		return nil, err
	}
	x := &textToSpeechStreamingSynthesizeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TextToSpeech_StreamingSynthesizeClient interface {
	Recv() (*StreamingSynthesizeSpeechResponse, error)
	grpc.ClientStream
}

type textToSpeechStreamingSynthesizeClient struct {
	grpc.ClientStream
}

func (x *textToSpeechStreamingSynthesizeClient) Recv() (*StreamingSynthesizeSpeechResponse, error) {
	m := new(StreamingSynthesizeSpeechResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TextToSpeechServer is the server API for TextToSpeech service.
type TextToSpeechServer interface {
	ListVoices(context.Context, *ListVoicesRequest) (*ListVoicesResponses, error)
	Synthesize(context.Context, *SynthesizeSpeechRequest) (*SynthesizeSpeechResponse, error)
	StreamingSynthesize(*SynthesizeSpeechRequest, TextToSpeech_StreamingSynthesizeServer) error
}

func RegisterTextToSpeechServer(s *grpc.Server, srv TextToSpeechServer) {
	s.RegisterService(&_TextToSpeech_serviceDesc, srv)
}

func _TextToSpeech_ListVoices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListVoicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TextToSpeechServer).ListVoices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tinkoff.cloud.tts.v1.TextToSpeech/ListVoices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TextToSpeechServer).ListVoices(ctx, req.(*ListVoicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TextToSpeech_Synthesize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SynthesizeSpeechRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TextToSpeechServer).Synthesize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tinkoff.cloud.tts.v1.TextToSpeech/Synthesize",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TextToSpeechServer).Synthesize(ctx, req.(*SynthesizeSpeechRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TextToSpeech_StreamingSynthesize_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SynthesizeSpeechRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TextToSpeechServer).StreamingSynthesize(m, &textToSpeechStreamingSynthesizeServer{stream})
}

type TextToSpeech_StreamingSynthesizeServer interface {
	Send(*StreamingSynthesizeSpeechResponse) error
	grpc.ServerStream
}

type textToSpeechStreamingSynthesizeServer struct {
	grpc.ServerStream
}

func (x *textToSpeechStreamingSynthesizeServer) Send(m *StreamingSynthesizeSpeechResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _TextToSpeech_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tinkoff.cloud.tts.v1.TextToSpeech",
	HandlerType: (*TextToSpeechServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListVoices",
			Handler:    _TextToSpeech_ListVoices_Handler,
		},
		{
			MethodName: "Synthesize",
			Handler:    _TextToSpeech_Synthesize_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamingSynthesize",
			Handler:       _TextToSpeech_StreamingSynthesize_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "tinkoff/cloud/tts/v1/tts.proto",
}

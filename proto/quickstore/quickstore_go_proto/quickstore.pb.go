// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/quickstore/quickstore.proto

package mako_quickstore

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	threshold_analyzer_go_proto "github.com/google/mako/proto/clients/analyzers/threshold_analyzer_go_proto"
	utest_analyzer_go_proto "github.com/google/mako/proto/clients/analyzers/utest_analyzer_go_proto"
	window_deviation_go_proto "github.com/google/mako/proto/clients/analyzers/window_deviation_go_proto"
	rolling_window_reducer_go_proto "github.com/google/mako/proto/helpers/rolling_window_reducer/rolling_window_reducer_go_proto"
	mako_go_proto "github.com/google/mako/spec/proto/mako_go_proto"
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

type QuickstoreOutput_Status int32

const (
	QuickstoreOutput_SUCCESS       QuickstoreOutput_Status = 1
	QuickstoreOutput_ERROR         QuickstoreOutput_Status = 2
	QuickstoreOutput_ANALYSIS_FAIL QuickstoreOutput_Status = 4
)

var QuickstoreOutput_Status_name = map[int32]string{
	1: "SUCCESS",
	2: "ERROR",
	4: "ANALYSIS_FAIL",
}

var QuickstoreOutput_Status_value = map[string]int32{
	"SUCCESS":       1,
	"ERROR":         2,
	"ANALYSIS_FAIL": 4,
}

func (x QuickstoreOutput_Status) Enum() *QuickstoreOutput_Status {
	p := new(QuickstoreOutput_Status)
	*p = x
	return p
}

func (x QuickstoreOutput_Status) String() string {
	return proto.EnumName(QuickstoreOutput_Status_name, int32(x))
}

func (x *QuickstoreOutput_Status) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(QuickstoreOutput_Status_value, data, "QuickstoreOutput_Status")
	if err != nil {
		return err
	}
	*x = QuickstoreOutput_Status(value)
	return nil
}

func (QuickstoreOutput_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1383215e6ee55b1f, []int{1, 0}
}

type QuickstoreInput struct {
	BenchmarkKey          *string                                               `protobuf:"bytes,1,opt,name=benchmark_key,json=benchmarkKey" json:"benchmark_key,omitempty"`
	TimestampMs           *float64                                              `protobuf:"fixed64,2,opt,name=timestamp_ms,json=timestampMs" json:"timestamp_ms,omitempty"`
	BuildId               *int64                                                `protobuf:"varint,20,opt,name=build_id,json=buildId" json:"build_id,omitempty"`
	DurationTimeMs        *float64                                              `protobuf:"fixed64,3,opt,name=duration_time_ms,json=durationTimeMs" json:"duration_time_ms,omitempty"`
	Tags                  []string                                              `protobuf:"bytes,4,rep,name=tags" json:"tags,omitempty"`
	HistoricalContextTags []string                                              `protobuf:"bytes,24,rep,name=historical_context_tags,json=historicalContextTags" json:"historical_context_tags,omitempty"`
	HoverText             *string                                               `protobuf:"bytes,14,opt,name=hover_text,json=hoverText" json:"hover_text,omitempty"`
	Description           *string                                               `protobuf:"bytes,6,opt,name=description" json:"description,omitempty"`
	TestPassId            *string                                               `protobuf:"bytes,18,opt,name=test_pass_id,json=testPassId" json:"test_pass_id,omitempty"`
	AnnotationList        []*mako_go_proto.RunAnnotation                        `protobuf:"bytes,7,rep,name=annotation_list,json=annotationList" json:"annotation_list,omitempty"`
	HyperlinkList         []*mako_go_proto.NamedData                            `protobuf:"bytes,8,rep,name=hyperlink_list,json=hyperlinkList" json:"hyperlink_list,omitempty"`
	AuxData               []*mako_go_proto.NamedData                            `protobuf:"bytes,17,rep,name=aux_data,json=auxData" json:"aux_data,omitempty"`
	IgnoreRangeList       []*mako_go_proto.LabeledRange                         `protobuf:"bytes,9,rep,name=ignore_range_list,json=ignoreRangeList" json:"ignore_range_list,omitempty"`
	TempDir               *string                                               `protobuf:"bytes,21,opt,name=temp_dir,json=tempDir" json:"temp_dir,omitempty"`
	DeleteSampleFiles     *bool                                                 `protobuf:"varint,22,opt,name=delete_sample_files,json=deleteSampleFiles,def=1" json:"delete_sample_files,omitempty"`
	AnalysisPass          *QuickstoreInput_ConditionalFields                    `protobuf:"bytes,10,opt,name=analysis_pass,json=analysisPass" json:"analysis_pass,omitempty"`
	AnalysisFail          *QuickstoreInput_ConditionalFields                    `protobuf:"bytes,11,opt,name=analysis_fail,json=analysisFail" json:"analysis_fail,omitempty"`
	RwrConfigs            []*rolling_window_reducer_go_proto.RWRConfig          `protobuf:"bytes,23,rep,name=rwr_configs,json=rwrConfigs" json:"rwr_configs,omitempty"`
	ThresholdInputs       []*threshold_analyzer_go_proto.ThresholdAnalyzerInput `protobuf:"bytes,12,rep,name=threshold_inputs,json=thresholdInputs" json:"threshold_inputs,omitempty"`
	WdaInputs             []*window_deviation_go_proto.WindowDeviationInput     `protobuf:"bytes,13,rep,name=wda_inputs,json=wdaInputs" json:"wda_inputs,omitempty"`
	UtestInputs           []*utest_analyzer_go_proto.UTestAnalyzerInput         `protobuf:"bytes,16,rep,name=utest_inputs,json=utestInputs" json:"utest_inputs,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}                                              `json:"-"`
	XXX_unrecognized      []byte                                                `json:"-"`
	XXX_sizecache         int32                                                 `json:"-"`
}

func (m *QuickstoreInput) Reset()         { *m = QuickstoreInput{} }
func (m *QuickstoreInput) String() string { return proto.CompactTextString(m) }
func (*QuickstoreInput) ProtoMessage()    {}
func (*QuickstoreInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_1383215e6ee55b1f, []int{0}
}

func (m *QuickstoreInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QuickstoreInput.Unmarshal(m, b)
}
func (m *QuickstoreInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QuickstoreInput.Marshal(b, m, deterministic)
}
func (m *QuickstoreInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuickstoreInput.Merge(m, src)
}
func (m *QuickstoreInput) XXX_Size() int {
	return xxx_messageInfo_QuickstoreInput.Size(m)
}
func (m *QuickstoreInput) XXX_DiscardUnknown() {
	xxx_messageInfo_QuickstoreInput.DiscardUnknown(m)
}

var xxx_messageInfo_QuickstoreInput proto.InternalMessageInfo

const Default_QuickstoreInput_DeleteSampleFiles bool = true

func (m *QuickstoreInput) GetBenchmarkKey() string {
	if m != nil && m.BenchmarkKey != nil {
		return *m.BenchmarkKey
	}
	return ""
}

func (m *QuickstoreInput) GetTimestampMs() float64 {
	if m != nil && m.TimestampMs != nil {
		return *m.TimestampMs
	}
	return 0
}

func (m *QuickstoreInput) GetBuildId() int64 {
	if m != nil && m.BuildId != nil {
		return *m.BuildId
	}
	return 0
}

func (m *QuickstoreInput) GetDurationTimeMs() float64 {
	if m != nil && m.DurationTimeMs != nil {
		return *m.DurationTimeMs
	}
	return 0
}

func (m *QuickstoreInput) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *QuickstoreInput) GetHistoricalContextTags() []string {
	if m != nil {
		return m.HistoricalContextTags
	}
	return nil
}

func (m *QuickstoreInput) GetHoverText() string {
	if m != nil && m.HoverText != nil {
		return *m.HoverText
	}
	return ""
}

func (m *QuickstoreInput) GetDescription() string {
	if m != nil && m.Description != nil {
		return *m.Description
	}
	return ""
}

func (m *QuickstoreInput) GetTestPassId() string {
	if m != nil && m.TestPassId != nil {
		return *m.TestPassId
	}
	return ""
}

func (m *QuickstoreInput) GetAnnotationList() []*mako_go_proto.RunAnnotation {
	if m != nil {
		return m.AnnotationList
	}
	return nil
}

func (m *QuickstoreInput) GetHyperlinkList() []*mako_go_proto.NamedData {
	if m != nil {
		return m.HyperlinkList
	}
	return nil
}

func (m *QuickstoreInput) GetAuxData() []*mako_go_proto.NamedData {
	if m != nil {
		return m.AuxData
	}
	return nil
}

func (m *QuickstoreInput) GetIgnoreRangeList() []*mako_go_proto.LabeledRange {
	if m != nil {
		return m.IgnoreRangeList
	}
	return nil
}

func (m *QuickstoreInput) GetTempDir() string {
	if m != nil && m.TempDir != nil {
		return *m.TempDir
	}
	return ""
}

func (m *QuickstoreInput) GetDeleteSampleFiles() bool {
	if m != nil && m.DeleteSampleFiles != nil {
		return *m.DeleteSampleFiles
	}
	return Default_QuickstoreInput_DeleteSampleFiles
}

func (m *QuickstoreInput) GetAnalysisPass() *QuickstoreInput_ConditionalFields {
	if m != nil {
		return m.AnalysisPass
	}
	return nil
}

func (m *QuickstoreInput) GetAnalysisFail() *QuickstoreInput_ConditionalFields {
	if m != nil {
		return m.AnalysisFail
	}
	return nil
}

func (m *QuickstoreInput) GetRwrConfigs() []*rolling_window_reducer_go_proto.RWRConfig {
	if m != nil {
		return m.RwrConfigs
	}
	return nil
}

func (m *QuickstoreInput) GetThresholdInputs() []*threshold_analyzer_go_proto.ThresholdAnalyzerInput {
	if m != nil {
		return m.ThresholdInputs
	}
	return nil
}

func (m *QuickstoreInput) GetWdaInputs() []*window_deviation_go_proto.WindowDeviationInput {
	if m != nil {
		return m.WdaInputs
	}
	return nil
}

func (m *QuickstoreInput) GetUtestInputs() []*utest_analyzer_go_proto.UTestAnalyzerInput {
	if m != nil {
		return m.UtestInputs
	}
	return nil
}

type QuickstoreInput_ConditionalFields struct {
	Tags                 []string `protobuf:"bytes,1,rep,name=tags" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QuickstoreInput_ConditionalFields) Reset()         { *m = QuickstoreInput_ConditionalFields{} }
func (m *QuickstoreInput_ConditionalFields) String() string { return proto.CompactTextString(m) }
func (*QuickstoreInput_ConditionalFields) ProtoMessage()    {}
func (*QuickstoreInput_ConditionalFields) Descriptor() ([]byte, []int) {
	return fileDescriptor_1383215e6ee55b1f, []int{0, 0}
}

func (m *QuickstoreInput_ConditionalFields) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QuickstoreInput_ConditionalFields.Unmarshal(m, b)
}
func (m *QuickstoreInput_ConditionalFields) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QuickstoreInput_ConditionalFields.Marshal(b, m, deterministic)
}
func (m *QuickstoreInput_ConditionalFields) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuickstoreInput_ConditionalFields.Merge(m, src)
}
func (m *QuickstoreInput_ConditionalFields) XXX_Size() int {
	return xxx_messageInfo_QuickstoreInput_ConditionalFields.Size(m)
}
func (m *QuickstoreInput_ConditionalFields) XXX_DiscardUnknown() {
	xxx_messageInfo_QuickstoreInput_ConditionalFields.DiscardUnknown(m)
}

var xxx_messageInfo_QuickstoreInput_ConditionalFields proto.InternalMessageInfo

func (m *QuickstoreInput_ConditionalFields) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

type QuickstoreOutput struct {
	Status               *QuickstoreOutput_Status        `protobuf:"varint,1,opt,name=status,enum=mako.quickstore.QuickstoreOutput_Status" json:"status,omitempty"`
	AnalyzerOutputList   []*mako_go_proto.AnalyzerOutput `protobuf:"bytes,2,rep,name=analyzer_output_list,json=analyzerOutputList" json:"analyzer_output_list,omitempty"`
	SummaryOutput        *string                         `protobuf:"bytes,3,opt,name=summary_output,json=summaryOutput" json:"summary_output,omitempty"`
	RunChartLink         *string                         `protobuf:"bytes,4,opt,name=run_chart_link,json=runChartLink" json:"run_chart_link,omitempty"`
	RunKey               *string                         `protobuf:"bytes,5,opt,name=run_key,json=runKey" json:"run_key,omitempty"`
	GeneratedSampleFiles []string                        `protobuf:"bytes,6,rep,name=generated_sample_files,json=generatedSampleFiles" json:"generated_sample_files,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *QuickstoreOutput) Reset()         { *m = QuickstoreOutput{} }
func (m *QuickstoreOutput) String() string { return proto.CompactTextString(m) }
func (*QuickstoreOutput) ProtoMessage()    {}
func (*QuickstoreOutput) Descriptor() ([]byte, []int) {
	return fileDescriptor_1383215e6ee55b1f, []int{1}
}

func (m *QuickstoreOutput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QuickstoreOutput.Unmarshal(m, b)
}
func (m *QuickstoreOutput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QuickstoreOutput.Marshal(b, m, deterministic)
}
func (m *QuickstoreOutput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuickstoreOutput.Merge(m, src)
}
func (m *QuickstoreOutput) XXX_Size() int {
	return xxx_messageInfo_QuickstoreOutput.Size(m)
}
func (m *QuickstoreOutput) XXX_DiscardUnknown() {
	xxx_messageInfo_QuickstoreOutput.DiscardUnknown(m)
}

var xxx_messageInfo_QuickstoreOutput proto.InternalMessageInfo

func (m *QuickstoreOutput) GetStatus() QuickstoreOutput_Status {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return QuickstoreOutput_SUCCESS
}

func (m *QuickstoreOutput) GetAnalyzerOutputList() []*mako_go_proto.AnalyzerOutput {
	if m != nil {
		return m.AnalyzerOutputList
	}
	return nil
}

func (m *QuickstoreOutput) GetSummaryOutput() string {
	if m != nil && m.SummaryOutput != nil {
		return *m.SummaryOutput
	}
	return ""
}

func (m *QuickstoreOutput) GetRunChartLink() string {
	if m != nil && m.RunChartLink != nil {
		return *m.RunChartLink
	}
	return ""
}

func (m *QuickstoreOutput) GetRunKey() string {
	if m != nil && m.RunKey != nil {
		return *m.RunKey
	}
	return ""
}

func (m *QuickstoreOutput) GetGeneratedSampleFiles() []string {
	if m != nil {
		return m.GeneratedSampleFiles
	}
	return nil
}

func init() {
	proto.RegisterEnum("mako.quickstore.QuickstoreOutput_Status", QuickstoreOutput_Status_name, QuickstoreOutput_Status_value)
	proto.RegisterType((*QuickstoreInput)(nil), "mako.quickstore.QuickstoreInput")
	proto.RegisterType((*QuickstoreInput_ConditionalFields)(nil), "mako.quickstore.QuickstoreInput.ConditionalFields")
	proto.RegisterType((*QuickstoreOutput)(nil), "mako.quickstore.QuickstoreOutput")
}

func init() { proto.RegisterFile("proto/quickstore/quickstore.proto", fileDescriptor_1383215e6ee55b1f) }

var fileDescriptor_1383215e6ee55b1f = []byte{
	// 948 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0x6d, 0x6f, 0xdb, 0x36,
	0x17, 0x85, 0x13, 0xc7, 0x2f, 0xf4, 0x9b, 0xc2, 0x24, 0x8d, 0x1a, 0xe0, 0x01, 0xdc, 0x3c, 0x1b,
	0x6a, 0xac, 0x9b, 0x32, 0x64, 0x45, 0xb1, 0x0d, 0xc3, 0x30, 0xcf, 0x69, 0x80, 0x64, 0x69, 0xbb,
	0xd1, 0x2e, 0x82, 0x7d, 0x12, 0x18, 0x91, 0xb1, 0x09, 0x4b, 0x94, 0x46, 0x52, 0x75, 0xbc, 0xff,
	0xb3, 0x9f, 0xb7, 0xff, 0x30, 0xf0, 0x52, 0x96, 0xe3, 0xb4, 0xd9, 0x97, 0x7d, 0x23, 0xcf, 0x3d,
	0xe7, 0x90, 0xf7, 0x92, 0xf7, 0xa2, 0x67, 0x99, 0x4a, 0x4d, 0x7a, 0xf2, 0x47, 0x2e, 0xa2, 0xb9,
	0x36, 0xa9, 0xe2, 0xf7, 0x96, 0x01, 0xc4, 0x70, 0x2f, 0xa1, 0xf3, 0x34, 0x58, 0xc3, 0x47, 0x5f,
	0x3b, 0x4d, 0x14, 0x0b, 0x2e, 0x8d, 0x3e, 0xa1, 0x92, 0xc6, 0xcb, 0x3f, 0xb9, 0xd2, 0x27, 0x66,
	0xa6, 0xb8, 0x9e, 0xa5, 0x31, 0x0b, 0x57, 0x98, 0xb3, 0x38, 0xfa, 0xf2, 0x31, 0x45, 0x6e, 0xb8,
	0x36, 0x0f, 0xd9, 0xc1, 0x63, 0xec, 0x85, 0x90, 0x2c, 0x5d, 0x84, 0x8c, 0x7f, 0x10, 0xd4, 0x88,
	0x54, 0x16, 0xfc, 0xa1, 0xe3, 0xcf, 0x78, 0x9c, 0x59, 0x96, 0x4a, 0xe3, 0x58, 0xc8, 0x69, 0x58,
	0xb0, 0x15, 0x67, 0x79, 0xc4, 0xd5, 0x23, 0x70, 0x61, 0x71, 0xa0, 0x33, 0x1e, 0x9d, 0x38, 0x1f,
	0x48, 0x17, 0x96, 0xc7, 0x7f, 0x35, 0x51, 0xef, 0xb7, 0x32, 0xf1, 0x0b, 0x99, 0xe5, 0x06, 0xff,
	0x1f, 0x75, 0x6e, 0xb8, 0x8c, 0x66, 0x09, 0x55, 0xf3, 0x70, 0xce, 0x97, 0x7e, 0xa5, 0x5f, 0x19,
	0x34, 0x49, 0xbb, 0x04, 0x7f, 0xe1, 0x4b, 0xfc, 0x0c, 0xb5, 0x8d, 0x48, 0xb8, 0x36, 0x34, 0xc9,
	0xc2, 0x44, 0xfb, 0x5b, 0xfd, 0xca, 0xa0, 0x42, 0x5a, 0x25, 0xf6, 0x46, 0xe3, 0xa7, 0xa8, 0x71,
	0x93, 0x8b, 0x98, 0x85, 0x82, 0xf9, 0xfb, 0xfd, 0xca, 0x60, 0x9b, 0xd4, 0x61, 0x7f, 0xc1, 0xf0,
	0x00, 0x79, 0x2c, 0x57, 0x90, 0x62, 0x68, 0x25, 0xd6, 0x61, 0x1b, 0x1c, 0xba, 0x2b, 0x7c, 0x22,
	0x12, 0xfe, 0x46, 0x63, 0x8c, 0xaa, 0x86, 0x4e, 0xb5, 0x5f, 0xed, 0x6f, 0x0f, 0x9a, 0x04, 0xd6,
	0xf8, 0x15, 0x3a, 0x9c, 0x09, 0x7b, 0x61, 0x11, 0xd1, 0x38, 0x8c, 0x52, 0x69, 0xf8, 0x9d, 0x09,
	0x81, 0xe6, 0x03, 0xed, 0x60, 0x1d, 0x1e, 0xb9, 0xe8, 0xc4, 0xea, 0xfe, 0x87, 0xd0, 0x2c, 0xfd,
	0xc0, 0x55, 0x68, 0x11, 0xbf, 0x0b, 0x59, 0x35, 0x01, 0x99, 0xf0, 0x3b, 0x83, 0xfb, 0xa8, 0xc5,
	0xb8, 0x8e, 0x94, 0xc8, 0xec, 0xf9, 0x7e, 0x0d, 0xe2, 0xf7, 0x21, 0xdc, 0x47, 0x6d, 0x78, 0xce,
	0x8c, 0x6a, 0x6d, 0xb3, 0xc2, 0x40, 0x41, 0x16, 0xfb, 0x95, 0x6a, 0x7d, 0xc1, 0xf0, 0x0f, 0xa8,
	0x47, 0xa5, 0x4c, 0x8d, 0x4b, 0x2d, 0x16, 0xda, 0xf8, 0xf5, 0xfe, 0xf6, 0xa0, 0x75, 0xba, 0x17,
	0x40, 0xd5, 0x49, 0x2e, 0x87, 0x65, 0x9c, 0x74, 0xd7, 0xdc, 0x2b, 0xa1, 0x0d, 0x7e, 0x85, 0xba,
	0xb3, 0x65, 0xc6, 0x55, 0x2c, 0xe4, 0xdc, 0x89, 0x1b, 0x20, 0xee, 0x39, 0xf1, 0x5b, 0x9a, 0x70,
	0x76, 0x46, 0x0d, 0x25, 0x9d, 0x92, 0x06, 0xba, 0x2f, 0x50, 0x83, 0xe6, 0x77, 0x21, 0xa3, 0x86,
	0xfa, 0xbb, 0x9f, 0x56, 0xd4, 0x69, 0x7e, 0x67, 0x17, 0xf8, 0x47, 0xb4, 0x2b, 0xa6, 0x32, 0x55,
	0x3c, 0x54, 0x54, 0x4e, 0xb9, 0x3b, 0xa6, 0x09, 0x22, 0xec, 0x44, 0x57, 0xf4, 0x86, 0xc7, 0x9c,
	0x11, 0x1b, 0x26, 0x3d, 0x47, 0x86, 0x0d, 0x9c, 0xf5, 0x14, 0x35, 0x0c, 0x4f, 0xb2, 0x90, 0x09,
	0xe5, 0x1f, 0x40, 0xfe, 0x75, 0xbb, 0x3f, 0x13, 0x0a, 0xbf, 0x44, 0x7b, 0x8c, 0xc7, 0xdc, 0xf0,
	0x50, 0xd3, 0x24, 0x8b, 0x79, 0x78, 0x2b, 0x62, 0xae, 0xfd, 0x27, 0xfd, 0xca, 0xa0, 0xf1, 0x7d,
	0xd5, 0xa8, 0x9c, 0x93, 0x5d, 0x47, 0x18, 0x43, 0xfc, 0xdc, 0x86, 0xf1, 0x35, 0xea, 0x40, 0x03,
	0x68, 0xa1, 0xa1, 0xb0, 0x3e, 0xea, 0x57, 0x06, 0xad, 0xd3, 0xd3, 0xe0, 0x41, 0x57, 0x06, 0x0f,
	0xfe, 0x69, 0x30, 0x4a, 0x25, 0x13, 0xb6, 0x76, 0x34, 0x3e, 0x17, 0x3c, 0x66, 0x9a, 0xb4, 0x57,
	0x46, 0xf6, 0x35, 0x36, 0x8c, 0x6f, 0xa9, 0x88, 0xfd, 0xd6, 0x7f, 0x37, 0x3e, 0xa7, 0x22, 0xc6,
	0xdf, 0xa2, 0x96, 0x5a, 0x28, 0xfb, 0xf1, 0x6e, 0xc5, 0x54, 0xfb, 0x87, 0x50, 0xbc, 0x43, 0x67,
	0x5b, 0xf4, 0x68, 0x40, 0xae, 0xc9, 0x08, 0xe2, 0x04, 0xa9, 0x85, 0x72, 0x4b, 0x8d, 0x19, 0xf2,
	0xd6, 0x23, 0x44, 0xd8, 0xc3, 0xb4, 0xdf, 0x06, 0xf9, 0x77, 0x4e, 0x5e, 0x8e, 0x82, 0xe0, 0x13,
	0xa3, 0x66, 0xb2, 0x82, 0x86, 0x05, 0x02, 0xd7, 0x25, 0xbd, 0x92, 0x0a, 0x7b, 0x8d, 0x2f, 0x11,
	0x5a, 0x30, 0xba, 0xf2, 0xef, 0x80, 0xff, 0x0b, 0xe7, 0xff, 0xd1, 0x80, 0xb9, 0x06, 0xe0, 0x6c,
	0xb5, 0x77, 0x8e, 0xcd, 0x05, 0xa3, 0xa5, 0x57, 0xdb, 0x8d, 0xb0, 0xc2, 0xcd, 0x03, 0xb7, 0xe7,
	0xce, 0xed, 0xc1, 0x70, 0x7b, 0x3f, 0xe1, 0xda, 0x6c, 0xde, 0xad, 0x05, 0x14, 0xe7, 0x75, 0xf4,
	0x1c, 0xed, 0x7e, 0x54, 0xda, 0xb2, 0xc1, 0x2b, 0xeb, 0x06, 0xbf, 0xac, 0x36, 0x7a, 0x9e, 0x77,
	0x59, 0x6d, 0xec, 0x79, 0xfb, 0x97, 0xd5, 0xc6, 0x8e, 0x57, 0x3b, 0xfe, 0x7b, 0x0b, 0x79, 0xeb,
	0x67, 0x7a, 0x97, 0x1b, 0x3b, 0xa8, 0x7e, 0x42, 0x35, 0x6d, 0xa8, 0xc9, 0x35, 0x4c, 0xa8, 0xee,
	0xe9, 0xe0, 0x5f, 0x5e, 0xd6, 0x49, 0x82, 0x31, 0xf0, 0x49, 0xa1, 0xc3, 0xe7, 0x68, 0x7f, 0x75,
	0xfb, 0x30, 0x05, 0x86, 0xeb, 0x87, 0x2d, 0xc8, 0x72, 0xdf, 0xf9, 0xad, 0x32, 0x72, 0x16, 0x04,
	0xd3, 0x8d, 0x3d, 0x34, 0xc5, 0xe7, 0xa8, 0xab, 0xf3, 0x24, 0xa1, 0x6a, 0x59, 0xd8, 0xc0, 0x34,
	0x6b, 0x92, 0x4e, 0x81, 0x16, 0x17, 0xfe, 0x0c, 0x75, 0x55, 0x2e, 0xc3, 0x68, 0x46, 0x95, 0x3d,
	0x48, 0xce, 0xfd, 0xaa, 0x1b, 0xad, 0x2a, 0x97, 0x23, 0x0b, 0x5e, 0x09, 0x39, 0xc7, 0x87, 0xa8,
	0x6e, 0x59, 0x76, 0xf2, 0xee, 0x40, 0xb8, 0xa6, 0x72, 0x69, 0x67, 0xee, 0x4b, 0xf4, 0x64, 0xca,
	0x25, 0x57, 0xd4, 0x70, 0xb6, 0xd9, 0x62, 0x35, 0x28, 0xde, 0x7e, 0x19, 0xbd, 0xd7, 0x5f, 0xc7,
	0xdf, 0xa0, 0x9a, 0xcb, 0x1a, 0xb7, 0x50, 0x7d, 0xfc, 0x7e, 0x34, 0x7a, 0x3d, 0x1e, 0x7b, 0x15,
	0xdc, 0x44, 0x3b, 0xaf, 0x09, 0x79, 0x47, 0xbc, 0x2d, 0xbc, 0x8b, 0x3a, 0xc3, 0xb7, 0xc3, 0xab,
	0xdf, 0xc7, 0x17, 0xe3, 0xf0, 0x7c, 0x78, 0x71, 0xe5, 0x55, 0x7f, 0xfe, 0x0a, 0xbd, 0x88, 0xd2,
	0x24, 0x98, 0xa6, 0xe9, 0x34, 0xe6, 0x81, 0x7d, 0x43, 0x21, 0xa7, 0x41, 0xc6, 0xd5, 0x6d, 0xaa,
	0x12, 0x2a, 0x23, 0xbe, 0xf1, 0xdb, 0xff, 0x09, 0x00, 0x00, 0xff, 0xff, 0xca, 0x40, 0x77, 0xc2,
	0x65, 0x07, 0x00, 0x00,
}

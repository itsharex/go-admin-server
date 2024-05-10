// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.0
// source: grpc/model/sys_stats.proto

package model

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

type CpuStat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PhysicalCores int32   `protobuf:"varint,1,opt,name=physicalCores,proto3" json:"physicalCores,omitempty"`
	LogicalCores  int32   `protobuf:"varint,2,opt,name=logicalCores,proto3" json:"logicalCores,omitempty"`
	UsedPercent   float32 `protobuf:"fixed32,3,opt,name=usedPercent,proto3" json:"usedPercent,omitempty"`
}

func (x *CpuStat) Reset() {
	*x = CpuStat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_model_sys_stats_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CpuStat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CpuStat) ProtoMessage() {}

func (x *CpuStat) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_model_sys_stats_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CpuStat.ProtoReflect.Descriptor instead.
func (*CpuStat) Descriptor() ([]byte, []int) {
	return file_grpc_model_sys_stats_proto_rawDescGZIP(), []int{0}
}

func (x *CpuStat) GetPhysicalCores() int32 {
	if x != nil {
		return x.PhysicalCores
	}
	return 0
}

func (x *CpuStat) GetLogicalCores() int32 {
	if x != nil {
		return x.LogicalCores
	}
	return 0
}

func (x *CpuStat) GetUsedPercent() float32 {
	if x != nil {
		return x.UsedPercent
	}
	return 0
}

type MemoryStat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total       float32 `protobuf:"fixed32,1,opt,name=total,proto3" json:"total,omitempty"`
	Used        float32 `protobuf:"fixed32,2,opt,name=used,proto3" json:"used,omitempty"`
	UsedPercent float32 `protobuf:"fixed32,3,opt,name=usedPercent,proto3" json:"usedPercent,omitempty"`
}

func (x *MemoryStat) Reset() {
	*x = MemoryStat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_model_sys_stats_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemoryStat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemoryStat) ProtoMessage() {}

func (x *MemoryStat) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_model_sys_stats_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemoryStat.ProtoReflect.Descriptor instead.
func (*MemoryStat) Descriptor() ([]byte, []int) {
	return file_grpc_model_sys_stats_proto_rawDescGZIP(), []int{1}
}

func (x *MemoryStat) GetTotal() float32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *MemoryStat) GetUsed() float32 {
	if x != nil {
		return x.Used
	}
	return 0
}

func (x *MemoryStat) GetUsedPercent() float32 {
	if x != nil {
		return x.UsedPercent
	}
	return 0
}

type DiskStat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total       float32 `protobuf:"fixed32,1,opt,name=total,proto3" json:"total,omitempty"`
	Used        float32 `protobuf:"fixed32,2,opt,name=used,proto3" json:"used,omitempty"`
	UsedPercent float32 `protobuf:"fixed32,3,opt,name=usedPercent,proto3" json:"usedPercent,omitempty"`
}

func (x *DiskStat) Reset() {
	*x = DiskStat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_model_sys_stats_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiskStat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiskStat) ProtoMessage() {}

func (x *DiskStat) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_model_sys_stats_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiskStat.ProtoReflect.Descriptor instead.
func (*DiskStat) Descriptor() ([]byte, []int) {
	return file_grpc_model_sys_stats_proto_rawDescGZIP(), []int{2}
}

func (x *DiskStat) GetTotal() float32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *DiskStat) GetUsed() float32 {
	if x != nil {
		return x.Used
	}
	return 0
}

func (x *DiskStat) GetUsedPercent() float32 {
	if x != nil {
		return x.UsedPercent
	}
	return 0
}

type SysStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip            string                 `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	Cpu           *CpuStat               `protobuf:"bytes,2,opt,name=cpu,proto3" json:"cpu,omitempty"`
	VirtualMemory *MemoryStat            `protobuf:"bytes,3,opt,name=virtualMemory,proto3" json:"virtualMemory,omitempty"`
	SwapMemory    *MemoryStat            `protobuf:"bytes,4,opt,name=swapMemory,proto3" json:"swapMemory,omitempty"`
	Disk          *DiskStat              `protobuf:"bytes,5,opt,name=disk,proto3" json:"disk,omitempty"`
	Timestamp     *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *SysStats) Reset() {
	*x = SysStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_model_sys_stats_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysStats) ProtoMessage() {}

func (x *SysStats) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_model_sys_stats_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysStats.ProtoReflect.Descriptor instead.
func (*SysStats) Descriptor() ([]byte, []int) {
	return file_grpc_model_sys_stats_proto_rawDescGZIP(), []int{3}
}

func (x *SysStats) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *SysStats) GetCpu() *CpuStat {
	if x != nil {
		return x.Cpu
	}
	return nil
}

func (x *SysStats) GetVirtualMemory() *MemoryStat {
	if x != nil {
		return x.VirtualMemory
	}
	return nil
}

func (x *SysStats) GetSwapMemory() *MemoryStat {
	if x != nil {
		return x.SwapMemory
	}
	return nil
}

func (x *SysStats) GetDisk() *DiskStat {
	if x != nil {
		return x.Disk
	}
	return nil
}

func (x *SysStats) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

var File_grpc_model_sys_stats_proto protoreflect.FileDescriptor

var file_grpc_model_sys_stats_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x73, 0x79, 0x73,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x75, 0x0a, 0x07, 0x43, 0x70, 0x75, 0x53, 0x74, 0x61, 0x74, 0x12,
	0x24, 0x0a, 0x0d, 0x70, 0x68, 0x79, 0x73, 0x69, 0x63, 0x61, 0x6c, 0x43, 0x6f, 0x72, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x70, 0x68, 0x79, 0x73, 0x69, 0x63, 0x61, 0x6c,
	0x43, 0x6f, 0x72, 0x65, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x61, 0x6c,
	0x43, 0x6f, 0x72, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x6c, 0x6f, 0x67,
	0x69, 0x63, 0x61, 0x6c, 0x43, 0x6f, 0x72, 0x65, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x75, 0x73, 0x65,
	0x64, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b,
	0x75, 0x73, 0x65, 0x64, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x22, 0x58, 0x0a, 0x0a, 0x4d,
	0x65, 0x6d, 0x6f, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12,
	0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x75,
	0x73, 0x65, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x75, 0x73, 0x65, 0x64, 0x50, 0x65, 0x72, 0x63, 0x65,
	0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x75, 0x73, 0x65, 0x64, 0x50, 0x65,
	0x72, 0x63, 0x65, 0x6e, 0x74, 0x22, 0x56, 0x0a, 0x08, 0x44, 0x69, 0x73, 0x6b, 0x53, 0x74, 0x61,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x75, 0x73, 0x65, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x75,
	0x73, 0x65, 0x64, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x0b, 0x75, 0x73, 0x65, 0x64, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x22, 0x87, 0x02,
	0x0a, 0x08, 0x53, 0x79, 0x73, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x12, 0x20, 0x0a, 0x03, 0x63, 0x70,
	0x75, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x43, 0x70, 0x75, 0x53, 0x74, 0x61, 0x74, 0x52, 0x03, 0x63, 0x70, 0x75, 0x12, 0x37, 0x0a, 0x0d,
	0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4d, 0x65, 0x6d, 0x6f,
	0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x52, 0x0d, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x4d,
	0x65, 0x6d, 0x6f, 0x72, 0x79, 0x12, 0x31, 0x0a, 0x0a, 0x73, 0x77, 0x61, 0x70, 0x4d, 0x65, 0x6d,
	0x6f, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x52, 0x0a, 0x73, 0x77,
	0x61, 0x70, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x12, 0x23, 0x0a, 0x04, 0x64, 0x69, 0x73, 0x6b,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x44,
	0x69, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x52, 0x04, 0x64, 0x69, 0x73, 0x6b, 0x12, 0x38, 0x0a,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x61, 0x72, 0x63, 0x68, 0x47, 0x65, 0x2f, 0x67, 0x6f,
	0x2d, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x61, 0x67,
	0x65, 0x6e, 0x74, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x62, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_model_sys_stats_proto_rawDescOnce sync.Once
	file_grpc_model_sys_stats_proto_rawDescData = file_grpc_model_sys_stats_proto_rawDesc
)

func file_grpc_model_sys_stats_proto_rawDescGZIP() []byte {
	file_grpc_model_sys_stats_proto_rawDescOnce.Do(func() {
		file_grpc_model_sys_stats_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_model_sys_stats_proto_rawDescData)
	})
	return file_grpc_model_sys_stats_proto_rawDescData
}

var file_grpc_model_sys_stats_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_grpc_model_sys_stats_proto_goTypes = []interface{}{
	(*CpuStat)(nil),               // 0: model.CpuStat
	(*MemoryStat)(nil),            // 1: model.MemoryStat
	(*DiskStat)(nil),              // 2: model.DiskStat
	(*SysStats)(nil),              // 3: model.SysStats
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_grpc_model_sys_stats_proto_depIdxs = []int32{
	0, // 0: model.SysStats.cpu:type_name -> model.CpuStat
	1, // 1: model.SysStats.virtualMemory:type_name -> model.MemoryStat
	1, // 2: model.SysStats.swapMemory:type_name -> model.MemoryStat
	2, // 3: model.SysStats.disk:type_name -> model.DiskStat
	4, // 4: model.SysStats.timestamp:type_name -> google.protobuf.Timestamp
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_grpc_model_sys_stats_proto_init() }
func file_grpc_model_sys_stats_proto_init() {
	if File_grpc_model_sys_stats_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_model_sys_stats_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CpuStat); i {
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
		file_grpc_model_sys_stats_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemoryStat); i {
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
		file_grpc_model_sys_stats_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DiskStat); i {
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
		file_grpc_model_sys_stats_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysStats); i {
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
			RawDescriptor: file_grpc_model_sys_stats_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_grpc_model_sys_stats_proto_goTypes,
		DependencyIndexes: file_grpc_model_sys_stats_proto_depIdxs,
		MessageInfos:      file_grpc_model_sys_stats_proto_msgTypes,
	}.Build()
	File_grpc_model_sys_stats_proto = out.File
	file_grpc_model_sys_stats_proto_rawDesc = nil
	file_grpc_model_sys_stats_proto_goTypes = nil
	file_grpc_model_sys_stats_proto_depIdxs = nil
}
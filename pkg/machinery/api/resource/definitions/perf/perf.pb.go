// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.2
// source: resource/definitions/perf/perf.proto

package perf

import (
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// CPUSpec represents the last CPU stats snapshot.
type CPUSpec struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	Cpu             []*CPUStat             `protobuf:"bytes,1,rep,name=cpu,proto3" json:"cpu,omitempty"`
	CpuTotal        *CPUStat               `protobuf:"bytes,2,opt,name=cpu_total,json=cpuTotal,proto3" json:"cpu_total,omitempty"`
	IrqTotal        uint64                 `protobuf:"varint,3,opt,name=irq_total,json=irqTotal,proto3" json:"irq_total,omitempty"`
	ContextSwitches uint64                 `protobuf:"varint,4,opt,name=context_switches,json=contextSwitches,proto3" json:"context_switches,omitempty"`
	ProcessCreated  uint64                 `protobuf:"varint,5,opt,name=process_created,json=processCreated,proto3" json:"process_created,omitempty"`
	ProcessRunning  uint64                 `protobuf:"varint,6,opt,name=process_running,json=processRunning,proto3" json:"process_running,omitempty"`
	ProcessBlocked  uint64                 `protobuf:"varint,7,opt,name=process_blocked,json=processBlocked,proto3" json:"process_blocked,omitempty"`
	SoftIrqTotal    uint64                 `protobuf:"varint,8,opt,name=soft_irq_total,json=softIrqTotal,proto3" json:"soft_irq_total,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *CPUSpec) Reset() {
	*x = CPUSpec{}
	mi := &file_resource_definitions_perf_perf_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CPUSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CPUSpec) ProtoMessage() {}

func (x *CPUSpec) ProtoReflect() protoreflect.Message {
	mi := &file_resource_definitions_perf_perf_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CPUSpec.ProtoReflect.Descriptor instead.
func (*CPUSpec) Descriptor() ([]byte, []int) {
	return file_resource_definitions_perf_perf_proto_rawDescGZIP(), []int{0}
}

func (x *CPUSpec) GetCpu() []*CPUStat {
	if x != nil {
		return x.Cpu
	}
	return nil
}

func (x *CPUSpec) GetCpuTotal() *CPUStat {
	if x != nil {
		return x.CpuTotal
	}
	return nil
}

func (x *CPUSpec) GetIrqTotal() uint64 {
	if x != nil {
		return x.IrqTotal
	}
	return 0
}

func (x *CPUSpec) GetContextSwitches() uint64 {
	if x != nil {
		return x.ContextSwitches
	}
	return 0
}

func (x *CPUSpec) GetProcessCreated() uint64 {
	if x != nil {
		return x.ProcessCreated
	}
	return 0
}

func (x *CPUSpec) GetProcessRunning() uint64 {
	if x != nil {
		return x.ProcessRunning
	}
	return 0
}

func (x *CPUSpec) GetProcessBlocked() uint64 {
	if x != nil {
		return x.ProcessBlocked
	}
	return 0
}

func (x *CPUSpec) GetSoftIrqTotal() uint64 {
	if x != nil {
		return x.SoftIrqTotal
	}
	return 0
}

// CPUStat represents a single cpu stat.
type CPUStat struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	User          float64                `protobuf:"fixed64,1,opt,name=user,proto3" json:"user,omitempty"`
	Nice          float64                `protobuf:"fixed64,2,opt,name=nice,proto3" json:"nice,omitempty"`
	System        float64                `protobuf:"fixed64,3,opt,name=system,proto3" json:"system,omitempty"`
	Idle          float64                `protobuf:"fixed64,4,opt,name=idle,proto3" json:"idle,omitempty"`
	Iowait        float64                `protobuf:"fixed64,5,opt,name=iowait,proto3" json:"iowait,omitempty"`
	Irq           float64                `protobuf:"fixed64,6,opt,name=irq,proto3" json:"irq,omitempty"`
	SoftIrq       float64                `protobuf:"fixed64,7,opt,name=soft_irq,json=softIrq,proto3" json:"soft_irq,omitempty"`
	Steal         float64                `protobuf:"fixed64,8,opt,name=steal,proto3" json:"steal,omitempty"`
	Guest         float64                `protobuf:"fixed64,9,opt,name=guest,proto3" json:"guest,omitempty"`
	GuestNice     float64                `protobuf:"fixed64,10,opt,name=guest_nice,json=guestNice,proto3" json:"guest_nice,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CPUStat) Reset() {
	*x = CPUStat{}
	mi := &file_resource_definitions_perf_perf_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CPUStat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CPUStat) ProtoMessage() {}

func (x *CPUStat) ProtoReflect() protoreflect.Message {
	mi := &file_resource_definitions_perf_perf_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CPUStat.ProtoReflect.Descriptor instead.
func (*CPUStat) Descriptor() ([]byte, []int) {
	return file_resource_definitions_perf_perf_proto_rawDescGZIP(), []int{1}
}

func (x *CPUStat) GetUser() float64 {
	if x != nil {
		return x.User
	}
	return 0
}

func (x *CPUStat) GetNice() float64 {
	if x != nil {
		return x.Nice
	}
	return 0
}

func (x *CPUStat) GetSystem() float64 {
	if x != nil {
		return x.System
	}
	return 0
}

func (x *CPUStat) GetIdle() float64 {
	if x != nil {
		return x.Idle
	}
	return 0
}

func (x *CPUStat) GetIowait() float64 {
	if x != nil {
		return x.Iowait
	}
	return 0
}

func (x *CPUStat) GetIrq() float64 {
	if x != nil {
		return x.Irq
	}
	return 0
}

func (x *CPUStat) GetSoftIrq() float64 {
	if x != nil {
		return x.SoftIrq
	}
	return 0
}

func (x *CPUStat) GetSteal() float64 {
	if x != nil {
		return x.Steal
	}
	return 0
}

func (x *CPUStat) GetGuest() float64 {
	if x != nil {
		return x.Guest
	}
	return 0
}

func (x *CPUStat) GetGuestNice() float64 {
	if x != nil {
		return x.GuestNice
	}
	return 0
}

// MemorySpec represents the last Memory stats snapshot.
type MemorySpec struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	MemTotal          uint64                 `protobuf:"varint,1,opt,name=mem_total,json=memTotal,proto3" json:"mem_total,omitempty"`
	MemUsed           uint64                 `protobuf:"varint,2,opt,name=mem_used,json=memUsed,proto3" json:"mem_used,omitempty"`
	MemAvailable      uint64                 `protobuf:"varint,3,opt,name=mem_available,json=memAvailable,proto3" json:"mem_available,omitempty"`
	Buffers           uint64                 `protobuf:"varint,4,opt,name=buffers,proto3" json:"buffers,omitempty"`
	Cached            uint64                 `protobuf:"varint,5,opt,name=cached,proto3" json:"cached,omitempty"`
	SwapCached        uint64                 `protobuf:"varint,6,opt,name=swap_cached,json=swapCached,proto3" json:"swap_cached,omitempty"`
	Active            uint64                 `protobuf:"varint,7,opt,name=active,proto3" json:"active,omitempty"`
	Inactive          uint64                 `protobuf:"varint,8,opt,name=inactive,proto3" json:"inactive,omitempty"`
	ActiveAnon        uint64                 `protobuf:"varint,9,opt,name=active_anon,json=activeAnon,proto3" json:"active_anon,omitempty"`
	InactiveAnon      uint64                 `protobuf:"varint,10,opt,name=inactive_anon,json=inactiveAnon,proto3" json:"inactive_anon,omitempty"`
	ActiveFile        uint64                 `protobuf:"varint,11,opt,name=active_file,json=activeFile,proto3" json:"active_file,omitempty"`
	InactiveFile      uint64                 `protobuf:"varint,12,opt,name=inactive_file,json=inactiveFile,proto3" json:"inactive_file,omitempty"`
	Unevictable       uint64                 `protobuf:"varint,13,opt,name=unevictable,proto3" json:"unevictable,omitempty"`
	Mlocked           uint64                 `protobuf:"varint,14,opt,name=mlocked,proto3" json:"mlocked,omitempty"`
	SwapTotal         uint64                 `protobuf:"varint,15,opt,name=swap_total,json=swapTotal,proto3" json:"swap_total,omitempty"`
	SwapFree          uint64                 `protobuf:"varint,16,opt,name=swap_free,json=swapFree,proto3" json:"swap_free,omitempty"`
	Dirty             uint64                 `protobuf:"varint,17,opt,name=dirty,proto3" json:"dirty,omitempty"`
	Writeback         uint64                 `protobuf:"varint,18,opt,name=writeback,proto3" json:"writeback,omitempty"`
	AnonPages         uint64                 `protobuf:"varint,19,opt,name=anon_pages,json=anonPages,proto3" json:"anon_pages,omitempty"`
	Mapped            uint64                 `protobuf:"varint,20,opt,name=mapped,proto3" json:"mapped,omitempty"`
	Shmem             uint64                 `protobuf:"varint,21,opt,name=shmem,proto3" json:"shmem,omitempty"`
	Slab              uint64                 `protobuf:"varint,22,opt,name=slab,proto3" json:"slab,omitempty"`
	SReclaimable      uint64                 `protobuf:"varint,23,opt,name=s_reclaimable,json=sReclaimable,proto3" json:"s_reclaimable,omitempty"`
	SUnreclaim        uint64                 `protobuf:"varint,24,opt,name=s_unreclaim,json=sUnreclaim,proto3" json:"s_unreclaim,omitempty"`
	KernelStack       uint64                 `protobuf:"varint,25,opt,name=kernel_stack,json=kernelStack,proto3" json:"kernel_stack,omitempty"`
	PageTables        uint64                 `protobuf:"varint,26,opt,name=page_tables,json=pageTables,proto3" json:"page_tables,omitempty"`
	NfSunstable       uint64                 `protobuf:"varint,27,opt,name=nf_sunstable,json=nfSunstable,proto3" json:"nf_sunstable,omitempty"`
	Bounce            uint64                 `protobuf:"varint,28,opt,name=bounce,proto3" json:"bounce,omitempty"`
	WritebackTmp      uint64                 `protobuf:"varint,29,opt,name=writeback_tmp,json=writebackTmp,proto3" json:"writeback_tmp,omitempty"`
	CommitLimit       uint64                 `protobuf:"varint,30,opt,name=commit_limit,json=commitLimit,proto3" json:"commit_limit,omitempty"`
	CommittedAs       uint64                 `protobuf:"varint,31,opt,name=committed_as,json=committedAs,proto3" json:"committed_as,omitempty"`
	VmallocTotal      uint64                 `protobuf:"varint,32,opt,name=vmalloc_total,json=vmallocTotal,proto3" json:"vmalloc_total,omitempty"`
	VmallocUsed       uint64                 `protobuf:"varint,33,opt,name=vmalloc_used,json=vmallocUsed,proto3" json:"vmalloc_used,omitempty"`
	VmallocChunk      uint64                 `protobuf:"varint,34,opt,name=vmalloc_chunk,json=vmallocChunk,proto3" json:"vmalloc_chunk,omitempty"`
	HardwareCorrupted uint64                 `protobuf:"varint,35,opt,name=hardware_corrupted,json=hardwareCorrupted,proto3" json:"hardware_corrupted,omitempty"`
	AnonHugePages     uint64                 `protobuf:"varint,36,opt,name=anon_huge_pages,json=anonHugePages,proto3" json:"anon_huge_pages,omitempty"`
	ShmemHugePages    uint64                 `protobuf:"varint,37,opt,name=shmem_huge_pages,json=shmemHugePages,proto3" json:"shmem_huge_pages,omitempty"`
	ShmemPmdMapped    uint64                 `protobuf:"varint,38,opt,name=shmem_pmd_mapped,json=shmemPmdMapped,proto3" json:"shmem_pmd_mapped,omitempty"`
	CmaTotal          uint64                 `protobuf:"varint,39,opt,name=cma_total,json=cmaTotal,proto3" json:"cma_total,omitempty"`
	CmaFree           uint64                 `protobuf:"varint,40,opt,name=cma_free,json=cmaFree,proto3" json:"cma_free,omitempty"`
	HugePagesTotal    uint64                 `protobuf:"varint,41,opt,name=huge_pages_total,json=hugePagesTotal,proto3" json:"huge_pages_total,omitempty"`
	HugePagesFree     uint64                 `protobuf:"varint,42,opt,name=huge_pages_free,json=hugePagesFree,proto3" json:"huge_pages_free,omitempty"`
	HugePagesRsvd     uint64                 `protobuf:"varint,43,opt,name=huge_pages_rsvd,json=hugePagesRsvd,proto3" json:"huge_pages_rsvd,omitempty"`
	HugePagesSurp     uint64                 `protobuf:"varint,44,opt,name=huge_pages_surp,json=hugePagesSurp,proto3" json:"huge_pages_surp,omitempty"`
	Hugepagesize      uint64                 `protobuf:"varint,45,opt,name=hugepagesize,proto3" json:"hugepagesize,omitempty"`
	DirectMap4K       uint64                 `protobuf:"varint,46,opt,name=direct_map4k,json=directMap4k,proto3" json:"direct_map4k,omitempty"`
	DirectMap2M       uint64                 `protobuf:"varint,47,opt,name=direct_map2m,json=directMap2m,proto3" json:"direct_map2m,omitempty"`
	DirectMap1G       uint64                 `protobuf:"varint,48,opt,name=direct_map1g,json=directMap1g,proto3" json:"direct_map1g,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *MemorySpec) Reset() {
	*x = MemorySpec{}
	mi := &file_resource_definitions_perf_perf_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MemorySpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemorySpec) ProtoMessage() {}

func (x *MemorySpec) ProtoReflect() protoreflect.Message {
	mi := &file_resource_definitions_perf_perf_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemorySpec.ProtoReflect.Descriptor instead.
func (*MemorySpec) Descriptor() ([]byte, []int) {
	return file_resource_definitions_perf_perf_proto_rawDescGZIP(), []int{2}
}

func (x *MemorySpec) GetMemTotal() uint64 {
	if x != nil {
		return x.MemTotal
	}
	return 0
}

func (x *MemorySpec) GetMemUsed() uint64 {
	if x != nil {
		return x.MemUsed
	}
	return 0
}

func (x *MemorySpec) GetMemAvailable() uint64 {
	if x != nil {
		return x.MemAvailable
	}
	return 0
}

func (x *MemorySpec) GetBuffers() uint64 {
	if x != nil {
		return x.Buffers
	}
	return 0
}

func (x *MemorySpec) GetCached() uint64 {
	if x != nil {
		return x.Cached
	}
	return 0
}

func (x *MemorySpec) GetSwapCached() uint64 {
	if x != nil {
		return x.SwapCached
	}
	return 0
}

func (x *MemorySpec) GetActive() uint64 {
	if x != nil {
		return x.Active
	}
	return 0
}

func (x *MemorySpec) GetInactive() uint64 {
	if x != nil {
		return x.Inactive
	}
	return 0
}

func (x *MemorySpec) GetActiveAnon() uint64 {
	if x != nil {
		return x.ActiveAnon
	}
	return 0
}

func (x *MemorySpec) GetInactiveAnon() uint64 {
	if x != nil {
		return x.InactiveAnon
	}
	return 0
}

func (x *MemorySpec) GetActiveFile() uint64 {
	if x != nil {
		return x.ActiveFile
	}
	return 0
}

func (x *MemorySpec) GetInactiveFile() uint64 {
	if x != nil {
		return x.InactiveFile
	}
	return 0
}

func (x *MemorySpec) GetUnevictable() uint64 {
	if x != nil {
		return x.Unevictable
	}
	return 0
}

func (x *MemorySpec) GetMlocked() uint64 {
	if x != nil {
		return x.Mlocked
	}
	return 0
}

func (x *MemorySpec) GetSwapTotal() uint64 {
	if x != nil {
		return x.SwapTotal
	}
	return 0
}

func (x *MemorySpec) GetSwapFree() uint64 {
	if x != nil {
		return x.SwapFree
	}
	return 0
}

func (x *MemorySpec) GetDirty() uint64 {
	if x != nil {
		return x.Dirty
	}
	return 0
}

func (x *MemorySpec) GetWriteback() uint64 {
	if x != nil {
		return x.Writeback
	}
	return 0
}

func (x *MemorySpec) GetAnonPages() uint64 {
	if x != nil {
		return x.AnonPages
	}
	return 0
}

func (x *MemorySpec) GetMapped() uint64 {
	if x != nil {
		return x.Mapped
	}
	return 0
}

func (x *MemorySpec) GetShmem() uint64 {
	if x != nil {
		return x.Shmem
	}
	return 0
}

func (x *MemorySpec) GetSlab() uint64 {
	if x != nil {
		return x.Slab
	}
	return 0
}

func (x *MemorySpec) GetSReclaimable() uint64 {
	if x != nil {
		return x.SReclaimable
	}
	return 0
}

func (x *MemorySpec) GetSUnreclaim() uint64 {
	if x != nil {
		return x.SUnreclaim
	}
	return 0
}

func (x *MemorySpec) GetKernelStack() uint64 {
	if x != nil {
		return x.KernelStack
	}
	return 0
}

func (x *MemorySpec) GetPageTables() uint64 {
	if x != nil {
		return x.PageTables
	}
	return 0
}

func (x *MemorySpec) GetNfSunstable() uint64 {
	if x != nil {
		return x.NfSunstable
	}
	return 0
}

func (x *MemorySpec) GetBounce() uint64 {
	if x != nil {
		return x.Bounce
	}
	return 0
}

func (x *MemorySpec) GetWritebackTmp() uint64 {
	if x != nil {
		return x.WritebackTmp
	}
	return 0
}

func (x *MemorySpec) GetCommitLimit() uint64 {
	if x != nil {
		return x.CommitLimit
	}
	return 0
}

func (x *MemorySpec) GetCommittedAs() uint64 {
	if x != nil {
		return x.CommittedAs
	}
	return 0
}

func (x *MemorySpec) GetVmallocTotal() uint64 {
	if x != nil {
		return x.VmallocTotal
	}
	return 0
}

func (x *MemorySpec) GetVmallocUsed() uint64 {
	if x != nil {
		return x.VmallocUsed
	}
	return 0
}

func (x *MemorySpec) GetVmallocChunk() uint64 {
	if x != nil {
		return x.VmallocChunk
	}
	return 0
}

func (x *MemorySpec) GetHardwareCorrupted() uint64 {
	if x != nil {
		return x.HardwareCorrupted
	}
	return 0
}

func (x *MemorySpec) GetAnonHugePages() uint64 {
	if x != nil {
		return x.AnonHugePages
	}
	return 0
}

func (x *MemorySpec) GetShmemHugePages() uint64 {
	if x != nil {
		return x.ShmemHugePages
	}
	return 0
}

func (x *MemorySpec) GetShmemPmdMapped() uint64 {
	if x != nil {
		return x.ShmemPmdMapped
	}
	return 0
}

func (x *MemorySpec) GetCmaTotal() uint64 {
	if x != nil {
		return x.CmaTotal
	}
	return 0
}

func (x *MemorySpec) GetCmaFree() uint64 {
	if x != nil {
		return x.CmaFree
	}
	return 0
}

func (x *MemorySpec) GetHugePagesTotal() uint64 {
	if x != nil {
		return x.HugePagesTotal
	}
	return 0
}

func (x *MemorySpec) GetHugePagesFree() uint64 {
	if x != nil {
		return x.HugePagesFree
	}
	return 0
}

func (x *MemorySpec) GetHugePagesRsvd() uint64 {
	if x != nil {
		return x.HugePagesRsvd
	}
	return 0
}

func (x *MemorySpec) GetHugePagesSurp() uint64 {
	if x != nil {
		return x.HugePagesSurp
	}
	return 0
}

func (x *MemorySpec) GetHugepagesize() uint64 {
	if x != nil {
		return x.Hugepagesize
	}
	return 0
}

func (x *MemorySpec) GetDirectMap4K() uint64 {
	if x != nil {
		return x.DirectMap4K
	}
	return 0
}

func (x *MemorySpec) GetDirectMap2M() uint64 {
	if x != nil {
		return x.DirectMap2M
	}
	return 0
}

func (x *MemorySpec) GetDirectMap1G() uint64 {
	if x != nil {
		return x.DirectMap1G
	}
	return 0
}

var File_resource_definitions_perf_perf_proto protoreflect.FileDescriptor

const file_resource_definitions_perf_perf_proto_rawDesc = "" +
	"\n" +
	"$resource/definitions/perf/perf.proto\x12\x1ftalos.resource.definitions.perf\"\xf5\x02\n" +
	"\aCPUSpec\x12:\n" +
	"\x03cpu\x18\x01 \x03(\v2(.talos.resource.definitions.perf.CPUStatR\x03cpu\x12E\n" +
	"\tcpu_total\x18\x02 \x01(\v2(.talos.resource.definitions.perf.CPUStatR\bcpuTotal\x12\x1b\n" +
	"\tirq_total\x18\x03 \x01(\x04R\birqTotal\x12)\n" +
	"\x10context_switches\x18\x04 \x01(\x04R\x0fcontextSwitches\x12'\n" +
	"\x0fprocess_created\x18\x05 \x01(\x04R\x0eprocessCreated\x12'\n" +
	"\x0fprocess_running\x18\x06 \x01(\x04R\x0eprocessRunning\x12'\n" +
	"\x0fprocess_blocked\x18\a \x01(\x04R\x0eprocessBlocked\x12$\n" +
	"\x0esoft_irq_total\x18\b \x01(\x04R\fsoftIrqTotal\"\xed\x01\n" +
	"\aCPUStat\x12\x12\n" +
	"\x04user\x18\x01 \x01(\x01R\x04user\x12\x12\n" +
	"\x04nice\x18\x02 \x01(\x01R\x04nice\x12\x16\n" +
	"\x06system\x18\x03 \x01(\x01R\x06system\x12\x12\n" +
	"\x04idle\x18\x04 \x01(\x01R\x04idle\x12\x16\n" +
	"\x06iowait\x18\x05 \x01(\x01R\x06iowait\x12\x10\n" +
	"\x03irq\x18\x06 \x01(\x01R\x03irq\x12\x19\n" +
	"\bsoft_irq\x18\a \x01(\x01R\asoftIrq\x12\x14\n" +
	"\x05steal\x18\b \x01(\x01R\x05steal\x12\x14\n" +
	"\x05guest\x18\t \x01(\x01R\x05guest\x12\x1d\n" +
	"\n" +
	"guest_nice\x18\n" +
	" \x01(\x01R\tguestNice\"\xb8\f\n" +
	"\n" +
	"MemorySpec\x12\x1b\n" +
	"\tmem_total\x18\x01 \x01(\x04R\bmemTotal\x12\x19\n" +
	"\bmem_used\x18\x02 \x01(\x04R\amemUsed\x12#\n" +
	"\rmem_available\x18\x03 \x01(\x04R\fmemAvailable\x12\x18\n" +
	"\abuffers\x18\x04 \x01(\x04R\abuffers\x12\x16\n" +
	"\x06cached\x18\x05 \x01(\x04R\x06cached\x12\x1f\n" +
	"\vswap_cached\x18\x06 \x01(\x04R\n" +
	"swapCached\x12\x16\n" +
	"\x06active\x18\a \x01(\x04R\x06active\x12\x1a\n" +
	"\binactive\x18\b \x01(\x04R\binactive\x12\x1f\n" +
	"\vactive_anon\x18\t \x01(\x04R\n" +
	"activeAnon\x12#\n" +
	"\rinactive_anon\x18\n" +
	" \x01(\x04R\finactiveAnon\x12\x1f\n" +
	"\vactive_file\x18\v \x01(\x04R\n" +
	"activeFile\x12#\n" +
	"\rinactive_file\x18\f \x01(\x04R\finactiveFile\x12 \n" +
	"\vunevictable\x18\r \x01(\x04R\vunevictable\x12\x18\n" +
	"\amlocked\x18\x0e \x01(\x04R\amlocked\x12\x1d\n" +
	"\n" +
	"swap_total\x18\x0f \x01(\x04R\tswapTotal\x12\x1b\n" +
	"\tswap_free\x18\x10 \x01(\x04R\bswapFree\x12\x14\n" +
	"\x05dirty\x18\x11 \x01(\x04R\x05dirty\x12\x1c\n" +
	"\twriteback\x18\x12 \x01(\x04R\twriteback\x12\x1d\n" +
	"\n" +
	"anon_pages\x18\x13 \x01(\x04R\tanonPages\x12\x16\n" +
	"\x06mapped\x18\x14 \x01(\x04R\x06mapped\x12\x14\n" +
	"\x05shmem\x18\x15 \x01(\x04R\x05shmem\x12\x12\n" +
	"\x04slab\x18\x16 \x01(\x04R\x04slab\x12#\n" +
	"\rs_reclaimable\x18\x17 \x01(\x04R\fsReclaimable\x12\x1f\n" +
	"\vs_unreclaim\x18\x18 \x01(\x04R\n" +
	"sUnreclaim\x12!\n" +
	"\fkernel_stack\x18\x19 \x01(\x04R\vkernelStack\x12\x1f\n" +
	"\vpage_tables\x18\x1a \x01(\x04R\n" +
	"pageTables\x12!\n" +
	"\fnf_sunstable\x18\x1b \x01(\x04R\vnfSunstable\x12\x16\n" +
	"\x06bounce\x18\x1c \x01(\x04R\x06bounce\x12#\n" +
	"\rwriteback_tmp\x18\x1d \x01(\x04R\fwritebackTmp\x12!\n" +
	"\fcommit_limit\x18\x1e \x01(\x04R\vcommitLimit\x12!\n" +
	"\fcommitted_as\x18\x1f \x01(\x04R\vcommittedAs\x12#\n" +
	"\rvmalloc_total\x18  \x01(\x04R\fvmallocTotal\x12!\n" +
	"\fvmalloc_used\x18! \x01(\x04R\vvmallocUsed\x12#\n" +
	"\rvmalloc_chunk\x18\" \x01(\x04R\fvmallocChunk\x12-\n" +
	"\x12hardware_corrupted\x18# \x01(\x04R\x11hardwareCorrupted\x12&\n" +
	"\x0fanon_huge_pages\x18$ \x01(\x04R\ranonHugePages\x12(\n" +
	"\x10shmem_huge_pages\x18% \x01(\x04R\x0eshmemHugePages\x12(\n" +
	"\x10shmem_pmd_mapped\x18& \x01(\x04R\x0eshmemPmdMapped\x12\x1b\n" +
	"\tcma_total\x18' \x01(\x04R\bcmaTotal\x12\x19\n" +
	"\bcma_free\x18( \x01(\x04R\acmaFree\x12(\n" +
	"\x10huge_pages_total\x18) \x01(\x04R\x0ehugePagesTotal\x12&\n" +
	"\x0fhuge_pages_free\x18* \x01(\x04R\rhugePagesFree\x12&\n" +
	"\x0fhuge_pages_rsvd\x18+ \x01(\x04R\rhugePagesRsvd\x12&\n" +
	"\x0fhuge_pages_surp\x18, \x01(\x04R\rhugePagesSurp\x12\"\n" +
	"\fhugepagesize\x18- \x01(\x04R\fhugepagesize\x12!\n" +
	"\fdirect_map4k\x18. \x01(\x04R\vdirectMap4k\x12!\n" +
	"\fdirect_map2m\x18/ \x01(\x04R\vdirectMap2m\x12!\n" +
	"\fdirect_map1g\x180 \x01(\x04R\vdirectMap1gBr\n" +
	"'dev.talos.api.resource.definitions.perfZGgithub.com/siderolabs/talos/pkg/machinery/api/resource/definitions/perfb\x06proto3"

var (
	file_resource_definitions_perf_perf_proto_rawDescOnce sync.Once
	file_resource_definitions_perf_perf_proto_rawDescData []byte
)

func file_resource_definitions_perf_perf_proto_rawDescGZIP() []byte {
	file_resource_definitions_perf_perf_proto_rawDescOnce.Do(func() {
		file_resource_definitions_perf_perf_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_resource_definitions_perf_perf_proto_rawDesc), len(file_resource_definitions_perf_perf_proto_rawDesc)))
	})
	return file_resource_definitions_perf_perf_proto_rawDescData
}

var file_resource_definitions_perf_perf_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_resource_definitions_perf_perf_proto_goTypes = []any{
	(*CPUSpec)(nil),    // 0: talos.resource.definitions.perf.CPUSpec
	(*CPUStat)(nil),    // 1: talos.resource.definitions.perf.CPUStat
	(*MemorySpec)(nil), // 2: talos.resource.definitions.perf.MemorySpec
}
var file_resource_definitions_perf_perf_proto_depIdxs = []int32{
	1, // 0: talos.resource.definitions.perf.CPUSpec.cpu:type_name -> talos.resource.definitions.perf.CPUStat
	1, // 1: talos.resource.definitions.perf.CPUSpec.cpu_total:type_name -> talos.resource.definitions.perf.CPUStat
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_resource_definitions_perf_perf_proto_init() }
func file_resource_definitions_perf_perf_proto_init() {
	if File_resource_definitions_perf_perf_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_resource_definitions_perf_perf_proto_rawDesc), len(file_resource_definitions_perf_perf_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_resource_definitions_perf_perf_proto_goTypes,
		DependencyIndexes: file_resource_definitions_perf_perf_proto_depIdxs,
		MessageInfos:      file_resource_definitions_perf_perf_proto_msgTypes,
	}.Build()
	File_resource_definitions_perf_perf_proto = out.File
	file_resource_definitions_perf_perf_proto_goTypes = nil
	file_resource_definitions_perf_perf_proto_depIdxs = nil
}

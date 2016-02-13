// Code generated by protoc-gen-go.
// source: types.proto
// DO NOT EDIT!

/*
Package api is a generated protocol buffer package.

It is generated from these files:
	types.proto
	api.proto

It has these top-level messages:
	Node
	Spec
	TaskStatus
	Task
	Job
	Update
	RegisterNodeRequest
	RegisterNodeResponse
	UpdateNodeStatusRequest
	UpdateNodeStatusResponse
	ListNodesRequest
	ListNodesResponse
	DrainNodeRequest
	DrainNodeResponse
	CreateTaskRequest
	CreateTaskResponse
	GetTasksRequest
	GetTasksResponse
	RemoveTaskRequest
	RemoveTaskResponse
	ListTasksRequest
	ListTasksResponse
	CreateJobRequest
	CreateJobResponse
	GetJobRequest
	GetJobResponse
	UpdateJobRequest
	UpdateJobResponse
	RemoveJobRequest
	RemoveJobResponse
	ListJobsRequest
	ListJobsResponse
	UpdateTaskStatusRequest
	UpdateTaskStatusResponse
	WatchTasksRequest
	WatchTasksResponse
*/
package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type NodeStatus int32

const (
	NodeStatus_READY NodeStatus = 0
	NodeStatus_DOWN  NodeStatus = 1
)

var NodeStatus_name = map[int32]string{
	0: "READY",
	1: "DOWN",
}
var NodeStatus_value = map[string]int32{
	"READY": 0,
	"DOWN":  1,
}

func (x NodeStatus) String() string {
	return proto.EnumName(NodeStatus_name, int32(x))
}
func (NodeStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type TaskStatus_State int32

const (
	TaskStatus_NEW       TaskStatus_State = 0
	TaskStatus_ASSIGNED  TaskStatus_State = 1
	TaskStatus_PREPARING TaskStatus_State = 2
	TaskStatus_READY     TaskStatus_State = 3
	TaskStatus_STARTING  TaskStatus_State = 4
	TaskStatus_RUNNING   TaskStatus_State = 5
	TaskStatus_SHUTDOWN  TaskStatus_State = 6
	TaskStatus_COMPLETE  TaskStatus_State = 7
	TaskStatus_FAILED    TaskStatus_State = 8
	TaskStatus_REJECTED  TaskStatus_State = 9
	TaskStatus_FINALIZE  TaskStatus_State = 10
	TaskStatus_DEAD      TaskStatus_State = 11
)

var TaskStatus_State_name = map[int32]string{
	0:  "NEW",
	1:  "ASSIGNED",
	2:  "PREPARING",
	3:  "READY",
	4:  "STARTING",
	5:  "RUNNING",
	6:  "SHUTDOWN",
	7:  "COMPLETE",
	8:  "FAILED",
	9:  "REJECTED",
	10: "FINALIZE",
	11: "DEAD",
}
var TaskStatus_State_value = map[string]int32{
	"NEW":       0,
	"ASSIGNED":  1,
	"PREPARING": 2,
	"READY":     3,
	"STARTING":  4,
	"RUNNING":   5,
	"SHUTDOWN":  6,
	"COMPLETE":  7,
	"FAILED":    8,
	"REJECTED":  9,
	"FINALIZE":  10,
	"DEAD":      11,
}

func (x TaskStatus_State) String() string {
	return proto.EnumName(TaskStatus_State_name, int32(x))
}
func (TaskStatus_State) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

type Node struct {
	Id      string     `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name    string     `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Ip      string     `protobuf:"bytes,4,opt,name=ip" json:"ip,omitempty"`
	Drained bool       `protobuf:"varint,5,opt,name=drained" json:"drained,omitempty"`
	Status  NodeStatus `protobuf:"varint,3,opt,name=status,enum=api.NodeStatus" json:"status,omitempty"`
}

func (m *Node) Reset()                    { *m = Node{} }
func (m *Node) String() string            { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()               {}
func (*Node) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// Spec defines the properties of a Job. As tasks are created, they gain the
// Job specification.
type Spec struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// remote specifies a URL to fetch task bundle data. This can either
	// be a git repository or blobster repository.
	Remote     string `protobuf:"bytes,2,opt,name=remote" json:"remote,omitempty"`
	Bundle     string `protobuf:"bytes,3,opt,name=bundle" json:"bundle,omitempty"`
	Entrypoint string `protobuf:"bytes,4,opt,name=entrypoint" json:"entrypoint,omitempty"`
	Instances  int64  `protobuf:"varint,5,opt,name=instances" json:"instances,omitempty"`
}

func (m *Spec) Reset()                    { *m = Spec{} }
func (m *Spec) String() string            { return proto.CompactTextString(m) }
func (*Spec) ProtoMessage()               {}
func (*Spec) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type TaskStatus struct {
	State   TaskStatus_State `protobuf:"varint,2,opt,name=state,enum=api.TaskStatus_State" json:"state,omitempty"`
	Message string           `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
}

func (m *TaskStatus) Reset()                    { *m = TaskStatus{} }
func (m *TaskStatus) String() string            { return proto.CompactTextString(m) }
func (*TaskStatus) ProtoMessage()               {}
func (*TaskStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type Task struct {
	Id     string      `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	JobId  string      `protobuf:"bytes,2,opt,name=job_id" json:"job_id,omitempty"`
	NodeId string      `protobuf:"bytes,3,opt,name=node_id" json:"node_id,omitempty"`
	Spec   *Spec       `protobuf:"bytes,4,opt,name=spec" json:"spec,omitempty"`
	Status *TaskStatus `protobuf:"bytes,5,opt,name=status" json:"status,omitempty"`
	// Networking state
	NetId      string   `protobuf:"bytes,6,opt,name=net_id" json:"net_id,omitempty"`
	EpId       string   `protobuf:"bytes,7,opt,name=ep_id" json:"ep_id,omitempty"`
	Ip         string   `protobuf:"bytes,8,opt,name=ip" json:"ip,omitempty"`
	Gateway    string   `protobuf:"bytes,9,opt,name=gateway" json:"gateway,omitempty"`
	DriverInfo []string `protobuf:"bytes,10,rep,name=driver_info" json:"driver_info,omitempty"`
}

func (m *Task) Reset()                    { *m = Task{} }
func (m *Task) String() string            { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()               {}
func (*Task) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Task) GetSpec() *Spec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *Task) GetStatus() *TaskStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type Job struct {
	Id   string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Spec *Spec  `protobuf:"bytes,2,opt,name=spec" json:"spec,omitempty"`
}

func (m *Job) Reset()                    { *m = Job{} }
func (m *Job) String() string            { return proto.CompactTextString(m) }
func (*Job) ProtoMessage()               {}
func (*Job) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Job) GetSpec() *Spec {
	if m != nil {
		return m.Spec
	}
	return nil
}

type Update struct {
	// Types that are valid to be assigned to Update:
	//	*Update_UpdateNode
	//	*Update_UpdateTask
	//	*Update_UpdateJob
	//	*Update_DeleteNode
	//	*Update_DeleteTask
	//	*Update_DeleteJob
	Update isUpdate_Update `protobuf_oneof:"update"`
}

func (m *Update) Reset()                    { *m = Update{} }
func (m *Update) String() string            { return proto.CompactTextString(m) }
func (*Update) ProtoMessage()               {}
func (*Update) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type isUpdate_Update interface {
	isUpdate_Update()
}

type Update_UpdateNode struct {
	UpdateNode *Node `protobuf:"bytes,1,opt,name=updateNode,oneof"`
}
type Update_UpdateTask struct {
	UpdateTask *Task `protobuf:"bytes,2,opt,name=updateTask,oneof"`
}
type Update_UpdateJob struct {
	UpdateJob *Job `protobuf:"bytes,4,opt,name=updateJob,oneof"`
}
type Update_DeleteNode struct {
	DeleteNode string `protobuf:"bytes,5,opt,name=deleteNode,oneof"`
}
type Update_DeleteTask struct {
	DeleteTask string `protobuf:"bytes,6,opt,name=deleteTask,oneof"`
}
type Update_DeleteJob struct {
	DeleteJob string `protobuf:"bytes,8,opt,name=deleteJob,oneof"`
}

func (*Update_UpdateNode) isUpdate_Update() {}
func (*Update_UpdateTask) isUpdate_Update() {}
func (*Update_UpdateJob) isUpdate_Update()  {}
func (*Update_DeleteNode) isUpdate_Update() {}
func (*Update_DeleteTask) isUpdate_Update() {}
func (*Update_DeleteJob) isUpdate_Update()  {}

func (m *Update) GetUpdate() isUpdate_Update {
	if m != nil {
		return m.Update
	}
	return nil
}

func (m *Update) GetUpdateNode() *Node {
	if x, ok := m.GetUpdate().(*Update_UpdateNode); ok {
		return x.UpdateNode
	}
	return nil
}

func (m *Update) GetUpdateTask() *Task {
	if x, ok := m.GetUpdate().(*Update_UpdateTask); ok {
		return x.UpdateTask
	}
	return nil
}

func (m *Update) GetUpdateJob() *Job {
	if x, ok := m.GetUpdate().(*Update_UpdateJob); ok {
		return x.UpdateJob
	}
	return nil
}

func (m *Update) GetDeleteNode() string {
	if x, ok := m.GetUpdate().(*Update_DeleteNode); ok {
		return x.DeleteNode
	}
	return ""
}

func (m *Update) GetDeleteTask() string {
	if x, ok := m.GetUpdate().(*Update_DeleteTask); ok {
		return x.DeleteTask
	}
	return ""
}

func (m *Update) GetDeleteJob() string {
	if x, ok := m.GetUpdate().(*Update_DeleteJob); ok {
		return x.DeleteJob
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Update) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Update_OneofMarshaler, _Update_OneofUnmarshaler, _Update_OneofSizer, []interface{}{
		(*Update_UpdateNode)(nil),
		(*Update_UpdateTask)(nil),
		(*Update_UpdateJob)(nil),
		(*Update_DeleteNode)(nil),
		(*Update_DeleteTask)(nil),
		(*Update_DeleteJob)(nil),
	}
}

func _Update_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Update)
	// update
	switch x := m.Update.(type) {
	case *Update_UpdateNode:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.UpdateNode); err != nil {
			return err
		}
	case *Update_UpdateTask:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.UpdateTask); err != nil {
			return err
		}
	case *Update_UpdateJob:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.UpdateJob); err != nil {
			return err
		}
	case *Update_DeleteNode:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.DeleteNode)
	case *Update_DeleteTask:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.DeleteTask)
	case *Update_DeleteJob:
		b.EncodeVarint(8<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.DeleteJob)
	case nil:
	default:
		return fmt.Errorf("Update.Update has unexpected type %T", x)
	}
	return nil
}

func _Update_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Update)
	switch tag {
	case 1: // update.updateNode
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Node)
		err := b.DecodeMessage(msg)
		m.Update = &Update_UpdateNode{msg}
		return true, err
	case 2: // update.updateTask
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Task)
		err := b.DecodeMessage(msg)
		m.Update = &Update_UpdateTask{msg}
		return true, err
	case 4: // update.updateJob
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Job)
		err := b.DecodeMessage(msg)
		m.Update = &Update_UpdateJob{msg}
		return true, err
	case 5: // update.deleteNode
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Update = &Update_DeleteNode{x}
		return true, err
	case 6: // update.deleteTask
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Update = &Update_DeleteTask{x}
		return true, err
	case 8: // update.deleteJob
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Update = &Update_DeleteJob{x}
		return true, err
	default:
		return false, nil
	}
}

func _Update_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Update)
	// update
	switch x := m.Update.(type) {
	case *Update_UpdateNode:
		s := proto.Size(x.UpdateNode)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Update_UpdateTask:
		s := proto.Size(x.UpdateTask)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Update_UpdateJob:
		s := proto.Size(x.UpdateJob)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Update_DeleteNode:
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.DeleteNode)))
		n += len(x.DeleteNode)
	case *Update_DeleteTask:
		n += proto.SizeVarint(6<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.DeleteTask)))
		n += len(x.DeleteTask)
	case *Update_DeleteJob:
		n += proto.SizeVarint(8<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.DeleteJob)))
		n += len(x.DeleteJob)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*Node)(nil), "api.Node")
	proto.RegisterType((*Spec)(nil), "api.Spec")
	proto.RegisterType((*TaskStatus)(nil), "api.TaskStatus")
	proto.RegisterType((*Task)(nil), "api.Task")
	proto.RegisterType((*Job)(nil), "api.Job")
	proto.RegisterType((*Update)(nil), "api.Update")
	proto.RegisterEnum("api.NodeStatus", NodeStatus_name, NodeStatus_value)
	proto.RegisterEnum("api.TaskStatus_State", TaskStatus_State_name, TaskStatus_State_value)
}

var fileDescriptor0 = []byte{
	// 615 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xad, 0xe3, 0xef, 0x09, 0xb4, 0xd6, 0xa2, 0x82, 0x0f, 0x50, 0x8a, 0x2f, 0x54, 0x54, 0xca,
	0xa1, 0xf0, 0x07, 0x4c, 0xe3, 0xb6, 0xae, 0x82, 0x1b, 0x6d, 0x1c, 0x55, 0x70, 0xa9, 0x9c, 0x7a,
	0xa9, 0x0c, 0x8d, 0x6d, 0xd9, 0x0e, 0x28, 0x07, 0x8e, 0x48, 0xfc, 0x0d, 0x7e, 0x25, 0x12, 0x27,
	0x76, 0x76, 0x6d, 0x9c, 0x52, 0x4e, 0xf5, 0x7b, 0xf3, 0x34, 0x6f, 0x76, 0xe6, 0x35, 0x30, 0x6c,
	0xd6, 0x25, 0xab, 0x47, 0x65, 0x55, 0x34, 0x05, 0x51, 0x93, 0x32, 0xf3, 0xbe, 0x81, 0x16, 0x15,
	0x29, 0x23, 0xdb, 0x30, 0xc8, 0x52, 0x57, 0xd9, 0x57, 0x0e, 0x6c, 0xca, 0xbf, 0x08, 0x01, 0x2d,
	0x4f, 0x96, 0xcc, 0x1d, 0x08, 0x46, 0x7c, 0x0b, 0x4d, 0xe9, 0x6a, 0xad, 0xa6, 0x24, 0x2e, 0x98,
	0x69, 0x95, 0x64, 0x39, 0x4b, 0x5d, 0x9d, 0x93, 0x16, 0xed, 0x20, 0x79, 0x09, 0x46, 0xdd, 0x24,
	0xcd, 0xaa, 0x76, 0x55, 0x5e, 0xd8, 0x3e, 0xda, 0x19, 0x71, 0xaf, 0x11, 0x1a, 0xcd, 0x04, 0x4d,
	0xdb, 0xb2, 0xf7, 0x43, 0x01, 0x6d, 0x56, 0xb2, 0xeb, 0xbf, 0x7e, 0xca, 0x86, 0xdf, 0x63, 0x30,
	0x2a, 0xb6, 0x2c, 0x9a, 0x6e, 0x8a, 0x16, 0x21, 0xbf, 0x58, 0xe5, 0xe9, 0x2d, 0x13, 0xdd, 0x39,
	0x2f, 0x11, 0xd9, 0x03, 0x60, 0x79, 0x53, 0xad, 0xcb, 0x22, 0xcb, 0x9b, 0x76, 0xce, 0x0d, 0x86,
	0x3c, 0x05, 0x3b, 0xcb, 0xb9, 0x71, 0x7e, 0xcd, 0x6a, 0x31, 0xb1, 0x4a, 0x7b, 0xc2, 0xfb, 0xa5,
	0x00, 0xc4, 0x49, 0xfd, 0x59, 0x4e, 0x48, 0x0e, 0x41, 0xc7, 0x19, 0xa5, 0xf7, 0xf6, 0xd1, 0xae,
	0x78, 0x41, 0x5f, 0x1f, 0xe1, 0x1f, 0x46, 0xa5, 0x06, 0x37, 0xb1, 0x64, 0x75, 0x9d, 0xdc, 0x74,
	0x23, 0x75, 0xd0, 0xfb, 0xa9, 0x80, 0x2e, 0xa4, 0xc4, 0x04, 0x35, 0x0a, 0x2e, 0x9d, 0x2d, 0xf2,
	0x00, 0x2c, 0x7f, 0x36, 0x0b, 0x4f, 0xa3, 0x60, 0xec, 0x28, 0xe4, 0x21, 0xd8, 0x53, 0x1a, 0x4c,
	0x7d, 0x1a, 0x46, 0xa7, 0xce, 0x80, 0xd8, 0xa0, 0xd3, 0xc0, 0x1f, 0xbf, 0x77, 0x54, 0xd4, 0xcd,
	0x62, 0x9f, 0xc6, 0x58, 0xd0, 0xc8, 0x10, 0x4c, 0x3a, 0x8f, 0x22, 0x04, 0xba, 0x28, 0x9d, 0xcd,
	0xe3, 0xf1, 0xc5, 0x65, 0xe4, 0x18, 0x88, 0x8e, 0x2f, 0xde, 0x4d, 0x27, 0x41, 0x1c, 0x38, 0x26,
	0x01, 0x30, 0x4e, 0xfc, 0x70, 0xc2, 0x9b, 0x5b, 0x58, 0xa1, 0xc1, 0x79, 0x70, 0x1c, 0x73, 0x64,
	0x23, 0x3a, 0x09, 0x23, 0x7f, 0x12, 0x7e, 0x08, 0x1c, 0x20, 0x16, 0x68, 0x63, 0xee, 0xe4, 0x0c,
	0xbd, 0xef, 0x03, 0xd0, 0xf0, 0x65, 0xf7, 0x42, 0xb0, 0x0b, 0xc6, 0xa7, 0x62, 0x71, 0xc5, 0x39,
	0x79, 0x00, 0x9d, 0xa3, 0x30, 0x25, 0x4f, 0xc0, 0xcc, 0xf9, 0x29, 0x91, 0x6f, 0x0f, 0x80, 0x90,
	0x17, 0x9e, 0x81, 0x56, 0xf3, 0x63, 0x8a, 0xd5, 0x0f, 0x8f, 0x6c, 0xb1, 0x32, 0xbc, 0x2e, 0x15,
	0xf4, 0x46, 0x2a, 0x74, 0x21, 0xd8, 0xf9, 0x67, 0xa7, 0x5d, 0x2a, 0xd0, 0x37, 0x67, 0x0d, 0xf6,
	0x37, 0xa4, 0x2f, 0x47, 0xbc, 0xfd, 0x23, 0xd0, 0x59, 0x89, 0xac, 0x29, 0x43, 0xc2, 0x4a, 0x4e,
	0xca, 0x50, 0x5a, 0x9b, 0xa1, 0xbc, 0xe1, 0xeb, 0xfe, 0x9a, 0xac, 0x5d, 0x5b, 0x9e, 0xa2, 0x85,
	0xe4, 0x39, 0x0c, 0xd3, 0x2a, 0xfb, 0xc2, 0xaa, 0xab, 0x2c, 0xff, 0x58, 0xb8, 0xb0, 0xaf, 0x62,
	0x3e, 0x24, 0x15, 0x72, 0xc6, 0x7b, 0x03, 0xea, 0x79, 0xb1, 0xb8, 0xb7, 0x85, 0xee, 0x55, 0x83,
	0xff, 0xbe, 0xca, 0xfb, 0xad, 0x80, 0x31, 0x2f, 0x53, 0x3c, 0xf1, 0x21, 0xc0, 0x4a, 0x7c, 0x61,
	0xd2, 0x45, 0x87, 0x4e, 0x8f, 0xc4, 0xd9, 0x16, 0xdd, 0x28, 0xf7, 0x62, 0x5c, 0xc0, 0x9d, 0xe6,
	0x48, 0xf4, 0x62, 0x71, 0x99, 0x03, 0xb0, 0x25, 0xe2, 0x03, 0xb6, 0xeb, 0xb5, 0x84, 0x96, 0x63,
	0x2e, 0xed, 0x8b, 0x64, 0x1f, 0x20, 0x65, 0xb7, 0xac, 0x9d, 0x01, 0x17, 0x6d, 0x63, 0xaf, 0x9e,
	0xeb, 0x15, 0xc2, 0xd8, 0xb8, 0xab, 0x10, 0x6e, 0x7b, 0x60, 0x4b, 0x84, 0x6e, 0x56, 0x2b, 0xe8,
	0xa9, 0xb7, 0x16, 0x18, 0xd2, 0xf0, 0xd5, 0x0b, 0x80, 0xfe, 0xbf, 0xba, 0x0f, 0xef, 0x96, 0x48,
	0x17, 0xa6, 0x53, 0x59, 0x18, 0xe2, 0xd7, 0xe6, 0xf5, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1c,
	0x01, 0x91, 0xf9, 0x7c, 0x04, 0x00, 0x00,
}

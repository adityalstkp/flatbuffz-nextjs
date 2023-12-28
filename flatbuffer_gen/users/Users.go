// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package users

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Users struct {
	_tab flatbuffers.Table
}

func GetRootAsUsers(buf []byte, offset flatbuffers.UOffsetT) *Users {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Users{}
	x.Init(buf, n+offset)
	return x
}

func FinishUsersBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.Finish(offset)
}

func GetSizePrefixedRootAsUsers(buf []byte, offset flatbuffers.UOffsetT) *Users {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Users{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func FinishSizePrefixedUsersBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.FinishSizePrefixed(offset)
}

func (rcv *Users) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Users) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Users) List(obj *User, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *Users) ListLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func UsersStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func UsersAddList(builder *flatbuffers.Builder, list flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(list), 0)
}
func UsersStartListVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func UsersEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
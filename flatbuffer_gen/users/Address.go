// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package users

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Address struct {
	_tab flatbuffers.Table
}

func GetRootAsAddress(buf []byte, offset flatbuffers.UOffsetT) *Address {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Address{}
	x.Init(buf, n+offset)
	return x
}

func FinishAddressBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.Finish(offset)
}

func GetSizePrefixedRootAsAddress(buf []byte, offset flatbuffers.UOffsetT) *Address {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Address{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func FinishSizePrefixedAddressBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.FinishSizePrefixed(offset)
}

func (rcv *Address) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Address) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Address) Street() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Address) City() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Address) Geo(obj *Geo) *Geo {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(Geo)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func AddressStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func AddressAddStreet(builder *flatbuffers.Builder, street flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(street), 0)
}
func AddressAddCity(builder *flatbuffers.Builder, city flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(city), 0)
}
func AddressAddGeo(builder *flatbuffers.Builder, geo flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(geo), 0)
}
func AddressEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}

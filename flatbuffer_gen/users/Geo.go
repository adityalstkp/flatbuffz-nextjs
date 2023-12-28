// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package users

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Geo struct {
	_tab flatbuffers.Table
}

func GetRootAsGeo(buf []byte, offset flatbuffers.UOffsetT) *Geo {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Geo{}
	x.Init(buf, n+offset)
	return x
}

func FinishGeoBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.Finish(offset)
}

func GetSizePrefixedRootAsGeo(buf []byte, offset flatbuffers.UOffsetT) *Geo {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Geo{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func FinishSizePrefixedGeoBuffer(builder *flatbuffers.Builder, offset flatbuffers.UOffsetT) {
	builder.FinishSizePrefixed(offset)
}

func (rcv *Geo) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Geo) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Geo) Lat() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Geo) Long() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func GeoStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func GeoAddLat(builder *flatbuffers.Builder, lat flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(lat), 0)
}
func GeoAddLong(builder *flatbuffers.Builder, long flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(long), 0)
}
func GeoEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}

package goenet

/*
#cgo CFLAGS: -Isrc/include

#include "enet/enet.h"
*/
import "C"

import (
	"unsafe"
)

type ENetPacket C.ENetPacket

func slice_to_cpointer(in []byte) unsafe.Pointer {
	if len(in) == 0 {
		return unsafe.Pointer(nil)
	}
	return unsafe.Pointer(&in[0])
}

// NewPacket create a ENetPacket.
// data can be nil, when it not nil, make sure len(data) >= dataLength!
func NewPacket(data []byte, dataLength int, flags ENetPacketFlag) *ENetPacket {
	return (*ENetPacket)(C.enet_packet_create(slice_to_cpointer(data), C.size_t(dataLength), C.enet_uint32(flags)))
}

func (p *ENetPacket) Data() []byte {
	return (*[1 << 30]byte)(unsafe.Pointer(p.data))[0:p.DataLength()]
}

func (p *ENetPacket) DataLength() int {
	return int(p.dataLength)
}

func (p *ENetPacket) Destroy() {
	C.enet_packet_destroy((*C.ENetPacket)(p))
}

func (p *ENetPacket) Resize(dataLength int) int {
	return int(C.enet_packet_resize((*C.ENetPacket)(p), C.size_t(dataLength)))
}

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

// Currently this method requires that the data be written in byte form.
// Ideally it would be nice to write arbitrary data such as one can do in the C ENet library.
func NewPacket(data []byte, dataLength int, flags ENetPacketFlag) *ENetPacket {
	return (*ENetPacket)(C.enet_packet_create(unsafe.Pointer(&data[0]), C.size_t(dataLength), C.enet_uint32(flags)))
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

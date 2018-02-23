package goenet

/*
#cgo CFLAGS: -Isrc/include

#include "enet/enet.h"
*/
import "C"

import (
    "unsafe"
)

type ENetPeer C.ENetPeer

// Private

func (p *ENetPeer) ConnectID() int {
    return int(p.connectID)
}

// SetData set a user data.
//
// https://golang.org/cmd/cgo/
// The C code must preserve this property: it must not store any Go pointers in Go memory, even temporarily.
//
func (p *ENetPeer) SetData(data uint) {
    p.data = unsafe.Pointer(uintptr(data))
}

// Data returns the user data.
func (p *ENetPeer) Data() uint {
    return uint(uintptr(p.data))
}

// Public

func (p *ENetPeer) Address() ENetAddress {
    return (ENetAddress)(p.address)
}

func (p *ENetPeer) Send(channelID int, packet *ENetPacket) int {
    return int(C.enet_peer_send((*C.ENetPeer)(p), C.enet_uint8(channelID), (*C.ENetPacket)(packet)))
}

func (p *ENetPeer) Receive(channelID *C.enet_uint8) *ENetPacket {
    return (*ENetPacket)(C.enet_peer_receive((*C.ENetPeer)(p), channelID))
}

func (p *ENetPeer) Ping() {
    C.enet_peer_ping((*C.ENetPeer)(p))
}

func (p *ENetPeer) PingInterval(pingInterval int) {
    C.enet_peer_ping_interval((*C.ENetPeer)(p), C.enet_uint32(pingInterval))
}

func (p *ENetPeer) Timeout(timeoutLimit, timeoutMinimum, timeoutMaximum int) {
    C.enet_peer_timeout((*C.ENetPeer)(p), C.enet_uint32(timeoutLimit), C.enet_uint32(timeoutMinimum), C.enet_uint32(timeoutMaximum))
}

func (p *ENetPeer) Reset() {
    C.enet_peer_reset((*C.ENetPeer)(p))
}

func (p *ENetPeer) Disconnect(data int) {
    C.enet_peer_disconnect((*C.ENetPeer)(p), C.enet_uint32(data))
}

func (p *ENetPeer) DisconnectNow(data int) {
    C.enet_peer_disconnect_now((*C.ENetPeer)(p), C.enet_uint32(data))
}

func (p *ENetPeer) DisconnectLater(data int) {
    C.enet_peer_disconnect_later((*C.ENetPeer)(p), C.enet_uint32(data))
}

func (p *ENetPeer) ThrottleConfigure(interval, acceleration, deceleration int) {
    C.enet_peer_throttle_configure((*C.ENetPeer)(p), C.enet_uint32(interval), C.enet_uint32(acceleration), C.enet_uint32(deceleration))
}

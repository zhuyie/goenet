package goenet

/*
#cgo CFLAGS: -Isrc/include

#include "enet/enet.h"
*/
import "C"

import (
    "unsafe"
)

type ENetAddress C.ENetAddress

func (a *ENetAddress) SetHost(h uint) {
    a.host = C.enet_uint32(h)
}

func (a *ENetAddress) SetPort(port uint) {
    a.port = C.enet_uint16(port)
}

func (a *ENetAddress) SetHostName(hostName string) int {
    hName := C.CString(hostName)
    defer C.free(unsafe.Pointer(hName))
    return int(C.enet_address_set_host((*C.ENetAddress)(a), hName))
}

func (a *ENetAddress) HostIp() string {
    hostName := C.malloc(16)
    defer C.free(hostName)
    if C.enet_address_get_host_ip((*C.ENetAddress)(a), (*C.char)(hostName), 16) != 0 {
        return ""
    }
    return C.GoString((*C.char)(hostName))
}

func (a *ENetAddress) Host() string {
    hostName := C.malloc(50)
    defer C.free(hostName)
    if C.enet_address_get_host((*C.ENetAddress)(a), (*C.char)(hostName), 50) != 0 {
        return ""
    }
    return C.GoString((*C.char)(hostName))
}

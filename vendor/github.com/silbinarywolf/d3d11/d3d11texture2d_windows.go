package d3d11

import (
	"syscall"
	"unsafe"
)

// ID3D11Texture2D
// 6f15aaf2-d208-4e89-9ab4-489535d34f9c
type ID3D11Texture2D struct {
	vtbl *id3d11Texture2DVtbl
}

var id3d11Texture2D_UUID = guid{0x6f15aaf2, 0xd208, 0x4e89, [8]byte{0x9a, 0xb4, 0x48, 0x95, 0x35, 0xd3, 0x4f, 0x9c}}

type id3d11Texture2DVtbl struct {
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr

	GetDevice               uintptr
	GetPrivateData          uintptr
	SetPrivateData          uintptr
	SetPrivateDataInterface uintptr
	GetType                 uintptr
	SetEvictionPriority     uintptr
	GetEvictionPriority     uintptr
	GetDesc                 uintptr
}

// Release has to be called when finished using the object to free its
// associated resources.
func (obj *ID3D11Texture2D) Release() uint32 {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.Release,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0,
	)
	return uint32(ret)
}

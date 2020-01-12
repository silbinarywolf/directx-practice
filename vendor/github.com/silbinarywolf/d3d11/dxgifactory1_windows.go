package d3d11

import (
	"syscall"
	"unsafe"
)

// IDXGIFactory1
// 770aae78-f26f-4dba-a829-253c83d1b387
type IDXGIFactory1 struct {
	vtbl *idxgiFactory1Vtbl
}

var idxgiFactory1_UUID = guid{0x770aae78, 0xf26f, 0x4dba, [8]byte{0xa8, 0x29, 0x25, 0x3c, 0x83, 0xd1, 0xb3, 0x87}}

type idxgiFactory1Vtbl struct {
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr

	SetPrivateData          uintptr
	SetPrivateDataInterface uintptr
	GetPrivateData          uintptr
	GetParent               uintptr
	EnumAdapters            uintptr
	MakeWindowAssociation   uintptr
	GetWindowAssociation    uintptr
	CreateSwapChain         uintptr
	CreateSoftwareAdapter   uintptr
	EnumAdapters1           uintptr
	IsCurrent               uintptr
}

func (obj *IDXGIFactory1) CreateSwapChain(device *Device, desc DXGI_SWAP_CHAIN_DESC) (*IDXGISwapChain, Error) {
	var r *IDXGISwapChain
	ret, _, _ := syscall.Syscall6(
		obj.vtbl.CreateSwapChain,
		4,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(device)),
		uintptr(unsafe.Pointer(&desc)),
		uintptr(unsafe.Pointer(&r)),
		0,
		0,
	)
	return r, toErr(ret)
}

func (obj *IDXGIFactory1) MakeWindowAssociation(window HWND, flags DXGI_MWA) Error {
	ret, _, _ := syscall.Syscall6(
		obj.vtbl.MakeWindowAssociation,
		3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(window),
		uintptr(flags),
		0,
		0,
		0,
	)
	return toErr(ret)
}

// Release has to be called when finished using the object to free its
// associated resources.
func (obj *IDXGIFactory1) Release() uint32 {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.Release,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0,
	)
	return uint32(ret)
}

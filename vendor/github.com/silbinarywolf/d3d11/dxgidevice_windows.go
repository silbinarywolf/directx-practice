package d3d11

import (
	"syscall"
	"unsafe"
)

// IDXGIDevice
// 54ec77fa-1377-44e6-8c32-88fd5f44c84c
type IDXGIDevice struct {
	vtbl *idxgiDeviceVtbl
}

var idxgiDevice_UUID = guid{0x54ec77fa, 0x1377, 0x44e6, [8]byte{0x8c, 0x32, 0x88, 0xfd, 0x5f, 0x44, 0xc8, 0x4c}}

type idxgiDeviceVtbl struct {
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr

	SetPrivateData          uintptr
	SetPrivateDataInterface uintptr
	GetPrivateData          uintptr
	GetParent               uintptr
	GetAdapter              uintptr
	CreateSurface           uintptr
	QueryResourceResidency  uintptr
	SetGPUThreadPriority    uintptr
	GetGPUThreadPriority    uintptr
}

func (obj *IDXGIDevice) GetAdapter() (*IDXGIAdapter, Error) {
	var r *IDXGIAdapter
	ret, _, _ := syscall.Syscall(
		obj.vtbl.GetAdapter,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&r)),
		0,
	)
	return r, toErr(ret)
}

// Release has to be called when finished using the object to free its
// associated resources.
func (obj *IDXGIDevice) Release() uint32 {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.Release,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0,
	)
	return uint32(ret)
}

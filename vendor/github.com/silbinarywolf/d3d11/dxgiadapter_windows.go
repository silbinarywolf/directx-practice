package d3d11

import (
	"syscall"
	"unsafe"
)

// IDXGIAdapter
// 2411e7e1-12ac-4ccf-bd14-9798e8534dc0
type IDXGIAdapter struct {
	vtbl *idxgiAdapterVtbl
}

var idxgiAdapter_UUID = guid{0x2411e7e1, 0x12ac, 0x4ccf, [8]byte{0xbd, 0x14, 0x97, 0x98, 0xe8, 0x53, 0x4d, 0xc0}}

type idxgiAdapterVtbl struct {
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr

	SetPrivateData          uintptr
	SetPrivateDataInterface uintptr
	GetPrivateData          uintptr
	GetParent               uintptr
	EnumOutputs             uintptr
	GetDesc                 uintptr
	CheckInterfaceSupport   uintptr
}

func (obj *IDXGIAdapter) GetParent() (*IDXGIFactory1, Error) {
	var r *IDXGIFactory1
	ret, _, _ := syscall.Syscall(
		obj.vtbl.GetParent,
		3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&idxgiFactory1_UUID)),
		uintptr(unsafe.Pointer(&r)),
	)
	return r, toErr(ret)
}

// Release has to be called when finished using the object to free its
// associated resources.
func (obj *IDXGIAdapter) Release() uint32 {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.Release,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0,
	)
	return uint32(ret)
}

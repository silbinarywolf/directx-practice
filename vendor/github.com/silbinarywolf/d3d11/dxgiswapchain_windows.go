package d3d11

import (
	"syscall"
	"unsafe"
)

// IDXGISwapChain
// 310d36a0-d2e7-4c0a-aa04-6a9d23b8886a
type IDXGISwapChain struct {
	vtbl *idxgiSwapChainVtbl
}

var idxgiSwapChain_UUID = guid{0x310d36a0, 0xd2e7, 0x4c0a, [8]byte{0xaa, 0x04, 0x6a, 0x9d, 0x23, 0xb8, 0x88, 0x6a}}

type idxgiSwapChainVtbl struct {
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr

	SetPrivateData          uintptr
	SetPrivateDataInterface uintptr
	GetPrivateData          uintptr
	GetParent               uintptr
	GetDevice               uintptr
	Present                 uintptr
	GetBuffer               uintptr
	SetFullscreenState      uintptr
	GetFullscreenState      uintptr
	GetDesc                 uintptr
	ResizeBuffers           uintptr
	ResizeTarget            uintptr
	GetContainingOutput     uintptr
	GetFrameStatistics      uintptr
	GetLastPresentCount     uintptr
}

// ID3D11Texture2D* pBackBuffer = nullptr;
// hr = g_pSwapChain->GetBuffer( 0, __uuidof( ID3D11Texture2D ), reinterpret_cast<void**>( &pBackBuffer ) );

func (obj *IDXGISwapChain) GetBuffer(buffer uint32) (*ID3D11Texture2D, Error) {
	var r *ID3D11Texture2D
	ret, _, _ := syscall.Syscall6(
		obj.vtbl.GetBuffer,
		4,
		uintptr(unsafe.Pointer(obj)),
		uintptr(buffer),
		uintptr(unsafe.Pointer(&id3d11Texture2D_UUID)),
		uintptr(unsafe.Pointer(&r)),
		0,
		0,
	)
	return r, toErr(ret)
}

func (obj *IDXGISwapChain) Present(syncInterval uint32, flags uint32) {
	syscall.Syscall(
		obj.vtbl.Present,
		3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(syncInterval),
		uintptr(flags),
	)
}

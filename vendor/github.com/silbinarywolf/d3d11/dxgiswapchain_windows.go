package d3d11

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

package d3d11

// 54ec77fa-1377-44e6-8c32-88fd5f44c84c
var idxgiDevice_UUID = guid{0x54ec77fa, 0x1377, 0x44e6, [8]byte{0x8c, 0x32, 0x88, 0xfd, 0x5f, 0x44, 0xc8, 0x4c}}

// IDXGIDevice
type IDXGIDevice struct {
	vtbl *idxgiDeviceVtbl
}

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

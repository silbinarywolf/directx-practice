package d3d11

// ID3D11Resource
// dc8e63f3-d12b-4952-b47b-5e45026a862d
type ID3D11Resource struct {
	vtbl *id3d11ResourceVtbl
}

type id3d11ResourceVtbl struct {
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

func (*ID3D11Resource) uuid() guid {
	// dc8e63f3-d12b-4952-b47b-5e45026a862d
	return guid{0xdc8e63f3, 0xd12b, 0x4952, [8]byte{0xb4, 0x7b, 0x5e, 0x45, 0x02, 0x6a, 0x86, 0x2d}}
}

package d3d11

// ID3D11DepthStencilView
// 9fdac92a-1876-48c3-afad-25b94f84a9b6
type ID3D11DepthStencilView struct {
	vtbl *id3d11ResourceVtbl
}

var id3d11DepthStencilView_UUID = guid{0x9fdac92a, 0x1876, 0x48c3, [8]byte{0xaf, 0xad, 0x25, 0xb9, 0x4f, 0x84, 0xa9, 0xb6}}

type id3d11DepthStencilViewVtbl struct {
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr

	GetDevice               uintptr
	GetPrivateData          uintptr
	SetPrivateData          uintptr
	SetPrivateDataInterface uintptr
	GetResource             uintptr
	GetDesc                 uintptr
}

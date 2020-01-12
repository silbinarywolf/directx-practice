package d3d11

import (
	"syscall"
	"unsafe"
)

type Device struct {
	vtbl *deviceVtbl
}

type deviceVtbl struct {
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr

	CreateBuffer                         uintptr
	CreateTexture1D                      uintptr
	CreateTexture2D                      uintptr
	CreateTexture3D                      uintptr
	CreateShaderResourceView             uintptr
	CreateUnorderedAccessView            uintptr
	CreateRenderTargetView               uintptr
	CreateDepthStencilView               uintptr
	CreateInputLayout                    uintptr
	CreateVertexShader                   uintptr
	CreateGeometryShader                 uintptr
	CreateGeometryShaderWithStreamOutput uintptr
	CreatePixelShader                    uintptr
	CreateHullShader                     uintptr
	CreateDomainShader                   uintptr
	CreateComputeShader                  uintptr
	CreateClassLinkage                   uintptr
	CreateBlendState                     uintptr
	CreateDepthStencilState              uintptr
	CreateRasterizerState                uintptr
	CreateSamplerState                   uintptr
	CreateQuery                          uintptr
	CreatePredicate                      uintptr
	CreateCounter                        uintptr
	CreateDeferredContext                uintptr
	OpenSharedResource                   uintptr
	CheckFormatSupport                   uintptr
	CheckMultisampleQualityLevels        uintptr
	CheckCounterInfo                     uintptr
	CheckCounter                         uintptr
	CheckFeatureSupport                  uintptr
	GetPrivateData                       uintptr
	SetPrivateData                       uintptr
	SetPrivateDataInterface              uintptr
	GetFeatureLevel                      uintptr
	GetDeviceRemovedReason               uintptr
	GetImmediateContext                  uintptr
	SetExceptionMode                     uintptr
	GetExceptionMode                     uintptr
}

func (obj *Device) QueryInterfaceIDXGIDevice() (*IDXGIDevice, Error) {
	var r *IDXGIDevice
	ret, _, _ := syscall.Syscall(
		obj.vtbl.QueryInterface,
		3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&idxgiDevice_UUID)),
		uintptr(unsafe.Pointer(&r)),
	)
	return r, toErr(ret)
}

func (obj *Device) CreateRenderTargetView(backBuffer *ID3D11Texture2D, desc *RENDER_TARGET_VIEW_DESC) (*ID3D11RenderTargetView, Error) {
	var r *ID3D11RenderTargetView
	ret, _, _ := syscall.Syscall6(
		obj.vtbl.CreateRenderTargetView,
		4,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(backBuffer)),
		uintptr(unsafe.Pointer(desc)),
		uintptr(unsafe.Pointer(&id3d11RenderTargetView_UUID)),
		uintptr(unsafe.Pointer(&r)),
		0,
	)
	return r, toErr(ret)
}

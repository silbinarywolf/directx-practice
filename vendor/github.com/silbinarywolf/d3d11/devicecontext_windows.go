package d3d11

import (
	"syscall"
	"unsafe"
)

type DeviceContext struct {
	vtbl *deviceContextVtbl
}

type deviceContextVtbl struct {
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr

	GetDevice               uintptr
	GetPrivateData          uintptr
	SetPrivateData          uintptr
	SetPrivateDataInterface uintptr
	VSSetConstantBuffers    uintptr
	PSSetShaderResources    uintptr
	PSSetShader             uintptr
	PSSetSamplers           uintptr
	VSSetShader             uintptr
	DrawIndexed             uintptr
	Draw                    uintptr
	Map                     uintptr
	Unmap                   uintptr
	PSSetConstantBuffers    uintptr
	IASetInputLayout        uintptr
	IASetVertexBuffers      uintptr
	IASetIndexBuffer        uintptr
	DrawIndexedInstanced    uintptr
	DrawInstanced           uintptr
	GSSetConstantBuffers    uintptr
	GSSetShader             uintptr
	IASetPrimitiveTopology  uintptr
	VSSetShaderResources    uintptr
	VSSetSamplers           uintptr
	Begin                   uintptr
	End                     uintptr
	GetData                 uintptr
	SetPredication          uintptr
	GSSetShaderResources    uintptr
	GSSetSamplers           uintptr

	OMSetRenderTargets                        uintptr
	OMSetRenderTargetsAndUnorderedAccessViews uintptr
	OMSetBlendState                           uintptr
	OMSetDepthStencilState                    uintptr
	SOSetTargets                              uintptr

	DrawAuto                      uintptr
	DrawIndexedInstancedIndirect  uintptr
	DrawInstancedIndirect         uintptr
	Dispatch                      uintptr
	DispatchIndirect              uintptr
	RSSetState                    uintptr
	RSSetViewports                uintptr
	RSSetScissorRects             uintptr
	CopySubresourceRegion         uintptr
	CopyResource                  uintptr
	UpdateSubresource             uintptr
	CopyStructureCount            uintptr
	ClearRenderTargetView         uintptr
	ClearUnorderedAccessViewUint  uintptr
	ClearUnorderedAccessViewFloat uintptr
	ClearDepthStencilView         uintptr
	GenerateMips                  uintptr
	SetResourceMinLOD             uintptr
	GetResourceMinLOD             uintptr
	ResolveSubresource            uintptr
	ExecuteCommandList            uintptr
	HSSetShaderResources          uintptr
	HSSetShader                   uintptr
	HSSetSamplers                 uintptr
	HSSetConstantBuffers          uintptr
	DSSetShaderResources          uintptr
	DSSetShader                   uintptr
	DSSetSamplers                 uintptr
	DSSetConstantBuffers          uintptr
	CSSetShaderResources          uintptr
	CSSetUnorderedAccessViews     uintptr
	CSSetShader                   uintptr
	CSSetSamplers                 uintptr
	CSSetConstantBuffers          uintptr
	VSGetConstantBuffers          uintptr
	PSGetShaderResources          uintptr
	PSGetShader                   uintptr
	PSGetSamplers                 uintptr
	VSGetShader                   uintptr
	PSGetConstantBuffers          uintptr
	IAGetInputLayout              uintptr
	IAGetVertexBuffers            uintptr
	IAGetIndexBuffer              uintptr
	GSGetConstantBuffers          uintptr
	GSGetShader                   uintptr
	IAGetPrimitiveTopology        uintptr
	VSGetShaderResources          uintptr
	VSGetSamplers                 uintptr
	GetPredication                uintptr

	GSGetShaderResources uintptr
	GSGetSamplers        uintptr

	OMGetRenderTargets                        uintptr
	OMGetRenderTargetsAndUnorderedAccessViews uintptr

	OMGetBlendState           uintptr
	OMGetDepthStencilState    uintptr
	SOGetTargets              uintptr
	RSGetState                uintptr
	RSGetViewports            uintptr
	RSGetScissorRects         uintptr
	HSGetShaderResources      uintptr
	HSGetShader               uintptr
	HSGetSamplers             uintptr
	HSGetConstantBuffers      uintptr
	DSGetShaderResources      uintptr
	DSGetShader               uintptr
	DSGetSamplers             uintptr
	DSGetConstantBuffers      uintptr
	CSGetShaderResources      uintptr
	CSGetUnorderedAccessViews uintptr
	CSGetShader               uintptr
	CSGetSamplers             uintptr
	CSGetConstantBuffers      uintptr
	ClearState                uintptr
	Flush                     uintptr
	GetType                   uintptr
	GetContextFlags           uintptr
	FinishCommandList         uintptr
}

func (obj *DeviceContext) OMSetRenderTargets(numViews uint32, renderTargetViews **ID3D11RenderTargetView, depthStencilView *ID3D11DepthStencilView) {
	syscall.Syscall6(
		obj.vtbl.OMSetRenderTargets,
		4,
		uintptr(unsafe.Pointer(obj)),
		uintptr(numViews),
		uintptr(unsafe.Pointer(renderTargetViews)),
		uintptr(unsafe.Pointer(depthStencilView)),
		0,
		0,
	)
}

func (obj *DeviceContext) RSSetViewports(numViewports uint32, viewports **VIEWPORT) {
	syscall.Syscall(
		obj.vtbl.RSSetViewports,
		3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(numViewports),
		uintptr(unsafe.Pointer(viewports)),
	)
}

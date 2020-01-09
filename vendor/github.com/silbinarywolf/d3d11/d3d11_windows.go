package d3d11

import (
	"syscall"
	"unsafe"
)

var (
	d3d11 = syscall.NewLazyDLL("d3d11.dll")

	createDevice = d3d11.NewProc("D3D11CreateDevice")
)

// CreateDevice
// https://docs.microsoft.com/en-us/windows/win32/api/d3d11/nf-d3d11-d3d11createdevice
func CreateDevice(
	pAdapter IDXGIAdapter,
	driverType DRIVER_TYPE,
	software HMODULE,
	flags uint32,
	featureLevels []FEATURE_LEVEL,
	sdkVersion uint32,
	immediateContext *DeviceContext,
) (*Device, FEATURE_LEVEL, error) {
	//CreateDevice( nullptr, g_driverType, nullptr, createDeviceFlags, featureLevels, numFeatureLevels,
	//                                D3D11_SDK_VERSION, &g_pd3dDevice, &g_featureLevel, &g_pImmediateContext );
	/*
		HRESULT D3D11CreateDevice(
		  IDXGIAdapter            *pAdapter,
		  D3D_DRIVER_TYPE         DriverType,
		  HMODULE                 Software,
		  UINT                    Flags,
		  const D3D_FEATURE_LEVEL *pFeatureLevels,
		  UINT                    FeatureLevels,
		  UINT                    SDKVersion,
		  ID3D11Device            **ppDevice,
		  D3D_FEATURE_LEVEL       *pFeatureLevel,
		  ID3D11DeviceContext     **ppImmediateContext
		);
	*/
	var device *Device
	var featureLevel FEATURE_LEVEL
	ret, _, _ := createDevice.Call(
		uintptr(pAdapter),
		uintptr(driverType),
		uintptr(software),
		uintptr(flags),
		uintptr(unsafe.Pointer(&featureLevels)),
		uintptr(uint32(len(featureLevels))),
		uintptr(sdkVersion),
		uintptr(unsafe.Pointer(&device)),
		uintptr(unsafe.Pointer(&featureLevel)),
		uintptr(unsafe.Pointer(&immediateContext)),
	)
	return device, featureLevel, toErr(ret)
}

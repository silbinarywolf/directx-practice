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
	pAdapter *IDXGIAdapter,
	driverType DRIVER_TYPE,
	software HMODULE,
	flags uint32,
	featureLevels []FEATURE_LEVEL,
	sdkVersion uint32,
) (*Device, FEATURE_LEVEL, *DeviceContext, Error) {
	var device *Device
	var immediateContext *DeviceContext
	var featureLevel FEATURE_LEVEL
	ret, _, _ := createDevice.Call(
		uintptr(unsafe.Pointer(pAdapter)),
		uintptr(driverType),
		uintptr(software),
		uintptr(flags),
		uintptr(unsafe.Pointer(&featureLevels[0])),
		uintptr(uint32(len(featureLevels))),
		uintptr(sdkVersion),
		uintptr(unsafe.Pointer(&device)),
		uintptr(unsafe.Pointer(&featureLevel)),
		uintptr(unsafe.Pointer(&immediateContext)),
	)
	return device, featureLevel, immediateContext, toErr(ret)
}

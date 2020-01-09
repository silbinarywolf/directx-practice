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
) (*Device, FEATURE_LEVEL, Error) {
	// NOTE(Jae): 2020-01-09
	// Copy slice to stack array temporarily
	// so this call works
	var featureLevelsArr [16]FEATURE_LEVEL
	for i, featureLevel := range featureLevels {
		featureLevelsArr[i] = featureLevel
	}
	var device *Device
	var featureLevel FEATURE_LEVEL
	ret, _, _ := createDevice.Call(
		uintptr(pAdapter),
		uintptr(driverType),
		uintptr(software),
		uintptr(flags),
		uintptr(unsafe.Pointer(&featureLevelsArr)),
		uintptr(uint32(len(featureLevels))),
		uintptr(sdkVersion),
		uintptr(unsafe.Pointer(&device)),
		uintptr(unsafe.Pointer(&featureLevel)),
		uintptr(unsafe.Pointer(&immediateContext)),
	)
	return device, featureLevel, toErr(ret)
}

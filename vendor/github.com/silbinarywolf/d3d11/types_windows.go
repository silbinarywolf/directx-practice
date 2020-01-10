package d3d11

type (
	handle uintptr
	// Windows
	HMODULE handle

	// D3D
	IDXGIAdapter  handle
	DeviceContext handle
	IDXGIFactory1 handle
)

// https://docs.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
type guid struct {
	data1 uint32
	data2 uint16
	data3 uint16
	data4 [8]byte
}

const (
	SDK_VERSION = 7
)

type DRIVER_TYPE handle

const (
	DRIVER_TYPE_UNKNOWN   DRIVER_TYPE = 0
	DRIVER_TYPE_HARDWARE  DRIVER_TYPE = DRIVER_TYPE_UNKNOWN + 1
	DRIVER_TYPE_REFERENCE DRIVER_TYPE = DRIVER_TYPE_HARDWARE + 1
	DRIVER_TYPE_NULL      DRIVER_TYPE = DRIVER_TYPE_REFERENCE + 1
	DRIVER_TYPE_SOFTWARE  DRIVER_TYPE = DRIVER_TYPE_NULL + 1
	DRIVER_TYPE_WARP      DRIVER_TYPE = DRIVER_TYPE_SOFTWARE + 1
)

type FEATURE_LEVEL handle

const (
	FEATURE_LEVEL_9_1  FEATURE_LEVEL = 0x9100
	FEATURE_LEVEL_9_2  FEATURE_LEVEL = 0x9200
	FEATURE_LEVEL_9_3  FEATURE_LEVEL = 0x9300
	FEATURE_LEVEL_10_0 FEATURE_LEVEL = 0xa000
	FEATURE_LEVEL_10_1 FEATURE_LEVEL = 0xa100
	FEATURE_LEVEL_11_0 FEATURE_LEVEL = 0xb000
	FEATURE_LEVEL_11_1 FEATURE_LEVEL = 0xb100
	//FEATURE_LEVEL_12_0 FEATURE_LEVEL = 0xc000
	//FEATURE_LEVEL_12_1 FEATURE_LEVEL = 0xc100
)

const (
	CREATE_DEVICE_SINGLETHREADED                           uint32 = 0x1
	CREATE_DEVICE_DEBUG                                    uint32 = 0x2
	CREATE_DEVICE_SWITCH_TO_REF                            uint32 = 0x4
	CREATE_DEVICE_PREVENT_INTERNAL_THREADING_OPTIMIZATIONS uint32 = 0x8
	CREATE_DEVICE_BGRA_SUPPORT                             uint32 = 0x20
)

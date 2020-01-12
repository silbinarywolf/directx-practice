package d3d11

const (
	SIMULTANEOUS_RENDER_TARGET_COUNT = 8
)

const (
	// S_OK indicates that no error occurred.
	S_OK = 0

	// ERR_CONFLICTINGRENDERSTATE indicates that the currently set render
	// states cannot be used together.
	ERR_CONFLICTINGRENDERSTATE = -2005530591

	// ERR_CONFLICTINGTEXTUREFILTER indicates that the current texture filters
	// cannot be used together.
	ERR_CONFLICTINGTEXTUREFILTER = -2005530594

	// ERR_CONFLICTINGTEXTUREPALETTE indicates that the current textures cannot
	// be used simultaneously.
	ERR_CONFLICTINGTEXTUREPALETTE = -2005530586

	// ERR_DEVICEHUNG indicates that the device that returned this code caused
	// the hardware adapter to be reset by the OS. Most applications should
	// destroy the device and quit. Applications that must continue should
	// destroy all video memory objects (surfaces, textures, state blocks etc)
	// and call Reset() to put the device in a default state. If the
	// application then continues rendering in the same way, the device will
	// return to this state.
	// Applies to Direct3D 9Ex only.
	ERR_DEVICEHUNG = -2005530508

	// ERR_DEVICELOST indicates that the device has been lost but cannot be
	// reset at this time. Therefore, rendering is not possible. A Direct3D
	// device object other than the one that returned this code caused the
	// hardware adapter to be reset by the OS. Delete all video memory objects
	// (surfaces, textures, state blocks) and call Reset() to return the device
	// to a default state. If the application continues rendering without a
	// reset, the rendering calls will succeed.
	ERR_DEVICELOST = -2005530520

	// ERR_DEVICENOTRESET indicates that the device has been lost but can be
	// reset at this time.
	ERR_DEVICENOTRESET = -2005530519

	// ERR_DEVICEREMOVED indicates that he hardware adapter has been removed.
	// Application must destroy the device, do enumeration of adapters and
	// create another Direct3D device. If application continues rendering
	// without calling Reset, the rendering calls will succeed.
	// Applies to Direct3D 9Ex only.
	ERR_DEVICEREMOVED = -2005530512

	// ERR_DRIVERINTERNALERROR Internal driver error. Applications should
	// destroy and recreate the device when receiving this error.
	ERR_DRIVERINTERNALERROR = -2005530585

	// ERR_DRIVERINVALIDCALL is not used.
	ERR_DRIVERINVALIDCALL = -2005530515

	// ERR_INVALIDCALL indicates that the method call is invalid. For example,
	// a method's parameter may not be an invalid pointer.
	ERR_INVALIDCALL = -2005530516

	// ERR_INVALIDDEVICE indicates that the requested device type is not valid.
	ERR_INVALIDDEVICE = -2005530517

	// ERR_MOREDATA indicates that there is more data available than the
	// specified buffer size can hold.
	ERR_MOREDATA = -2005530521

	// ERR_NOTAVAILABLE indicates that this device does not support the queried
	// technique.
	ERR_NOTAVAILABLE = -2005530518

	// ERR_NOTFOUND indicates that the requested item was not found.
	ERR_NOTFOUND = -2005530522

	// _OK indicates that no error occurred.
	OK = S_OK

	// ERR_OUTOFVIDEOMEMORY indicates that Direct3D does not have enough
	// display memory to perform the operation. The device is using more
	// resources in a single scene than can fit simultaneously into video
	// memory. Present, PresentEx, or CheckDeviceState can return this error.
	// Recovery is similar to ERR_DEVICEHUNG, though the application may want
	// to reduce its per-frame memory usage as well to avoid having the error
	// recur.
	ERR_OUTOFVIDEOMEMORY = -2005532292

	// ERR_TOOMANYOPERATIONS indicates that the application is requesting more
	// texture-filtering operations than the device supports.
	ERR_TOOMANYOPERATIONS = -2005530595

	// ERR_UNSUPPORTEDALPHAARG indicates that the device does not support a
	// specified texture-blending argument for the alpha channel.
	ERR_UNSUPPORTEDALPHAARG = -2005530596

	// ERR_UNSUPPORTEDALPHAOPERATION indicates that the device does not support
	// a specified texture-blending operation for the alpha channel.
	ERR_UNSUPPORTEDALPHAOPERATION = -2005530597

	// ERR_UNSUPPORTEDCOLORARG indicates that the device does not support a
	// specified texture-blending argument for color values.
	ERR_UNSUPPORTEDCOLORARG = -2005530598

	// ERR_UNSUPPORTEDCOLOROPERATION indicates that the device does not support
	// a specified texture-blending operation for color values.
	ERR_UNSUPPORTEDCOLOROPERATION = -2005530599

	// ERR_UNSUPPORTEDFACTORVALUE indicates that the device does not support
	// the specified texture factor value. Not used; provided only to support
	// older drivers.
	ERR_UNSUPPORTEDFACTORVALUE = -2005530593

	// ERR_UNSUPPORTEDTEXTUREFILTER indicates that the device does not support
	// the specified texture filter.
	ERR_UNSUPPORTEDTEXTUREFILTER = -2005530590

	// ERR_WASSTILLDRAWING indicates that the previous blit operation that is
	// transferring information to or from this surface is incomplete.
	ERR_WASSTILLDRAWING = -2005532132

	// ERR_WRONGTEXTUREFORMAT indicates that the pixel format of the texture
	// surface is not valid.
	ERR_WRONGTEXTUREFORMAT = -2005530600

	// ERR_UNSUPPORTEDOVERLAY indicates that the device does not support
	// overlay for the specified size or display mode.
	// Direct3D 9Ex under Windows 7 only.
	ERR_UNSUPPORTEDOVERLAY = -2005530501

	// ERR_UNSUPPORTEDOVERLAYFORMAT indicates that the device does not support
	// overlay for the specified surface format.
	// Direct3D 9Ex under Windows 7 only.
	ERR_UNSUPPORTEDOVERLAYFORMAT = -2005530500

	// ERR_CANNOTPROTECTCONTENT indicates that the specified content cannot be
	// protected.
	// Direct3D 9Ex under Windows 7 only.
	ERR_CANNOTPROTECTCONTENT = -2005530499

	// ERR_UNSUPPORTEDCRYPTO indicates that the specified cryptographic
	// algorithm is not supported.
	// Direct3D 9Ex under Windows 7 only.
	ERR_UNSUPPORTEDCRYPTO = -2005530498

	// E_INVALIDARG indicates that an invalid parameter was passed to the
	// returning function.
	E_INVALIDARG = -2147024809

	ERR_CANNOT_MODIFY_INDEX_BUFFER = 2900
	ERR_INVALID_MESH               = 2901
	ERR_CANNOT_ATTR_SORT           = 2902
	ERR_SKINNING_NOT_SUPPORTED     = 2903
	ERR_TOO_MANY_INFLUENCES        = 2904
	ERR_INVALID_DATA               = 2905
	ERR_LOADED_MESH_HAS_NO_DATA    = 2906
	ERR_DUPLICATE_NAMED_FRAGMENT   = 2907
	ERR_CANNOT_REMOVE_LAST_ITEM    = 2908
)

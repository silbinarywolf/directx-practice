package main

import (
	"errors"
	"fmt"
	"runtime"
	"syscall"

	"github.com/gonutz/w32"
	"github.com/silbinarywolf/d3d11"
)

const (
	windowWidth  = 1280
	windowHeight = 720
)

var (
	gDriverType       = d3d11.DRIVER_TYPE_NULL
	gFeatureLevel     = d3d11.FEATURE_LEVEL_11_0
	gImmediateContext d3d11.DeviceContext
)

func init() {
	runtime.LockOSThread()
}

// https://github.com/walbourn/directx-sdk-samples/blob/master/Direct3D11Tutorials/Tutorial01/Tutorial01.cpp
func main() {
	/*var windowHandle w32.HWND
	var window w32.HWND
	{
		var err error
		windowHandle, err = openWindow("class name", handleEvent, 0, 0, windowWidth, windowHeight)
		if err != nil {
			panic(err)
		}
		window := w32.HWND(windowHandle)
		w32.SetWindowText(window, "My New Window")
		w32.ShowCursor(false)
	}*/

	driverTypes := []d3d11.DRIVER_TYPE{
		d3d11.DRIVER_TYPE_HARDWARE,
		d3d11.DRIVER_TYPE_WARP,
		d3d11.DRIVER_TYPE_REFERENCE,
	}
	featureLevels := []d3d11.FEATURE_LEVEL{
		d3d11.FEATURE_LEVEL_11_1,
		d3d11.FEATURE_LEVEL_11_0,
		d3d11.FEATURE_LEVEL_10_1,
		d3d11.FEATURE_LEVEL_10_0,
	}
	var (
		device       *d3d11.Device
		featureLevel d3d11.FEATURE_LEVEL
		err          d3d11.Error
	)
	driverType := d3d11.DRIVER_TYPE_NULL
	for driverTypeIndex, _ := range driverTypes {
		driverType = driverTypes[driverTypeIndex]
		device, featureLevel, err = d3d11.CreateDevice(
			0,
			driverType,
			0,
			d3d11.CREATE_DEVICE_DEBUG,
			featureLevels,
			d3d11.SDK_VERSION,
			&gImmediateContext,
		)
		if err != nil {
			if err.Code() != d3d11.E_INVALIDARG {
				panic(err)
			}
			// DirectX 11.0 platforms will not recognize D3D_FEATURE_LEVEL_11_1 so we need to retry without it
			device, featureLevel, err = d3d11.CreateDevice(
				0,
				driverType,
				0,
				d3d11.CREATE_DEVICE_DEBUG,
				featureLevels[:1],
				d3d11.SDK_VERSION,
				&gImmediateContext,
			)
			if err != nil &&
				err.Code() != d3d11.E_INVALIDARG {
				// Fail
				panic(err)
			}
		}
		fmt.Printf(`
			Device: %v
			Feature Level: %v
			Error: %v
		`, device, featureLevel, err)
		if err == nil {
			break
		}
	}
	if err != nil {
		panic(err)
	}

	/*for {
		// oh no
	}*/
}

func openWindow(
	className string,
	callback messageCallback,
	x, y, width, height int,
) (w32.HWND, error) {
	windowProc := syscall.NewCallback(callback)

	class := w32.WNDCLASSEX{
		WndProc:   windowProc,
		Cursor:    w32.LoadCursor(0, w32.MakeIntResource(w32.IDC_ARROW)),
		ClassName: syscall.StringToUTF16Ptr(className),
	}
	if w32.RegisterClassEx(&class) == 0 {
		return 0, errors.New("RegisterClassEx failed")
	}

	window := w32.CreateWindowEx(
		0,
		syscall.StringToUTF16Ptr(className),
		nil,
		w32.WS_OVERLAPPEDWINDOW|w32.WS_VISIBLE,
		x, y, width, height,
		0, 0, 0, nil,
	)
	if window == 0 {
		return 0, errors.New("CreateWindowEx failed")
	}

	return window, nil
}

func handleEvent(window w32.HWND, message uint32, w, l uintptr) uintptr {
	switch message {
	case w32.WM_PAINT:
		return 1
	case w32.WM_KEYDOWN:
		if !isKeyRepeat(l) {
			switch w {
			}
		}
		/*if !isKeyRepeat(l) {
			switch w {
			case w32.VK_LEFT:
				game.HandleInput(InputEvent{GoLeft, true, charIndex})
			case w32.VK_RIGHT:
				game.HandleInput(InputEvent{GoRight, true, charIndex})
			case w32.VK_UP, w32.VK_SPACE:
				game.HandleInput(InputEvent{Jump, true, charIndex})
			case w32.VK_ESCAPE:
				game.HandleInput(InputEvent{QuitGame, true, charIndex})
			}
		}*/
		return 1
	case w32.WM_KEYUP:
		if !isKeyRepeat(l) {
			switch w {
			}
		}
		/*switch w {
		case w32.VK_LEFT:
			game.HandleInput(InputEvent{GoLeft, false, charIndex})
		case w32.VK_RIGHT:
			game.HandleInput(InputEvent{GoRight, false, charIndex})
		case w32.VK_UP, w32.VK_SPACE:
			game.HandleInput(InputEvent{Jump, false, charIndex})
		case w32.VK_F11:
			toggleFullscreen(window)
		case w32.VK_ESCAPE:
			game.HandleInput(InputEvent{QuitGame, false, charIndex})
			w32.PostQuitMessage(0)
		}*/
		return 1
	case w32.WM_SIZE:
		/*if camera != nil {
			windowW, windowH = lowWord(uint(l)), highWord(uint(l))
			camera.setWindowSize(windowW, windowH)
		}*/
		return 1
	case w32.WM_DESTROY:
		w32.PostQuitMessage(0)
		return 1
	default:
		return w32.DefWindowProc(window, message, w, l)
	}
}

func isKeyRepeat(l uintptr) bool {
	return l&(1<<30) != 0
}

type messageCallback func(window w32.HWND, msg uint32, w, l uintptr) uintptr

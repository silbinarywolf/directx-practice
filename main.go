package main

import (
	"errors"
	"runtime"
	"syscall"

	"github.com/gonutz/w32"
	"github.com/silbinarywolf/d3d11"
)

const (
	windowWidth  = 1280
	windowHeight = 720
)

func init() {
	runtime.LockOSThread()
}

// https://github.com/walbourn/directx-sdk-samples/blob/master/Direct3D11Tutorials/Tutorial01/Tutorial01.cpp
func main() {
	windowHandle, err := openWindow("class name", handleEvent, 0, 0, windowWidth, windowHeight)
	if err != nil {
		panic(err)
	}
	window := w32.HWND(windowHandle)

	w32.SetWindowText(window, "My New Window")

	w32.ShowCursor(false)

	d3d11.CreateDevice()
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

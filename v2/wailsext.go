// Package wails is the main package of the Wails project.
// It is used by client applications.
package wails

import (
	"github.com/wailsapp/wails/v2/internal/frontend/desktop/windows/winc/w32"
	"unsafe"
)

// 获取缩放比率
func GetDpi() uint {
	var rgrc w32.RECT
	monitor := w32.MonitorFromRect(&rgrc, w32.MONITOR_DEFAULTTONULL)
	var monitorInfo w32.MONITORINFO
	monitorInfo.CbSize = uint32(unsafe.Sizeof(monitorInfo))
	if monitor != 0 && w32.GetMonitorInfo(monitor, &monitorInfo) {
		var dpiX, dpiY uint
		w32.GetDPIForMonitor(monitor, w32.MDT_EFFECTIVE_DPI, &dpiX, &dpiY)
		return dpiX
	}
	return 0
}

// 获取屏幕大小
func GetScreen() (int, int) {
	// SM_CXFRAME: 窗口左右边框宽度
	framex := w32.GetSystemMetrics(w32.SM_CXSCREEN)
	framey := w32.GetSystemMetrics(w32.SM_CYSCREEN)
	return framex, framey
}

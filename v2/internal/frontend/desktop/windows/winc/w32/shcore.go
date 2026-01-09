//go:build windows

package w32

import (
	"fmt"
	"golang.org/x/sys/windows/registry"
	"syscall"
	"unsafe"
)

var (
	modshcore = syscall.NewLazyDLL("shcore.dll")

	procGetDpiForMonitor = modshcore.NewProc("GetDpiForMonitor")
)

func HasGetDPIForMonitorFunc() bool {
	err := procGetDpiForMonitor.Find()
	return err == nil
}

func GetDPIForMonitor(hmonitor HMONITOR, dpiType MONITOR_DPI_TYPE, dpiX *UINT, dpiY *UINT) uintptr {
	ret, _, _ := procGetDpiForMonitor.Call(
		hmonitor,
		uintptr(dpiType),
		uintptr(unsafe.Pointer(dpiX)),
		uintptr(unsafe.Pointer(dpiY)))

	return ret
}

func GetTextScaleFactor() (int, error) {
	// 方法1：直接读取
	key, err := registry.OpenKey(
		registry.CURRENT_USER,
		`Software\Microsoft\Accessibility`,
		registry.READ,
	)
	if err != nil {
		fmt.Printf("错误: %v (可能使用默认值 100%%)\n", err)
		return 100, err
	}
	defer key.Close()

	factor, _, err := key.GetIntegerValue("TextScaleFactor")
	if err != nil {
		return 100, err
	}

	//fmt.Printf("文本缩放比例: %d%%\n", factor)
	return int(factor), nil
}

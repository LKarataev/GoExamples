package ElemGetter

import (
	"fmt"
	"unsafe"
)

func getElement(arr []int, idx int) (int, error) {
	if idx < 0 || idx >= len(arr) {
		return 0, fmt.Errorf("index out of range")
	}
	return *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&arr[0])) + uintptr(idx)*unsafe.Sizeof(arr[0]))), nil
}

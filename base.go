package gowindows

import (
	"fmt"
	"syscall"
	"unsafe"

	"reflect"
)

const ptrSize = unsafe.Sizeof(uintptr(0))

const ERROR_SUCCESS syscall.Errno = 0

const INVALID_HANDLE_VALUE Handle = Handle(^uintptr(0))

// 转换为 []byte 切片
// 注意，请自己保证内存引用，按文档，reflect.SliceHeader 不会保存 data 的指针，可能会被垃圾回收。
// 另外记得 cgo C 库里面有这个实现。
func ToBytes(data uintptr, len, cap int) []byte {
	var o []byte
	sliceHeader := (*reflect.SliceHeader)((unsafe.Pointer(&o)))
	sliceHeader.Cap = cap
	sliceHeader.Len = len
	sliceHeader.Data = data
	return o
}

// 更改切片尺寸
// 一些 windows 结构内数组长度不确定，这个函数可以强制转换出指定长度的切片
// 输入值：  v 必须是切片的指针
func ChangeSliceSize(v interface{}, Len, Cap int) error {
	if reflect.TypeOf(v).Kind() != reflect.Ptr {
		return fmt.Errorf("v必须是切片的指针。")
	}

	elem := reflect.ValueOf(v).Elem()
	if elem.Kind() != reflect.Slice {
		return fmt.Errorf("v必须是切片的指针。")
	}

	sliceHeader := (*reflect.SliceHeader)((unsafe.Pointer(elem.UnsafeAddr())))
	sliceHeader.Cap = Cap
	sliceHeader.Len = Len
	return nil
}

type Word = uint16

type Bool = int32

type DWord = uint32
type ULong = uint32

type HMODULE = Handle

type WCHAR = wchar_t
type wchar_t = uint16

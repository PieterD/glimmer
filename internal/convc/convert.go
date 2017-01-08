package convc

/*
#include <stdlib.h>
*/
import "C"

import (
	"reflect"
	"unsafe"
)

// Directly convert a string to a byte slice.
// This is highly unsafe, do not modify the resulting slice!
func StringToBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}

// Reverse of StringToBytes.
//
// Directly convert a byte slice to a string.
// This is highly unsafe, do not modify the original slice!
func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// Convert the given string to a byte pointer.
// This is highly unsafe.
func StringToPointer(s string) *uint8 {
	return &StringToBytes(s)[0]
}

// Convert the given byte slice to a byte pointer.
func BytesToPointer(b []byte) *uint8 {
	return &b[0]
}

// Convert the given pointer and size to a string.
// This is highly unsafe.
func PointerToString(p *uint8, size int) string {
	return *(*string)(unsafe.Pointer(&reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(p)),
		Len:  size,
		Cap:  size,
	}))
}

// Convert the given pointer and size to a byte slice.
// This is highly unsafe.
func PointerToBytes(p *uint8, size int) []byte {
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(p)),
		Len:  size,
		Cap:  size,
	}))
}

// Malloc the given number of bytes.
// The resulting memory is not known to the Go garbage collector.
// Call the resulting free function to deallocate the allocated memory.
func Malloc(size int) (bytes []byte, free func()) {
	ptr := (*uint8)(unsafe.Pointer(C.calloc(C.size_t(size), 1)))
	return PointerToBytes(ptr, size), func() { C.free(unsafe.Pointer(ptr)) }
}

// Copy the given string to a newly Malloc'd C string.
func StringToC(s string) (p *uint8, free func()) {
	a, f := Malloc(len(s) + 1)
	copy(a, s)
	return BytesToPointer(a), f
}

// Copy the given byte slice to a newly Malloc'd C string.
func BytesToC(b []byte) (p *uint8, free func()) {
	a, f := Malloc(len(b) + 1)
	copy(a, b)
	return BytesToPointer(a), f
}

// Copy the given strings to a newly Malloc'd C buffer.
// It returns a list of nul-terminated C strings.
func MultiStringToC(ss ...string) (p **uint8, free func()) {
	if len(ss) == 0 {
		return nil, func() {}
	}
	// Allocate enough space for pointers and nul characters
	var q *uint8
	ptrsize := len(ss) * int(unsafe.Sizeof(q))
	size := ptrsize + len(ss)
	// And for all strings
	for _, s := range ss {
		size += len(s)
	}
	bytes, free := Malloc(size)
	pointers := *(*[]*uint8)(unsafe.Pointer(&reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(BytesToPointer(bytes))),
		Len:  len(ss),
		Cap:  len(ss),
	}))
	bytes = bytes[ptrsize:]
	b := bytes
	for i, s := range ss {
		cur := b[:len(s)+1 : len(s)+1]
		copy(cur, s)
		pointers[i] = BytesToPointer(cur)
		b = b[len(s)+1:]
	}
	return (**uint8)(unsafe.Pointer(&pointers[0])), free
}

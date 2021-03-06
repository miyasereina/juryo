// Code generated by gen/gen_pointer.go. DO NOT EDIT.

package proto

import "unsafe"

var boolPtrCodec = codec{
	size:   sizeOfBoolPtr,
	encode: encodeBoolPtr,
	decode: decodeBoolPtr,
}

func sizeOfBoolPtr(p unsafe.Pointer, f *structField) int {
	p = deref(p)
	if p != nil {
		return sizeOfBoolRequired(p, f)
	}
	return 0
}

func encodeBoolPtr(b []byte, p unsafe.Pointer, f *structField) []byte {
	p = deref(p)
	if p != nil {
		return encodeBoolRequired(b, p, f)
	}
	return b
}

func decodeBoolPtr(b []byte, p unsafe.Pointer) (int, error) {
	v := (*unsafe.Pointer)(p)
	if *v == nil {
		*v = unsafe.Pointer(new(bool))
	}
	return decodeBool(b, *v)
}

var stringPtrCodec = codec{
	size:   sizeOfStringPtr,
	encode: encodeStringPtr,
	decode: decodeStringPtr,
}

func sizeOfStringPtr(p unsafe.Pointer, f *structField) int {
	p = deref(p)
	if p != nil {
		return sizeOfStringRequired(p, f)
	}
	return 0
}

func encodeStringPtr(b []byte, p unsafe.Pointer, f *structField) []byte {
	p = deref(p)
	if p != nil {
		return encodeStringRequired(b, p, f)
	}
	return b
}

func decodeStringPtr(b []byte, p unsafe.Pointer) (int, error) {
	v := (*unsafe.Pointer)(p)
	if *v == nil {
		*v = unsafe.Pointer(new(string))
	}
	return decodeString(b, *v)
}

var int32PtrCodec = codec{
	size:   sizeOfInt32Ptr,
	encode: encodeInt32Ptr,
	decode: decodeInt32Ptr,
}

func sizeOfInt32Ptr(p unsafe.Pointer, f *structField) int {
	p = deref(p)
	if p != nil {
		return sizeOfInt32Required(p, f)
	}
	return 0
}

func encodeInt32Ptr(b []byte, p unsafe.Pointer, f *structField) []byte {
	p = deref(p)
	if p != nil {
		return encodeInt32Required(b, p, f)
	}
	return b
}

func decodeInt32Ptr(b []byte, p unsafe.Pointer) (int, error) {
	v := (*unsafe.Pointer)(p)
	if *v == nil {
		*v = unsafe.Pointer(new(int32))
	}
	return decodeInt32(b, *v)
}

var uint32PtrCodec = codec{
	size:   sizeOfUint32Ptr,
	encode: encodeUint32Ptr,
	decode: decodeUint32Ptr,
}

func sizeOfUint32Ptr(p unsafe.Pointer, f *structField) int {
	p = deref(p)
	if p != nil {
		return sizeOfUint32Required(p, f)
	}
	return 0
}

func encodeUint32Ptr(b []byte, p unsafe.Pointer, f *structField) []byte {
	p = deref(p)
	if p != nil {
		return encodeUint32Required(b, p, f)
	}
	return b
}

func decodeUint32Ptr(b []byte, p unsafe.Pointer) (int, error) {
	v := (*unsafe.Pointer)(p)
	if *v == nil {
		*v = unsafe.Pointer(new(uint32))
	}
	return decodeUint32(b, *v)
}

var int64PtrCodec = codec{
	size:   sizeOfInt64Ptr,
	encode: encodeInt64Ptr,
	decode: decodeInt64Ptr,
}

func sizeOfInt64Ptr(p unsafe.Pointer, f *structField) int {
	p = deref(p)
	if p != nil {
		return sizeOfInt64Required(p, f)
	}
	return 0
}

func encodeInt64Ptr(b []byte, p unsafe.Pointer, f *structField) []byte {
	p = deref(p)
	if p != nil {
		return encodeInt64Required(b, p, f)
	}
	return b
}

func decodeInt64Ptr(b []byte, p unsafe.Pointer) (int, error) {
	v := (*unsafe.Pointer)(p)
	if *v == nil {
		*v = unsafe.Pointer(new(int64))
	}
	return decodeInt64(b, *v)
}

var uint64PtrCodec = codec{
	size:   sizeOfUint64Ptr,
	encode: encodeUint64Ptr,
	decode: decodeUint64Ptr,
}

func sizeOfUint64Ptr(p unsafe.Pointer, f *structField) int {
	p = deref(p)
	if p != nil {
		return sizeOfUint64Required(p, f)
	}
	return 0
}

func encodeUint64Ptr(b []byte, p unsafe.Pointer, f *structField) []byte {
	p = deref(p)
	if p != nil {
		return encodeUint64Required(b, p, f)
	}
	return b
}

func decodeUint64Ptr(b []byte, p unsafe.Pointer) (int, error) {
	v := (*unsafe.Pointer)(p)
	if *v == nil {
		*v = unsafe.Pointer(new(uint64))
	}
	return decodeUint64(b, *v)
}

var zigzag32PtrCodec = codec{
	size:   sizeOfZigzag32Ptr,
	encode: encodeZigzag32Ptr,
	decode: decodeZigzag32Ptr,
}

func sizeOfZigzag32Ptr(p unsafe.Pointer, f *structField) int {
	p = deref(p)
	if p != nil {
		return sizeOfZigzag32Required(p, f)
	}
	return 0
}

func encodeZigzag32Ptr(b []byte, p unsafe.Pointer, f *structField) []byte {
	p = deref(p)
	if p != nil {
		return encodeZigzag32Required(b, p, f)
	}
	return b
}

func decodeZigzag32Ptr(b []byte, p unsafe.Pointer) (int, error) {
	v := (*unsafe.Pointer)(p)
	if *v == nil {
		*v = unsafe.Pointer(new(int32))
	}
	return decodeZigzag32(b, *v)
}

var zigzag64PtrCodec = codec{
	size:   sizeOfZigzag64Ptr,
	encode: encodeZigzag64Ptr,
	decode: decodeZigzag64Ptr,
}

func sizeOfZigzag64Ptr(p unsafe.Pointer, f *structField) int {
	p = deref(p)
	if p != nil {
		return sizeOfZigzag64Required(p, f)
	}
	return 0
}

func encodeZigzag64Ptr(b []byte, p unsafe.Pointer, f *structField) []byte {
	p = deref(p)
	if p != nil {
		return encodeZigzag64Required(b, p, f)
	}
	return b
}

func decodeZigzag64Ptr(b []byte, p unsafe.Pointer) (int, error) {
	v := (*unsafe.Pointer)(p)
	if *v == nil {
		*v = unsafe.Pointer(new(int64))
	}
	return decodeZigzag64(b, *v)
}

var fixed32PtrCodec = codec{
	size:   sizeOfFixed32Ptr,
	encode: encodeFixed32Ptr,
	decode: decodeFixed32Ptr,
}

func sizeOfFixed32Ptr(p unsafe.Pointer, f *structField) int {
	p = deref(p)
	if p != nil {
		return sizeOfFixed32Required(p, f)
	}
	return 0
}

func encodeFixed32Ptr(b []byte, p unsafe.Pointer, f *structField) []byte {
	p = deref(p)
	if p != nil {
		return encodeFixed32Required(b, p, f)
	}
	return b
}

func decodeFixed32Ptr(b []byte, p unsafe.Pointer) (int, error) {
	v := (*unsafe.Pointer)(p)
	if *v == nil {
		*v = unsafe.Pointer(new(uint32))
	}
	return decodeFixed32(b, *v)
}

var fixed64PtrCodec = codec{
	size:   sizeOfFixed64Ptr,
	encode: encodeFixed64Ptr,
	decode: decodeFixed64Ptr,
}

func sizeOfFixed64Ptr(p unsafe.Pointer, f *structField) int {
	p = deref(p)
	if p != nil {
		return sizeOfFixed64Required(p, f)
	}
	return 0
}

func encodeFixed64Ptr(b []byte, p unsafe.Pointer, f *structField) []byte {
	p = deref(p)
	if p != nil {
		return encodeFixed64Required(b, p, f)
	}
	return b
}

func decodeFixed64Ptr(b []byte, p unsafe.Pointer) (int, error) {
	v := (*unsafe.Pointer)(p)
	if *v == nil {
		*v = unsafe.Pointer(new(uint64))
	}
	return decodeFixed64(b, *v)
}

var float32PtrCodec = codec{
	size:   sizeOfFloat32Ptr,
	encode: encodeFloat32Ptr,
	decode: decodeFloat32Ptr,
}

func sizeOfFloat32Ptr(p unsafe.Pointer, f *structField) int {
	p = deref(p)
	if p != nil {
		return sizeOfFloat32Required(p, f)
	}
	return 0
}

func encodeFloat32Ptr(b []byte, p unsafe.Pointer, f *structField) []byte {
	p = deref(p)
	if p != nil {
		return encodeFloat32Required(b, p, f)
	}
	return b
}

func decodeFloat32Ptr(b []byte, p unsafe.Pointer) (int, error) {
	v := (*unsafe.Pointer)(p)
	if *v == nil {
		*v = unsafe.Pointer(new(float32))
	}
	return decodeFloat32(b, *v)
}

var float64PtrCodec = codec{
	size:   sizeOfFloat64Ptr,
	encode: encodeFloat64Ptr,
	decode: decodeFloat64Ptr,
}

func sizeOfFloat64Ptr(p unsafe.Pointer, f *structField) int {
	p = deref(p)
	if p != nil {
		return sizeOfFloat64Required(p, f)
	}
	return 0
}

func encodeFloat64Ptr(b []byte, p unsafe.Pointer, f *structField) []byte {
	p = deref(p)
	if p != nil {
		return encodeFloat64Required(b, p, f)
	}
	return b
}

func decodeFloat64Ptr(b []byte, p unsafe.Pointer) (int, error) {
	v := (*unsafe.Pointer)(p)
	if *v == nil {
		*v = unsafe.Pointer(new(float64))
	}
	return decodeFloat64(b, *v)
}

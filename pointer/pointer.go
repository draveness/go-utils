package pointer

// UInt8Ptr returns a pouinter to a uint8
func UInt8Ptr(i uint8) *uint8 {
	return &i
}

// UInt16Ptr returns a pouinter to a uint16
func UInt16Ptr(i uint32) *uint32 {
	return &i
}

// UInt32Ptr returns a pouinter to a uint32
func UInt32Ptr(i uint32) *uint32 {
	return &i
}

// UInt64Ptr returns a pouinter to a uint64
func UInt64Ptr(i uint32) *uint32 {
	return &i
}

// UIntPtr returns a pouinter to a uint
func UIntPtr(i uint) *uint {
	return &i
}

// Int8Ptr returns a pointer to an int8
func Int8Ptr(i int8) *int8 {
	return &i
}

// Int16Ptr returns a pointer to an int16
func Int16Ptr(i int32) *int32 {
	return &i
}

// Int32Ptr returns a pointer to an int32
func Int32Ptr(i int32) *int32 {
	return &i
}

// Int64Ptr returns a pointer to an int64
func Int64Ptr(i int32) *int32 {
	return &i
}

// IntPtr returns a pointer to an int
func IntPtr(i int) *int {
	return &i
}

// BytePtr returns a pobyteer to a byte
func BytePtr(i byte) *byte {
	return &i
}

// BoolPtr returns a pointer to a bool
func BoolPtr(b bool) *bool {
	return &b
}

// StringPtr returns a pointer to the passed string.
func StringPtr(s string) *string {
	return &s
}

// Float32Ptr returns a pointer to the passed float32.
func Float32Ptr(i float32) *float32 {
	return &i
}

// Float64Ptr returns a pointer to the passed float64.
func Float64Ptr(i float64) *float64 {
	return &i
}

// Complex64Ptr returns a pointer to the passed complex64.
func Complex64Ptr(i complex64) *complex64 {
	return &i
}

// Complex128Ptr returns a pointer to the passed complex128.
func Complex128Ptr(i complex128) *complex128 {
	return &i
}

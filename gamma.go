package gamma

import (
	"math/bits"
)

// Encode encodes a list of numbers into a stream of gamma codes
func Encode(numbers []uint32) (stream []uint64) {
	var buf uint64
	idx := uint32(64)

	// compute code stream
	for _, number := range numbers {
		// compute code value and code length
		len := 63 - 2*uint32(bits.LeadingZeros32(number))

		// add code to stream
		if idx >= len {
			// code fits in buf
			buf = buf<<len | uint64(number)
			idx -= len
		} else {
			// code does not fit in buf
			buf = buf<<idx | uint64(number)>>(len-idx)
			stream = append(stream, buf)
			len -= idx
			buf = buf<<len | uint64(number)&(1<<len-1)
			idx = 64 - len
		}
	}

	// padding code stream
	stream = append(stream, buf<<idx)

	return
}

// Encode2 encodes a list of numbers into a stream of gamma codes
func Encode2(numbers []uint32) (stream []uint64) {
	var buf uint64
	idx := uint32(64)

	// compute code stream
	for _, number := range numbers {
		// compute code value and code length
		code, len := encode(number)

		// add code to stream
		if idx >= len {
			// code fits in buf
			buf = buf<<len | code
			idx -= len
		} else {
			// code does not fit in buf
			buf = buf<<idx | code>>(len-idx)
			stream = append(stream, buf)
			len -= idx
			buf = buf<<len | code&(1<<len-1)
			idx = 64 - len
		}
	}

	// padding code stream
	stream = append(stream, buf<<idx)

	return
}

// encode ...
func encode(number uint32) (code uint64, len uint32) {
	return uint64(number), 63 - 2*uint32(bits.LeadingZeros32(number))
}

// Decode decodes a stream of gamma codes into a list of numbers
func Decode(stream []uint64) (numbers []uint32) {
	var bufLen, partialCode, missingLen uint64
	for _, buf := range stream {
		if bufLen != 0 {
			// buffer contains partial codes
			if partialCode != 0 {
				// the read fragment contains 1-bits
				numbers = append(numbers, uint32(partialCode|(buf>>(64-missingLen))))
			} else {
				// the read fragment contains only 0-bits
				missingLen = 1 + bufLen + 2*uint64(bits.LeadingZeros64(buf))
				numbers = append(numbers, uint32(buf>>(64-missingLen)))
			}
			buf <<= missingLen
			bufLen = 64 - missingLen
		} else {
			// buffer does not contain partial code
			bufLen = 64
		}

		// as long as there are 1-bits in the buffer, ...
		for buf != 0 {
			codeLen := 1 + 2*uint64(bits.LeadingZeros64(buf))
			if bufLen >= codeLen {
				// code fully resides in buffer
				numbers = append(numbers, uint32(buf>>(64-codeLen)))
				buf <<= codeLen
				bufLen -= codeLen
			} else {
				// code partially in buffer, prepare partialCode for alignment with missing fragment
				missingLen = bufLen - codeLen
				partialCode, buf = (buf>>(64-codeLen))&^(1<<(missingLen-1)-1), 0
			}
		}
	}
	return
}

// Decode2 decodes a stream of gamma codes into a list of numbers
func Decode2(stream []uint64) (numbers []uint32) {
	var bufLen, partialCode, missingLen uint64
	for _, buf := range stream {
		if bufLen != 0 {
			// buffer contains partial codes
			if partialCode != 0 {
				// the read fragment contains 1-bits
				numbers = append(numbers, uint32(partialCode|(buf>>(64-missingLen))))
			} else {
				// the read fragment contains only 0-bits
				missingLen = 1 + bufLen + 2*uint64(bits.LeadingZeros64(buf))
				numbers = append(numbers, uint32(buf>>(64-missingLen)))
			}
			buf <<= missingLen
			bufLen = 64 - missingLen
		} else {
			// buffer does not contain partial code
			bufLen = 64
		}

		// as long as there are 1-bits in the buffer, ...
		for buf != 0 {
			codeLen := 1 + 2*uint64(bits.LeadingZeros64(buf))
			if bufLen >= codeLen {
				// code fully resides in buffer
				numbers = append(numbers, uint32(buf>>(64-codeLen)))
				buf <<= codeLen
				bufLen -= codeLen
			} else {
				// code partially in buffer, prepare partialCode for alignment with missing fragment
				missingLen = bufLen - codeLen
				partialCode, buf = (buf>>(64-codeLen))&^(1<<(missingLen-1)-1), 0
			}
		}
	}
	return
}

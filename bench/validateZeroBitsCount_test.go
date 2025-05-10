package bench

import (
	"encoding/binary"
	"math/bits"
	"testing"
)

func generateZeroBitsTestCases() [][]byte {
	testCases := make([][]byte, 32)
	for i := range 32 {
		testCases[i] = make([]byte, 4)
		if i != 32 {
			testCases[i][i/8] = 1 << (7 - i%8)
		}
	}
	return testCases
}

func BenchmarkValidateZeroBits(b *testing.B) {
	testCases := generateZeroBitsTestCases()
	b.Run("binaryBits", func(b *testing.B) {
		for i := range b.N {
			count := i % 31
			test := testCases[count]

			if count != bits.LeadingZeros32(binary.BigEndian.Uint32(test)) {
				b.Fatal("unexpected number of zero bits")
			}
		}
	})
	b.Run("mask", func(b *testing.B) {
		for i := range b.N {
			count := i % 31
			test := testCases[count]

			_ = test[3]
			u32 := (uint32(test[0]) << 24) | (uint32(test[1]) << 16) | (uint32(test[2]) << 8) | uint32(test[3])
			mask := uint32(0xffffffff << (32 - count))
			if (u32 & mask) != 0 {
				b.Fatal("unexpected number of zero bits")
			}
		}

	})
	b.Run("loop", func(b *testing.B) {
		for i := range b.N {
			count := i % 31
			test := testCases[count]

			fullBytes := count/8 - 1
			for j := range fullBytes {
				if test[j] != 0 {
					b.Fatal("unexpected number of zero bits")
				}
			}
			bitsRemaining := count % 8 // Number of bits remaining
			if bitsRemaining != 0 && (test[fullBytes+1]&byte(0xff<<(8-bitsRemaining))) != 0 {
				b.Fatal("unexpected number of zero bits")
			}
		}
	})
}

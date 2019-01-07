package main

import (
	"log"
	"math/rand"
	"testing"

	"fmt"

	"encoding/binary"

	"blainsmith.com/go/seahash"
	farm "github.com/dgryski/go-farm"
	"github.com/cespare/xxhash"
	"github.com/minio/highwayhash"
)

const (
	nInput       = 64
	maxStringLen = 1 << 20
)

var sizes = []int{8, 8192, maxStringLen}

func newRand() *rand.Rand { return rand.New(rand.NewSource(0)) }

func randBytes(r *rand.Rand) [][]byte {
	s := make([][]byte, nInput)
	for i := range s {
		s[i] = make([]byte, maxStringLen)
		if _, err := r.Read(s[i]); err != nil {
			log.Panic(err)
		}
	}
	return s
}

func randUints(r *rand.Rand) []uint64 {
	s := make([]uint64, nInput)
	for i := range s {
		s[i] = r.Uint64()
	}
	return s
}

func BenchmarkSeahashString(b *testing.B) {
	r := newRand()
	data := randBytes(r)
	for _, size := range sizes {
		b.Run(fmt.Sprintf("%d", size), func(b *testing.B) {
			idx := 0
			for i := 0; i < b.N; i++ {
				_ = seahash.Sum64(data[idx][:size])
				idx++
				if idx >= len(data) {
					idx = 0
				}
			}
		})
	}
}

func BenchmarkSeahashUInts1(b *testing.B) {
	r := newRand()
	data := randUints(r)
	idx := 0
	for i := 0; i < b.N; i++ {
		buf := [8]byte{}
		binary.LittleEndian.PutUint64(buf[:], data[idx])
		_ = seahash.Sum64(buf[:])
		idx++
		if idx >= len(data) {
			idx = 0
		}
	}
}

func BenchmarkFarmHashString(b *testing.B) {
	r := newRand()
	data := randBytes(r)
	for _, size := range sizes {
		b.Run(fmt.Sprintf("%d", size), func(b *testing.B) {
			idx := 0
			for i := 0; i < b.N; i++ {
				_ = farm.Hash64(data[idx][:size])
				idx++
				if idx >= len(data) {
					idx = 0
				}
			}
		})
	}
}

func BenchmarkFarmHashUInts1(b *testing.B) {
	r := newRand()
	data := randUints(r)
	idx := 0
	for i := 0; i < b.N; i++ {
		_ = farm.Hash64WithSeed(nil, data[idx])
		idx++
		if idx >= len(data) {
			idx = 0
		}
	}
}

func BenchmarkXXHashString(b *testing.B) {
	r := newRand()
	data := randBytes(r)
	for _, size := range sizes {
		b.Run(fmt.Sprintf("%d", size), func(b *testing.B) {
			idx := 0
			for i := 0; i < b.N; i++ {
				_ = xxhash.Sum64(data[idx][:size])
				idx++
				if idx >= len(data) {
					idx = 0
				}
			}
		})
	}
}

func BenchmarkXXHashUInts1(b *testing.B) {
	r := newRand()
	data := randUints(r)
	idx := 0
	for i := 0; i < b.N; i++ {
		buf := [8]byte{}
		binary.LittleEndian.PutUint64(buf[:], data[idx])
		_ = xxhash.Sum64(buf[:])
		idx++
		if idx >= len(data) {
			idx = 0
		}
	}
}

func BenchmarkHighwayHashString(b *testing.B) {
	r := newRand()
	data := randBytes(r)
	var key [32]byte
	for _, size := range sizes {
		b.Run(fmt.Sprintf("%d", size), func(b *testing.B) {
			idx := 0
			for i := 0; i < b.N; i++ {
				_ = highwayhash.Sum64(data[idx][:size], key[:])
				idx++
				if idx >= len(data) {
					idx = 0
				}
			}
		})
	}
}

func BenchmarkHighwayHashUInts1(b *testing.B) {
	r := newRand()
	data := randUints(r)
	idx := 0
	for i := 0; i < b.N; i++ {
		buf := [32]byte{}
		binary.LittleEndian.PutUint64(buf[:], data[idx])
		_ = highwayhash.Sum64(nil, buf[:])
		idx++
		if idx >= len(data) {
			idx = 0
		}
	}
}

// Copyright 2019 Changkun Ou. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package sync

import (
	"math"
	"sync/atomic"
	"unsafe"
)

// AddFloat64 add delta to given address atomically
func AddFloat64(addr *float64, delta float64) (new float64) {
	var old float64
	for {
		old = math.Float64frombits(atomic.LoadUint64((*uint64)(unsafe.Pointer(addr))))
		if atomic.CompareAndSwapUint64((*uint64)(unsafe.Pointer(addr)),
			math.Float64bits(old), math.Float64bits(old+delta)) {
			break
		}
	}
	return
}

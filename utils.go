package main

import (
	"reflect"
	"unsafe"
	"math"
	"x-go-binding.googlecode.com/hg/xgb"
)

type IdList []xgb.Id

func (l IdList) Contains(id xgb.Id) bool {
	for _, i := range l {
		if i == id {
			return true
		}
	}
	return false
}

func propReplyAtoms(prop *xgb.GetPropertyReply) IdList {
	if prop == nil || prop.ValueLen == 0 {
		return nil
	}
	atom_size := uintptr(prop.Format / 8)
	if atom_size != reflect.TypeOf(xgb.Id(0)).Size() {
		panic("Property reply has wrong format for atoms")
	}
	num_atoms := prop.ValueLen / uint32(atom_size)
	return (*[1<<24]xgb.Id)(unsafe.Pointer(&prop.Value[0]))[:num_atoms]
}

func Int16(x int) int16 {
	if x > math.MaxInt16 || x < math.MinInt16 {
		panic("Can't convert int to int16")
	}
	return int16(x)
}

func Uint16(x int) uint16 {
	if x > math.MaxUint16 || x < 0 {
		panic("Can't convert int to uint16")
	}
	return uint16(x)
}

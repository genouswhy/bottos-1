// Copyright 2017~2022 The Bottos Authors
// This file is part of the Bottos Chain library.
// Created by Rocket Core Team of Bottos.

// This program is free software: you can distribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with Bottos.  If not, see <http://www.gnu.org/licenses/>.

// Copyright 2017 The go-interpreter Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package exec

import (
	"math"
	"unsafe"
	"bytes"
	"encoding/binary"
)

func (vm *VM) i32Wrapi64() {
	vm.pushUint32(uint32(vm.popUint64()))
}

func (vm *VM) i32TruncSF32() {
	vm.pushInt32(int32(math.Trunc(float64(vm.popFloat32()))))
}

func (vm *VM) i32TruncUF32() {
	vm.pushUint32(uint32(math.Trunc(float64(vm.popFloat32()))))
}

func (vm *VM) i32TruncSF64() {
	vm.pushInt32(int32(math.Trunc(vm.popFloat64())))
}

func (vm *VM) i32TruncUF64() {
	vm.pushUint32(uint32(math.Trunc(vm.popFloat64())))
}

func (vm *VM) i64ExtendSI32() {
	vm.pushInt64(int64(vm.popInt32()))
}

func (vm *VM) i64ExtendUI32() {
	vm.pushUint64(uint64(vm.popUint32()))
}

func (vm *VM) i64TruncSF32() {
	vm.pushInt64(int64(math.Trunc(float64(vm.popFloat32()))))
}

func (vm *VM) i64TruncUF32() {
	vm.pushUint64(uint64(math.Trunc(float64(vm.popFloat32()))))
}

func (vm *VM) i64TruncSF64() {
	vm.pushInt64(int64(math.Trunc(vm.popFloat64())))
}

func (vm *VM) i64TruncUF64() {
	vm.pushUint64(uint64(math.Trunc(vm.popFloat64())))
}

func (vm *VM) f32ConvertSI32() {
	vm.pushFloat32(float32(vm.popInt32()))
}

func (vm *VM) f32ConvertUI32() {
	vm.pushFloat32(float32(vm.popUint32()))
}

func (vm *VM) f32ConvertSI64() {
	vm.pushFloat32(float32(vm.popInt64()))
}

func (vm *VM) f32ConvertUI64() {
	vm.pushFloat32(float32(vm.popUint64()))
}

func (vm *VM) f32DemoteF64() {
	vm.pushFloat32(float32(vm.popFloat64()))
}

func (vm *VM) f64ConvertSI32() {
	vm.pushFloat64(float64(vm.popInt32()))
}

func (vm *VM) f64ConvertUI32() {
	vm.pushFloat64(float64(vm.popUint32()))
}

func (vm *VM) f64ConvertSI64() {
	vm.pushFloat64(float64(vm.popInt64()))
}

func (vm *VM) f64ConvertUI64() {
	vm.pushFloat64(float64(vm.popUint64()))
}

func (vm *VM) f64PromoteF32() {
	vm.pushFloat64(float64(vm.popFloat32()))
}

func BytesToStr(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func F32ToBytes(f32 float32) []byte {
	bytes := make([]byte, 4)
	bits  := math.Float32bits(f32)

	binary.LittleEndian.PutUint32(bytes, bits)

	return bytes
}

func BytesToF32(b []byte) float32 {
	f32  := math.Float32frombits(binary.LittleEndian.Uint32(b))
	return f32
}

func F64ToBytes(f64 float64) []byte {
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, math.Float64bits(f64))
	return bytes
}

func BytesToF64(b []byte) float64 {
	f64  := math.Float64frombits(binary.LittleEndian.Uint64(b))
	return f64
}

func I32ToBytes(i32 uint32) []byte {
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.LittleEndian, i32)
	return bytesBuffer.Bytes()
}

func I64ToBytes(i64 uint64) []byte {
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.LittleEndian, i64)
	return bytesBuffer.Bytes()
}
// Copyright 2017 The go-interpreter Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package exec

import (
	"github.com/go-interpreter/wagon/exec/internal/compile"
	ops "github.com/go-interpreter/wagon/wasm/operators"
	"math"
)

var (
	GasBlock uint64 = 29
	GasIf    uint64 = 6
	GasLoop  uint64 = 23
	GasBreak uint64 = 4
	GasJump  uint64 = 4

	GasSwitch       uint64 = 12
	GasCall         uint64 = 104
	GasCallImport   uint64 = 30
	GasCallIndirect uint64 = 430
	GasGetLocal     uint64 = 1
	GasSetLocal     uint64 = 4
	GasGetGlobal    uint64 = 8
	GasSetGlobal    uint64 = 8
	GasLoad         uint64 = 7
	GasStore        uint64 = 10
	GasConst        uint64 = 1
	GasUnary        uint64 = 4
	GasBinary       uint64 = 8
	GasSelect       uint64 = 10
	GasDrop         uint64 = 5
	GasReturn       uint64 = 4
	GasHost         uint64 = 3
	GasNop          uint64 = 1
	GasUnreachable  uint64 = math.MaxInt64

	// TODO checktime
	GasCurrentMemory uint64 = 5
	GasGrowMemory    uint64 = 10
)

func (vm *VM) addCallGas(gas uint64) {
	cost := vm.CostGas + gas

	if cost < vm.CostGas {
		panic("gas cost overflow")
	}

	vm.CostGas = cost
	if vm.MaxGas < vm.CostGas {
		panic("gas cost overflow")
	}
}

func (vm *VM) addOpGas(op byte) {
	cost := vm.CostGas + vm.gasTable[op]

	if cost < vm.CostGas {
		panic("gas cost overflow")
	}

	vm.CostGas = cost
	if vm.MaxGas < vm.CostGas {
		panic("gas cost overflow")
	}
}

func (vm *VM) InitGasTable(maxGas uint64) {

	vm.MaxGas = maxGas
	vm.CostGas = 0

	for index := range vm.gasTable {
		vm.gasTable[index] = math.MaxUint64
	}

	vm.initGasTable()
}

func (vm *VM) initGasTable() {

	vm.gasTable[ops.I32Clz] = GasUnary
	vm.gasTable[ops.I32Ctz] = GasUnary
	vm.gasTable[ops.I32Popcnt] = GasUnary
	vm.gasTable[ops.I32Add] = GasBinary
	vm.gasTable[ops.I32Sub] = GasBinary
	vm.gasTable[ops.I32Mul] = GasBinary
	vm.gasTable[ops.I32DivS] = GasBinary
	vm.gasTable[ops.I32DivU] = GasBinary
	vm.gasTable[ops.I32RemS] = GasBinary
	vm.gasTable[ops.I32RemU] = GasBinary
	vm.gasTable[ops.I32And] = GasBinary
	vm.gasTable[ops.I32Or] = GasBinary
	vm.gasTable[ops.I32Xor] = GasBinary
	vm.gasTable[ops.I32Shl] = GasBinary
	vm.gasTable[ops.I32ShrS] = GasBinary
	vm.gasTable[ops.I32ShrU] = GasBinary
	vm.gasTable[ops.I32Rotl] = GasBinary
	vm.gasTable[ops.I32Rotr] = GasBinary
	vm.gasTable[ops.I32Eqz] = GasBinary
	vm.gasTable[ops.I32Eq] = GasBinary
	vm.gasTable[ops.I32Ne] = GasBinary
	vm.gasTable[ops.I32LtS] = GasBinary
	vm.gasTable[ops.I32LtU] = GasBinary
	vm.gasTable[ops.I32GtS] = GasBinary
	vm.gasTable[ops.I32GtU] = GasBinary
	vm.gasTable[ops.I32LeS] = GasBinary
	vm.gasTable[ops.I32LeU] = GasBinary
	vm.gasTable[ops.I32GeS] = GasBinary
	vm.gasTable[ops.I32GeU] = GasBinary

	vm.gasTable[ops.I64Clz] = GasUnary
	vm.gasTable[ops.I64Ctz] = GasUnary
	vm.gasTable[ops.I64Popcnt] = GasUnary
	vm.gasTable[ops.I64Add] = GasBinary
	vm.gasTable[ops.I64Sub] = GasBinary
	vm.gasTable[ops.I64Mul] = GasBinary
	vm.gasTable[ops.I64DivS] = GasBinary
	vm.gasTable[ops.I64DivU] = GasBinary
	vm.gasTable[ops.I64RemS] = GasBinary
	vm.gasTable[ops.I64RemU] = GasBinary
	vm.gasTable[ops.I64And] = GasBinary
	vm.gasTable[ops.I64Or] = GasBinary
	vm.gasTable[ops.I64Xor] = GasBinary
	vm.gasTable[ops.I64Shl] = GasBinary
	vm.gasTable[ops.I64ShrS] = GasBinary
	vm.gasTable[ops.I64ShrU] = GasBinary
	vm.gasTable[ops.I64Rotl] = GasBinary
	vm.gasTable[ops.I64Rotr] = GasBinary
	vm.gasTable[ops.I64Eqz] = GasBinary
	vm.gasTable[ops.I64Eq] = GasBinary
	vm.gasTable[ops.I64Ne] = GasBinary
	vm.gasTable[ops.I64LtS] = GasBinary
	vm.gasTable[ops.I64LtU] = GasBinary
	vm.gasTable[ops.I64GtS] = GasBinary
	vm.gasTable[ops.I64GtU] = GasBinary
	vm.gasTable[ops.I64LeS] = GasBinary
	vm.gasTable[ops.I64LeU] = GasBinary
	vm.gasTable[ops.I64GeS] = GasBinary
	vm.gasTable[ops.I64GeU] = GasBinary

	//vm.gasTable[ops.F32Eq] = vm.f32Eq
	//vm.gasTable[ops.F32Ne] = vm.f32Ne
	//vm.gasTable[ops.F32Lt] = vm.f32Lt
	//vm.gasTable[ops.F32Gt] = vm.f32Gt
	//vm.gasTable[ops.F32Le] = vm.f32Le
	//vm.gasTable[ops.F32Ge] = vm.f32Ge
	//vm.gasTable[ops.F32Abs] = vm.f32Abs
	//vm.gasTable[ops.F32Neg] = vm.f32Neg
	//vm.gasTable[ops.F32Ceil] = vm.f32Ceil
	//vm.gasTable[ops.F32Floor] = vm.f32Floor
	//vm.gasTable[ops.F32Trunc] = vm.f32Trunc
	//vm.gasTable[ops.F32Nearest] = vm.f32Nearest
	//vm.gasTable[ops.F32Sqrt] = vm.f32Sqrt
	//vm.gasTable[ops.F32Add] = vm.f32Add
	//vm.gasTable[ops.F32Sub] = vm.f32Sub
	//vm.gasTable[ops.F32Mul] = vm.f32Mul
	//vm.gasTable[ops.F32Div] = vm.f32Div
	//vm.gasTable[ops.F32Min] = vm.f32Min
	//vm.gasTable[ops.F32Max] = vm.f32Max
	//vm.gasTable[ops.F32Copysign] = vm.f32Copysign

	//vm.gasTable[ops.F64Eq] = vm.f64Eq
	//vm.gasTable[ops.F64Ne] = vm.f64Ne
	//vm.gasTable[ops.F64Lt] = vm.f64Lt
	//vm.gasTable[ops.F64Gt] = vm.f64Gt
	//vm.gasTable[ops.F64Le] = vm.f64Le
	//vm.gasTable[ops.F64Ge] = vm.f64Ge
	//vm.gasTable[ops.F64Abs] = vm.f64Abs
	//vm.gasTable[ops.F64Neg] = vm.f64Neg
	//vm.gasTable[ops.F64Ceil] = vm.f64Ceil
	//vm.gasTable[ops.F64Floor] = vm.f64Floor
	//vm.gasTable[ops.F64Trunc] = vm.f64Trunc
	//vm.gasTable[ops.F64Nearest] = vm.f64Nearest
	//vm.gasTable[ops.F64Sqrt] = vm.f64Sqrt
	//vm.gasTable[ops.F64Add] = vm.f64Add
	//vm.gasTable[ops.F64Sub] = vm.f64Sub
	//vm.gasTable[ops.F64Mul] = vm.f64Mul
	//vm.gasTable[ops.F64Div] = vm.f64Div
	//vm.gasTable[ops.F64Min] = vm.f64Min
	//vm.gasTable[ops.F64Max] = vm.f64Max
	//vm.gasTable[ops.F64Copysign] = vm.f64Copysign

	vm.gasTable[ops.I32Const] = GasConst
	vm.gasTable[ops.I64Const] = GasConst
	//vm.gasTable[ops.F32Const] = vm.f32Const
	//vm.gasTable[ops.F64Const] = vm.f64Const

	//vm.gasTable[ops.I32ReinterpretF32] = vm.i32ReinterpretF32
	//vm.gasTable[ops.I64ReinterpretF64] = vm.i64ReinterpretF64
	//vm.gasTable[ops.F32ReinterpretI32] = vm.f32ReinterpretI32
	//vm.gasTable[ops.F64ReinterpretI64] = vm.f64ReinterpretI64

	vm.gasTable[ops.I32WrapI64] = GasUnary
	//vm.gasTable[ops.I32TruncSF32] = vm.i32TruncSF32
	//vm.gasTable[ops.I32TruncUF32] = vm.i32TruncUF32
	//vm.gasTable[ops.I32TruncSF64] = vm.i32TruncSF64
	//vm.gasTable[ops.I32TruncUF64] = vm.i32TruncUF64
	vm.gasTable[ops.I64ExtendSI32] = GasUnary
	vm.gasTable[ops.I64ExtendUI32] = GasUnary
	//vm.gasTable[ops.I64TruncSF32] = vm.i64TruncSF32
	//vm.gasTable[ops.I64TruncUF32] = vm.i64TruncUF32
	//vm.gasTable[ops.I64TruncSF64] = vm.i64TruncSF64
	//vm.gasTable[ops.I64TruncUF64] = vm.i64TruncUF64
	//vm.gasTable[ops.F32ConvertSI32] = vm.f32ConvertSI32
	//vm.gasTable[ops.F32ConvertUI32] = vm.f32ConvertUI32
	//vm.gasTable[ops.F32ConvertSI64] = vm.f32ConvertSI64
	//vm.gasTable[ops.F32ConvertUI64] = vm.f32ConvertUI64
	//vm.gasTable[ops.F32DemoteF64] = vm.f32DemoteF64
	//vm.gasTable[ops.F64ConvertSI32] = vm.f64ConvertSI32
	//vm.gasTable[ops.F64ConvertUI32] = vm.f64ConvertUI32
	//vm.gasTable[ops.F64ConvertSI64] = vm.f64ConvertSI64
	//vm.gasTable[ops.F64ConvertUI64] = vm.f64ConvertUI64
	//vm.gasTable[ops.F64PromoteF32] = vm.f64PromoteF32

	vm.gasTable[ops.I32Load] = GasLoad
	vm.gasTable[ops.I64Load] = GasLoad
	//vm.gasTable[ops.F32Load] = vm.f32Load
	//vm.gasTable[ops.F64Load] = vm.f64Load
	vm.gasTable[ops.I32Load8s] = GasLoad
	vm.gasTable[ops.I32Load8u] = GasLoad
	vm.gasTable[ops.I32Load16s] = GasLoad
	vm.gasTable[ops.I32Load16u] = GasLoad
	vm.gasTable[ops.I64Load8s] = GasLoad
	vm.gasTable[ops.I64Load8u] = GasLoad
	vm.gasTable[ops.I64Load16s] = GasLoad
	vm.gasTable[ops.I64Load16u] = GasLoad
	vm.gasTable[ops.I64Load32s] = GasLoad
	vm.gasTable[ops.I64Load32u] = GasLoad
	vm.gasTable[ops.I32Store] = GasStore
	vm.gasTable[ops.I64Store] = GasStore
	//vm.gasTable[ops.F32Store] = vm.f32Store
	//vm.gasTable[ops.F64Store] = vm.f64Store
	vm.gasTable[ops.I32Store8] = GasStore
	vm.gasTable[ops.I32Store16] = GasStore
	vm.gasTable[ops.I64Store8] = GasStore
	vm.gasTable[ops.I64Store16] = GasStore
	vm.gasTable[ops.I64Store32] = GasStore
	vm.gasTable[ops.CurrentMemory] = GasCurrentMemory
	vm.gasTable[ops.GrowMemory] = GasGrowMemory

	vm.gasTable[ops.Drop] = GasDrop
	vm.gasTable[ops.Select] = GasSelect

	vm.gasTable[ops.GetLocal] = GasGetLocal
	vm.gasTable[ops.SetLocal] = GasSetLocal
	vm.gasTable[ops.TeeLocal] = GasSetLocal
	vm.gasTable[ops.GetGlobal] = GasGetGlobal
	vm.gasTable[ops.SetGlobal] = GasSetGlobal

	vm.gasTable[ops.Unreachable] = GasUnreachable
	vm.gasTable[ops.Nop] = GasNop

	vm.gasTable[ops.Call] = GasCall
	vm.gasTable[ops.CallIndirect] = GasCallIndirect

	vm.gasTable[ops.BrIf] = GasBlock
	vm.gasTable[ops.Br] = GasBlock
	vm.gasTable[ops.If] = GasIf
	vm.gasTable[ops.Else] = GasIf
	vm.gasTable[ops.End] = GasIf
	vm.gasTable[ops.Loop] = GasLoop
	vm.gasTable[ops.Block] = GasBlock

	vm.gasTable[compile.OpJmp] = GasJump
	vm.gasTable[compile.OpJmpZ] = GasJump
	vm.gasTable[compile.OpJmpNz] = GasJump
	vm.gasTable[compile.OpDiscard] = GasJump
	vm.gasTable[compile.OpDiscardPreserveTop] = GasJump
	vm.gasTable[ops.Return] = GasJump

	vm.gasTable[ops.BrTable] = GasBlock
}

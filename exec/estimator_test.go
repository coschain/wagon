package exec_test

import (
	"bytes"
	"fmt"
	"github.com/go-interpreter/wagon/exec"
	"github.com/go-interpreter/wagon/wasm"
	"io/ioutil"
	"log"
	"reflect"
	"testing"
)

func TestNewEstimator(t *testing.T) {
	add := func(proc *exec.Process, a, b int32) int32 {
		c := a + b
		fmt.Printf("result = %v\n", c)
		return c
	}
	wasmFile := "./testdata/add.wasm"
	code, err := ioutil.ReadFile(wasmFile)
	if err != nil {
		t.Fatal(err)
	}
	m := wasm.NewModule()
	m.Types = &wasm.SectionTypes{
		Entries: []wasm.FunctionSig{
			{
				Form:        0, // value for the 'func' type constructor
				ParamTypes:  []wasm.ValueType{wasm.ValueTypeI32, wasm.ValueTypeI32},
				ReturnTypes: []wasm.ValueType{wasm.ValueTypeI32},
			},
		},
	}
	m.FunctionIndexSpace = []wasm.Function{
		{
			Sig:  &m.Types.Entries[0],
			Host: reflect.ValueOf(add),
			Body: &wasm.FunctionBody{}, // create a dummy wasm body (the actual value will be taken from Host.)
			Gas:  200,
		},
	}
	m.Export = &wasm.SectionExports{
		Entries: map[string]wasm.ExportEntry{
			"add": {
				FieldStr: "add",
				Kind:     wasm.ExternalFunction,
				Index:    0,
			},
		},
	}

	vmModule, err := wasm.ReadModule(bytes.NewReader(code), func(name string) (module *wasm.Module, e error) {
		return m, nil
	})

	vm, err := exec.NewEstimator(vmModule)
	if err != nil {
		log.Fatalf("could not create wagon vm: %v", err)
	}

	var entryIndex = -1
	for name, entry := range vmModule.Export.Entries {
		if name == "main" && entry.Kind == wasm.ExternalFunction {
			entryIndex = int(entry.Index)
		}
	}

	_, err = vm.ExecCode(int64(entryIndex))
	if err != nil {
		log.Fatalf("exec code error: %v", err)
	}
	spent := vm.CostGas
	if spent <= 200 {
		log.Fatalf("do not estimate")
	}
	log.Println(spent)
}

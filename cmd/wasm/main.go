//go:build js && wasm

package main

import (
	"encoding/json"
	"fmt"
	"syscall/js"
)

// Build with GOOS=js GOARCH=wasm go build -o ../../assets/json.wasm -a
func main() {
	fmt.Println("Hello wasm!1")
	js.Global().Set("formatJSON", js.FuncOf(jsonWrapper))
	select {}
}

func prettyJson(input string) (string, error) {
	var raw any
	if err := json.Unmarshal([]byte(input), &raw); err != nil {
		return "", err
	}

	pretty, err := json.MarshalIndent(raw, "", "\t")
	if err != nil {
		return "", err
	}
	return string(pretty), nil
}

func jsonWrapper(this js.Value, args []js.Value) any {
	if len(args) != 1 {
		return []any{nil, "invalid no. of arguments passed"}
	}
	document := js.Global().Get("document")
	if !document.Truthy() {
		return []any{nil, "unable to get document object"}
	}
	jsonOutputBox := document.Call("getElementById", "jsonoutput")
	if !jsonOutputBox.Truthy() {
		return []any{nil, "unable to get json output text area"}
	}

	input := args[0].String()
	fmt.Printf("input: %v\n", input)
	pretty, err := prettyJson(input)
	if err != nil {
		errStr := fmt.Sprintf("unable to convert to json %s\n", err)
		return []any{nil, errStr}
	}
	jsonOutputBox.Set("value", pretty)
	return []any{pretty, nil}
}

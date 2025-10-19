//go:build js && wasm

package main

import (
	"encoding/json"
	"fmt"
	"syscall/js"
)

func main() {
	fmt.Println("Hello wasm!!!!1")
	js.Global().Set("formatJson", jsonWrapper())
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

func jsonWrapper() js.Func {
	jsonFunc := js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) != 1 {
			return "Invalid no. of arguments passed"
		}
		input := args[0].String()
		fmt.Printf("input: %v\n", input)
		pretty, err := prettyJson(input)
		if err != nil {
			return fmt.Sprintf("Unable to prettify json %v", err)
		}
		return pretty
	})

	return jsonFunc
}

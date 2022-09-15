//go:build js && wasm

package console

import "syscall/js"

func Error(message any) {
    js.Global().Get("console").Call("error", message)
}

func Info(message any) {
    js.Global().Get("console").Call("info", message)
}

func Log(message any) {
    js.Global().Get("console").Call("log", message)
}

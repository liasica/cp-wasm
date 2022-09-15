//go:build js && wasm

package encryption

import (
    "fmt"
    "github.com/chatpuppy/encryption-wasm/console"
    "github.com/liasica/go-encryption/aes"
    "syscall/js"
)

// AesEncrypt aes encrypt data (string) using shared key
// args [0] is shared key hex
// args [1] is string data
// returns base64 string
func AesEncrypt(_ js.Value, args []js.Value) (restult any) {
    restult = js.Undefined()
    if len(args) != 2 {
        console.Error("arguments error")
        return
    }

    b, err := getKey(args[0].String())
    if err != nil {
        console.Error(err.Error())
        return
    }

    restult, err = aes.EncryptToBase64([]byte(args[1].String()), b)
    if err != nil {
        console.Error(fmt.Sprintf("data encrypt error: %v", err))
        return
    }

    return
}

// AesDecrypt aes decrypt base64 data using shared key
// args [0] is shared key hex
// args [1] is base64 string data
// returns string value
func AesDecrypt(_ js.Value, args []js.Value) (restult any) {
    restult = js.Undefined()

    if len(args) != 2 {
        console.Error("arguments error")
        return
    }

    b, err := getKey(args[0].String())
    if err != nil {
        console.Error(err.Error())
        return
    }

    var data []byte
    data, err = aes.DecryptFromBase64(args[1].String(), b)
    if err != nil {
        console.Error(fmt.Sprintf("data encrypt error: %v", err))
        return
    }

    restult = string(data)

    return
}

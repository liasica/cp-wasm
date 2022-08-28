//go:build js && wasm

package encryption

import (
    "github.com/chatpuppy/encryption-wasm/console"
    "github.com/liasica/go-encryption/ecdh"
    "regexp"
    "syscall/js"
)

func isValidKey(key string) bool {
    return regexp.MustCompile(`^0x[a-fA-F0-9]+$`).MatchString(key)
}

func Share(_ js.Value, args []js.Value) (shared any) {
    shared = js.Undefined()
    if len(args) != 2 {
        console.Error("arguments error")
        return
    }
    pubStr := args[0].String()
    if !isValidKey(pubStr) {
        console.Error("public key error")
        return
    }

    privStr := args[1].String()
    if !isValidKey(privStr) {
        console.Error("private key error")
        return
    }

    var err error
    shared, err = ecdh.ShareKey(pubStr, privStr)
    if err != nil {
        console.Error(err)
    }

    return
}

func Generate(_ js.Value, _ []js.Value) any {
    priv, pub, err := ecdh.GenerateKey()
    if err != nil {
        console.Error(err)
        return nil
    }

    var s string
    s, err = ecdh.PrivateKeyEncode(priv)
    if err != nil {
        console.Error(err)
        return nil
    }

    return map[string]any{
        "public":  ecdh.PublicKeyEncode(pub),
        "private": s,
    }
}

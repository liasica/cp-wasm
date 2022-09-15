package encryption

import (
    "errors"
    "fmt"
    "github.com/liasica/go-encryption/hexutil"
)

func getKey(hex string) (b []byte, err error) {
    if !isValidKey(hex) {
        err = errors.New("key error")
        return
    }

    b, err = hexutil.Decode(hex)
    if err != nil {
        err = errors.New(fmt.Sprintf("key decode error: %v", err))
        return
    }

    return
}

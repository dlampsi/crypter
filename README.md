# crypter &nbsp;[![GoDoc](https://godoc.org/github.com/dlampsi/crypter?status.svg)](https://godoc.org/github.com/dlampsi/crypter) [![Build Status](https://travis-ci.org/dlampsi/crypter.svg?branch=master)](https://travis-ci.org/dlampsi/crypter)

Simple go funcs to crypt data

## Examples

```go
import "github.com/dlampsi/crypter"
```

Entrypt and decrypt secret string

```go
// Set secret
secret := "MySecretPassword556"
// Generate secret salt key 32 symbols
salt := GenerateRandString(32)

// Encrypt secret
encrypted, err := Encrypt([]byte(secret), []byte(salt))
if err != nil {
    // Handle error
}
fmt.Printf("Enrcypted: %s\n", string(encrypted))

// Decrypt secret
decrypted, err := err := Decrypt(encrypted, []byte(salt))
if err != nil {
    // Handle error
}
fmt.Printf("Decrypted: %s\n", string(decrypted))
```

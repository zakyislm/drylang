package cryptohandler

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"drylang/core"
	"encoding/base64"
	"encoding/hex"
	"io"
)

// BuiltinEnc handles enc.hex, enc.dehex, enc.b64, enc.deb64, enc.aes, enc.deaes
func BuiltinEnc(vm core.VMCore, args []core.Value, line, col int) (core.Value, error) {
	if len(args) < 2 || args[0].Type != core.ValString {
		return core.UnknownValue, vm.Errorf("want enc(method, value)")
	}
	method := args[0].Data.(string)

	switch method {
	case "hex":
		if args[1].Type != core.ValString {
			return core.UnknownValue, vm.Errorf("enc.hex wants string")
		}
		return core.StringVal(hex.EncodeToString([]byte(args[1].Data.(string)))), nil

	case "dehex":
		if args[1].Type != core.ValString {
			return core.UnknownValue, vm.Errorf("enc.dehex wants string")
		}
		b, err := hex.DecodeString(args[1].Data.(string))
		if err != nil {
			return core.UnknownValue, vm.Errorf("dehex fail: %v", err)
		}
		return core.StringVal(string(b)), nil

	case "b64":
		if args[1].Type != core.ValString {
			return core.UnknownValue, vm.Errorf("enc.b64 wants string")
		}
		return core.StringVal(base64.StdEncoding.EncodeToString([]byte(args[1].Data.(string)))), nil

	case "deb64":
		if args[1].Type != core.ValString {
			return core.UnknownValue, vm.Errorf("enc.deb64 wants string")
		}
		b, err := base64.StdEncoding.DecodeString(args[1].Data.(string))
		if err != nil {
			return core.UnknownValue, vm.Errorf("deb64 fail: %v", err)
		}
		return core.StringVal(string(b)), nil

	case "aes":
		// enc("aes", "plaintext", "16-or-32-byte-key")
		if len(args) < 3 {
			return core.UnknownValue, vm.Errorf("enc.aes wants (plaintext, key)")
		}
		plaintext := []byte(args[1].Data.(string))
		key := []byte(args[2].Data.(string))

		block, err := aes.NewCipher(padKey(key))
		if err != nil {
			return core.UnknownValue, vm.Errorf("aes fail: %v", err)
		}
		gcm, err := cipher.NewGCM(block)
		if err != nil {
			return core.UnknownValue, vm.Errorf("aes gcm fail: %v", err)
		}
		nonce := make([]byte, gcm.NonceSize())
		if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
			return core.UnknownValue, vm.Errorf("aes nonce fail: %v", err)
		}
		ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)
		return core.StringVal(base64.StdEncoding.EncodeToString(ciphertext)), nil

	case "deaes":
		// enc("deaes", "ciphertext_b64", "key")
		if len(args) < 3 {
			return core.UnknownValue, vm.Errorf("enc.deaes wants (ciphertext, key)")
		}
		cipherB64 := args[1].Data.(string)
		key := []byte(args[2].Data.(string))

		ciphertext, err := base64.StdEncoding.DecodeString(cipherB64)
		if err != nil {
			return core.UnknownValue, vm.Errorf("deaes b64 fail: %v", err)
		}
		block, err := aes.NewCipher(padKey(key))
		if err != nil {
			return core.UnknownValue, vm.Errorf("deaes fail: %v", err)
		}
		gcm, err := cipher.NewGCM(block)
		if err != nil {
			return core.UnknownValue, vm.Errorf("deaes gcm fail: %v", err)
		}
		nonceSize := gcm.NonceSize()
		if len(ciphertext) < nonceSize {
			return core.UnknownValue, vm.Errorf("deaes: ciphertext too short")
		}
		nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
		plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
		if err != nil {
			return core.UnknownValue, vm.Errorf("deaes decrypt fail: %v", err)
		}
		return core.StringVal(string(plaintext)), nil

	default:
		return core.UnknownValue, vm.Errorf("unknown enc method: %s", method)
	}
}

// padKey pads or truncates key to 32 bytes (AES-256)
func padKey(key []byte) []byte {
	k := make([]byte, 32)
	copy(k, key)
	return k
}

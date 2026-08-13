package cryptohandler

import (
	"crypto/md5"
	"crypto/sha256"
	"drylang/core"
	"encoding/hex"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// BuiltinHash handles hash.bcrypt, hash.sha256, hash.md5, hash.fit
func BuiltinHash(vm core.VMCore, args []core.Value, line, col int) (core.Value, error) {
	if len(args) < 2 || args[0].Type != core.ValString {
		return core.UnknownValue, vm.Errorf("want hash(method, value)")
	}
	method := args[0].Data.(string)

	switch method {
	case "bcrypt":
		if args[1].Type != core.ValString {
			return core.UnknownValue, vm.Errorf("hash.bcrypt wants string")
		}
		plain := args[1].Data.(string)
		hashed, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
		if err != nil {
			return core.UnknownValue, vm.Errorf("bcrypt fail: %v", err)
		}
		return core.StringVal(string(hashed)), nil

	case "sha256":
		if args[1].Type != core.ValString {
			return core.UnknownValue, vm.Errorf("hash.sha256 wants string")
		}
		h := sha256.Sum256([]byte(args[1].Data.(string)))
		return core.StringVal(hex.EncodeToString(h[:])), nil

	case "md5":
		if args[1].Type != core.ValString {
			return core.UnknownValue, vm.Errorf("hash.md5 wants string")
		}
		h := md5.Sum([]byte(args[1].Data.(string)))
		return core.StringVal(hex.EncodeToString(h[:])), nil

	case "fit":
		// hash("fit", "bcrypt", "plain_password", "hashed_password")
		if len(args) < 4 {
			return core.UnknownValue, vm.Errorf("hash.fit wants (algo, plain, hashed)")
		}
		algo := args[1].Data.(string)
		plain := args[2].Data.(string)
		hashed := args[3].Data.(string)

		switch algo {
		case "bcrypt":
			err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))
			return core.BoolVal(err == nil), nil
		case "sha256":
			h := sha256.Sum256([]byte(plain))
			return core.BoolVal(hex.EncodeToString(h[:]) == hashed), nil
		case "md5":
			h := md5.Sum([]byte(plain))
			return core.BoolVal(hex.EncodeToString(h[:]) == hashed), nil
		default:
			return core.UnknownValue, vm.Errorf("unknown hash algo for fit: %s", algo)
		}

	default:
		return core.UnknownValue, fmt.Errorf("unknown hash method: %s", method)
	}
}

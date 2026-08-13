package http

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"drylang/core"
	"encoding/hex"
	"hash"
	"strings"
)

func verifyGithub(payload, signatureHeader, secret string) bool {
	// signatureHeader looks like "sha256=abcdef..."
	parts := strings.SplitN(signatureHeader, "=", 2)
	if len(parts) != 2 {
		return false
	}
	
	algo := parts[0]
	sigHex := parts[1]
	
	var mac hash.Hash
	if algo == "sha256" {
		mac = hmac.New(sha256.New, []byte(secret))
	} else if algo == "sha1" {
		mac = hmac.New(sha1.New, []byte(secret))
	} else {
		return false
	}
	
	mac.Write([]byte(payload))
	expectedMAC := mac.Sum(nil)
	expectedHex := hex.EncodeToString(expectedMAC)
	
	return hmac.Equal([]byte(sigHex), []byte(expectedHex))
}

func verifyStripe(payload, signatureHeader, secret string) bool {
	// Stripe header: t=1611111111,v1=abcdef...,v0=12345...
	parts := strings.Split(signatureHeader, ",")
	var t string
	var v1 []string
	
	for _, p := range parts {
		kv := strings.SplitN(strings.TrimSpace(p), "=", 2)
		if len(kv) == 2 {
			if kv[0] == "t" {
				t = kv[1]
			} else if kv[0] == "v1" {
				v1 = append(v1, kv[1])
			}
		}
	}
	
	if t == "" || len(v1) == 0 {
		return false
	}
	
	signedPayload := t + "." + payload
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(signedPayload))
	expectedHex := hex.EncodeToString(mac.Sum(nil))
	
	for _, sig := range v1 {
		if hmac.Equal([]byte(sig), []byte(expectedHex)) {
			return true
		}
	}
	return false
}

// BuiltinHook handles hook.verify
func BuiltinHook(v core.VMCore, args []core.Value, line, col int) (core.Value, error) {
	if len(args) < 1 || args[0].Type != core.ValString {
		return core.UnknownValue, v.Errorf("want hook(method, ...args)")
	}
	method := args[0].Data.(string)

	switch method {
	case "verify":
		// hook.verify(provider, payload, signatureHeader, secret)
		if len(args) != 5 {
			return core.UnknownValue, v.Errorf("hook.verify wants (provider, payload, header, secret)")
		}
		if args[1].Type != core.ValString || args[2].Type != core.ValString || args[3].Type != core.ValString || args[4].Type != core.ValString {
			return core.UnknownValue, v.Errorf("hook.verify args must be strings")
		}

		provider := args[1].Data.(string)
		payload := args[2].Data.(string)
		header := args[3].Data.(string)
		secret := args[4].Data.(string)

		var valid bool
		if provider == "github" {
			valid = verifyGithub(payload, header, secret)
		} else if provider == "stripe" {
			valid = verifyStripe(payload, header, secret)
		} else {
			return core.UnknownValue, v.Errorf("hook.verify unsupported provider: %s (use github or stripe)", provider)
		}

		return core.Value{Type: core.ValBool, Data: valid}, nil

	default:
		return core.UnknownValue, v.Errorf("unknown hook method: %s", method)
	}
}

package core

import (
	"drylang/core"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

// BuiltinJwt handles jwt.gen and jwt.fit
func BuiltinJwt(vm core.VMCore, args []core.Value, line, col int) (core.Value, error) {
	if len(args) < 2 || args[0].Type != core.ValString {
		return core.UnknownValue, vm.Errorf("want jwt(method, ...args)")
	}
	method := args[0].Data.(string)

	switch method {
	case "gen":
		// jwt("gen", { "id": 1 }, "secret")
		if len(args) < 3 {
			return core.UnknownValue, vm.Errorf("jwt.gen wants (payload, secret)")
		}
		if args[1].Type != core.ValMap {
			return core.UnknownValue, vm.Errorf("jwt.gen payload must be a map")
		}
		if args[2].Type != core.ValString {
			return core.UnknownValue, vm.Errorf("jwt.gen secret must be a string")
		}

		payloadMap := args[1].Data.(map[string]core.Value)
		secret := []byte(args[2].Data.(string))

		claims := jwt.MapClaims{}
		for k, v := range payloadMap {
			claims[k] = extractGoValue(v)
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		signedToken, err := token.SignedString(secret)
		if err != nil {
			return core.UnknownValue, vm.Errorf("jwt.gen fail: %v", err)
		}
		return core.StringVal(signedToken), nil

	case "fit":
		// jwt("fit", "tokenString", "secret")
		if len(args) < 3 {
			return core.UnknownValue, vm.Errorf("jwt.fit wants (token, secret)")
		}
		if args[1].Type != core.ValString {
			return core.UnknownValue, vm.Errorf("jwt.fit token must be a string")
		}
		if args[2].Type != core.ValString {
			return core.UnknownValue, vm.Errorf("jwt.fit secret must be a string")
		}

		tokenStr := args[1].Data.(string)
		secret := []byte(args[2].Data.(string))

		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}
			return secret, nil
		})

		if err != nil || !token.Valid {
			return core.BoolVal(false), nil // Or could return unknown
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			resultMap := make(map[string]core.Value)
			for k, v := range claims {
				resultMap[k] = toDryLangValue(v)
			}
			return core.MapVal(resultMap), nil
		}
		return core.BoolVal(false), nil

	default:
		return core.UnknownValue, vm.Errorf("unknown jwt method: %s", method)
	}
}

func extractGoValue(val core.Value) interface{} {
	switch val.Type {
	case core.ValString:
		return val.Data.(string)
	case core.ValNumber:
		return val.Data.(float64)
	case core.ValBool:
		return val.Data.(bool)
	case core.ValArray:
		arr := val.Data.([]core.Value)
		goArr := make([]interface{}, len(arr))
		for i, v := range arr {
			goArr[i] = extractGoValue(v)
		}
		return goArr
	case core.ValMap:
		m := val.Data.(map[string]core.Value)
		goMap := make(map[string]interface{})
		for k, v := range m {
			goMap[k] = extractGoValue(v)
		}
		return goMap
	default:
		return nil
	}
}

func toDryLangValue(v interface{}) core.Value {
	switch val := v.(type) {
	case string:
		return core.StringVal(val)
	case float64:
		return core.NumberVal(val)
	case bool:
		return core.BoolVal(val)
	case []interface{}:
		arr := make([]core.Value, len(val))
		for i, item := range val {
			arr[i] = toDryLangValue(item)
		}
		return core.ArrayVal(arr)
	case map[string]interface{}:
		m := make(map[string]core.Value)
		for k, item := range val {
			m[k] = toDryLangValue(item)
		}
		return core.MapVal(m)
	default:
		return core.UnknownValue
	}
}

package iohandler

import (
	"bytes"
	"drylang/core"
	"io"
	"net/http"
	"strings"
)

func BuiltinReq(vm core.VMCore, args []core.Value, line, col int) (core.Value, error) {
	var result core.Value
	if len(args) < 1 || args[0].Type != core.ValString {
		return core.UnknownValue, vm.Errorf("E300 at %d:%d: "+"want url string", line, col)
	}
	reqURL := args[0].Data.(string)
	method := "GET"
	var bodyReader io.Reader = nil
	var headers map[string]string

	if len(args) > 1 && args[1].Type == core.ValMap {
		opts := args[1].Data.(map[string]core.Value)
		if v, ok := opts["method"]; ok && v.Type == core.ValString {
			mRaw := strings.ToUpper(v.Data.(string))
			if mRaw == "PO" {
				method = "POST"
			} else if mRaw == "G" {
				method = "GET"
			} else if mRaw == "PUT" {
				method = "PUT"
			} else if mRaw == "DEL" {
				method = "DELETE"
			} else if mRaw == "PAT" {
				method = "PATCH"
			} else if mRaw == "OPT" {
				method = "OPTIONS"
			} else if mRaw == "H" {
				method = "HEAD"
			} else {
				method = mRaw
			}
		}
		if v, ok := opts["body"]; ok && v.Type == core.ValString {
			bodyReader = bytes.NewBufferString(v.Data.(string))
		}
		if v, ok := opts["headers"]; ok && v.Type == core.ValMap {
			headers = make(map[string]string)
			for hk, hv := range v.Data.(map[string]core.Value) {
				if hv.Type == core.ValString {
					headers[hk] = hv.Data.(string)
				}
			}
		}
	}

	req, err := http.NewRequest(method, reqURL, bodyReader)
	if err != nil {
		return core.UnknownValue, vm.Errorf("E300 at %d:%d: req fail: %v", line, col, err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64)")
	if headers != nil {
		for k, v := range headers {
			req.Header.Set(k, v)
		}
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return core.UnknownValue, vm.Errorf("E300 at %d:%d: req fail: %v", line, col, err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return core.UnknownValue, vm.Errorf("E300 at %d:%d: read fail: %v", line, col, err)
	}
	result = core.StringVal(string(body))
	return result, nil
}

package iohandler

import (
	"drylang/core"
	"io"
	"net/http"
)

func BuiltinReq(vm core.VMCore, args []core.Value, line, col int) (core.Value, error) {
	var result core.Value
	if len(args) < 1 || args[0].Type != core.ValString {
		return core.UnknownValue, vm.Errorf("E300 at %d:%d: "+"want url string", line, col)
	}
	reqURL := args[0].Data.(string)
	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		return core.UnknownValue, vm.Errorf("E300 at %d:%d: "+"req fail: %v", err, line, col)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64)")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return core.UnknownValue, vm.Errorf("E300 at %d:%d: "+"req fail: %v", err, line, col)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return core.UnknownValue, vm.Errorf("E300 at %d:%d: "+"read fail: %v", err, line, col)
	}
	result = core.StringVal(string(body))
	return result, nil
}

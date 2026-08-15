package core

import (
	"drylang/core"
	"regexp"
)

// BuiltinRgx handles rgx.match, rgx.find, rgx.rep
func BuiltinRgx(vm core.VMCore, args []core.Value, line, col int) (core.Value, error) {
	if len(args) < 3 || args[0].Type != core.ValString {
		return core.UnknownValue, vm.Errorf("want rgx(method, pattern, string, ...args)")
	}
	method := args[0].Data.(string)
	
	if args[1].Type != core.ValString {
		return core.UnknownValue, vm.Errorf("rgx pattern must be string")
	}
	pattern := args[1].Data.(string)

	re, err := regexp.Compile(pattern)
	if err != nil {
		return core.UnknownValue, vm.Errorf("rgx compile fail: %v", err)
	}

	switch method {
	case "match":
		if args[2].Type != core.ValString {
			return core.UnknownValue, vm.Errorf("rgx.match target must be string")
		}
		target := args[2].Data.(string)
		return core.BoolVal(re.MatchString(target)), nil

	case "find":
		if args[2].Type != core.ValString {
			return core.UnknownValue, vm.Errorf("rgx.find target must be string")
		}
		target := args[2].Data.(string)
		matches := re.FindAllString(target, -1)
		
		res := make([]core.Value, len(matches))
		for i, m := range matches {
			res[i] = core.StringVal(m)
		}
		return core.ArrayVal(res), nil

	case "rep":
		// rgx("rep", "pattern", "replacement", "target")
		if len(args) < 4 {
			return core.UnknownValue, vm.Errorf("rgx.rep wants (pattern, repl, target)")
		}
		if args[2].Type != core.ValString {
			return core.UnknownValue, vm.Errorf("rgx.rep repl must be string")
		}
		if args[3].Type != core.ValString {
			return core.UnknownValue, vm.Errorf("rgx.rep target must be string")
		}
		repl := args[2].Data.(string)
		target := args[3].Data.(string)
		
		result := re.ReplaceAllString(target, repl)
		return core.StringVal(result), nil

	default:
		return core.UnknownValue, vm.Errorf("unknown rgx method: %s", method)
	}
}

package iohandler

import (
	"drylang/core"
	"encoding/csv"
	"os"
)

// BuiltinDoc handles doc("csv_write", path, 2d_array) and doc("csv_read", path)
func BuiltinDoc(v core.VMCore, args []core.Value, line, col int) (core.Value, error) {
	if len(args) < 1 || args[0].Type != core.ValString {
		return core.UnknownValue, v.Errorf("want doc(method, ...args)")
	}
	method := args[0].Data.(string)

	switch method {
	case "csv_write":
		if len(args) != 3 {
			return core.UnknownValue, v.Errorf("doc.csv_write wants (path, array)")
		}
		if args[1].Type != core.ValString || args[2].Type != core.ValArray {
			return core.UnknownValue, v.Errorf("doc.csv_write args must be (string, array)")
		}

		path := args[1].Data.(string)
		rows := args[2].Data.([]core.Value)

		file, err := os.Create(path)
		if err != nil {
			return core.UnknownValue, v.Errorf("csv create error: %s", err)
		}
		defer file.Close()

		writer := csv.NewWriter(file)
		for _, rowVal := range rows {
			if rowVal.Type != core.ValArray {
				continue
			}
			cols := rowVal.Data.([]core.Value)
			var strCols []string
			for _, colVal := range cols {
				strCols = append(strCols, colVal.String())
			}
			if err := writer.Write(strCols); err != nil {
				return core.UnknownValue, v.Errorf("csv write error: %s", err)
			}
		}
		writer.Flush()
		return core.Value{Type: core.ValBool, Data: true}, nil

	case "csv_read":
		if len(args) != 2 || args[1].Type != core.ValString {
			return core.UnknownValue, v.Errorf("doc.csv_read wants (path)")
		}
		path := args[1].Data.(string)

		file, err := os.Open(path)
		if err != nil {
			return core.UnknownValue, v.Errorf("csv open error: %s", err)
		}
		defer file.Close()

		reader := csv.NewReader(file)
		records, err := reader.ReadAll()
		if err != nil {
			return core.UnknownValue, v.Errorf("csv read error: %s", err)
		}

		var outRows []core.Value
		for _, rec := range records {
			var outCols []core.Value
			for _, c := range rec {
				outCols = append(outCols, core.Value{Type: core.ValString, Data: c})
			}
			outRows = append(outRows, core.ArrayVal(outCols))
		}
		return core.ArrayVal(outRows), nil

	default:
		return core.UnknownValue, v.Errorf("unknown doc method: %s", method)
	}
}

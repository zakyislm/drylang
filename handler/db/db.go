package db

import (
	"database/sql"
	"drylang/core"
	"fmt"
	"strings"
	"sync"
)

var dbCache sync.Map

// GetDB returns a cached *sql.DB for the given driver and dsn.
func GetDB(driver, dsn string) (*sql.DB, error) {
	key := driver + "://" + dsn
	if cached, ok := dbCache.Load(key); ok {
		return cached.(*sql.DB), nil
	}
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return nil, err
	}
	// By default, just store it. dbpool.go will allow configuration.
	dbCache.Store(key, db)
	return db, nil
}

func BuiltinDb(vm core.VMCore, args []core.Value, line, col int) (core.Value, error) {
	var result core.Value
	if len(args) < 3 {
		return core.UnknownValue, vm.Errorf("E300 at %d:%d: "+"want driver, dsn, query", line, col)
	}
	driver := args[0].String()
	dsn := args[1].String()
	query := args[2].String()

	if !core.GetSandbox().AllowDB() {
		return core.UnknownValue, vm.Errorf("E300 at %d:%d: "+"db denied", line, col)
	}

	db, err := GetDB(driver, dsn)
	if err != nil {
		return core.UnknownValue, vm.Errorf("E300 at %d:%d: "+"db open err: %v", err, line, col)
	}


	var qargs []interface{}
	for i := 3; i < len(args); i++ {
		if args[i].Type == core.ValNumber {
			qargs = append(qargs, args[i].Data.(float64))
		} else if args[i].Type == core.ValBool {
			qargs = append(qargs, args[i].Data.(bool))
		} else {
			qargs = append(qargs, args[i].String())
		}
	}

	isSelect := strings.HasPrefix(strings.ToUpper(strings.TrimSpace(query)), "SELECT")
	if isSelect || strings.HasPrefix(strings.ToUpper(strings.TrimSpace(query)), "PRAGMA") || strings.HasPrefix(strings.ToUpper(strings.TrimSpace(query)), "SHOW") || strings.HasPrefix(strings.ToUpper(strings.TrimSpace(query)), "DESCRIBE") {
		rows, err := db.Query(query, qargs...)
		if err != nil {
			return core.UnknownValue, vm.Errorf("E300 at %d:%d: "+"db query err: %v", err, line, col)
		}
		defer rows.Close()

		cols, _ := rows.Columns()
		var resultArr []core.Value

		for rows.Next() {
			columns := make([]interface{}, len(cols))
			columnPointers := make([]interface{}, len(cols))
			for i := range columns {
				columnPointers[i] = &columns[i]
			}

			if err := rows.Scan(columnPointers...); err != nil {
				return core.UnknownValue, vm.Errorf("E300 at %d:%d: "+"db scan err: %v", err, line, col)
			}

			rowMap := make(map[string]core.Value)
			for i, colName := range cols {
				val := columns[i]
				if val == nil {
					rowMap[colName] = core.UnknownValue
					continue
				}
				switch v := val.(type) {
				case []byte:
					rowMap[colName] = core.StringVal(string(v))
				case string:
					rowMap[colName] = core.StringVal(v)
				case int64:
					rowMap[colName] = core.NumberVal(float64(v))
				case int:
					rowMap[colName] = core.NumberVal(float64(v))
				case int32:
					rowMap[colName] = core.NumberVal(float64(v))
				case float64:
					rowMap[colName] = core.NumberVal(v)
				case bool:
					rowMap[colName] = core.BoolVal(v)
				default:
					rowMap[colName] = core.StringVal(fmt.Sprintf("%v", v))
				}
			}
			resultArr = append(resultArr, core.Value{Type: core.ValMap, Data: rowMap})
		}
		result = core.Value{Type: core.ValArray, Data: resultArr}
	} else {
		res, err := db.Exec(query, qargs...)
		if err != nil {
			return core.UnknownValue, vm.Errorf("E300 at %d:%d: "+"db exec err: %v", err, line, col)
		}

		lastId, _ := res.LastInsertId()
		rowsAffected, _ := res.RowsAffected()

		m := make(map[string]core.Value)
		m["last_insert_id"] = core.NumberVal(float64(lastId))
		m["rows_affected"] = core.NumberVal(float64(rowsAffected))
		result = core.Value{Type: core.ValMap, Data: m}
	}
	return result, nil
}

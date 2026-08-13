package db

import (
	"drylang/core"
	"time"
)

// BuiltinDbpool handles dbpool("config", driver, dsn, maxOpen, maxIdle, maxLifeTimeMins)
func BuiltinDbpool(vm core.VMCore, args []core.Value, line, col int) (core.Value, error) {
	if len(args) < 2 {
		return core.UnknownValue, vm.Errorf("E300 at %d:%d: want driver, dsn", line, col)
	}

	driver := args[0].String()
	dsn := args[1].String()

	db, err := GetDB(driver, dsn)
	if err != nil {
		return core.UnknownValue, vm.Errorf("E300 at %d:%d: db open err: %v", err, line, col)
	}

	if len(args) >= 3 && args[2].Type == core.ValNumber {
		db.SetMaxOpenConns(int(args[2].Data.(float64)))
	}
	if len(args) >= 4 && args[3].Type == core.ValNumber {
		db.SetMaxIdleConns(int(args[3].Data.(float64)))
	}
	if len(args) >= 5 && args[4].Type == core.ValNumber {
		db.SetConnMaxLifetime(time.Duration(args[4].Data.(float64)) * time.Minute)
	}

	return core.BoolVal(true), nil
}

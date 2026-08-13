package iohandler

import (
	"drylang/core"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Minimalist native cron parser for "* * * * *"
type CronSchedule struct {
	Min, Hour, Dom, Mon, Dow []int
}

func parseField(field string, min, max int) ([]int, error) {
	if field == "*" || field == "?" {
		res := make([]int, max-min+1)
		for i := min; i <= max; i++ {
			res[i-min] = i
		}
		return res, nil
	}

	var res []int
	parts := strings.Split(field, ",")
	for _, p := range parts {
		step := 1
		stepParts := strings.Split(p, "/")
		if len(stepParts) == 2 {
			s, err := strconv.Atoi(stepParts[1])
			if err != nil {
				return nil, err
			}
			step = s
			p = stepParts[0]
		}
		
		if p == "*" {
			for i := min; i <= max; i += step {
				res = append(res, i)
			}
			continue
		}

		rangeParts := strings.Split(p, "-")
		if len(rangeParts) == 2 {
			rMin, _ := strconv.Atoi(rangeParts[0])
			rMax, _ := strconv.Atoi(rangeParts[1])
			for i := rMin; i <= rMax; i += step {
				res = append(res, i)
			}
			continue
		}

		val, err := strconv.Atoi(p)
		if err != nil {
			return nil, err
		}
		res = append(res, val)
	}
	return res, nil
}

func parseCron(expr string) (*CronSchedule, error) {
	fields := strings.Fields(expr)
	if len(fields) != 5 {
		return nil, fmt.Errorf("cron expression must have 5 fields")
	}

	mins, err := parseField(fields[0], 0, 59)
	if err != nil { return nil, err }
	hours, err := parseField(fields[1], 0, 23)
	if err != nil { return nil, err }
	doms, err := parseField(fields[2], 1, 31)
	if err != nil { return nil, err }
	mons, err := parseField(fields[3], 1, 12)
	if err != nil { return nil, err }
	dows, err := parseField(fields[4], 0, 6)
	if err != nil { return nil, err }

	return &CronSchedule{
		Min:  mins,
		Hour: hours,
		Dom:  doms,
		Mon:  mons,
		Dow:  dows,
	}, nil
}

func contains(arr []int, val int) bool {
	for _, a := range arr {
		if a == val {
			return true
		}
	}
	return false
}

func (s *CronSchedule) Matches(t time.Time) bool {
	return contains(s.Min, t.Minute()) &&
		contains(s.Hour, t.Hour()) &&
		contains(s.Dom, t.Day()) &&
		contains(s.Mon, int(t.Month())) &&
		contains(s.Dow, int(t.Weekday()))
}

func parseCronAndTick(cronExpr string, fn *core.Closure, args []core.Value, v core.VMCore) error {
	var sched *CronSchedule
	var err error
	
	// Support both @every X and * * * * *
	if strings.HasPrefix(cronExpr, "@every ") {
		durStr := strings.TrimSpace(strings.TrimPrefix(cronExpr, "@every "))
		parsed, err := time.ParseDuration(durStr)
		if err == nil {
			ticker := time.NewTicker(parsed)
			go func() {
				for range ticker.C {
					execVM := v.Clone()
					execVM.Push(core.Value{Type: core.ValFn, Data: fn})
					for _, arg := range args {
						execVM.Push(arg)
					}
					execVM.CallFunction(fn.Fn, len(args))
				}
			}()
			return nil
		}
		return err
	}

	sched, err = parseCron(cronExpr)
	if err != nil {
		return err
	}

	go func() {
		for {
			now := time.Now()
			// Wait until next minute boundary
			nextMin := now.Truncate(time.Minute).Add(time.Minute)
			time.Sleep(time.Until(nextMin))

			if sched.Matches(time.Now()) {
				execVM := v.Clone()
				execVM.Push(core.Value{Type: core.ValFn, Data: fn})
				for _, arg := range args {
					execVM.Push(arg)
				}
				execVM.CallFunction(fn.Fn, len(args))
			}
		}
	}()
	return nil
}

// BuiltinCron handles cron.add
func BuiltinCron(v core.VMCore, args []core.Value, line, col int) (core.Value, error) {
	if len(args) < 1 || args[0].Type != core.ValString {
		return core.UnknownValue, v.Errorf("want cron(method, ...args)")
	}
	method := args[0].Data.(string)

	switch method {
	case "add":
		// cron.add("* * * * *", fn, arg1, ...)
		if len(args) < 3 {
			return core.UnknownValue, v.Errorf("cron.add wants (expr, fn, ...args)")
		}
		if args[1].Type != core.ValString {
			return core.UnknownValue, v.Errorf("cron.add arg 1 must be string (expr)")
		}
		if args[2].Type != core.ValFn {
			return core.UnknownValue, v.Errorf("cron.add arg 2 must be fn")
		}

		expr := args[1].Data.(string)
		fn := args[2].Data.(*core.Closure)
		fnArgs := args[3:]

		err := parseCronAndTick(expr, fn, fnArgs, v)
		if err != nil {
			return core.UnknownValue, v.Errorf("cron parse error: %v", err)
		}

		return core.UnknownValue, nil

	default:
		return core.UnknownValue, v.Errorf("unknown cron method: %s", method)
	}
}

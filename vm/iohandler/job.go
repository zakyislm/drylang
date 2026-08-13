package iohandler

import (
	"drylang/core"
)

type JobTask struct {
	Fn   *core.Closure
	Args []core.Value
}

var jobQueue chan JobTask
var maxWorkers = 10
var currentWorkers = 0

func startWorker(v core.VMCore) {
	go func() {
		for task := range jobQueue {
			execVM := v.Clone()
			execVM.Push(core.Value{Type: core.ValFn, Data: task.Fn})
			for _, arg := range task.Args {
				execVM.Push(arg)
			}
			execVM.CallFunction(task.Fn.Fn, len(task.Args))
		}
	}()
}

// BuiltinJob handles job.push, job.config
func BuiltinJob(v core.VMCore, args []core.Value, line, col int) (core.Value, error) {
	if len(args) < 1 || args[0].Type != core.ValString {
		return core.UnknownValue, v.Errorf("want job(method, ...args)")
	}
	method := args[0].Data.(string)

	switch method {
	case "config":
		if len(args) >= 2 && args[1].Type == core.ValNumber {
			maxWorkers = int(args[1].Data.(float64))
		}
		queueSize := 1000
		if len(args) >= 3 && args[2].Type == core.ValNumber {
			queueSize = int(args[2].Data.(float64))
		}
		if jobQueue == nil {
			jobQueue = make(chan JobTask, queueSize)
		}
		return core.UnknownValue, nil

	case "push":
		if len(args) < 2 {
			return core.UnknownValue, v.Errorf("job.push wants (fn, ...args)")
		}
		if args[1].Type != core.ValFn {
			return core.UnknownValue, v.Errorf("job.push arg 1 must be fn")
		}
		fn := args[1].Data.(*core.Closure)
		fnArgs := args[2:]

		if jobQueue == nil {
			jobQueue = make(chan JobTask, 1000)
		}

		if currentWorkers < maxWorkers {
			currentWorkers++
			startWorker(v)
		}

		jobQueue <- JobTask{Fn: fn, Args: fnArgs}
		return core.UnknownValue, nil

	default:
		return core.UnknownValue, v.Errorf("unknown job method: %s", method)
	}
}

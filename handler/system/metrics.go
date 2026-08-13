package system

import (
	"expvar"
	"sync/atomic"
)

// Runtime metrics exposed via /debug/vars on any dryLang HTTP server.
var (
	Requests  atomic.Int64 // total HTTP requests handled
	Errors    atomic.Int64 // runtime errors thrown
	Builtins  atomic.Int64 // builtin calls executed
	AsyncJobs atomic.Int64 // async tasks spawned
)

func init() {
	expvar.Publish("dry_requests", expvar.Func(func() interface{} { return Requests.Load() }))
	expvar.Publish("dry_errors", expvar.Func(func() interface{} { return Errors.Load() }))
	expvar.Publish("dry_builtins", expvar.Func(func() interface{} { return Builtins.Load() }))
	expvar.Publish("dry_async_jobs", expvar.Func(func() interface{} { return AsyncJobs.Load() }))
}

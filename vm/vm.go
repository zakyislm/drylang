package vm

import (
	"bufio"
	"drylang/core"
	"sync"
)

// VM executes dryLang bytecode.
type VM struct {
	mu           sync.Mutex
	gmu          sync.RWMutex // guards globals map for concurrent clone/exec
	chunk        *core.Chunk
	fns          []*core.CompiledFn
	globals      map[string]core.Value
	stack        []core.Value
	sp           int // stack pointer
	ip           int // instruction pointer
	frames       []callFrame
	tryStack     []tryFrame
	stdinScanner *bufio.Scanner

	// Async state
	asyncWg    sync.WaitGroup
	asyncPools map[*core.CompiledFn]chan asyncTask
	poolMutex  sync.Mutex
}

// New creates a new VM.
func New(chunk *core.Chunk, fns []*core.CompiledFn) *VM {
	vm := &VM{
		chunk:      chunk,
		fns:        fns,
		globals:    make(map[string]core.Value),
		stack:      make([]core.Value, 4096),
		sp:         0,
		ip:         0,
		asyncPools: make(map[*core.CompiledFn]chan asyncTask),
	}
	vm.SetGlobal("null", core.UnknownValue)
	registerBuiltinModules(vm)
	return vm
}

func (vm *VM) Update(chunk *core.Chunk, fns []*core.CompiledFn) {
	vm.chunk = chunk
	vm.fns = fns
}

// Run executes the main chunk.
func (vm *VM) Run() error {
	// Push base frame for global scope so top-level blocks (loop/try) can use locals
	if len(vm.frames) == 0 {
		vm.frames = append(vm.frames, callFrame{
			closure: &core.Closure{Fn: &core.CompiledFn{Chunk: vm.chunk}, Env: make(map[string]core.Value)},
			ip:      0,
			bp:      vm.sp,
			chunk:   vm.chunk,
		})
		for i := 0; i < vm.chunk.LocalsCount; i++ {
			vm.push(core.UnknownValue)
		}
	}
	return vm.execute(vm.chunk)
}

func (vm *VM) execute(chunk *core.Chunk) error {
	vm.chunk = chunk
	vm.ip = 0
	baseDepth := len(vm.frames)

	for {
		err := vm.executeInner(vm.chunk)
		if err == nil {
			return nil
		}

		if len(vm.tryStack) > 0 {
			tf := vm.tryStack[len(vm.tryStack)-1]
			if tf.frameDepth == baseDepth {
				vm.tryStack = vm.tryStack[:len(vm.tryStack)-1]
				vm.sp = tf.sp
				vm.push(core.StringVal(err.Error()))
				vm.ip = tf.catchIP
				vm.chunk = chunk
				continue
			}
		}
		return err
	}
}

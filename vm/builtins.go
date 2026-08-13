package vm

import (
	"drylang/core"
	hcore "drylang/handler/core"
	"drylang/handler/concurrency"
	"drylang/handler/data"
	"drylang/handler/db"
	hhttp "drylang/handler/http"
	"drylang/handler/messaging"
	"drylang/handler/state"
	"drylang/handler/system"
)

// registerBuiltinModules seeds the globals with builtin module markers so
// `math.sqrt(...)`-style module access resolves via OpDotGet.
func registerBuiltinModules(vm *VM) {
	for name, id := range core.BuiltinNames {
		vm.SetGlobal(name, core.Value{
			Type: core.ValBuiltinModule,
			Data: core.BuiltinModule{ModuleID: int(id)},
		})
	}
}

// executeBuiltin dispatches a builtin call. Arguments are popped from the stack;
// the result is pushed. All implementations live in handler packages.
func (vm *VM) executeBuiltin(id core.BuiltinID, argCount int, line, col int) error {
	args := make([]core.Value, argCount)
	for i := argCount - 1; i >= 0; i-- {
		args[i] = vm.pop()
	}

	res, err := vm.callBuiltin(id, args, line, col)
	if err != nil {
		return err
	}

	vm.push(res)
	return nil
}

func (vm *VM) callBuiltin(id core.BuiltinID, args []core.Value, line, col int) (core.Value, error) {
	switch id {
	case core.BuiltinLen:
		return hcore.BuiltinLen(vm, args, line, col)
	case core.BuiltinGet:
		return hcore.BuiltinGet(vm, args, line, col)
	case core.BuiltinAdd:
		return hcore.BuiltinAdd(vm, args, line, col)
	case core.BuiltinNum:
		return hcore.BuiltinNum(vm, args, line, col)
	case core.BuiltinStr:
		return hcore.BuiltinStr(vm, args, line, col)
	case core.BuiltinAbs:
		return hcore.BuiltinAbs(vm, args, line, col)
	case core.BuiltinMin:
		return hcore.BuiltinMin(vm, args, line, col)
	case core.BuiltinMax:
		return hcore.BuiltinMax(vm, args, line, col)
	case core.BuiltinRnd:
		return hcore.BuiltinRnd(vm, args, line, col)
	case core.BuiltinCap:
		return hcore.BuiltinCap(vm, args, line, col)
	case core.BuiltinLow:
		return hcore.BuiltinLow(vm, args, line, col)
	case core.BuiltinTrm:
		return hcore.BuiltinTrm(vm, args, line, col)
	case core.BuiltinSpl:
		return hcore.BuiltinSpl(vm, args, line, col)
	case core.BuiltinJ:
		return hcore.BuiltinJ(vm, args, line, col)
	case core.BuiltinMod:
		return hcore.BuiltinMod(vm, args, line, col)
	case core.BuiltinHas:
		return hcore.BuiltinHas(vm, args, line, col)
	case core.BuiltinSort:
		return hcore.BuiltinSort(vm, args, line, col)
	case core.BuiltinPop:
		return hcore.BuiltinPop(vm, args, line, col)
	case core.BuiltinRm:
		return hcore.BuiltinRm(vm, args, line, col)
	case core.BuiltinKey:
		return hcore.BuiltinKey(vm, args, line, col)
	case core.BuiltinVal:
		return hcore.BuiltinVal(vm, args, line, col)
	case core.BuiltinRan:
		return hcore.BuiltinRan(vm, args, line, col)
	case core.BuiltinQ:
		return hcore.BuiltinQ(vm, args, line, col)
	case core.BuiltinR:
		return hcore.BuiltinR(vm, args, line, col)
	case core.BuiltinW:
		return hcore.BuiltinW(vm, args, line, col)
	case core.BuiltinNow:
		return hcore.BuiltinNow(vm, args, line, col)
	case core.BuiltinDate:
		return hcore.BuiltinDate(vm, args, line, col)
	case core.BuiltinReq:
		return hhttp.BuiltinReq(vm, args, line, col)
	case core.BuiltinJson:
		return data.BuiltinJson(vm, args, line, col)
	case core.BuiltinArg:
		return hcore.BuiltinArg(vm, args, line, col)
	case core.BuiltinEnv:
		return hcore.BuiltinEnv(vm, args, line, col)
	case core.BuiltinCmd:
		return hcore.BuiltinCmd(vm, args, line, col)
	case core.BuiltinDir:
		return hcore.BuiltinDir(vm, args, line, col)
	case core.BuiltinDel:
		return hcore.BuiltinDel(vm, args, line, col)
	case core.BuiltinDie:
		return hcore.BuiltinDie(vm, args, line, col)
	case core.BuiltinOp:
		return vm.builtinOp(args, line, col)
	case core.BuiltinDb:
		return db.BuiltinDb(vm, args, line, col)
	case core.BuiltinMath:
		return hcore.BuiltinMath(vm, args, line, col)
	case core.BuiltinIn:
		return vm.builtinIn(args, line, col)
	case core.BuiltinPt:
		return hcore.BuiltinPt(vm, args, line, col)
	case core.BuiltinHash:
		return hcore.BuiltinHash(vm, args, line, col)
	case core.BuiltinEnc:
		return hcore.BuiltinEnc(vm, args, line, col)
	case core.BuiltinJwt:
		return hcore.BuiltinJwt(vm, args, line, col)
	case core.BuiltinRgx:
		return hcore.BuiltinRgx(vm, args, line, col)
	case core.BuiltinFmt:
		return hcore.BuiltinFmt(vm, args, line, col)
	case core.BuiltinValid:
		return data.BuiltinValid(vm, args, line, col)
	case core.BuiltinRt:
		return hhttp.BuiltinRt(vm, args, line, col)
	case core.BuiltinSys:
		return system.BuiltinSys(vm, args, line, col)
	case core.BuiltinPipe:
		return concurrency.BuiltinPipe(vm, args, line, col)
	case core.BuiltinCron:
		return concurrency.BuiltinCron(vm, args, line, col)
	case core.BuiltinJob:
		return concurrency.BuiltinJob(vm, args, line, col)
	case core.BuiltinRate:
		return concurrency.BuiltinRate(vm, args, line, col)
	case core.BuiltinSess:
		return state.BuiltinSess(vm, args, line, col)
	case core.BuiltinHook:
		return hhttp.BuiltinHook(vm, args, line, col)
	case core.BuiltinFlag:
		return state.BuiltinFlag(vm, args, line, col)
	case core.BuiltinImg:
		return data.BuiltinImg(vm, args, line, col)
	case core.BuiltinDoc:
		return data.BuiltinDoc(vm, args, line, col)
	case core.BuiltinTmpl:
		return data.BuiltinTmpl(vm, args, line, col)
	case core.BuiltinMail:
		return messaging.BuiltinMail(vm, args, line, col)
	case core.BuiltinMem:
		return state.BuiltinMem(vm, args, line, col)
	case core.BuiltinWs:
		return hhttp.BuiltinWs(vm, args, line, col)
	case core.BuiltinRpc:
		return messaging.BuiltinRpc(vm, args, line, col)
	case core.BuiltinMet:
		return system.BuiltinMet(vm, args, line, col)
	case core.BuiltinGeo:
		return data.BuiltinGeo(vm, args, line, col)
	case core.BuiltinFlow:
		return state.BuiltinFlow(vm, args, line, col)
	case core.BuiltinTest:
		return system.BuiltinTest(vm, args, line, col)
	case core.BuiltinDbpool:
		return db.BuiltinDbpool(vm, args, line, col)
	case core.BuiltinLog:
		return system.BuiltinLog(vm, args, line, col)
	}

	return core.UnknownValue, vm.runtimeErr("E300", line, col, "unknown builtin")
}

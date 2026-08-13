package core

// BuiltinID identifies built-in functions.
type BuiltinID int

const (
	BuiltinLen BuiltinID = iota
	BuiltinGet
	BuiltinAdd
	BuiltinNum
	BuiltinStr
	BuiltinAbs
	BuiltinMin
	BuiltinMax
	BuiltinRnd
	BuiltinCap
	BuiltinLow
	BuiltinTrm
	BuiltinSpl
	BuiltinJ
	BuiltinMod
	BuiltinHas
	BuiltinSort
	BuiltinPop
	BuiltinRm
	BuiltinKey
	BuiltinVal
	BuiltinRan
	BuiltinQ
	BuiltinR
	BuiltinW
	BuiltinNow
	BuiltinDate
	BuiltinReq
	BuiltinJson
	BuiltinArg
	BuiltinEnv
	BuiltinCmd
	BuiltinDir
	BuiltinDel
	BuiltinDie
	BuiltinOp
	BuiltinDb
	BuiltinMath
	BuiltinIn
	BuiltinPt
	BuiltinHash
	BuiltinEnc
	BuiltinJwt
	BuiltinRgx
	BuiltinFmt
	BuiltinValid
	BuiltinRt
	BuiltinSys
	BuiltinPipe
	BuiltinCron
	BuiltinJob
	BuiltinRate
	BuiltinSess
	BuiltinHook
	BuiltinFlag
	BuiltinImg
	BuiltinDoc
	BuiltinTmpl
	BuiltinMail
	BuiltinMem
	BuiltinWs
	BuiltinRpc
	BuiltinMet
	BuiltinGeo
	BuiltinFlow
	BuiltinTest
	BuiltinDbpool
	BuiltinLog
)

// BuiltinNames maps function names to builtin IDs.
var BuiltinNames = map[string]BuiltinID{
	"len":  BuiltinLen,
	"get":  BuiltinGet,
	"add":  BuiltinAdd,
	"num":  BuiltinNum,
	"str":  BuiltinStr,
	"abs":  BuiltinAbs,
	"min":  BuiltinMin,
	"max":  BuiltinMax,
	"rnd":  BuiltinRnd,
	"cap":  BuiltinCap,
	"low":  BuiltinLow,
	"trm":  BuiltinTrm,
	"spl":  BuiltinSpl,
	"j":    BuiltinJ,
	"mod":  BuiltinMod,
	"has":  BuiltinHas,
	"sort": BuiltinSort,
	"pop":  BuiltinPop,
	"rm":   BuiltinRm,
	"key":  BuiltinKey,
	"val":  BuiltinVal,
	"ran":  BuiltinRan,
	"q":    BuiltinQ,
	"r":    BuiltinR,
	"w":    BuiltinW,
	"now":  BuiltinNow,
	"date": BuiltinDate,
	"req":  BuiltinReq,
	"json": BuiltinJson,
	"arg":  BuiltinArg,
	"env":  BuiltinEnv,
	"cmd":  BuiltinCmd,
	"dir":  BuiltinDir,
	"del":  BuiltinDel,
	"die":  BuiltinDie,
	"op":   BuiltinOp,
	"db":   BuiltinDb,
	"math": BuiltinMath,
	"in":   BuiltinIn,
	"pt":   BuiltinPt,
	"hash": BuiltinHash,
	"enc":  BuiltinEnc,
	"jwt":  BuiltinJwt,
	"rgx":  BuiltinRgx,
	"fmt":  BuiltinFmt,
	"valid": BuiltinValid,
	"rt":   BuiltinRt,
	"sys":  BuiltinSys,
	"pipe": BuiltinPipe,
	"cron": BuiltinCron,
	"job":  BuiltinJob,
	"rate": BuiltinRate,
	"sess": BuiltinSess,
	"flag": BuiltinFlag,
	"img":  BuiltinImg,
	"doc":  BuiltinDoc,
	"tmpl": BuiltinTmpl,
	"mail": BuiltinMail,
	"mem":  BuiltinMem,
	"ws":   BuiltinWs,
	"rpc":  BuiltinRpc,
	"met":  BuiltinMet,
	"geo":  BuiltinGeo,
	"flow": BuiltinFlow,
	"test": BuiltinTest,
	"dbpool": BuiltinDbpool,
	"log": BuiltinLog,
}

type StructDef struct {
	Name       string
	Fields     []string
	Visibility string
}

type ClassDef struct {
	Name       string
	ParentNames []string
	Parents    []*ClassDef
	Fields     []string
	Methods    map[string]ClassMethod
	Visibility string
}

type ClassMethod struct {
	Chunk      *Chunk
	Visibility string
	IsAsync    bool
}


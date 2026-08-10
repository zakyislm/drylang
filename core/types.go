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


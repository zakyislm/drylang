# DryLang Dictionary

This file indexes all keywords, built-in functions, operators, delimiters, literals, and error codes in DryLang, mapping their semantics to the compiler pipeline. Designed with Write Less Get More.

> **Structure note (post-Phase 1 refactor):** paths below reflect the current layout.
> - Builtin implementations: `handler/{core,db,http,messaging,concurrency,state,data,system}/`
> - Opcode handlers: `vm/op/` (e.g. `op_math_add.go`)
> - VM core: `vm/{vm,execute,op_calls,builtins,builtin_server,stack,frames,globals,runtime_error}.go`
> - Statement parsing: `parser/stmt/` · Expression parsing: `parser/expr/`
> - Compiler feature files: `compiler/{emit,expr,expr_helpers,fn_compile,stmt_*}.go`
> - Sandbox config: `core/sandbox.go` (env: `DRY_ALLOW_CMD`, `DRY_ALLOW_DB`, `DRY_ALLOW_URL`, `DRY_CORS_ORIGIN`, `DRY_LOG_JSON`)
> - Errors: `errfmt.DryError{Code,Line,Col,Message}` · Metrics: expvar `/debug/vars`
> - Build/version: `make build` · `-ldflags "-X drylang/cli.Version=..."`

## Keywords

### `asn`
**Means:** Async function modifier.
- Lexer: [lexer/token.go:32](lexer/token.go#L32)
- AST: [ast/](ast/)
- Parser: [parser/stmt/fn.go](parser/stmt/fn.go)
- Compiler: [compiler/fn_compile.go](compiler/fn_compile.go)
- VM Exec: [vm/vm.go](vm/vm.go)

### `awt`
**Means:** Await async function call.
- Lexer: [lexer/token.go:33](lexer/token.go#L33)
- AST: [ast/](ast/)
- Parser: [parser/stmt/awt.go](parser/stmt/awt.go)
- Compiler: [compiler/compiler.go](compiler/compiler.go)
- VM Exec: [vm/vm.go](vm/vm.go)

### `cl`
**Means:** Class declaration.
- Lexer: [lexer/token.go:38](lexer/token.go#L38)
- AST: [ast/](ast/)
- Parser: [parser/stmt/cl.go](parser/stmt/cl.go)
- Compiler: [compiler/stmt_decl.go](compiler/stmt_decl.go)
- VM Exec: [vm/vm.go](vm/vm.go)

### `cns`
**Means:** Constant declaration.
- Lexer: [lexer/token.go:20](lexer/token.go#L20)
- AST: [ast/](ast/)
- Parser: [parser/stmt/cns.go](parser/stmt/cns.go)
- Compiler: [compiler/stmt_vars.go](compiler/stmt_vars.go)
- VM Exec: [vm/vm.go](vm/vm.go)

### `con`
**Means:** Continue to next loop iteration.
- Lexer: [lexer/token.go:31](lexer/token.go#L31)
- AST: [ast/](ast/)
- Parser: [parser/stmt/cond.go](parser/stmt/cond.go)
- Compiler: [compiler/stmt_funcs.go](compiler/stmt_funcs.go)
- VM Exec: [vm/vm.go](vm/vm.go)

### `done`
**Means:** Break out of a loop.
- Lexer: [lexer/token.go:30](lexer/token.go#L30)
<<<<<<< HEAD
- AST: [ast/](ast/)
- Parser: [parser/stmt/done.go](parser/stmt/done.go)
- Compiler: [compiler/stmt_funcs.go](compiler/stmt_funcs.go)
=======
- AST: [ast/ast.go](ast/ast.go)
- Parser: [parser/parser.go:165](parser/parser.go#L165)
- Compiler: [compiler/compiler.go:233](compiler/compiler.go#L233)
### `cns`
**Means:** Constant declaration.
- Lexer: [lexer/token.go:159](lexer/token.go#L159)
- AST: [ast/ast.go](ast/ast.go)
- Parser: [parser/parser.go:146](parser/parser.go#L146)
- Compiler: [compiler/compiler.go](compiler/compiler.go)
- VM Exec: [vm/vm.go](vm/vm.go)

### `t`
**Means:** Boolean true.
- Lexer: [lexer/token.go:160](lexer/token.go#L160)
- AST: [ast/ast.go](ast/ast.go)
- Parser: [parser/parser.go:175](parser/parser.go#L175)
- Compiler: [compiler/compiler.go:252](compiler/compiler.go#L252)
- VM Exec: [vm/vm.go:196](vm/vm.go#L196)

### `f`
**Means:** Boolean false.
- Lexer: [lexer/token.go:161](lexer/token.go#L161)
- AST: [ast/ast.go](ast/ast.go)
- Parser: [parser/parser.go:148](parser/parser.go#L148)
- Compiler: [compiler/compiler.go:201](compiler/compiler.go#L201)
- VM Exec: [vm/vm.go:198](vm/vm.go#L198)

### `fn`
**Means:** Function declaration.
- Lexer: [lexer/token.go:162](lexer/token.go#L162)
- AST: [ast/ast.go](ast/ast.go)
- Parser: [parser/parser.go:148](parser/parser.go#L148)
- Compiler: [compiler/compiler.go:201](compiler/compiler.go#L201)
- VM Exec: [vm/vm.go](vm/vm.go)

### `rev`
**Means:** Return/exit statement.
- Lexer: [lexer/token.go:163](lexer/token.go#L163)
- AST: [ast/ast.go](ast/ast.go)
- Parser: [parser/parser.go:157](parser/parser.go#L157)
- Compiler: [compiler/compiler.go](compiler/compiler.go)
- VM Exec: [vm/vm.go](vm/vm.go)

### `if`
**Means:** If conditional.
- Lexer: [lexer/token.go:164](lexer/token.go#L164)
- AST: [ast/ast.go](ast/ast.go)
- Parser: [parser/parser.go:159](parser/parser.go#L159)
- Compiler: [compiler/compiler.go:204](compiler/compiler.go#L204)
- VM Exec: [vm/vm.go](vm/vm.go)

### `elif`
**Means:** Else-if conditional.
- Lexer: [lexer/token.go:165](lexer/token.go#L165)
- AST: [ast/ast.go](ast/ast.go)
- Parser: [parser/parser.go](parser/parser.go)
- Compiler: [compiler/compiler.go](compiler/compiler.go)
>>>>>>> bb7a0c19659becbf257b6f9be22e9f5969a929b6
- VM Exec: [vm/vm.go](vm/vm.go)

### `el`
**Means:** Else conditional.
- Lexer: [lexer/token.go:26](lexer/token.go#L26)
<<<<<<< HEAD
- AST: [ast/](ast/)
- Parser: [parser/stmt/if.go](parser/stmt/if.go)
- Compiler: [compiler/stmt_control.go](compiler/stmt_control.go)
=======
- Lexer: [lexer/token.go:166](lexer/token.go#L166)
- AST: [ast/ast.go](ast/ast.go)
- Parser: [parser/parser.go](parser/parser.go)
- Compiler: [compiler/compiler.go](compiler/compiler.go)
>>>>>>> bb7a0c19659becbf257b6f9be22e9f5969a929b6
- VM Exec: [vm/vm.go](vm/vm.go)

### `elif`
**Means:** Else-if conditional.
- Lexer: [lexer/token.go:26](lexer/token.go#L26)
- AST: [ast/](ast/)
- Parser: [parser/stmt/if.go](parser/stmt/if.go)
- Compiler: [compiler/stmt_control.go](compiler/stmt_control.go)
- VM Exec: [vm/vm.go](vm/vm.go)

### `err`
**Means:** Catch/error block.
- Lexer: [lexer/token.go:35](lexer/token.go#L35)
- AST: [ast/](ast/)
- Parser: [parser/stmt/err.go](parser/stmt/err.go)
- Compiler: [compiler/compiler.go](compiler/compiler.go)
- VM Exec: [vm/vm.go](vm/vm.go)
### `f`
**Means:** Boolean false.
- Lexer: [lexer/token.go:22](lexer/token.go#L22)
- AST: [ast/](ast/)
- Parser: [parser/expr/bool.go](parser/expr/bool.go)
- Compiler: [compiler/expr.go](compiler/expr.go)
- VM Exec: [vm/vm.go](vm/vm.go)

### `fn`
**Means:** Function declaration.
- Lexer: [lexer/token.go:23](lexer/token.go#L23)
- AST: [ast/](ast/)
- Parser: [parser/stmt/fn.go](parser/stmt/fn.go)
- Compiler: [compiler/fn_compile.go](compiler/fn_compile.go)
- VM Exec: [vm/vm.go](vm/vm.go)

### `if`
**Means:** If conditional.
- Lexer: [lexer/token.go:25](lexer/token.go#L25)
- AST: [ast/](ast/)
- Parser: [parser/stmt/if.go](parser/stmt/if.go)
- Compiler: [compiler/stmt_control.go](compiler/stmt_control.go)
- VM Exec: [vm/vm.go](vm/vm.go)

### `lp`
**Means:** Loop block.
- Lexer: [lexer/token.go:29](lexer/token.go#L29)
- AST: [ast/](ast/)
- Parser: [parser/stmt/lp.go](parser/stmt/lp.go)
- Compiler: [compiler/stmt_control.go](compiler/stmt_control.go)
- VM Exec: [vm/vm.go](vm/vm.go)

### `mul`
**Means:** Spawn multiple async tasks.
- Lexer: [lexer/token.go:40](lexer/token.go#L40)
- AST: [ast/](ast/)
- Parser: [parser/stmt/async_call.go](parser/stmt/async_call.go)
- Compiler: [compiler/stmt_funcs.go](compiler/stmt_funcs.go)
- VM Exec: [vm/op_calls.go](vm/op_calls.go)

### `on`
**Means:** Switch/case block.
- Lexer: [lexer/token.go:28](lexer/token.go#L28)
<<<<<<< HEAD
- AST: [ast/](ast/)
- Parser: [parser/stmt/on.go](parser/stmt/on.go)
- Compiler: [compiler/stmt_control.go](compiler/stmt_control.go)
=======
- AST: [ast/ast.go](ast/ast.go)
- Parser: [parser/parser.go:161](parser/parser.go#L161)
- Compiler: [compiler/compiler.go:207](compiler/compiler.go#L207)
### `on`
**Means:** Switch/case block.
- Lexer: [lexer/token.go:167](lexer/token.go#L167)
- AST: [ast/ast.go](ast/ast.go)
- Parser: [parser/parser.go:161](parser/parser.go#L161)
- Compiler: [compiler/compiler.go:207](compiler/compiler.go#L207)
- VM Exec: [vm/vm.go](vm/vm.go)

### `lp`
**Means:** Loop block.
- Lexer: [lexer/token.go:168](lexer/token.go#L168)
- AST: [ast/ast.go](ast/ast.go)
- Parser: [parser/parser.go:38](parser/parser.go#L38)
- Compiler: [compiler/compiler.go](compiler/compiler.go)
- VM Exec: [vm/vm.go](vm/vm.go)

### `done`
**Means:** Break out of a loop.
- Lexer: [lexer/token.go:169](lexer/token.go#L169)
- AST: [ast/ast.go](ast/ast.go)
- Parser: [parser/parser.go:165](parser/parser.go#L165)
- Compiler: [compiler/compiler.go](compiler/compiler.go)
- VM Exec: [vm/vm.go](vm/vm.go)

### `con`
**Means:** Continue to next loop iteration.
- Lexer: [lexer/token.go:170](lexer/token.go#L170)
- AST: [ast/ast.go](ast/ast.go)
- Parser: [parser/parser.go:167](parser/parser.go#L167)
- Compiler: [compiler/compiler.go](compiler/compiler.go)
- VM Exec: [vm/vm.go:163](vm/vm.go#L163)

### `asn`
**Means:** Async function modifier.
- Lexer: [lexer/token.go:171](lexer/token.go#L171)
- AST: [ast/ast.go](ast/ast.go)
- Parser: [parser/parser.go:150](parser/parser.go#L150)
- Compiler: [compiler/compiler.go](compiler/compiler.go)
- VM Exec: [vm/vm.go](vm/vm.go)

### `awt`
**Means:** Await async function call.
- Lexer: [lexer/token.go:172](lexer/token.go#L172)
- AST: [ast/ast.go](ast/ast.go)
- Parser: [parser/parser.go:173](parser/parser.go#L173)
- Compiler: [compiler/compiler.go](compiler/compiler.go)
- VM Exec: [vm/vm.go](vm/vm.go)

### `try`
**Means:** Try block for error handling.
- Lexer: [lexer/token.go:173](lexer/token.go#L173)
- AST: [ast/ast.go](ast/ast.go)
- Parser: [parser/parser.go:175](parser/parser.go#L175)
- Compiler: [compiler/compiler.go:252](compiler/compiler.go#L252)
- VM Exec: [vm/vm.go:385](vm/vm.go#L385)

### `err`
**Means:** Catch/error block.
- Lexer: [lexer/token.go:174](lexer/token.go#L174)
- AST: [ast/ast.go](ast/ast.go)
- Parser: [parser/parser.go:177](parser/parser.go#L177)
- Compiler: [compiler/compiler.go](compiler/compiler.go)
>>>>>>> bb7a0c19659becbf257b6f9be22e9f5969a929b6
- VM Exec: [vm/vm.go](vm/vm.go)

### `pv`
**Means:** Private access modifier.
- Lexer: [lexer/token.go:36](lexer/token.go#L36)
<<<<<<< HEAD
- AST: [ast/](ast/)
- Parser: [parser/stmt/pv.go](parser/stmt/pv.go)
=======
- Lexer: [lexer/token.go:175](lexer/token.go#L175)
- AST: [ast/ast.go](ast/ast.go)
- Parser: [parser/parser.go:181](parser/parser.go#L181)
>>>>>>> bb7a0c19659becbf257b6f9be22e9f5969a929b6
- Compiler: [compiler/compiler.go](compiler/compiler.go)
- VM Exec: [vm/vm.go](vm/vm.go)

### `rev`
**Means:** Return/exit statement.
- Lexer: [lexer/token.go:24](lexer/token.go#L24)
- AST: [ast/](ast/)
- Parser: [parser/stmt/rev.go](parser/stmt/rev.go)
- Compiler: [compiler/stmt_funcs.go](compiler/stmt_funcs.go)
- VM Exec: [vm/vm.go](vm/vm.go)

### `t`
**Means:** Boolean true.
- Lexer: [lexer/token.go:21](lexer/token.go#L21)
- AST: [ast/](ast/)
- Parser: [parser/expr/bool.go](parser/expr/bool.go)
- Compiler: [compiler/expr.go](compiler/expr.go)
- VM Exec: [vm/vm.go](vm/vm.go)

### `try`
**Means:** Try block for error handling.
- Lexer: [lexer/token.go:34](lexer/token.go#L34)
- AST: [ast/](ast/)
- Parser: [parser/stmt/t.go](parser/stmt/t.go)
- Compiler: [compiler/stmt_errors.go](compiler/stmt_errors.go)
- VM Exec: [vm/op/op_errors_try.go](vm/op/op_errors_try.go)

### `uni`
**Means:** Single async task.
- Lexer: [lexer/token.go:41](lexer/token.go#L41)
- AST: [ast/](ast/)
- Parser: [parser/stmt/async_call.go](parser/stmt/async_call.go)
- Compiler: [compiler/stmt_funcs.go](compiler/stmt_funcs.go)
- VM Exec: [vm/vm.go](vm/vm.go)

### `unknown`
**Means:** Unknown boolean literal.
- Lexer: [lexer/token.go:39](lexer/token.go#L39)
- AST: [ast/](ast/)
- Parser: [parser/expr/unknown.go](parser/expr/unknown.go)
- Compiler: [compiler/expr.go](compiler/expr.go)
- VM Exec: [vm/execute.go](vm/execute.go)

### `use`
**Means:** Import file/lib/package.
- Lexer: [lexer/token.go:37](lexer/token.go#L37)
<<<<<<< HEAD
- AST: [ast/](ast/)
- Parser: [parser/stmt/use.go](parser/stmt/use.go)
=======
- AST: [ast/ast.go](ast/ast.go)
- Parser: [parser/parser.go:179](parser/parser.go#L179)
- Compiler: [compiler/compiler.go:261](compiler/compiler.go#L261)
### `use`
**Means:** Import/use module.
- Lexer: [lexer/token.go:176](lexer/token.go#L176)
- AST: [ast/ast.go](ast/ast.go)
- Parser: [parser/parser.go:179](parser/parser.go#L179)
- Compiler: [compiler/compiler.go](compiler/compiler.go)
- VM Exec: [vm/vm.go](vm/vm.go)

### `cl`
**Means:** Clear output/console.
- Lexer: [lexer/token.go:177](lexer/token.go#L177)
- AST: [ast/ast.go](ast/ast.go)
- Parser: [parser/parser.go:183](parser/parser.go#L183)
- Compiler: [compiler/compiler.go](compiler/compiler.go)
- VM Exec: [vm/vm.go:367](vm/vm.go#L367)

### `unknown`
**Means:** Unknown boolean literal.
- Lexer: [lexer/token.go:178](lexer/token.go#L178)
- AST: [ast/ast.go](ast/ast.go)
- Parser: [parser/parser.go](parser/parser.go)
- Compiler: [compiler/compiler.go](compiler/compiler.go)
- VM Exec: [vm/vm.go:200](vm/vm.go#L200)

### `mul`
**Means:** Spawn multiple async tasks.
- Lexer: [lexer/token.go:179](lexer/token.go#L179)
- AST: [ast/ast.go](ast/ast.go)
- Parser: [parser/parser.go:169](parser/parser.go#L169)
- Compiler: [compiler/compiler.go](compiler/compiler.go)
- VM Exec: [vm/vm.go:213](vm/vm.go#L213)

### `uni`
**Means:** Unidirectional detached async execution.
- Lexer: [lexer/token.go:180](lexer/token.go#L180)
- AST: [ast/ast.go](ast/ast.go)
- Parser: [parser/parser.go:171](parser/parser.go#L171)
>>>>>>> bb7a0c19659becbf257b6f9be22e9f5969a929b6
- Compiler: [compiler/compiler.go](compiler/compiler.go)
- VM Exec: [vm/vm.go](vm/vm.go)

## Built-ins

*(Note: Built-ins are evaluated natively by the compiler via `OpBuiltin`, therefore AST and Parser nodes are generic `CallExpr` in those stages.)*

### `abs`
**Means:** Absolute value of a number.
- Core: [core/types.go:12](core/types.go#L12)
- VM Exec: [handler/core/abs.go:8](handler/core/abs.go#L8)

### `add`
**Means:** Append item to array.
- Core: [core/types.go:9](core/types.go#L9)
- VM Exec: [handler/core/add.go:7](handler/core/add.go#L7)

### `arg`
**Means:** Get CLI arguments.
- Core: [core/types.go:36](core/types.go#L36)
- VM Exec: [handler/core/arg.go:8](handler/core/arg.go#L8)

### `cap`
**Means:** Capitalize string (Title Case).
- Core: [core/types.go:16](core/types.go#L16)
- VM Exec: [handler/core/cap.go:8](handler/core/cap.go#L8)

### `cmd`
**Means:** Execute shell command.
- Core: [core/types.go:38](core/types.go#L38)
- VM Exec: [handler/core/cmd.go:9](handler/core/cmd.go#L9)

### `cron`
**Means:** Scheduled background tasks.
- Core: [core/types.go:56](core/types.go#L56)
- VM Exec: [handler/concurrency/cron.go:159](handler/concurrency/cron.go#L159)

### `date`
**Means:** Format timestamp to date string.
- Core: [core/types.go:33](core/types.go#L33)
- VM Exec: [handler/core/date.go:8](handler/core/date.go#L8)

### `db`
**Means:** Database operations entrypoint.
- Core: [core/types.go:43](core/types.go#L43)
- VM Exec: [handler/db/db.go:28](handler/db/db.go#L28)

### `dbpool`
**Means:** Database connection pool configuration.
- Core: [core/types.go:73](core/types.go#L73)
- VM Exec: [handler/db/dbpool.go:9](handler/db/dbpool.go#L9)

### `del`
**Means:** Delete variable/reference.
- Core: [core/types.go:40](core/types.go#L40)
- VM Exec: [handler/core/del.go:8](handler/core/del.go#L8)

### `die`
**Means:** Exit program with error code.
- Core: [core/types.go:41](core/types.go#L41)
- VM Exec: [handler/core/die.go:9](handler/core/die.go#L9)

### `dir`
**Means:** List directory contents.
- Core: [core/types.go:39](core/types.go#L39)
- VM Exec: [handler/core/dir.go:8](handler/core/dir.go#L8)

### `doc`
**Means:** Document generation.
- Core: [core/types.go:63](core/types.go#L63)
- VM Exec: [handler/data/doc.go:10](handler/data/doc.go#L10)

### `enc`
**Means:** Encoding functions (e.g. b64).
- Core: [core/types.go:48](core/types.go#L48)
- VM Exec: [handler/core/enc.go:14](handler/core/enc.go#L14)

### `env`
**Means:** Get environment variables.
- Core: [core/types.go:37](core/types.go#L37)
- VM Exec: [handler/core/env.go:8](handler/core/env.go#L8)

### `flag`
**Means:** CLI flag parsing.
- Core: [core/types.go:61](core/types.go#L61)
- VM Exec: [handler/state/flag.go:16](handler/state/flag.go#L16)

### `flow`
**Means:** Flow control.
- Core: [core/types.go:71](core/types.go#L71)
- VM Exec: [handler/state/flow.go:18](handler/state/flow.go#L18)

### `fmt`
**Means:** String formatting.
- Core: [core/types.go:51](core/types.go#L51)
- VM Exec: [handler/core/fmt.go:9](handler/core/fmt.go#L9)

### `geo`
**Means:** Geolocation functions.
- Core: [core/types.go:70](core/types.go#L70)
- VM Exec: [handler/data/geo.go:11](handler/data/geo.go#L11)

### `get`
**Means:** Pull characteristic/type info of a value.
- Core: [core/types.go:8](core/types.go#L8)
- VM Exec: [handler/core/get.go:7](handler/core/get.go#L7)

### `has`
**Means:** Check if string contains substring or map has key.
- Core: [core/types.go:22](core/types.go#L22)
- VM Exec: [handler/core/has.go:8](handler/core/has.go#L8)

### `hash`
**Means:** Hashing functions (e.g. md5, sha256).
- Core: [core/types.go:47](core/types.go#L47)
- VM Exec: [handler/core/hash.go:14](handler/core/hash.go#L14)

### `img`
**Means:** Image processing.
- Core: [core/types.go:62](core/types.go#L62)
- VM Exec: [handler/data/img.go:29](handler/data/img.go#L29)

### `in`
**Means:** Read input from standard input.
- Core: [core/types.go:45](core/types.go#L45)
- VM Exec: [handler/core/in.go:10](handler/core/in.go#L10)

### `j`
**Means:** Join array of strings.
- Core: [core/types.go:20](core/types.go#L20)
- VM Exec: [handler/core/j.go:7](handler/core/j.go#L7)

### `job`
**Means:** Background job queuing.
- Core: [core/types.go:57](core/types.go#L57)
- VM Exec: [handler/concurrency/job.go:30](handler/concurrency/job.go#L30)

### `json`
**Means:** Parse JSON string to object, or stringify object to JSON.
- Core: [core/types.go:35](core/types.go#L35)
- VM Exec: [handler/data/json.go:8](handler/data/json.go#L8)

### `jwt`
**Means:** JWT generation and validation.
- Core: [core/types.go:49](core/types.go#L49)
- VM Exec: [handler/core/jwt.go:11](handler/core/jwt.go#L11)

### `key`
**Means:** Get keys of a map.
- Core: [core/types.go:26](core/types.go#L26)
- VM Exec: [handler/core/key.go:7](handler/core/key.go#L7)

### `len`
**Means:** Get length of string, array, or map.
- Core: [core/types.go:7](core/types.go#L7)
- VM Exec: [handler/core/len.go:7](handler/core/len.go#L7)

### `log`
**Means:** Structured logging.
- Core: [core/types.go:74](core/types.go#L74)
- VM Exec: [handler/system/log.go:10](handler/system/log.go#L10)


### `low`
**Means:** Lowercase string.
- Core: [core/types.go:17](core/types.go#L17)
- VM Exec: [handler/core/low.go:8](handler/core/low.go#L8)

### `mail`
**Means:** Email sending.
- Core: [core/types.go:65](core/types.go#L65)
- VM Exec: [handler/messaging/mail.go:10](handler/messaging/mail.go#L10)

### `math`
**Means:** Math operations module/namespace.
- Core: [core/types.go:44](core/types.go#L44)
- VM Exec: [handler/core/math.go:8](handler/core/math.go#L8)

### `max`
**Means:** Maximum of two numbers.
- Core: [core/types.go:14](core/types.go#L14)
- VM Exec: [handler/core/max.go:8](handler/core/max.go#L8)

### `mem`
**Means:** Memory and caching operations.
- Core: [core/types.go:66](core/types.go#L66)
- VM Exec: [handler/state/mem.go:36](handler/state/mem.go#L36)

### `met`
**Means:** Metrics collection.
- Core: [core/types.go:69](core/types.go#L69)
- VM Exec: [handler/system/met.go:9](handler/system/met.go#L9)

### `min`
**Means:** Minimum of two numbers.
- Core: [core/types.go:13](core/types.go#L13)
- VM Exec: [handler/core/min.go:8](handler/core/min.go#L8)

### `mod`
**Means:** Modify.
- Core: [core/types.go:21](core/types.go#L21)
- VM Exec: [handler/core/mod.go:8](handler/core/mod.go#L8)

### `now`
**Means:** Get current timestamp.
- Core: [core/types.go:32](core/types.go#L32)
- VM Exec: [handler/core/now.go:8](handler/core/now.go#L8)

### `num`
**Means:** Convert string to number.
- Core: [core/types.go:10](core/types.go#L10)
- VM Exec: [handler/core/num.go:8](handler/core/num.go#L8)

### `op`
**Means:** Open file or stream.
- Core: [core/types.go:42](core/types.go#L42)
- VM Exec: [vm/builtin_server.go](vm/builtin_server.go)

### `pipe`
**Means:** Data streaming and piping.
- Core: [core/types.go:55](core/types.go#L55)
- VM Exec: [handler/concurrency/pipe.go:8](handler/concurrency/pipe.go#L8)

### `pop`
**Means:** Pop last item from array.
- Core: [core/types.go:24](core/types.go#L24)
- VM Exec: [handler/core/pop.go:7](handler/core/pop.go#L7)

### `pt`
**Means:** Print to standard output.
- Core: [core/types.go:46](core/types.go#L46)
- VM Exec: [handler/core/pt.go:8](handler/core/pt.go#L8)

### `q`
**Means:** Quit/exit program forcefully.
- Core: [core/types.go:29](core/types.go#L29)
- VM Exec: [handler/core/q.go:8](handler/core/q.go#L8)

### `r`
**Means:** Read file contents.
- Core: [core/types.go:15](core/types.go#L15)
- VM Exec: [handler/core/r.go:8](handler/core/r.go#L8)

### `ran`
**Means:** Generate random float between 0 and 1.
- Core: [core/types.go:28](core/types.go#L28)
- VM Exec: [handler/core/ran.go:8](handler/core/ran.go#L8)

### `rate`
**Means:** Rate limiting operations.
- Core: [core/types.go:58](core/types.go#L58)
- VM Exec: [handler/concurrency/rate.go:37](handler/concurrency/rate.go#L37)

### `req`
**Means:** Make HTTP request.
- Core: [core/types.go:34](core/types.go#L34)
- VM Exec: [handler/http/req.go:11](handler/http/req.go#L11)

### `rgx`
**Means:** Regular expression operations.
- Core: [core/types.go:50](core/types.go#L50)
- VM Exec: [handler/core/regex.go:9](handler/core/regex.go#L9)

### `rm`
**Means:** Remove item from array by index, or delete file by path.
- Core: [core/types.go:25](core/types.go#L25)
- VM Exec: [handler/core/rm.go:8](handler/core/rm.go#L8)

### `rnd`
**Means:** Round number to nearest integer.
- Core: [core/types.go:15](core/types.go#L15)
- VM Exec: [handler/core/rnd.go:8](handler/core/rnd.go#L8)

### `rpc`
**Means:** RPC communications.
- Core: [core/types.go:68](core/types.go#L68)
- VM Exec: [handler/messaging/rpc.go:13](handler/messaging/rpc.go#L13)

### `rt`
**Means:** Realtime routing.
- Core: [core/types.go:53](core/types.go#L53)
- VM Exec: [handler/http/router.go:20](handler/http/router.go#L20)

### `sess`
**Means:** Session management.
- Core: [core/types.go:59](core/types.go#L59)
- VM Exec: [handler/state/session.go:36](handler/state/session.go#L36)

### `sort`
**Means:** Sort array.
- Core: [core/types.go:23](core/types.go#L23)
- VM Exec: [handler/core/sort.go:8](handler/core/sort.go#L8)

### `spl`
**Means:** Split string by separator.
- Core: [core/types.go:19](core/types.go#L19)
- VM Exec: [handler/core/spl.go:8](handler/core/spl.go#L8)

### `str`
**Means:** Convert number to string.
- Core: [core/types.go:11](core/types.go#L11)
- VM Exec: [handler/core/str.go:7](handler/core/str.go#L7)

### `sys`
**Means:** System and OS operations.
- Core: [core/types.go:54](core/types.go#L54)
- VM Exec: [handler/system/sys.go:12](handler/system/sys.go#L12)

### `test`
**Means:** Testing utilities.
- Core: [core/types.go:72](core/types.go#L72)
- VM Exec: [handler/system/test.go:10](handler/system/test.go#L10)

### `tmpl`
**Means:** Template rendering.
- Core: [core/types.go:64](core/types.go#L64)
- VM Exec: [handler/data/tmpl.go:43](handler/data/tmpl.go#L43)

### `trm`
**Means:** Trim whitespace from string.
- Core: [core/types.go:18](core/types.go#L18)
- VM Exec: [handler/core/trm.go:8](handler/core/trm.go#L8)

### `val`
**Means:** Get values of a map.
- Core: [core/types.go:27](core/types.go#L27)
- VM Exec: [handler/core/val.go:7](handler/core/val.go#L7)

### `valid`
**Means:** Data validation.
- Core: [core/types.go:52](core/types.go#L52)
- VM Exec: [handler/data/valid.go:10](handler/data/valid.go#L10)

### `w`
**Means:** Write data to file.
- Core: [core/types.go:31](core/types.go#L31)
- VM Exec: [handler/core/w.go:8](handler/core/w.go#L8)

### `ws`
**Means:** WebSockets support.
- Core: [core/types.go:67](core/types.go#L67)
<<<<<<< HEAD
- VM Exec: [handler/http/ws.go:19](handler/http/ws.go#L19)
=======
- VM Exec: [vm/iohandler/ws.go:19](vm/iohandler/ws.go#L19)
### `len`
**Means:** Get length of string, array, or map.
- Core: [core/types.go:51](core/types.go#L51)
- VM Exec: [vm/vm.go:485](vm/vm.go#L485)

### `get`
**Means:** Get type of a value as string.
- Core: [core/types.go:52](core/types.go#L52)
- VM Exec: [vm/vm.go:500](vm/vm.go#L500)

### `add`
**Means:** Append item to array.
- Core: [core/types.go:53](core/types.go#L53)
- VM Exec: [vm/vm.go:506](vm/vm.go#L506)

### `num`
**Means:** Convert string to number.
- Core: [core/types.go:54](core/types.go#L54)
- VM Exec: [vm/vm.go:519](vm/vm.go#L519)

### `str`
**Means:** Convert number to string.
- Core: [core/types.go:55](core/types.go#L55)
- VM Exec: [vm/vm.go:529](vm/vm.go#L529)

### `abs`
**Means:** Absolute value of a number.
- Core: [core/types.go:56](core/types.go#L56)
- VM Exec: [vm/vm.go:535](vm/vm.go#L535)

### `min`
**Means:** Minimum of two numbers.
- Core: [core/types.go:57](core/types.go#L57)
- VM Exec: [vm/vm.go:541](vm/vm.go#L541)

### `max`
**Means:** Maximum of two numbers.
- Core: [core/types.go:58](core/types.go#L58)
- VM Exec: [vm/vm.go:547](vm/vm.go#L547)

### `rnd`
**Means:** Round number to nearest integer.
- Core: [core/types.go:59](core/types.go#L59)
- VM Exec: [vm/vm.go:553](vm/vm.go#L553)

### `cap`
**Means:** Capitalize string (Title Case).
- Core: [core/types.go:60](core/types.go#L60)
- VM Exec: [vm/vm.go:559](vm/vm.go#L559)

### `low`
**Means:** Lowercase string.
- Core: [core/types.go:61](core/types.go#L61)
- VM Exec: [vm/vm.go:565](vm/vm.go#L565)

### `trm`
**Means:** Trim whitespace from string.
- Core: [core/types.go:62](core/types.go#L62)
- VM Exec: [vm/vm.go:571](vm/vm.go#L571)

### `spl`
**Means:** Split string by separator.
- Core: [core/types.go:63](core/types.go#L63)
- VM Exec: [vm/vm.go:577](vm/vm.go#L577)

### `j`
**Means:** Join array of strings.
- Core: [core/types.go:64](core/types.go#L64)
- VM Exec: [vm/vm.go:588](vm/vm.go#L588)

### `mod`
**Means:** Modulo operation.
- Core: [core/types.go:65](core/types.go#L65)
- VM Exec: [vm/vm.go:599](vm/vm.go#L599)

### `has`
**Means:** Check if string contains substring or map has key.
- Core: [core/types.go:66](core/types.go#L66)
- VM Exec: [vm/vm.go:605](vm/vm.go#L605)

### `sort`
**Means:** Sort array.
- Core: [core/types.go:67](core/types.go#L67)
- VM Exec: [vm/vm.go:636](vm/vm.go#L636)

### `pop`
**Means:** Pop last item from array.
- Core: [core/types.go:68](core/types.go#L68)
- VM Exec: [vm/vm.go:650](vm/vm.go#L650)

### `rm`
**Means:** Remove item from array by index, or delete file by path.
- Core: [core/types.go:69](core/types.go#L69)
- VM Exec: [vm/vm.go:660](vm/vm.go#L660)

### `key`
**Means:** Get keys of a map.
- Core: [core/types.go:70](core/types.go#L70)
- VM Exec: [vm/vm.go:683](vm/vm.go#L683)

### `val`
**Means:** Get values of a map.
- Core: [core/types.go:71](core/types.go#L71)
- VM Exec: [vm/vm.go:694](vm/vm.go#L694)

### `ran`
**Means:** Generate random float between 0 and 1.
- Core: [core/types.go:72](core/types.go#L72)
- VM Exec: [vm/vm.go:705](vm/vm.go#L705)

### `q`
**Means:** Quit/exit program forcefully.
- Core: [core/types.go:73](core/types.go#L73)
- VM Exec: [vm/vm.go:708](vm/vm.go#L708)

### `r`
**Means:** Read file contents.
- Core: [core/types.go:74](core/types.go#L74)
- VM Exec: [vm/vm.go:553](vm/vm.go#L553)

### `w`
**Means:** Write data to file.
- Core: [core/types.go:75](core/types.go#L75)
- VM Exec: [vm/vm.go:726](vm/vm.go#L726)

### `now`
**Means:** Get current timestamp.
- Core: [core/types.go:76](core/types.go#L76)
- VM Exec: [vm/vm.go:954](vm/vm.go#L954)

### `date`
**Means:** Format timestamp to date string.
- Core: [core/types.go:77](core/types.go#L77)
- VM Exec: [vm/vm.go:957](vm/vm.go#L957)

### `req`
**Means:** Make HTTP request.
- Core: [core/types.go:78](core/types.go#L78)
- VM Exec: [vm/vm.go:969](vm/vm.go#L969)

### `json`
**Means:** Parse JSON string to object, or stringify object to JSON.
- Core: [core/types.go:79](core/types.go#L79)
- VM Exec: [vm/vm.go:976](vm/vm.go#L976)

### `arg`
**Means:** Get CLI arguments.
- Core: [core/types.go:80](core/types.go#L80)
- VM Exec: [vm/vm.go:998](vm/vm.go#L998)

### `env`
**Means:** Get environment variables.
- Core: [core/types.go:81](core/types.go#L81)
- VM Exec: [vm/vm.go:1012](vm/vm.go#L1012)

### `cmd`
**Means:** Execute shell command.
- Core: [core/types.go:82](core/types.go#L82)
- VM Exec: [vm/vm.go:1018](vm/vm.go#L1018)

### `dir`
**Means:** List directory contents.
- Core: [core/types.go:83](core/types.go#L83)
- VM Exec: [vm/vm.go:1039](vm/vm.go#L1039)

### `del`
**Means:** Delete variable/reference.
- Core: [core/types.go:84](core/types.go#L84)
- VM Exec: [vm/vm.go:1053](vm/vm.go#L1053)

### `die`
**Means:** Exit program with error code.
- Core: [core/types.go:85](core/types.go#L85)
- VM Exec: [vm/vm.go:1063](vm/vm.go#L1063)

### `op`
**Means:** Open file or stream.
- Core: [core/types.go:86](core/types.go#L86)
- VM Exec: [vm/vm.go:736](vm/vm.go#L736)

### `db`
**Means:** Database operations entrypoint.
- Core: [core/types.go:87](core/types.go#L87)
- VM Exec: [vm/vm.go:863](vm/vm.go#L863)

### `math`
**Means:** Math operations module/namespace.
- Core: [core/types.go:88](core/types.go#L88)
- VM Exec: [vm/vm.go:1073](vm/vm.go#L1073)

### `in`
**Means:** Read input from standard input.
- Core: [core/types.go:89](core/types.go#L89)
- VM Exec: [vm/vm.go:1088](vm/vm.go#L1088)

### `pt`
**Means:** Print to standard output.
- Core: [core/types.go:90](core/types.go#L90)
- VM Exec: [vm/vm.go:1080](vm/vm.go#L1080)

>>>>>>> bb7a0c19659becbf257b6f9be22e9f5969a929b6

## Literals & Specials

### `EOF`
**Means:** End of file.
- Lexer: [lexer/token.go:7](lexer/token.go#L7)
- Parser: [parser/parser.go:118](parser/parser.go#L118)

### `ILLEGAL`
**Means:** Illegal character encountered.
- Lexer: [lexer/token.go:8](lexer/token.go#L8)

### `IDENT`
**Means:** Identifier (variable/function names).
- Lexer: [lexer/token.go:11](lexer/token.go#L11)
- AST: [ast/ident.go](ast/ident.go)
- Parser: [parser/expr/ident.go](parser/expr/ident.go)
- Compiler: [compiler/expr.go](compiler/expr.go)

### `NUMBER`
**Means:** Numeric literal (e.g. 42, 3.14).
- Lexer: [lexer/token.go:12](lexer/token.go#L12)
- AST: [ast/num.go](ast/num.go)
- Parser: [parser/expr/num.go](parser/expr/num.go)
- Compiler: [compiler/expr.go](compiler/expr.go)

### `STRING`
**Means:** String literal (e.g. "hello", 'hello').
- Lexer: [lexer/token.go:13](lexer/token.go#L13)
- AST: [ast/str.go](ast/str.go)
- Parser: [parser/expr/str.go](parser/expr/str.go)
- Compiler: [compiler/expr.go](compiler/expr.go)

### `RAW_STRING`
**Means:** Raw string literal without escapes.
- Lexer: [lexer/token.go:14](lexer/token.go#L14)
- AST: [ast/str.go](ast/str.go)
- Parser: [parser/expr/str.go](parser/expr/str.go)

### `STRING_PART`
**Means:** Segment of a string before or between interpolations.
- Lexer: [lexer/token.go:15](lexer/token.go#L15)
- AST: [ast/interp.go](ast/interp.go)
- Parser: [parser/expr/interp.go](parser/expr/interp.go)

### `INTERP_START` (`${`)
**Means:** Start of string interpolation expression.
- Lexer: [lexer/token.go:16](lexer/token.go#L16)
- Parser: [parser/expr/interp.go](parser/expr/interp.go)

### `INTERP_END` (`}`)
**Means:** End of string interpolation expression.
- Lexer: [lexer/token.go:17](lexer/token.go#L17)
- Parser: [parser/expr/interp.go](parser/expr/interp.go)


## Operators

### 1-Char Operators
- `=` (`TOKEN_ASSIGN`): Assignment. [lexer/token.go:44](lexer/token.go#L44)
- `+` (`TOKEN_PLUS`): Addition. [lexer/token.go:45](lexer/token.go#L45)
- `-` (`TOKEN_MINUS`): Subtraction. [lexer/token.go:46](lexer/token.go#L46)
- `*` (`TOKEN_STAR`): Multiplication. [lexer/token.go:47](lexer/token.go#L47)
- `/` (`TOKEN_SLASH`): Division. [lexer/token.go:48](lexer/token.go#L48)
- `%` (`TOKEN_PERCENT`): Modulo. [lexer/token.go:49](lexer/token.go#L49)
- `<` (`TOKEN_LT`): Less than. [lexer/token.go:50](lexer/token.go#L50)
- `>` (`TOKEN_GT`): Greater than. [lexer/token.go:51](lexer/token.go#L51)
- `&` (`TOKEN_AND`): Logical AND. [lexer/token.go:52](lexer/token.go#L52)
- `|` (`TOKEN_OR`): Logical OR. [lexer/token.go:53](lexer/token.go#L53)
- `!` (`TOKEN_NOT`): Logical NOT. [lexer/token.go:54](lexer/token.go#L54)
- `?` (`TOKEN_QUESTION`): Ternary / Optional. [lexer/token.go:55](lexer/token.go#L55)
- `?.` (`TOKEN_QMARK_DOT`): Safe navigation. [lexer/token.go:56](lexer/token.go#L56)
- `??` (`TOKEN_QQ`): Null coalescing. [lexer/token.go:57](lexer/token.go#L57)
- `.` (`TOKEN_DOT`): Property access / comment delimiter. [lexer/token.go:58](lexer/token.go#L58)

### 2-Char Operators
- `!=` (`TOKEN_NOT_EQ`): Not equal. [lexer/token.go:61](lexer/token.go#L61)
- `<=` (`TOKEN_LT_EQ`): Less than or equal. [lexer/token.go:62](lexer/token.go#L62)
- `>=` (`TOKEN_GT_EQ`): Greater than or equal. [lexer/token.go:63](lexer/token.go#L63)
- `->` (`TOKEN_ARROW`): Arrow function / return type mapping. [lexer/token.go:64](lexer/token.go#L64)
- `<-` (`TOKEN_LARROW`): Channel receive / backward mapping. [lexer/token.go:65](lexer/token.go#L65)

## Delimiters

- `(` (`TOKEN_LPAREN`): Open parenthesis. [lexer/token.go:68](lexer/token.go#L68)
- `)` (`TOKEN_RPAREN`): Close parenthesis. [lexer/token.go:69](lexer/token.go#L69)
- `{` (`TOKEN_LBRACE`): Open brace. [lexer/token.go:70](lexer/token.go#L70)
- `}` (`TOKEN_RBRACE`): Close brace. [lexer/token.go:71](lexer/token.go#L71)
- `[` (`TOKEN_LBRACKET`): Open bracket. [lexer/token.go:72](lexer/token.go#L72)
- `]` (`TOKEN_RBRACKET`): Close bracket. [lexer/token.go:73](lexer/token.go#L73)
- `,` (`TOKEN_COMMA`): Comma separator. [lexer/token.go:74](lexer/token.go#L74)
- `:` (`TOKEN_COLON`): Colon delimiter. [lexer/token.go:75](lexer/token.go#L75)
- `;` (`TOKEN_SEMICOLON`): Statement terminator. [lexer/token.go:76](lexer/token.go#L76)

## Error Codes
### Syntax & Parser Errors (E1XX)
- `E102` : Expected closing brace `}`. [parser/parser.go:78](parser/parser.go#L78)
- `E103` : Expected closing parenthesis `)`. [parser/parser.go:80](parser/parser.go#L80)
- `E104` : Expected closing bracket `]`. [parser/parser.go:82](parser/parser.go#L82)
- `E105` : Needs `fn` keyword before declaration. [parser/parser.go:146](parser/parser.go#L146)
- `E107` : Expected identifier or expression. [parser/parser.go:84](parser/parser.go#L84)
- `E108` : Needs closing character (unclosed quote, bracket, or string interp). [parser/parser.go:86](parser/parser.go#L86)
- `E109` : Illegal character encountered in lexer. [parser/parser.go:88](parser/parser.go#L88)
- `E110` : Bad number format. [compiler/expr.go](compiler/expr.go)

### Compiler Errors (E2XX)
- `E203` : Stray `done` (break) outside of a loop. [compiler/stmt_funcs.go](compiler/stmt_funcs.go)
- `E204` : Stray `con` (continue) outside of a loop. [compiler/stmt_funcs.go](compiler/stmt_funcs.go)

### Runtime & VM Errors (E3XX)
- `E300` : Generic runtime panic / VM execution error. [vm/runtime_error.go](vm/runtime_error.go)
- `E301` : Internal error for variable not found (usually caught and converted to E300). [vm/op/op_vars_getglobal.go](vm/op/op_vars_getglobal.go)

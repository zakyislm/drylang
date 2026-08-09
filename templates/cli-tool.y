// template: cli-tool
// desc: CLI tool with argument parsing and help

args = arg()

if len(args) = 0 {
    pt "Usage: y cli-tool.y <command> [args]"
    pt ""
    pt "Commands:"
    pt "  greet <name>    Say hello"
    pt "  upper <text>    Convert to uppercase"
    pt "  lower <text>    Convert to lowercase"
    pt "  count <text>    Count characters"
    pt "  reverse <text>  Reverse a string"
    die("")
}

command = args[0]

if command = "greet" {
    if len(args) < 2 { die("greet needs a name") }
    pt "Hello, ${args[1]}!"
} elif command = "upper" {
    if len(args) < 2 { die("upper needs text") }
    pt cap(args[1])
} elif command = "lower" {
    if len(args) < 2 { die("lower needs text") }
    pt low(args[1])
} elif command = "count" {
    if len(args) < 2 { die("count needs text") }
    pt "Length: ${len(args[1])}"
} elif command = "reverse" {
    if len(args) < 2 { die("reverse needs text") }
    text = args[1]
    chars = spl(text, "")
    result = ""
    idx = len(chars) - 1
    lp = {
        if idx < 0 { done }
        result = result + chars[idx]
        idx = idx - 1
    }
    pt result
} el {
    die("Unknown command: ${command}")
}

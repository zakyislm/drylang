// 10_strings.y - String built-ins

msg "  hello drylang  "

pt "Original: '" + msg + "'"
pt "Trimmed: '" + trm(msg) + "'"
pt "Capitalized: " + cap(trm(msg))
pt "Contains 'dry': " + str(has(msg, "dry"))

parts = spl("a,b,c", ",")
pt "Split 'a,b,c': " + str(parts)
pt "Join: " + j(parts, "-")
pt "Replace: " + mod("hello world", "world", "zaky")

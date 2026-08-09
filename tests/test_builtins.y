// Test all new builtins
pt "--- math ---"
pt math("sqrt", 16)
pt math("pow", 2, 10)
pt math("ceil", 4.1)
pt math("floor", 4.9)

pt "--- time ---"
pt now()
d = date()
pt d["format"]

pt "--- env ---"
pt env("USERNAME")

pt "--- dir ---"
pt dir(".")

pt "--- arg ---"
pt arg()

pt "--- json ---"
data = json("{\"name\":\"Zaky\",\"age\":17}")
pt data["name"]
pt data["age"]

pt "All builtins OK!"

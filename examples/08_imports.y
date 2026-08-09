// 08_imports.y - Modularity using 'use'

// This will load and run the variables example, merging its AST here
use "02_variables.y"

pt ""
pt "--- Imported Data ---"
// We can use the variables and constants defined in 02_variables.y!
pt "Imported PI: " + str(PI)
pt "Imported username: " + username

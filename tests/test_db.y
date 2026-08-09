// Connect to SQLite DB (creates file if not exists)
pt "Creating table..."
db("sqlite", "test.db", "CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, name TEXT, age INTEGER)")

pt "Inserting data..."
res = db("sqlite", "test.db", "INSERT INTO users (name, age) VALUES (?, ?)", "Zaky", 25)
pt res

pt "Querying data..."
users = db("sqlite", "test.db", "SELECT * FROM users")
pt users

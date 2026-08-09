// 12_database.y
// Demonstrates built-in SQLite database support using db()

DB_FILE "test_db.sqlite"

pt "--- Setup ---"
// Create a table
db("sqlite", DB_FILE, "CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, name TEXT, age INTEGER)")
pt "Table created."

pt "--- Insert ---"
// Insert some records (using parameterized queries to prevent SQL injection)
res1 db("sqlite", DB_FILE, "INSERT INTO users (name, age) VALUES (?, ?)", "Zaky", 17)
pt "Inserted Zaky, rows affected: " + str(res1["rows_affected"])

res2 db("sqlite", DB_FILE, "INSERT INTO users (name, age) VALUES (?, ?)", "Alice", 25)
pt "Inserted Alice, rows affected: " + str(res2["rows_affected"])

pt "--- Query ---"
// Query all users. Returns an array of maps.
users db("sqlite", DB_FILE, "SELECT * FROM users")
lp len(users) {
    pt "User " + str(users[i]["id"]) + ": " + users[i]["name"] + " (Age: " + str(users[i]["age"]) + ")"
}

pt "--- Update ---"
// Update a record
db("sqlite", DB_FILE, "UPDATE users SET age = ? WHERE name = ?", 18, "Zaky")
pt "Updated Zaky's age."

pt "--- Query with WHERE ---"
// Query with parameters
adults db("sqlite", DB_FILE, "SELECT * FROM users WHERE age >= ?", 18)
pt "Found " + str(len(adults)) + " adults:"
lp len(adults) {
    pt "- " + adults[i]["name"]
}

pt "--- Cleanup ---"
// Delete the test database file
del(DB_FILE)
pt "Cleaned up database file."

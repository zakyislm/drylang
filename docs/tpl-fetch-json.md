# Fetch Json

```javascript
// template: fetch-json
// desc: Fetches data from a REST API and parses the JSON response

pt "Fetching random joke..."

try {
     // Make HTTP GET request
    resp = req("https://official-joke-api.appspot.com/random_joke")
    
     // Parse string into dryLang Map
    data = json(resp)
    
     // Extract fields
    setup = data["setup"]
    punchline = data["punchline"]
    
     // Display
    pt "Q: ${setup}"
    q(2000) // wait 2 seconds for effect
    pt "A: ${punchline}"
    
} err(e) {
    pt "Failed to fetch joke: ${e}"
}

```

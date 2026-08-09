# Structs

```javascript
// 05_structs.y - Structs without keywords

// Define a struct structure simply by capitalizing
User {
    name
    age
    active
}

// Instantiate
zaky = User("Zaky", 17, t)

pt "Name: ${zaky.name}"
pt "Age: ${zaky.age}"
pt "Active: ${zaky.active}"

// Modify field
zaky.active = f
pt "Active now: ${zaky.active}"

```

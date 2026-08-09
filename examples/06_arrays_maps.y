// 06_arrays_maps.y - Arrays and Maps

// Arrays
fruits = ["apple", "banana", "cherry"]
pt "First fruit: ${fruits[0]}"
pt "Total fruits: ${len(fruits)}"

add(fruits, "mango")
pt "After add: ${fruits}"

pop(fruits)
pt "After pop: ${fruits}"

// Maps (Dictionaries)
user = {"id": 1, "role": "admin"}
pt "User role: " + user["role"]

pt "Keys: " + str(key(user))
pt "Values: " + str(val(user))

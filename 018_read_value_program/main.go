package main

import "fmt"

func main() {
	points := map[string]int{
		"a": 10,
		"b": 20,
	}
	fmt.Println("a", points["a"])
	fmt.Println("b", points["b"])
	fmt.Println("c", points["c"])

	valC, okC := points["b"]
	fmt.Println(valC, okC)
	valD, okD := points["c"]
	fmt.Println(valD, okD)

	prices := map[string]int{
		"xyx": 500,
		"dzy": 500,
	}

	total := 0

	for item, price := range prices {
		fmt.Println(item, price)
		total = total + price
	}

	fmt.Println(total)

}



// Bro, **maps in Go = objects/dictionaries in JavaScript**.
	
// Think:

// ```js
// const ages = {
//   sangam: 65,
//   jone: 75
// }
// ```

In Go:

```go
ages := map[string]int{
    "sangam": 65,
    "jone": 75,
}
```

---

# What is a Map?

A map stores data as:

```text
Key  -> Value

"name" -> "Abhi"
"age"  -> 22
```

Like a real dictionary:

```text
Word      Meaning
Apple  -> Fruit
Dog    -> Animal
```

Go map:

```go
map[keyType]valueType
```

Example:

```go
map[string]int
```

Means:

```text
Key    = string
Value  = int
```

---

# Your Example

```go
ages := map[string]int{
    "sangam": 65,
    "jone":   75,
}
```

Memory looks like:

```text
ages

"sangam" -> 65
"jone"   -> 75
```

---

# Accessing Values

```go
fmt.Println(ages["sangam"])
```

Output:

```text
65
```

Because:

```text
"sangam" -> 65
```

---

# len()

```go
fmt.Println(len(ages))
```

Output:

```text
2
```

Because map contains:

```text
1. sangam
2. jone
```

Total = 2

---

# This Line

```go
fmt.Println(ages["sangam"], len(ages))
```

Output:

```text
65 2
```

---

# What is make()?

Imagine:

### Variable Declaration

```go
var score map[string]int
```

This creates:

```text
score = nil
```

Like:

```text
score -> nothing
```

Visual:

```text
score
  |
 nil
```

---

# Check It

```go
var score map[string]int

fmt.Println(score)
```

Output:

```text
map[]
```

Actually map is nil internally.

---

# Reading from Nil Map

```go
fmt.Println(score["math"])
```

Output:

```text
0
```

Why?

Because int default value is:

```text
0
```

---

# Writing to Nil Map ❌

```go
score["math"] = 90
```

Panic!

```text
panic: assignment to entry in nil map
```

Because map memory was never created.

Like:

```text
Trying to put clothes
inside a cupboard
that doesn't exist.
```

---

# Solution: make()

```go
score = make(map[string]int)
```

Now Go creates memory.

Visual:

```text
score
 |
 v

empty map
```

Now you can store values.

```go
score["math"] = 90
score["science"] = 80
```

Map becomes:

```text
"math"    -> 90
"science" -> 80
```

---

# Full Example

```go
score := make(map[string]int)

score["math"] = 90
score["science"] = 80

fmt.Println(score)
```

Output:

```text
map[math:90 science:80]
```

---

# Adding New Values

```go
ages["ram"] = 50
```

Before:

```text
sangam -> 65
jone   -> 75
```

After:

```text
sangam -> 65
jone   -> 75
ram    -> 50
```

---

# Updating Value

```go
ages["sangam"] = 100
```

Before:

```text
sangam -> 65
```

After:

```text
sangam -> 100
```

---

# Delete

```go
delete(ages, "jone")
```

Before:

```text
sangam -> 65
jone   -> 75
```

After:

```text
sangam -> 65
```

---

# Check if Key Exists

Suppose:

```go
ages := map[string]int{
    "sangam": 65,
}
```

You ask:

```go
fmt.Println(ages["jone"])
```

Output:

```text
0
```

Problem:

```text
Did jone exist?
OR
Was value actually 0?
```

Go gives special syntax:

```go
age, exists := ages["jone"]

fmt.Println(age)
fmt.Println(exists)
```

Output:

```text
0
false
```

Another example:

```go
age, exists := ages["sangam"]
```

Output:

```text
65
true
```

---

# Real-Life Example 1

Student Marks

```go
marks := map[string]int{
    "Abhi": 90,
    "Ram":  80,
}
```

```text
Abhi -> 90
Ram  -> 80
```

---

# Real-Life Example 2

Phone Book

```go
phone := map[string]string{
    "Mom": "999999999",
    "Dad": "888888888",
}
```

```text
Mom -> 999999999
Dad -> 888888888
```

---

# Real-Life Example 3

Product Prices

```go
prices := map[string]int{
    "Laptop": 50000,
    "Mouse":  500,
}
```

```text
Laptop -> 50000
Mouse  -> 500
```

---

# Two Ways to Create Maps

### Method 1

```go
ages := map[string]int{
    "sangam": 65,
    "jone": 75,
}
```

Already contains data.

---

### Method 2

```go
ages := make(map[string]int)
```

Starts empty.

```go
ages["sangam"] = 65
ages["jone"] = 75
```

---

# Easy Memory Trick

```text
map = storage

key -> value
```

```text
Dictionary
Word -> Meaning
```

```text
Phonebook
Name -> Number
```

```text
Student
Name -> Marks
```

All are maps.

---

# Mini Quiz

What is output?

```go
ages := map[string]int{
    "sangam": 65,
    "jone": 75,
}

delete(ages, "jone")

fmt.Println(len(ages))
```

Answer:

```text
1
```

Because only:

```text
sangam -> 65
```

remains. Got it ? 😎

# goscore : Underscore like library for golang!

## Overview
 
[![GoDoc](https://godoc.org/github.com/Ar11rA/goscore?status.svg)](https://godoc.org/github.com/Ar11rA/goscore)
[![Go Report Card](https://goreportcard.com/badge/github.com/Ar11rA/goscore)](https://goreportcard.com/report/github.com/Ar11rA/goscore)
![Go](https://github.com/Ar11rA/goscore/workflows/Go/badge.svg?branch=master)

- Make Generic lists in golang
- Write functional code
- Functions supported
    * Map
    * Reduce
    * Filter
    * Each 
    * Reverse
    * Find
    * Delete
    * Unique
    * Sort

## Install

```
go get github.com/Ar11rA/goscore
```

## Usage 

1. Simple - Array of strings or numbers

```go
package score

import (
	"fmt"
	"github.com/Ar11rA/goscore"
	)

func main() {
	l1 := goscore.List{1,2,3,4,5,6,7,8,9,10}
	l2 := l1.Map(func(i interface{}) interface{} {
		return i.(int) * 3
	})
	fmt.Println(l2)
	l3 := l2.Filter(func(i interface{}) bool {
		return i.(int) % 2 == 0
	})
	fmt.Println(l3)
	p := l3.Reduce(func(acc interface{}, i interface{}) interface{} {
		return acc.(int) * i.(int)
	}, 1)
	fmt.Println(p)
	fmt.Println(l3.Find(12))
}
```

Output

```shell script
[3 6 9 12 15 18 21 24 27 30]
[6 12 18 24 30]
933120
[1] true
```

2. Complex - Array of maps

```go
import (
	"fmt"
	"github.com/Ar11rA/goscore"
	)

func main() {
	l1 := goscore.List{
		map[string]interface{}{
			"name": "john",
			"age": 20,
			"division": "CS",
		},
		map[string]interface{}{
			"name": "jane",
			"age": 30,
			"division": "EC",
		},
		map[string]interface{}{
			"name": "sam",
			"age": 35,
			"division": "CS",
		},
	}
	fmt.Println(l1)
	l2 := l1.Filter(func(i interface{}) bool {
		return i.(map[string]interface{})["age"].(int) < 25
	})
	fmt.Println(l2)
	l3 := l1.Map(func(i interface{}) interface{} {
		nm := map[string]interface{}{
			"name": i.(map[string]interface{})["name"].(string),
			"age": i.(map[string]interface{})["age"].(int),
			"checked": true,
		}
		return nm
	})
	fmt.Println(l3)
	l4 := l1.Reduce(func(acc interface{}, i interface{}) interface{} {
		acc.(map[string][]string)[i.(map[string]interface{})["division"].(string)] = append(acc.(map[string][]string)[i.(map[string]interface{})["division"].(string)], i.(map[string]interface{})["name"].(string))
		return acc
	}, map[string][]string{})
	fmt.Println(l4)
}

```

Output
```shell script
[map[age:20 division:CS name:john] map[age:30 division:EC name:jane] map[age:35 division:CS name:sam]]
[map[age:20 division:CS name:john]]
[map[age:20 checked:true name:john] map[age:30 checked:true name:jane] map[age:35 checked:true name:sam]]
map[CS:[john sam] EC:[jane]]
```

Run `go run main.go` to see the results
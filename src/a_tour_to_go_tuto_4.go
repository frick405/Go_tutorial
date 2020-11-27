// The name starting with capital is exported in Golang
/* // Wrong case
package main

import (
    "fmt"
    "math"
)

func main() {
	fmt.Println(math.pi)
}*/

//Right case
package main

import (
    "fmt"
    "math"
)

func main() {
	fmt.Println(math.Pi)
}


// If variable type is not declared, then type is inferred by right value type
package main

import (
	"fmt"
)

func main() {
		/*v_1 := 42
		v_2 := v_1
		fmt.Printf("v_1 is of type %T\n", v_1)
		fmt.Printf("v_2 is of type %T\n", v_2)
		 or we can code below way*/
		var v_1 int = 1
		var v_2 int = v_1

		fmt.Printf("v_1 is of type %T\n", v_1)
		fmt.Printf("v_2 is of type %T\n", v_2)
}



// UNIX 系统命令 echo 实现

package main

import (
	"fmt"
	"os"
	"strings"
)

/*func main() {

	var s, sep string
	sep = " "

	for _, arg := range os.Args[1:] {
		s += sep + arg
	}

	fmt.Println(s)
}
*/

func main() {

	fmt.Println(strings.Join(os.Args[:], " "))
}

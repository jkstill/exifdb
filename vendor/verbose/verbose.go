
package verbose

import (
	"fmt"
)

func Println(v bool, s string) {

	if v {
		fmt.Println("verbose: " + s)
	}
}

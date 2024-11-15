// +build !linux

package extensions

import "fmt"

func Thing() {
	fmt.Println("Hello from non-linux")
}

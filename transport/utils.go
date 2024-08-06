package transport

import "fmt"

func addressFormat(a, p string) string {
	return fmt.Sprintf("%s:%s", a, p)
}

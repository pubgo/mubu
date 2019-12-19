package mubu

import (
	"fmt"
)

var _url = func(url string) func(...interface{}) string {
	return func(params ...interface{}) string {
		return fmt.Sprintf(url, params...)
	}
}

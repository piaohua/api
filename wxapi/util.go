package wxapi

import "fmt"

// ToQueryString convert the map[string]string to query string parameters
func ToQueryString(param map[string]string) string {
	var args string
	for k, v := range param {
		if args == "" {
			args += fmt.Sprintf("?%s=%s", k, v)
		} else {
			args += fmt.Sprintf("&%s=%s", k, v)
		}
	}

	return args
}

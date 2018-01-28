package smsbao

import "net/url"

func urlencoder(str string) string {
	return url.QueryEscape(str)
}

package internal

import (
	"net/url"
)

func ParseReqForm(f url.Values, key string) string {
	for k, v := range f { // range over map
		for _, value := range v { // range over []string
			if key == k {
				return value
			}
		}
	}
	return ""

}

package tool

import "github.com/tidwall/gjson"

func ReadStringField(raw, key string) string {
	r := gjson.Get(raw, key)
	return r.Raw
}

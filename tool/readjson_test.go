package tool

import "testing"

func TestReadStringField(t *testing.T) {
	raw := `{"name":"ZhangSan"}`
	t.Log(ReadStringField(raw, "name"))
}

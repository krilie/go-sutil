package util_str

import "testing"

func TestToJsonPretty(t *testing.T) {
	pretty := ToJsonPretty(123231)
	t.Log(pretty)
}

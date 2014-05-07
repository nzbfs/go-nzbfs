//
package yenc

import "testing"

func TestDecode(t *testing.T) {
	data := "Some data"

	var decoded = Decode(data)

	if data != decoded {
		t.Errorf("Decode(%s) != %s", data, decoded)
	}
}

// Testing for the inline yenc encoder
package yenc

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestDecode(t *testing.T) {
	var fileBytes []byte
	var _ error

	fileBytes, _ = ioutil.ReadFile("singlepart_test_data_only.yenc")
	//fileBytes, _ = ioutil.ReadFile("multipart_test1.yenc")

	var decoded = Decode(fileBytes)
	fmt.Print("length of fileBytes is ", len(fileBytes), "\n")
	fmt.Print("length of decoded is ", len(decoded), "\n")

	for x := 0; x < len(decoded); x++ {
		fmt.Printf("%c", decoded[x])
	}

	ioutil.WriteFile("singlepart_test1.txt", decoded, 0644)
}

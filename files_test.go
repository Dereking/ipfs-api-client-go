package IPFSClient

import (
	//"fmt"
	"testing"
)

func Test_FilesLs(t *testing.T) {
	client := NewIPFSClientLocal()
	res, err := client.FilesLs("/", false, false)
	if err != nil {
		t.Fatalf(err.Error())
	}
	//fmt.Println("Test_FilesLs", res)
	//t.Println(res)

	assertEqual(t, len(res.Entries), 1)
}

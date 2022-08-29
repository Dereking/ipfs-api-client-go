package IPFSClient

import (
	"testing"
)

func Test_Add(t *testing.T) {

	req := NewAddReq()
	req.SrcFilePath = "./test.txt"
	req.TargetPath = "/test.txt"

	//r, err := Add(req)
	client := NewIPFSClientLocal()
	res, err := client.Add(req)
	if err != nil {
		t.Fatalf(err.Error())
	}
	//t.Log("add ", res, req)
	assertEqual(t, res,
		`{"Name":"test.txt","Hash":"QmUL7wDowvNk3y7KeEYFAATmz43727FwXKhBJJrqQu813a","Size":"11"}
`)

}

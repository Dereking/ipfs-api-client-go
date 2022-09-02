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

func Test_Cat(t *testing.T) {

	req := NewCatReq()
	req.IpfsPath = "QmTEzo7FYzUCd5aq7bGKoMLfsbbsebpfARRZd4Znejb25R"

	//r, err := Add(req)
	client := NewIPFSClientLocal()
	res, err := client.Cat(req)
	if err != nil {
		t.Fatalf(err.Error())
	}
	t.Log("cat ", res, req)
	assertEqual(t, res,
		`{"Name":"test.txt","Hash":"QmUL7wDowvNk3y7KeEYFAATmz43727FwXKhBJJrqQu813a","Size":"11"}
`)

}

func Test_Commands(t *testing.T) {
	req := NewCommandsReq()
	req.ShowFlags = true

	//r, err := Add(req)
	client := NewIPFSClientLocal()
	res, err := client.Commands(req)
	if err != nil {
		t.Fatalf(err.Error())
	}
	t.Log("Test_Commands ", res)
	// assertEqual(t, res,
	// 	`{"Name":"test.txt","Hash":"QmUL7wDowvNk3y7KeEYFAATmz43727FwXKhBJJrqQu813a","Size":"11"}
	//`)
}

func Test_DhtQuery(t *testing.T) {

	//r, err := Add(req)
	client := NewIPFSClientLocal()
	_, err := client.DhtQuery("12D3KooWJju1H2zf2iZZBD3Vq5E5mN2D8ipxomsKVLC3yQ3dDzvh", false)
	if err != nil {
		t.Fatalf(err.Error())
	}
	//t.Log("Test_DhtQuery ", len(res) )
	// assertEqual(t, res,
	// 	`{"Name":"test.txt","Hash":"QmUL7wDowvNk3y7KeEYFAATmz43727FwXKhBJJrqQu813a","Size":"11"}
	//`)
}

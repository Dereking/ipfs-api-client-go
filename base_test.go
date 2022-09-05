package IPFSClient

import (
	"log"
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
	assertEqual(t, res.Hash,
		`QmUL7wDowvNk3y7KeEYFAATmz43727FwXKhBJJrqQu813a`)

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
	assertEqual(t, res, `123
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
		t.Fatalf("Test_DhtQuery err: %s", err.Error())
	}
	//t.Log("Test_DhtQuery ", len(res) )
	// assertEqual(t, res,
	// 	`{"Name":"test.txt","Hash":"QmUL7wDowvNk3y7KeEYFAATmz43727FwXKhBJJrqQu813a","Size":"11"}
	//`)
}

func Test_Get(t *testing.T) {

	req := NewGetReq()
	req.IpfsPath = "QmTEzo7FYzUCd5aq7bGKoMLfsbbsebpfARRZd4Znejb25R"

	client := NewIPFSClientLocal()
	res, err := client.Get(req)
	if err != nil {
		t.Fatalf(err.Error())
	}
	t.Log("cat ", res, req)
	// assertEqual(t, res, `string=QmTEzo7FYzUCd5aq7bGKoMLfsbbsebpfARRZd4Znejb25R0000644000000000000000000000000414304367170017241 0ustar0000000000000000123
	//     `)

}

func Test_Id(t *testing.T) {

	req := NewIdReq()
	//req.IpfsPath = "QmTEzo7FYzUCd5aq7bGKoMLfsbbsebpfARRZd4Znejb25R"

	client := NewIPFSClientLocal()
	res, err := client.Id(req)
	if err != nil {
		t.Fatalf(err.Error())
	}
	t.Log("id ", res, req)
	//log.Println("id ", res, req)
	// assertEqual(t, res, `string=QmTEzo7FYzUCd5aq7bGKoMLfsbbsebpfARRZd4Znejb25R0000644000000000000000000000000414304367170017241 0ustar0000000000000000123
	//     `)

}

func Test_LogLevel(t *testing.T) {

	req := NewLogLevelReq()
	//req.IpfsPath = "QmTEzo7FYzUCd5aq7bGKoMLfsbbsebpfARRZd4Znejb25R"

	client := NewIPFSClientLocal()
	res, err := client.LogLevel(req)
	if err != nil {
		t.Fatalf("LogLevel fail:" + err.Error())
	}
	//t.Log("id ", res, req)
	log.Println("LogLevel ", res, req)
	// assertEqual(t, res, `string=QmTEzo7FYzUCd5aq7bGKoMLfsbbsebpfARRZd4Znejb25R0000644000000000000000000000000414304367170017241 0ustar0000000000000000123
	//     `)

}

func Test_LogLs(t *testing.T) {

	client := NewIPFSClientLocal()
	res, err := client.LogLs()
	if err != nil {
		t.Fatalf("Test_LogLs fail:" + err.Error())
	}
	t.Log("logls ", res)
	//log.Println("Test_LogLs ", res)
	// assertEqual(t, res, `string=QmTEzo7FYzUCd5aq7bGKoMLfsbbsebpfARRZd4Znejb25R0000644000000000000000000000000414304367170017241 0ustar0000000000000000123
	//     `)

}

func Test_Ls(t *testing.T) {

	req := NewLsReq()
	req.IpfsPath = "QmTEzo7FYzUCd5aq7bGKoMLfsbbsebpfARRZd4Znejb25R--"

	client := NewIPFSClientLocal()
	res, err := client.Ls(req)
	if err != nil {
		t.Fatalf("Ls fail:" + err.Error())
	}
	t.Log("ls ", res, req)
	//log.Println("Ls ", res, req)
	// assertEqual(t, res, `string=QmTEzo7FYzUCd5aq7bGKoMLfsbbsebpfARRZd4Znejb25R0000644000000000000000000000000414304367170017241 0ustar0000000000000000123
	//     `)

}

func Test_Ping(t *testing.T) {

	req := NewPingReq("12D3KooWFS3Q9oZn5z3VnELmUoYr4qVDyEsNb75Wqnzumxwtkv8s")
	//req.IpfsPath = "QmTEzo7FYzUCd5aq7bGKoMLfsbbsebpfARRZd4Znejb25R--"

	client := NewIPFSClientLocal()
	res, err := client.Ping(req)
	if err != nil {
		t.Fatalf("Ls fail:" + err.Error())
	}
	t.Log("Ping ", res, req)
	//log.Println("Ls ", res, req)
	// assertEqual(t, res, `string=QmTEzo7FYzUCd5aq7bGKoMLfsbbsebpfARRZd4Znejb25R0000644000000000000000000000000414304367170017241 0ustar0000000000000000123
	//     `)

}



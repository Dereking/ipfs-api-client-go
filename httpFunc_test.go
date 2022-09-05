package IPFSClient

import (
	//"io/ioutil"
	//"log"
	//"net/http"
	//"net/url"
	"testing"
)

func Test_PostForm(t *testing.T) {
	query := make(map[string][]string)
	form := make(map[string][]string)

	res, err := PostForm("http://127.0.0.1:5001/api/v0/files/ls", query, form)
	if err != nil {
		t.Fatal(err)
	}
	//log.Println(string(res))
	assertEqual(t, string(res),
		`{"Entries":[{"Name":"test","Type":0,"Size":0,"Hash":""}]}
`)
}

func Test_PostFormWithFile(t *testing.T) {

	query := make(map[string][]string)
	form := make(map[string][]string)
	res, err := PostFormWithFile("http://127.0.0.1:5001/api/v0/add", query, form, "path", "/tmp/123.txt")
	if err != nil {
		t.Fatal(err)
	}
	//log.Println(string(res))
	assertEqual(t, string(res),
		`{"Name":"tmp/123.txt","Hash":"QmTEzo7FYzUCd5aq7bGKoMLfsbbsebpfARRZd4Znejb25R","Size":"12"}
{"Name":"tmp","Hash":"QmUhdcbR4WzvKmL2THRhB8NQjm5oSbNysvdmfxcX8oDcft","Size":"65"}
`)
}

package IPFSClient

import (
	//"fmt"
	"testing"
)

func Test_StructToHttpData(t *testing.T) {
	tests := map[string]struct {
		req       interface{}
		expectedQ string
		expectedF string
	}{
		`StructToHttpData(AddReq)`: {
			req:       AddReq{SrcFilePath: "./test.txt", TargetPath: "/test.txt", HashAlgorithm: Sha2_256},
			expectedQ: "&cid-version=0&fscache=false&inlhashine=sha2-256&inline=false&inline-limit=0&nocopy=false&only-hash=false&pin=false&progress=false&quiet=false&quieter=false&raw-leaves=false&silent=false&trickle=false&wrap-with-directory=false",
			expectedF: "",
		},
		`CatReq`: {
			req:       CatReq{IpfsPath: "sdferwedgfgf"},
			expectedQ: "&arg=sdferwedgfgf&length=0&offset=0&progress=false",
			expectedF: "",
		},
	}
	for name, tt := range tests {
		//fmt.Println(name, tt)
		t.Run(name, func(t *testing.T) {
			//	t.Log(name, tt)
			q, f, err := StructToHttpData(tt.req)
			if err != nil {
				t.Errorf("StructToHttpData err: %s", err.Error())
				return
			}

			if q != tt.expectedQ {
				t.Errorf("got Q: %s, expected: %s", q, tt.expectedQ)
			}
			if f != tt.expectedF {
				t.Errorf("got F: [%s], expected: [%s]", f, tt.expectedF)
			}

			//t.Log(q, f)
		})
	}
}

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
			expectedQ: "&quiet=false&quieter=false&silent=false&progress=false&trickle=false&only-hash=false&wrap-with-directory=false&chunker=&pin=false&raw-leaves=false&nocopy=false&fscache=false&cid-version=0&inlhashine=sha2-256&inline=false&inline-limit=0",
			expectedF: "",
		},
		`CatReq`: {
			req:       CatReq{IpfsPath: "sdferwedgfgf"},
			expectedQ: "&arg=sdferwedgfgf&offset=0&length=0&progress=false",
			expectedF: "",
		},
	}
	for name, tt := range tests {
		//fmt.Println(name, tt)
		t.Run(name, func(t *testing.T) {
			//	t.Log(name, tt)
			q, f := StructToHttpData(tt.req)

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

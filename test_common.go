package IPFSClient

import (
	"io"
	"log"
	"net"
	"net/http"
	"strings"
)

type TestCommonInterface interface {
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Fail()
	FailNow()
	Failed() bool
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	Helper()
	Log(args ...interface{})
	Logf(format string, args ...interface{})
	Name() string
	Skip(args ...interface{})
	SkipNow()
	Skipf(format string, args ...interface{})
	Skipped() bool
}

func assertEqual(t TestCommonInterface, a, b interface{}) {

	t.Helper()

	valA, okA := a.(string)
	valB, okB := b.(string)
	if okA && okB { //string
		if strings.Compare(valA, valB) != 0 {
			t.Errorf("String Not Equal. %d %d", a, b)
		}
	} else {
		if a != b {
			t.Errorf("Not Equal. %d %d", a, b)
		}
	}
}

func runTmpHttp() (*http.Server, int, error) {
	mux := http.NewServeMux()
	mux.HandleFunc("/1.sh", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, `#!/bin/bash
echo 123 > /tmp/123`)
	})

	server := &http.Server{
		Addr:    ":0",
		Handler: mux,
	}
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		return nil, 0, err
	}
	log.Println("Using port:", listener.Addr().(*net.TCPAddr).Port)
	go server.Serve(listener)

	return server, listener.Addr().(*net.TCPAddr).Port, nil
}

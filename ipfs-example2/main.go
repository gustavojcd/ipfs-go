package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	shell "github.com/ipfs/go-ipfs-api"
)

type TimeSeriesDatum struct {
	Id    uint64 `json:"id"`
	Value uint64 `json:"value"`
}

func main() {
	sh := shell.NewShell("localhost:5001")
	tsd := &TimeSeriesDatum{
		Id:    1,
		Value: 123,
	}
	tsdBin, _ := json.Marshal(tsd)
	reader := bytes.NewReader(tsdBin)

	cid, err := sh.Add(reader)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		os.Exit(1)
	}
	fmt.Printf("added %s\n", cid)

	data, err := sh.Cat(cid)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		os.Exit(1)
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(data)
	newStr := buf.String()

	res := &TimeSeriesDatum{}
	json.Unmarshal([]byte(newStr), &res)
	fmt.Println(res)
}

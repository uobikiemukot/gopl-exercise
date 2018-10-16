package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Info struct {
	Num        int    `json:"num"`
	Transcript string `json:"transcript"`
}

const (
	Num  = 2059
	Path = "./xkcd"
)

func loadIndex(info *[Num]Info) error {
	for i := 0; i < Num; i++ {
		if i == 403 {
			// info[403] <===> https://xkcd.com/404/info.0.json
			// 404 Not Found
			// Maybe joke?
			continue
		}

		path := fmt.Sprintf("%s/%d.json", Path, i+1)
		b, err := ioutil.ReadFile(path)
		if err != nil {
			return fmt.Errorf("ioutil.ReadFile() failed (path: %s): %s", path, err)
		}

		err = json.Unmarshal(b, &info[i])
		if err != nil {
			return fmt.Errorf("json.Unmarshal() failed (path: %s): %s", path, err)
		}
	}

	return nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: xkcd SEARCH_TERM\n")
		os.Exit(1)
	}

	var info [Num]Info
	err := loadIndex(&info)
	if err != nil {
		fmt.Fprintf(os.Stderr, "loadIndex() failed: %s\n", err)
		os.Exit(2)
	}

	for i := range info {
		if strings.Index(strings.ToLower(info[i].Transcript), strings.ToLower(os.Args[1])) >= 0 {
			fmt.Printf("==> https://xkcd.com/%d/ <==\n", info[i].Num)
			fmt.Printf("%s\n\n", info[i].Transcript)
		}
	}
}

package main

import (
	"fmt"
	"github.com/tangjicheng46/powergrid/yt"
)

func main() {
	url := "https://www.youtube.com/watch?v=qt2lLJ3MvkM"
	err := yt.DownloadWithRecord(url)
	if err != nil {
		fmt.Printf("fail download: %s, error: %s\n", url, err)
	}
	fmt.Printf("success download: %s\n", url)
	return
}

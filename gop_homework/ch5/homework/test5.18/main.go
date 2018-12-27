//练习5.18：不修改fetch的行为，重写fetch函数，要求使用defer机制关闭文件。
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

//!+
// Fetch downloads the URL and returns the
// name and length of the local file.
func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	n, err = io.Copy(f, resp.Body)
	// Close file, but prefer error from Copy, if any.
	defer func() {
		if closeErr := f.Close(); err == nil {
			err = closeErr
		}
		filename = url
	}()

	return // 使用这种方式的return return结果在defer之后，因为 return后面没有东西。
}

//!-

func main() {
	for _, url := range os.Args[1:] {
		local, n, err := fetch(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch %s: %v\n", url, err)
			continue
		}
		fmt.Fprintf(os.Stderr, "%s => %s (%d bytes).\n", url, local, n)
	}
}
//go run main.go https://coastroad.net/donate
//https://coastroad.net/donate => https://coastroad.net/donate (3135 bytes).
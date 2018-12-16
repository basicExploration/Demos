//练习 4.13： 使用开放电影数据库的JSON服务接口，
// 允许你检索和下载 https://omdbapi.com/ 上电影的名字和对应的海报图像。
// 编写一个poster工具，通过命令行输入的电影名字，下载对应的海报。
package main

import (
	"encoding/json"
	"fmt"
	"github.com/urfave/cli"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type PosterContent struct {
	Title  string
	Poster string
}

func main() {
	middleDate := make([]byte, 0)
	cliApp := cli.NewApp()
	cliApp.Name = "Poster"
	cliApp.Version = "v0.1.0"
	cliApp.Author = "[googege](https://github.com/googege)"
	cliApp.Usage = "the movie poster tool"
	cliApp.Commands = []cli.Command{
		{
			Name:    "get",
			Aliases: []string{"g"},
			Usage:   "get the https://omdbapi.com 's movie name & image ",
			Action: func(c *cli.Context) error {
				args := c.Args()
				resp, err := http.Get("https://omdbapi.com?apikey=c770ccb7&t=" + args[0] + "&y=" + args[1])
				if err != nil {
					log.Fatal(err)
				}
				defer resp.Body.Close()
				da, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					log.Fatal(err)
				}
				middleDate = append(middleDate, da...)
				return nil
			},
		},
	}
	cliApp.Run(os.Args) // 要放前面 不知道为什么放后面就不能运行。
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		a := new(PosterContent)
		json.Unmarshal(middleDate, a)
		HTMLData := "<p>" + a.Title + "</p>" + "<image " + "src=" + "'" + a.Poster + "'" + "/>"
		fmt.Fprint(writer, HTMLData)

	})
	http.ListenAndServe(":8090", nil)

}

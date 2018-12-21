//练习 4.11： 编写一个工具，允许用户在命令行创建、读取、更新和关闭GitHub上的issue，当必要的时候自动打开用户默认的编辑器用于输入文本信息。
package main

import (
	"encoding/json"
	"fmt"
	"github.com/urfave/cli"
	"io/ioutil"
	"net/http"
	"strings"
	"os"
)

func init() {
	cliApp := cli.NewApp()
	cliApp.Version = "v0.1.0"
	cliApp.Name = "Issue helper"
	cliApp.Author = "[googege](https://github.com/googege)"
	cliApp.Usage = "this is a tool that can help you control your github's issue ,such as creating reading closing updating," +
		"and more importantly, it can open your IDE when you need it."
	cliApp.Commands = []cli.Command{
		{
			Name:    "creat",
			Aliases: []string{"cr"},
			Usage:   "creat your issue",
			Action:  CreatIssue,
		},
		{
			Name:    "read",
			Aliases: []string{"r"},
			Usage:   "read your issue",

			Action: ReadIssue,
		},
		{
			Name:    "update",
			Aliases: []string{"u"},
			Usage:   "update your issue",
			Action:  UpdateIssue,
		},
		{
			Name:    "close",
			Aliases: []string{"cl"},
			Usage:   "close your issue",
			Action:  CloseIssue,
		},
	}
	cliApp.Run(os.Args)

}

type JsonData struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

func main() {

}

//CreatIssue 读取issue
func CreatIssue(c *cli.Context) error {
	fmt.Println("ceshi",c.Args().Get(0))
	data, _ := json.Marshal(JsonData{"test1", "test"})
	read := strings.NewReader(string(data))
	res, err := http.Post("https://api.github.com/repos/googege/goFiles/issues"+"?access_token=你的token", "text/json", read)
	d, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(d))
	res.Body.Close()
	fmt.Println("err :=======", err)
	return err
}

//ReadIssue 读取issue
func ReadIssue(c *cli.Context) error {
	//PASS
	return nil
}

//UpdateIssue 读取issue
func UpdateIssue(c *cli.Context) error {
	//PASS
	return nil
}

//CloseIssue 读取issue
func CloseIssue(c *cli.Context) error {
	//PASS
	return nil
}

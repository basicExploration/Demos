package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
)

type Nice struct {
	Name   string `json:"name"`
	Year   int    `json:"year",omitempty`
	Height int    `json:"height"`
	Weight int    `json:"weight"`
}

func main() {
	nice := new(Nice)
	nice.Name = "JackieHoo"
	//nice.Year = 2018
	nice.Height = 178
	nice.Weight = 78
	jsonValue, err := json.Marshal(nice)
	if err != nil {
		fmt.Println("在得到json编码的返回值的时候发生了错误，错误代码是:", err)
	}
	file, err := os.OpenFile("./testJsonData.json", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("在创建写入testJsonData.json文件时发生错误：", err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		if i < 0 {
			fmt.Errorf("非法入侵")
			panic("非法入侵")
		}
		if i == 0 {
			writer.WriteByte('{')
		}
		writer.WriteString("\"" + "value" + strconv.FormatInt(int64(i), 10) + "\"" + ":")
		for _, v := range jsonValue {
			err = writer.WriteByte(v)
			if err == io.EOF {
				fmt.Println("信息的传入已经结束")
				break
			}
		}
		if i != 9 && i >= 0 {
			writer.WriteByte(',')
		}
		writer.WriteByte('\n')

		if i == 9 {
			writer.WriteByte('}')
		}
	}
	writer.Flush() //在所有的数据都传入结束后，一次性将数据打入文件中。
	fmt.Println("成功！")

}


###

{"value0":{"name":"JackieHoo","year":0,"height":178,"weight":78},
"value1":{"name":"JackieHoo","year":0,"height":178,"weight":78},
"value2":{"name":"JackieHoo","year":0,"height":178,"weight":78},
"value3":{"name":"JackieHoo","year":0,"height":178,"weight":78},
"value4":{"name":"JackieHoo","year":0,"height":178,"weight":78},
"value5":{"name":"JackieHoo","year":0,"height":178,"weight":78},
"value6":{"name":"JackieHoo","year":0,"height":178,"weight":78},
"value7":{"name":"JackieHoo","year":0,"height":178,"weight":78},
"value8":{"name":"JackieHoo","year":0,"height":178,"weight":78},
"value9":{"name":"JackieHoo","year":0,"height":178,"weight":78}
}

//练习 7.9： 使用html/template包 (§4.6) 替代printTracks将tracks展示成一个HTML表格。
// 将这个解决方案用在前一个练习中，让每次点击一个列的头部产生一个HTTP请求来排序这个表格。
package main

import (
	"fmt"
	"html/template"
	"os"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   time.Time
	Length int
}

func printTracks(tracks []*Track) {
	va := `
<table>
<tr>
<td>
{{.Title}}
</td>
<td>
{{.Artist}}
</td>
<td>
{{.Album}}
</td>
<td>
{{.Year}}
</td>
<td>
{{.Length}}
</td>
</tr>
<tr>
<td>
</td>
<td>
</td>
<td>
</td>
<td>
</td>
<td>
</td>
</tr>

</table>


`
	t1, err := template.New("foo").Parse(va)
	if err != nil {
		fmt.Println(err)
	}

	//const format = "%v\t%v\t%v\t%v\t%v\t\n"
	//tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	//fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	//fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		err = t1.ExecuteTemplate(os.Stdout, "foo", t)
		//fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	//tw.Flush() // calculate column widths and print table
}

func main() {
	printTracks([]*Track{
		&Track{
			"12", "233", "2332", time.Now(), 23,
		},
		&Track{
			"fdfdfd", "2fsdfsd33", "233dsfsdf2", time.Now(), 665,
		},
	})
}

package main


import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// 游戏读取数据，读取游戏配置数据
func ReadCsv(fileName string) bool {
	// 获取数据，按照文件
	cntb, err := ioutil.ReadFile(fileName)
	if err != nil {
		return false
	}
	// 读取文件数据
	r2 := csv.NewReader(strings.NewReader(string(cntb)))
	ss, _ := r2.ReadAll()
	//fmt.Println(ss)
	sz := len(ss)
	// 循环取数据
	fmt.Println(sz)
	nameMap := map[string][]string{}
	for i := 0; i < sz; i++ {
		// fmt.Println(ss[i])
		// fmt.Println(ss[i][0]) //  key的数据  可以作为map的数据的值
	}



	//创建文件
	f, err := os.Create("过滤.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	// 写入UTF-8 BOM
	f.WriteString("\xEF\xBB\xBF")
	//创建一个新的写入文件流
	w := csv.NewWriter(f)


	for _, b := range nameMap {
		w.Write(b)
	}


	//写入数据
	//w.WriteAll(data)
	w.Flush()




	return true
}

func main() {
	var a interface{}
	a = 1
	val, ok := a.(string)
	fmt.Println(ok)
	fmt.Println(val)

	//fmt.Println(strings.Replace("12345 67", " ", "", -1))
}
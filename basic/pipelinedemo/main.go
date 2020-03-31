package main

import (
	"bufio"
	"fmt"
	"learn_go/pipeline"
	"os"
)

func main() {
	// MergeDemo()
	// 第一步: 造数据, 生成100个随机数, 写入到文件
    const fileName  = "small.in"
    const count = 64
    // 第一步: 将随机生产的数字保存到small.in文件
    // 构造第一个数据源
    file, e := os.Create(fileName)
    if e != nil {
        panic(e)
    }
    defer file.Close()
    dataSource := pipeline.RandomSource(count)
    writer := bufio.NewWriter(file)
    pipeline.WriteSink(writer, dataSource)
    writer.Flush()
    
    // 第二步: 从文件中读取文件内容, 在控制台打印
    // 从第一个数据源读取出数据
    f, e := os.Open(fileName)
    if e != nil {
        panic(e)
    }
    defer f.Close()
    readerSource := pipeline.ReaderSource(bufio.NewReader(f), -1)

    var num = 0
    for rs := range readerSource{
        fmt.Println(rs)
        num ++
        if num > 100 {
            break
        }
    }
}


// MergeDemo func
func MergeDemo()  {
	p := pipeline.Merge(
		pipeline.InMemSort(
			pipeline.ArraySource(3, 2, 6, 7, 4)),
			pipeline.InMemSort(
				pipeline.ArraySource(3, 2, 6, 7, 1)))
	for v := range p {
		fmt.Println(v)
	}
}
package pipeline

import (
	"encoding/binary"
	"io"
	"math/rand"
	"sort"
)

// ArraySource func
func ArraySource(a ...int) <-chan int {
	out := make(chan int)
	go func ()  {
		for _, v := range a {
			out <- v
		}
		close(out)
	}()
	return out
}


// InMemSort func
func InMemSort(in <-chan int) <-chan int {
	out := make(chan int)
	go func()  {
		a := []int{}
		for v := range in {
			a = append(a, v)
		}
		// Sort
		sort.Ints(a)

		// Output

		for _, v := range a {
			out <- v
		}
		close(out)
	}()
	return out
}

// Merge func
func Merge(in1, in2 <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		v1, ok1 := <-in1
		v2, ok2 := <-in2
		for ok1 || ok2 {
			if !ok2 || (ok1 && v1 <= v2){
				out <- v1
				v1, ok1 = <-in1
			} else {
				out <- v2
				v2, ok2 = <-in2
			}
		}
		close(out)
	}()
	return out
}


// ReaderSource func
func ReaderSource(reader io.Reader, chunkSize int) <- chan int{
    out := make(chan int)
    go func() {
        // 这里定义为8个字节, 原因是我的机器是64位的, 所以int也是64位, 那么对应的字节数就是8个字节
		buffer := make([]byte, 8)
		bytesRead := 0
        for   {
            // reader返回两个参数, 第一个是读取到的字节数, 第二个是err异常
			n, err := reader.Read(buffer)
			bytesRead += n
            if n > 0  {
                // 如果读到了, 就把读到的东西发给channel
				u := binary.BigEndian.Uint64(buffer)
                out <- int(u)
            }

            if err != nil || (chunkSize != -1 && bytesRead >= chunkSize) {
                break
            }
        }
        close(out)
    }()
    return out
}



// WriteSink func
func WriteSink(writer io.Writer, in <-chan int) {
    for v := range in {
        b := make([]byte, 8)
        binary.BigEndian.PutUint64(b, uint64(v))
        writer.Write(b)
    }
}


// RandomSource func
func RandomSource(count int) chan int {
	out := make(chan int)
	go func() {
		// 生成count个随机数
		for i := 0; i<count ; i++ {
			out <- rand.Int()
		}

		close(out)
	}()
	return out
}


// MergeN func
func MergeN(inputs ...<-chan int) <-chan int {
	if len(inputs) == 1 {
		return inputs[0]
	}

	m := len(inputs) / 2
	return Merge(MergeN(inputs[:m]...), MergeN(inputs[m:]...))
}
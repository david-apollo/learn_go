package fetcher

import (
	"time"
	"golang.org/x/text/encoding/unicode"
	"log"
	"golang.org/x/text/encoding"
	"golang.org/x/net/html/charset"
	"bufio"
	"io/ioutil"
	"golang.org/x/text/transform"
	"fmt"
	"net/http"
)


var rateLmiter = time.Tick(10 * time.Millisecond)

// Fetch func
func Fetch(url string) ([]byte, error) {
	<- rateLmiter
	resp, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Wrong status code: %d", resp.StatusCode)
	}

	bodyReader := bufio.NewReader(resp.Body)
	e := determineEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}


func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("Fetcher error: %v", err)
		return unicode.UTF8
	}

	e, _, _ := charset.DetermineEncoding(bytes, "")

	return e
}
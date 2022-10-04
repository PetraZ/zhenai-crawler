package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main() {
	res, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		fmt.Println(err)
	}
	// when function ends, close the Closer.
	defer res.Body.Close()
	bs, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	reg, err := regexp.Compile(`<a href="(http://www.zhenai.com/zhenghun/[^"]*)[^>]*>([^<]*)</a>`)
	if err != nil {
		fmt.Println(err)
	}

	dic := make(map[string]string)
	matches := reg.FindAllStringSubmatch(string(bs), 1000)
	for _, m := range matches {
		url := m[1]
		name := m[2]
		dic[name] = url
	}
	fmt.Printf("%v", dic)
}

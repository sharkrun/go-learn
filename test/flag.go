package main

import (
	"flag"
	"fmt"
	"strings"
)

var music_file *string = flag.String("file", "musicfile", "Use -file <filesource>")

func test(){
	map1 := make(map[string]string)
	map1["a"] = "AAA"
	map1["b"] = "BBB"
	map1["c"] = "CCC"
	for v := range map1 {
		fmt.Println(map1)
		fmt.Println(v)
		fmt.Println(map1["a"])
	}

	m := strings.Split("/cluster/clustermanage/dev/apply_num", "/")
	fmt.Println("m:", m, m[0], m[1])
	fmt.Println("a[0]", m[0])
	fmt.Println("len(m):", len(m))
	s :=[] int {1,2,3 }
	fmt.Println("s:", s)
}

func main() {
	flag.Parse()
	fmt.Println(*music_file)

	test()
}


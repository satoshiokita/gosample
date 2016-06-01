package gosample

// Go言語では公開する場合は、関数名の先頭を大文字

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

// Message テストメッセージ
var Message = "こんにちはhello world!!"

func hello() {
	fmt.Printf("hello\n")
}

func sum(i int, j int) {
	fmt.Println(i + j)
}

func retSum(i, j int) int {
	return i + j
}

func swap(i, j int) (int, int) {
	return j, i
}

func test1(message string) {
	fmt.Println(strings.ToUpper(Message))
	fmt.Println(message)

	var i int
	fmt.Println(i)

	a, b := 10, 100
	if a > b {
		fmt.Println("a is larger than b")
	} else if a < b {
		fmt.Println(" a is smaller than b")
	} else {
		fmt.Println("a equals b")
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	n := 0
	for n < 10 {
		fmt.Printf("n = %d\n", n)
		n++
	}

	for {
		fmt.Printf("for break\n")
		break
	}

	hello()
	sum(1, 2)
	fmt.Println(retSum(5, 6))

	x, y := 3, 4
	x, y = swap(x, y)
	fmt.Println(x, y)

	x, _ = swap(x, y)
	fmt.Println(x)
}

func div(i, j int) (int, error) {
	if j == 0 {
		return 0, errors.New("divied by zero")
	}
	return i / j, nil
}

// OpenFile 引数filenameのファイルを開く
func OpenFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("file not found.")
		return
	}
	fmt.Println("opened file.", file)
	file.Close()

	//////////
	n, err := div(10, 0)
	if err != nil {
		log.Fatal(err)
		fmt.Println("XXX")
	}
	fmt.Println("n=", n)
}

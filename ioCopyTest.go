package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("./proverbs.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	if _, err := io.Copy(os.Stdout, file); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

//func main() {
//	proverbs := new(bytes.Buffer)
//	proverbs.WriteString("Channels orchestrate mutexes serialize\n")
//	proverbs.WriteString("Cgo is not Go\n")
//	proverbs.WriteString("Errors are values\n")
//	proverbs.WriteString("Don't panic\n")
//
//	file, err := os.Create("./proverbs.txt")
//	if err != nil {
//		fmt.Println(err)
//		os.Exit(1)
//	}
//	defer file.Close()
//
//	// io.Copy 完成了从 proverbs 读取数据并写入 file 的流程
//	if _, err := io.Copy(file, proverbs); err != nil {
//		fmt.Println(err)
//		os.Exit(1)
//	}
//	fmt.Println("file created")
//}
package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	var b bytes.Buffer
	//“使用fmt包里的Fprintf函数将字符串"World!"
	// 追加到Buffer类型变量里”
	b.Write([]byte("Hello "))

	fmt.Fprintf(&b, "World!")
	b.WriteTo(os.Stdout)
}

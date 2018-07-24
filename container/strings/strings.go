package strings

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s:="Yes哈哈哈"
	fmt.Println(len(s))
	//fmt.Printf("%X", s)
	for i,b:=range []byte(s){
		fmt.Printf("(%d %X) ", i, b)
	}

	fmt.Println()

	for i, ch:=range s{ // ch is rune
		fmt.Printf("(%d %X) ", i, ch)
	}

	fmt.Println()

	fmt.Println("Rune Coount:", utf8.RuneCountInString(s))
	bytes:=[]byte(s)
	for len(bytes)>0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}

	fmt.Println()

	// 重新开辟内存，每个字符4字节
	for i,ch:= range []rune(s){
		fmt.Printf("(%d, %c) ", i, ch)
	}
}

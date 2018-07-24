package basic

import (
	"testing"
	"fmt"
)

func TestAdd(t *testing.T) {
	tests := []struct{a, b, c int}{
		{1, 2, 3},
		{4, 5, 6},
		{123, 345, 468},
		{1, 345, 468},
	}

	for _, tt :=range tests{
		if actual := add(tt.a, tt.b); actual!=tt.c{
			t.Errorf("add(%d, %d); get %d; expected %d\n", tt.a, tt.b, actual, tt.c)
		}
	}
}

func BenchmarkAdd(b *testing.B){

	inA, inB:=123, 345
	outC:=468

	// 之前操作不计入计时
	b.ResetTimer()
	for i:=0;i<b.N;i++ {
		actual := add(inA, inB)
		if actual != outC {
			b.Errorf("add(%d, %d); get %d; expected %d\n", inA, inB, actual, outC)
		}
	}
}

func ExampleAdd() {
	fmt.Println(add(1,2))
	fmt.Println(add(3,3))
	fmt.Println(add(5,6))

	// Output:
	// 3
	// 6
	// 11
}

package leetcode

import (
	"fmt"
	"testing"
)



func TestMatricRotate(t *testing.T){
	matrix := [][]int{[]int{1,2,3},[]int{4,5,6},[]int{7,8,9}}
	rotate(matrix)
	fmt.Println(matrix)
}

func rotate(matrix [][]int)  {
	l := len(matrix[0])
	mod := l%2
	step := 0
	if mod == 0{
		step = l/2
	}
	if mod == 1{
		step = (l-1)/2
	}
	if step < 1{
		return
	}

	for i:=0;i<step;i++{
		for j:=i;j<(l-i-1);j++{
			x:=i
			y:=j
			tmp := matrix[x][y]
			for k:=0;k<4;k++{
				x,y = y,(l-1)-x
				tmp,matrix[x][y] = matrix[x][y],tmp
			}
		}
	}
}
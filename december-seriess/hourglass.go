package main

import (
"fmt"
"math"
)

func main() {
    a := [][]int{
        {1, 1, 1, 0, 0, 0},
        {0, 1, 0, 0, 0, 0},
        {1, 1, 1, 0, 0, 0},
        {0, 0, 2, 4, 4, 0},
        {0, 0, 0, 2, 0, 0},
        {0, 0, 1, 2, 4, 0},
    }
   max := math.MinInt32
   for i:=0;i<=len(a)-3;i++{
	for j:=0;j<=len(a[0])-3;j++{
		sum := a[i][j]+a[i][j+1]+a[i][j+2]+a[i+1][j+1]+a[i+2][j]+a[i+2][j+1]+a[i+2][j+2]
		if sum >max{
			max=sum
		}
	}
   }
   fmt.Print(max)
}

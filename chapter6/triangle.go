package solution

// you can also use imports, for example:
// import "fmt"
// import "os"
import "sort"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
    // write your code in Go 1.4
    sort.Sort(sort.IntSlice(A))
    if len(A)<3{
        return 0
    }
    for index,p := range A{
        if index == len(A)-2{
            break
        }
        q := A[index+1]
        r := A[index+2]
        if p+q > r && q+r>p && r+p>q{
            return 1
        }
    }
    return 0
}


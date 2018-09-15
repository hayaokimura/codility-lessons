package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(S string) int {
    // write your code in Go 1.4
    var stack []rune
    for _,r := range S{
        if string(r) == "("{
            stack = append(stack,r)
        }else{
            if len(stack) > 0{
                stack = stack[:len(stack)-1]
            }else{
                return 0
            }
        }
    }
    if len(stack) != 0{
        return 0
    }
    return 1
}

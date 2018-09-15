package solution

// you can also use imports, for example:
//import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int, B []int) int {
    // write your code in Go 1.4
    var upStream,downStream []int 
    for i:=0;i<len(A);i++{
        if B[i] == 1{
            downStream = append(downStream,A[i])
        }else{
            if len(downStream) > 0{
                if downStream[len(downStream)-1] < A[i]{
                    downStream = downStream[:len(downStream)-1]
                    i--
                }
            }else{
                upStream = append(upStream,A[i])
            }
            
        }
    }
    return len(upStream)+len(downStream)
}


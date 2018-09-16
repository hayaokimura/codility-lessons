package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

type stoneStack struct {
	Stack      []int
	TotalStone int
	NowHeight  int
}

func Solution(H []int) int {
	// write your code in Go 1.4
	stone := stoneStack{
		TotalStone: 0,
		NowHeight:  0,
	}
	for _, height := range H {
		if stone.NowHeight < height {
			stone.addStone(height)
		} else if stone.NowHeight > height {
			for stone.NowHeight > height {
				stone.cutStone()
			}
			stone.addStone(height)
		}
	}
	return stone.TotalStone
}
func (stone *stoneStack) addStone(height int) {
	stoneSize := height - stone.NowHeight
	if stoneSize > 0 {
		stone.Stack = append(stone.Stack, stoneSize)
		stone.NowHeight += stoneSize
		stone.TotalStone++
	}
}
func (stone *stoneStack) cutStone() {
	stoneSize := stone.Stack[len(stone.Stack)-1]
	stone.Stack = stone.Stack[:len(stone.Stack)-1]
	stone.NowHeight -= stoneSize
}

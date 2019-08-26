package main

import (
	"fmt"
	"os"
)

type point struct {
	i, j int
}

var onestep = []point{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

//(去发现)找门
func (t *point) onestep(p point) point {
	return point{t.i + p.i, t.j + p.j}

}

//是否越界
func (t point) in(maze [][]int) (int, bool) {
	if t.i < 0 || t.i >= len(maze) { // 行是否越界
		return 0, true
	}
	if t.j < 0 || t.j >= len(maze[0]) { //列是否越界
		return 0, true
	}

	return maze[t.i][t.j], false
}

func walk(start, end point, maze [][]int) [][]int {
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}

	Q := []point{start}
	for len(Q) > 0 { //探索队列
		zero := Q[0] //0号探员
		if zero == end {
			break
		}
		Q = Q[1:]
		for _, p := range onestep {
			next := zero.onestep(p)
			//是否越界，墙是特殊边界
			if v, ok := next.in(maze); ok || v == 1 { //1是墙壁
				continue
			}
			//走过的不能走
			if v, ok := next.in(steps); ok || v != 0 { //非0是已经走过的路
				continue
			}
			//不走回头路
			if next == start {
				continue
			}

			Q = append(Q, next)
			steps[next.i][next.j] = steps[zero.i][zero.j] + 1

		}
	}
	return steps

}

func main() {
	file, _ := os.Open("map.txt")
	row, col := 0, 0
	fmt.Fscanf(file, "%d %d", &row, &col)
	fmt.Println(row, col)
	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
	L1:
		for j := range maze[i] {
			mazedata := 0
			_, err := fmt.Fscanf(file, "%d", &mazedata)
			if err != nil {
				if err.Error() == "unexpected newline" {
					goto L1
				} else {
					fmt.Println(err)
				}
			}
			maze[i][j] = mazedata
		}
	}
	for _, r := range maze {
		fmt.Println(r)
	}
	fmt.Println()

	steps := walk(point{0, 0}, point{len(maze) - 1, len(maze[0]) - 1}, maze)

	for _, m := range steps {
		for _, n := range m {
			fmt.Printf("%2d ", n)
		}
		fmt.Println()
	}

}

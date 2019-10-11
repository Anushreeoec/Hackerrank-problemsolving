package main

import (
   "fmt"
//   "io"
 //  "os"
)


func main() {
   var n int
   var k int
   var rq int
   var rc int
    _, _ = fmt.Scan(&n)
    _, _ = fmt.Scan(&k)
    _, _ = fmt.Scan(&rq)
    _, _ = fmt.Scan(&rc)

    obstacles := make([][]int, k)

   for i := 0; i < k; i++ {
	   obstacles[i] = make([]int, 2)
      for j := 0; j < 2; j++ {
	      _, _ = fmt.Scan(&obstacles[i][j])
      }
   }

   result := queensAttack(n, k, rq, rc, obstacles)

   fmt.Println(result)
}

func queensAttack(n int, k int, rq int, rc int, obstacles [][]int) int {
	m := make(map[int][]int)
	for i := 0; i < int(k); i++ {
		a := obstacles[i][0]
		arr := m[a]
		m[a] = append(arr, obstacles[i][1])
	}

	count := movesTopLeft(n, rq, rc, m)
	count += movesTop(n, rq, rc, m)
	count += movesTopRight(n, rq, rc, m)
	count += movesRight(n, rq, rc, m)
	count += movesDownRight(n, rq, rc, m)
	count += movesDown(n, rq, rc, m)
	count += movesDownLeft(n, rq, rc, m)
	count += movesLeft(n, rq, rc, m)

	return count
}

func movesTopLeft (n int, rq int, rc int, m map[int][]int) int {
	c := 0
	for rq < n && rc > 1 {
	    rq = rq + 1
	    rc = rc - 1

	    check := checkIfObstacle(m, rq, rc)
	    if check == 1 {
	       return c
	    } else {
	      c++
	    }
	}


	return c
}

func movesTop(n int, rq int, rc int, m map[int][]int) int {
	c := 0
	for rq < n {
	    rq = rq + 1

	    check := checkIfObstacle(m, rq, rc)
	    if check == 1 {
		return c
	    } else {
		c++
	    }
	}

	return c
}

func movesTopRight(n int, rq int, rc int, m map[int][]int) int {
        c := 0
        for rq < n && rc < n{
            rq = rq + 1
	    rc = rc + 1

            check := checkIfObstacle(m, rq, rc)
            if check == 1 {
                return c
            } else {
                c++
            }
        }

        return c

}

func movesRight(n int, rq int, rc int, m map[int][]int) int {
        c := 0
        for rc < n {
            rc = rc + 1

            check := checkIfObstacle(m, rq, rc)
            if check == 1 {
                return c
            } else {
                c++
            }
        }

        return c

}

func movesDownRight(n int, rq int, rc int, m map[int][]int) int {
        c := 0
        for rc < n && rq > 1{
            rq = rq - 1
            rc = rc + 1

            check := checkIfObstacle(m, rq, rc)
            if check == 1 {
                return c
            } else {
                c++
            }
        }

        return c

}

func movesDown(n int, rq int, rc int, m map[int][]int) int {
        c := 0
        for rq > 1 {
            rq = rq - 1

            check := checkIfObstacle(m, rq, rc)
            if check == 1 {
                return c
            } else {
                c++
            }
        }

        return c

}


func movesDownLeft(n int, rq int, rc int, m map[int][]int) int {
        c := 0
        for rq > 1 && rc > 1 {
            rq = rq - 1
            rc = rc - 1

            check := checkIfObstacle(m, rq, rc)
            if check == 1 {
                return c
            } else {
                c++
            }
        }

        return c

}

func movesLeft(n int, rq int, rc int, m map[int][]int) int {
        c := 0
        for rc > 1 {
            rc = rc - 1

            check := checkIfObstacle(m, rq, rc)
            if check == 1 {
                return c
            } else {
                c++
            }
        }

        return c

}



func checkIfObstacle(m map[int][]int, rq int, rc int) int {
	var arr []int
	for _, v := range m[rq] {
	   arr = append(arr, v)
	}
	for _, v := range arr {
	   if v == rc {
	      return 1
	   }
	}
	return 0
}


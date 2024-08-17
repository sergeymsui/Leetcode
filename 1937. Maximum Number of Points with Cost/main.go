// 1937. Максимальное количество очков со стоимостью

// Вам дана m x nцелочисленная матрица points( с индексом 0 ). Начиная с 0точек, вы хотите максимизировать количество точек, которые вы можете получить из матрицы.
// Чтобы получить очки, вы должны выбрать одну ячейку в каждой строке . Выбор ячейки с координатами (r, c)добавит вам points[r][c] очков.
// Однако вы потеряете очки, если выберете ячейку, которая находится слишком далеко от ячейки, выбранной в предыдущей строке. 
// Для каждых двух соседних строк rи r + 1(где 0 <= r < m - 1), выбор ячеек с координатами и будет вычитаться из вашего счета.(r, c1)(r + 1, c2) abs(c1 - c2)

// Укажите максимальное количество баллов, которое вы можете набрать .

// abs(x)определяется как:
// x для x >= 0.
// -x для x < 0.

package main

import "fmt"

func maxPoints(points [][]int) int64 {
	v := len(points)
	h := len(points[0])

	dp := []int{}

	for i := 0; i < h; i++ {
		dp = append(dp, points[0][i])
	}

	for i := 1; i < v; i++ {

		left_arr := []int{}
		right_arr := []int{}

		left_arr = append(left_arr, dp[0])
		right_arr = append(right_arr, dp[h-1]-(h-1))

		for j := 1; j < h; j++ {
			left_arr = append(left_arr, max(left_arr[j-1], dp[j]+j))
		}

		for j := h - 2; j >= 0; j-- {
			right_arr = append(right_arr, max(right_arr[len(right_arr)-1], dp[j]-j))
		}

		newDp := []int{}

		// fmt.Println(len(left_arr))
		// fmt.Println(len(right_arr))

		for j := 0; j < h; j++ {
			newDp = append(newDp, max(left_arr[j]-j, right_arr[h-1-j]+j)+points[i][j])
		}

		dp = newDp
	}

	maxValue := dp[0]

	for i := 1; i < len(dp); i++ {
		maxValue = max(maxValue, dp[i])
	}

	return int64(maxValue)
}

func main() {

	arr_case := [][]int{
		[]int{1, 2, 3},
		[]int{1, 5, 1},
		[]int{3, 1, 1},
	}

	fmt.Println(maxPoints(arr_case))

}

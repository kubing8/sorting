package sorts

/*
	type TypeNumber interface {
		int | int64 | float64
	}
 */

//Асимптомтика (худший / средний / лучший): O(n^2) / O(n^2) / O(n)
func Shaker[T TypeNumber](arr []T) []T {
	if len(arr) <= 1 {
		return arr
	}
	beg := 0
	end := len(arr) - 1

	for beg < end {
		for i := beg; i < end; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}
		end--

		for i := end; i > beg; i-- {
			if arr[i-1] > arr[i] {
				arr[i], arr[i-1] = arr[i-1], arr[i]
			}
		}
		beg++
	}
	return arr
}
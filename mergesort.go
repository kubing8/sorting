package sorts

/*
	type TypeNumber interface {
		int | int64 | float64
	}
*/

//Асимптомтика (худший / средний / лучший): O(n^2) / O(nLogn) / O(nLogn)
func Merge[T TypeNumber](arr []T) []T {
	if len(arr) < 2 {
		return arr
	}
	mid := len(arr) / 2

	return mergeSort(Merge(arr[:mid]), Merge(arr[mid:]))
}

func mergeSort[T TypeNumber](left, right []T) []T{
	length, i, j := len(left)+len(right), 0, 0
	merged := make([]T, length, length)

	for k := 0; k < length; k++ {
		if j >= len(right) || (i < len(left) && left[i] <= right[j]) {
			merged[k] = left[i]
			i++
		} else {
			merged[k] = right[j]
			j++
		}
	}
	return merged
}

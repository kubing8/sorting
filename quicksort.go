package sorts

/*
	type TypeNumber interface {
		int | int64 | float64
	}
*/

//Асимптомтика (худший / средний / лучший): O(n^2) / O(nLogn) / O(nLogn)
func Quick[T TypeNumber](arr []T) []T {
	if len(arr) < 2 {
		return arr
	}
	return quickAlg(arr,0, len(arr)-1)
}

func quickAlg[T TypeNumber](arr []T, beg, end int) []T {
	if beg < end {
		partitionIndex := partition(arr, beg, end)

		quickAlg(arr, beg, partitionIndex-1)
		quickAlg(arr, partitionIndex+1, end)
	}

	return arr
}

func partition[T TypeNumber](arr []T, beg int, end int) int {
	index := beg

	for i := beg + 1; i <= end; i++ {
		if arr[i] < arr[beg] {
			index++
			arr[i], arr[index] = arr[index], arr[i]
		}
	}

	arr[beg], arr[index] = arr[index], arr[beg]
	return index
}
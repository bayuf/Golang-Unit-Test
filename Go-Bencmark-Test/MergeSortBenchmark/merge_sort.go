package mergesort

import "fmt"

func MergeSort(datas []int) []int {
	// Langsung mengembalikan data jika hanya ada satu data
	if len(datas) <= 1 {
		return datas
	}

	// membagi data menjadi dua bagian. bagian kiri dan kanan
	mid := len(datas) / 2
	left := make([]int, mid)
	right := make([]int, len(datas)-mid)

	copy(left, datas[:mid])
	copy(right, datas[mid:])

	// Recursif membagi data sampai terkecil
	left = MergeSort(left)
	right = MergeSort(right)

	// data kecil siap untuk di urutkan
	return Merge(left, right)
}

func Merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	for i < len(left) {
		result = append(result, left[i])
		i++
	}

	for j < len(right) {
		result = append(result, right[j])
		j++
	}

	return result

}

func StartSort(data []int) {
	// datas := []int{9, 8, 4, 1, 3, 4, 1, 9, 0, 7, 5, 4, 2, 4, 0, 10, 14, 11, 15, 13, 9, 15, 3, 0, 12}

	datas := data
	fmt.Println("=========Data Acak=========\n", datas)

	fmt.Println("=========Data Urut=========\n", MergeSort(datas))

}

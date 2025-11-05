package mergesort

import "testing"

func BenchmarkMergeSortSub(b *testing.B) {
	// Subtest
	b.Run("Merge1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			StartSort([]int{9, 9, 3, 5, 1, 6, 8, 5, 4, 3, 6, 2, 3, 5, 0, 1})
		}
	})
	b.Run("Merge2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			StartSort([]int{100, 78, 34, 56, 23, 98, 45, 23, 43, 22, 85, 54, 34, 81, 11})
		}
	})
}
func BenchmarkMergeSort(b *testing.B) {
	datas := []int{10, 9, 8, 7, 6, 3, 2, 4, 1, 0}
	for i := 0; i < b.N; i++ {
		StartSort(datas)
	}
}

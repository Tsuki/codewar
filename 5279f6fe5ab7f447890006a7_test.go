package kata_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type PosPeaks struct {
	Pos   []int
	Peaks []int
}

func PickPeaks(array []int) PosPeaks {
	pos := make([]int, 0)
	peaks := make([]int, 0)
	var cur int
	for i := 1; i < len(array)-1; i++ {
		if array[i-1] >= array[i] {
			continue
		}
		if array[i-1] < array[i] && array[i] > array[i+1] {
			pos = append(pos, i)
			peaks = append(peaks, array[i])
			continue
		}
		if array[i] == array[i+1] {
			cur = i
			for i < len(array)-2 {
				i++
				if array[i] > array[i+1] {
					pos = append(pos, cur)
					peaks = append(peaks, array[i])
					break
				} else if array[i] < array[i+1] {
					break
				}
			}
		}
	}
	return PosPeaks{pos, peaks}
}

func pickPeaksTest(array []int, posPeaks PosPeaks) {
	var ans = PickPeaks(array)
	Expect(ans).To(Equal(posPeaks))
}

var _ = Describe("Test Example", func() {
	It("should support finding peaks", func() {
		pickPeaksTest(
			[]int{1, 2, 3, 6, 4, 1, 2, 3, 2, 1},
			PosPeaks{Pos: []int{3, 7}, Peaks: []int{6, 3}},
		)
	})

	It("should support finding peaks, but should ignore peaks on the edge of the array", func() {
		pickPeaksTest(
			[]int{3, 2, 3, 6, 4, 1, 2, 3, 2, 1, 2, 3},
			PosPeaks{Pos: []int{3, 7}, Peaks: []int{6, 3}},
		)
	})

	It("should support finding peaks; if the peak is a plateau, it should only return the position of the first element of the plateau", func() {
		pickPeaksTest(
			[]int{3, 2, 3, 6, 4, 1, 2, 3, 2, 1, 2, 2, 2, 1},
			PosPeaks{Pos: []int{3, 7, 10}, Peaks: []int{6, 3, 2}},
		)
	})

	It("should support finding peaks; if the peak is a plateau, it should only return the position of the first element of the plateau", func() {
		pickPeaksTest(
			[]int{2, 1, 3, 1, 2, 2, 2, 2, 1},
			PosPeaks{Pos: []int{2, 4}, Peaks: []int{3, 2}},
		)
	})

	It("should support finding peaks, but should ignore peaks on the edge of the array", func() {
		pickPeaksTest(
			[]int{2, 1, 3, 1, 2, 2, 2, 2},
			PosPeaks{Pos: []int{2}, Peaks: []int{3}},
		)
	})

	It("should support finding peaks, but should ignore peaks on the edge of the array", func() {
		pickPeaksTest(
			[]int{6, -1, -2, 5, 13, 6, 9, -2, 3, 3, 5, -5},
			PosPeaks{Pos: []int{4, 6, 10}, Peaks: []int{13, 9, 5}},
		)
	})
	It("should support finding peaks, but should ignore peaks on the edge of the array", func() {
		pickPeaksTest(
			[]int{9, 9, 2, 5, 2, 4, -5, -4, 8, 1},
			PosPeaks{Pos: []int{3, 5, 8}, Peaks: []int{5, 4, 8}},
		)
	})
	It("should support finding peaks, but should ignore peaks on the edge of the array", func() {
		pickPeaksTest(
			[]int{3, -5, 14, 11, 11, -4, 12, 12, 7},
			PosPeaks{Pos: []int{2, 6}, Peaks: []int{14, 12}},
		)
	})
	It("should support finding peaks, but should ignore peaks on the edge of the array", func() {
		pickPeaksTest(
			[]int{2, 1, 3, 2, 2, 2, 2, 1},
			PosPeaks{Pos: []int{2}, Peaks: []int{3}},
		)
	})
})

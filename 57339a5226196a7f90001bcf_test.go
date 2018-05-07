package kata

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("N-Point Crossover", func() {

	It("should work when there is no cross-point index", func() {
		xs, ys := Crossover([]int{}, []int{1, 1, 1, 1, 1}, []int{2, 2, 2, 2, 2})
		Expect(xs).To(Equal([]int{1, 1, 1, 1, 1}))
		Expect(ys).To(Equal([]int{2, 2, 2, 2, 2}))
	})

	It("should work with 1 cross-point index", func() {
		xs, ys := Crossover([]int{1}, []int{1, 1, 1, 1, 1}, []int{2, 2, 2, 2, 2})
		Expect(xs).To(Equal([]int{1, 2, 2, 2, 2}))
		Expect(ys).To(Equal([]int{2, 1, 1, 1, 1}))
	})

	It("should work with 1 repeated cross-point index", func() {
		xs, ys := Crossover([]int{1, 1}, []int{1, 1, 1, 1, 1}, []int{2, 2, 2, 2, 2})
		Expect(xs).To(Equal([]int{1, 2, 2, 2, 2}))
		Expect(ys).To(Equal([]int{2, 1, 1, 1, 1}))
	})

	It("should work for 2 cross-point indices", func() {
		xs, ys := Crossover([]int{1, 3}, []int{1, 1, 1, 1, 1}, []int{2, 2, 2, 2, 2})
		Expect(xs).To(Equal([]int{1, 2, 2, 1, 1}))
		Expect(ys).To(Equal([]int{2, 1, 1, 2, 2}))
	})

	It("should work for 3 cross-point indices", func() {
		xs, ys := Crossover([]int{1, 3, 5}, []int{1, 1, 1, 1, 1, 1, 1}, []int{2, 2, 2, 2, 2, 2, 2})
		Expect(xs).To(Equal([]int{1, 2, 2, 1, 1, 2, 2}))
		Expect(ys).To(Equal([]int{2, 1, 1, 2, 2, 1, 1}))
	})

	It("should work for unordered and repeated cross-point indices", func() {
		xs, ys := Crossover([]int{3, 5, 1, 1, 3}, []int{1, 1, 1, 1, 1, 1, 1}, []int{2, 2, 2, 2, 2, 2, 2})
		Expect(xs).To(Equal([]int{1, 2, 2, 1, 1, 2, 2}))
		Expect(ys).To(Equal([]int{2, 1, 1, 2, 2, 1, 1}))
	})

})

func Crossover(ns []int, xs []int, ys []int) ([]int, []int) {
	var crossover = false
	for k := range xs {
		for _, v := range ns {
			if k == v {
				crossover = !crossover
				break
			}
		}
		if crossover {
			xs[k], ys[k] = ys[k], xs[k]
		}
	}
	return xs, ys
}

package kata_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "fmt"
)

func doJosephus(items []interface{}, k int, exp []interface{}) {
	Println("input:", items, k)
	var ans = Josephus(items, k)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("Sample Tests", func() {
	It("should handle basic test cases", func() {
		items := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		result := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		k := 1
		doJosephus(items, k, result)
		items = []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		result = []interface{}{2, 4, 6, 8, 10, 3, 7, 1, 9, 5}
		k = 2
		doJosephus(items, k, result)
		items = []interface{}{1, 2, 3, 4, 5, 6, 7}
		result = []interface{}{3, 6, 2, 7, 5, 1, 4}
		k = 3
		doJosephus(items, k, result)
	})

	It("should handle basic test cases", func() {
		items := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
		result := []interface{}{10, 7, 8, 13, 5, 4, 12, 11, 3, 15, 14, 9, 1, 6, 2}
		k := 40
		doJosephus(items, k, result)
	})
})

func Josephus(items []interface{}, k int) []interface{} {
	result := make([]interface{}, 0)
	y := 0
	for len(items) > 0 {
		if y = k; k > len(items) {
			y = k % len(items)
		}
		if y == 0 {
			y = len(items)
		}
		result = append(result, items[y-1])
		items = append(items[y:], items[:y-1]...)
	}
	return result
}

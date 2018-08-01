package kata_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Test Example", func() {
	It("should handle basic cases", func() {
		Expect(ProductFib(4895)).To(Equal([3]uint64{55, 89, 1}))
	})
	It("should handle basic cases", func() {
		Expect(ProductFib(5895)).To(Equal([3]uint64{89, 144, 0}))
	})
	It("should handle basic cases", func() {
		Expect(ProductFib(74049690)).To(Equal([3]uint64{6765, 10946, 1}))
	})
	It("should handle basic cases", func() {
		Expect(ProductFib(84049690)).To(Equal([3]uint64{10946, 17711, 0}))
	})
})

func ProductFib(prod uint64) [3]uint64 {
	fib := []uint64{0, 1, 1, 2, 3}
	result := uint64(0)
	var times, sum uint64
	for times < prod {
		times = 1
		sum = 0
		for _, item := range fib[len(fib)-2:] {
			sum += item
		}
		fib = append(fib, sum)
		for _, item := range fib[len(fib)-2:] {
			times *= item
		}
	}
	if prod == times {
		result = 1
	}
	return [3]uint64{fib[len(fib)-2], fib[len(fib)-1], result}
}

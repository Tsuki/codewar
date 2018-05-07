package kata

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SumEvenFibonacci", func() {
	It("should return 10 for input 8", func() {
		Expect(SumEvenFibonacci(8)).To(Equal(10))
	})
	It("should return 60696 for input 111111", func() {
		Expect(SumEvenFibonacci(111111)).To(Equal(60696))
	})
	It("should return 4613732 for input 8675309", func() {
		Expect(SumEvenFibonacci(8675309)).To(Equal(4613732))
	})
})

//func SumEvenFibonacci(limit int) int {
//	if limit < 2 {
//		return 0
//	}
//	result := make([]int, 100)
//	result = append(result, 0)
//	result = append(result, 1)
//	for result[len(result)-1] < limit {
//		result = append(result, result[len(result)-1]+result[len(result)-2])
//	}
//	done := 0
//	for _, v := range result {
//		if v%2 == 0 {
//			done += v
//		}
//	}
//	return done
//}

func SumEvenFibonacci(limit int) int {
	sum, a, b := 2, 1, 2
	for a+b <= limit {
		a, b = b, a+b
		if b%2 == 0 {
			sum += b
		}
	}
	return sum
}

package kata_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"math/big"
	"math"
)

func dotest(n int, exp float64) {
	var ans = Going(n)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("Test Example", func() {

	It("should handle basic cases", func() {
		dotest(5, 1.275)
		dotest(6, 1.2125)
		dotest(7, 1.173214)
		dotest(8, 1.146651)
		dotest(30, 1.034525)
	})
})

func Going(n int) float64 {
	sum, fat := new(big.Float), big.NewFloat(1)
	for i := 1; i <= n; i++ {
		fat = fat.Mul(fat, big.NewFloat(float64(i)))
		sum = sum.Add(sum, fat)
	}
	result := sum.Quo(sum, fat)
	result2, _ := result.Float64()
	return math.Floor(result2*1e6) / 1e6
}

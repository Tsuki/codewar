package kata_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"math"
)

func EquableTriangle(a, b, c int) bool {
	s := (a + b + c) / 2
	return math.Sqrt(float64(s*(s-a)*(s-b)*(s-c))) == float64(a+b+c)
}

var _ = Describe("Example Test Cases", func() {
	It("should return true", func() {
		Expect(EquableTriangle(5, 12, 13)).To(Equal(true))
	})
	It("should return false", func() {
		Expect(EquableTriangle(2, 3, 4)).To(Equal(false))
	})
})

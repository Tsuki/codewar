package kata_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"strings"
)

var _ = Describe("test roman to decimal converter", func() {
	It("should give decimal number from roman", func() {
		Expect(Decode("XXI")).To(Equal(21))

	})
	It("should give decimal number from roman", func() {
		Expect(Decode("I")).To(Equal(1))
	})
	It("should give decimal number from roman", func() {
		Expect(Decode("IV")).To(Equal(4))
	})
	It("should give decimal number from roman", func() {
		Expect(Decode("MMVIII")).To(Equal(2008))
	})
	It("should give decimal number from roman", func() {
		Expect(Decode("MDCLXVI")).To(Equal(1666))
	})
})

func Decode(roman string) int {
	var result []int
	var sum int
	arrs := strings.Split(roman, "")
	for _, value := range arrs {
		switch value {
		case "I":
			result = append(result, 1)
		case "V":
			result = append(result, 5)
		case "X":
			result = append(result, 10)
		case "L":
			result = append(result, 50)
		case "C":
			result = append(result, 100)
		case "D":
			result = append(result, 500)
		case "M":
			result = append(result, 1000)
		}
	}
	for i := 0; i < len(result)-1; i++ {
		if result[i] < result[i+1] {
			result[i] = -result[i]
		}
	}
	for _, value := range result {
		sum += value
	}
	return sum
}

package kata_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Example Tests", func() {
	It("passes example tests", func() {
		Expect(ValidParentheses("()")).To(Equal(true))
		Expect(ValidParentheses(")")).To(Equal(false))
	})
})

func ValidParentheses(parens string) bool {
	open, clo := 0, 0
	for _, v := range parens {
		if v == 40 {
			open++
		} else {
			clo++
		}
		if clo > open {
			return false
		}
	}
	return open == clo
}

package kata

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("RemoveChar", func() {
	Describe("Fixed Tests", func() {
		It("eloquent", func() {
			Expect(RemoveChar("eloquent")).To(Equal("loquen"))
		})
		It("country", func() {
			Expect(RemoveChar("country")).To(Equal("ountr"))
		})
		It("person", func() {
			Expect(RemoveChar("person")).To(Equal("erso"))
		})
		It("place", func() {
			Expect(RemoveChar("place")).To(Equal("lac"))
		})
	})
})

func RemoveChar(word string) string {
	return word[1 : len(word)-1]
}

package kata_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Test Example", func() {
	It("should return true", func() {
		Expect(Solution("abcdef")).To(Equal([]string{"ab", "cd", "ef"}))
	})
	It("should return false", func() {
		Expect(Solution("abc")).To(Equal([]string{"ab", "c_"}))
	})
})

//func Solution(str string) []string {
//	result := make([]string, 0)
//	arr := strings.Split(str, "")
//	for i := 0; i+1 <= len(str); i += 2 {
//		if i+2 <= len(str) {
//			result = append(result, strings.Join(arr[i:i+2], ""))
//		} else {
//			result = append(result, strings.Join(append(arr[i:i+1], "_"), ""))
//		}
//	}
//	return result
//}
func Solution(str string) []string {
	result := make([]string, 0)
	if len(str)%2 != 0 {
		str += "_"
	}
	for i := 0; i+1 <= len(str); i += 2 {
		result = append(result, str[i:i+2])
	}
	return result
}

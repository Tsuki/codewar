package kata_test

import (
	"strings"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func LongestConsec(strarr []string, k int) string {
	if k <= 0 || len(strarr) == 0 || len(strarr) < k {
		return ""
	}
	index, max := 0, 0
	for i := 0; i <= len(strarr)-k; i++ {
		if len(strings.Join(strarr[i:i+k], "")) > max {
			index, max = i, len(strings.Join(strarr[i:i+k], ""))
		}
	}
	return strings.Join(strarr[index:index+k], "")
}
func doTest(strarr []string, k int, exp string) {
	var ans = LongestConsec(strarr, k)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("Test Example", func() {

	It("should handle basic cases", func() {
		doTest([]string{"zone", "abigail", "theta", "form", "libe", "zas"}, 2, "abigailtheta")
		doTest([]string{"ejjjjmmtthh", "zxxuueeg", "aanlljrrrxx", "dqqqaaabbb", "oocccffuucccjjjkkkjyyyeehh"}, 1,
			"oocccffuucccjjjkkkjyyyeehh")
		doTest([]string{}, 3, "")
		doTest([]string{"itvayloxrp", "wkppqsztdkmvcuwvereiupccauycnjutlv", "vweqilsfytihvrzlaodfixoyxvyuyvgpck"}, 2,
			"wkppqsztdkmvcuwvereiupccauycnjutlvvweqilsfytihvrzlaodfixoyxvyuyvgpck")
	})
})

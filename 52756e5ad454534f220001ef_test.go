package kata_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func testLCS(s1, s2, s string) {
	var ans = LCS(s1, s2)
	Expect(ans).To(Equal(s))
}

var _ = Describe("Sample Test Cases", func() {

	XIt("LCS should work on sample case #1", func() {
		testLCS("a", "b", "")
	})

	XIt("LCS should work on sample case #1", func() {
		testLCS("MZJAWXU", "XMJYAUZ", "MJAU")
	})

	XIt("LCS should work on sample case #2", func() {
		testLCS("abcdef", "abc", "abc")
	})

	XIt("LCS should work on sample case #3", func() {
		testLCS("132535365", "123456789", "12356")
	})
})
//FIXME https://en.wikipedia.org/wiki/Longest_common_subsequence_problem#First_property
func LCS(s1, s2 string) string {
	result := ""
	a := make([][]uint8, len(s2)+1)
	for i := range a {
		a[i] = make([]uint8, len(s1)+1)
	}
	for k, v := range s1 {
		for k2, v2 := range s2 {
			if v == v2 {
				a[k2+1][k+1] = a[k2][k] + 1
			} else if a[k2][k+1] < a[k2+1][k] {
				a[k2+1][k+1] = a[k2+1][k]
			} else {
				a[k2+1][k+1] = a[k2][k+1]
			}
		}
	}
	for h := len(s2); h > 1; h-- {
		for v := len(s1); v > 1; v-- {
			if a[h][v] == a[h-1][v] {
				//fmt.Println(a[h][v])
			}
		}
	}
	//for _, v := range a {
	//	fmt.Println(v)
	//}
	return result
}

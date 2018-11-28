package leetcode_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var result ListNode
	var next *ListNode
	var becometen bool
	next = &result
	for l1 != nil || l2 != nil {
		sum := 0
		if l1 == nil {
			l1 = &ListNode{Val: 0}
		} else if l2 == nil {
			l2 = &ListNode{Val: 0}
		}
		sum = l1.Val + l2.Val
		if becometen {
			sum = sum + 1
		}
		becometen = sum >= 10
		if becometen {
			next.Val = sum - 10
		} else {
			next.Val = sum
		}
		if l1.Next == nil && l2.Next == nil {
			if becometen {
				next.Next = &ListNode{Val: 1}
			}
			return &result
		}
		next.Next = &ListNode{}
		l1 = l1.Next
		l2 = l2.Next
		next = next.Next
	}
	return &result
}
func TestAddTwoNumber(t *testing.T) {
	var a = &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	var b = &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	t.Logf("%#v", addTwoNumbers(a, b))
	//Expect(addTwoNumbers(a, b)).To(Equal(result))
}
func TestAddTwoNumber2(t *testing.T) {
	var a = &ListNode{1, &ListNode{8, nil}}
	var b = &ListNode{0, nil}
	t.Logf("%#v", addTwoNumbers(a, b))
	//Expect(addTwoNumbers(a, b)).To(Equal(result))
}

var _ = Describe("Basic tests", func() {
	It("should return the correct values", func() {
		var a = &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
		var b = &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
		var result = &ListNode{7, &ListNode{0, &ListNode{8, nil}}}
		Expect(addTwoNumbers(a, b)).To(ConsistOf(result))
	})
})

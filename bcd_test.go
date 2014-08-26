package bcd

import (
	. "github.com/r7kamura/gospel"
	"testing"
)

func TestDescribe(t *testing.T) {
	Describe(t, "bcd to int", func() {
		Context("bcd 0", func() {
			It("should be 0", func() {
				Expect(BcdToInt(0)).To(Equal, 0)
			})
		})
		Context("bcd 20", func() {
			It("should be 14", func() {
				Expect(BcdToInt(20)).To(Equal, 14)
			})
		})
	})
	Describe(t, "int to bcd", func() {
		Context("int 0", func() {
			It("should be 0", func() {
				Expect(BcdToInt(0)).To(Equal, 0)
			})
		})
		Context("int 14", func() {
			It("should be 20", func() {
				Expect(IntToBcd(14)).To(Equal, 20)
			})
		})
	})
}

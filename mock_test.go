package emock_test

import (
	. "github.com/ebabani/emock"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/ebabani/emock/matchers"

	"fmt"
)

var _ = Describe("Mock", func() {
	Describe("MockFunc", func() {
		Describe("when mocking a function with one return", func() {
			var myFunc func(a int, b string) string

			var ret string
			var mockObj *Mock

			BeforeEach(func() {
				myFunc = func(a int, b string) string {
					return fmt.Sprintf("A IS %+v B is %+v", a, b)
				}

				mockObj = MockFunc(&myFunc)
				ret = myFunc(123, "ERGIN")
			})

			It("should return an empty string", func() {
				Expect(ret).To(BeEmpty())
			})

			It("should record the right args", func() {
				Expect(mockObj.GetArgsForCall(0)).To(MatchArgs(123, "ERGIN"))
			})

			It("should have a call count of 1", func() {
				Expect(mockObj.CallCount()).To(Equal(1))
			})

			Describe("Restore", func() {
				BeforeEach(func() {
					mockObj.Restore()
					ret = myFunc(123, "ERGIN")
				})

				It("should call the original function", func() {
					Expect(ret).To(Equal("A IS 123 B is ERGIN"));
				})
			})
		})

		Describe("when mocking a function with multiple returns", func() {
			var myFunc func(a int, b string) (string, func(int));

			var ret1 string
			var ret2 func(int)

			var mockObj *Mock

			var called bool;
			var funcReturn = func(int) { called = true }
			BeforeEach(func() {
				called = false;

				myFunc = func(a int, b string) (string, func(int)){
					return fmt.Sprintf("A IS %+v B is %+v", a, b), funcReturn
				}

				mockObj = MockFunc(&myFunc)
				ret1, ret2 = myFunc(123, "ERGIN")
			})

			It("should return zero values for both returns", func() {
				Expect(ret1).To(BeEmpty())
				Expect(ret2).To(BeZero())
			})

			Describe("restore", func() {
				BeforeEach(func() {
					mockObj.Restore()
					ret1, ret2 = myFunc(123, "ERGIN")
				})

				It("should call the original function", func() {
					Expect(ret1).To(Equal("A IS 123 B is ERGIN"))
					ret2(0)
					Expect(called).To(BeTrue())
				})
			})
		})

		Describe("when mocking a function with no returns", func() {
			var mockObj *Mock

			var called bool
			var myFunc func(a int, b string)

			BeforeEach(func() {
				called = false
				myFunc = func(a int, b string) {
					called = true
				}

				mockObj = MockFunc(&myFunc)
				myFunc(123, "ERGIN")
			})

			It("should not call the original function", func() {
				Expect(called).To(BeFalse())
			})

			Describe("restore", func() {
				BeforeEach(func() {
					mockObj.Restore()
					myFunc(123, "ERGIN")
				})

				It("should call the original function", func() {
					Expect(called).To(BeTrue())
				})
			})
		})

		Describe("when mocking a function with a channel return", func() {
			var mockObj *Mock

			var called bool
			var myFunc func(a int, b string) (chan bool)

			var ret chan bool;

			BeforeEach(func() {
				called = false
				myFunc = func(a int, b string) (chan bool) {
					ret := make(chan bool, 1)
					ret <- true
					return ret
				}

				mockObj = MockFunc(&myFunc)
				ret = myFunc(123, "ERGIN")
			})

			It("should not call the original function", func() {
				Expect(ret).To(BeNil())
			})

			Describe("restore", func() {
				BeforeEach(func() {
					mockObj.Restore()
					ret = myFunc(123, "ERGIN")
				})

				It("should call the original function", func() {
					Expect(ret).To(Receive(Equal(true)))
				})
			})
		})
	})

	Describe("SetReturns", func() {
		Describe("when calling the mock with custom return values set", func() {
				var myFunc func(a int, b string) string

			var ret string
			var mockObj *Mock

			BeforeEach(func() {
				myFunc = func(a int, b string) string {
					return fmt.Sprintf("A IS %+v B is %+v", a, b)
				}

				mockObj = MockFunc(&myFunc)
				mockObj.SetReturns("AAAA")
				ret = myFunc(123, "ERGIN")
			})

			It("should return the specified return values", func() {
				Expect(ret).To(Equal("AAAA"))
			})

			It("should record the right args", func() {
				Expect(mockObj.GetArgsForCall(0)).To(MatchArgs(123, "ERGIN"))
			})

			It("should have a call count of 1", func() {
				Expect(mockObj.CallCount()).To(Equal(1))
			})

			Describe("Restore", func() {
				BeforeEach(func() {
					mockObj.Restore()
					ret = myFunc(123, "ERGIN")
				})

				It("should call the original function", func() {
					Expect(ret).To(Equal("A IS 123 B is ERGIN"));
				})
			})
		})
	})

	Describe("SetReturnFunc", func() {
		Describe("when mocking a function with one return", func() {
			var myFunc func(a int, b string) string

			var ret string
			var mockObj *Mock

			BeforeEach(func() {
				myFunc = func(a int, b string) string {
					return fmt.Sprintf("A IS %+v B is %+v", a, b)
				}

				mockObj = MockFunc(&myFunc)
				mockObj.SetReturnFunc(func(a int, b string) string {return "AAAA"})
				ret = myFunc(123, "ERGIN")
			})

			It("should return an empty string", func() {
				Expect(ret).To(Equal("AAAA"))
			})

			It("should record the right args", func() {
				Expect(mockObj.GetArgsForCall(0)).To(MatchArgs(123, "ERGIN"))
			})

			It("should have a call count of 1", func() {
				Expect(mockObj.CallCount()).To(Equal(1))
			})

			Describe("Restore", func() {
				BeforeEach(func() {
					mockObj.Restore()
					ret = myFunc(123, "ERGIN")
				})

				It("should call the original function", func() {
					Expect(ret).To(Equal("A IS 123 B is ERGIN"));
				})
			})
		})
	})
})

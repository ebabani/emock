# xmock
Lightweight mocking library for function calls in Go

The purpose of this library is to allow easy mocking of function calls, and ease dependency injection.

## Get it: 

`
go get github.com/ebabani/xmock
`

## Docs

http://godoc.org/github.com/ebabani/xmock


## Mocking:

### Example
The following example uses ginkgo and gomega for assertions
```
myFunc = func(a int, b string) string {
	return fmt.Sprintf("A IS %+v B is %+v", a, b)
}

It("should return an empty string", func() {
  mockObj = MockFunc(&myFunc) // (1)
  defer mockObj.Restore() // (2)
  
  ret := myFunc(123, "Cool stuff") // (3)
  
  Expect(mockObj.GetArgsForCall(0)).To(MatchArgs(123, "Cool stuff")) // (4)
  Expect(mockObj.CallCount()).To(Equal(1)) // (5)
  Expect(ret).To(BeEmpty()) // (6)
})
```

#### 1 - Create a mock object
Mock a function at the given address, and return a mock object. The mock object can be used to check how many times the focked function was called, and with what arguments. (See lines 4 and 5)

#### 2 - Restore the original function implementation
Restore the function to behave as normal instead of the mock implemetation.

#### 3 - Call our function with args.
When mocked, by default the function will return zero values, but you can also set custom return values. 

#### 4 - Check the args of the first call to `myFuncc`. 
GetArgsForCall(i) will return the args used in thte ith call to the mocked function. The args are returned as an []interface{}

MatchArgs is a custom gomega matcher to make it easier to check if the function was called with the right args. 

#### 5 Call count of the mocked function
How many times the mocked function was called.

#### 6 Check the return value of the mocked function
By default when smoething is mocked it will return zero values for all its returns. (See https://golang.org/ref/spec#The_zero_value)

You can use `mockObj.SetReturns` to set custom returns, or `mockObj.SetReturnFunc` to replace the mocked function with another one with the same signature. 




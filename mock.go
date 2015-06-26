package emock
import (
	"log"
	"reflect"
)

// If the function being mocked returns a func, that function will be nil
type Mock struct {
	originalFuncAddr reflect.Value
	originalFuncValue reflect.Value

	Calls [][]interface{}

	returns []interface{}
}

// Restore the original behaviour of the function that was being mocked.
func (self *Mock) Restore() {
	lg("OriginalFuncAddr", self.originalFuncAddr)
	lg("OriginalFuncValue", self.originalFuncValue)
	self.originalFuncAddr.Elem().Set(self.originalFuncValue)
}

// Returns the args for the n'th call to the mocked function
func (self *Mock) GetArgsForCall(n int) (interface{}) {
	args := self.Calls[n]
	return args
}

// Returns the number of calls made to the mocked function.
func (self *Mock) CallCount() int {
	return len(self.Calls)
}

func lg(data string, item interface{}) {
	return
	log.Printf("*** %s %+V", data, item)
	log.Printf("*** %s %+v", data, item)
}

// Mock the function at the specified address.
//
// By default the function will return zero values.
func MockFunc(funcAddress interface{}) *Mock {
	mock := Mock{}

	lg("funcToMock", funcAddress)

	funcValue := reflect.ValueOf(funcAddress)
	lg("funcValue", funcValue)

	mock.originalFuncAddr = funcValue

	funcType := funcValue.Type()
	lg("funcType", funcType)

	funcElem := funcValue.Elem()
	lg("funcElem", funcElem)

	mock.originalFuncValue = reflect.New(funcValue.Elem().Type()).Elem()
	mock.originalFuncValue.Set(funcElem)

	lg("mock.OriginalFuncAddr", mock.originalFuncAddr)
	lg("mock.OriginalFuncValue", mock.originalFuncValue)


	funcElemType := funcElem.Type()
	lg("funcElemType", funcElemType)

	fake := makeFuncStub(&mock, nil)


	mock.originalFuncAddr.Elem().Set(reflect.MakeFunc(mock.originalFuncValue.Type(), fake))
	return &mock
}

// Set what the mocked function should return. Panic if the number of return values
// doesn't match the original function, or if the types don't match
func (self *Mock) SetReturns(returns ...interface{}) {
	fake := makeFuncStub(self, returns)
	self.originalFuncAddr.Elem().Set(reflect.MakeFunc(self.originalFuncValue.Type(), fake))
}

// Replace the mocked function with the provided one. Panics if the given function
// signature does not match the original one.
func (self *Mock) SetReturnFunc(function interface{}) {
	wrap := wrapFunc(self, function)
	self.originalFuncAddr.Elem().Set(reflect.MakeFunc(self.originalFuncValue.Type(), wrap))
}

func makeFuncStub(mock *Mock, returns []interface{}) func(in []reflect.Value) []reflect.Value {
	funcType := mock.originalFuncValue.Type()

	if (returns != nil && len(returns) != funcType.NumOut()) {
		panic("SetReturns: Wrong number of returns arguments.")
	}

	fake := func(in []reflect.Value) []reflect.Value {
		inputArgs := []interface{}{}

		for _, value := range in {
			inputArgs = append(inputArgs, value.Interface())
		}

		mock.Calls = append(mock.Calls, inputArgs)

		var outputs []reflect.Value
		if (returns == nil) {
			for i := 0; i < funcType.NumOut(); i++ {
				lg("index", i)
				newOutput := reflect.New(funcType.Out(i)).Elem()
				outputs = append(outputs, newOutput)
			}
		} else {
			for i := 0; i < funcType.NumOut(); i++ {
				outputs = append(outputs, reflect.ValueOf(returns[i]))
			}
		}

		return outputs
	}

	return fake
}

func wrapFunc(mock *Mock, function interface{}) func(in []reflect.Value) []reflect.Value {
	fake := func(in []reflect.Value) []reflect.Value {
		inputArgs := []interface{}{}

		for _, value := range in {
			inputArgs = append(inputArgs, value.Interface())
		}

		mock.Calls = append(mock.Calls, inputArgs)
		return reflect.ValueOf(function).Call(in)
	}

	return fake
}
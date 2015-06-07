package xmock
import (
	"log"
	"reflect"
)

// If the function being mocked returns a func, that function will be nil
type Mock struct {
	OriginalFuncAddr reflect.Value
	OriginalFuncValue reflect.Value

	Calls [][]interface{}

	returns []interface{}
}

func (self *Mock) Restore() {
	lg("OriginalFuncAddr", self.OriginalFuncAddr)
	lg("OriginalFuncValue", self.OriginalFuncValue)

	self.OriginalFuncAddr.Elem().Set(self.OriginalFuncValue)
}

func (self *Mock) GetArgsForCall(n int) (interface{}) {
	args := self.Calls[n]
	return args
}

func (self *Mock) CallCount() int {
	return len(self.Calls)
}

func lg(data string, item interface{}) {
	log.Printf("*** %s %+V", data, item)
	log.Printf("*** %s %+v", data, item)
}

func MockFunc(funcToMock interface{}) *Mock {
	mock := Mock{}

	lg("funcToMock", funcToMock)

	funcValue := reflect.ValueOf(funcToMock)
	lg("funcValue", funcValue)

	mock.OriginalFuncAddr = funcValue

	funcType := funcValue.Type()
	lg("funcType", funcType)

	funcElem := funcValue.Elem()
	lg("funcElem", funcElem)

	mock.OriginalFuncValue = reflect.New(funcValue.Elem().Type()).Elem()
	mock.OriginalFuncValue.Set(funcElem)

	lg("mock.OriginalFuncAddr", mock.OriginalFuncAddr)
	lg("mock.OriginalFuncValue", mock.OriginalFuncValue)


	funcElemType := funcElem.Type()
	lg("funcElemType", funcElemType)

	fake := makeFuncStub(&mock, nil)


	mock.OriginalFuncAddr.Elem().Set(reflect.MakeFunc(mock.OriginalFuncValue.Type(), fake))
	return &mock
}

func (self *Mock) SetReturns(returns ...interface{}) {
	fake := makeFuncStub(self, returns)
	self.OriginalFuncAddr.Elem().Set(reflect.MakeFunc(self.OriginalFuncValue.Type(), fake))
}

func (self *Mock) SetReturnFunc(function interface{}) {
	wrap := wrapFunc(self, function)
	self.OriginalFuncAddr.Elem().Set(reflect.MakeFunc(self.OriginalFuncValue.Type(), wrap))
}

func makeFuncStub(mock *Mock, returns []interface{}) func(in []reflect.Value) []reflect.Value {
	funcType := mock.OriginalFuncValue.Type()

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
		lg("ASD", "ASD")
		return reflect.ValueOf(function).Call(in)
	}

	return fake
}
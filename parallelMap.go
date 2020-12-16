package main

import (
	"errors"
	"fmt"
	"reflect"
	"sync"
)
var (
	ErrMapSourceNotArray   = errors.New("input value is not an array")
	ErrMapTransformNil     = errors.New("transform function cannot be nil")
	ErrMapTransformNotFunc = errors.New("transform argument must be a function")
	ErrMapResultTypeNil    = errors.New("map result type cannot be nil")
)

func ParallelMap(source interface{}, transform interface{}, T reflect.Type) (interface{}, error) {
	srcV := reflect.ValueOf(source)
	kind := srcV.Kind()
	if kind != reflect.Slice && kind != reflect.Array {
		return nil, ErrMapSourceNotArray
	}
	if transform == nil {
		return nil, ErrMapTransformNil
	}
	tv := reflect.ValueOf(transform)
	if tv.Kind() != reflect.Func {
		return nil, ErrMapTransformNotFunc
	}
	if T == nil {
		return nil, ErrMapResultTypeNil
	}
	// kinda equivalent to = make([]T, srcv.Len())
	result := reflect.MakeSlice(reflect.SliceOf(T), srcV.Len(), srcV.Cap())
	// create a waitgroup with length = source array length
	// we'll reduce the counter each time an entry finished processing
	wg := &sync.WaitGroup{}
	wg.Add(srcV.Len())
	for i := 0; i < srcV.Len(); i++ {
		// one go routine for each entry
		go func(idx int, entry reflect.Value) {
			//Call the transformation and store the result value
			tfResults := tv.Call([]reflect.Value{entry})
			//Store the transformation result into array of result
			resultEntry := result.Index(idx)
			if len(tfResults) > 0 {
				resultEntry.Set(tfResults[0])
			} else {
				resultEntry.Set(reflect.Zero(T))
			}
			//this go routine is done
			wg.Done()
		}(i, srcV.Index(i))
	}
	wg.Wait()
	return result.Interface(), nil
}

func main(){
	i,err := ParallelMap([]int{1,2,3}, func(num int) int { return num+1 }, reflect.TypeOf(1))
	if err!= nil{
		fmt.Println(err)
	}
	fmt.Println(i)

}
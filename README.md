# commonfunctions
repo contains go lang implementations of common functions from other languages.

## sum(arr []interface{}) returns float64, error
The sum function returns the sum of all elements of an array if the array consists of integers or floating point numbers.
Usage:
```
  arr := int[]{1, 2, 3}
  total, err := sum(arr)
```
## ParallelMap(source interface{}, transform interface{}) (interface{}, error)
ParallelMap is a simple implementation of Maps inspired by blog post:-
```https://medium.com/better-programming/implementing-map-in-go-40c55d34dbb4```
The function takes three parameters:-
1. Array
2. Function to be implemented on array
3. Type of element

Usage:

```
i,err := ParallelMap([]int{1,2,3}, func(num int) int { return num+1 }, reflect.TypeOf(1))
	if err!= nil{
		fmt.Println(err)
	}
fmt.Println(i) //output - [2 3 4]


## enumerate(arr []interface{}) returns interface{}
The enumerate function returns an array with paired elements from array with their index
Usage:
```
fmt.Println(enumerate([]interface{"1.2","a"}))
//output - [{0 1.2} {1 a}]
```
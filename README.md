# commonfunctions
repo contains go lang implementations of common functions from other languages.

## sum(arr []interface{}) returns float64, error
The sum function returns the sum of all elements of an array if the array consists of integers or floating point numbers.
Usage:
```
  arr := int[]{1, 2, 3}
  total, err := sum(arr)
```

## enumerate(arr []interface{}) returns interface{}
The enumerate function returns an array with paired elements from array with their index
Usage:
```
fmt.Println(enumerate([]interface{"1.2","a"}))
//output - [{0 1.2} {1 a}]
```
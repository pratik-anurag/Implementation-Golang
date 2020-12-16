package enumerate

//struct to create a pair value
type returnTuple struct{
	index int
	value interface{}
}

func enumerate(list []interface{}) interface{} {
	//if length of list is empty returning nil
	if len(list)==0 {
		return nil
	}

	//creating array of struct
	var enumerateArray []returnTuple
	for key,val:= range list{
		returnT := returnTuple{index: key,value: val}
		enumerateArray = append(enumerateArray,returnT)
	}
	return enumerateArray
}
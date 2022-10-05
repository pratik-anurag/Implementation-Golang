package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func foo()(string,error){
	statusCode := rand.Int()
	if(statusCode%2 == 0){
		return "success",nil
	}else{
		err := errors.New("failure")
		return "failure",err
	}
}

func main(){
	ch := make(chan string)
	errch := make(chan error)
	for k:=0;k<10;k++{
		go func(){
			status,err:= foo()
			if err!=nil{
				errch<-err
			}else{
				ch <- status
			}
		}()

	}
	go func(){
		for{
				select {
				case c:= <-ch:
					fmt.Println(c)
				}
			}
		}()
	go func(){
		for{
			select {
			case e:= <-errch:
				fmt.Println(e)
			}
		}
	}()
	time.Sleep(100*time.Millisecond)

}

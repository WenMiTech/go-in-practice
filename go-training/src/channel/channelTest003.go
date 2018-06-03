package main
import "fmt"

func main(){
   data:=make(chan int,3)
   quit:=make(chan bool)

   go func(){
     for d := range data{  //block here where data id empty
       fmt.Println(d)
     }
     quit <- true     // send true to quit
     
   }()

    go func(){
     for d := range data{  //block here where data id empty
       fmt.Println(d)
     }
     quit <- true     // send true to quit

   }()
   data <- 1
   data <- 2
   close(data)  //must close the data chan or goroutine will always block
   
   <- quit // block here 
}

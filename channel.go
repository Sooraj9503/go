package main
import "fmt"

func myfunc(ch chan int){
	fmt.Println(234 + <-ch)
}

func myfunc1(mychnl chan string){
	for v := 0; v < 4; v ++{
		mychnl <- "hey!"
	}
	close(mychnl)
}

func main(){
	fmt.Println("Start main method")
	//Creating a channel
	ch := make(chan string)

	//Calling Go routine
	go myfunc1(ch)

	//When the value of ok is set to true means the channel is open and it can send
	//or receive data when the value of ok is set to false means the channel is closed

	for{
		res, ok := <-ch
		if!ok{
			fmt.Println("Channel Close", ok)
			break
		}
		fmt.Println("Channel Open", res, ok)
	}
	fmt.Println("End Main method")
}
package main

import ("fmt"; "time")


func printMessage(message string) {
	for i := 0; i < 5; i++ {
		fmt.Println(message)
		time.Sleep(500*time.Millisecond)
	}
}

func printMessageWithChannels(message string, channel chan string) {
	for i := 0; i < 5; i++ {
		fmt.Println(message)
		time.Sleep(500*time.Millisecond)
		channel <- fmt.Sprintf("%v times %v",message,i)
	}
}


func main(){
	 	// By simply adding the go keyword before the function we execute the function in another goroutine
		// The 1st printMessage function and the 2nd are executed at the same time, becuase the 1st one is executed in a seperate goroutine
		// and the 2nd one is running in the main goroutine
		// go printMessage("Hello! World")
		// printMessage("Hello! Balu")

		// Uncomment the code below to test and comment the rest of the out to test
		// In this case nothing happens becuase, both the printMessage functions are being executed in 2 seperate goroutines
		// And nothing is being run on the main goroutine
		// When the main goroutine is complete everything else (other goroutines) are terminated
		// go printMessage("Hello! World")
		// go printMessage("Hello! Balu")

		// This way we will always need to have something running in the main goroutine to make sure all the other goroutines can finish execution

		// Channels
		// It is possible to send data from a goroutine kinda like a return in a func using a channel
		// A channel is just like any normal variable just with the "chan" keyword before it
		channel := make(chan string)	
		go printMessageWithChannels("Hello! Channel",channel)
		printMessage("Hello! Balu")	
		response := <-channel
		fmt.Println(response)	
}
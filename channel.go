package main

func main() {
	data := make(chan int)
	done := make(chan bool)

	go consumer(data, done)
	go producer(data)

	<- done
}

func consumer(data chan int, done chan bool) {
	for x := range data {
		println("recv", x)
	}

	done <- true
}

func producer(data chan int) {
	for i := 0; i < 4; i++ {
		data <- i
	}
	close(data)
}
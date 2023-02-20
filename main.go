package main

import "simple-payment/delivery"

func main() {
		delivery.NewServer(":8080").Run()
}
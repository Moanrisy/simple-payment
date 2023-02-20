package main

import "simple-payment/delivery"

// @title Simple Payment API
func main() {
	delivery.NewServer(":8080").Run()
}
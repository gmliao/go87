package main

import "fmt"

// interface type bot defines a function getGreeting() that returns a string
type bot interface {
	getGreeting() string
}

// concrete types englishBot and spanishBot implement the interface bot
type englishBot struct{}

// concrete types englishBot and spanishBot implement the interface bot
// compiler will complain if concrete type does not implement all functions in the interface
func (englishBot) getGreeting() string {
	// very custom logic for generating an english greeting
	return "Hi there!"
}

type chineseBot struct{}

func (chineseBot) getGreeting() string {
	// very custom logic for generating a chinese greeting
	return "你好!"
}

// function printGreeting() takes a bot interface and prints the greeting
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func main() {
	eb := englishBot{}
	cb := chineseBot{}
	printGreeting(eb)
	printGreeting(cb)
}

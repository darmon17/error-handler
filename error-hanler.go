package error_handler

import "fmt"

// Hello mengembalikan sebuah salam untuk nama orang tertentu.
func Hello(name string) string {
	// Kembalikan sebuah salam yang berisi `name` dalam sebuah pesan.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
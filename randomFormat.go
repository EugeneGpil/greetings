package greetings

import "math/rand"

func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Здравствуйте, %v",
	}

	randomFormat := formats[rand.Intn(len(formats))]

	return randomFormat
}

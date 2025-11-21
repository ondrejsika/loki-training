package random_utils

import "math/rand"

func RandomChoice(choices []string) string {
	return choices[rand.Intn(len(choices))]
}

func GetRandomPath() string {
	return RandomChoice([]string{
		"/api/v1/users",
		"/api/v1/orders",
		"/api/v1/products",
		"/api/v1/cart",
		"/api/v1/checkout",
	})
}

func GetRandomMethod() string {
	return RandomChoice([]string{
		"GET",
		"POST",
	})
}

func GetRandomMessage() string {
	return RandomChoice([]string{
		"User logged in successfully",
		"Order created successfully",
		"Product added to cart",
		"Checkout completed",
		"User profile updated",
	})
}

func GetRandomLogLevel() string {
	return RandomChoice([]string{
		"debug",
		"info",
		"warning",
		"error",
		"fatal",
	})
}

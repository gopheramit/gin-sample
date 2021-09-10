package main

func main() {
	r := RegisterRoutes()

	r.Run(":3000")

}

package main

func main() {
	new_router := newRouter()
	new_router.Logger.Fatal(new_router.Start(":8080"))
}
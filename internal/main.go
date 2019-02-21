package main

import (
	server "search/hiendy"
)

func main() {

	server := server.Server{}
	defer server.Begin()

}

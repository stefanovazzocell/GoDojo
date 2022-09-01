package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./httpProxy [forward addr]")
		fmt.Println("Example: ./httpProxy http://localhost:8080")
		return
	}
	proxy := NewProxy(os.Args[1])
	http.ListenAndServe("0.0.0.0:8888", proxy)
}

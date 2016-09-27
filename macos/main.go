package main

import "fmt"
import "github.com/dontpanic92/wxGo/macos/app"
import "os"

func main() {
	if len(os.Args) == 1 {
		fmt.Println("macos <name>")
		return
	}
	name := os.Args[1]
	fmt.Println("building skeleton for", name)
	app.Skeleton(name)
}

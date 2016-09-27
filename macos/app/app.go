package app

import "fmt"
import "os"

func Skeleton(name string) {
	fmt.Println("creating dir")
	os.Mkdir(name, os.ModePerm)
	os.Mkdir(name+"/Contents", os.ModePerm)
	os.Mkdir(name+"/Contents/MacOS", os.ModePerm)
	os.Mkdir(name+"/Contents/Resources", os.ModePerm)

	file, _ := os.Create(name + "/Contents/Info.plist")
	fmt.Println(file)
}

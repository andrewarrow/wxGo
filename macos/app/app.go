package app

import "fmt"
import "os"
import "io/ioutil"
import "strings"

func Skeleton(name string) {
	fmt.Println("creating dir")
	os.Mkdir(name, os.ModePerm)
	os.Mkdir(name+"/Contents", os.ModePerm)
	os.Mkdir(name+"/Contents/MacOS", os.ModePerm)
	os.Mkdir(name+"/Contents/Resources", os.ModePerm)

	file, _ := os.Create(name + "/Contents/Info.plist")
	fmt.Println(file)

	template, _ := os.Open("Info.plist.template")
	data, _ := ioutil.ReadAll(template)
	text := string(data)
	for _, line := range strings.Split(text, "\n") {
		line = strings.Replace(line, "{name}", name, -1)
		line = strings.Replace(line, "{slug}", name, -1)
		file.Write([]byte(line + "\n"))
	}
}

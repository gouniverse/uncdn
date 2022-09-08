package templates

import (
	"io/ioutil"
	"log"
)

func ToString(path string) string {
	contentBytes, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println("Template: " + path + " NOT FOUND")
		return ""
	}
	content := string(contentBytes)
	// if t.Minified {
	// 	content = minify(path, content)
	// }
	return content
}

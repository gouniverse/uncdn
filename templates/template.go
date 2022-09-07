package templates

import (
	"embed"
	"log"
)

//go:embed *
var files embed.FS

func ToString(path string) string {
	contentBytes, err := files.ReadFile(path)
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

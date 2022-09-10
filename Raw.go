package uncdn

import "github.com/gouniverse/uncdn/templates"

func Raw(path string) string {
	return templates.ToString("vend/" + path)
}

package uncdn

import "github.com/gouniverse/uncdn/templates"

func NotifyJs() string {
	return templates.ToString("vend/notifyjs/notifyjs.min.js")
}

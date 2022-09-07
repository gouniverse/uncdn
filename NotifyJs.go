package ucdn

import "github.com/gouniverse/uncdn/templates"

func NotifyJs() string {
	return templates.ToString("vendor/notifyjs/notifyjs.min.js")
}

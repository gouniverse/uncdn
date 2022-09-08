package uncdn

import "github.com/gouniverse/uncdn/templates"

func NotifyJs() string {
	return templates.ToString("templates/vendor/notifyjs/notifyjs.min.js")
}

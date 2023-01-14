package uncdn

import "github.com/gouniverse/uncdn/templates"

func TrumbowygCss() string {
	return templates.ToString("vend/trumbowyg/2.26.0/trumbowyg.min.css")
}


func TrumbowygJs() string {
	return templates.ToString("vend/trumbowyg/2.26.0/trumbowyg.min.js")
}

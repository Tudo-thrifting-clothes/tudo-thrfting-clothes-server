package util

import "github.com/gosimple/slug"

func GenerateSlug(name string) string {
	// generate slug
	slug := slug.MakeLang(name, "en")
	return slug
}

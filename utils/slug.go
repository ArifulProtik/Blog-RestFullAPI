package utils

import "github.com/gosimple/slug"
// GenSlug Generate Unique Slug
func GenSlug(title string) string {
	slug := slug.Make(title)
	uid := UIDGen()
	return slug + uid
}
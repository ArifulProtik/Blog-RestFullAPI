package utils

import (
	"encoding/json"
	"forum-api/model"

	"github.com/liip/sheriff"
)

// Postmarshal marshals the post model
func Postmarshal(groups []string, posts []model.Post) ([]byte, error) {
	o := &sheriff.Options{
		Groups: groups,
	}
	data, err := sheriff.Marshal(o, posts)
	if err != nil {
		return nil, err
	}

	return json.MarshalIndent(data, "", "  ")
}

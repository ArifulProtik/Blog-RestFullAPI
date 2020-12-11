package utils

import (
	"strings"

	"github.com/google/uuid"
)

// UIDGen Generate Unique ID
func UIDGen() string {
	uid := uuid.New()
	uidstring := strings.Replace(uid.String(), "-", "", -1)
	return uidstring

}

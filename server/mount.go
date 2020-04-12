package server

import (
	"github.com/docker/docker/api/types/mount"
)

// Mount represents a Server Mount.
type Mount struct {
	ID       int        `json:"id"`
	Target   string     `json:"target"`
	Source   string     `json:"source"`
	ReadOnly bool       `json:"readonly"`
	Type     mount.Type `json:"type"`
}

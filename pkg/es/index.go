// Package es
// Before CRUD, make sure the auto_create_index configuration of the ES cluster is enabled.
package es

import (
	"beegoframe/pkg/es/internal"
)

var (
	Account internal.Index = "account"
)

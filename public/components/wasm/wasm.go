//go:build !openbsd

package wasm

import (
	// Bring in the internal plugin definitions.
	_ "github.com/benthosdev/benthos/v4/internal/impl/wasm"
)

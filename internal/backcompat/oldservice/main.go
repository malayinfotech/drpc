// Copyright (C) 2021 Storx Labs, Inc.
// See LICENSE for copying information.

// oldservice runs the old version of the backwards compatibility check.
package main

import (
	"context"

	"drpc/internal/backcompat"
)

func main() { backcompat.Main(context.Background()) }

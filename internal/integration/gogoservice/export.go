// Copyright (C) 2021 Storx Labs, Inc.
// See LICENSE for copying information.

// Package service is a testing service for the integration tests.
package service

import (
	"reflect"

	"drpc"
)

// Equal returns true if the two messages are equal.
func Equal(a, b drpc.Message) bool { return reflect.DeepEqual(a, b) }

// Encoding is the drpc.Encoding used for this service.
var Encoding drpcEncoding_File_service_proto

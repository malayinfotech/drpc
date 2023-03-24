// Copyright (C) 2022 Storx Labs, Inc.
// See LICENSE for copying information.

package main

import (
	"context"

	"drpc"
	"drpc/drpcmetadata"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

type streamWrapper struct {
	drpc.Stream
	ctx context.Context
}

func (s *streamWrapper) Context() context.Context { return s.ctx }

type otelHandler struct {
	handler drpc.Handler
}

func (t *otelHandler) HandleRPC(stream drpc.Stream, rpc string) (err error) {
	metadata, ok := drpcmetadata.Get(stream.Context())
	if ok {
		ctx := otel.GetTextMapPropagator().Extract(stream.Context(), propagation.MapCarrier(metadata))
		ctx, span := tracer.Start(ctx, "HandleRPC")
		defer span.End()
		stream = &streamWrapper{Stream: stream, ctx: ctx}
	}
	return t.handler.HandleRPC(stream, rpc)
}

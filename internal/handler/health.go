package handler

import "gofr.dev/pkg/gofr"

func Health(ctx *gofr.Context) (interface{}, error) {
	return map[string]string{
		"status": "ok",
	}, nil
}

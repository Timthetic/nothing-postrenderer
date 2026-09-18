package main

import (
	"log/slog"

	"github.com/Timthetic/nothing-postrenderer/internal/postrenderer"
)

func main() {
	if err := postrenderer.RunPostrenderer(); err != nil {
		slog.Error("Error running postrenderer")
	}
}

package internal

import (
	"github.com/disco07/gRPC/internal/calculator"
	"go.uber.org/fx"
)

var Module = fx.Options(
	calculator.Module,
)

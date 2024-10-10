package internal

import (
	"github.com/disco07/gRPC/internal/calculator"
	"github.com/disco07/gRPC/internal/users"
	"go.uber.org/fx"
)

var Module = fx.Options(
	calculator.Module,
	users.Module,
)

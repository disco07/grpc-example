package calculator

import (
	"go.uber.org/fx"
)

var Module = fx.Provide(
	NewCalculator,
)

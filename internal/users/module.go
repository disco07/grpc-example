package users

import "go.uber.org/fx"

var Module = fx.Provide(
	NewUsers,
)

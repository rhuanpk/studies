package init

import (
	// Packages to import 1th.
	_ "dev/pkg/router"

	// Packages to import 2th.
	_ "dev/internal/server/all"
	_ "dev/internal/server/check"
	_ "dev/internal/server/middles"
)

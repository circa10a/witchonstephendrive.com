package config

import (
	"io/fs"

	witchPrometheusMiddleware "github.com/circa10a/witchonstephendrive.com/routes/middleware/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// InitEchoConfig sets initial echo config such as middleware and logger
func (w *WitchConfig) InitEchoConfig(frontendAssets fs.FS, apiDocAssets fs.FS) *echo.Echo {
	// New instance of echo
	e := echo.New()

	// Prometheus metrics
	if w.Metrics {
		prometheus := witchPrometheusMiddleware.NewPrometheus(w.APIBaseURL)
		prometheus.Use(e)
	}

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "\033[36mINFO\033[0m[${time_rfc3339}] ${status} ${method} ${path} (${remote_ip}) ${latency_human}\n",
		Output: e.Logger.Output(),
	}))

	return e
}

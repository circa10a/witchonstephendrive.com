package config

import (
	"io/fs"

	witchGeofencingMiddleware "github.com/circa10a/witchonstephendrive.com/internal/routes/middleware/geofencing"
	witchPrometheusMiddleware "github.com/circa10a/witchonstephendrive.com/internal/routes/middleware/prometheus"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// InitEchoConfig sets initial echo config such as middleware and logger
func (w *WitchConfig) InitEchoConfig(frontendAssets fs.FS, apiDocAssets fs.FS) *echo.Echo {
	// New instance of echo
	e := echo.New()
	// Disable massive startup banner
	e.HideBanner = true

	// Prometheus metrics
	if w.MetricsEnabled {
		prometheus := witchPrometheusMiddleware.NewPrometheus(w.APIBaseURL)
		prometheus.Use(e)
	}

	// Geofencing only POST (state changing) routes
	if w.GeofencingEnabled {
		e.Use(witchGeofencingMiddleware.IsClientAllowed(w.GeofencingClient))
	}

	// Make route logging easier to read/match logrus without the shitty middleware
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "\033[36mINFO\033[0m[${time_rfc3339}] ${status} ${method} ${path} (${remote_ip}) ${latency_human}\n",
		Output: e.Logger.Output(),
	}))

	// Get real IP from proxy (caddy) to properly use geofencing
	e.IPExtractor = echo.ExtractIPFromXFFHeader()
	return e
}

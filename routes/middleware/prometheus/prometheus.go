package prometheus

import (
	"strings"

	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
)

// NewPrometheus creates a new prometheus instance
// But ensure route names are fully populated for accuracy and more cardinality
func NewPrometheus(apiBaseURL string) *prometheus.Prometheus {
	prom := prometheus.NewPrometheus("echo", nil)
	prom.RequestCounterURLLabelMappingFunc = func(c echo.Context) string {
		url := c.Request().URL.Path
		color := c.Param("color")
		sound := c.Param("sound")
		// Strip API Base URL path(/api/1)
		url = strings.Replace(url, apiBaseURL, "", 1)
		// /color/:color => /color/red
		if color != "" {
			url = strings.Replace(url, ":color", color, 1)
		}
		// /sound/:sound => /sound/this-is-halloween
		if sound != "" {
			url = strings.Replace(url, ":sound", sound, 1)
		}
		return url
	}
	return prom
}

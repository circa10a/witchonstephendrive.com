package middleware

import (
	"strings"

	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
)

// NewPrometheusMiddleware creates a new prometheus instance
// But ensure route names are fully populated for accuracy and more cardinality
func NewPrometheusMiddlware() *prometheus.Prometheus {
	prom := prometheus.NewPrometheus("echo", nil)
	prom.RequestCounterURLLabelMappingFunc = func(c echo.Context) string {
		url := c.Request().URL.Path
		color := c.Param("color")
		sound := c.Param("sound")
		if color != "" {
			url = strings.Replace(url, ":color", color, 1)
		}
		if sound != "" {
			url = strings.Replace(url, ":sound", sound, 1)
		}
		return url
	}
	return prom
}

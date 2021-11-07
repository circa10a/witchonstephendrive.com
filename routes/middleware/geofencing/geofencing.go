package geofencing

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/qioalice/ipstack"
	log "github.com/sirupsen/logrus"
)

// GeofencingAllowed responds with a boolean to indicate whether a client is allowed to make changes
// based on proximity to the server
type GeofencingAllowed struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// IsClientAllowed looks up the coordinates of a client to see if it's nearby
func IsClientAllowed(latitude, longitude float32, clientIPCache map[string]bool) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			method := c.Request().Method
			// We only care about routes that effect state
			if method != http.MethodPost {
				return next(c)
			}

			// Client ip, strip port
			ipAddress := strings.Split(c.Request().Host, ":")[0]
			// Ensure not a private ip
			if strings.HasPrefix(ipAddress, "192") ||
				strings.HasPrefix(ipAddress, "172") ||
				strings.HasPrefix(ipAddress, "10") {
				log.Debugf("Private ip of %s detected. Skipping geofencing middleware", ipAddress)
				return next(c)
			}

			// Remove the last decimal to make sure we're not too exact
			currentLat := fmt.Sprintf("%.4f", latitude)
			currentLong := fmt.Sprintf("%.4f", longitude)

			// If address in cache and allowed
			if val, ok := clientIPCache[ipAddress]; ok {
				if val {
					return next(c)
				}
			}

			// Look up client location
			clientCoordinates, err := ipstack.IP(ipAddress)
			if err != nil {
				log.Error(err)
				return next(c)
			}
			clientLat := fmt.Sprintf("%.4f", clientCoordinates.Latitide)
			clientLong := fmt.Sprintf("%.4f", clientCoordinates.Longitude)

			// Is client close enough and allowed to continue?
			clientIPCache[ipAddress] = currentLat == clientLat && currentLong == clientLong
			if !clientIPCache[ipAddress] {
				return c.JSON(http.StatusForbidden, GeofencingAllowed{
					Success: false,
					Message: "Client not within proximity to the server. Not allowing changes",
				})
			}
			return next(c)
		}
	}
}

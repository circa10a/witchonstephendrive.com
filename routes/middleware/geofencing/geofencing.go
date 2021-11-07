package geofencing

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/qioalice/ipstack"
	log "github.com/sirupsen/logrus"
)

// GeofencingAllowedResponse responds with a boolean to indicate whether a client is allowed to make changes
// based on proximity to the server
type GeofencingAllowedResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// widenCoordinatesToString converts 5 to 3 decimal points and givens back a string for comparison
func widenCoordinatesToString(location float32) string {
	return fmt.Sprintf("%.3f", location)
}

// IsClientAllowed looks up the coordinates of a client to see if it's nearby
func IsClientAllowed(latitude, longitude float32, clientIPCache map[string]bool) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			method := c.Request().Method
			// We only care about routes that affect state
			if method != http.MethodPost {
				return next(c)
			}

			// Client ip, strip port
			ipAddress := strings.Split(c.Request().Host, ":")[0]
			// Ensure not a private ip
			if strings.HasPrefix(ipAddress, "192.") ||
				strings.HasPrefix(ipAddress, "172.") ||
				strings.HasPrefix(ipAddress, "10.") {
				log.Debugf("Private ip of %s detected. Skipping geofencing middleware", ipAddress)
				return next(c)
			}

			// Remove the last decimal to make sure we're not too exact
			currentLat := widenCoordinatesToString(latitude)
			currentLong := widenCoordinatesToString(longitude)

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

			clientLat := widenCoordinatesToString(clientCoordinates.Latitide)
			clientLong := widenCoordinatesToString(clientCoordinates.Longitude)

			// Is client close enough and allowed to continue?
			clientIPCache[ipAddress] = currentLat == clientLat && currentLong == clientLong
			// If not close enough, reject request
			if !clientIPCache[ipAddress] {
				return c.JSON(http.StatusForbidden, GeofencingAllowedResponse{
					Success: false,
					Message: "Client not within proximity to the server. Not allowing changes",
				})
			}
			// If close enough, allow to proceed
			return next(c)
		}
	}
}

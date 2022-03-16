package geofencing

import (
	"net/http"

	"github.com/circa10a/go-geofence"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

// GeofencingAllowedResponse responds with a boolean to indicate whether a client is allowed to make changes
// based on proximity to the server
type GeofencingAllowedResponse struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

// IsClientAllowed looks up the coordinates of a client to see if it's nearby
func IsClientAllowed(geofenceClient geofence.Geofence) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// We only care about routes that affect state
			if c.Request().Method != http.MethodPost {
				return next(c)
			}

			// Client ip, strip port
			ipAddress := c.RealIP()

			isAllowed, err := geofenceClient.IsIPAddressNear(ipAddress)
			if err != nil {
				log.Error(err)
			}

			log.Debugf("IP address: %s, isAllowed: %t", ipAddress, isAllowed)

			if !isAllowed {
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

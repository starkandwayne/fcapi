package memory

import (
	"github.com/labstack/echo/v4"

	"github.com/starkandwayne/fcapi/core"
)

func NewServices() *core.Services {
	return &core.Services{
		Brokers: NewBrokerService(),
		Orgs:    NewOrganizationService(),
		Spaces:  NewSpaceService(),
		Logger:  func(e echo.Context) {},
	}
}

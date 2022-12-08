package initialize

import (
	"wellpict-be/app/controller"
)

var InitFunctions = []InitFunc{
	controller.InitializeUserController,
	controller.InitializeAdminController,
}

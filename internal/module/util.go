package module

import (
	utilDomain "github.com/rudianto-dev/gotemp-api-service/internal/domain/util"
	utilHandler "github.com/rudianto-dev/gotemp-api-service/internal/handler/api/util"
)

func (m *Module) UtilHandlerAPI() utilDomain.IHandlerAPI {
	return utilHandler.NewHandler(m.infra.Logger)
}

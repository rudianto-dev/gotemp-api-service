package module

import (
	userDomain "github.com/rudianto-dev/gotemp-api-service/internal/domain/user"
	userHandler "github.com/rudianto-dev/gotemp-api-service/internal/handler/api/user"
	userUseCase "github.com/rudianto-dev/gotemp-api-service/internal/usecase/user"
)

func (m *Module) UserHandlerAPI() userDomain.IHandlerAPI {
	userUseCase := userUseCase.NewUseCase(m.infra.Logger, m.userRepo)
	return userHandler.NewHandler(m.infra.Logger, userUseCase)
}

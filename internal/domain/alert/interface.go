package alert

import (
	"context"

	alertRepository "github.com/rudianto-dev/gotemp-api-service/internal/domain/alert/repository"
)

type Repository interface {
	Send(ctx context.Context, alertEntity *alertRepository.AlertEntity) error
}

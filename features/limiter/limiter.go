package limiter

import (
	"context"

	"github.com/xtls/xray-core/common/errors"
	"github.com/xtls/xray-core/features"
)

// Manager is the interface for limiter manager.
type Manager interface {
	features.Feature

	// DeviceLimitExceeded check client limits by its identifier.
	DeviceLimitExceeded(ctx context.Context) bool
}

func CheckLimits(ctx context.Context, m Manager) error {
	if m.DeviceLimitExceeded(ctx) {
		return errors.New("device limit exceeded").AtInfo()
	}
	return nil
}

// ManagerType returns the type of Manager interface. Can be used to implement common.HasType.
func ManagerType() interface{} {
	return (*Manager)(nil)
}

type NoopManager struct{}

func (n NoopManager) Type() interface{} {
	return ManagerType()
}

func (n NoopManager) Start() error {
	return nil
}

func (n NoopManager) Close() error {
	return nil
}

func (n NoopManager) DeviceLimitExceeded(context.Context) bool {
	return false
}

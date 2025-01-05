package limiter

import (
	"context"

	"github.com/xtls/xray-core/common"
	"github.com/xtls/xray-core/common/errors"
	"github.com/xtls/xray-core/common/protocol"
	"github.com/xtls/xray-core/common/session"
	"github.com/xtls/xray-core/core"
	"github.com/xtls/xray-core/features/limiter"
	"github.com/xtls/xray-core/features/policy"
	"github.com/xtls/xray-core/features/stats"
)

// Manager is an implementation of feature limiter.Manager.
type Manager struct {
	policyManager policy.Manager
	statsManager  stats.Manager
}

// Type implements common.HasType.
func (*Manager) Type() interface{} {
	return limiter.ManagerType()
}

// Start implements common.Runnable.
func (*Manager) Start() error {
	return nil
}

// Close implements common.Runnable.
func (*Manager) Close() error {
	return nil
}

// DeviceLimitExceeded implements feature limiter.Manager.
func (m *Manager) DeviceLimitExceeded(ctx context.Context) bool {
	sessionInbound := session.InboundFromContext(ctx)
	var user *protocol.MemoryUser
	if sessionInbound != nil {
		user = sessionInbound.User
	}

	if user != nil && user.Email != "" {
		p := m.policyManager.ForLevel(user.Level)
		if p.DeviceCount <= 0 {
			return false
		}

		if !p.Stats.UserOnline {
			errors.LogWarning(ctx, "device limitation required user online statistics in policy level")
			return false
		}

		name := "user>>>" + user.Email + ">>>online"
		onlineMap := m.statsManager.GetOnlineMap(name)
		if onlineMap != nil && onlineMap.Count() > int(p.DeviceCount) {
			return true
		}
	}

	return false
}

func init() {
	common.Must(common.RegisterConfig((*Config)(nil), func(ctx context.Context, cfg interface{}) (interface{}, error) {
		m := new(Manager)
		if err := core.RequireFeatures(ctx, func(pm policy.Manager, sm stats.Manager) error {
			m.policyManager = pm
			m.statsManager = sm
			return nil
		}); err != nil {
			return nil, err
		}
		return m, nil
	}))
}

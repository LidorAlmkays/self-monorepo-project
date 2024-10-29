package left

import (
	"github.com/LidorAlmkays/self-monorepo-project/apps/user/internal/application"
)

type BaseServer interface {
	ListenAndServe(userApi application.UserPort) error
}

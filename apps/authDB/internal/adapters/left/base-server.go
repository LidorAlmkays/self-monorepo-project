package left

import (
	"github.com/LidorAlmkays/self-monorepo-project/apps/authDB/internal/application"
)

type BaseServer interface {
	ListenAndServe(userApi application.UserPort) error
}

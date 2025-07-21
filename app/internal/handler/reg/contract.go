package reg

import "market/app/internal/entity"

type Registry interface {
	Registration(user entity.User) (entity.User, error)
}

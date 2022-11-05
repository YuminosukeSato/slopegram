package usecase

import (
	"github.com/YuminosukeSato/slopegram/background/domain"
)

type UserRepository interface {
    Store(domain.User)
    Select() []domain.User
    Delete(id string)
}

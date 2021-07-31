package interactor

import (
	"context"

	"github.com/Goalt/FileSharer/internal/domain"
)

type FileInteractor interface {
	Upload(ctx context.Context, file *domain.File) (domain.Token, error)
	Download(ctx context.Context, token domain.Token) (*domain.File, error)
}

type fileInteractor struct {
}

func NewFileInteractor() FileInteractor {
	return &fileInteractor{}
}

func (ci *fileInteractor) Upload(ctx context.Context, file *domain.File) (domain.Token, error) {
	return domain.Token(domain.Token{}), nil
}

func (ci *fileInteractor) Download(ctx context.Context, token domain.Token) (*domain.File, error) {
	return nil, nil
}

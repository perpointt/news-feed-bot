package fetcher

import (
	"context"
	"news-feed-bot/internal/model"
	"time"
)

type ArticleStorage interface {
	Store(ctx context.Context, article model.Article) error
}

type SourceProvider interface {
	Sources(ctx context.Context) ([]model.Source, error)
}

type Source interface {
	ID() int64
	Name() string
	Fetch(ctx context.Context) ([]model.Item, error)
}

type Fetcher struct {
	articles ArticleStorage
	sources  SourceProvider

	fetchInterval  time.Duration
	filterKeywords []string
}

func New(
	articleStorage ArticleStorage,
	sourceProvider SourceProvider,
	fetchInterval time.Duration,
	filterKeywords []string) *Fetcher {
	return &Fetcher{
		articles:       articleStorage,
		sources:        sourceProvider,
		fetchInterval:  fetchInterval,
		filterKeywords: filterKeywords,
	}
}

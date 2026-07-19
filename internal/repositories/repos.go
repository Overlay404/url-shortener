package repositories

import (
	"context"
	"fmt"
	"log/slog"
	"url-shortener/internal/config"
	"url-shortener/internal/models"

	"github.com/redis/go-redis/v9"
)

var Client *redis.Client

type Repos struct {
	ctx  *context.Context
	conf *config.Config
	log  *slog.Logger
}

func New(ctx *context.Context, conf *config.Config, logger *slog.Logger) (r *Repos) {
	return &Repos{
		ctx:  ctx,
		conf: conf,
		log:  logger,
	}
}

func (r *Repos) Init() {
	Client = redis.NewClient(&redis.Options{
		Addr:     r.conf.Redis.Addr,
		Password: r.conf.Redis.Password,
		DB:       r.conf.Redis.DB,
	})
}

func (r *Repos) Set(urlShort string, link *models.Link) error {
	err := Client.Set(*r.ctx, urlShort, &link, 0).Err()
	if err != nil {
		return fmt.Errorf("cannot set value in Redis %w", err)
	}
	r.log.Debug("write in redis value", slog.Any("link", &link))
	return nil
}

func (r *Repos) Get(urlShort string) (link *models.Link, err error) {
	err = Client.Get(*r.ctx, urlShort).Scan(&link)
	if err != nil {
		return nil, fmt.Errorf("cannot set value in Redis %w", err)
	}
	r.log.Debug("get value from redis", slog.Any("link", &link))
	return link, nil
}

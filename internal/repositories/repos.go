package repositories

import (
	"context"
	"fmt"
	"log/slog"
	"url-shortener/internal/config"
	"url-shortener/internal/models"
	"url-shortener/pkg"

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

func (r *Repos) Set(link *models.Link) error {
	urlShort := pkg.GenShortUrl(6)
	err := Client.Set(*r.ctx, urlShort, link, 0).Err()
	if err != nil {
		return fmt.Errorf("cannot set value in Redis %w", err)
	}
	r.log.Debug("write in redis value", slog.Any("link", &link))
	pkg.WriteInFile(urlShort, r.log)
	return nil
}

//не закончил: получение по ссылке значение, преобразование link из byte[], возрат *link
func (r *Repos) Get() error {
	err := Client.Set(*r.ctx, "key", "value", 0).Err()
	if err != nil {
		return fmt.Errorf("cannot set value in Redis %w", err)
	}
	return nil
}

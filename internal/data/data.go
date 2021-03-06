package data

import (
	"context"
	"kratos/internal/conf"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/google/wire"

	"github.com/go-redis/redis/extra/redisotel"
	"github.com/go-redis/redis/v8"

	"kratos/internal/data/ent"

	// init mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo, NewArticleRepo)

// Data .
type Data struct {
	// TODO warpped database client
	db  *ent.Client
	rdb *redis.Client
}

// NewData .
func NewData(conf *conf.Data, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper("data", logger)
	client, err := ent.Open(
		conf.Database.Driver,
		conf.Database.Source,
	)
	if err != nil {
		log.Errorf("failed opening connection to sqlite: %v", err)
		return nil, nil, err
	}
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Errorf("failed creating schema resources: %v", err)
		return nil, nil, err
	}

	rdb := redis.NewClient(&redis.Options{
		Addr: conf.Redis.Addr,
		// Password:     conf.Redis.Password,
		// DB:           int(conf.Redis.Db),
		// DialTimeout:  conf.Redis.DialTimeout.AsDuration(),
		WriteTimeout: conf.Redis.WriteTimeout.AsDuration(),
		ReadTimeout:  conf.Redis.ReadTimeout.AsDuration(),
	})
	rdb.AddHook(redisotel.TracingHook{})
	d := &Data{
		db:  client,
		rdb: rdb,
	}
	return d, func() {
		log.Info("message", "closing the data resources")
		if err := d.db.Close(); err != nil {
			log.Error(err)
		}
		if err := d.rdb.Close(); err != nil {
			log.Error(err)
		}
	}, nil

	// cleanup := func() {
	// 	logger.Log("msg", "closing the data resources")
	// }
	// return &Data{}, cleanup, nil
}

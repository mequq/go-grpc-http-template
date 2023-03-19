package data

import (
	"context"
	"strconv"

	"entgo.io/ent/dialect"
	_ "github.com/mattn/go-sqlite3"
	"github.com/mequq/go-grpc-http-template/config"
	"github.com/mequq/go-grpc-http-template/internal/ent"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

// DataSouce is the struct that holds all the data sources
type DataSource struct {
	logger *zap.Logger
	cfg    *config.ViperConfig
	redis  *redis.Client
}

// NewDataSource creates a new DataSource
func NewDataSource(logger *zap.Logger, cfg *config.ViperConfig) (*DataSource, error) {
	ds := &DataSource{
		logger: logger,
		cfg:    cfg,
		redis:  nil,
	}
	err := ds.Init()
	if err != nil {
		return nil, err
	}
	return ds, nil
}

func (ds *DataSource) Init() error {
	if ds.cfg.DatasourceConfig.Redis.Enabled {
		err := ds.initRedis()
		if err != nil {
			return err
		}
	}
	return nil
}

func (ds *DataSource) Close() error {
	if ds.cfg.DatasourceConfig.Redis.Enabled {
		err := ds.redis.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

func (ds *DataSource) initRedis() error {
	ds.redis = redis.NewClient(&redis.Options{
		Addr:     ds.cfg.DatasourceConfig.Redis.Host + ":" + strconv.Itoa(ds.cfg.DatasourceConfig.Redis.Port),
		Password: ds.cfg.DatasourceConfig.Redis.Password,
		DB:       ds.cfg.DatasourceConfig.Redis.Database,
	})
	_, err := ds.redis.Ping(context.Background()).Result()
	if err != nil {
		ds.logger.Error("redis ping failed", zap.Error(err))
		return err
	}
	return nil
}

// initSqlite initializes the sqlite database
func (ds *DataSource) initSqlite() error {

	client, err := ent.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		ds.logger.Error("failed opening connection to sqlite", zap.Error(err))
		return err
	}
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		return err
	}
	return nil

}

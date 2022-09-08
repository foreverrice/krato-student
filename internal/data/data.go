package data

import (
	"finance/internal/conf"
	"finance/internal/data/ent"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"

	// init mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewBalanceRepo, NewCartClient)

// Data .
type Data struct {
	finance *ent.Client
	// TODO wrapped database client
}

// NewData .
func NewData(c *conf.Data, finance *ent.Client, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{
		finance: finance,
	}, cleanup, nil
}

func NewCartClient(conf *conf.Data, logger log.Logger) *ent.Client {
	log := log.NewHelper(log.With(logger, "module", "user-service/data/ent"))
	client, err := ent.Open(
		conf.Database.Driver,
		conf.Database.Source,
	)
	if err != nil {
		log.Fatalf("failed opening connection to db: %v", err)
	}
	client = client.Debug()

	// Run the auto migration tool.
	//if err := client.Schema.Create(context.Background(), migrate.WithForeignKeys(false)); err != nil {
	//	log.Fatalf("failed creating schema resources: %v", err)
	//}
	return client
}

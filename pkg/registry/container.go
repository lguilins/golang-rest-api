package registry

import (
	"log"

	"github.com/sarulabs/di"
)

// Services contains the definitions of the application services.
var Services = []di.Def{
	{
		Name:  "mysqldb-conn",
		Build: buildMysqlDBConn,
	},
	{
		Name:  "user-usecase",
		Build: buildUserUsecase,
	},
	{
		Name:  "user-repository",
		Build: buildUserMysqlRepository,
	},
}

// NewContainer create a new instance of di.Container
func NewContainer() di.Container {
	builder, err := di.NewBuilder()
	if err != nil {
		log.Fatalf("Failed to build the container %v", err)
	}

	err = builder.Add(Services...)
	if err != nil {
		log.Fatalf("Failed to registry services %v", err)
	}

	app := builder.Build()

	return app
}

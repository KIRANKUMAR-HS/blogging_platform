package main

import (
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"

	// "github.com/spf13/viper"

	"github.com/KIRANKUMAR-HS/blogging_platform/internal/apihandler"
	"github.com/KIRANKUMAR-HS/blogging_platform/internal/authservice"
	"github.com/KIRANKUMAR-HS/blogging_platform/internal/config"
	"github.com/KIRANKUMAR-HS/blogging_platform/internal/logger"
	db "github.com/KIRANKUMAR-HS/blogging_platform/internal/psql"
	"github.com/KIRANKUMAR-HS/blogging_platform/internal/router"
)

const (
	secretKey = "my-32-character-ultra-secure-and-ultra-long-secret"
)

func init() {
	// init config and logger before executing main
	config.Init()
	logger.Init()
}

func main() {

	Psqlconn := fmt.Sprintf(
		// "host=localhost port=5432 user=kiran password=Pass@1234 dbname=blogging sslmode=disable",

		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		viper.GetString("database.psql.host"),
		viper.GetInt("database.psql.port"),
		viper.GetString("database.psql.user"),
		viper.GetString("database.psql.password"),
		viper.GetString("database.psql.dbname"),
	)
	fmt.Println(Psqlconn)

	PsqlClint, err := db.NewPsqlClint(Psqlconn)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create psql clint")
	}

	auth, err := authservice.NewAuthService(PsqlClint, secretKey)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create authservice")
	}

	Handler, err := apihandler.NewBlogServer(PsqlClint, auth)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create causes grpc server")
	}
	r, err := router.NewRouter(Handler, auth)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create router")

	}

	// Start the server on port 8080
	fmt.Println("Starting server on :8080")
	if err := Handler.Start(viper.GetString("api_server.url_adress"), r); err != nil {
		log.Fatal().Err(err).Msg("failed to create causes grpc server")
	}

}

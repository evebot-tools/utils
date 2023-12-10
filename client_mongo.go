package utils

import (
	"time"

	"github.com/kamva/mgm/v3"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func InitMongoDBClient() {
	mongoOpts := options.Client()
	mongoOpts.SetReadPreference(readpref.SecondaryPreferred())
	mongoOpts.ApplyURI(GetEnv("MONGODB_URI", ""))
	err := mgm.SetDefaultConfig(&mgm.Config{CtxTimeout: 60 * time.Second}, GetEnv("MONGODB_DATABASE", "evetools"), mongoOpts)

	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect to mongodb")
	}
}

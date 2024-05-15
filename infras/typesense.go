package infras

import (
	"context"
	"errors"
	"searchproject/configs"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/typesense/typesense-go/typesense"
	"github.com/typesense/typesense-go/typesense/api"
)

type (
	TypesenseClient   *typesense.Client
	MultipleTypesense map[string]*typesense.Client
)

func MustTSNewClient(cfg *configs.Config) TypesenseClient {
	client, err := tsNewClient(cfg.TS.ApiKey, cfg.TS.Nodes)
	if err != nil {
		log.Fatal().Msg("unable to create typesense client")
	}
	return TypesenseClient(client)
}

func MultipleTSNewClient(cfg *configs.Config) MultipleTypesense {
	multipleTS := make(MultipleTypesense)
	for instanceName, tsCfg := range cfg.MultipleTS {
		var err error
		instanceName = strings.ToLower(instanceName)
		multipleTS[instanceName], err = tsNewClient(tsCfg.ApiKey, tsCfg.Nodes)
		if err != nil {
			log.Warn().Msgf("failed to initiate '%s' typesense instance", instanceName)
		}
	}
	return multipleTS
}

func tsNewClient(apiKey string, nodes []api.Node) (*typesense.Client, error) {
	if len(nodes) == 0 {
		log.Fatal().Msg("empty typesense nodes")
	}

	client := typesense.NewClient(
		typesense.WithServer(nodes...),
		typesense.WithAPIKey(apiKey),
	)

	if err := checkTSHealth(client); err != nil {
		log.Info().Msgf("unable to check typesense health: %v", err)
		return nil, err
	}

	log.Info().
		Str("node", nodes[0].Host).
		Msg("Connected to Typesense")

	return client, nil
}

func checkTSHealth(client *typesense.Client) error {
	health, err := client.Health().Retrieve(context.Background())
	if err != nil {
		return err
	}
	if health["ok"] != "true" {
		return errors.New("typesense health check failed")
	}
	return nil
}

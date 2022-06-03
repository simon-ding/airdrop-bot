package aws

import (
	"airdrop-bot/log"
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/pkg/errors"
)

func Load(dir string) error {

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithSharedConfigFiles(
			[]string{dir},
		),
		config.WithRegion(string(types.RegionNameApSoutheast1)),
	)
	if err != nil {
		return errors.Wrap(err, "load aws config")
	}
	client := lightsail.NewFromConfig(cfg)

	outputs, err := client.GetStaticIps(context.TODO(), &lightsail.GetStaticIpsInput{})

	log.Info(outputs, err)
	return nil
}

type Client struct {
	lightsail *lightsail.Client
}

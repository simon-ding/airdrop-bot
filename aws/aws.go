package aws

import (
	"airdrop-bot/db"
	"airdrop-bot/log"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/pkg/errors"
)

func CreateLightsailClient(instanceName, region, cfgDir string) (*Client, error) {

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithSharedConfigFiles(
			[]string{cfgDir},
		),
		//types.RegionNameUsWest2
		config.WithRegion(region),
	)
	if err != nil {
		return nil, errors.Wrap(err, "load aws config")
	}
	client := lightsail.NewFromConfig(cfg)
	c := &Client{Client: client, instanceName: instanceName}
	return c, nil
}

type Client struct {
	*lightsail.Client
	instanceName string
}

func (c *Client) CreateIp(name string) error {
	in := lightsail.AllocateStaticIpInput{
		StaticIpName: &name,
	}
	out, err := c.AllocateStaticIp(context.TODO(), &in)
	if err != nil {
		return errors.Wrap(err, "new ip")
	}

	log.Infof("new ip created: %+v", out)
	return nil
}

func (c *Client) DetachIp(name string) error {
	in := lightsail.DetachStaticIpInput{StaticIpName: &name}

	out, err := c.Client.DetachStaticIp(context.TODO(), &in)
	if err != nil {
		return errors.Wrap(err, "detach ip")
	}
	log.Infof("ip %s detach success: %+v", name, out)
	return nil
}

func (c *Client) AttachIp(name string) error {
	in := lightsail.AttachStaticIpInput{
		InstanceName: &c.instanceName,
		StaticIpName: &name,
	}

	out, err := c.Client.AttachStaticIp(context.TODO(), &in)
	if err != nil {
		return errors.Wrap(err, "attach ip")
	}
	log.Infof("attach ip %s to instance %s success: %+v", name, c.instanceName, out)
	return nil
}

func (c *Client) GetAttachedIp() (string, string, error) {
	in := lightsail.GetStaticIpsInput{}
	out, err := c.Client.GetStaticIps(context.TODO(), &in)
	if err != nil {
		return "", "", errors.Wrap(err, "get static ips")
	}
	for _, ip := range out.StaticIps {
		if ip.IsAttached != nil && *ip.IsAttached {
			return *ip.Name, *ip.IpAddress, err
		}
	}
	return "", "", fmt.Errorf("no attached ip")
}

func (c *Client) GetAllIps() {
	in := lightsail.GetStaticIpsInput{}
	out, err := c.Client.GetStaticIps(context.TODO(), &in)
	if err != nil {
		return
	}
	for _, ip := range out.StaticIps {
		db.SaveOrUpdateIp(*ip.Name, *ip.IpAddress)
	}

}

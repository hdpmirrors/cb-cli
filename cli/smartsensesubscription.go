package cli

import (
	"errors"
	"strconv"
	"time"

	log "github.com/Sirupsen/logrus"
	s "github.com/hortonworks/hdc-cli/client_cloudbreak/smartsensesubscriptions"
	"github.com/hortonworks/hdc-cli/models_cloudbreak"
	"github.com/urfave/cli"
)

var SmartSenseSubscriptionHeader []string = []string{"ID", "SubscriptionID"}

type SmartSenseSubscription struct {
	Id             string `json:"ID" yaml:"ID"`
	SubscriptionId string `json:"SubscriptionID" yaml:"SubscriptionID"`
}

func (b *SmartSenseSubscription) DataAsStringArray() []string {
	return []string{b.Id, b.SubscriptionId}
}

func CreateSmartSenseSubscription(c *cli.Context) error {
	defer timeTrack(time.Now(), "creation of SmartSenseSubscription")
	checkRequiredFlags(c)

	subscriptionId := c.String(FlSmartSenseSubscriptionID.Name)
	if len(subscriptionId) == 0 {
		logMissingParameterAndExit(c, []string{FlSmartSenseSubscriptionID.Name})
	}

	cbClient := NewCloudbreakOAuth2HTTPClient(c.String(FlServer.Name), c.String(FlUsername.Name), c.String(FlPassword.Name))
	createSmartSenseSubscriptionImpl(subscriptionId,
		cbClient.Cloudbreak.Smartsensesubscriptions.PostPrivateSmartSenseSubscription)
	return nil
}

func createSmartSenseSubscriptionImpl(subscriptionId string,
	postPrivateSmartSenseSubscription func(*s.PostPrivateSmartSenseSubscriptionParams) (*s.PostPrivateSmartSenseSubscriptionOK, error)) {

	log.Infof("[createSmartSenseSubscription] create SmartSense subscription for: %s ", subscriptionId)
	resp, err := postPrivateSmartSenseSubscription(
		&s.PostPrivateSmartSenseSubscriptionParams{Body: &models_cloudbreak.SmartSenseSubscriptionJSON{SubscriptionID: &subscriptionId}})

	if err != nil {
		logErrorAndExit(err)
	}

	log.Infof("[createSmartSenseSubscription] SmartSense subscription created: %v", *resp.Payload)
}

func DeleteSmartSenseSubscription(c *cli.Context) error {
	defer timeTrack(time.Now(), "deletion of SmartSenseSubscription")
	checkRequiredFlags(c)

	id := c.Int64(FlSmartSenseSubscriptionID.Name)
	if id == 0 {
		logMissingParameterAndExit(c, []string{FlSmartSenseSubscriptionID.Name})
	}

	cbClient := NewCloudbreakOAuth2HTTPClient(c.String(FlServer.Name), c.String(FlUsername.Name), c.String(FlPassword.Name))
	deleteSmartSenseSubscriptionImpl(cbClient.Cloudbreak.Smartsensesubscriptions.DeleteSmartSenseSubscriptionByID, id)
	log.Infof("[deleteSmartSenseSubscription] SmartSense subscription has been deleted with id: %s ", id)
	return nil
}

func deleteSmartSenseSubscriptionImpl(deleteSmartSenseSubscriptionByID func(params *s.DeleteSmartSenseSubscriptionByIDParams) error, id int64) {
	err := deleteSmartSenseSubscriptionByID(&s.DeleteSmartSenseSubscriptionByIDParams{ID: id})

	if err != nil {
		logErrorAndExit(err)
	}
}

func DescribeSmartSenseSubscription(c *cli.Context) error {
	defer timeTrack(time.Now(), "describe of SmartSenseSubscription")
	checkRequiredFlags(c)

	cbClient := NewCloudbreakOAuth2HTTPClient(c.String(FlServer.Name), c.String(FlUsername.Name), c.String(FlPassword.Name))
	output := Output{Format: c.String(FlOutput.Name)}
	describeSmartSenseSubscriptionImpl(cbClient.Cloudbreak.Smartsensesubscriptions.GetSmartSenseSubscription, output)
	return nil
}

func describeSmartSenseSubscriptionImpl(
	getDefaultSmartSenseSubscription func(params *s.GetSmartSenseSubscriptionParams) (*s.GetSmartSenseSubscriptionOK, error),
	output Output) {

	ss, err := getDefaultSmartSenseSubscription(&s.GetSmartSenseSubscriptionParams{})

	if err != nil {
		logErrorAndExit(err)
	}

	log.Infof("[describeSmartSenseSubscription] describe SmartSense subscription for: %v ", ss.Payload.SubscriptionID)
	output.Write(SmartSenseSubscriptionHeader, &SmartSenseSubscription{Id: strconv.FormatInt(ss.Payload.ID, 10), SubscriptionId: *ss.Payload.SubscriptionID})
}

func getAvailableSmartSenseSubscription(c *cli.Context) (*models_cloudbreak.SmartSenseSubscriptionJSON, error) {
	cbClient := NewCloudbreakOAuth2HTTPClient(c.String(FlServer.Name), c.String(FlUsername.Name), c.String(FlPassword.Name))
	ss, err := cbClient.Cloudbreak.Smartsensesubscriptions.GetSmartSenseSubscription(&s.GetSmartSenseSubscriptionParams{})
	if err != nil {
		return nil, err
	} else if *ss.Payload.AutoGenerated {
		return nil, errors.New("SmartSense subscription generated by Cloudbreak")
	}
	return ss.Payload, nil
}

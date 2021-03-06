package sdk

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
)

func AWSCredsFromConfig(conf *Config) (*session.Session, *sts.Credentials) {
	return AWSCredsFromValues(conf.Region, conf.Profile, conf.AssumeRole)
}

func AWSCredsFromValues(region, profile string, assumeRole *AssumeRoleConfig) (*session.Session, *sts.Credentials) {
	sess := NewSession(region, profile)

	if assumeRole == nil {
		return sess, nil
	}

	assumed, creds, err := AssumeRole(sess, *assumeRole)
	if err != nil {
		panic(err)
	}

	return assumed, creds
}

package config

import (
	"errors"
	"flag"
	"fmt"
)

type Config struct {
	Remote string
	Sha    string
}

func Initialize(config *Config) error {
	flag.StringVar(&config.Remote, "remote", "", "-remote <git origin remote>")
	flag.StringVar(&config.Sha, "sha", "", "-sha <commit sha>")

	flag.Parse()

	return VerifyRequired(config)
}

func VerifyRequired(c *Config) error {
	errs := make([]string, 0, 2)
	if c.Remote == "" {
		errs = append(errs, "Remote not set")
	}
	if c.Sha == "" {
		errs = append(errs, "Sha not set")
	}

	if len(errs) > 0 {
		finalErrorText := "The following errors occurs on init:"
		for _, err := range errs {
			finalErrorText = fmt.Sprintf("%s\n\t%s", finalErrorText, err)
		}

		return errors.New(finalErrorText)
	}

	return nil
}

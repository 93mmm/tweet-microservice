package config

import (
	"errors"
	"fmt"
)

func genError(name string) error {
	return fmt.Errorf("Env %v is not valid", name)
}

func (c *config) validate() error {
	fmt.Println(c)
	var errs []error

	if len(c.App.Host) == 0 {
		errs = append(errs, genError("APP_HOST"))
	}

	if c.App.Port == 0 {
		errs = append(errs, genError("APP_PORT"))
	}

	if len(c.Mongo.Host) == 0 {
		errs = append(errs, genError("MONGO_HOST"))
	}

	if c.Mongo.Port == 0 {
		errs = append(errs, genError("MONGO_PORT"))
	}

	if errs != nil {
		return errors.Join(errs...)
	}
	return nil
}


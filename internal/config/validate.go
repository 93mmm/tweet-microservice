package config

import (
	"errors"
	"fmt"
	"log"
)

func genError(name string) error {
	return fmt.Errorf("Env %v is not valid or missing", name)
}

func isDefault(in any) bool {
	switch val := in.(type) {
	case string:
		return len(val) == 0
	case int:
		return val == 0
	default:
		log.Println("[WARN] Unknown type of field in config")
		return false
	}
}

func (c *config) validate() error {
	fmt.Println(c)
	var errs []error

	check := func(in any, env string) {
		if isDefault(in) {
			errs = append(errs, genError(env))
		}
	}
	check(c.App.Host, "APP_HOST")
	check(c.App.Port, "APP_PORT")
	check(c.Mongo.Host, "MONGO_HOST")
	check(c.Mongo.Port, "MONGO_PORT")

	if errs != nil {
		return errors.Join(errs...)
	}
	return nil
}


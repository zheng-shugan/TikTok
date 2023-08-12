package config

import (
	"testing"
)

func TestLoadConfigFromToml(t *testing.T) {

	err := LoadConfigFromYaml("../etc/config.yaml")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(C().Apps.HTTP.Name)
}

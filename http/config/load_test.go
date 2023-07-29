package conf

import (
	"os"
	"testing"
)

func TestLoadConfigFromToml(t *testing.T) {

	err := LoadConfigFromToml("../etc/demo.toml")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(C().App.Name)
}

func TestLoadConfigFromEnv(t *testing.T) {

	os.Setenv("MYSQL_DATABASE", "unit_test")
	err := LoadConfigFromEnv()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(C().MySQL.Database)
}

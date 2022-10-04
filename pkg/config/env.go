package config

import "github.com/spf13/viper"

func ConfigureEnv(v *viper.Viper) {
	v.SetEnvPrefix("")
	v.AllowEmptyEnv(true)
}

func EnvironmentalVariables(v *viper.Viper) {
	// e.g. v.BindEnv("SOME_ENVIRONMENTAL_VARIALBE", "databases.scylla.hosts")
}

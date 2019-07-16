package config

import "fmt"

func Help() {
	fmt.Println("https://xuri.me/toml-to-go")
}

func IsDebug() bool {
	return Default().Debug
}

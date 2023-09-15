package main

import "github.com/AnhQuanTrl/proflow/internal/config"

func main() {

}

func run() error {
	var cfg config.AppConfig
	cfg, err := config.InitConfig()
	if err != nil {
		return err
	}

	
}

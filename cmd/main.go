package main

import (
	"fmt"

	"github.com/cdutwhu/dev-nrt-splitter/config"
)

func main() {

	cfg := config.GetConfig("../config/config.toml", "./config/config.toml", "./config.toml")
	fmt.Println(cfg.InFolder)

	// csvfile := "./in/system_reports/actSystemDomainScores.csv"
	// csvtool.Split(csvfile, "./out", false, "School", "YrLevel", "Domain")

}

package conf

import (
	"encoding/xml"
	"os"

	"devsrv/logger"
)

//Configuration 配置项
type Configuration struct {
	DbHostip   string `xml:"dbhostip"`
	DbUsername string `xml:"dbusername"`
	DbPassword string `xml:"dbpassword"`
	DbName     string `xml:"dbname"`
}

//GConf 全局配置，来自配置文件
var GConf Configuration

func init() {
	xmlFile, err := os.Open("../etc/devsrv_conf.xml")
	if err != nil {
		logger.Info("Error opening file:", err)
		return
	}
	defer xmlFile.Close()
	if err := xml.NewDecoder(xmlFile).Decode(&GConf); err != nil {
		logger.Info("Error Decode file:", err)
		return
	}
}

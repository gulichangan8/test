package cmd

import (
	"test/api"
	"test/dao"
)

func main() {
	dao.InitMySql()
	api.InitEngine()
}

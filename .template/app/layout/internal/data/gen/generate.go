package main

import (
	"flag"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
	"kratos-demo/app/layout/internal/conf"
)

var flagConf string

func init() {
	flag.StringVar(&flagConf, "conf", "../../../configs", "config path, eg: -conf config.yaml")
}
func main() {
	c := config.New(
		config.WithSource(
			file.NewSource(flagConf),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	g := gen.NewGenerator(gen.Config{
		OutPath: "./query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	gormDB, _ := gorm.Open(mysql.Open(bc.Data.Database.Source))
	g.UseDB(gormDB)

	// generate all table from database
	//g.ApplyBasic(g.GenerateAllTable()...)
	g.ApplyBasic(g.GenerateModel(""))

	g.Execute()
}

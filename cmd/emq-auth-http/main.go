package main

import (
	"context"
	"fmt"
	"github.com/lishimeng/app-starter"
	etc2 "github.com/lishimeng/app-starter/etc"
	"github.com/lishimeng/emq-auth-http/cmd"
	"github.com/lishimeng/emq-auth-http/internal/api"
	"github.com/lishimeng/emq-auth-http/internal/etc"
	"github.com/lishimeng/go-log"
	"time"
)
import _ "github.com/lib/pq"

func main() {
	//orm.Debug = true

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	fmt.Println("start app:")
	fmt.Println(cmd.AppName)
	fmt.Println(cmd.Version)

	err := _main()
	if err != nil {
		fmt.Println(err)
	}
	time.Sleep(time.Second * 2)
}

func _main() (err error) {
	configName := "config"

	application := app.New()

	err = application.Start(func(ctx context.Context, builder *app.ApplicationBuilder) error {

		var err error
		err = builder.LoadConfig(&etc.Config, func(loader etc2.Loader) {
			loader.SetFileSearcher(configName, ".").SetEnvPrefix("").SetEnvSearcher()
		})
		if err != nil {
			return err
		}
		etc.Config.Web.Listen = ":80"
		//dbConfig := persistence.PostgresConfig{
		//	UserName:  etc.Config.Db.User,
		//	Password:  etc.Config.Db.Password,
		//	Host:      etc.Config.Db.Host,
		//	Port:      etc.Config.Db.Port,
		//	DbName:    etc.Config.Db.Database,
		//	InitDb:    true,
		//	AliasName: "default",
		//	SSL:       etc.Config.Db.Ssl,
		//}

		builder.
			//EnableDatabase(dbConfig.Build(),
			//model.Tables()...).
			SetWebLogLevel("debug").
			EnableWeb(etc.Config.Web.Listen, api.Route)
		//ComponentBefore(setup.JobClearExpireTask).
		//ComponentBefore(setup.BeforeStarted).
		//ComponentAfter(setup.AfterStarted)

		return err
	}, func(s string) {
		log.Info(s)
	})
	return
}

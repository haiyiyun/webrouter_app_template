package main

import (
	"flag"
	"os"
	"path"
	"runtime"

	"project_name/internal/app"

	"github.com/haiyiyun/config"
	"github.com/haiyiyun/log"
	"github.com/haiyiyun/webrouter"
)

func main() {
	appConfFile := flag.String("config.app", "../config/project_name/app.conf", "app config file")
	flag.Parse()
	appConf := app.Config{}
	config.Files(*appConfFile).Load(&appConf)

	serverDirPath := ""
	if appConf.LocateRelativeExecPath {
		serverDirPath = path.Dir(os.Args[0]) + "/"
	}

	log.SetLevel(appConf.LogDebugLevel)
	if appConf.LogDir != "" {
		if _, err := os.Stat(serverDirPath + appConf.LogDir); err != nil {
			os.MkdirAll(serverDirPath+appConf.LogDir, 0700)
		}
	}

	if appConf.LogFile != "" {
		logFi, err := os.OpenFile(serverDirPath+appConf.LogDir+"/"+appConf.LogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0700)
		if err != nil {
			log.Fatalln(err)
		}

		log.SetOutput(logFi)
	}

	runtime.GOMAXPROCS(appConf.MaxProcs)
	webrouter.DefaultServer.ReadTimeout = appConf.ServerReadTimeOut.Duration
	webrouter.DefaultServer.WriteTimeout = appConf.ServerWriteTimeOut.Duration
	webrouter.SetBeforeMethodName("Init")
	webrouter.SetBeforeMethodName("Before")

	if log.LEVEL_DEBUG&log.Levels() == 0 {
		log.SetFlags(log.Lproduction)
	}

	defer webrouter.Close()

	log.Info("Running at", appConf.BindAddr)
	if appConf.SslCertFile != "" && appConf.SslKeyFile != "" {
		if err := webrouter.ListenAndServeTLS(appConf.BindAddr, serverDirPath+appConf.SslCertFile, serverDirPath+appConf.SslKeyFile, nil); err != nil {
			log.Fatalln("ListenAndServeTLS error: ", err)
		}
	} else {
		if err := webrouter.ListenAndServe(appConf.BindAddr, nil); err != nil {
			log.Fatalln("ListenAndServe error: ", err)
		}
	}
}

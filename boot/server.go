package boot

import (
	g "OnlineJudge/app/global"
	"OnlineJudge/app/router"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func SeverSetup() {
	config := g.Config.Server
	gin.SetMode(config.Mode)
	routers := router.InitRouter()
	server := &http.Server{
		Addr:              config.Addr(),
		Handler:           routers,
		TLSConfig:         nil,
		ReadTimeout:       config.GetReadTimeout(),
		ReadHeaderTimeout: 0,
		WriteTimeout:      config.GetWriteTimeout(),
		IdleTimeout:       0,
		MaxHeaderBytes:    1 << 20,
	}
	g.Logger.Info("initialize server successfully", zap.String("port", config.Addr()))
	g.Logger.Error(server.ListenAndServe().Error())
}

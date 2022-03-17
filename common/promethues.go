package common

import (
	"fmt"
	"net/http"
	"strconv"

	log "github.com/micro/go-micro/v2/logger"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Promethues(port int) {
	httt.Handle("/metrics", promhttp.Handler())
	// 启动web服务
	go func() {
		err := http.ListenAndServe(fmt.Sprintf("%s:%d","0.0.0.0:", port), nil)
		if err != nil {
			log.Fatal("启动失败")
		}
		log.Debug("监控启动，端口为："+strconv.Itoa(port))
	}

}

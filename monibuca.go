package main

//	启动 monibuca

import (
	"context"
	"flag"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	. "github.com/Monibuca/engine/v3"
	// _ "github.com/Monibuca/plugin-ffmpeg"
	// _ "github.com/Monibuca/plugin-cluster"
	_ "github.com/Monibuca/plugin-gateway/v3"

	_ "github.com/Monibuca/plugin-gb28181/v3"
	_ "github.com/Monibuca/plugin-hdl/v3"
	_ "github.com/Monibuca/plugin-hls/v3"
	_ "github.com/Monibuca/plugin-jessica/v3"
	_ "github.com/Monibuca/plugin-logrotate/v3"
	_ "github.com/Monibuca/plugin-record/v3"
	_ "github.com/Monibuca/plugin-rtmp/v3"
	_ "github.com/Monibuca/plugin-rtsp/v3"
	_ "github.com/Monibuca/plugin-summary"
	_ "github.com/Monibuca/plugin-ts/v3"
	_ "github.com/Monibuca/plugin-webrtc/v3"
)

func monibuca_go(configname string) {
	configaddr := "configs/" + configname
	//addr := flag.String("c", configname, "config file")
	addr := &configaddr

	flag.Parse()
	ctx, cancel := context.WithCancel(context.Background())
	go waiter(cancel)
	if _, err := os.Stat(*addr); err == nil {
		Run(ctx, *addr)
	} else {
		Run(ctx, filepath.Join(filepath.Dir(os.Args[0]), *addr))
	}
}
func Monibuca(str string) {
	if str == "" {
		monibuca_go("config.toml")
	} else {
		monibuca_go(str)
	}
}

func waiter(cancel context.CancelFunc) {
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(sigc)
	<-sigc
	cancel()
}

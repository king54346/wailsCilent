package main

import (
	"embed"
	"nx/config"
	"syscall"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"

	"nx/utils"
)

//go:embed all:frontend/dist
var assets embed.FS

//go:embed tools/*
var tools embed.FS

//go:embed conf/*
var conf embed.FS

func main() {
	// 释放mtp_tools.exe文件
	err := utils.ExtractTools(tools)
	if err != nil {
		println("Extract Tools Error:", err.Error())
		syscall.Exit(1)
	}

	// 释放config文件
	err = utils.ExtractConfig(conf)
	if err != nil {
		println("Extract Config Error:", err.Error())
		syscall.Exit(1)
	}

	// 加载yaml文件
	cfg, err := config.LoadConfigFile()
	if err != nil {
		println("Load Config Error: ", err.Error())
	}

	// Create an instance of the app structure
	app := NewApp(cfg)

	// Create application with options
	err = wails.Run(&options.App{
		Title:  "nx",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}

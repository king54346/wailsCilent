package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"nx/config"
	"nx/mtp"
	"nx/utils"
)

// App struct
type App struct {
	ctx    context.Context
	config *config.Config
	Tasks  *utils.Tasks
}

// NewApp creates a new App application struct
func NewApp(config *config.Config) *App {
	return &App{config: config, Tasks: NewTasks()}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) GetAnnouncement() (msg utils.Msg) {
	url := a.config.GetRemoteAddress("/api/v1/user/announcement")

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error sending request to API endpoint. %+v\n", err)
		return utils.Msg{Code: 500, Msg: "Error sending request to API endpoint"}
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Couldn't read response body. %+v\n", err)
		return utils.Msg{Code: 500, Msg: "Couldn't read response body"}
	}

	err = json.Unmarshal(body, &msg)
	if err != nil {
		fmt.Printf("Error decoding JSON. %+v\n", err)
		return utils.Msg{Code: 500, Msg: "Error decoding JSON"}
	}

	if msg.Code != 0 {
		// return data with []string
		msg.Data = []string{}
	}
	return msg
}

func (a *App) UpdateAnnouncement(data []string) (msg utils.Msg) {
	url := a.config.GetRemoteAddress("/api/v1/admin/announcement")

	jsonStr, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("Error encoding JSON. %+v\n", err)
		return utils.Msg{Code: 500, Msg: "Error encoding JSON"}
	}

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		fmt.Printf("Error creating request. %+v\n", err)
		return utils.Msg{Code: 500, Msg: "Error creating request"}
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error sending request to API endpoint. %+v\n", err)
		return utils.Msg{Code: 500, Msg: "Error sending request to API endpoint"}
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Couldn't read response body. %+v\n", err)
		return utils.Msg{Code: 500, Msg: "Couldn't read response body"}
	}

	err = json.Unmarshal(body, &msg)
	if err != nil {
		fmt.Printf("Error decoding JSON. %+v\n", err)
		return utils.Msg{Code: 500, Msg: "Error decoding JSON"}
	}

	return msg
}

func (a *App) GetDeviceList() []mtp.Device {
	return mtp.GetDeviceList()
}

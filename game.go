package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	netURL "net/url"
	"nx/utils"
)

func (a *App) GetGameList(searchText string) (games []utils.Game) {
	url := a.config.GetRemoteAddress("/api/v1/admin/game/list")

	if searchText != "" {
		// URL编码searchText
		url += fmt.Sprintf("?search=%s", netURL.QueryEscape(searchText))
	}

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error sending request to API endpoint. %+v\n", err)
		return nil
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Couldn't read response body. %+v\n", err)
		return nil
	}

	err = json.Unmarshal(body, &games)
	if err != nil {
		fmt.Printf("Error decoding JSON. %+v\n", err)
		return nil
	}

	return games
}

func (a *App) AddGame(games []utils.Game) (msg utils.Msg) {
	url := a.config.GetRemoteAddress("/api/v1/admin/game/add")

	// 检查game的size是否填写，没有填写就自动统计（GB）
	for index, game := range games {
		if game.Size == 0 {
			games[index].Size = game.SizeCount()
		}
	}

	// 将结构体切片转换为json字符串
	jsonStr, err := json.Marshal(games)
	if err != nil {
		fmt.Printf("Error encoding JSON. %+v\n", err)
		return utils.Msg{Code: 500, Msg: "Error encoding JSON"}
	}

	// 发送post请求
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		fmt.Printf("Error sending request to API endpoint. %+v\n", err)
		return utils.Msg{Code: 500, Msg: "Error sending request to API endpoint"}
	}
	defer resp.Body.Close()

	// 读取响应体
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

func (a *App) UpdateGame(game utils.Game) (msg utils.Msg) {
	url := a.config.GetRemoteAddress(fmt.Sprintf("/api/v1/admin/game/update/%d", game.Id))

	if game.Size == 0 {
		game.Size = game.SizeCount()
	}

	// 将结构体转换为json字符串
	jsonStr, err := json.Marshal(game)
	if err != nil {
		fmt.Printf("Error encoding JSON. %+v\n", err)
		return utils.Msg{Code: 500, Msg: "Error encoding JSON"}
	}

	// 发送put请求
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

	// 读取响应体
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

func (a *App) DeleteGame(id []int) (msg utils.Msg) {
	url := a.config.GetRemoteAddress("/api/v1/admin/game/delete")

	// 将结构体切片转换为json字符串
	// {"id": id}
	jsonStr, err := json.Marshal(utils.GameDelete{Id: id})
	if err != nil {
		fmt.Printf("Error encoding JSON. %+v\n", err)
		return utils.Msg{Code: 500, Msg: "Error encoding JSON"}
	}

	// 发送delete请求
	req, err := http.NewRequest("DELETE", url, bytes.NewBuffer(jsonStr))
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

	// 读取响应体
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

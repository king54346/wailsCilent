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

func (a *App) GetWorkOrder(query utils.WorkOrderQuery) (result any) {
	url := a.config.GetRemoteAddress("/api/v1/admin/workorder/list")

	// query is get params
	url += fmt.Sprintf("?pageNo=%d&pageSize=%d&status=%d&search=%s", query.PageNo, query.PageSize, query.Status, netURL.QueryEscape(query.Search))

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

	var tmp utils.WorkOrderQueryResult
	err = json.Unmarshal(body, &tmp)
	if err != nil {
		fmt.Printf("Error decoding JSON. %+v\n", err)
		return utils.Msg{Code: 500, Msg: "Error decoding JSON"}
	}

	return tmp
}

func (a *App) DeleteWorkOrder(id []int) (msg utils.Msg) {
	url := a.config.GetRemoteAddress("/api/v1/admin/workorder/delete")

	jsonStr, err := json.Marshal(utils.WorkOrderDelete{Id: id})
	if err != nil {
		fmt.Printf("Error encoding JSON. %+v\n", err)
		return utils.Msg{Code: 500, Msg: "Error encoding JSON"}
	}

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

func (a *App) GetOrderGameList(id int) []utils.WorkOrderGameList {
	url := a.config.GetRemoteAddress("/api/v1/admin/workorder/game/list")
	url += fmt.Sprintf("?id=%d", id)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error sending request to API endpoint. %+v\n", err)
		return []utils.WorkOrderGameList{}
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Couldn't read response body. %+v\n", err)
		return []utils.WorkOrderGameList{}
	}

	var tmp []utils.WorkOrderGameList
	err = json.Unmarshal(body, &tmp)
	if err != nil {
		fmt.Printf("Error decoding JSON. %+v\n", err)
		return []utils.WorkOrderGameList{}
	}

	return tmp
}

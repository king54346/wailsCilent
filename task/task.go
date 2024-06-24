package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"nx/config"
	"nx/mtp"
	"nx/utils"
	"strings"
)

type UpdateGameStatusStruct struct {
	OrderID  int    `json:"order_id"`
	GameID   int    `json:"game_id"`
	Status   int    `json:"status"`
	ErrorMsg string `json:"error_msg"`
}

type UpdateTaskStatusStruct struct {
	OrderID int `json:"id"`
	Status  int `json:"status"`
}

func UpdateOrderGameStatus(cfg *config.Config, workOrderID, gameID, status int, errMsg string) error {
	url := cfg.GetRemoteAddress("/api/v1/admin/workorder/game/update")

	updateStruct := UpdateGameStatusStruct{
		OrderID:  workOrderID,
		GameID:   gameID,
		Status:   status,
		ErrorMsg: errMsg,
	}

	data, err := json.Marshal(updateStruct)
	if err != nil {
		fmt.Printf("UpdateOrderGameStatus Error: %v\n", err)
		return err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		fmt.Printf("UpdateOrderGameStatus Error: %v\n", err)
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Couldn't read response body. %+v\n", err)
		return err
	}

	var msg utils.Msg
	err = json.Unmarshal(body, &msg)
	if err != nil {
		fmt.Printf("Error decoding JSON. %v\n", err)
		return err
	}

	if msg.Code != 0 {
		fmt.Printf("UpdateOrderGameStatus Error: %v\n", msg.Msg)
		return fmt.Errorf(msg.Msg)
	}

	return nil

}

func UpdateTaskStatus(cfg *config.Config, workOrderID, status int) error {
	url := cfg.GetRemoteAddress("/api/v1/admin/workorder/update")

	updateStruct := UpdateTaskStatusStruct{
		OrderID: workOrderID,
		Status:  status,
	}

	data, err := json.Marshal(updateStruct)
	if err != nil {
		fmt.Printf("UpdateTaskStatus Error: %v\n", err)
		return err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		fmt.Printf("UpdateTaskStatus Error: %v\n", err)
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Couldn't read response body. %+v\n", err)
		return err
	}

	var msg utils.Msg
	err = json.Unmarshal(body, &msg)
	if err != nil {
		fmt.Printf("Error decoding JSON. %v\n", err)
		return err
	}

	if msg.Code != 0 {
		fmt.Printf("UpdateTaskStatus Error: %v\n", msg.Msg)
		return fmt.Errorf(msg.Msg)
	}

	return nil
}

func SendFilesTask(workOrderID int, machineID string, tasks *utils.Tasks, cfg *config.Config) {
	var task utils.Task
	var games []utils.Game
	task = tasks.Tasks[workOrderID]
	games = task.Games

	// Get Directory ID
	parentID, err := mtp.FindObject(machineID, cfg.DirName)
	if err != nil {
		fmt.Printf("workOrderID: %d, FindObject Error: %v", workOrderID, err)
		UpdateTaskStatus(cfg, workOrderID, 2)
		// remove task
		tasks.Lock.Lock()
		delete(tasks.Tasks, workOrderID)
		tasks.Lock.Unlock()
		return
	}

	UpdateTaskStatus(cfg, workOrderID, 1)
	var noError bool = true
	for _, game := range games {
		task = tasks.Tasks[workOrderID]

		// Check if task.Run is false
		tasks.Lock.Lock()
		if !task.Run {
			UpdateTaskStatus(cfg, workOrderID, 4)
			delete(tasks.Tasks, workOrderID)
			tasks.Lock.Unlock()
			return
		}
		tasks.Lock.Unlock()

		UpdateOrderGameStatus(cfg, workOrderID, game.Id, 1, "")

		// Send file
		if !utils.CheckPathExist(game.Filepath) {
			fmt.Printf("workOrderID: %d, game ID: %d, File not exist\n", workOrderID, game.Id)
			UpdateOrderGameStatus(cfg, workOrderID, game.Id, 2, "目标文件不存在！请更新游戏路径！")
			continue
		}

		switch utils.CheckPathDirOrFile(game.Filepath) {
		case "file":
			// Send file
			status := 3
			errMsg := ""
			err = mtp.SendFile(machineID, parentID, game.Filepath)
			if err != nil {
				fmt.Printf("workOrderID: %d, game ID: %d, SendFile Error: %v\n", workOrderID, game.Id, err.Error())
				noError = false
				status = 2
				errMsg = fmt.Sprintf("文件 %s 发送失败，错误信息：%v", game.Filepath, err.Error())
			}
			UpdateOrderGameStatus(cfg, workOrderID, game.Id, status, errMsg)

		case "dir":
			// Walk directory and send files
			var errors []string
			files, err := utils.GetFilesListInDir(game.Filepath)
			if err != nil {
				fmt.Printf("workOrderID: %d, game ID: %d, GetFilesListInDir Error: %v\n", workOrderID, game.Id, err)
				noError = false
				UpdateOrderGameStatus(cfg, workOrderID, game.Id, 2, fmt.Sprintf("获取文件列表失败！错误信息：%v", err.Error()))
				continue
			}

			for _, file := range files {
				err = mtp.SendFile(machineID, parentID, file)
				if err != nil {
					fmt.Printf("workOrderID: %d, game ID: %d, File: %v, SendFile Error: %v\n", workOrderID, game.Id, file, err.Error())
					errors = append(errors, fmt.Sprintf("文件 %s 发送失败，错误信息：%v", file, err.Error()))
					noError = false
				}
			}
			if len(errors) > 0 {
				fmt.Printf("workOrderID: %d, game ID: %d, SendFiles Error: %v\n", workOrderID, game.Id, errors)
				UpdateOrderGameStatus(cfg, workOrderID, game.Id, 2, fmt.Sprintf("发送文件失败！错误信息：%v", strings.Join(errors, "\n")))
			} else {
				UpdateOrderGameStatus(cfg, workOrderID, game.Id, 3, "")
			}
		default:
			fmt.Printf("workOrderID: %d, game ID: %d, Filepath is not a file or directory: %v\n", workOrderID, game.Id, game.Filepath)
			noError = false
			UpdateOrderGameStatus(cfg, workOrderID, game.Id, 2, "目标文件类型错误！请更新游戏路径！")
		}
	}

	if noError {
		fmt.Printf("workOrderID: %d, SendFilesTask success!\n", workOrderID)
		UpdateTaskStatus(cfg, workOrderID, 3)
	} else {
		fmt.Printf("workOrderID: %d, SendFilesTask failed!\n", workOrderID)
		UpdateTaskStatus(cfg, workOrderID, 2)
	}

	tasks.Lock.Lock()
	delete(tasks.Tasks, workOrderID)
	tasks.Lock.Unlock()
}

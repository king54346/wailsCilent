package main

import (
	nxTask "nx/task"
	"nx/utils"
	"sync"
)

func NewTasks() *utils.Tasks {
	return &utils.Tasks{Tasks: map[int]utils.Task{}, Lock: sync.Mutex{}}
}

func (a *App) AddTask(workOrderID int, machineID string) utils.Msg {
	result := utils.GetWorkOrderGameList(workOrderID, a.config)

	task := utils.Task{Games: []utils.Game{}, Run: true}
	switch v := result.(type) {
	case utils.Msg:
		return v
	case []utils.Game:
		if len(v) == 0 {
			return utils.Msg{Code: 404, Msg: "对应工单没有勾选任何游戏！"}
		}
		task.Games = v
	default:
	}

	if a.Tasks == nil {
		a.Tasks = NewTasks()
	}

	a.Tasks.Lock.Lock()
	// check workOrderID is in Tasks
	if _, ok := (*a.Tasks).Tasks[workOrderID]; ok {
		a.Tasks.Lock.Unlock()
		return utils.Msg{Code: 400, Msg: "任务已存在！"}
	}
	a.Tasks.Tasks[workOrderID] = task
	a.Tasks.Lock.Unlock()

	go nxTask.SendFilesTask(workOrderID, machineID, a.Tasks, a.config)

	return utils.Msg{Code: 200, Msg: "添加任务成功！"}
}

func (a *App) CancelTask(workOrderID int) utils.Msg {
	if a.Tasks == nil {
		return utils.Msg{Code: 404, Msg: "任务不存在！"}
	}

	a.Tasks.Lock.Lock()
	defer a.Tasks.Lock.Unlock()

	if _, ok := a.Tasks.Tasks[workOrderID]; ok {
		task := a.Tasks.Tasks[workOrderID]
		task.Run = false
		a.Tasks.Tasks[workOrderID] = task
		return utils.Msg{Code: 200, Msg: "取消任务成功！"}
	}

	return utils.Msg{Code: 404, Msg: "任务不存在！"}
}

package utils

import "sync"

type Msg struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

type Game struct {
	Id       int      `json:"id"`
	Name     string   `json:"name"`
	Size     float64  `json:"size"`
	Tags     []string `json:"tags"`
	Filepath string   `json:"filepath"`
}

type GameDelete struct {
	Id []int `json:"id"`
}

type WorkOrderQuery struct {
	PageNo   int    `json:"pageNo"`
	PageSize int    `json:"pageSize"`
	Search   string `json:"search"`
	Status   int    `json:"status"`
}

type WorkOrderItem struct {
	Id      int    `json:"id"`
	Phone   string `json:"phone"`
	TbOrder string `json:"tb_order"`
	Machine string `json:"machine"`
	Sdcard  int    `json:"sdcard"`
	Games   []int  `json:"games"`
	Status  string `json:"status"`
}

type WorkOrderResultInfo struct {
	TotalCount int             `json:"totalCount"`
	TotalPage  int             `json:"totalPage"`
	PageNo     int             `json:"pageNo"`
	PageSize   int             `json:"pageSize"`
	Data       []WorkOrderItem `json:"data"`
}

type WorkOrderQueryResult struct {
	Result WorkOrderResultInfo `json:"result"`
}

type WorkOrderUpdater struct {
	Id     int `json:"id"`
	Status int `json:"status"`
}

type WorkOrderDelete struct {
	Id []int `json:"id"`
}

type WorkOrderGameList struct {
	Id       int     `json:"id"`
	Name     string  `json:"name"`
	Size     float32 `json:"size"`
	Status   string  `json:"status"`
	ErrorMsg string  `json:"error_msg"`
}

type Task struct {
	Games []Game
	Run   bool
}

type Tasks struct {
	Tasks map[int]Task
	Lock  sync.Mutex
}

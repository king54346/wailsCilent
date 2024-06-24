package utils

import (
	"bytes"
	"crypto/md5"
	"embed"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"nx/config"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
)

func GetMD5(data []byte) string {
	// 创建一个 MD5 哈希对象
	hash := md5.New()

	// 将数据写入哈希对象
	hash.Write(data)

	// 获取最终的哈希值
	hashInBytes := hash.Sum(nil)

	// 将哈希值转换为十六进制字符串
	hashInString := hex.EncodeToString(hashInBytes)
	return hashInString
}

func CheckPathExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}
	return true
}

// 检查目标路径是文件夹还是文件，"dir"代表目录，"file"代表文件，"none"代表不存在
func CheckPathDirOrFile(path string) string {
	f, err := os.Stat(path)
	if err != nil {
		return "none"
	}

	if f.IsDir() {
		return "dir"
	}

	return "file"
}

func GetFilesListInDir(dir string) ([]string, error) {
	var files []string

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// 如果是文件，则添加到文件列表中
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})

	return files, err
}

func ExtractTools(tools embed.FS) error {
	// 检查本地文件tools/mtp_tools.exe是否存在
	data, err := os.ReadFile("tools/mtp_tools.exe")
	if err != nil {
		// 本地文件不存在，从embed.FS中提取
		data, err = tools.ReadFile("tools/mtp_tools.exe")
		if err != nil {
			return err
		}

		if !CheckPathExist("tools") {
			err = os.MkdirAll("tools", 0755)
			if err != nil {
				return err
			}
		}

		err = os.WriteFile("tools/mtp_tools.exe", data, 0755)
		if err != nil {
			return err
		}

		return nil
	}

	MD5InDisk := GetMD5(data)

	// 计算 tools/mtp_tools.exe 的 MD5 值
	data, err = tools.ReadFile("tools/mtp_tools.exe")
	if err != nil {
		return err
	}

	MD5InProgram := GetMD5(data)

	if MD5InDisk != MD5InProgram {
		// 本地文件与embed.FS中的文件MD5值不同，更新本地文件
		err = os.WriteFile("tools/mtp_tools.exe", data, 0755)
		if err != nil {
			return err
		}

		return nil
	}

	return nil
}

func ExtractConfig(conf embed.FS) error {
	// 检查本地文件conf/config.json是否存在
	_, err := os.ReadFile("conf/config.yaml")
	if err != nil {
		// 本地文件不存在，从embed.FS中提取
		data, err := conf.ReadFile("conf/config.yaml")
		if err != nil {
			return err
		}

		if !CheckPathExist("conf") {
			err = os.MkdirAll("conf", 0755)
			if err != nil {
				return err
			}
		}

		err = os.WriteFile("conf/config.yaml", data, 0644)
		if err != nil {
			return err
		}
		return nil
	}
	return nil
}

func RunCommand(name string, args ...string) (returnCode int, output bytes.Buffer, err error) {
	cmd := exec.Command(name, args...)

	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	cmd.Stdout = &output

	err = cmd.Run()
	if err != nil {
		returnCode = -1
		return
	}

	returnCode = cmd.ProcessState.ExitCode()
	return
}

func GetWorkOrderGameList(workOrderID int, config *config.Config) interface{} {
	url := config.GetRemoteAddress("/api/v1/admin/game/get?id=" + fmt.Sprintf("%d", workOrderID))

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error sending request to API endpoint. %+v\n", err)
		return Msg{Code: 500, Msg: "请求远程数据失败！请检查网络连接！"}
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Couldn't read response body. %+v\n", err)
		return Msg{Code: 500, Msg: "读取远程响应数据失败！"}
	}

	var game []Game
	err = json.Unmarshal(body, &game)
	if err != nil {
		fmt.Printf("Error decoding JSON. %+v\n", err)
		return Msg{Code: 500, Msg: "解析响应数据失败！"}
	}

	return game
}

func (g *Game) SizeCount() float64 {
	var size float64

	switch CheckPathDirOrFile(g.Filepath) {
	case "file":
		fileInfo, err := os.Stat(g.Filepath)
		if err != nil {
			return 0
		}
		size = float64(fileInfo.Size()) / 1024 / 1024 / 1024
	case "dir":
		files, err := GetFilesListInDir(g.Filepath)
		if err != nil {
			return 0
		}

		for _, file := range files {
			fileInfo, err := os.Stat(file)
			if err != nil {
				return 0
			}
			size += float64(fileInfo.Size()) / 1024 / 1024 / 1024
		}

	default:
		return 0
	}

	return size
}

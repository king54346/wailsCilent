package mtp

import (
	"fmt"
	"nx/utils"
	"regexp"
	"strings"
)

type Device struct {
	ID     string
	VID    string
	PID    string
	UID    string
	Name   string
	Serial string
}

func GetDeviceList() []Device {
	returnCode, output, err := utils.RunCommand("./tools/mtp_tools.exe", "list")
	if err != nil {
		fmt.Println("Execute command to list devices error! Error: ", err.Error())
		return []Device{}
	} else if returnCode != 0 {
		fmt.Println("Execute command to list devices error! Error: ", output.String())
		return []Device{}
	}

	reg := regexp.MustCompile(`ID: (.*vid_(\S{4})&pid_(\S{4}).*\{([0-9a-zA-Z]{8}-[0-9a-zA-Z]{4}-[0-9a-zA-Z]{4}-[0-9a-zA-Z]{4}-[0-9a-zA-Z]{12})\})`)

	matches := reg.FindAllStringSubmatch(output.String(), -1)

	devices := make([]Device, 0)

	for _, match := range matches {
		device := Device{
			ID:     match[1],
			VID:    match[2],
			PID:    match[3],
			UID:    match[4],
			Name:   GetFriendlyName(match[1]),
			Serial: GetSerialNumber(match[1]),
		}
		devices = append(devices, device)
	}

	return devices
}

func GetFriendlyName(ID string) string {
	returnCode, output, err := utils.RunCommand("./tools/mtp_tools.exe", "name", ID)
	if err != nil {
		fmt.Println("Execute command to get name error! Error: ", err.Error())
		return ""
	} else if returnCode != 0 {
		fmt.Println("Execute command to get name error! Error: ", output.String())
		return ""
	}

	fmt.Println("output: ", strings.TrimSpace(strings.TrimPrefix(output.String(), "Name: ")))

	return strings.TrimSpace(strings.TrimPrefix(output.String(), "Name: "))
}

func GetSerialNumber(ID string) string {
	returnCode, output, err := utils.RunCommand("./tools/mtp_tools.exe", "serial", ID)
	if err != nil {
		fmt.Println("Execute command to get serial number error! Error: ", err.Error())
		return ""
	} else if returnCode != 0 {
		fmt.Println("Execute command to get serial number error! Error: ", output.String())
		return ""
	}

	return strings.TrimSpace(strings.TrimPrefix(output.String(), "Serial Number: "))
}

func FindObject(ID string, objectName string) (string, error) {
	returnCode, output, err := utils.RunCommand("./tools/mtp_tools.exe", "find", ID, objectName)
	if err != nil {
		fmt.Println("Execute command to find object error! Error: ", err.Error())
		return "", err
	} else if returnCode != 0 {
		fmt.Println("Execute command to find object error! Error: ", output.String())
		return "", fmt.Errorf("execute command find returns not zero, return code: %v", returnCode)
	}

	return strings.TrimSpace(strings.TrimPrefix(output.String(), "Object ID: ")), nil
}

func SendFile(ID string, parentID string, filepath string) error {
	returnCode, output, err := utils.RunCommand("./tools/mtp_tools.exe", "send", ID, parentID, filepath)
	if err != nil {
		fmt.Println("Execute command to send file error! Error: ", err.Error())
		return err
	} else if returnCode != 0 {
		fmt.Println("Execute command to send file error! Error: ", output.String())
		return fmt.Errorf("execute command to send file returns not zero, return code: %v", returnCode)
	}
	return nil
}

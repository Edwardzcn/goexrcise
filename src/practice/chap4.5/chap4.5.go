package main

import (
	"log"
	"strings"
)

func main() {
	iniData := []string{
		"; Cut down copy of Mozilla application.ini file",
		"",
		"[App]",
		"Vendor=Mozilla",
		"Name=Iceweasel",
		"Profile=mozilla/firefox",
		"Version=3.5.16",
		"[Gecko]",
		"MinVersion=1.9.1",
		"MaxVersion=1.9.1.*",
		"[XRE]",
		"EnableProfileMigrator=0",
		"EnableExtensionManager=1",
	}
	ini := PaseIni(iniData)
	PrintIni(ini)
}

func PaseIni(iniData []string) (returnData map[string]map[string]string) {
	var firstMapKey string
	// 返回值命名也仅仅初始化字典为零值，即nil 仍需要make
	returnData = make(map[string]map[string]string)
	for _, line := range iniData {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, ";") {
			continue
		}
		if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
			firstMapKey = line[1 : len(line)-1]
			if _, find := returnData[firstMapKey]; find == false {
				returnData[firstMapKey] = make(map[string]string)
			}
		} else if index := strings.Index(line, "="); index > -1 {
			secondMapKey := line[:index]
			secondMapValue := line[index+1:]
			returnData[firstMapKey][secondMapKey] = secondMapValue
		} else {
			log.Print("failed to parse line:", line)
		}
	}
	return returnData
}

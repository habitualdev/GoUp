// +build linux
package install

import (
	"GoUp/get"
	"os/exec"
	"regexp"
	"strings"
)



func Install(dlLinks []string) {
	reCompile, _ := regexp.Compile("linux")
	for _, link := range dlLinks {
		if reCompile.Match([]byte(link)){
			get.DownloadFile("go.tar.gz", link)
			break
		}
	}
	cmdSlice := strings.Fields("rm -rf /usr/local/go && tar -C /usr/local -xzf go.tar.gz")
	c := exec.Command(cmdSlice[0], cmdSlice[1:]...)
	c.Run()
}
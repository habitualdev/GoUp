// +build darwin
package install

import (
	"GoUp/get"
	"os/exec"
	"regexp"
)

func Install(dlLinks []string) {
	reCompile, _ := regexp.Compile("darwin")
	for _, link := range dlLinks {
		if reCompile.Match([]byte(link)){
			get.DownloadFile("go.pkg", link)
			break
		}
	}
	c := exec.Command("sudo","open","go.pkg")
	c.Run()
}
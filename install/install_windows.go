// +build windows

package install

import (
	"GoUp/get"
	"os"
	"os/exec"
	"regexp"
)

func Start(args ...string) (p *os.Process, err error) {
	if args[0], err = exec.LookPath(args[0]); err == nil {
		var procAttr os.ProcAttr
		procAttr.Files = []*os.File{os.Stdin,
			os.Stdout, os.Stderr}
		p, err := os.StartProcess(args[0], args, &procAttr)
		if err == nil {
			return p, nil
		}
	}
	return nil, err
}

func Install(dlLinks []string){

	reCompile, _ := regexp.Compile("windows")
	for _, link := range dlLinks {
		if reCompile.Match([]byte(link)){
			get.DownloadFile("go.msi", link)
			break
		}
	}

	Start("msiexec", "/a", "go.msi", "/passive")
}
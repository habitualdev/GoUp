package main

import (
	"GoUp/get"
	"GoUp/install"
)

func main() {

	dlLinks := get.GetLink()

	install.Install(dlLinks)

}

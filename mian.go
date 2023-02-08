package main

import (
	gHttp "yxProject/http"
	glog "yxProject/log"
	gText "yxProject/text"
	gtime "yxProject/time"
)

func main() {
	gHttp.Test()
	gtime.Test()
	glog.Test()
	gText.Test()

}

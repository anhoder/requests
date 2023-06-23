package main

import "github.com/go-musicfox/requests"

func main() {

	resp, _ := requests.Get("http://go.xiulian.net.cn")
	println(resp.Text())
}

package main

import "github.com/anhoder/requests"

func main() {

	resp, _ := requests.Get("http://go.xiulian.net.cn")
	println(resp.Text())
}

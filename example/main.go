package main

import (
	docker "dockerapi"
	"fmt"
)

func main() {

	//创建一个客户端，通常来说整个app只需要一个客户端
	cli := docker.New()
	defer cli.Close()

	// 创建容器
	// 这里可选配置宿主机向容器内映射的端口：`outPort`。
	// 若不配置OutPort，则会在容器Start之后将docker映射的宿主机随机端口赋值给`outPort`
	image1 := docker.Image{
		ImageName: "mashiroc/ctfeasysql", //你的docker镜像名字
		InPort:    80,                    //容器内向外映射的端口
	}

	// 若配置宿主机映射端口
	_ = docker.Image{
		ImageName: "OtherImage",
		InPort:    80,
		OutPort:   4949,
	}
	c, err := cli.NewContainer(image1)
	if err != nil {
		panic(err)
	}

	if err := c.Start(); err != nil {
		panic(err)
	}
	fmt.Println("your out port:", c.OutPort())

	// 建议Stop或Kill后Remove
	// 这里可以Stop或Kill
	if err := c.Stop(); err != nil {
		panic(err)
	}

	if err := c.Remove(); err != nil {
		panic(err)
	}
}

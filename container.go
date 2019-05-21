package dockerapi

import "strconv"

// Images Image array
type Images []Image

//Image 镜像配置
type Image struct {
	ImageName string
	InPort    uint16
	OutPort   uint16
}

//Containers container array
type Containers []Container

//Container docker容器
type Container struct {
	*Client
	id        string
	imageName string
	inPort    string
	outPort   string
	isStart   bool
}

//InPort 获得内部端口
func (c *Container) InPort() uint16 {
	num, _ := strconv.ParseUint(c.inPort, 10, 16)
	return uint16(num)
}

//OutPort 获得外部端口
func (c *Container) OutPort() uint16 {
	num, _ := strconv.ParseUint(c.outPort, 10, 16)
	return uint16(num)
}

//ImageName 获得镜像名
func (c *Container) ImageName() string {
	return c.imageName
}

//ID 获得镜像ID
func (c *Container) ID() string {
	return c.id
}

//IsStart 是否已经启动
func (c *Container) IsStart() bool {
	return c.isStart
}

package dockerapi

import (
	"github.com/docker/docker/api/types"
	"github.com/docker/go-connections/nat"
	"time"
)

// Start 启动一个容器
func (c *Container) Start() (err error) {

	//启动容器
	if err = c.cli.ContainerStart(c.ctx, c.id, types.ContainerStartOptions{}); err != nil {
		return
	}

	// 若无宿主机端口配置，则通过docker container inspect获得宿主机端口
	if len(c.outPort) == 0 {
		//拿到启动后的json参数
		containerJSON, err := c.cli.ContainerInspect(c.ctx, c.id)
		if err != nil {
			return
		}

		// 取端口
		// 因为启动端口是容器来做的，只做了一对映射映射，这里就直接用inPort取出映射的数组中的第一个port
		containerPorts := containerJSON.NetworkSettings.Ports
		portBinding := containerPorts[nat.Port(c.inPort)]
		c.outPort = portBinding[0].HostPort
	}
	c.isStart = true
	return
}

// Stop 关闭容器
func (c *Container) Stop() (err error) {
	timeout := time.Second * 10
	err = c.cli.ContainerStop(c.ctx, c.id, &timeout)
	c.isStart = false
	return
}

// Kill 杀掉容器
func (c *Container) Kill() (err error) {
	err = c.cli.ContainerKill(c.ctx, c.id, "KILL")
	return
}

// Remove 移除容器资源
// 建议在kill或stop后remove
func (c *Container) Remove() (err error) {
	opt := types.ContainerRemoveOptions{}
	err = c.cli.ContainerRemove(c.ctx, c.id, opt)
	return
}



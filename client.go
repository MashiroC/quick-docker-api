package dockerapi

import (
	"context"
	dockerContainer "github.com/docker/docker/api/types/container"
	dockerClient "github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"github.com/pkg/errors"
	"strconv"
	"strings"
)

// Client 操作dockerapi的客户端
type Client struct {
	cli *dockerClient.Client
	ctx context.Context
}

// New 创建一个Client
func New() (c *Client) {
	cli, err := dockerClient.NewClientWithOpts(dockerClient.WithVersion("1.37"))
	if err != nil {
		panic(err)
	}

	c = &Client{
		cli: cli,
		ctx: context.Background(),
	}
	return
}

// NewContainer 根据镜像创建一个容器
// 若镜像未配置outPort 会赋予一个随机的outPort
func (c *Client) NewContainer(i Image) (con *Container, err error) {

	handle := &Container{
		imageName: i.ImageName,
		inPort:    strconv.Itoa(int(i.InPort)),
		Client:    c,
	}

	// 配置端口映射
	exports := make(nat.PortSet, 1)
	port, _ := nat.NewPort("tcp", handle.inPort)

	exports[port] = struct{}{}

	ports := make(nat.PortMap)
	pb := make([]nat.PortBinding, 1)
	pb[0] = nat.PortBinding{}
	//若配置了OutPort端口
	if i.OutPort != 0 {
		out := strconv.Itoa(int(i.OutPort))
		pb[0].HostPort = out
		con.outPort = out
	}
	ports[port] = pb

	config := &dockerContainer.Config{
		Image:        handle.imageName,
		ExposedPorts: exports}
	hostConfig := &dockerContainer.HostConfig{PortBindings: ports}

	//创建容器
	body, err := c.cli.ContainerCreate(c.ctx, config, hostConfig, nil, "")
	if err != nil {
		return
	}

	handle.id = body.ID

	if len(body.Warnings) != 0 {
		err = errors.New(strings.Join(body.Warnings, ";"))
		return
	}
	return handle, nil
}

// Close 关闭客户端
func (c *Client) Close() {
	if err := c.cli.Close(); err != nil {
		panic(err)
	}
}

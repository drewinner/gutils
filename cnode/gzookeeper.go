package cnode

/**
*   @Time : 2019/6/27 9:17
*   @Author : wanghongli
*   @File : gzookeeper
*   @Software: GoLand
 */

import (
	"errors"
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

type ZKClient struct {
	path        string        //路径
	name        string        //名称
	addr        []string      //集群地址
	conn        *zk.Conn      //zk connections
	version     int32         //最新版本号
	versionLock *sync.RWMutex //版本锁
	closed      bool          //是否关闭
	lock        *sync.RWMutex
	snapshot    map[string][]byte //当前快照
}

/**
*	创建zk连接
*	@param path路径 addrs zk集群地址
 */
func NewZKClient(path string, addrs ...string) (client *ZKClient, err error) {
	if len(addrs) == 0 {
		return nil, errors.New("zk addr error")
	}
	client = new(ZKClient)
	client.path = path
	client.closed = false
	client.lock = new(sync.RWMutex)
	client.versionLock = new(sync.RWMutex)
	err = client.connect(addrs)
	if err != nil {
		fmt.Printf("the err is :%+v", err)
		return nil, err
	}
	return
}

/**
*	连接zk节点
*	@param:addrs zk集群地址
 */
func (c *ZKClient) connect(addrs []string) (err error) {
	if len(addrs) == 0 {
		return errors.New("zk addr error")
	}
	c.conn, _, err = zk.Connect(addrs, time.Second)
	if err != nil {
		return err
	}
	c.addr = addrs
	return nil
}

/**
*	注册临时节点
*	@param:name 临时节点名称 value为自定义信息 权重，开关，ip地址等
 */
func (c *ZKClient) Register(name string, value []byte) (err error) {
	c.lock.RLock()
	defer c.lock.RUnlock()
	if c.closed {
		return errors.New("Closed zk instance ")
	}
	c.name = name
	//确保使用的目录存在
	err = c.ensurePath(c.path)
	if err != nil {
		return err
	}
	flags := int32(zk.FlagEphemeral) //短暂，session断开则改节点也被删除
	acl := zk.WorldACL(zk.PermAll)   //所有都可以访问
	dest := filepath.Join(c.path, name)
	if strings.Contains(dest, `\`) {
		dest = strings.Replace(dest, `\`, `/`, -1)
	}
	_, err = c.conn.Create(dest, value, flags, acl)
	if err != nil {
		fmt.Printf("Register error:%+v", err)
		return err
	}
	return nil
}

/**
*	判断节点自己的存在性
 */
func (c *ZKClient) Exists() (exists bool, err error) {
	c.lock.RLock()
	defer c.lock.RUnlock()
	if c.closed {
		return false, errors.New("Closed zk instance ")
	}
	path := filepath.Join(c.path, c.name)
	if strings.Contains(path, `\`) {
		path = strings.Replace(path, `\`, `/`, -1)
	}
	exists, _, err = c.conn.Exists(path)
	if err != nil {
		return false, err
	}
	//c.setVersion(stat.Version)
	return exists, nil
}

/**
*	更新节点的数据
 */
func (c *ZKClient) Update(value []byte) (err error) {
	c.lock.RLock()
	defer c.lock.RUnlock()
	if c.closed {
		return errors.New("Closed zk instance ")
	}
	err = c.ensurePath(c.path) //确保使用的目录存在
	if nil != err {
		return err
	}
	dest := filepath.Join(c.path, c.name)
	if strings.Contains(dest, `\`) {
		dest = strings.Replace(dest, `\`, `/`, -1)
	}
	fmt.Println("version1:", c.getVersion())
	stat, err := c.conn.Set(dest, value, c.getVersion())
	if err != nil {
		return err
	}
	fmt.Println("version2:", stat.Version)
	c.setVersion(stat.Version)
	return nil
}

/**
*	返回节点列表
 */
func (c *ZKClient) Nodes() (nodeValues map[string][]byte, err error) {
	c.lock.RLock()
	defer c.lock.RUnlock()

	if c.closed {
		return nil, errors.New("Closed zk instance ")
	}
	nodes, _, err := c.conn.Children(c.path)
	if nil != err {
		return nil, err
	}
	//c.setVersion(stat.Version)
	nodeValues, err = c.getNodeValues(nodes)
	if nil != err {
		return nil, err
	}
	return nodeValues, nil
}

/**
*	Mirror 指定路径的snapshots chan, 每个snapshot如 map[string][]byte{"ip:port":"value","ip:port":"value"}
*	有问题，改完后坚挺不到子集
 */
func (c *ZKClient) Mirror() (snapshots chan map[string][]byte, errors chan error) {
	c.lock.RLock()
	defer c.lock.RUnlock()
	snapshots = make(chan map[string][]byte)
	errors = make(chan error)
	go func() {
		for {
			nodes, _, events, err := c.conn.ChildrenW(c.path) // 变化后，获取所有节点名
			if nil != err {
				errors <- err
				continue
			}
			snapshot, err := c.getNodeValues(nodes) // 获取节点的值
			if nil != err {
				errors <- err
				continue
			}
			c.snapshot = snapshot
			snapshots <- snapshot
			event := <-events
			if nil != event.Err {
				errors <- err
				return
			}
		}
	}()
	return snapshots, errors
}

// nodeMirror 监听指定节点的变化
func (c *ZKClient) nodeMirror(name string) (snapshots chan []byte, errors chan error) {
	snapshots = make(chan []byte)
	errors = make(chan error)
	go func() {
		for {
			path := filepath.Join(c.path, name)
			if strings.Contains(path, `\`) {
				path = strings.Replace(path, `\`, `/`, -1)
			}
			node, _, events, err := c.conn.GetW(path)
			if nil != err {
				errors <- err
				continue
			}
			//c.setVersion(stat.Version)
			snapshots <- node
			event := <-events
			if nil != event.Err {
				errors <- err
				continue
			}
		}
	}()
	return snapshots, errors
}

/**
*	关闭连接
 */
func (c *ZKClient) Close() (err error) {
	c.lock.Lock()
	defer c.lock.Unlock()
	if c.closed {
		return errors.New("Closed zk instance ")
	}
	dest := filepath.Join(c.path, c.name)
	if strings.Contains(dest, `\`) {
		dest = strings.Replace(dest, `\`, `/`, -1)
	}
	err = c.conn.Delete(dest, c.getVersion())
	if nil != err {
		return err
	}
	defer c.conn.Close()
	c.closed = true
	return nil
}

/**
*	获取节点信息
 */
func (c *ZKClient) getNodeValues(nodes []string) (nodeValues map[string][]byte, err error) {
	nodeValues = make(map[string][]byte)
	for _, node := range nodes {
		path := filepath.Join(c.path, node)
		if strings.Contains(path, `\`) {
			path = strings.Replace(path, `\`, `/`, -1)
		}
		bytes, _, err := c.conn.Get(path)
		if nil != err {
			return nil, err
		}
		//c.setVersion(stat.Version)
		nodeValues[node] = bytes
	}
	return nodeValues, nil
}

/**
*	ensurePath path不存在时创建
*	param:path 路径
 */
func (c *ZKClient) ensurePath(path string) (err error) {
	slash := "/"
	tmpPath := slash
	directories := strings.Split(path, slash)
	for _, directory := range directories {
		tmpPath = filepath.Join(tmpPath, directory)
		if strings.Contains(tmpPath, `\`) {
			tmpPath = strings.Replace(tmpPath, `\`, "/", -1)
		}
		exists, _, err := c.conn.Exists(tmpPath)
		if nil != err {
			return err
		}
		//c.setVersion(stat.Version)
		if !exists {
			flags := int32(0)
			acl := zk.WorldACL(zk.PermAll)
			_, err = c.conn.Create(tmpPath, nil, flags, acl)
		}
	}
	return err
}

func (c *ZKClient) getVersion() int32 {
	c.versionLock.RLock()
	defer c.versionLock.RUnlock()
	return c.version
}
func (c *ZKClient) setVersion(version int32) {
	c.versionLock.Lock()
	defer c.versionLock.Unlock()
	c.version = version
}

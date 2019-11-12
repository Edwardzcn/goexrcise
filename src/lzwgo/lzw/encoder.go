package lzw

// 整体思路

import (
	"errors"
	"fmt"
	"io/ioutil"
	"path"
)

// Encoder ASCII
// 字典最大容载256,初始128(ASCII表)
// Encoder is a struct
type Encoder struct {
	readPath  string
	writePath string
	befComp   []byte
	aftComp   []byte
	preString []byte
	// dict      map[byte]byte
	capacity int

	root *TreeNode
}

// 在Tree上不断找延伸节点,初始应该全为空

type TreeNode struct {
	// 对应字典的索引值
	index int
	// 指针数组
	// 或者字典
	childNode [128]*TreeNode
	// childNode map[int]*TreeNode
}

func (c *Encoder) Init(read_name string, write_name string) {
	c.readPath = read_name
	c.writePath = write_name
	c.root = new(TreeNode)
	c.root.index = -1
	c.capacity = 128
	// c.root.childNode = make(map[int]*TreeNode, 128)
	// fmt.Println(c.root.childNode)
	// 初始ASCII码表存上
	// 初始化根结点指针
	for i := 0; i <= 127; i++ {
		// c.dict[byte(i)] = byte(i)
		c.root.childNode[i] = new(TreeNode)
		c.root.childNode[i].index = i
		// c.root.childNode[i]
		// fmt.Println(c.root.childNchildNode[addchar]
	}
}

func (c *Encoder) ReadFile() (err error) {
	// 注意接收者参数是指针,要做修改
	// fmt.Println("#", c.readPath, "#")
	c.befComp, err = ioutil.ReadFile(c.readPath)
	if err != nil {
		return err
	}
	fmt.Printf("Read file from %s to buffer ok!\n", path.Base(c.readPath))
	return nil
}

func (c Encoder) WriteFile() (err error) {
	err = ioutil.WriteFile(c.writePath, c.aftComp, 0644)
	if err != nil {
		return err
	}
	fmt.Printf("Write file to %s ok!\n", path.Base(c.writePath))
	return nil
}

func (c *Encoder) Circle() {
	fmt.Println("Start Encode Circle!")
	for key, char := range c.befComp {
		outIndex, addflag, _ := c.Digest(char)
		if addflag == true {
			c.preString = []byte{}
			c.preString = append(c.preString, char)
			c.aftComp = append(c.aftComp, outIndex)
			if key == len(c.befComp)-1 {
				c.aftComp = append(c.aftComp, char)
			}
		} else {
			// 在末尾的时候特判
			if key == len(c.befComp)-1 {
				// c.preString = []byte{}
				c.aftComp = append(c.aftComp, outIndex)
			}
		}
	}
	// fmt.Println("Test print:")
	// fmt.Println("before compression: ", c.befComp)
	// fmt.Println("after compression: ", c.aftComp)
}

func (c *Encoder) Digest(scanChar byte) (outIndex byte, flag bool, err error) {
	// 将Pre字符串和扫描的字符组合成为checkString
	checkString := append(c.preString, scanChar)
	// 这列可以换成 lastroot 有待改进
	nowPtr, isFind := c.FindTreeptr(checkString, c.root)
	if isFind == true {
		// 变更pre continue
		// 不做输出
		c.preString = checkString
		return byte(nowPtr.index), false, nil
	} else {
		// 加入新结点
		outIndex, err := c.AddTreeptr(nowPtr, scanChar)
		// 先输出
		return outIndex, true, err
	}

}

func (c *Encoder) FindTreeptr(leftString []byte, treeRoot *TreeNode) (addptr *TreeNode, flag bool) {
	tempChar := leftString[0]
	if treeRoot.childNode[tempChar] != nil {
		// 已经有节点
		if len(leftString) == 1 {
			// 返回本结点
			return treeRoot.childNode[tempChar], true
		} else {
			// 递归
			return c.FindTreeptr(leftString[1:], treeRoot.childNode[tempChar])
		}
	} else {
		// 未有节点
		// 返回父亲结点（输出过程需要知道index）
		return treeRoot, false
	}
}

func (c *Encoder) AddTreeptr(faptr *TreeNode, addchar byte) (outIndex byte, err error) {
	outIndex = byte(faptr.index)
	if c.capacity >= 256 {
		return outIndex, errors.New("Dictionary Full!")
	}
	// 由于获取父结点index需要所以这么操作
	faptr.childNode[addchar] = new(TreeNode)
	faptr.childNode[addchar].index = int(c.capacity)
	c.capacity++
	// c.dict[newIndex] =
	return outIndex, nil
}

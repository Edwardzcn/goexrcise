package lzw

// 整体思路

import (
	"fmt"
	"io/ioutil"
	"path"
)

// Decoder ASCII
// 字典最大容载256,初始128(ASCII表)
// Decoder is a struct
type Decoder struct {
	readPath    string
	writePath   string
	befComp     []byte
	aftComp     []byte
	storeString []byte
	dict        map[byte]string
}

// Treenode 定义过了

func (c *Decoder) Init(read_name string, write_name string) {
	c.readPath = read_name
	c.writePath = write_name
	c.dict = make(map[byte]string, 256)
	// c.root.childNode = make(map[int]*TreeNode, 128)
	// fmt.Println(c.root.childNode)
	// 初始ASCII码表存上
	// 初始化根结点指针
	for i := 0; i <= 127; i++ {
		c.dict[byte(i)] = string(byte(i))
	}
	// 打印初始化字典
	fmt.Print(c.dict)
}

func (c *Decoder) ReadFile() (err error) {
	// 注意接收者参数是指针,要做修改
	// fmt.Println("#", c.readPath, "#")
	c.befComp, err = ioutil.ReadFile(c.readPath)
	if err != nil {
		return err
	}
	fmt.Printf("Read file from %s to buffer ok!\n", path.Base(c.readPath))
	return nil
}

func (c Decoder) WriteFile() (err error) {
	err = ioutil.WriteFile(c.writePath, c.aftComp, 0644)
	if err != nil {
		return err
	}
	fmt.Printf("Write file to %s ok!\n", path.Base(c.writePath))
	return nil
}

//BytesCombine 多个[]byte数组合并成一个[]byte
// func BytesCombine(pBytes ...[]byte) []byte {
//     return bytes.Join(pBytes, []byte(""))
// }

func (c *Decoder) Circle() {
	fmt.Println("Start Decoder Circle!")
	for _, value := range c.befComp {
		if len(c.storeString) == 0 {
			c.aftComp = append(c.aftComp, value)
			// 第一个 特判 变更storeString
			c.storeString = append(c.storeString, value)
			continue
		}
		// 其他情况
		// (错误)写： aftComp是byte切片   storeString 也是byte切片
		// 先从字典获取现在值对应的字符串
		lookUpString := c.dict[value]
		// 加入输出
		c.aftComp = append(c.aftComp, []byte(lookUpString)...)

		// XX? 变成XXX
		// 把查到的字典加进来
		if flag := c.CheckDicFull(); flag == true {
			continue
		} else {

			c.StoreDic(c.storeString, lookUpString)
			// 将storeString 转化为当前 XX？
			c.storeString = []byte(lookUpString)
		}

	}
}

func (c *Decoder) CheckDicFull() (flag bool) {
	if len(c.dict) == 256 {
		return true
	}
	return false
}

func (c *Decoder) StoreDic(original []byte, prefix string) {
	// 原始+这次扫描的第一个byte（）
	newcontent := append(original, prefix[0])
	length := len(c.dict)
	fmt.Println(length, string(newcontent))
	c.dict[byte(length)] = string(newcontent)
}

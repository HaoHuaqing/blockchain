package main

import (
	"fmt"
	"strconv"
	"bytes"
	"crypto/sha256"
	"time"
	"os"
)
/*
区块结构
 */

type Block struct {
	Timestamp	 int64	//时间戳
	Data 		 []byte //当前区块存放的信息
	PrevBlockHash []byte //上一个区块的Hash
	Hash 		 []byte //当前区块的Hash
}

func(this *Block) SetHash(){
	//将本区块的TimeStemp+Data+PreBlockHash--->Hash
	timestamp := []byte(strconv.FormatInt(this.Timestamp,10))
	//将三个二进制拼接
	headers := bytes.Join([][]byte{this.PrevBlockHash, this.Data, timestamp}, []byte{})
	//拼接之后的headers进行SHA256加密
	hash := sha256.Sum256(headers)
	this.Hash = hash[:]
}
/*
新建一个区块的API
 */
 func NewBlock(data string, prevBlockHash []byte) *Block {
 	//生成一个区块
 	block := Block{}
 	//给当前区块赋值
 	block.Timestamp = time.Now().Unix()
 	block.Data = []byte(data)
 	block.PrevBlockHash = prevBlockHash
 	//给当前区块进行hash加密
 	block.SetHash()
 	return &block
 }

 /*
 定义一个区块链结构
  */
  type BlockChain struct{
  	Blocks []*Block //有序的区块
  }
  func (this *BlockChain) AddBlock(data string){
  	//得到前区块
  	prevBlock := this.Blocks[len(this.Blocks)-1]
  	//根据data创建一个新的区块
  	newBlock := NewBlock(data, prevBlock.Hash)
  	//根据前区块和新区块，添加到区块链blocks中
  	this.Blocks = append(this.Blocks, newBlock)
  }
  //区块链 = 创世块-->区块-->区块

  //新建一个创世块
  func NewGenesisBlock() *Block {
  	genesisBlock := Block{}
  	genesisBlock.Data = []byte("Genesis block")
  	genesisBlock.PrevBlockHash = []byte{}

  	return &genesisBlock
  }
//新建一个区块链
  func NewBlockChain() *BlockChain{
  	return &BlockChain{[]*Block{NewGenesisBlock()}}
  }

func main() {
	//创建一个区块链bc
	bc := NewBlockChain()

	//用户输入指令1,2，other
	var cmd string
	for {
		fmt.Println("按‘1’添加一条信息到区块链中")
		fmt.Println("按‘2’遍历当前区块链中有哪些区块信息")
		fmt.Println("按其他按键退出")
		fmt.Scanf("%s\n", &cmd)

		switch cmd {
		case "1":
			input := make([]byte, 1024)
			//添加一个区块
			fmt.Println("请输入区块链的行为数据")
			os.Stdin.Read(input)
			bc.AddBlock(string(input))
		case "2":
			for i, block := range bc.Blocks{
				fmt.Println("===================")
				fmt.Println("第", i, "个区块的信息：")
				fmt.Printf("PrevHash: %x\n", block.PrevBlockHash)
				fmt.Printf("Data: %s\n", block.Data)
				fmt.Printf("Hash: %x\n", block.Hash)
				fmt.Println("===================")
			}
		default:
			fmt.Println("您已退出")
			return
		}
	}
}
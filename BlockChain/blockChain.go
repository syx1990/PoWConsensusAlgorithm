package BlockChain

import (
	"PoW/Block"
	"fmt"
)

type Node struct {
	NextNode *Node //指针域
	Data *Block.Block //数据域
}

var HeadNode *Node

//创建头节点，保存创世区块
func CreatHeadNode(data *Block.Block) *Node  {
	var headNode *Node = new(Node)
	headNode.NextNode = nil
	headNode.Data = data
	return headNode
}

//挖矿成功以后添加区块
func AddNode(data *Block.Block,prefNode *Node) *Node {
	var newNode *Node = new(Node)
	newNode.Data = data
	newNode.NextNode = nil
	prefNode.NextNode = newNode
	return newNode
}

func ShowBlockChain(headNode *Node)  {
	var node *Node = headNode
	fmt.Print(node.Data)
	for {
		if node.NextNode != nil {
			node = node.NextNode
			fmt.Print("----",node.Data,"----")
		}else if node.NextNode == nil{
			break
		}
	}
}



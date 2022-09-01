package main


import(
	"../Block"
	"fmt"
	"strings"
	"time"
	"../BlockChain"
	"strconv"
)
	
//工作量证明函数 进行哈希碰撞
//传指针 保证操作对象的同一性

func pow(diff int, block *Block.Block) string{
	//不停的去挖
	count := 0
	for{
		hash := Block.GenerateHashcode(*block)
		//fmt.Println(hash)
		//判断前导零个数
		if strings.HasPrefix(hash,strings.Repeat("0",diff)){
			fmt.Println("经过"+strconv.Itoa(count)+"次后挖矿成功")
			return hash
		} else {
			block.Nonce++
			count++
		}
	}
}
	
func NextBlock(data string, oldblock Block.Block)Block.Block {
	var nextblock Block.Block
	nextblock.PreHash = oldblock.HashCode
	nextblock.TimeStamp = time.Now().String()
	nextblock.Index = 2 //区块高度
	nextblock.Diff = 4
	nextblock.Data = data
	nextblock.Nonce = 0
	nextblock.HashCode = pow(nextblock.Diff,&nextblock)

	return nextblock
}


func main(){
	firstblock := Block.CreateFirstBlock("创世区块")
	fmt.Println(firstblock)
	headernode := BlockChain.CreatHeadernode(&firstblock)
	secondblock := NextBlock("下一区块",firstblock)
	fmt.Println("下一区块哈希是"+secondblock.HashCode)
	nextnode := BlockChain.Addnode(headernode,&secondblock)
	fmt.Println(nextnode.Data)
	PrintChain(headernode)
}

func PrintChain(node *BlockChain.Node){
	var tempnode *BlockChain.Node= new(BlockChain.Node)
	if node.Data != nil{
		fmt.Println(node.Data)
	}
	tempnode = node.Nextnode
	for(tempnode != nil){
		fmt.Println(tempnode.Data)
		tempnode = tempnode.Nextnode
	}
}
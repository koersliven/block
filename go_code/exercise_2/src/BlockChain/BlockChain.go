package BlockChain
import "../Block"

type Node struct{
	Nextnode *Node//指针域

	Data *Block.Block//数据域

}
//创建头节点
func CreatHeadernode(block *Block.Block) *Node{
	var headernode *Node =new(Node)
	headernode.Nextnode = nil
	headernode.Data = block

	return headernode
}


//添加节点

func Addnode(node *Node,newblock *Block.Block) *Node{
	var nextnode *Node =new(Node)
	//指针域为空 数据域添加新的block	
	nextnode.Nextnode = nil
	nextnode.Data =newblock	

	node.Nextnode = nextnode
	return nextnode
	
}
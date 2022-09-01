package Block

import(
	"time"
	"strconv"
	"crypto/sha256"
	"encoding/hex"
)
type Block struct{
	PreHash string
	//last block's HashCode
	HashCode string
	//this block's Hashcode
	TimeStamp string
	//time

	Diff int//the diffcultiy of block

	Index int//the height of the chain
	Nonce int//random number
	Data string //the information of purchase
}

func CreateFirstBlock (data string) Block{
	var firstblock Block
	firstblock.PreHash = "0"
	firstblock.TimeStamp = time.Now().String()
	firstblock.Diff = 4
	firstblock.Data = data
	firstblock.Index = 1
	firstblock.Nonce = 0

	firstblock.HashCode = GenerateHashcode(firstblock)

	return firstblock


}

func GenerateHashcode(block Block) string {
	var hashdata = strconv.Itoa(block.Index) + strconv.Itoa(block.Nonce) + 
	block.TimeStamp + strconv.Itoa(block.Diff) + block.Data

	var sha = sha256.New()
	sha.Write([]byte(hashdata))
	hashed := sha.Sum(nil)

	return hex.EncodeToString(hashed)
}

//生成下一个区块函数


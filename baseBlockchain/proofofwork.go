package baseBlockchain

// 计算块工作证明
import (
	"blockchain/tools"
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
)

// 最大记录数
var (
	maxNonce = math.MaxInt64
)

// 目标比特
const targetBits = 24

// 工作证明
type ProofOfWork struct {
	block  *Block
	target *big.Int
}

// 创建并返回一个工作证明
func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))

	pow := &ProofOfWork{b, target}

	return pow
}

// 准备数据
func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.Data,
			tools.IntToHex(pow.block.Timestamp),
			tools.IntToHex(int64(targetBits)),
			tools.IntToHex(int64(nonce)),
		},
		[]byte{},
	)

	return data
}

// 执行一个工作证明
func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	fmt.Printf("区块包含 \"%s\"\n", pow.block.Data)
	for nonce < maxNonce {
		data := pow.prepareData(nonce)

		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.target) == -1 {
			break
		} else {
			nonce++
		}
	}
	fmt.Print("\n\n")

	return nonce, hash[:]
}

// 验证工作证明是否有效
func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int

	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	isValid := hashInt.Cmp(pow.target) == -1

	return isValid
}

package utils

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"golang.org/x/crypto/sha3"

	"erinaceus_data_feeds/utils/hex"
)

const (
	// FastN is a shorter N parameter for testing
	FastN = 2
	// FastP is a shorter P parameter for testing
	FastP = 1
)

type ScryptParams struct{ N, P int }

var FastScryptParams = ScryptParams{N: FastN, P: FastP}

// EVMWordByteLen the length of an EVM Word Byte
const EVMWordByteLen = 32

// ZeroAddress is an address of all zeroes, otherwise in Ethereum as
// 0x0000000000000000000000000000000000000000
var ZeroAddress = common.Address{}

// EmptyHash is a hash of all zeroes, otherwise in Ethereum as
// 0x0000000000000000000000000000000000000000000000000000000000000000
var EmptyHash = common.Hash{}

func IsAllZeros(bytes [32]byte) bool {
	for _, b := range bytes {
		if b != 0 {
			return false
		}
	}
	return true
}

func RandomAddress() common.Address {
	b := make([]byte, 20)
	_, _ = rand.Read(b) // Assignment for errcheck. Only used in tests so we can ignore.
	return common.BytesToAddress(b)
}

// IsEmptyAddress checks that the address is empty, synonymous with the zero
// account/address. No logs can come from this address, as there is no contract
// present there.
//
// See https://stackoverflow.com/questions/48219716/what-is-address0-in-solidity
// for the more info on the zero address.
func IsEmptyAddress(addr common.Address) bool {
	return addr == ZeroAddress
}

func RandomBytes32() (r [32]byte) {
	b := make([]byte, 32)
	_, _ = rand.Read(b[:]) // Assignment for errcheck. Only used in tests so we can ignore.
	copy(r[:], b)
	return
}

func Bytes32ToSlice(a [32]byte) (r []byte) {
	r = append(r, a[:]...)
	return
}

// Uint256ToBytes is x represented as the bytes of a uint256
func Uint256ToBytes(x *big.Int) (uint256 []byte, err error) {
	if x.Cmp(MaxUint256) > 0 {
		return nil, fmt.Errorf("too large to convert to uint256")
	}
	uint256 = common.LeftPadBytes(x.Bytes(), EVMWordByteLen)
	if x.Cmp(big.NewInt(0).SetBytes(uint256)) != 0 {
		panic("failed to round-trip uint256 back to source big.Int")
	}
	return uint256, err
}

// NewHash return random Keccak256
func NewHash() common.Hash {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return common.BytesToHash(b)
}

// PadByteToHash returns a hash with zeros padded on the left of the given byte.
func PadByteToHash(b byte) common.Hash {
	var h [32]byte
	h[31] = b
	return h
}

// Uint256ToBytes32 returns the bytes32 encoding of the big int provided
func Uint256ToBytes32(n *big.Int) []byte {
	if n.BitLen() > 256 {
		panic("vrf.uint256ToBytes32: too big to marshal to uint256")
	}
	return common.LeftPadBytes(n.Bytes(), 32)
}

// MustHash returns the keccak256 hash, or panics on failure.
func MustHash(in string) common.Hash {
	out, err := Keccak256([]byte(in))
	if err != nil {
		panic(err)
	}
	return common.BytesToHash(out)
}

// HexToUint256 returns the uint256 represented by s, or an error if it doesn't
// represent one.
func HexToUint256(s string) (*big.Int, error) {
	rawNum, err := hexutil.Decode(s)
	if err != nil {
		return nil, fmt.Errorf("error while parsing %s as hex: %w", s, err)
	}
	rv := big.NewInt(0).SetBytes(rawNum) // can't be negative number
	if err := CheckUint256(rv); err != nil {
		return nil, err
	}
	return rv, nil
}

var zero = big.NewInt(0)

// CheckUint256 returns an error if n is out of bounds for a uint256
func CheckUint256(n *big.Int) error {
	if n.Cmp(zero) < 0 || n.Cmp(MaxUint256) >= 0 {
		return fmt.Errorf("number out of range for uint256")
	}
	return nil
}

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// Keccak256 is a simplified interface for the legacy SHA3 implementation that
// Ethereum uses.
func Keccak256(in []byte) ([]byte, error) {
	hash := sha3.NewLegacyKeccak256()
	_, err := hash.Write(in)
	return hash.Sum(nil), err
}

func Keccak256Fixed(in []byte) [32]byte {
	hash := sha3.NewLegacyKeccak256()
	// Note this Keccak256 cannot error https://github.com/golang/crypto/blob/master/sha3/sha3.go#L126
	// if we start supporting hashing algos which do, we can change this API to include an error.
	hash.Write(in)
	var h [32]byte
	copy(h[:], hash.Sum(nil))
	return h
}

// EIP55CapitalizedAddress returns true iff possibleAddressString has the correct
// capitalization for an Ethereum address, per EIP 55
func EIP55CapitalizedAddress(possibleAddressString string) bool {
	possibleAddressString = hex.EnsurePrefix(possibleAddressString)
	EIP55Capitalized := common.HexToAddress(possibleAddressString).Hex()
	return possibleAddressString == EIP55Capitalized
}

// ParseEthereumAddress returns addressString as a go-ethereum Address, or an
// error if it's invalid, e.g. if EIP 55 capitalization check fails
func ParseEthereumAddress(addressString string) (common.Address, error) {
	if !common.IsHexAddress(addressString) {
		return common.Address{}, fmt.Errorf(
			"not a valid Ethereum address: %s", addressString)
	}
	address := common.HexToAddress(addressString)
	if !EIP55CapitalizedAddress(addressString) {
		return common.Address{}, fmt.Errorf(
			"%s treated as Ethereum address, but it has an invalid capitalization! "+
				"The correctly-capitalized address would be %s, but "+
				"check carefully before copying and pasting! ",
			addressString, address.Hex())
	}
	return address, nil
}

func LoadEnv(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" || strings.HasPrefix(line, "#") { // Ignore empty lines and comments
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			log.Printf("Ignoring malformed line: %s\n", line)
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		if err := os.Setenv(key, value); err != nil {
			return err
		}
	}

	return scanner.Err()
}

// HasQuotes checks if the first and last characters are either " or '.
func HasQuotes(input []byte) bool {
	return len(input) >= 2 &&
		((input[0] == '"' && input[len(input)-1] == '"') ||
			(input[0] == '\'' && input[len(input)-1] == '\''))
}

// TrimQuotes removes the first and last character if they are both either
// " or ', otherwise it is a noop.
func TrimQuotes(input []byte) []byte {
	if HasQuotes(input) {
		return input[1 : len(input)-1]
	}
	return input
}

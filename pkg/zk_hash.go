package pkg

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark-crypto/hash"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
	"github.com/consensys/gnark/std/hash/mimc"
	"math/big"
	//"io"
)

type Circuit struct {
	Source frontend.Variable
	Hash     frontend.Variable `gnark:",public"`
}

type MyRandReader struct{
	Random  []byte
}

func (m MyRandReader) Read(p []byte) (n int, err error) {
	genLen := len(p)

	zeroLen := genLen - len(m.Random)

	slice := make([]byte, zeroLen)
	if genLen > len(m.Random){
		p = append(m.Random, slice...)

	} else {
		p = m.Random[:genLen]
	}
	rand.Reader = MyRandReader{}
	return len(p), nil
}

type MyPolyReader struct{
}

func (m MyPolyReader) Read(p []byte) (n int, err error) {
	length := len(p)
	for i := 0; i < length; i++ {
		p[i] = byte(i + 1)
	}
	return len(p), nil
}

func (circuit *Circuit) Define(api frontend.API) error {
	// hash function
	mimc, _ := mimc.NewMiMC(api)
	mimc.Write(circuit.Source)
	api.AssertIsEqual(circuit.Hash, mimc.Sum())

	return nil
}

func HashCalc(source string) string{
	hi,_ := big.NewInt(0).SetString(source, 10)

	h := hash.MIMC_BN254.New()
	h.Write(hi.Bytes())
	rd := h.Sum(nil)
	r1 := big.NewInt(0).SetBytes(rd).String()

	return r1
}

//生成电路、proof、verifyKey
func GenerateProve(source string, random []byte)  ([]byte, []byte){
	rand.Reader = MyPolyReader{}
	h := sha256.New()
	h.Write([]byte(source))
	r := h.Sum(nil)
	source = hex.EncodeToString(r)
	n := new(big.Int)
	n, ok := n.SetString(source, 16)
	if !ok {
		panic("xxx")
	}
	source = n.String()

	// 外部系统生成零知识证明电路，编译电路
	var circuit Circuit
	ccs, err := frontend.Compile(ecc.BN254, r1cs.NewBuilder, &circuit)
	if err != nil {
		panic(err)
	}
	// 根据编译电路生成proofKey
	pk, vk, err := groth16.Setup(ccs)
	if err != nil {
		panic(err)
	}

	assignment :=  Circuit{
		Source: source,
		Hash:   HashCalc(source),
	}

	witness, err := frontend.NewWitness(&assignment, ecc.BN254)
	if err != nil {
		panic(err)
	}

	rand.Reader = MyRandReader{Random:random}
	proof, err := groth16.Prove(ccs, pk, witness)
	if err != nil {
		panic(err)
	}
	var proofBuffer bytes.Buffer
	proofBuffer.Reset()
	proof.WriteRawTo(&proofBuffer)
	proofBuffer.Bytes()

	var vkBuffer bytes.Buffer
	vkBuffer.Reset()
	vk.WriteRawTo(&vkBuffer)
	vkBuffer.Bytes()

	//proof, verifyKey
	return proofBuffer.Bytes(), vkBuffer.Bytes()
}

func VerifyProof(value string, verifyKey []byte, publicWitness []byte) (bool, error) {
	assignment1 :=  Circuit{
		Hash:     value,
	}
	publicWitness1,err := frontend.NewWitness(&assignment1, ecc.BN254, frontend.PublicOnly())
	if err != nil {
		return false, err
	}

	proof := groth16.NewProof(ecc.BN254)
	proof.ReadFrom(bytes.NewBuffer(publicWitness))

	vk := groth16.NewVerifyingKey(ecc.BN254)
	vk.ReadFrom(bytes.NewBuffer(verifyKey))

	err = groth16.Verify(proof, vk, publicWitness1)
	if err != nil {
		return false, err
	}
	return true, nil
}

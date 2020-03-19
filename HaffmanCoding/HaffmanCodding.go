package main

import (
	"bufio"
	"fmt"
	"github.com/AlgoAndStruct/myContainers/binaryTree"
	"github.com/AlgoAndStruct/myContainers/heap"
	"os"
)

type letter struct {
	frequency int
	ch rune
}

type EncryptedInfo struct {
	frequencyTable map[rune]int
	letterTable map[rune][]rune
	str []rune
}

func ComposeFrequencyTable(unencrypted string) map[rune]int {
	frequencyTable := make(map[rune]int)
	for _, ch := range unencrypted {
		frequencyTable[rune(ch)]++
	}
	return frequencyTable
}

func MakeParent(leftChild, rightChild *binaryTree.Element) *binaryTree.Element {
	parent := binaryTree.MakeParent(leftChild, rightChild)
	parent.Value = letter{
		leftChild.Value.(letter).frequency + rightChild.Value.(letter).frequency,
		'0',
	}
	return parent
}

func isLetterElement(elem *binaryTree.Element) bool {
	return elem.Value.(letter).ch != '0'
}

func FillLetterTable(elem *binaryTree.Element, code []rune, letterTable map[rune][]rune) []rune {
	if isLetterElement(elem) {
		letterTable[elem.Value.(letter).ch] = append(letterTable[elem.Value.(letter).ch], code...)
		if len(code) > 0 {
			code = code[:len(code)-1]
		}
	}
	if elem.ChildLeft != nil {
		code = append(code, rune('0'))
		code = FillLetterTable(elem.ChildLeft, code, letterTable)
	}
	if elem.ChildRight != nil {
		code = append(code, rune('1'))
		code = FillLetterTable(elem.ChildRight, code, letterTable)
	}
	return code
}

func ComposeEncryptedInfo(encryp *EncryptedInfo, HaffmanTree *binaryTree.Tree, unencrypted string) *EncryptedInfo {
	var code []rune

	if !HaffmanTree.Root.HasChilds() {
		code = append(code, rune('0'))
	}
	encryp.letterTable = make(map[rune][]rune)
	FillLetterTable(HaffmanTree.Root, code, encryp.letterTable)
	for _, letter := range unencrypted {
		encryp.str = append(encryp.str, encryp.letterTable[rune(letter)]...)
	}
	return encryp
}

func HaffmanAlgorithm(unencrypted string) *EncryptedInfo {
	var encryp EncryptedInfo
	var leftChild, rightChild *binaryTree.Element
	var freqPrQueue heap.MinHeap

	freqPrQueue.Init()
	encryp.frequencyTable = ComposeFrequencyTable(unencrypted)
	for ch, frequency := range encryp.frequencyTable { // make a priority queue of tree`s elements
		freqPrQueue.Insert(frequency, binaryTree.MakeElement(letter{frequency, ch}))
	}
	if freqPrQueue.Size() == 0 {
		return nil
	}

	for freqPrQueue.Size() != 1 { // make a Haffman tree
		if ok := freqPrQueue.ExtractMin(); ok != nil { leftChild = ok.(*binaryTree.Element) }
		if ok := freqPrQueue.ExtractMin(); ok != nil { rightChild = ok.(*binaryTree.Element) }
		parent := MakeParent(leftChild, rightChild)
		freqPrQueue.Insert(parent.Value.(letter).frequency, parent)
	}
	HaffmanTree := binaryTree.MakeRoot(freqPrQueue.ExtractMin().(*binaryTree.Element))
	return ComposeEncryptedInfo(&encryp, HaffmanTree, unencrypted)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	encryptedInfo := HaffmanAlgorithm(scanner.Text())
	if encryptedInfo == nil {
		return
	}

	fmt.Println(len(encryptedInfo.frequencyTable), len(encryptedInfo.str))
	for k, v := range encryptedInfo.letterTable {
		fmt.Printf("%c: %s\n", k, string(v))
	}
	fmt.Println(string(encryptedInfo.str))

}

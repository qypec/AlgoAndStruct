/*
	По данной непустой строке s длины не более 10^4 , состоящей из строчных букв латинского алфавита,
	постройте оптимальный беспрефиксный код. В первой строке выведите количество различных букв k,
	встречающихся в строке, и размер получившейся закодированной строки.
	В следующих k строках запишите коды букв в формате "letter: code".
	В последней строке выведите закодированную строку.
 */

package main

import (
	"bufio"
	"container/list"
	"fmt"
	"github.com/AlgoAndStruct/myContainers/binaryTree"
	"os"
)

// Custom method for List
// Adds an item to a list in sorted order
func PushSort(lst *list.List, value interface{}) *list.Element {
	newFrequency := value.(*binaryTree.Element).Value.(letter).frequency
	for e := lst.Front(); e != nil; e = e.Next() {
		if newFrequency <= e.Value.(*binaryTree.Element).Value.(letter).frequency {
			return lst.InsertBefore(value, e)
		}
	}
	return lst.PushBack(value)
}

type letter struct {
	frequency int
	ch rune
}

type EncryptedInfo struct {
	frequencyTable map[rune]int // <letter>: <number of letter in the string>
	letterTable map[rune][]rune // <letter>: <letter code>
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
	}
	if elem.ChildLeft != nil {
		code = append(code, rune('0'))
		code = FillLetterTable(elem.ChildLeft, code, letterTable)
	}
	if elem.ChildRight != nil {
		code = append(code, rune('1'))
		code = FillLetterTable(elem.ChildRight, code, letterTable)
	}
	if len(code) > 0 {
		code = code[:len(code)-1]
	}
	return code
}

func ComposeEncryptedInfo(encryp *EncryptedInfo, HaffmanTree *binaryTree.Tree, unencrypted string) *EncryptedInfo {
	var code []rune

	if !HaffmanTree.Root.HasChilds() { // kostil`
		code = append(code, rune('0'))
	}
	encryp.letterTable = make(map[rune][]rune)
	FillLetterTable(HaffmanTree.Root, code, encryp.letterTable)

	for _, letter := range unencrypted { // compose a encrypted string
		encryp.str = append(encryp.str, encryp.letterTable[rune(letter)]...)
	}
	return encryp
}

func HaffmanAlgorithm(unencrypted string) *EncryptedInfo {
	var encryp EncryptedInfo
	var leftChild, rightChild *binaryTree.Element

	sortLst := list.New()
	encryp.frequencyTable = ComposeFrequencyTable(unencrypted)
	for ch, frequency := range encryp.frequencyTable { // make a sorting list of tree`s elements
		PushSort(sortLst, binaryTree.MakeElement(letter{frequency, ch}))
	}
	if sortLst.Len() == 0 {
		return nil
	}

	for sortLst.Len() != 1 { // make a Haffman tree
		if ok := sortLst.Front(); ok != nil {
			leftChild = sortLst.Remove(sortLst.Front()).(*binaryTree.Element)
		}
		if ok := sortLst.Front(); ok != nil {
			rightChild = sortLst.Remove(sortLst.Front()).(*binaryTree.Element)
		}
		parent := MakeParent(leftChild, rightChild)
		PushSort(sortLst, parent)
	}
	HaffmanTree := binaryTree.MakeRoot(sortLst.Front().Value.(*binaryTree.Element))
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

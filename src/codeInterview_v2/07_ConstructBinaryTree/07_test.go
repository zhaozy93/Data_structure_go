package code_07

import (
	"log"
	"strings"
	"testing"
)

func TestSort(t *testing.T) {
	preOrder := "GDAFEMHZ"
	inOrder := "ADEFGHMZ"
	root := ConstructBinaryTreeByPreandInOrder(strings.Split(preOrder, ""), strings.Split(inOrder, ""))
	log.Println(root)
	log.Println(root.leftC)
	log.Println(root.leftC.leftC)
	log.Println(root.leftC.rightC)
	log.Println(root.leftC.rightC.leftC)
	log.Println(root.rightC)
	log.Println(root.rightC.leftC)
	log.Println(root.rightC.rightC)

}

package node

import "fmt"

type NodeType string

const (
	NODE_TYPE_DATA      NodeType = "DATA"
	NODE_TYPE_PROCEDURE NodeType = "PROCEDURE"
)

func GetNameAndType(fullName string) (NodeType, string) {
	if len(fullName) > 4 && fullName[:4] == "data" {
		dataName := fullName[5:]
		return NODE_TYPE_DATA, dataName
	} else if len(fullName) > 9 && fullName[:9] == "procedure" {
		name := fullName[10:]
		return NODE_TYPE_PROCEDURE, name
	}
	return NodeType(""), ""
}

func TypeAndStringKey(t NodeType, name string) string {
	return fmt.Sprintf("%s-%s", t, name)
}

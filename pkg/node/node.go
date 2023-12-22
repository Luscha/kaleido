package node

import "fmt"

type NodeType string

const (
	NODE_TYPE_DATA          NodeType = "DATA"
	NODE_TYPE_PROCEDURE     NodeType = "PROCEDURE"
	NODE_TYPE_SUB_PROCEDURE NodeType = "SUB_PROCEDURE"
	NODE_TYPE_ACTION        NodeType = "ACTION"
)

func GetNameAndType(fullName string) (NodeType, string) {
	if len(fullName) > 4 && fullName[:4] == "data" {
		dataName := fullName[5:]
		return NODE_TYPE_DATA, dataName
	} else if len(fullName) > 9 && fullName[:9] == "procedure" {
		name := fullName[10:]
		return NODE_TYPE_PROCEDURE, name
	} else if len(fullName) > 14 && fullName[:14] == "real_procedure" {
		name := fullName[15:]
		return NODE_TYPE_SUB_PROCEDURE, name
	} else if len(fullName) > 6 && fullName[:6] == "action" {
		name := fullName[7:]
		return NODE_TYPE_ACTION, name
	}
	return NodeType(""), ""
}

func TypeAndStringKey(t NodeType, name string) string {
	return fmt.Sprintf("%s-%s", t, name)
}

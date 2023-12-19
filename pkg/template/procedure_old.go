package template

import (
	"bytes"
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
	"sync"
	"text/template"

	"github.pitagora/pkg/datasource"
	"github.pitagora/pkg/node"
)

type Arguments map[string]any

func (a Arguments) HasArguments(nodeType node.NodeType, name string) bool {
	return true
}

var reg = regexp.MustCompile(`"\|{\|(.*?)\|}\|"`)
var regCurlOpen = regexp.MustCompile(`"{`)

func ArgumentPrefix(nodeType node.NodeType, name string) string {
	switch nodeType {
	case node.NODE_TYPE_DATA:
		return fmt.Sprintf("%s.%s", "data", name)
	case node.NODE_TYPE_PROCEDURE:
		return fmt.Sprintf("%s.%s", "procedure", name)
	}
	return ""
}

func Resolve(in any, args Arguments, prefix string, out any) error {
	stringIn, err := json.Marshal(in)
	if nil != err {
		return err
	}
	fmt.Println(string(stringIn))

	preparedArgs, err := PrepareArguments(args, prefix)
	if err != nil {
		return err
	}

	unescaped := reg.ReplaceAllString(string(stringIn), `|{|$1|}|`)
	fmt.Println(string(unescaped))
	funcMap := template.FuncMap{
		"nillable": func(val interface{}) string {
			if val == nil {
				return "null"
			}
			return fmt.Sprintf("%v", val)
		},
		"string": func(val interface{}) string {
			if val == nil {
				return "null"
			}
			return fmt.Sprintf("\"%v\"", val)
		},
	}

	tmpl, err := template.New("json").Delims("|{|", "|}|").Funcs(funcMap).Parse(unescaped)
	if err != nil {
		return err
	}

	var tpl bytes.Buffer
	err = tmpl.Execute(&tpl, preparedArgs)
	if err != nil {
		return err
	}

	fmt.Printf("Processed Data: %s\n", string(tpl.Bytes()))

	return json.Unmarshal(tpl.Bytes(), out)
}

func PrepareArguments(args Arguments, prefix string) (Arguments, error) {
	prepared := Arguments{}
	thisArg := &prepared

	for key, value := range args {
		if !strings.HasPrefix(key, prefix) {
			continue
		}

		var stringify string
		if _, ok := value.(string); ok {
			stringify = value.(string)
		} else {
			bytestring, err := json.Marshal(value)
			if err != nil {
				panic(err)
			}
			stringify = string(bytestring)
		}
		levels := strings.Split(strings.TrimPrefix(key, fmt.Sprintf("%s.", prefix)), ".")
		for i, level := range levels {
			if i == len(levels)-1 {
				(*thisArg)[level] = stringify
			} else if next, ok := (*thisArg)[level]; ok {
				a := (next).(Arguments)
				thisArg = &a
			} else {
				a := Arguments{}
				(*thisArg)[level] = a
				thisArg = &a
			}
		}

		thisArg = &prepared
	}
	return prepared, nil
}

func MergeArgumentsForData(ds datasource.DataSource, globalArguments Arguments, intermediateRes *sync.Map) (Arguments, error) {
	merged := Arguments{}
	for key, value := range globalArguments {
		merged[key] = value
	}

	for _, dep := range ds.Depends {
		specialName := node.TypeAndStringKey(node.GetNameAndType(dep.Value))
		res, ok := intermediateRes.Load(specialName)
		if !ok {
			return Arguments{}, fmt.Errorf("%s not found in results", specialName)
		}

		merged[fmt.Sprintf("%s%s", ArgumentPrefix(node.NODE_TYPE_DATA, ds.Name), dep.Template)] = string(res.([]byte))
	}
	return merged, nil
}

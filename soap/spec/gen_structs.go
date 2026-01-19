// gen_structs.go
// Command-line tool to parse an XML file and generate Go structs for all elements.
// Usage: go run gen_structs.go <xml_file> <namespace> <package>

package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

type Field struct {
	Name      string
	Type      string
	Tag       string
	Comment   string
	IsPointer bool
}

type StructDef struct {
	Name   string
	Fields []Field
}

func camelCase(s string) string {
	parts := regexp.MustCompile(`[_-]`).Split(s, -1)
	for i, p := range parts {
		if len(p) > 0 {
			parts[i] = strings.ToUpper(p[:1]) + p[1:]
		}
	}
	return strings.Join(parts, "")
}

func parseElement(e *xmlNode, structs map[string]*StructDef, ns, oldNS string) string {
	name := camelCase(e.XMLName.Local)
	if _, exists := structs[name]; exists {
		return name
	}
	fields := []Field{}
	// Attributes
	for _, attr := range e.Attrs {
		attrName := camelCase(attr.Name.Local)
		fields = append(fields, Field{
			Name:      attrName,
			Type:      "*string",
			Tag:       fmt.Sprintf("xml:\"%s,attr\"", attr.Name.Local),
			IsPointer: true,
		})
	}
	// Children
	for _, child := range e.Children {
		childType := "*string"
		if len(child.Children) > 0 || len(child.Attrs) > 0 {
			childType = "*" + camelCase(child.XMLName.Local)
			parseElement(child, structs, ns, oldNS)
		}
		fields = append(fields, Field{
			Name:      camelCase(child.XMLName.Local),
			Type:      childType,
			Tag:       fmt.Sprintf("xml:\"%s,omitempty\"", child.XMLName.Local),
			IsPointer: true,
		})
	}
	structs[name] = &StructDef{Name: name, Fields: fields}
	return name
}

type xmlNode struct {
	XMLName  xml.Name
	Attrs    []xml.Attr
	Children []*xmlNode
}

func buildTree(d *xml.Decoder, se xml.StartElement) (*xmlNode, error) {
	node := &xmlNode{XMLName: se.Name, Attrs: se.Attr}
	for {
		t, err := d.Token()
		if err != nil {
			return nil, err
		}
		switch tok := t.(type) {
		case xml.StartElement:
			child, err := buildTree(d, tok)
			if err != nil {
				return nil, err
			}
			node.Children = append(node.Children, child)
		case xml.EndElement:
			if tok.Name == se.Name {
				return node, nil
			}
		}
	}
}

func printStruct(s *StructDef, isRoot bool, ns, rootTag string) {
	fmt.Printf("type %s struct {\n", s.Name)
	if isRoot {
		fmt.Printf("\tXMLName xml.Name `xml:\"%s:%s\"`\n", ns, rootTag)
		fmt.Printf("\tXMLNamespace *string `xml:\"xmlns,attr\"`\n")
	}
	used := map[string]bool{}
	for _, f := range s.Fields {
		if used[f.Name] {
			continue
		}
		used[f.Name] = true
		fmt.Printf("\t%s %s `%s` %s\n", f.Name, f.Type, f.Tag, f.Comment)
	}
	fmt.Println("}")
}

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage: go run gen_structs.go <xml_file> <namespace> <package>")
		os.Exit(1)
	}
	xmlPath, ns, pkg := os.Args[1], os.Args[2], os.Args[3]
	data, err := ioutil.ReadFile(xmlPath)
	if err != nil {
		panic(err)
	}
	d := xml.NewDecoder(strings.NewReader(string(data)))
	var root *xmlNode
	for {
		t, err := d.Token()
		if err != nil {
			break
		}
		if se, ok := t.(xml.StartElement); ok {
			root, err = buildTree(d, se)
			if err != nil {
				panic(err)
			}
			break
		}
	}
	if root == nil {
		panic("No root element found")
	}
	rootTag := root.XMLName.Local
	structs := map[string]*StructDef{}
	parseElement(root, structs, ns, root.XMLName.Space)
	fmt.Printf("package %s\n\nimport (\n\t\"encoding/xml\"\n)\n\n", pkg)
	order := []string{}
	for name := range structs {
		order = append(order, name)
	}
	for i, name := range order {
		isRoot := i == len(order)-1
		printStruct(structs[name], isRoot, ns, rootTag)
		fmt.Println()
	}
}

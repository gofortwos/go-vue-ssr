package ast_from_api

import (
	"encoding/json"
	"errors"
	"github.com/bysir-zl/go-vue-ssr/internal/pkg/log"
	"github.com/bysir-zl/go-vue-ssr/internal/pkg/obj"
	"reflect"
)

type Node map[string]interface{}
type X interface {
}
type Program struct {
	Body []Node `json:"body"`
}

type ExpressionStatement struct {
	Expression Node `json:"expression"`
}

type Identifier struct {
	Name string `json:"name"`
}

type LogicalExpression struct {
	Operator string `json:"operator"` // && ||
	Left     Node   `json:"left"`
	Right    Node   `json:"right"`
}

type BinaryExpression struct {
	Operator string `json:"operator"` // +
	Left     Node   `json:"left"`
	Right    Node   `json:"right"`
}

type UnaryExpression struct {
	Operator string `json:"operator"` // !
	Prefix   bool   `json:"prefix"`
	Argument Node   `json:"argument"`
}

// 字面量, " " , 1, 1.1
type Literal struct {
	Value interface{} `json:"value"`
	Raw   string      `json:"raw"`
}

// 对象
type ObjectExpression struct {
	Properties []Node `json:"properties"`
}

// 对象的成员
type Property struct {
	Key   Node `json:"key"` // 一般都是Identifier
	Value Node `json:"value"`
}

// a.b.c这样的读取成员变量表达式
type MemberExpression struct {
	Object   Node `json:"object"`
	Property Node `json:"property"`
}

func (p Property) GetKey() string {
	switch t := p.Key.Assert().(type) {
	case Identifier:
		return t.Name
	case Literal:
		return t.Value.(string)
	default:
		panic(t)
	}

	return ""
}

// 获取a.b.c.d
func (p MemberExpression) GetKey() string {
	currKey := ""
	switch t := p.Property.Assert().(type) {
	case Identifier:
		currKey = t.Name
	case Literal:
		currKey = t.Value.(string)
	default:
		panic(t)
	}

	parentKey := ""
	switch t := p.Object.Assert().(type) {
	case MemberExpression:
		parentKey += t.GetKey() + "."
	case Identifier:
		parentKey += t.Name + "."
	case Literal:
		parentKey += t.Value.(string) + "."
	}

	return parentKey + currKey
}

// a.b.c这样的读取成员变量表达式
type CallExpression struct {
	Arguments []Node `json:"arguments"`
	Callee    Node   `json:"callee"`
}

func (c CallExpression) GetFuncName() string {
	switch t := c.Callee.Assert().(type) {
	case Identifier:
		return t.Name
	}
	return ""
}

var nodeMap = map[string]interface{}{
	"Program":             Program{},
	"ExpressionStatement": ExpressionStatement{},
	"BinaryExpression":    BinaryExpression{},
	"LogicalExpression":   LogicalExpression{},
	"Identifier":          Identifier{},
	"UnaryExpression":     UnaryExpression{},
	"Literal":             Literal{},
	"ObjectExpression":    ObjectExpression{},
	"Property":            Property{},
	"MemberExpression":    MemberExpression{},
	"CallExpression":      CallExpression{},
}

func (n Node) Assert() interface{} {
	typ,ok := n["type"].(string)
	if !ok{
		return nil
	}
	entity, ok := nodeMap[typ]
	if !ok {
		log.Errorf("unhand type:%s, %+v", typ, n)
		return nil
	}

	vPoint := reflect.New(reflect.TypeOf(entity))
	_, err := obj.Copy(n, vPoint.Interface())
	if err != nil {
		log.Error(err)
		return nil
	}

	return vPoint.Elem().Interface()
}

func GetAST(code string) (node Node, err error) {
	bs, _ := json.Marshal(map[string]string{
		"code": code,
	})
	status, res, err := client.Post("", bs, nil, &node)
	if err != nil {
		return
	}

	if status != 200 {
		err = errors.New(string(res))
		return
	}
	return
}
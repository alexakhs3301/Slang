package object

import (
	"Goslang/ast"
	"bytes"
	"fmt"
	"strings"
)

type ObjectType string

type Object interface {
	Type() ObjectType
	Inspect() string
}

const (
	INTEGER_OBJ    = "INTEGER"
	BOOLEAN_OBJ    = "BOOLEAN"
	NULL_OBJ       = "NULL"
	STRING_OBJ     = "STRING"
	RETURN_VAL_OBJ = "RETURN_VAL"
	ERROR_OBJ      = "ERROR"
	FUNCTION_OBJ   = "FUNCTION"
)

// Integer object
type Integer struct {
	Value int64
}

func (i *Integer) Inspect() string  { return fmt.Sprintf("%d", i.Value) }
func (i *Integer) Type() ObjectType { return INTEGER_OBJ }

// Boolean object
type Boolean struct {
	Value bool
}

func (b *Boolean) Inspect() string  { return fmt.Sprintf("%t", b.Value) }
func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJ }

// Null object
type Null struct{}

func (n *Null) Inspect() string  { return "null" }
func (n *Null) Type() ObjectType { return NULL_OBJ }

// String Object
type String struct {
	Value string
}

func (s *String) Inspect() string  { return s.Value }
func (s *String) Type() ObjectType { return STRING_OBJ }

// ReturnVal object
type ReturnVal struct {
	Value Object
}

func (rv *ReturnVal) Type() ObjectType { return RETURN_VAL_OBJ }
func (rv *ReturnVal) Inspect() string  { return rv.Value.Inspect() }

// Error object
type Error struct {
	Message string
}

func (e *Error) Type() ObjectType { return ERROR_OBJ }
func (e *Error) Inspect() string  { return "ERROR:" + e.Message }

// Function object
type Function struct {
	Name       *ast.Identifier
	Parameters []*ast.FunctionParameter
	ReturnType *ast.TypeAnnotation
	Block      *ast.BlockStatement
	Env        *Environment
}

func (f *Function) Type() ObjectType { return FUNCTION_OBJ }
func (f *Function) Inspect() string {
	var out bytes.Buffer
	params := []string{}

	for _, p := range f.Parameters {
		params = append(params, p.Name.String()+":"+p.Type.String())
	}

	out.WriteString("fn")
	if f.Name != nil {
		out.WriteString(f.Name.String())
	}
	out.WriteString("(")
	out.WriteString(strings.Join(params, ","))
	out.WriteString(")")

	if f.ReturnType != nil {
		out.WriteString(":" + f.ReturnType.String())
	}

	out.WriteString("{\n")
	out.WriteString(f.Block.String())
	out.WriteString("\n}")

	return out.String()
}

package entity

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"text/template"

	"github.com/hzylyq/hole/core"
)

func NewEntityFromFile(file string) (core.IEntity, error) {
	f, err := parser.ParseFile(token.NewFileSet(), file, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	ent := new(schema)

	ent.PackName = f.Name.Name
	var structName string

	ast.Inspect(f, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.TypeSpec:
			if x.Type == nil {
				return true
			}
			structName = x.Name.Name
		}

		s, ok := n.(*ast.StructType)
		if !ok {
			return true
		}

		ent.TableName = structName

		for _, f := range s.Fields.List {
			ent.Fields = append(ent.Fields, field{
				Name:  f.Names[0].Name,
				Value: f.Type,
			})
			fmt.Printf("Field: %s\n", f.Names[0].Name)
			fmt.Printf("Tag:   %s\n", f.Tag.Value)
		}

		return false
	})

	return ent, nil
}

func (s *schema) CreateRepo() (string, error) {
	temp, err := template.New(s.TableName).Parse(repoTemplate)
	if err != nil {
		return "", err
	}

	temp.Execute()

	return "", nil
}

func (s *schema) CreateApi() (string, error) {
	return "", nil
}

func (s *schema) CreateService() (string, error) {
	return "", nil
}

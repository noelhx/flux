// DO NOT EDIT: This file is autogenerated via the builtin command.

package sql

import (
	flux "github.com/influxdata/flux"
	ast "github.com/influxdata/flux/ast"
)

func init() {
	flux.RegisterPackage(pkgAST)
}

var pkgAST = &ast.Package{
	BaseNode: ast.BaseNode{
		Errors: nil,
		Loc:    nil,
	},
	Files: []*ast.File{&ast.File{
		BaseNode: ast.BaseNode{
			Errors: nil,
			Loc: &ast.SourceLocation{
				End: ast.Position{
					Column: 11,
					Line:   4,
				},
				File:   "sql.flux",
				Source: "package sql\n\nbuiltin from\nbuiltin to",
				Start: ast.Position{
					Column: 1,
					Line:   1,
				},
			},
		},
		Body: []ast.Statement{&ast.BuiltinStatement{
			BaseNode: ast.BaseNode{
				Errors: nil,
				Loc: &ast.SourceLocation{
					End: ast.Position{
						Column: 13,
						Line:   3,
					},
					File:   "sql.flux",
					Source: "builtin from",
					Start: ast.Position{
						Column: 1,
						Line:   3,
					},
				},
			},
			ID: &ast.Identifier{
				BaseNode: ast.BaseNode{
					Errors: nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 13,
							Line:   3,
						},
						File:   "sql.flux",
						Source: "from",
						Start: ast.Position{
							Column: 9,
							Line:   3,
						},
					},
				},
				Name: "from",
			},
		}, &ast.BuiltinStatement{
			BaseNode: ast.BaseNode{
				Errors: nil,
				Loc: &ast.SourceLocation{
					End: ast.Position{
						Column: 11,
						Line:   4,
					},
					File:   "sql.flux",
					Source: "builtin to",
					Start: ast.Position{
						Column: 1,
						Line:   4,
					},
				},
			},
			ID: &ast.Identifier{
				BaseNode: ast.BaseNode{
					Errors: nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 11,
							Line:   4,
						},
						File:   "sql.flux",
						Source: "to",
						Start: ast.Position{
							Column: 9,
							Line:   4,
						},
					},
				},
				Name: "to",
			},
		}},
		Imports:  nil,
		Metadata: "parser-type=rust",
		Name:     "sql.flux",
		Package: &ast.PackageClause{
			BaseNode: ast.BaseNode{
				Errors: nil,
				Loc: &ast.SourceLocation{
					End: ast.Position{
						Column: 12,
						Line:   1,
					},
					File:   "sql.flux",
					Source: "package sql",
					Start: ast.Position{
						Column: 1,
						Line:   1,
					},
				},
			},
			Name: &ast.Identifier{
				BaseNode: ast.BaseNode{
					Errors: nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 12,
							Line:   1,
						},
						File:   "sql.flux",
						Source: "sql",
						Start: ast.Position{
							Column: 9,
							Line:   1,
						},
					},
				},
				Name: "sql",
			},
		},
	}},
	Package: "sql",
	Path:    "sql",
}

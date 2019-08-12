// DO NOT EDIT: This file is autogenerated via the builtin command.

package experimental

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
					Column: 14,
					Line:   7,
				},
				File:   "experimental.flux",
				Source: "package experimental\n\nbuiltin addDuration\nbuiltin subDuration\n\n// An experimental version of group that has mode: \"extend\"\nbuiltin group",
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
						Column: 20,
						Line:   3,
					},
					File:   "experimental.flux",
					Source: "builtin addDuration",
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
							Column: 20,
							Line:   3,
						},
						File:   "experimental.flux",
						Source: "addDuration",
						Start: ast.Position{
							Column: 9,
							Line:   3,
						},
					},
				},
				Name: "addDuration",
			},
		}, &ast.BuiltinStatement{
			BaseNode: ast.BaseNode{
				Errors: nil,
				Loc: &ast.SourceLocation{
					End: ast.Position{
						Column: 20,
						Line:   4,
					},
					File:   "experimental.flux",
					Source: "builtin subDuration",
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
							Column: 20,
							Line:   4,
						},
						File:   "experimental.flux",
						Source: "subDuration",
						Start: ast.Position{
							Column: 9,
							Line:   4,
						},
					},
				},
				Name: "subDuration",
			},
		}, &ast.BuiltinStatement{
			BaseNode: ast.BaseNode{
				Errors: nil,
				Loc: &ast.SourceLocation{
					End: ast.Position{
						Column: 14,
						Line:   7,
					},
					File:   "experimental.flux",
					Source: "builtin group",
					Start: ast.Position{
						Column: 1,
						Line:   7,
					},
				},
			},
			ID: &ast.Identifier{
				BaseNode: ast.BaseNode{
					Errors: nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 14,
							Line:   7,
						},
						File:   "experimental.flux",
						Source: "group",
						Start: ast.Position{
							Column: 9,
							Line:   7,
						},
					},
				},
				Name: "group",
			},
		}},
		Imports: nil,
		Name:    "experimental.flux",
		Package: &ast.PackageClause{
			BaseNode: ast.BaseNode{
				Errors: nil,
				Loc: &ast.SourceLocation{
					End: ast.Position{
						Column: 21,
						Line:   1,
					},
					File:   "experimental.flux",
					Source: "package experimental",
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
							Column: 21,
							Line:   1,
						},
						File:   "experimental.flux",
						Source: "experimental",
						Start: ast.Position{
							Column: 9,
							Line:   1,
						},
					},
				},
				Name: "experimental",
			},
		},
	}},
	Package: "experimental",
	Path:    "experimental",
}

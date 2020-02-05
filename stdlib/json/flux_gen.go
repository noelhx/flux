// DO NOT EDIT: This file is autogenerated via the builtin command.

package json

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
					Column: 15,
					Line:   9,
				},
				File:   "json.flux",
				Source: "package json\n\n// encode converts a value into JSON bytes\n// Time values are encoded using RFC3339.\n// Duration values are encoded in number of milleseconds since the epoch.\n// Regexp values are encoded as their string representation.\n// Bytes values are encodes as base64-encoded strings.\n// Function values cannot be encoded and will produce an error.\nbuiltin encode",
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
						Column: 15,
						Line:   9,
					},
					File:   "json.flux",
					Source: "builtin encode",
					Start: ast.Position{
						Column: 1,
						Line:   9,
					},
				},
			},
			ID: &ast.Identifier{
				BaseNode: ast.BaseNode{
					Errors: nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 15,
							Line:   9,
						},
						File:   "json.flux",
						Source: "encode",
						Start: ast.Position{
							Column: 9,
							Line:   9,
						},
					},
				},
				Name: "encode",
			},
		}},
		Imports:  nil,
		Metadata: "parser-type=rust",
		Name:     "json.flux",
		Package: &ast.PackageClause{
			BaseNode: ast.BaseNode{
				Errors: nil,
				Loc: &ast.SourceLocation{
					End: ast.Position{
						Column: 13,
						Line:   1,
					},
					File:   "json.flux",
					Source: "package json",
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
							Column: 13,
							Line:   1,
						},
						File:   "json.flux",
						Source: "json",
						Start: ast.Position{
							Column: 9,
							Line:   1,
						},
					},
				},
				Name: "json",
			},
		},
	}},
	Package: "json",
	Path:    "json",
}

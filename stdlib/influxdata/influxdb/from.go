// From is an operation that mocks the real implementation of InfluxDB's from.
// It is used in Flux to compile queries that resemble real queries issued against InfluxDB.
// Implementors of the real from are expected to replace its implementation via flux.ReplacePackageValue.
package influxdb

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/influxdata/flux"
	"github.com/influxdata/flux/ast"
	"github.com/influxdata/flux/codes"
	"github.com/influxdata/flux/csv"
	"github.com/influxdata/flux/execute"
	"github.com/influxdata/flux/internal/errors"
	"github.com/influxdata/flux/memory"
	"github.com/influxdata/flux/plan"
	"github.com/influxdata/flux/semantic"
	"github.com/influxdata/flux/stdlib/universe"
)

const FromKind = "from"

const DefaultHost = "http://localhost:9999"

type FromOpSpec struct {
	Org    string
	Bucket string
	Host   *string
	Token  *string
}

func init() {
	fromSignature := semantic.FunctionPolySignature{
		Parameters: map[string]semantic.PolyType{
			"org":    semantic.String,
			"bucket": semantic.String,
			"host":   semantic.String,
			"token":  semantic.String,
		},
		Required: []string{"bucket"},
		Return:   flux.TableObjectType,
	}

	flux.RegisterPackageValue("influxdata/influxdb", FromKind, flux.FunctionValue(FromKind, createFromOpSpec, fromSignature))
	flux.RegisterOpSpec(FromKind, newFromOp)
	plan.RegisterProcedureSpec(FromKind, newFromProcedure, FromKind)
	execute.RegisterSource(FromKind, createFromSource)
	plan.RegisterPhysicalRules(
		PushDownRangeRule{},
	)
}

func createFromOpSpec(args flux.Arguments, a *flux.Administration) (flux.OperationSpec, error) {
	spec := new(FromOpSpec)
	if b, _, e := args.GetString("bucket"); e != nil {
		return nil, e
	} else {
		spec.Bucket = b
	}
	if o, ok, err := args.GetString("org"); err != nil {
		return nil, err
	} else if ok {
		spec.Org = o
	}

	if h, ok, err := args.GetString("host"); err != nil {
		return nil, err
	} else if ok {
		spec.Host = &h
	}

	if token, ok, err := args.GetString("token"); err != nil {
		return nil, err
	} else if ok {
		spec.Token = &token
	}
	return spec, nil
}

func newFromOp() flux.OperationSpec {
	return new(FromOpSpec)
}

func (s *FromOpSpec) Kind() flux.OperationKind {
	return FromKind
}

type FromProcedureSpec struct {
	plan.DefaultCost

	Org    string
	Bucket string
	Host   string
	Token  *string
	Range  *universe.RangeProcedureSpec
}

func newFromProcedure(qs flux.OperationSpec, pa plan.Administration) (plan.ProcedureSpec, error) {
	spec, ok := qs.(*FromOpSpec)
	if !ok {
		return nil, errors.Newf(codes.Internal, "invalid spec type %T", qs)
	}

	host := DefaultHost
	if spec.Host != nil {
		host = *spec.Host
	}
	return &FromProcedureSpec{
		Org:    spec.Org,
		Bucket: spec.Bucket,
		Host:   host,
		Token:  spec.Token,
	}, nil
}

func (s *FromProcedureSpec) Kind() plan.ProcedureKind {
	return FromKind
}

func (s *FromProcedureSpec) Copy() plan.ProcedureSpec {
	ns := new(FromProcedureSpec)
	*ns = *s
	if s.Range != nil {
		ns.Range = s.Range.Copy().(*universe.RangeProcedureSpec)
	}
	return ns
}

type source struct {
	id   execute.DatasetID
	spec *FromProcedureSpec
	deps flux.Dependencies
	mem  *memory.Allocator
	ts   execute.TransformationSet
}

func createFromSource(s plan.ProcedureSpec, id execute.DatasetID, a execute.Administration) (execute.Source, error) {
	spec := s.(*FromProcedureSpec)
	if spec.Range == nil {
		return nil, errors.Newf(codes.Invalid, "bounds must be set")
	}

	// These parameters are only required for the remote influxdb
	// source. If running flux within influxdb, these aren't
	// required.
	if spec.Org == "" {
		return nil, errors.Newf(codes.Invalid, "org must be set")
	}

	// Host isn't a required parameter, but it must be set for this
	// specific implementation.
	if spec.Host == "" {
		return nil, errors.Newf(codes.Invalid, "host must be set")
	}

	deps := flux.GetDependencies(a.Context())
	return &source{
		id:   id,
		spec: spec,
		deps: deps,
		mem:  a.Allocator(),
	}, nil
}

func (s *source) AddTransformation(t execute.Transformation) {
	s.ts = append(s.ts, t)
}

func (s *source) Run(ctx context.Context) {
	err := s.run(ctx)
	s.ts.Finish(s.id, err)
}

func (s *source) run(ctx context.Context) error {
	req, err := s.newRequest(ctx)
	if err != nil {
		return err
	}

	client, err := s.deps.HTTPClient()
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	} else if resp.StatusCode != 200 {
		io.Copy(os.Stderr, resp.Body)
		return errors.New(codes.Internal, "todo(jsternberg): read the error message")
	}
	return s.processResults(resp.Body)
}

func (s *source) newRequest(ctx context.Context) (*http.Request, error) {
	u, err := url.Parse(s.spec.Host)
	if err != nil {
		return nil, err
	}
	u.Path += "/api/v2/query"
	u.RawQuery = func() string {
		params := make(url.Values)
		params.Set("org", s.spec.Org)
		return params.Encode()
	}()

	// Validate that the produced url is allowed.
	urlv, err := s.deps.URLValidator()
	if err != nil {
		return nil, err
	}

	if err := urlv.Validate(u); err != nil {
		return nil, err
	}

	body, err := s.newRequestBody()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", u.String(), bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	if s.spec.Token != nil {
		req.Header.Set("Authorization", "Token "+*s.spec.Token)
	}
	req.Header.Set("Content-Type", "application/json")
	return req.WithContext(ctx), nil
}

func (s *source) newRequestBody() ([]byte, error) {
	var req struct {
		AST     *ast.Package `json:"ast"`
		Dialect struct {
			Header         bool     `json:"header"`
			DateTimeFormat string   `json:"dateTimeFormat"`
			Annotations    []string `json:"annotations"`
		} `json:"dialect"`
	}
	req.AST = &ast.Package{
		Package: "main",
		Files: []*ast.File{{
			Package: &ast.PackageClause{
				Name: &ast.Identifier{Name: "main"},
			},
			Name: "query.flux",
			Body: []ast.Statement{
				&ast.ExpressionStatement{Expression: s.buildQuery()},
			},
		}},
	}
	req.Dialect.Header = true
	req.Dialect.DateTimeFormat = "RFC3339Nano"
	req.Dialect.Annotations = []string{"group", "datatype", "default"}
	return json.Marshal(req)
}

func (s *source) buildQuery() ast.Expression {
	return &ast.PipeExpression{
		Argument: &ast.CallExpression{
			Callee:    &ast.Identifier{Name: "from"},
			Arguments: []ast.Expression{s.fromArgs()},
		},
		Call: &ast.CallExpression{
			Callee:    &ast.Identifier{Name: "range"},
			Arguments: []ast.Expression{s.rangeArgs()},
		},
	}
}

func (s *source) fromArgs() *ast.ObjectExpression {
	return &ast.ObjectExpression{
		Properties: []*ast.Property{{
			Key:   &ast.Identifier{Name: "bucket"},
			Value: &ast.StringLiteral{Value: s.spec.Bucket},
		}},
	}
}

func (s *source) rangeArgs() *ast.ObjectExpression {
	toLiteral := func(t flux.Time) ast.Literal {
		if t.IsRelative {
			// TODO(jsternberg): This seems wrong. Relative should be a values.Duration
			// and not a time.Duration.
			d := flux.ConvertDuration(t.Relative)
			return &ast.DurationLiteral{Values: d.AsValues()}
		}
		return &ast.DateTimeLiteral{Value: t.Absolute}
	}

	args := make([]*ast.Property, 0, 2)
	args = append(args, &ast.Property{
		Key:   &ast.Identifier{Name: "start"},
		Value: toLiteral(s.spec.Range.Bounds.Start),
	})
	if !s.spec.Range.Bounds.Stop.IsZero() {
		args = append(args, &ast.Property{
			Key:   &ast.Identifier{Name: "stop"},
			Value: toLiteral(s.spec.Range.Bounds.Stop),
		})
	}
	return &ast.ObjectExpression{Properties: args}
}

func (s *source) processResults(r io.ReadCloser) error {
	defer func() { _ = r.Close() }()

	config := csv.ResultDecoderConfig{Allocator: s.mem}
	dec := csv.NewMultiResultDecoder(config)
	results, err := dec.Decode(r)
	if err != nil {
		return err
	}
	defer results.Release()

	for results.More() {
		res := results.Next()
		if err := res.Tables().Do(func(table flux.Table) error {
			return s.ts.Process(s.id, table)
		}); err != nil {
			return err
		}
	}
	results.Release()
	return results.Err()
}

type PushDownRangeRule struct{}

func (p PushDownRangeRule) Name() string {
	return "PushDownRangeRule"
}

func (p PushDownRangeRule) Pattern() plan.Pattern {
	return plan.Pat(universe.RangeKind, plan.Pat(FromKind))
}

func (p PushDownRangeRule) Rewrite(node plan.Node) (plan.Node, bool, error) {
	fromNode := node.Predecessors()[0]
	fromSpec := fromNode.ProcedureSpec().(*FromProcedureSpec)
	if fromSpec.Range != nil {
		return node, false, nil
	}

	rangeSpec := node.ProcedureSpec().(*universe.RangeProcedureSpec)
	newFromSpec := fromSpec.Copy().(*FromProcedureSpec)
	newFromSpec.Range = rangeSpec
	n, err := plan.MergeToPhysicalNode(node, fromNode, newFromSpec)
	if err != nil {
		return nil, false, err
	}
	return n, true, nil
}

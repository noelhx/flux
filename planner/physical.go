package planner

import (
	"errors"
	"fmt"
	"math"
)

// PhysicalPlanner performs transforms a logical plan to a physical plan,
// by applying any registered physical rules.
type PhysicalPlanner interface {
	Plan(lplan *PlanSpec) (*PlanSpec, error)
}

// NewPhysicalPlanner creates a new physical planner with the specified options.
// The new planner will be configured to apply any physical rules that have been registered.
func NewPhysicalPlanner(options ...PhysicalOption) PhysicalPlanner {
	pp := &physicalPlanner{
		heuristicPlanner:   newHeuristicPlanner(),
		defaultMemoryLimit: math.MaxInt64,
	}

	rules := make([]Rule, len(ruleNameToPhysicalRule))
	i := 0
	for _, v := range ruleNameToPhysicalRule {
		rules[i] = v
		i++
	}

	pp.addRules(rules)

	// Options may add or remove rules, so process them after we've
	// added registered rules.
	for _, opt := range options {
		opt.apply(pp)
	}

	return pp
}

func (pp *physicalPlanner) Plan(spec *PlanSpec) (*PlanSpec, error) {
	transformedSpec, err := pp.heuristicPlanner.Plan(spec)
	if err != nil {
		return nil, err
	}

	// Convert yields into result list
	// TODO: Implement this via a transformation rule
	final, err := removeYields(transformedSpec)
	if err != nil {
		return nil, err
	}

	// Compute time bounds for nodes in the plan
	if err := final.BottomUpWalk(ComputeBounds); err != nil {
		return nil, err
	}

	// Update memory quota
	if final.Resources.MemoryBytesQuota == 0 {
		final.Resources.MemoryBytesQuota = pp.defaultMemoryLimit
	}

	// Update concurrency quota
	if final.Resources.ConcurrencyQuota == 0 {
		final.Resources.ConcurrencyQuota = len(spec.Results)
	}

	return final, nil
}

// TODO: This procedure should be encapsulated in a yield rewrite rule
func removeYields(plan *PlanSpec) (*PlanSpec, error) {
	for root := range plan.Roots {

		name := DefaultYieldName

		if yield, ok := root.ProcedureSpec().(YieldProcedureSpec); ok {

			name = yield.YieldName()

			if len(root.Predecessors()) != 1 {
				return nil, errors.New("yield must have exactly one predecessor")
			}

			if _, ok := plan.Results[name]; ok {
				return nil, fmt.Errorf("found duplicate yield name %q", name)
			}

			newRoot := root.Predecessors()[0]
			newRoot.RemoveSuccessor(root)
			plan.Replace(root, newRoot)
			plan.Results[name] = newRoot
			continue
		}

		if _, ok := plan.Results[name]; ok {
			return nil, fmt.Errorf("found duplicate yield name %q", name)
		}

		plan.Results[name] = root
	}
	return plan, nil
}

type physicalPlanner struct {
	*heuristicPlanner
	defaultMemoryLimit int64
}

// PhysicalOption is an option to configure the behavior of the physical planner.
type PhysicalOption interface {
	apply(*physicalPlanner)
}

type physicalOption func(*physicalPlanner)

func (opt physicalOption) apply(p *physicalPlanner) {
	opt(p)
}

// WithDefaultMemoryLimit sets the default memory limit for plans generated by the planner.
// If the query spec explicitly sets a memory limit, that limit is used instead of the default.
func WithDefaultMemoryLimit(memBytes int64) PhysicalOption {
	return physicalOption(func(p *physicalPlanner) {
		p.defaultMemoryLimit = memBytes
	})
}

// PhysicalProcedureSpec is similar to its logical counterpart but must provide a method to determine cost.
type PhysicalProcedureSpec interface {
	Kind() ProcedureKind
	Copy() ProcedureSpec
	Cost(inStats []Statistics) (cost Cost, outStats Statistics)
}

// PhysicalPlanNode represents a physical operation in a plan.
type PhysicalPlanNode struct {
	edges
	bounds
	id   NodeID
	Spec PhysicalProcedureSpec

	// The attributes required from inputs to this node
	RequiredAttrs []PhysicalAttributes

	// The attributes provided to consumers of this node's output
	OutputAttrs PhysicalAttributes
}

// ID returns a human-readable id for this plan node.
func (ppn *PhysicalPlanNode) ID() NodeID {
	return ppn.id
}

// ProcedureSpec returns the procedure spec for this plan node.
func (ppn *PhysicalPlanNode) ProcedureSpec() ProcedureSpec {
	return ppn.Spec
}

// Kind returns the procedure kind for this plan node.
func (ppn *PhysicalPlanNode) Kind() ProcedureKind {
	return ppn.Spec.Kind()
}

func (ppn *PhysicalPlanNode) ShallowCopy() PlanNode {
	newNode := new(PhysicalPlanNode)
	newNode.edges = ppn.edges.shallowCopy()
	newNode.id = ppn.id + "_copy"
	// TODO: the type assertion below... is it needed?
	newNode.Spec = ppn.Spec.Copy().(PhysicalProcedureSpec)
	return newNode
}

// Cost provides the self-cost (i.e., does not include the cost of its predecessors) for
// this plan node.  Caller must provide statistics of predecessors to this node.
func (ppn *PhysicalPlanNode) Cost(inStats []Statistics) (cost Cost, outStats Statistics) {
	return ppn.Spec.Cost(inStats)
}

// PhysicalAttributes encapsulates sny physical attributes of the result produced
// by a physical plan node, such as collation, etc.
type PhysicalAttributes struct {
}

// CreatePhysicalNode creates a single physical plan node from a procedure spec.
// The newly created physical node has no incoming or outgoing edges.
func CreatePhysicalNode(id NodeID, spec PhysicalProcedureSpec) *PhysicalPlanNode {
	return &PhysicalPlanNode{
		id:   id,
		Spec: spec,
	}
}

package predicates

import (
	"k8s.io/api/core/v1"
	"k8s.io/kubernetes/pkg/scheduler/cache"
)

// Predicate checks whether pods are eligible to be scheduled on nodes. Each
// predicate is responsible for a specific eligibility check,
// e.g. requested CPU <= available CPU.
type Predicate struct {
	// name is used in config to select predicates that are enabled or disabled.
	name string
	// fit checks whether the pod fits on the node.
	fit fitFunc
	// precompute is optional. It populates Metadata fields used by this predicate.
	precompute precomputeFunc

	// NOTE: This type is a struct with unexported fields specifically so that
	// other packages cannot create a Predicate value and assign it to one of our
	// global predicate variables.
}

// Name returns the name of the Predicate.
func (p Predicate) Name() string { return p.name }

// Fit checks whether the given pod fits on the given node. If the pod does not
// fit, a list of FitErrors is returned.
func (p Predicate) Fit(pod *v1.Pod, meta *Metadata, node *cache.NodeInfo) (bool, []FitError, error) {
	// TODO: All existing predicate implementations tolerate nil Metadata. Should
	// we keep that convention? It would be simpler to remove that complexity from
	// predicates and put it here.
	return p.fit(pod, meta, node)
}

// Precompute updates the given Metadata object with data that is relevant to
// this Predicate and specific to the given pod.
func (p Predicate) Precompute(pod *v1.Pod, nodes map[string]*cache.NodeInfo, meta *Metadata) {
	if p.precompute != nil {
		p.precompute(pod, nodes, meta)
	}
}

// fitFunc is a function that checks whether a pod fits on a node.
type fitFunc = func(*v1.Pod, *Metadata, *cache.NodeInfo) (bool, []FitError, error)

// precomputeFunc is a function that populates part of a Metadata object.
type precomputeFunc = func(*v1.Pod, map[string]*cache.NodeInfo, *Metadata)

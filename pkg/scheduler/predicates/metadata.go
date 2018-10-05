package predicates

import (
	"k8s.io/api/core/v1"
	"k8s.io/kubernetes/pkg/scheduler/cache"
)

// Metadata contains predicate data for a given particular pod. Predicates use
// Metadata when evaluating nodes for that pod. Metadata is built from a
// snapshot of the scheduler's cache so that predicates can evaluate with a
// consistent view of the cluster.
type Metadata struct {
	// TODO: Copy fields from old predicateMetadata struct.
	todo string
}

// NewMetadata computes the predicate metadata for a given pod. `predicates` is
// the list of predicates that will be used to evaluate nodes. The nodes map
// should be obtained from `UpdateNodeNameToInfoMap()` in the scheduler cache.
func NewMetadata(predicates []Predicate, pod *v1.Pod, nodes map[string]*cache.NodeInfo) *Metadata {
	meta := &Metadata{
		// TODO: initialize common metadata
	}
	for _, p := range predicates {
		p.Precompute(pod, nodes, meta)
	}
	return meta
}

// TODO: AddPod, RemovePod, and ShallowCopy implementations

// AddPod changes Metadata as if `addedPod` was added to the system.
func (m *Metadata) AddPod(addedPod *v1.Pod, node *cache.NodeInfo) error { return nil }

// RemovePod changes Metadata as if `deletedPod` was deleted from the system.
func (m *Metadata) RemovePod(deletedPod *v1.Pod) error { return nil }

// ShallowCopy copies a Metadata, creating copies of its maps and slices, but
// copying pointers as-is.
func (m *Metadata) ShallowCopy() *Metadata { return nil }

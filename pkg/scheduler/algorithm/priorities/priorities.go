package priorities

import (
	"k8s.io/api/core/v1"
	schedulerapi "k8s.io/kubernetes/pkg/scheduler/api"
	schedulernodeinfo "k8s.io/kubernetes/pkg/scheduler/nodeinfo"
)

// PriorityMapFunction is a function that computes per-node results for a given node.
// TODO: Figure out the exact API of this method.
// TODO: Change interface{} to a specific type.
type PriorityMapFunction func(pod *v1.Pod, meta interface{}, nodeInfo *schedulernodeinfo.NodeInfo) (schedulerapi.HostPriority, error)

// PriorityReduceFunction is a function that aggregated per-node results and computes
// final scores for all nodes.
// TODO: Figure out the exact API of this method.
// TODO: Change interface{} to a specific type.
type PriorityReduceFunction func(pod *v1.Pod, meta interface{}, nodeNameToInfo map[string]*schedulernodeinfo.NodeInfo, result schedulerapi.HostPriorityList) error

// PriorityMetadataProducer is a function that computes metadata for a given pod. This
// is now used for only for priority functions. For predicates please use PredicateMetadataProducer.
type PriorityMetadataProducer func(pod *v1.Pod, nodeNameToInfo map[string]*schedulernodeinfo.NodeInfo) interface{}

// PriorityFunction is a function that computes scores for all nodes.
// DEPRECATED
// Use Map-Reduce pattern for priority functions.
type PriorityFunction func(pod *v1.Pod, nodeNameToInfo map[string]*schedulernodeinfo.NodeInfo, nodes []*v1.Node) (schedulerapi.HostPriorityList, error)

// PriorityConfig is a config used for a priority function.
type PriorityConfig struct {
	Name   string
	Map    PriorityMapFunction
	Reduce PriorityReduceFunction
	// TODO: Remove it after migrating all functions to
	// Map-Reduce pattern.
	Function PriorityFunction
	Weight   int
}

// EmptyPriorityMetadataProducer returns a no-op PriorityMetadataProducer type.
func EmptyPriorityMetadataProducer(pod *v1.Pod, nodeNameToInfo map[string]*schedulernodeinfo.NodeInfo) interface{} {
	return nil
}

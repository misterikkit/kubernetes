// Package predicates defines functions used to check whether pods can be scheduled onto nodes.
//
// Writing a New Predicate
//
// New predicates should be defined as package level variables, using the Predicate struct. e.g.
//
//  // PodToleratesNodeTaints compares pod tolerations against node taints.
//  // Note: Explicitly declaring the var's type ensures that godoc is organized correctly.
//  var PodToleratesNodeTaints Predicate = Predicate{
//  	name: "PodToleratesNodeTaints",
//  	fit: func(pod *v1.Pod, meta *Metadata, node *cache.NodeInfo) (bool, []FitError, error) {
//  		// Comparison code goes here.
//  		return fits, failReasons, nil
//  	},
//  }
//
// Each predicate and its unit tests should go in new files. e.g. "taints.go" and "taints_test.go"
//
// TODO: Should we make it easy to write stateful predicates?
//
// Predicate Evaluation Order
//
// TODO: write section
//
// Using Metadata
//
// TODO: write about the proper care and handling of Metadata
package predicates

// TODO: These are here for godoc demo purposes only. Delete them.

// CheckInterestingCondition is a very interesting predicate.
var CheckInterestingCondition Predicate = Predicate{name: "CheckInterestingCondition"}

// CheckBoringCondition is not very interesting.
var CheckBoringCondition Predicate = Predicate{name: "CheckBoringCondition"}

// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package redshift

import (
	"github.com/awslabs/aws-sdk-go/internal/waiter"
)
var waiterClusterAvailable *waiter.Config

func (c *Redshift) WaitUntilClusterAvailable(input *DescribeClustersInput) error {
	if waiterClusterAvailable == nil {
		waiterClusterAvailable = &waiter.Config{
			Operation:   "DescribeClusters",
			Delay:       60,
			MaxAttempts: 30,
			Acceptors: []waiter.WaitAcceptor{
				waiter.WaitAcceptor{
					State:    "success",
					Matcher:  "pathAll",
					Argument: "Clusters[].ClusterStatus",
					Expected: "available",
				},
				waiter.WaitAcceptor{
					State:    "failure",
					Matcher:  "pathAny",
					Argument: "Clusters[].ClusterStatus",
					Expected: "deleting",
				},
				waiter.WaitAcceptor{
					State:    "retry",
					Matcher:  "error",
					Argument: "",
					Expected: "ClusterNotFound",
				},
				
			},
		}
	}

	w := waiter.Waiter{
		Client: c,
		Input:  input,
		Config: waiterClusterAvailable,
	}
	return w.Wait()
}

var waiterClusterDeleted *waiter.Config

func (c *Redshift) WaitUntilClusterDeleted(input *DescribeClustersInput) error {
	if waiterClusterDeleted == nil {
		waiterClusterDeleted = &waiter.Config{
			Operation:   "DescribeClusters",
			Delay:       60,
			MaxAttempts: 30,
			Acceptors: []waiter.WaitAcceptor{
				waiter.WaitAcceptor{
					State:    "success",
					Matcher:  "error",
					Argument: "",
					Expected: "ClusterNotFound",
				},
				waiter.WaitAcceptor{
					State:    "failure",
					Matcher:  "pathAny",
					Argument: "Clusters[].ClusterStatus",
					Expected: "creating",
				},
				waiter.WaitAcceptor{
					State:    "failure",
					Matcher:  "pathAny",
					Argument: "Clusters[].ClusterStatus",
					Expected: "rebooting",
				},
				
			},
		}
	}

	w := waiter.Waiter{
		Client: c,
		Input:  input,
		Config: waiterClusterDeleted,
	}
	return w.Wait()
}

var waiterSnapshotAvailable *waiter.Config

func (c *Redshift) WaitUntilSnapshotAvailable(input *DescribeClusterSnapshotsInput) error {
	if waiterSnapshotAvailable == nil {
		waiterSnapshotAvailable = &waiter.Config{
			Operation:   "DescribeClusterSnapshots",
			Delay:       15,
			MaxAttempts: 20,
			Acceptors: []waiter.WaitAcceptor{
				waiter.WaitAcceptor{
					State:    "success",
					Matcher:  "pathAll",
					Argument: "Snapshots[].Status",
					Expected: "available",
				},
				waiter.WaitAcceptor{
					State:    "failure",
					Matcher:  "pathAny",
					Argument: "Snapshots[].Status",
					Expected: "failed",
				},
				waiter.WaitAcceptor{
					State:    "failure",
					Matcher:  "pathAny",
					Argument: "Snapshots[].Status",
					Expected: "deleted",
				},
				
			},
		}
	}

	w := waiter.Waiter{
		Client: c,
		Input:  input,
		Config: waiterSnapshotAvailable,
	}
	return w.Wait()
}

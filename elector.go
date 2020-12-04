package goelector

import (
	"context"
	"sync/atomic"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/leaderelection"
	"k8s.io/client-go/tools/leaderelection/resourcelock"
)

var (
	leading       int32
	leaderelector *leaderelection.LeaderElector
)

// StartWithCallbacks with user defined callbacks for set/unset
// Caller is responsible for calling cancel on context to release lease
func StartWithCallbacks(ctx context.Context, cfg *Config, nodeID string, client *kubernetes.Clientset, startFn *func(context.Context), stopFn *func(), newFn *func(string)) error {
	return start(ctx, cfg, nodeID, client, startFn, stopFn, newFn)
}

// Start with default callback implementation which is a basic atomic store/load int32
//
// Use when only set/unset leader is needed
// Caller is responsible for calling cancel on context to release lease
func Start(ctx context.Context, cfg *Config, nodeID string, client *kubernetes.Clientset) error {
	return start(ctx, cfg, nodeID, client, nil, nil, nil)
}

func start(ctx context.Context, cfg *Config, nodeID string, client *kubernetes.Clientset, startFn *func(context.Context), stopFn *func(), newFn *func(string)) error {
	var lock = &resourcelock.LeaseLock{
		LeaseMeta: metav1.ObjectMeta{
			Name:      cfg.Lock,
			Namespace: cfg.Namespace,
		},
		Client:     client.CoordinationV1(),
		LockConfig: resourcelock.ResourceLockConfig{Identity: nodeID},
	}

	cb := leaderelection.LeaderCallbacks{}
	if startFn == nil {
		cb.OnStartedLeading = startedLeading
	} else {
		cb.OnStartedLeading = *startFn
	}

	if stopFn == nil {
		cb.OnStoppedLeading = stoppedLeading
	} else {
		cb.OnStoppedLeading = *stopFn
	}

	if newFn != nil {
		cb.OnNewLeader = *newFn
	}

	var err error
	leaderelector, err = leaderelection.NewLeaderElector(leaderelection.LeaderElectionConfig{
		Lock:            lock,
		ReleaseOnCancel: true,
		LeaseDuration:   time.Duration(cfg.LeaseDuration) * time.Second,
		RenewDeadline:   time.Duration(cfg.RenewDeadline) * time.Second,
		RetryPeriod:     time.Duration(cfg.RetryPeriod) * time.Second,
		Callbacks:       cb,
	})
	if err != nil {
		return err
	}

	leaderelector.Run(ctx)
	return nil
}

func startedLeading(ctx context.Context) {
	atomic.StoreInt32(&leading, 1)
}

func stoppedLeading() {
	atomic.StoreInt32(&leading, 0)
}

// IsLeader returns if leading currently is reporting being the leader
func IsLeader() bool {
	return atomic.LoadInt32(&leading) == 1 || leaderelector.IsLeader()
}

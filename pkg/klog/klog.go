package klog

import (
	"context"

	k8slog "k8s.io/klog/v2"
)

func C(ctx context.Context) {
	k8slog.FromContext(ctx)
}

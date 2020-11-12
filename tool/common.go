package tool

import (
	"io"
	"sync"

	"github.com/fatih/color"
	"github.com/hashicorp/go-getter"
	"github.com/vbauerster/mpb/v5"
	"github.com/vbauerster/mpb/v5/decor"
)

var _ getter.ProgressTracker = &ProgressTracker{}

// ProgressTracker wraps a github.com/cheggaaa/pb.Pool
// in order to display download progress for one or multiple
// downloads.
type ProgressTracker struct {
	group    *sync.WaitGroup
	progress *mpb.Progress
}

// NewProgressTracker creates new tracker
func NewProgressTracker() *ProgressTracker {
	group := &sync.WaitGroup{}

	return &ProgressTracker{
		group: group,
		progress: mpb.New(
			mpb.WithWidth(80),
			mpb.WithWaitGroup(group),
		),
	}
}

// Add adds delta
func (t *ProgressTracker) Add(delta int) {
	t.group.Add(delta)
}

// Wait waits all bars
func (t *ProgressTracker) Wait() {
	t.progress.Wait()
}

// TrackProgress instantiates a new progress bar that will
// display the progress of stream until closed.
// total can be 0.
func (t *ProgressTracker) TrackProgress(name string, current, total int64, stream io.ReadCloser) io.ReadCloser {
	bar := t.progress.AddBar(total,
		mpb.PrependDecorators(
			decor.Name(name, decor.WCSyncSpace),
			decor.Percentage(decor.WCSyncSpace),
		),
		mpb.AppendDecorators(
			decor.OnComplete(
				decor.EwmaETA(decor.ET_STYLE_GO, 60, decor.WCSyncWidth), color.HiGreenString("done"),
			),
		),
	)

	return bar.ProxyReader(stream)
}

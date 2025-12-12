package system

import (
	"github.com/joshuakb2/moby/api/types/build"
	"github.com/joshuakb2/moby/api/types/container"
	"github.com/joshuakb2/moby/api/types/image"
	"github.com/joshuakb2/moby/api/types/volume"
)

// DiskUsageObject represents an object type used for disk usage query filtering.
type DiskUsageObject string

const (
	// ContainerObject represents a container DiskUsageObject.
	ContainerObject DiskUsageObject = "container"
	// ImageObject represents an image DiskUsageObject.
	ImageObject DiskUsageObject = "image"
	// VolumeObject represents a volume DiskUsageObject.
	VolumeObject DiskUsageObject = "volume"
	// BuildCacheObject represents a build-cache DiskUsageObject.
	BuildCacheObject DiskUsageObject = "build-cache"
)

// DiskUsage contains response of Engine API:
// GET "/system/df"
type DiskUsage struct {
	ImageUsage      *image.DiskUsage     `json:"ImageUsage,omitempty"`
	ContainerUsage  *container.DiskUsage `json:"ContainerUsage,omitempty"`
	VolumeUsage     *volume.DiskUsage    `json:"VolumeUsage,omitempty"`
	BuildCacheUsage *build.DiskUsage     `json:"BuildCacheUsage,omitempty"`
}

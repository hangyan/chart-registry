package storage

import (
	"fmt"
	"github.com/hangyan/chart-registry/pkg/storage/registry"
	"path/filepath"
	"time"
)

type Object struct {
	Path         string
	Content      []byte
	LastModified time.Time

	Name string
}

func NewObject(chart *registry.ChartObject) *Object {
	var object Object
	object.Path = chart.Chart.Metadata.Name + "-" + chart.Chart.Metadata.Version + ".tgz"
	object.Content = chart.ChartContent
	// todo: fix
	object.LastModified = time.Now()
	object.Name = chart.Name
	return &object

}

// HasExtension determines whether or not an object contains a file extension
func (object Object) HasExtension(extension string) bool {
	return filepath.Ext(object.Path) == fmt.Sprintf(".%s", extension)
}

type Backend interface {
	ListObjects(prefix string) ([]Object, error)
	GetObject(path string) (Object, error)
	PutObject(path string, content []byte) error
	DeleteObject(path string) error
}

// ObjectSliceDiff provides information on what has changed since last calling ListObjects
type ObjectSliceDiff struct {
	Change  bool
	Removed []Object
	Added   []Object
	Updated []Object
}

// GetObjectSliceDiff takes two objects slices and returns an ObjectSliceDiff
func GetObjectSliceDiff(os1 []Object, os2 []Object) ObjectSliceDiff {
	var diff ObjectSliceDiff
	for _, o1 := range os1 {
		found := false
		for _, o2 := range os2 {
			if o1.Path == o2.Path {
				found = true
				if !o1.LastModified.Equal(o2.LastModified) {
					diff.Updated = append(diff.Updated, o2)
				}
				break
			}
		}
		if !found {
			diff.Removed = append(diff.Removed, o1)
		}
	}
	for _, o2 := range os2 {
		found := false
		for _, o1 := range os1 {
			if o2.Path == o1.Path {
				found = true
				break
			}
		}
		if !found {
			diff.Added = append(diff.Added, o2)
		}
	}
	diff.Change = len(diff.Removed)+len(diff.Added)+len(diff.Updated) > 0
	return diff
}

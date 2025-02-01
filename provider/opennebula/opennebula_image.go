//go:build opennebula || one || !onlyprovider

package opennebula

import (
	"github.com/nanovms/ops/lepton"
	"github.com/nanovms/ops/types"
)

// BuildImage ...
func (p *Opennebula) BuildImage(ctx *lepton.Context) (string, error) {
	return "", nil
}

// BuildImageWithPackage ...
func (p *Opennebula) BuildImageWithPackage(ctx *lepton.Context, pkgpath string) (string, error) {
	return "", nil
}

// CreateImage ...
func (p *Opennebula) CreateImage(ctx *lepton.Context, imagePath string) error {
	return nil
}

// ListImages ...
func (p *Opennebula) ListImages(ctx *lepton.Context, filter string) error {
	return nil
}

// GetImages ...
func (p *Opennebula) GetImages(ctx *lepton.Context, filter string) ([]lepton.CloudImage, error) {
	return nil, nil
}

// DeleteImage ...
func (p *Opennebula) DeleteImage(ctx *lepton.Context, imagename string) error {
	return nil
}

// ResizeImage ...
func (p *Opennebula) ResizeImage(ctx *lepton.Context, imagename string, hbytes string) error {
	return nil
}

// SyncImage ...
func (p *Opennebula) SyncImage(config *types.Config, target lepton.Provider, imagename string) error {
	return nil
}

// CustomizeImage ...
func (p *Opennebula) CustomizeImage(ctx *lepton.Context) (string, error) {
	return "", nil
}

//go:build opennebula || one || !onlyprovider

package opennebula

import (
	"github.com/nanovms/ops/lepton"
	"github.com/nanovms/ops/types"
)

// CreateVolume ...
func (p *Opennebula) CreateVolume(ctx *lepton.Context, cv types.CloudVolume, data, provider string) (lepton.NanosVolume, error) {
	return lepton.NanosVolume{}, nil
}

// GetAllVolumes ...
func (p *Opennebula) GetAllVolumes(ctx *lepton.Context) (*[]lepton.NanosVolume, error) {
	return nil, nil
}

// DeleteVolume ...
func (p *Opennebula) DeleteVolume(ctx *lepton.Context, volumeName string) error {
	return nil
}

// AttachVolume ...
func (p *Opennebula) AttachVolume(ctx *lepton.Context, instanceName, volumeName string, attachID int) error {
	return nil
}

// DetachVolume ...
func (p *Opennebula) DetachVolume(ctx *lepton.Context, instanceName, volumeName string) error {
	return nil
}

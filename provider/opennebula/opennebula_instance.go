//go:build opennebula || one || !onlyprovider

package opennebula

import "github.com/nanovms/ops/lepton"

// CreateInstance ...
func (p *Opennebula) CreateInstance(ctx *lepton.Context) error {
	return nil
}

// ListInstances ...
func (p *Opennebula) ListInstances(ctx *lepton.Context) error {
	return nil
}

// InstanceStats ...
func (p *Opennebula) InstanceStats(ctx *lepton.Context, instancename string, watch bool) error {
	return nil
}

// GetInstances ...
func (p *Opennebula) GetInstances(ctx *lepton.Context) ([]lepton.CloudInstance, error) {
	return nil, nil
}

// GetInstanceByName ...
func (p *Opennebula) GetInstanceByName(ctx *lepton.Context, name string) (*lepton.CloudInstance, error) {
	return nil, nil
}

// DeleteInstance ...
func (p *Opennebula) DeleteInstance(ctx *lepton.Context, instancename string) error {
	return nil
}

// StopInstance ...
func (p *Opennebula) StopInstance(ctx *lepton.Context, instancename string) error {
	return nil
}

// StartInstance ...
func (p *Opennebula) StartInstance(ctx *lepton.Context, instancename string) error {
	return nil
}

// RebootInstance ...
func (p *Opennebula) RebootInstance(ctx *lepton.Context, instancename string) error {
	return nil
}

// GetInstanceLogs ...
func (p *Opennebula) GetInstanceLogs(ctx *lepton.Context, instancename string) (string, error) {
	return "", nil
}

// PrintInstanceLogs ...
func (p *Opennebula) PrintInstanceLogs(ctx *lepton.Context, instancename string, watch bool) error {
	return nil
}

//go:build opennebula || one || !onlyprovider

package opennebula

import (
	"os"

	"github.com/nanovms/ops/types"

	"github.com/OpenNebula/one/src/oca/go/src/goca"
)

// ProviderName ...
const ProviderName = "opennebula"

// Opennebula Provider to interact with opennebula cloud infrastructure
type Opennebula struct {
	Controller *goca.Controller
}

// NewProvider ...
func NewProvider() *Opennebula {
	return &Opennebula{}
}

// Initialize ...
func (p *Opennebula) Initialize(config *types.ProviderConfig) error {
	p.Controller = goca.NewController(
		goca.NewDefaultClient(
			goca.NewConfig(
				os.Getenv("ONE_AUTH_USER"),
				os.Getenv("ONE_AUTH_PASS"),
				os.Getenv("ONE_XMLRPC"),
			),
		),
	)

	return nil
}

package discovery

import (
	eos "github.com/eoscanada/eos-go"
	"github.com/eoscanada/eos-go/ecc"
)

// Discovery is data published autonomously by Block Producer
// candidates. LaunchData will contain the topology and build up the
// graph.  Otherwise, this data is simply metadata about a BP.
type Discovery struct {
	// Testnet is true if this discovery file represents a testing
	// network.
	Testnet bool `json:"testnet"`
	// Mainnet is true if this discovery file represents the main net
	// (or a production network). One of Testnet and Mainnet must be
	// `true`, and are mutually exclusive.
	Mainnet bool `json:"mainnet"`

	EOSIOAccountName      string        `json:"eosio_account_name"`
	EOSIOABPSigningKey    ecc.PublicKey `json:"eosio_appointed_block_producer_signing_key"`
	EOSIOInitialAuthority struct {
		Owner    eos.Authority `json:"owner"`
		Active   eos.Authority `json:"active"`
		Recovery eos.Authority `json:"recovery"`
	} `json:"eosio_initial_authority"`

	OrganizationName string `json:"organization_name"`

	LaunchData LaunchData `json:"launch_data"`

	ClonedFrom string `json:"-"`
}
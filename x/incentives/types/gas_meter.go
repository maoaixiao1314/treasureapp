package types

import (
	"github.com/ethereum/go-ethereum/common"
	treasurenet "github.com/treasurenet/types"
)

// NewGasMeter returns an instance of GasMeter
func NewGasMeter(
	contract common.Address,
	participant common.Address,
	cumulativeGas uint64,
) GasMeter {
	return GasMeter{
		Contract:      contract.String(),
		Participant:   participant.String(),
		CumulativeGas: cumulativeGas,
	}
}

// Validate performs a stateless validation of a Incentive
func (gm GasMeter) Validate() error {
	if err := treasurenet.ValidateAddress(gm.Contract); err != nil {
		return err
	}

	return treasurenet.ValidateAddress(gm.Participant)
}

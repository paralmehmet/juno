package simulation

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"

	appparams "github.com/CosmosContracts/juno/v22/app/params"
	"github.com/CosmosContracts/juno/v22/x/tokenfactory/types"
)

func RandDenomCreationFeeParam(r *rand.Rand) sdk.Coins {
	amount := r.Int63n(10_000_000)
	return sdk.NewCoins(sdk.NewCoin(appparams.BondDenom, sdk.NewInt(amount)))
}

func RandomizedGenState(simstate *module.SimulationState) {
	tfGenesis := types.DefaultGenesis()

	_, err := simstate.Cdc.MarshalJSON(tfGenesis)
	if err != nil {
		panic(err)
	}

	simstate.GenState[types.ModuleName] = simstate.Cdc.MustMarshalJSON(tfGenesis)
}

package keeper

import (
	"github.com/chandiniv1/cosmos-LMS/x/lms/types"
	//"github.com/chandiniv1/cosmos-LMS/x/lms/keeper/keeper.go"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
)

var _ types.ApplyLeaveRequest = Keeper{}

func NewKeeper(
	storeKey storetypes.StoreKey,

	cdc codec.Codec,
) *Keeper {

}

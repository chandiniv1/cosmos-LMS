package keeper

import (
	"github.com/chandiniv1/cosmos-LMS/x/lms/types"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"

	//sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ types.AcceptLeaveRequest=*types.NewAcceptLeaveRequest()

type Keeper struct {
	storeKey storetypes.StoreKey

	cdc codec.Codec
}






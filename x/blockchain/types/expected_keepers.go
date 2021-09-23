package types

import sdk "github.com/cosmos/cosmos-sdk/types"
import "github.com/cosmos/cosmos-sdk/x/bank/types"
// BankKeeper defines the expected bank keepe
type BankKeeper interface {
	ExportGenesis(sdk.Context) *types.GenesisState
	SendCoins(ctx sdk.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) error
	MintCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error
	BurnCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error
	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
}
package keeper

import (
	"context"

	"github.com/cosmonaut/nameservice/x/nameservice/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) SetName(goCtx context.Context, msg *types.MsgSetName) (*types.MsgSetNameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	whois, _ := k.GetWhois(ctx, msg.Name)
	if !(msg.Creator == whois.Creator) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner")
	}
	newWhois := types.Whois{
		Index:   msg.Name,
		Name:    msg.Name,
		Value:   msg.Value,
		Creator: whois.Creator,
		Price:   whois.Price,
	}
	k.SetWhois(ctx, newWhois)
	return &types.MsgSetNameResponse{}, nil
}

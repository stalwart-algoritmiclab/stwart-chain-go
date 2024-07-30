/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/tree/main/LICENSES
 */

package keeper

import (
	"context"
	"strconv"

	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/feepolicy/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateTariffs(goCtx context.Context, msg *types.MsgCreateTariffs) (*types.MsgCreateTariffsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, found := k.Keeper.securedKeeper.GetAddressesByAddress(ctx, msg.Creator)
	if !found {
		return nil, errorsmod.Wrapf(sdkerrors.ErrUnauthorized, "address %s is not allowed", msg.Creator)
	}

	// Check if the value already exists
	tariffs, isFound := k.GetTariffs(
		ctx,
		msg.Denom,
	)
	if !isFound {
		tariffs = types.Tariffs{
			Creator: msg.Creator,
			Denom:   msg.Denom,
			Tariffs: []*types.Tariff{},
		}
	}

	var foundTariff types.Tariff
	tariffID := len(tariffs.Tariffs)
	for _, t := range tariffs.Tariffs {
		// check found tariff for existing fee trigger from amount
		if t.Amount == msg.Tariffs.Amount {
			foundTariff = *t
			for _, fee := range t.Fees {
				if fee.AmountFrom == msg.Tariffs.Fees[0].AmountFrom {
					return nil, errorsmod.Wrapf(sdkerrors.ErrConflict, "amount %s is already set, use UpdateTariffs method", msg.Tariffs.Amount)
				}
			}
		}
	}

	for i, fee := range msg.Tariffs.Fees {
		msg.Tariffs.Fees[i].Id = uint64(i)
		msg.Tariffs.Fees[i].Creator = msg.Creator

		stakeRewardFloat, err := strconv.ParseFloat(fee.StakeReward, 64)
		if err != nil {
			return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "stake reward %s is invalid", fee.StakeReward)
		}

		refRewardFloat, err := strconv.ParseFloat(fee.RefReward, 64)
		if err != nil {
			return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "ref reward %s is invalid", fee.RefReward)
		}

		if stakeRewardFloat+refRewardFloat > 1 {
			return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "sum of stakeReward and refReward greater than 1")
		}
	}

	if foundTariff.Amount != "" {
		foundTariff.Fees = append(foundTariff.Fees, msg.Tariffs.Fees[0])
		tariffs.Tariffs[foundTariff.Id] = &foundTariff
	} else {

		// add tariff
		newEntry := types.Tariff{
			Id:            uint64(tariffID),
			Amount:        msg.Tariffs.Amount,
			Denom:         msg.Tariffs.Denom,
			MinRefBalance: msg.Tariffs.MinRefBalance,
			Fees:          msg.Tariffs.Fees,
		}
		tariffs.Tariffs = append(tariffs.Tariffs, &newEntry)
	}

	k.SetTariffs(
		ctx,
		tariffs,
	)
	return &types.MsgCreateTariffsResponse{}, nil
}

func (k msgServer) UpdateTariffs(goCtx context.Context, msg *types.MsgUpdateTariffs) (*types.MsgUpdateTariffsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, found := k.Keeper.securedKeeper.GetAddressesByAddress(ctx, msg.Creator)
	if !found {
		return nil, errorsmod.Wrapf(sdkerrors.ErrUnauthorized, "address %s is not allowed", msg.Creator)
	}

	// Check if the value exists
	tariffs, isFound := k.GetTariffs(
		ctx,
		msg.Denom,
	)
	if !isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	if msg.Tariffs.Id > uint64(len(tariffs.Tariffs)-1) {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "tariff id is not set")
	}

	tariff := tariffs.Tariffs[msg.Tariffs.Id]
	tariff.Denom = msg.Tariffs.Denom
	tariff.Amount = msg.Tariffs.Amount
	tariff.MinRefBalance = msg.Tariffs.MinRefBalance

	if msg.Tariffs.Fees[0].Id > uint64(len(tariff.Fees)-1) {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "fee id is not set")
	}

	tariff.Fees = msg.Tariffs.Fees

	// for i, fee := range tariff.Fees {
	//	for j, newFee := range msg.Tariffs.Fees {
	//		if fee.Id == newFee.Id {
	//			tariff.Fees[i] = msg.Tariffs.Fees[j]
	//		}
	//	}
	// }

	k.SetTariffs(ctx, tariffs)

	return &types.MsgUpdateTariffsResponse{}, nil
}

func (k msgServer) DeleteTariffs(goCtx context.Context, msg *types.MsgDeleteTariffs) (*types.MsgDeleteTariffsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, found := k.Keeper.securedKeeper.GetAddressesByAddress(ctx, msg.Creator)
	if !found {
		return nil, errorsmod.Wrapf(sdkerrors.ErrUnauthorized, "address %s is not allowed", msg.Creator)
	}

	// Check if the value exists
	tariffs, isFound := k.GetTariffs(
		ctx,
		msg.Denom,
	)
	if !isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	var (
		tariffID uint64
		feeID    uint64
		err      error
	)
	if msg.TariffID != "" {
		tariffID, err = strconv.ParseUint(msg.TariffID, 10, 64)
		if err != nil {
			return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "invalid tariff id (%s)", msg.TariffID)
		}
	}

	if msg.FeeID != "" {
		feeID, err = strconv.ParseUint(msg.FeeID, 10, 64)
		if err != nil {
			return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "invalid fee id (%s)", msg.FeeID)
		}
	}

	isDeleted := false
	if msg.TariffID == "" {
		k.RemoveTariffs(
			ctx,
			msg.Denom,
		)
		isDeleted = true
	} else {
		// find tariff
		for i, t := range tariffs.Tariffs {
			if t.Id == tariffID {
				if msg.FeeID == "" {
					// remove tariff
					tariffs.Tariffs = append(tariffs.Tariffs[:i], tariffs.Tariffs[i+1:]...)

					// update tariff ids
					for i, _ := range tariffs.Tariffs {
						tariffs.Tariffs[i].Id = uint64(i)
					}

					// update tariffs with removed item
					k.SetTariffs(ctx, tariffs)

					isDeleted = true
					break
				} else {
					// find fee
					for j, fee := range t.Fees {
						if fee.Id == feeID {
							// remove fee
							tariffs.Tariffs[i].Fees = append(tariffs.Tariffs[i].Fees[:j], tariffs.Tariffs[i].Fees[j+1:]...)

							// update fee ids
							for index, _ := range tariffs.Tariffs[i].Fees {
								tariffs.Tariffs[i].Fees[index].Id = uint64(index)
							}

							// update tariffs with removed item
							k.SetTariffs(ctx, tariffs)

							isDeleted = true
							break
						}
					}
				}
			}
		}
	}

	if !isDeleted {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "Nothing deleted")
	}

	if err := ctx.EventManager().EmitTypedEvents(msg); err != nil {
		return nil, err
	}

	return &types.MsgDeleteTariffsResponse{}, nil
}

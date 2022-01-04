package keeper

import (
	"fmt"

	abci "github.com/tendermint/tendermint/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/distribution/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

// AllocateTokens handles distribution of the collected fees 处理已收取费用的分配
// bondedVotes is a list of (validator address, validator voted on last block flag) for all  是绑定集中所有验证器的列表（验证器地址，验证器在最后一个块上投票标志）
// validators in the bonded set.
func (k Keeper) AllocateTokens(
	ctx sdk.Context, sumPreviousPrecommitPower, totalPreviousPower int64,
	previousProposer sdk.ConsAddress, bondedVotes []abci.VoteInfo,
) {

	logger := k.Logger(ctx)

	// fetch and clear the collected fees for distribution, since this is 获取并清除收取的分发费用，因为这是在 BeginBlock 中调用，收取的费用将来自前一个块
	// called in BeginBlock, collected fees will be from the previous block
	// (and distributed to the previous proposer) 并分发给先前的投标人
	feeCollector := k.authKeeper.GetModuleAccount(ctx, k.feeCollectorName)
	feesCollectedInt := k.bankKeeper.GetAllBalances(ctx, feeCollector.GetAddress())
	feesCollected := sdk.NewDecCoinsFromCoins(feesCollectedInt...)
	fmt.Println("feeCollector:", feeCollector)
	fmt.Println("feesCollectedInt:", feesCollectedInt)
	fmt.Println("feesCollected:", feesCollected)
	// transfer collected fees to the distribution module account 将收取的费用转入分销模块账户
	fmt.Println("feeCollectorName:", k.feeCollectorName)
	err := k.bankKeeper.SendCoinsFromModuleToModule(ctx, k.feeCollectorName, types.ModuleName, feesCollectedInt)
	if err != nil {
		panic(err)
	}

	// temporary workaround to keep CanWithdrawInvariant happy
	// general discussions here: https://github.com/cosmos/cosmos-sdk/issues/2906#issuecomment-441867634
	feePool := k.GetFeePool(ctx)
	if totalPreviousPower == 0 {
		feePool.CommunityPool = feePool.CommunityPool.Add(feesCollected...)
		k.SetFeePool(ctx, feePool)
		return
	}

	// calculate fraction votes  sumPreviousPrecommitPower 是根据区块中的投票信息计算出来的投票权重的和   totalPreviousPower-》通过LastCommitInfo来统计总的投票权重previousTotalPower
	previousFractionVotes := sdk.NewDec(sumPreviousPrecommitPower).Quo(sdk.NewDec(totalPreviousPower))
	fmt.Println("previousFractionVotes:", previousFractionVotes)
	// calculate previous proposer reward 计算先前的提议者奖励
	baseProposerReward := k.GetBaseProposerReward(ctx) //提案人费率    区块提案者获得当前区块奖励的固定比例作为基础奖励 默认值为%1(根据项目需求设置成%30)
	fmt.Println("baseProposerReward:", baseProposerReward)
	bonusProposerReward := k.GetBonusProposerReward(ctx) //当前分发奖金提议者奖励率    区块提案者的额外奖励--》当所有活跃验证者都进行了投票并且所有的投票都被打包进区块时，区块提案者可以得到的额外奖励比例最大 默认为%4
	fmt.Println("bonusProposerReward:", bonusProposerReward)
	//proposermutiplier这个值是计算额外奖励的，区块提案者的基本奖励比例固定，但额外奖励比例是浮动的，这个方法是计算两部分奖励的比例之和
	//计算公式：proposerMultiplier = baseProposerReward+bonusProposerReward*（sumPrecommitPower/totalPower）sumPrecommitPower/totalPower 这个值就是previousFractionVotes
	proposerMultiplier := baseProposerReward.Add(bonusProposerReward.MulTruncate(previousFractionVotes))
	fmt.Println("proposerMultiplier:", proposerMultiplier)
	//修改proposer的奖励比例 本次区块奖励*Proposer获得比例（30%）
	//proposerMultiplier
	//NewproposerReward := feesCollected.MulDecTruncate(0.300000000000000000)
	proposerReward := feesCollected.MulDecTruncate(proposerMultiplier)
	fmt.Println("proposerReward:", proposerReward)
	// pay previous proposer 向proposer付款
	remaining := feesCollected
	proposerValidator := k.stakingKeeper.ValidatorByConsAddr(ctx, previousProposer)
	//fmt.Println("proposerValidator:", proposerValidator)
	if proposerValidator != nil {
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventTypeProposerReward,
				sdk.NewAttribute(sdk.AttributeKeyAmount, proposerReward.String()),
				sdk.NewAttribute(types.AttributeKeyValidator, proposerValidator.GetOperator().String()),
			),
		)

		k.AllocateTokensToValidator(ctx, proposerValidator, proposerReward)

		remaining = remaining.Sub(proposerReward)
		fmt.Println("proposer后的remaining:", remaining)
	} else {
		// previous proposer can be unknown if say, the unbonding period is 1 block, so
		// e.g. a validator undelegates at block X, it's removed entirely by
		// block X+1's endblock, then X+2 we need to refer to the previous
		// proposer for X+1, but we've forgotten about them.
		logger.Error(fmt.Sprintf(
			"WARNING: Attempt to allocate proposer rewards to unknown proposer %s. "+
				"This should happen only if the proposer unbonded completely within a single block, "+
				"which generally should not happen except in exceptional circumstances (or fuzz testing). "+
				"We recommend you investigate immediately.",
			previousProposer.String()))
	}

	// calculate fraction allocated to validators
	communityTax := k.GetCommunityTax(ctx) //社区税
	tatReward := k.GetTatReward(ctx)       //TAT奖励费率
	fmt.Println("tatReward:", tatReward)
	voteMultiplier := sdk.OneDec().Sub(proposerMultiplier).Sub(communityTax).Sub(tatReward)
	fmt.Println("voteMultiplier:", voteMultiplier)
	// allocate tokens proportionally to voting power 与投票权成比例地分配代币
	// TODO consider parallelizing later, ref https://github.com/cosmos/cosmos-sdk/pull/3099#discussion_r246276376
	var tattotalpower int64
	var newunitallpower int64
	fmt.Printf("bondedVotes:%+v\n", bondedVotes)
	//TAT奖励分配
	for _, vote := range bondedVotes {
		validator := k.stakingKeeper.ValidatorByConsAddr(ctx, vote.Validator.Address)
		TatPower := validator.GetTatPower()
		newtatpower := TatPower.Int64()
		NewUnitPower := validator.GetNewUnitPower()
		NewUnitPower2 := NewUnitPower.Int64()
		tattotalpower += newtatpower
		newunitallpower += NewUnitPower2
	}
	fmt.Println("tattotalpower：", tattotalpower)
	if tattotalpower != int64(0) {
		for _, vote := range bondedVotes {
			validator := k.stakingKeeper.ValidatorByConsAddr(ctx, vote.Validator.Address)
			fmt.Println("voteMultiplier:", voteMultiplier)
			//获取tatpower的总量
			//params := k.stakingKeeper.GetParams(ctx)
			TatPower := validator.GetTatPower()
			newtatpower := TatPower.Int64()
			// fmt.Println("关键测试staking params:", params)
			fmt.Println("关键测试tatpower", TatPower)
			// TODO consider microslashing for missing votes. 考虑对丢失的选票进行微削减。
			// ref https://github.com/cosmos/cosmos-sdk/issues/2525#issuecomment-430838701
			tatpowerFraction := sdk.NewDec(newtatpower).QuoTruncate(sdk.NewDec(tattotalpower))
			fmt.Printf("tatpowerFraction:%+v\n", tatpowerFraction)
			tatreward := feesCollected.MulDecTruncate(tatReward).MulDecTruncate(tatpowerFraction)
			fmt.Printf("reward:%+v\n", tatreward)
			k.AllocateTokensToValidator(ctx, validator, tatreward)
			remaining = remaining.Sub(tatreward)
			fmt.Println("tat分配后remaining:", remaining)
		}
	}
	for _, vote := range bondedVotes {
		validator := k.stakingKeeper.ValidatorByConsAddr(ctx, vote.Validator.Address)
		fmt.Println("voteMultiplier:", voteMultiplier)
		// TODO consider microslashing for missing votes. 考虑对丢失的选票进行微削减。
		// ref https://github.com/cosmos/cosmos-sdk/issues/2525#issuecomment-430838701
		//计算staking的奖励需要将unit的质押的抛除掉，来达到staking和bid的相互隔离
		fmt.Println("奖励测试totalPreviousPower:", totalPreviousPower)
		newunit := validator.GetNewUnitPower().Int64()
		newpower := vote.Validator.Power - newunit
		newtotalPreviousPower := totalPreviousPower - newunitallpower
		//powerFraction := sdk.NewDec(vote.Validator.Power).QuoTruncate(sdk.NewDec(totalPreviousPower))
		powerFraction := sdk.NewDec(newpower).QuoTruncate(sdk.NewDec(newtotalPreviousPower))
		fmt.Printf("powerFraction:%+v\n", powerFraction)
		reward := feesCollected.MulDecTruncate(voteMultiplier).MulDecTruncate(powerFraction)
		fmt.Printf("reward:%+v\n", reward)
		k.AllocateTokensToValidator(ctx, validator, reward)
		remaining = remaining.Sub(reward)
		fmt.Println("validator后remaining的奖励:", remaining)
	}

	// allocate community funding 分配社区资金
	feePool.CommunityPool = feePool.CommunityPool.Add(remaining...)
	k.SetFeePool(ctx, feePool)
}

// AllocateTokensToValidator allocate tokens to a particular validator, splitting according to commission 将令牌分配给特定的验证器，根据佣金进行拆分
func (k Keeper) AllocateTokensToValidator(ctx sdk.Context, val stakingtypes.ValidatorI, tokens sdk.DecCoins) {
	// split tokens between validator and delegators according to commission 根据佣金在验证者和委托者之间分配代币
	commission := tokens.MulDec(val.GetCommission())
	shared := tokens.Sub(commission)
	fmt.Println("佣金:", commission)
	fmt.Println("减去佣金:", shared)
	// update current commission 更新当前佣金
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeCommission,
			sdk.NewAttribute(sdk.AttributeKeyAmount, commission.String()),
			sdk.NewAttribute(types.AttributeKeyValidator, val.GetOperator().String()),
		),
	)
	currentCommission := k.GetValidatorAccumulatedCommission(ctx, val.GetOperator())
	currentCommission.Commission = currentCommission.Commission.Add(commission...)
	k.SetValidatorAccumulatedCommission(ctx, val.GetOperator(), currentCommission)

	// update current rewards  更新当前奖励
	currentRewards := k.GetValidatorCurrentRewards(ctx, val.GetOperator())
	currentRewards.Rewards = currentRewards.Rewards.Add(shared...)
	k.SetValidatorCurrentRewards(ctx, val.GetOperator(), currentRewards)

	// update outstanding rewards  更新杰出奖励
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeRewards,
			sdk.NewAttribute(sdk.AttributeKeyAmount, tokens.String()),
			sdk.NewAttribute(types.AttributeKeyValidator, val.GetOperator().String()),
		),
	)
	outstanding := k.GetValidatorOutstandingRewards(ctx, val.GetOperator())
	outstanding.Rewards = outstanding.Rewards.Add(tokens...)
	k.SetValidatorOutstandingRewards(ctx, val.GetOperator(), outstanding)
}

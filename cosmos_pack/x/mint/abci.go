package mint

import (
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/mint/keeper"
	"github.com/cosmos/cosmos-sdk/x/mint/types"
)

// BeginBlocker mints new tokens for the previous block. BeginBlock为上一个区块铸造新代币。
func BeginBlocker(ctx sdk.Context, k keeper.Keeper) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)

	// fetch stored minter & params
	minter := k.GetMinter(ctx)
	params := k.GetParams(ctx)
	fmt.Printf("minter = %+v\n", minter)
	//fmt.Println("params:", params)
	fmt.Printf("params = %+v\n", params)
	// recalculate inflation rate 重新计算通货膨胀率
	totalStakingSupply := k.StakingTokenSupply(ctx)
	gasUsed := ctx.BlockGasMeter().GasConsumed()
	mainGasUsed := ctx.GasMeter().GasConsumed()
	fmt.Println("totalStakingSupply:", totalStakingSupply)
	fmt.Println("gasUsed:", gasUsed)
	fmt.Printf("mainGasUsed:%+v\n", mainGasUsed)
	bondedRatio := k.BondedRatio(ctx)

	/*
	  重新计算区块奖励
	  前两年的区块奖励是固定的5个unit 也就是5*10*18
	  第三年开始  3rd   year， 2nd  年TAT实际年增长率 B = 2nd  TAT 总量 / 1st TAT 总量 ，实际增长率与目标增长率的偏差
	  delta =  B/（1+10%）；当 delta >= 1，不调整block reward；当delta < 1，则 3rd 年 block reward = delta * n0 Unit per block；
	  也就是说第三年开始，将前两年的TAT总量/第一年的TAT总量再除以(1+10%)  如果这个值>=1 则每一个块的奖励还是5unit 如果这个值小于1 则第三年的区块奖励就是5unit*这个值
	*/
	//截止上月累计TAT生产量
	//AccumulateTat := int64(2000000000000000000)
	//unit的奖励
	AccumulateTat := params.PerReward
	fmt.Printf("AccumulateTat:%+v\n", AccumulateTat)
	//累计Unit发放量  TruncateInt64
	//UnitGrant := sdk.NewDec(int64(2000000))
	//UnitGrant := int64(params.UnitGrant)
	//上周日均出块数
	// AfterWeek := 200
	// fmt.Println("UnitGrant:", UnitGrant)
	NewAcc := sdk.NewInt(AccumulateTat)
	// fmt.Println("NewAcc:", NewAcc)
	//当前系统unit上限
	//NewAccumulateTat := int(float32(AccumulateTat) * params.Probability)
	//fmt.Println(reflect.TypeOf(params.Probability))
	fmt.Printf("Probability = %+v\n", params.Probability)
	//Probability := sdk.NewDecWithPrec(1, 2)
	//NewAccumulateTat := float32(AccumulateTat) * Probability
	//fmt.Println("NewAccumulateTat:", NewAccumulateTat)
	fmt.Println("bondedRatio:", bondedRatio)                                          //bondedRatio当前链上资产抵押比例
	minter.TatProbability = minter.NextProbabilityRate(params)                        //TatProbability表示TAT占比
	minter.Inflation = minter.NextInflationRate(params, bondedRatio)                  //inflation字段表示当前区块的年通胀率
	minter.AnnualProvisions = minter.NextAnnualProvisions(params, totalStakingSupply) //AnnualProvisions 表示根据当前区块适用的年通胀率和链上资产总量计算得到的在当前年通胀率下每年新铸造的链上资产数量
	//当前系统Unit上限 = 截止上月累计TAT生产量 * 1%
	NewAccumulateTat := minter.NewNextAnnualProvisions(params, NewAcc)
	minter.NewAnnualProvisions = NewAccumulateTat

	//newann := NewAccumulateTat.TruncateInt64() - UnitGrant
	//fmt.Println("newann:", newann)
	//minter.NewAnnualProvisions = sdk.NewDec(AccumulateTat)
	fmt.Printf("NewAccumulateTat = %+v\n", NewAccumulateTat)
	//NewTwoNumber := 365 * 2 * AfterWeek
	//fmt.Println("NewTwoNumber.unit64:", uint64(NewTwoNumber))
	// mint coins, update supply 铸币并且更新链上的资产总量
	mintedCoin := minter.BlockProvision(params)
	mintedCoins := sdk.NewCoins(mintedCoin)
	//新的链上资产
	NewmintedCoin := minter.NewBlockProvision(params, uint64(params.BlocksPerYear))
	//累计TAT一年的产量 UnitGrant  当区块高度达到一年高度时候就监听合约来获取TAT产量，当达到第二年高度时，就再次监听一次，此时的TAT产量是两年的总产量，两年总产量-第一年的产量(UnitGrant),用来计算第三年的通胀率
	//UnitGrant += NewmintedCoin.Amount.Int64()
	//params.UnitGrant = uint64(UnitGrant)
	//更新累加值
	k.SetParams(ctx, params)
	NewmintedCoins := sdk.NewCoins(NewmintedCoin)
	//minter.UnitGrant = sdk.NewDec(UnitGrant)
	fmt.Println("mintedCoin", mintedCoin)
	fmt.Println("mintedCoins", mintedCoins)
	fmt.Println("NewmintedCoin:", NewmintedCoin)
	fmt.Println("minter.UnitGrant:", minter.UnitGrant)
	// err := k.MintCoins(ctx, mintedCoins)
	// if err != nil {
	// 	panic(err)
	// }
	fmt.Printf("minter(new) = %+v\n", minter)
	k.SetMinter(ctx, minter)
	minternew := k.GetMinter(ctx)
	fmt.Printf("minternew:%+v\n", minternew)
	err := k.MintCoins(ctx, NewmintedCoins)
	if err != nil {
		panic(err)
	}
	// send the minted coins to the fee collector account 将新铸造的链上资产发送到交易费收集账户
	// err = k.AddCollectedFees(ctx, mintedCoins)
	// if err != nil {
	// 	panic(err)
	// }
	err = k.AddCollectedFees(ctx, NewmintedCoins)
	if err != nil {
		panic(err)
	}

	// if mintedCoin.Amount.IsInt64() {
	// 	defer telemetry.ModuleSetGauge(types.ModuleName, float32(mintedCoin.Amount.Int64()), "minted_tokens")
	// }

	if NewmintedCoin.Amount.IsInt64() {
		defer telemetry.ModuleSetGauge(types.ModuleName, float32(NewmintedCoin.Amount.Int64()), "minted_tokens")
	}
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeMint,
			sdk.NewAttribute(types.AttributeKeyBondedRatio, bondedRatio.String()),
			sdk.NewAttribute(types.AttributeKeyInflation, minter.Inflation.String()),
			sdk.NewAttribute(types.AttributeKeyAnnualProvisions, minter.AnnualProvisions.String()),
			sdk.NewAttribute(types.AttributeKeyNewAnnualProvisions, minter.NewAnnualProvisions.String()),
			sdk.NewAttribute(sdk.AttributeKeyAmount, NewmintedCoin.Amount.String()),
		),
	)
}

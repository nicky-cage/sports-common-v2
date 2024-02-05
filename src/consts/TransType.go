package consts

//帐变明细表 parent_id=0的不用管
const (
	TransTypeRecharge        = 1 //存款 <-  balance=0 frozen=0 available=0
	TransTypeRechargeOnline  = 2 //在线存款 <- 存款 balance=1 frozen=0 available=1
	TransTypeRechargeOffline = 3 //离线存款 <- 存款 balance=1 frozen=0 available=1
	TransTypeRechargeDeduct  = 4 //第三方充值手误返还 <- 存款 balance=-1 frozen=0 available=-1
	TransTypeWithdraw        = 5 //提款 <-  balance=0 frozen=1 available=-1

	TransTypeWithdrawFreeze                   = 6  //提现冻结 <- 提款 balance=0 frozen=1 available=-1
	TransTypeWithdrawUnFreeze                 = 7  //提现解冻 <- 提款 balance=0 frozen=-1 available=1
	TransTypeWithdrawOk                       = 39 //提现成功 <- 提款 balance=-1 frozen=0 available=-1
	TransTypeTransfer                         = 8  //转账 <-  balance=0 frozen=0 available=0
	TransTypeTransferIntoGame                 = 9  //转入游戏 <- 转账 balance=-1 frozen=0 available=-1
	TransTypeTransferIntoPlatform             = 10 //转出游戏 <- 转账 balance=1 frozen=0 available=1
	TransTypeTransferIntoGameFreeze           = 11 //主账户转入游戏冻结 <- 转账 balance=0 frozen=1 available=-1
	TransTypeTransferIntoGameUnFreeze         = 40 //主账户转入游戏解冻 <- 转账 balance=0 frozen=-1 available=1
	TransType12                               = 12 //优惠 <-  balance=0 frozen=0 available=0
	TransTypeTransferFirstDeposit             = 13 //首存优惠 <- 优惠 balance=0 frozen=0 available=0
	TransTypeTransferDeposit                  = 14 //存送优惠 <- 优惠 balance=0 frozen=0 available=0
	TransTypeTransferRegistered               = 15 //注册优惠 <- 优惠 balance=0 frozen=0 available=0
	TransTypeTransferRescueBonus              = 16 //救援奖金 <- 优惠 balance=0 frozen=0 available=0
	TransTypeTransferVipPackage               = 17 //VIP礼包 <- 优惠 balance=0 frozen=0 available=0
	TransType18                               = 18 //特别优惠 <- 优惠 balance=0 frozen=0 available=0
	TransType19                               = 19 //邀请朋友优惠 <- 优惠 balance=0 frozen=0 available=0
	TransType20                               = 20 //手动派发奖金 <- 优惠 balance=0 frozen=0 available=0
	TransType21                               = 21 //回归优惠 <- 优惠 balance=0 frozen=0 available=0
	TransTypeAdjustment                       = 22 //调整 <-  balance=0 frozen=0 available=0
	TransTypeAdjustmentActivityPlus           = 25 //活动(加) <- 调整 balance=1 frozen=0 available=1
	TransTypeAdjustmentDividendPlus           = 26 //红利(加) <- 调整 balance=1 frozen=0 available=1
	TransTypeAdjustmentBonusPlus              = 27 //奖金(加) <- 调整 balance=1 frozen=0 available=1
	TransTypeAdjustmentOtherPlus              = 28 //其他(加) <- 调整 balance=1 frozen=0 available=1
	TransTypeRebateBonus                      = 30 //反水红利 <- 红利 balance=1 frozen=0 available=1
	TransTypeVipDividend                      = 31 //VIP红利 <- 红利 balance=1 frozen=0 available=1
	TransTypeOfflineDepositBonus              = 32 //离线存款红利 <- 红利 balance=0 frozen=0 available=0
	TransTypeInviteFriendsBonus               = 33 //邀请好友红利 <- 红利 balance=1 frozen=0 available=1
	TransType34                               = 34 //云闪付红利 <- 红利 balance=0 frozen=0 available=0
	TransType35                               = 35 //活动存款红利 <- 红利 balance=0 frozen=0 available=0
	TransTypeEmptyMoney                       = 36 //钱包清零 <-  balance=0 frozen=0 available=0
	TransType37                               = 37 //中心钱包清零 <- 钱包清零 balance=0 frozen=0 available=0
	TransType38                               = 38 //游戏钱包清零 <- 钱包清零 balance=0 frozen=0 available=0
	TransTypeRechargeOfflineMistake           = 41 //离线存款失败扣款 <- 存款 balance=-1 frozen=0 available=-1
	TransTypeRechargeOfflineKhb               = 42 //快汇宝存款红利 <-红利  balance=0 frozen=0 available=0
	TransTypeRechargeOfflineKhbDuction        = 43 //快汇宝存款红利扣款 <- 红利 balance=-1 frozen=0 available=-1
	TransTypeWithdrawReturn                   = 44 //提现失败退回 <- 提款 balance=0 frozen=-1 available=1
	TransTypeAdjustmentDepositPlus            = 45 //存款(加) <- 调整 balance=1 frozen=0 available=1
	TransTypeAdjustmentWithdrawPlus           = 46 //提款(加) <- 调整 balance=1 frozen=0 available=1
	TransTypeAdjustmentDepositLess            = 47 //存款(减) <- 调整 balance=-1 frozen=0 available=-1
	TransTypeAdjustmentWithdrawLess           = 48 //提款(减) <- 调整 balance=-1 frozen=0 available=-1
	TransTypeAdjustmentActivityLess           = 49 //活动(减) <- 调整 balance=-1 frozen=0 available=-1
	TransTypeAdjustmentDividendLess           = 50 //红利(减) <- 调整 balance=-1 frozen=0 available=-1
	TransTypeAdjustmentBonusLess              = 51 //奖金(减) <- 调整 balance=-1 frozen=0 available=-1
	TransTypeAdjustmentOtherLess              = 52 //其他(减) <- 调整 balance=-1 frozen=0 available=-1
	TransTypeManualPointsPlus                 = 53 //手动上分(加) <- 存款 balance=1 frozen=0 available=1
	TransTypeManualPointsLess                 = 54 //手动下分(减) <- 提款 balance=-1 frozen=0 available=-1
	TransTypeRechargeWithdrawPlus             = 55 //提款失败确认(加) <- 提款 balance=1 frozen=-1 available=1
	TransTypeRechargeWithdrawLess             = 56 //提款成功确认(减) <- 提款 balance=0 frozen=-1 available=0
	TransTypeAdjustmentPlus                   = 57 //调整增加<- 调整 balance=1 frozen=0 available=1
	TransTypeAdjustmentLess                   = 58 //调整减少  <- 提款 balance=-1 frozen=0 available=-1
	TransTypeRechargeAgentWithdrawPlus        = 59 //代理提款失败确认 <- 提款 agent_balance=0 agent_frozen=-1 agent_available=1
	TransTypeRechargeAgentWithdrawLess        = 60 //代理提款成功确认<- 提款 agent_balance=-1 agent_frozen=0 agent_available=-1
	TransTypeAgentWithdraw                    = 61 //代理提款 <-  balance=0 frozen=0 available=0 agent_balance=-1 agent_frozen=1 agent_available=-1
	TransTypeRechargeAgentGrantCommissionPlus = 62 //代理发放佣金 agent_balance=1 agent_frozen=0 agent_available=1

)

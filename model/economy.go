package model

const (
	CompleteProfileRewardAmount          = 100
	SignupRewardAmount                   = 5
	signupGrantFromAmount                = 25
	signupGrantToAmount                  = 100
	SigninRewardAmount                   = 10
	SigninAdditionalDailyRewardAmount    = 5
	SigninMaxAdditionalDailyRewardAmount = 15
	NFTMintCopyFeeGeneral                = 2
	NFTMintCopyFeeForWaxp                = 1
	NFTMintCopyFeeForPolygonETH          = 3
	NewCollectionFee                     = 100
	NFTVoteFee                           = 1
	ReferralSignupBonusRewardAmount      = 50
	ReferrerBonusRewardAmount            = 500

	//SellNftRewardAmount                  = 30
)

func GetRandomSingupGrantAmount() int {
	return GetRandomNumberInRange(signupGrantFromAmount, signupGrantToAmount)
}

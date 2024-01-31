package model

const (
	CompleteProfileRewardAmount          = 100
	SignupRewardAmount                   = 5
	signupGrantFromAmount                = 25
	signupGrantToAmount                  = 100
	SigninRewardAmount                   = 25
	SigninAdditionalDailyRewardAmount    = 5
	SigninMaxAdditionalDailyRewardAmount = 25
	NFTMintCopyFee                       = 1
	NewCollectionFee                     = 100
	NFTVoteFee                           = 1
	//SellNftRewardAmount                  = 30
)

func GetRandomSingupGrantAmount() int {
	return GetRandomNumberInRange(signupGrantFromAmount, signupGrantToAmount)
}

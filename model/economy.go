package model

const (
	CompleteProfileRewardAmount          = 500
	SignupRewardAmount                   = 5
	signupGrantFromAmount                = 25
	signupGrantToAmount                  = 100
	SigninRewardAmount                   = 25
	SigninAdditionalDailyRewardAmount    = 5
	SigninMaxAdditionalDailyRewardAmount = 75
	//SellNftRewardAmount                  = 30
)

func GetRandomSingupGrantAmount() int {
	return GetRandomNumberInRange(signupGrantFromAmount, signupGrantToAmount)
}

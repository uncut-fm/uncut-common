package model

const (
	PLATFORM_TREAUSURY_ID = 0
	TOTAL_TREASURY_ID     = -1
	NULL_TREASURY_ID      = -2

	CompleteProfileRewardAmount          = 100
	SignupRewardAmount                   = 5
	signupGrantFromAmount                = 25
	SignupGrantToAmount                  = 100
	SigninRewardAmount                   = 10
	SigninAdditionalDailyRewardAmount    = 5
	SigninMaxAdditionalDailyRewardAmount = 15
	NFTMintCopyFeeGeneral                = 2
	NFTMintCopyFeeForWaxp                = 2
	NFTMintCopyFeeForPolygonETH          = 3
	NewCollectionFee                     = 100
	NewCollectionFeeForWaxp              = 50
	NewCollectionFeeForPolygon           = 100
	NFTVoteFee                           = 5
	PromoteCuratedNftFee                 = 200
	ReferralSignupBonusRewardAmount      = 100
	ReferrerBonusRewardAmount            = 250

	ArtxUsdRate   = 0.003 // 1 ARTX = 0.003 USD.
	UsdToArtxRate = 333   // 1 USD = 333.33 ARTX.

	//SellNftRewardAmount                  = 30
)

func GetRandomSingupGrantAmount() int {
	return GetRandomNumberInRange(signupGrantFromAmount, SignupGrantToAmount)
}

func ConvertArtxToUsd(artxAmount float64) float64 {
	return artxAmount * ArtxUsdRate
}

func ConvertUsdToArtx(usdAmount float64) float64 {
	return usdAmount * UsdToArtxRate
}

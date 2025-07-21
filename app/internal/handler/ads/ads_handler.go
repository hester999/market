package ads

type AdsHandler struct {
	ads Ads
}

func NewAdsHandler(ads Ads) *AdsHandler {
	return &AdsHandler{
		ads: ads,
	}
}

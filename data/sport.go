package data

type Sport struct {
	ID               uint64 `json:"id"`
	Name             string `json:"name"`
	FeatureAvailable bool   `json:"featureAvailable"`
}

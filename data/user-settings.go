package data

import (
	"pickside/service/db"
	"pickside/service/db/queries"
)

type UserSettings struct {
	ID                    uint64 `json:"-"`
	PreferredSport        string `json:"preferredSport"`
	PreferredLocale       string `json:"preferredLocale"`
	PreferredTheme        string `json:"preferredTheme"`
	PreferredRegion       string `json:"preferredRegion"`
	AllowLocationTracking bool   `json:"allowLocationTracking"`
	ShowAge               bool   `json:"showAge"`
	ShowEmail             bool   `json:"showEmail"`
	ShowPhone             bool   `json:"showPhone"`
	ShowGroups            bool   `json:"showGroups"`
}

func GetMeSettings(userID uint64) (*UserSettings, error) {
	dbInstance := db.GetDB()

	var settings UserSettings

	err := dbInstance.QueryRow(queries.SelectAllUserSettingsWhereUserIDEquals, userID).Scan(
		&settings.AllowLocationTracking,
		&settings.PreferredLocale,
		&settings.PreferredRegion,
		&settings.PreferredSport,
		&settings.PreferredTheme,
		&settings.ShowAge,
		&settings.ShowEmail,
		&settings.ShowGroups,
		&settings.ShowPhone,
	)

	return &settings, err
}

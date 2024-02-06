package data

import (
	"me/pickside/db"
	"me/pickside/db/queries"
)

type UserSettings struct {
	ID                    uint64 `json:"-"`
	PreferredSport        string `json:"preferred_sport"`
	PreferredLocale       string `json:"preferred_locale"`
	PreferredTheme        string `json:"preferred_theme"`
	PreferredRegion       string `json:"preferred_region"`
	AllowLocationTracking bool   `json:"allow_location_tracking"`
	ShowAge               bool   `json:"show_age"`
	ShowEmail             bool   `json:"show_email"`
	ShowPhone             bool   `json:"show_phone"`
	ShowGroups            bool   `json:"show_groups"`
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

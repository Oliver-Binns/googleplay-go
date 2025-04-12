package users

type User struct {
	Name                       string                     `json:"name"`
	Email                      string                     `json:"email"`
	DeveloperAccountPermission []DeveloperLevelPermission `json:"developerAccountPermissions"`
	PrimaryLocale              []Grant                    `json:"grants"`
}

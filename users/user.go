package users

type User struct {
	Name                       string                     `json:"name"`
	Email                      string                     `json:"email"`
	AccessState                AccessState                `json:"accessState"`
	DeveloperAccountPermission []DeveloperLevelPermission `json:"developerAccountPermissions"`
	PrimaryLocale              []Grant                    `json:"grants"`
}

package users

type User struct {
	Name                       string                     `json:"name"`
	Email                      string                     `json:"email"`
	DeveloperAccountPermission []DeveloperLevelPermission `json:"developerAccountPermissions"`
	Partial                    bool                       `json:"partial,omitempty"`
	AccessState                AccessState                `json:"accessState,omitempty"`
	PrimaryLocale              []Grant                    `json:"grants,omitempty"`
}

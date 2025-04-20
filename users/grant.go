package users

type Grant struct {
	Name                string               `json:"name"`
	PackageName         string               `json:"packageName,omitempty"`
	AppLevelPermissions []AppLevelPermission `json:"appLevelPermissions"`
}

package googleplay

type Grant struct {
	Name                string               `json:"name"`
	PackageName         string               `json:"packageName"`
	AppLevelPermissions []AppLevelPermission `json:"appLevelPermissions"`
}

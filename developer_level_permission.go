package googleplay

type DeveloperLevelPermission string

const (
	UnspecifiedDeveloperLevelPermission DeveloperLevelPermission = "DEVELOPER_LEVEL_PERMISSION_UNSPECIFIED"
	CanSeeAllApps                                                = "CAN_SEE_ALL_APPS"
	CanViewFinancialDataGlobal                                   = "CAN_VIEW_FINANCIAL_DATA_GLOBAL"
	CanManagePermissionsGlobal                                   = "CAN_MANAGE_PERMISSIONS_GLOBAL"
	CanEditGamesGlobal                                           = "CAN_EDIT_GAMES_GLOBAL"
	CanPublishGamesGlobal                                        = "CAN_PUBLISH_GAMES_GLOBAL"
	CanReplyToReviewsGlobal                                      = "CAN_REPLY_TO_REVIEWS_GLOBAL"
	CanManagePublicAPKsGlobal                                    = "CAN_MANAGE_PUBLIC_APKS_GLOBAL"
	CanManageTrackAPKsGlobal                                     = "CAN_MANAGE_TRACK_APKS_GLOBAL"
	CanManageTrackUsersGlobal                                    = "CAN_MANAGE_TRACK_USERS_GLOBAL"
	CanManagePublicListingGlobal                                 = "CAN_MANAGE_PUBLIC_LISTING_GLOBAL"
	CanManageDraftAppsGlobal                                     = "CAN_MANAGE_DRAFT_APPS_GLOBAL"
	CanCreateManagedPlayAppsGlobal                               = "CAN_CREATE_MANAGED_PLAY_APPS_GLOBAL"
	CanChangeManagedPlaySettingGlobal                            = "CAN_CHANGE_MANAGED_PLAY_SETTING_GLOBAL"
	CanManageOrdersGlobal                                        = "CAN_MANAGE_ORDERS_GLOBAL"
	CanManageAppContentGlobal                                    = "CAN_MANAGE_APP_CONTENT_GLOBAL"
	CanViewNonFinancialDataGlobal                                = "CAN_VIEW_NON_FINANCIAL_DATA_GLOBAL"
	CanViewAppQualityGlobal                                      = "CAN_VIEW_APP_QUALITY_GLOBAL"
	CanManageDeeplinksGlobal                                     = "CAN_MANAGE_DEEPLINKS_GLOBAL"
)

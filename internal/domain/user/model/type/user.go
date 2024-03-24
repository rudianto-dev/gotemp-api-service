package user

type Status int8

var (
	USER_STATUS_UNDEFINED     Status = 0
	USER_STATUS_PREREGISTERED Status = 1
	USER_STATUS_ACTIVE        Status = 2
	USER_STATUS_SUSPEND       Status = 3
)

type RoleType int8

var (
	USER_ROLE_TYPE_UNDEFINED RoleType = 0
	USER_ROLE_TYPE_OWNER     RoleType = 1
	USER_ROLE_TYPE_OUTLET    RoleType = 2
	USER_ROLE_TYPE_STAFF     RoleType = 3
)

type Create struct {
	Name     string
	Username string
	Status   Status
	RoleType RoleType
}

type Edit struct {
	ID   string
	Name string
}

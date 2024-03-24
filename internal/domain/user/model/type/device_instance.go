package user

type CreateDeviceInstance struct {
	DeviceID   string
	InstanceID string
	UserID     string
}

type UpdateDeviceInstance struct {
	ID         string
	DeviceID   string
	InstanceID string
	UserID     string
}

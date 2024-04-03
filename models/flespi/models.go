package flespi

type GetDevicesInfo struct {
	Result []struct {
		MessagesTTL   int    `json:"messages_ttl"`
		ProtocolID    int    `json:"protocol_id"`
		DeviceTypeID  int    `json:"device_type_id"`
		ID            int    `json:"id"`
		MediaTTL      int    `json:"media_ttl"`
		Name          string `json:"name"`
		Configuration struct {
			Ident           string `json:"ident"`
			SettingsPolling string `json:"settings_polling"`
		} `json:"configuration"`
		Cid            int `json:"cid"`
		MediaRotate    int `json:"media_rotate"`
		MessagesRotate int `json:"messages_rotate"`
	} `json:"result"`
}

type ResponseToClient struct {
	Data WebHookResponse `json:"data"`
}

type WebHookResponse struct {
	ChannelID                  int     `json:"channel.id"`
	CodecID                    int     `json:"codec.id"`
	DeviceID                   int     `json:"device.id"`
	DeviceName                 string  `json:"device.name"`
	DeviceTypeID               int     `json:"device.type.id"`
	Din                        int     `json:"din"`
	Din1                       bool    `json:"din.1"`
	EngineIgnitionStatus       bool    `json:"engine.ignition.status"`
	EventPriorityEnum          int     `json:"event.priority.enum"`
	ExternalPowersourceVoltage float64 `json:"external.powersource.voltage"`
	GnssStateEnum              int     `json:"gnss.state.enum"`
	GnssStatus                 bool    `json:"gnss.status"`
	GsmCellid                  int     `json:"gsm.cellid"`
	GsmLac                     int     `json:"gsm.lac"`
	GsmMcc                     int     `json:"gsm.mcc"`
	GsmMnc                     int     `json:"gsm.mnc"`
	GsmNetworkModeEnum         int     `json:"gsm.network.mode.enum"`
	GsmNetworkRoamingStatus    bool    `json:"gsm.network.roaming.status"`
	GsmOperatorCode            string  `json:"gsm.operator.code"`
	GsmSignalLevel             int     `json:"gsm.signal.level"`
	Ident                      string  `json:"ident"`
	MovementStatus             bool    `json:"movement.status"`
	Peer                       string  `json:"peer"`
	PositionAltitude           int     `json:"position.altitude"`
	PositionDirection          int     `json:"position.direction"`
	PositionHdop               float64 `json:"position.hdop"`
	PositionLatitude           float64 `json:"position.latitude"`
	PositionLongitude          float64 `json:"position.longitude"`
	PositionSatellites         int     `json:"position.satellites"`
	PositionSpeed              int     `json:"position.speed"`
	PositionValid              bool    `json:"position.valid"`
	ProtocolID                 int     `json:"protocol.id"`
	SegmentVehicleMileage      float64 `json:"segment.vehicle.mileage"`
	ServerTimestamp            float64 `json:"server.timestamp"`
	Timestamp                  int     `json:"timestamp"`
	TripStateEnum              int     `json:"trip.state.enum"`
	TripStatus                 bool    `json:"trip.status"`
}

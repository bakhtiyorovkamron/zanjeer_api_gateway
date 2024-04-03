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
	Data WebHookResponseToClient `json:"data"`
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
type WebHookResponseToClient struct {
	ChannelID                  int     `json:"channelId"`
	CodecID                    int     `json:"codecId"`
	DeviceID                   int     `json:"deviceId"`
	DeviceName                 string  `json:"deviceName"`
	DeviceTypeID               int     `json:"deviceTypeId"`
	Din                        int     `json:"din"`
	Din1                       bool    `json:"din1"`
	EngineIgnitionStatus       bool    `json:"engineIgnitionStatus"`
	EventPriorityEnum          int     `json:"eventPriorityEnum"`
	ExternalPowersourceVoltage float64 `json:"externalPowersourceVoltage"`
	GnssStateEnum              int     `json:"gnssStateEnum"`
	GnssStatus                 bool    `json:"gnssStatus"`
	GsmCellid                  int     `json:"gsmCellid"`
	GsmLac                     int     `json:"gsmLac"`
	GsmMcc                     int     `json:"gsmMcc"`
	GsmMnc                     int     `json:"gsmMnc"`
	GsmNetworkModeEnum         int     `json:"gsmNetworkModeEnum"`
	GsmNetworkRoamingStatus    bool    `json:"gsmNetworkRoamingStatus"`
	PositionAltitude           int     `json:"positionAltitude"`
	PositionDirection          int     `json:"positionDirection"`
	PositionHdop               float64 `json:"positionHdop"`
	PositionLatitude           float64 `json:"positionLatitude"`
	PositionLongitude          float64 `json:"positionLongitude"`
	PositionSatellites         int     `json:"positionSatellites"`
	PositionSpeed              int     `json:"positionSpeed"`
	PositionValid              bool    `json:"positionValid"`
	ProtocolID                 int     `json:"protocolId"`
	SegmentVehicleMileage      float64 `json:"segmentVehicleMileage"`
	Timestamp                  int     `json:"timestamp"`
}

package hass

type Device struct {
	Identifiers  []string `json:"identifiers,omitempty"`
	Connections  []string `json:"connections,omitempty"`
	Manufacturer string   `json:"manufacturer,omitempty"`
	Model        string   `json:"model,omitempty"`
	Name         string   `json:"name,omitempty"`
	SwVersion    string   `json:"sw_version,omitempty"`
}

type BasicConfig struct {
	Device   *Device `json:"device,omitempty"`
	UniqueID string  `json:"unique_id,omitempty"`
}

type Cover struct {
	BasicConfig
	CommandTopic   string `json:"command_topic,omitempty"`
	Name           string `json:"name,omitempty"`
	PositionTopic  string `json:"position_topic,omitempty"`
	PositionOpen   int    `json:"position_open"`
	PositionClosed int    `json:"position_closed"`
}

type DInput struct {
	BasicConfig
	Name       string `json:"name"`
	StateTopic string `json:"state_topic"`
}

type HVAC struct {
	BasicConfig
	Name string `json:"name,omitempty"`

	ActionTopic string `json:"action_topic,omitempty"`

	CurrentTemperatureTopic string `json:"current_temperature_topic,omitempty"`
	TemperatreCommandTopic  string `json:"temperature_command_topic,omitempty"`
	//TemperatureUnit         string  `json:"temperature_unit,omitempty"`
	TemperatureStateTopic string   `json:"temperature_state_topic,omitempty"`
	MaxTemp               float64  `json:"max_temp,omitempty"`
	MinTemp               float64  `json:"min_temp,omitempty"`
	TempStep              float64  `json:"temp_step,omitempty"`
	Modes                 []string `json:"modes,omitempty"`
	ModeCommandTopic      string   `json:"mode_command_topic,omitempty"`
}

type Light struct {
	BasicConfig
	CommandTopic           string `json:"command_topic,omitempty"`
	Name                   string `json:"name,omitempty"`
	BrightnessCommandTopic string `json:"brightness_command_topic,omitempty"`
	BrightnessScale        int32  `json:"brightness_scale"`
	BrightnessStateTopic   string `json:"brightness_state_topic,omitempty"`
	OnCommandType          string `json:"on_command_type,omitempty"`
	StateTopic             string `json:"state_topic,omitempty"`
}

//https://www.home-assistant.io/integrations/sensor.mqtt/
type Sensor struct {
	BasicConfig
	Name              string `json:"name,omitempty"`
	UnitOfMeasurement string `json:"unit_of_measurement"`
	Topic             string `json:"state_topic"`
	Icon              string `json:"icon"`
}

//https://www.home-assistant.io/integrations/switch.mqtt/
type Switch struct {
	BasicConfig
	CommandTopic string `json:"command_topic,omitempty"`
	Name         string `json:"name,omitempty"`
}

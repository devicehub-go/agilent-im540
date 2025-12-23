package protocol

type ResetOptions uint8
type Status uint8
type Offset uint8
type BayardCurrent uint8
type AutoOffset uint8
type CTR uint8
type FailureControl bool
type AnodeVoltage uint8
type AmplifierRange uint8
type Resolution uint8
type CathodeVoltage uint8
type EmissionCurrent uint8
type InterfaceBoard uint8
type MainFrequency uint8

const (
	VoltageSupplyErrors   ResetOptions = 0b00000001
	VoltageSupplyWarnings ResetOptions = 0b00000010
	IonivacSupplyErrors   ResetOptions = 0b00000100
	IonivacSupplyWarnings ResetOptions = 0b00001000
	PendingSupplyErrors   ResetOptions = 0b00010000
	EmissionOff           ResetOptions = 0b00100000
	AllErrors             ResetOptions = 0b10000000

	SwitchOff Status = 0
	SwitchOn  Status = 1
	None      Status = 2

	IgnoreOffset    Offset = 0
	ApplyOffset     Offset = 1
	OffsetIsRunning Offset = 2

	BayardAuto  BayardCurrent = 0
	Bayard100uA BayardCurrent = 1
	Bayard1mA   BayardCurrent = 2
	Bayard10mA  BayardCurrent = 3
	BayardNone  BayardCurrent = 4

	OffsetAutoOn   AutoOffset = 0
	OffsetAutoOff  AutoOffset = 1
	OffsetAutoNone AutoOffset = 2

	CTR_0_01_MBAR CTR = 0
	CTR_0_01_TORR CTR = 1
	CTR_0_02_T0RR CTR = 2
	CTR_0_05_TORR CTR = 3
	CTR_0_10_MBAR CTR = 4
	CTR_0_10_TORR CTR = 5
	CTR_0_25_TORR CTR = 6
	CTR_0_50_TORR CTR = 7
	CTR_1_00_MBAR CTR = 8
	CTR_1_00_TORR CTR = 9
	CTR_2_00_TORR CTR = 10
	CTR_10_MBAR   CTR = 11
	CTR_10_TORR   CTR = 12
	CTR_100_MBAR  CTR = 13
	CTR_100_TORR  CTR = 14
	CTR_1000_MBAR CTR = 15
	CTR_1100_MBAR CTR = 16
	CTR_1000_TOR  CTR = 17
	CTR_NONE      CTR = 18

	AuomaticSelection FailureControl = true

	ANODE_AUTO AnodeVoltage = 0
	ANODE_220V AnodeVoltage = 1
	ANODE_480V AnodeVoltage = 2
	ANODE_NONE AnodeVoltage = 3

	AMPL_AUTO  AmplifierRange = 0
	AMPL_100FA AmplifierRange = 1
	AMPL_1PA   AmplifierRange = 2
	AMPL_2PA   AmplifierRange = 11
	AMPL_10PA  AmplifierRange = 3
	AMPL_100PA AmplifierRange = 4
	AMPL_1NA   AmplifierRange = 5
	AMPL_10NA  AmplifierRange = 6
	AMPL_100NA AmplifierRange = 7
	AMPL_1UA   AmplifierRange = 8
	AMPL_10UA  AmplifierRange = 9
	AMPL_100UA AmplifierRange = 10
	AMPL_NONE  AmplifierRange = 12

	RESOLUTION_AUTO  Resolution = 0
	RESOLUTION_6BIT  Resolution = 1
	RESOLUTION_8BIT  Resolution = 2
	RESOLUTION_10BIT Resolution = 3
	RESOLUTION_11BIT Resolution = 4
	RESOLUTION_12BIT Resolution = 5
	RESOLUTION_14BIT Resolution = 6
	RESOLUTION_NONE  Resolution = 7

	CATHODE_AUTO CathodeVoltage = 0
	CATHODE_10V  CathodeVoltage = 1
	CATHODE_20V  CathodeVoltage = 2
	CATHODE_80V  CathodeVoltage = 3
	CATHODE_100V CathodeVoltage = 4
	CATHODE_NONE CathodeVoltage = 5

	EMISSION_AUTO   EmissionCurrent = 0
	EMISSION_0_1_MA EmissionCurrent = 1
	EMISSION_1_MA   EmissionCurrent = 2
	EMISSION_1_6MA  EmissionCurrent = 3
	EMISSION_10MA   EmissionCurrent = 4
	EMISSION_45MA   EmissionCurrent = 5
	EMISSION_90MA   EmissionCurrent = 6
	EMISSION_NONE   EmissionCurrent = 7

	IF540X_AUTO          InterfaceBoard = 0
	IF540X_INSTALLED     InterfaceBoard = 1
	IF540X_NOT_INSTALLED InterfaceBoard = 2
	IF540X_NONE          InterfaceBoard = 3

	MAIN_FREQ_AUTO MainFrequency = 0
	MAIN_FREQ_50HZ MainFrequency = 1
	MAIN_FREQ_60HZ MainFrequency = 2
	MAIN_FREQ_NONE MainFrequency = 3
)

type GlobalError struct {
	Watchdog            bool
	ROM                 bool
	RAM                 bool
	EEPROM              bool
	SPI                 bool
	NewSensor           bool
	HighPressure        bool
	KeyboardShutdown    bool
	Overtemperature     bool
	SensorStatusChanged bool
	Sensor1Status       bool
	Sensor2Status       bool
	Sensor3Status       bool
	Sensor4Status       bool
	PowerSupply         bool
	IonivacPowerSupply  bool
}

type PowerSupplyError struct {
	Error             bool
	AnodeVoltage      bool
	CathodeVoltage    bool
	ReflectorVoltage  bool
	AnodeCurrent      bool
	FilamentVoltage   bool
	FilamentCurrent   bool
	CathodeRegularAbs bool
	CathodeRegularDev bool
}

type VoltageSupplyError struct {
	Plus5VAnalog bool
	Minus5V      bool
	Plus24V      bool
	Plus15V      bool
	Plus5V       bool
	Plus24VCh3   bool
	Plus24VCh4   bool
	Plus24VKL    bool
	Plus5VRS232  bool
	Plus15VVB    bool
	Minus15VVB   bool
}

type Measurement struct {
	Pressure      float64
	MeasurementOk bool
	Underflow     bool
	Overflow      bool
	NoSensor      bool
	SensorError   bool
	EmissionOn    bool
	DegassingOn   bool
	IsSelected    bool
}

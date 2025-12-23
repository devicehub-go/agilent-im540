package protocol

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/devicehub-go/agilent-im540/internal/utils"
)

func parseMeasurements(measurement string) ([]Measurement, error) {
	values := strings.Split(measurement, ",")

	measurements := make([]Measurement, len(values)/2)
	for i := 0; i < len(values); i += 2 {
		status, err := utils.HexToByteArray(values[i])
		if err != nil {
			return nil, err
		}

		if len(status) != 8 {
			return nil, fmt.Errorf("invalid status length: %d", len(status))
		}

		pressure, err := strconv.ParseFloat(values[i+1], 64)
		if err != nil {
			return nil, err
		}

		measurements[i/2] = Measurement{
			Pressure:      pressure,
			MeasurementOk: status[7] == 1,
			Underflow:     status[6] == 1,
			Overflow:      status[5] == 1,
			NoSensor:      status[4] == 1,
			SensorError:   status[3] == 1,
			EmissionOn:    status[2] == 1,
			DegassingOn:   status[1] == 1,
			IsSelected:    status[0] == 1,
		}
	}

	return measurements, nil
}

/*
Gets the degassing status
*/
func (a *AgilentIM540) GetDegassing() (Status, error) {
	response, err := a.Request("GDS")
	if err != nil {
		return None, err
	}

	if response == "1" {
		return SwitchOn, nil
	}

	return SwitchOff, nil
}

/*
Changes the degassing status (0: off and 1: on)
*/
func (a *AgilentIM540) SetDegassing(option Status) error {
	message := fmt.Sprintf("DGS,%d", option)
	response, err := a.Request(message)
	if err != nil {
		return err
	}
	if response != string(option) {
		return fmt.Errorf("was not possible to set the degassing")
	}
	return nil
}

/*
Gets the emission status
*/
func (a *AgilentIM540) GetEmission() (Status, error) {
	response, err := a.Request("EMI")
	if err != nil {
		return None, err
	}

	values := strings.Split(response, ",")
	if values[1] == "1" {
		return SwitchOn, nil
	}

	return SwitchOff, nil
}

/*
Changes the emission status for the desired channel (0: off and 1: on)
*/
func (a *AgilentIM540) SetEmission(channel uint8, option Status) error {
	message := fmt.Sprintf("EMI,%d,%d", channel, option)
	response, err := a.Request(message)

	if err != nil {
		return err
	}
	if response != string(option) {
		return fmt.Errorf("was not possible to set the emission")
	}
	return nil
}

/*
Gets the offset correction status
*/
func (a *AgilentIM540) GetOffset(channel uint8) (Offset, error) {
	message := fmt.Sprintf("OFC,%d", channel)
	response, err := a.Request(message)

	if err != nil {
		return IgnoreOffset, err
	}
	if response == "1" {
		return ApplyOffset, nil
	}
	if response == "2" {
		return OffsetIsRunning, nil
	}

	return IgnoreOffset, nil
}

/*
Sets the status of offset correction
*/
func (a *AgilentIM540) SetOffset(channel uint8, option Offset) error {
	message := fmt.Sprintf("OFC,%d,%d", channel, option)
	response, err := a.Request(message)

	if err != nil {
		return err
	}
	if response != string(option) {
		return fmt.Errorf("was not possible to set the offset")
	}

	return nil
}

/*
Gets the status and pressure of an addressed sensor
*/
func (a *AgilentIM540) GetPressure(channel uint8) (Measurement, error) {
	message := fmt.Sprintf("PRS,%d", channel)
	response, err := a.Request(message)

	if err != nil {
		return Measurement{}, err
	}

	measurements, err := parseMeasurements(response)
	if err != nil {
		return Measurement{}, err
	}

	return measurements[0], nil
}

/*
Gets the status and pressure of all available sensors
*/
func (a *AgilentIM540) GetPressures() ([]Measurement, error) {
	response, err := a.Request("PRX")
	if err != nil {
		return []Measurement{}, err
	}

	return parseMeasurements(response)
}

/*
Turns on the Talk Only mode with the specified period

Parameters:
  - period: the period in seconds (0.1 to 60)
*/
func (a *AgilentIM540) TurnsTalkOnlyOn(period float32) error {
	if period < 0.1 || 60 < period {
		return fmt.Errorf("period must be between 0.1 and 60, got %.1f", period)
	}
	message := fmt.Sprintf("TRA, 0, %.1f", period)
	_, err := a.Request(message)
	return err
}

/*
Turns off the Talk Only mode
*/
func (a *AgilentIM540) TurnsTalkOnlyOff() error {
	response, err := a.Request("TRA, 0, 0")
	if err != nil {
		return err
	}
	if response != "0.0" {
		return fmt.Errorf("was not possible to turn off the Talk Only mode")
	}
	return nil
}

/*
Get the output rate of the Talk Only mode
*/
func (a *AgilentIM540) GetTalkOnlyRate() (float32, error) {
	response, err := a.Request("TRA, 0")
	if err != nil {
		return 0, err
	}
	rate, err := strconv.ParseFloat(response, 32)
	if err != nil {
		return 0, err
	}
	return float32(rate), nil
}

/*
Read pressure in Talk Only mode
*/
func (a *AgilentIM540) ReadTalkOnly() ([]Measurement, error) {
	msg, err := a.Communication.ReadUntil(CRLF)
	if err != nil {
		return []Measurement{}, err
	}

	response := string(msg)
	if strings.Contains(response, NAK+LF) {
		return nil, fmt.Errorf("invalid response: % x", msg)
	}

	values := strings.ReplaceAll(response, CRLF, "")
	return parseMeasurements(values)
}

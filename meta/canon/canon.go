// Package canon provides data types and functions for representing canon makernote values
package canon

// ContinuousDrive is part of the CanonCameraSettings field
type ContinuousDrive int16

func (ccd ContinuousDrive) String() string {
	return mapCanonContinuousDriveString[ccd]
}

var mapCanonContinuousDriveString = map[ContinuousDrive]string{
	0:  "Single",
	1:  "Continuous",
	2:  "Movie",
	3:  "Continuous, Speed Priority",
	4:  "Continuous, Low",
	5:  "Continuous, High",
	6:  "Silent Single",
	9:  "Single, Silent",
	10: "Continuous, Silent",
}

// FocusMode is part of the CanonCameraSettings field
type FocusMode int16

func (fm FocusMode) String() string {
	return mapCanonFocusModeString[fm]
}

var mapCanonFocusModeString = map[FocusMode]string{
	0:   "One-shot AF",
	1:   "AI Servo AF",
	2:   "AI Focus AF",
	3:   "Manual Focus",
	4:   "Single",
	5:   "Continuous",
	6:   "Manual Focus",
	16:  "Pan Focus",
	256: "AF + MF",
	512: "Movie Snap Focus",
	519: "Movie Servo AF",
}

// MeteringMode is part of the CanonCameraSettings field
type MeteringMode int16

func (mm MeteringMode) String() string {
	return mapCanonMeteringModeString[mm]
}

var mapCanonMeteringModeString = map[MeteringMode]string{
	0: "Default",
	1: "Spot",
	2: "Average",
	3: "Evaluative",
	4: "Partial",
	5: "Center-weighted average",
}

// FocusRange is part of the CanonCameraSettings field
type FocusRange int16

func (fr FocusRange) String() string {
	return mapCanonFocusRangeString[fr]
}

var mapCanonFocusRangeString = map[FocusRange]string{
	0: "Manual",
	1: "Auto",
	2: "Not Known",
	3: "Macro",
	4: "Very Close",
	5: "Close	   	",
	6:  "Middle Range",
	7:  "Far Range",
	8:  "Pan Focus",
	9:  "Super Macro",
	10: "Infinity",
}

// ExposureMode is part of the CanonCameraSettings field
type ExposureMode int16

func (em ExposureMode) String() string {
	return mapCanonExposureModeString[em]
}

var mapCanonExposureModeString = map[ExposureMode]string{
	0: "Easy",
	1: "Program AE",
	2: "Shutter speed priority AE",
	3: "Aperture-priority AE",
	4: "Manual",
	5: "Depth-of-field AE",
	6: "M-Dep",
	7: "Bulb",
	8: "Flexible-priority AE",
}

// FocusDistance -
type FocusDistance [2]int16

// NewFocusDistance creates a new FocusDistance with the upper
// and lower limits
func NewFocusDistance(upper, lower uint16) FocusDistance {
	return FocusDistance{int16(upper), int16(lower)}
}

// BracketMode - Canon Makernote Backet Mode
type BracketMode int16

func (bm BracketMode) String() string {
	return mapCanonBracketModeString[bm]
}

// Active - returns true if BracketMode is On
func (bm BracketMode) Active() bool {
	return bm != 0
}

var mapCanonBracketModeString = map[BracketMode]string{
	0: "Off",
	1: "AEB",
	2: "FEB",
	3: "ISO",
	4: "WB",
}

// AESetting - Canon Makernote AutoExposure Setting
type AESetting int16

func (ae AESetting) String() string {
	return mapCanonAESettingString[ae]
}

var mapCanonAESettingString = map[AESetting]string{
	0: "Normal AE",
	1: "Exposure Compensation",
	2: "AE Lock",
	3: "AE Lock + Exposure Compensation",
	4: "No AE",
}

// AFAreaMode - Canon Autofocus Area Mode
type AFAreaMode int16

func (caf AFAreaMode) String() string {
	return mapCanonAFAreaMode[caf]
}

var mapCanonAFAreaMode = map[AFAreaMode]string{
	0:  "Off (Manual Focus)",
	1:  "AF Point Expansion (surround)",
	2:  "Single-point AF",
	4:  "Auto",
	5:  "Face Detect AF",
	6:  "Face + Tracking",
	7:  "Zone AF",
	8:  "AF Point Expansion (4 point)",
	9:  "Spot AF",
	10: "AF Point Expansion (8 point)",
	11: "Flexizone Multi (49 point)",
	12: "Flexizone Multi (9 point)",
	13: "Flexizone Single",
	14: "Large Zone AF",
}

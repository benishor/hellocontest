package core

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/ftl/conval"
	"github.com/ftl/hamradio/callsign"
	"github.com/ftl/hamradio/dxcc"
	"github.com/ftl/hamradio/locator"
)

// QSO contains the details about one radio contact.
type QSO struct {
	Callsign      callsign.Callsign
	Time          time.Time
	Frequency     Frequency
	Band          Band
	Mode          Mode
	MyReport      RST
	MyNumber      QSONumber
	MyExchange    []string
	TheirReport   RST
	TheirNumber   QSONumber
	TheirExchange []string
	LogTimestamp  time.Time
	DXCC          dxcc.Prefix
	Points        int
	Duplicate     bool
}

func (qso *QSO) String() string {
	return fmt.Sprintf("%s|%-10s|%5.0fkHz|%4s|%-4s|%s|%s|%s|%s|%s|%s|%2d|%t", qso.Time.Format("15:04"), qso.Callsign.String(), qso.Frequency/1000.0, qso.Band, qso.Mode, qso.MyReport, qso.MyNumber.String(), strings.Join(qso.MyExchange, " "), qso.TheirReport, qso.TheirNumber.String(), strings.Join(qso.TheirExchange, " "), qso.Points, qso.Duplicate)
}

// Frequency in Hz.
type Frequency float64

func (f Frequency) String() string {
	return fmt.Sprintf("%.0fHz", float64(f))
}

// Band represents an amateur radio band.
type Band string

// All HF bands.
const (
	NoBand   Band = ""
	Band160m Band = "160m"
	Band80m  Band = "80m"
	Band60m  Band = "60m"
	Band40m  Band = "40m"
	Band30m  Band = "30m"
	Band20m  Band = "20m"
	Band17m  Band = "17m"
	Band15m  Band = "15m"
	Band12m  Band = "12m"
	Band10m  Band = "10m"
)

// Bands are all supported HF bands.
var Bands = []Band{Band160m, Band80m, Band60m, Band40m, Band30m, Band20m, Band17m, Band15m, Band12m, Band10m}

func (band *Band) String() string {
	return string(*band)
}

// Mode represents the mode.
type Mode string

// All relevant modes.
const (
	NoMode      Mode = ""
	ModeCW      Mode = "CW"
	ModeSSB     Mode = "SSB"
	ModeFM      Mode = "FM"
	ModeRTTY    Mode = "RTTY"
	ModeDigital Mode = "DIGI"
)

// Modes are all relevant modes.
var Modes = []Mode{ModeCW, ModeSSB, ModeFM, ModeRTTY, ModeDigital}

func (mode *Mode) String() string {
	return string(*mode)
}

// RST represents a signal report using the "Readability/Signalstrength/Tone" system.
type RST string

func (rst *RST) String() string {
	return string(*rst)
}

// QSONumber is the unique number of a QSO in the log, either on my or on their side.
type QSONumber int

func (nr *QSONumber) String() string {
	return fmt.Sprintf("%03d", *nr)
}

// Clock represents a source of the current time.
type Clock interface {
	Now() time.Time
}

// Workmode is either search&pounce or run.
type Workmode int

// All work modes.
const (
	SearchPounce Workmode = iota
	Run
)

// EntryField represents an entry field in the visual part.
type EntryField string

// The entry fields.
const (
	CallsignField EntryField = "callsign"
	BandField     EntryField = "band"
	ModeField     EntryField = "mode"
	OtherField    EntryField = "other"

	myExchangePrefix    string = "myExchange_"
	theirExchangePrefix string = "theirExchange_"
)

func (f EntryField) IsMyExchange() bool {
	return strings.HasPrefix(string(f), myExchangePrefix)
}

func (f EntryField) IsTheirExchange() bool {
	return strings.HasPrefix(string(f), theirExchangePrefix)
}

func IsExchangeField(name string) bool {
	return strings.HasPrefix(name, myExchangePrefix) || strings.HasPrefix(name, theirExchangePrefix)
}

func (f EntryField) ExchangeIndex() int {
	s := string(f)
	var a string
	switch {
	case strings.HasPrefix(s, myExchangePrefix):
		a = s[len(myExchangePrefix):]
	case strings.HasPrefix(s, theirExchangePrefix):
		a = s[len(theirExchangePrefix):]
	default:
		return -1
	}
	result, err := strconv.Atoi(a)
	if err != nil {
		return -1
	}
	return result
}

func (f EntryField) NextExchangeField() EntryField {
	s := string(f)
	var a string
	var prefix string
	switch {
	case strings.HasPrefix(s, myExchangePrefix):
		prefix = myExchangePrefix
		a = s[len(myExchangePrefix):]
	case strings.HasPrefix(s, theirExchangePrefix):
		prefix = theirExchangePrefix
		a = s[len(theirExchangePrefix):]
	default:
		return ""
	}
	i, err := strconv.Atoi(a)
	if err != nil {
		return ""
	}
	return EntryField(prefix + strconv.Itoa(i+1))
}

func MyExchangeField(index int) EntryField {
	return EntryField(fmt.Sprintf("%s%d", myExchangePrefix, index))
}

func TheirExchangeField(index int) EntryField {
	return EntryField(fmt.Sprintf("%s%d", theirExchangePrefix, index))
}

type ExchangeField struct {
	Field            EntryField
	CanContainSerial bool
	Properties       conval.ExchangeField

	Short    string
	Name     string
	Hint     string
	ReadOnly bool
}

func DefinitionsToExchangeFields(fieldDefinitions []conval.ExchangeField, exchangeEntryField func(int) EntryField) []ExchangeField {
	result := make([]ExchangeField, 0, len(fieldDefinitions))
	for i, fieldDefinition := range fieldDefinitions {
		short := strings.Join(fieldDefinition.Strings(), "/")
		field := ExchangeField{
			Field:      exchangeEntryField(i + 1),
			Properties: fieldDefinition,
			Short:      short,
		}
		for _, property := range fieldDefinition {
			if property == conval.SerialNumberProperty {
				field.CanContainSerial = true
			}
		}
		result = append(result, field)
	}
	return result
}

// KeyerValues contains the values that can be used as variables in the keyer templates.
type KeyerValues struct {
	TheirCall  string
	MyNumber   QSONumber
	MyReport   RST
	MyXchange  string
	MyExchange string
}

// AnnotatedCallsign contains a callsign with additional information retrieved from databases and the logbook.
type AnnotatedCallsign struct {
	Callsign         callsign.Callsign
	Assembly         MatchingAssembly
	Duplicate        bool
	Worked           bool
	ExactMatch       bool
	Points           int
	Multis           int
	PredictedXchange string

	Comparable interface{}
	Compare    func(interface{}, interface{}) bool
}

func (c AnnotatedCallsign) LessThan(o AnnotatedCallsign) bool {
	if c.ExactMatch && !o.ExactMatch {
		return true
	}

	if c.Compare == nil {
		return false
	}
	if c.Comparable == nil || o.Comparable == nil {
		return false
	}

	return c.Compare(c.Comparable, o.Comparable)
}

type MatchingOperation int

const (
	Matching MatchingOperation = iota
	Insert
	Delete
	Substitute
	FalseFriend
)

type MatchingPart struct {
	OP    MatchingOperation
	Value string
}

type MatchingAssembly []MatchingPart

func (m MatchingAssembly) String() string {
	var result string
	for _, match := range m {
		if match.OP != Delete {
			result += match.Value
		}
	}
	return result
}

type Settings interface {
	Station() Station
	Contest() Contest
}

type Station struct {
	Callsign callsign.Callsign
	Operator callsign.Callsign
	Locator  locator.Locator
}

type Contest struct {
	Definition             *conval.Definition
	ExchangeValues         []string
	GenerateSerialExchange bool

	MyExchangeFields         []ExchangeField
	MyReportExchangeField    ExchangeField
	MyNumberExchangeField    ExchangeField
	TheirExchangeFields      []ExchangeField
	TheirReportExchangeField ExchangeField
	TheirNumberExchangeField ExchangeField

	CallHistoryFilename string
	CallHistoryField    string

	// deprecated from here
	Name                string
	EnterTheirNumber    bool
	EnterTheirXchange   bool
	RequireTheirXchange bool
	AllowMultiBand      bool
	AllowMultiMode      bool

	SameCountryPoints       int
	SameContinentPoints     int
	SpecificCountryPoints   int
	SpecificCountryPrefixes []string
	OtherPoints             int

	Multis              Multis
	XchangeMultiPattern string
	CountPerBand        bool

	CabrilloQSOTemplate string
}

func (c *Contest) UpdateExchangeFields() {
	c.MyExchangeFields = nil
	c.MyReportExchangeField = ExchangeField{}
	c.MyNumberExchangeField = ExchangeField{}
	c.TheirExchangeFields = nil
	c.TheirReportExchangeField = ExchangeField{}
	c.TheirNumberExchangeField = ExchangeField{}

	if c.Definition == nil {
		return
	}

	fieldDefinitions := c.Definition.ExchangeFields()

	c.MyExchangeFields = DefinitionsToExchangeFields(fieldDefinitions, MyExchangeField)
	for i, field := range c.MyExchangeFields {
		switch {
		case field.Properties.Contains(conval.RSTProperty):
			c.MyReportExchangeField = field
		case field.Properties.Contains(conval.SerialNumberProperty):
			if c.GenerateSerialExchange {
				field.ReadOnly = true
				field.Short = "#"
				field.Hint = "Serial Number"
				c.MyExchangeFields[i] = field
			}
			c.MyNumberExchangeField = field
		}
	}

	c.TheirExchangeFields = DefinitionsToExchangeFields(fieldDefinitions, TheirExchangeField)
	for _, field := range c.TheirExchangeFields {
		switch {
		case field.Properties.Contains(conval.RSTProperty):
			c.TheirReportExchangeField = field
		case field.Properties.Contains(conval.SerialNumberProperty):
			c.TheirNumberExchangeField = field
		}
	}
}

type Multis struct {
	DXCC    bool
	WPX     bool
	Xchange bool
}

type Keyer struct {
	SPMacros  []string
	RunMacros []string
	WPM       int
}

type Score struct {
	ScorePerBand map[Band]BandScore
	TotalScore   BandScore
	OverallScore BandScore
}

func (s Score) String() string {
	buf := bytes.NewBufferString("")
	fmt.Fprintf(buf, "Band SpcQ CtyQ ConQ OthQ Dupe Pts     P/Q  CQ ITU Cty PFX Xch Mult Q/M  Result \n")
	fmt.Fprintf(buf, "-------------------------------------------------------------------------------\n")
	for _, band := range Bands {
		if score, ok := s.ScorePerBand[band]; ok {
			fmt.Fprintf(buf, "%4s %s\n", band, score)
		}
	}
	fmt.Fprintf(buf, "-------------------------------------------------------------------------------\n")
	fmt.Fprintf(buf, "Tot  %s\n", s.TotalScore)
	fmt.Fprintf(buf, "Ovr  %s\n", s.OverallScore)
	return buf.String()
}

type BandScore struct {
	SpecificCountryQSOs int
	SameCountryQSOs     int
	SameContinentQSOs   int
	OtherQSOs           int
	Duplicates          int
	Points              int
	CQZones             int
	ITUZones            int
	DXCCEntities        int
	WPXPrefixes         int
	XchangeValues       int
	Multis              int
}

func (s BandScore) String() string {
	return fmt.Sprintf("%4d %4d %4d %4d %4d %7d %4.1f %2d %3d %3d %3d %3d %4d %4.1f %7d", s.SpecificCountryQSOs, s.SameCountryQSOs, s.SameContinentQSOs, s.OtherQSOs, s.Duplicates, s.Points, s.PointsPerQSO(), s.CQZones, s.ITUZones, s.DXCCEntities, s.WPXPrefixes, s.XchangeValues, s.Multis, s.QSOsPerMulti(), s.Result())
}

func (s *BandScore) Add(other BandScore) {
	s.SpecificCountryQSOs += other.SpecificCountryQSOs
	s.SameCountryQSOs += other.SameCountryQSOs
	s.SameContinentQSOs += other.SameContinentQSOs
	s.OtherQSOs += other.OtherQSOs
	s.Duplicates += other.Duplicates
	s.Points += other.Points
	s.CQZones += other.CQZones
	s.ITUZones += other.ITUZones
	s.DXCCEntities += other.DXCCEntities
	s.WPXPrefixes += other.WPXPrefixes
	s.XchangeValues += other.XchangeValues
	s.Multis += other.Multis
}

func (s *BandScore) QSOs() int {
	return s.SpecificCountryQSOs + s.SameCountryQSOs + s.SameContinentQSOs + s.OtherQSOs + s.Duplicates
}

func (s *BandScore) PointsPerQSO() float64 {
	qsos := s.QSOs()
	if qsos == 0 {
		return 0
	}
	return float64(s.Points) / float64(qsos)
}

func (s *BandScore) QSOsPerMulti() float64 {
	qsos := s.QSOs()
	if qsos == 0 {
		return 0
	}
	if s.Multis == 0 {
		return 0
	}
	return float64(qsos) / float64(s.Multis)
}

func (s *BandScore) Result() int {
	return s.Points * s.Multis
}

// Hour is used as reference to calculate the number of QSOs per hour.
type Hour time.Time

// HourOf returns the given time to the hour.
func HourOf(t time.Time) Hour {
	return Hour(time.Date(
		t.Year(),
		t.Month(),
		t.Day(),
		t.Hour(),
		0,
		0,
		0,
		t.Location(),
	))
}

// QSOsPerHour is the rate of QSOs per one hour
type QSOsPerHour int

// QSOsPerHours contains the complete QSO rate statistic mapping each Hour in the contest to the rate of QSOs within this Hour
type QSOsPerHours map[Hour]QSOsPerHour

// QSORate contains all statistics regarding the rate of QSOs in a contest.
type QSORate struct {
	LastHourRate QSOsPerHour
	Last5MinRate QSOsPerHour
	QSOsPerHours QSOsPerHours
	SinceLastQSO time.Duration
}

func (r QSORate) SinceLastQSOFormatted() string {
	total := int(r.SinceLastQSO.Truncate(time.Second).Seconds())
	hours := int(total / (60 * 60))
	minutes := int(total/60) % 60
	seconds := int(total % 60)
	switch {
	case total < 60:
		return fmt.Sprintf("%2ds", seconds)
	case total < 60*60:
		return fmt.Sprintf("%02d:%02d", minutes, seconds)
	default:
		return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
	}
}

type Service int

const (
	NoService Service = iota
	TCIService
	HamlibService
	CWDaemonService
	DXCCService
	SCPService
	CallHistoryService
)

type ServiceStatusListener interface {
	StatusChanged(service Service, avialable bool)
}

type ServiceStatusListenerFunc func(Service, bool)

func (f ServiceStatusListenerFunc) StatusChanged(service Service, available bool) {
	f(service, available)
}

type AsyncRunner func(func())

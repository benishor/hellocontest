package keyer

import (
	"bytes"
	"log"
	"strings"
	"text/template"

	"github.com/ftl/hamradio/callsign"

	"github.com/ftl/hellocontest/core"
)

// New returns a new Keyer that has no patterns or templates defined yet.
func New(client CWClient, myCall callsign.Callsign, speed int) *Keyer {
	return &Keyer{
		myCall:    myCall,
		speed:     speed,
		patterns:  make(map[int]string),
		templates: make(map[int]*template.Template),
		client:    client,
		values:    noValues}
}

type Keyer struct {
	view   View
	client CWClient
	values KeyerValueProvider

	myCall    callsign.Callsign
	speed     int
	patterns  map[int]string
	templates map[int]*template.Template
}

// View represents the visual parts of the keyer.
type View interface {
	ShowMessage(...interface{})
	SetPattern(int, string)
	SetSpeed(int)
}

// CWClient defines the interface used by the Keyer to output the CW.
type CWClient interface {
	Connect() error
	IsConnected() bool
	Speed(int)
	Send(text string)
	Abort()
}

// KeyerValueProvider provides the variable values for the Keyer templates on demand.
type KeyerValueProvider func() core.KeyerValues

func (k *Keyer) SetView(view View) {
	k.view = view
	for i, pattern := range k.patterns {
		k.view.SetPattern(i, pattern)
	}
	k.view.SetSpeed(k.speed)
}

func (k *Keyer) SetValues(values KeyerValueProvider) {
	k.values = values
}

func (k *Keyer) EnterSpeed(speed int) {
	k.speed = speed
	if !k.client.IsConnected() {
		return
	}
	log.Printf("speed entered: %d", speed)
	k.client.Speed(k.speed)
}

func (k *Keyer) EnterPattern(index int, pattern string) {
	k.patterns[index] = pattern
	var err error
	k.templates[index], err = template.New("").Parse(pattern)
	if err != nil {
		k.view.ShowMessage(err)
	} else {
		k.view.ShowMessage()
	}
}

func (k *Keyer) SetPatterns(patterns []string) {
	for i, pattern := range patterns {
		k.patterns[i] = pattern
		k.templates[i], _ = template.New("").Parse(patterns[i])
		if k.view != nil {
			k.view.SetPattern(i, pattern)
		}
	}
}

func (k *Keyer) GetPattern(index int) string {
	return k.patterns[index]
}

func (k *Keyer) GetText(index int) (string, error) {
	buffer := bytes.NewBufferString("")
	template, ok := k.templates[index]
	if !ok {
		return "", nil
	}
	err := template.Execute(buffer, k.fillins())
	if err != nil {
		return "", err
	}
	return buffer.String(), nil
}

func (k *Keyer) fillins() map[string]string {
	values := k.values()
	return map[string]string{
		"MyCall":    k.myCall.String(),
		"MyReport":  softcut(values.MyReport.String()),
		"MyNumber":  softcut(values.MyNumber.String()),
		"MyXchange": values.MyXchange,
		"TheirCall": values.TheirCall,
	}
}

func (k *Keyer) Send(index int) {
	message, err := k.GetText(index)
	if err != nil {
		k.view.ShowMessage(err)
		return
	}
	k.send(message)
}

func (k *Keyer) SendQuestion(q string) {
	s := strings.TrimSpace(q) + "?"
	k.send(s)
}

func (k *Keyer) send(s string) {
	if !k.client.IsConnected() {
		err := k.client.Connect()
		if err != nil {
			k.view.ShowMessage(err)
			return
		}
		k.client.Speed(k.speed)
	}

	log.Printf("sending %s", s)
	k.client.Send(s)
}

func (k *Keyer) Stop() {
	if !k.client.IsConnected() {
		return
	}
	log.Println("abort sending")
	k.client.Abort()
}

// softcut replaces 0 and 9 with their "cut" counterparts t and n.
func softcut(s string) string {
	cuts := map[string]string{
		"0": "t",
		"9": "n",
	}
	result := s
	for digit, cut := range cuts {
		result = strings.Replace(result, digit, cut, -1)
	}
	return result
}

// cut replaces digits with the "cut" counterparts. (see http://wiki.bavarian-contest-club.de/wiki/Contest-FAQ#Was_sind_.22Cut_Numbers.22.3F)
func cut(s string) string {
	cuts := map[string]string{
		"0": "t",
		"1": "a",
		"2": "u",
		"3": "v",
		"5": "e",
		"7": "g",
		"8": "d",
		"9": "n",
	}
	result := s
	for digit, cut := range cuts {
		result = strings.Replace(result, digit, cut, -1)
	}
	return result
}

func noValues() core.KeyerValues {
	return core.KeyerValues{}
}

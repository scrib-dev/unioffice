// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package pml

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/scrib-dev/unioffice"
	"github.com/scrib-dev/unioffice/schema/soo/dml"
)

type CT_TransitionStartSoundAction struct {
	// Loop Sound
	LoopAttr *bool
	// Sound
	Snd *dml.CT_EmbeddedWAVAudioFile
}

func NewCT_TransitionStartSoundAction() *CT_TransitionStartSoundAction {
	ret := &CT_TransitionStartSoundAction{}
	ret.Snd = dml.NewCT_EmbeddedWAVAudioFile()
	return ret
}

func (m *CT_TransitionStartSoundAction) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.LoopAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "loop"},
			Value: fmt.Sprintf("%d", b2i(*m.LoopAttr))})
	}
	e.EncodeToken(start)
	sesnd := xml.StartElement{Name: xml.Name{Local: "p:snd"}}
	e.EncodeElement(m.Snd, sesnd)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TransitionStartSoundAction) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.Snd = dml.NewCT_EmbeddedWAVAudioFile()
	for _, attr := range start.Attr {
		if attr.Name.Local == "loop" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.LoopAttr = &parsed
			continue
		}
	}
lCT_TransitionStartSoundAction:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "snd"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "snd"}:
				if err := d.DecodeElement(m.Snd, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_TransitionStartSoundAction %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_TransitionStartSoundAction
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_TransitionStartSoundAction and its children
func (m *CT_TransitionStartSoundAction) Validate() error {
	return m.ValidateWithPath("CT_TransitionStartSoundAction")
}

// ValidateWithPath validates the CT_TransitionStartSoundAction and its children, prefixing error messages with path
func (m *CT_TransitionStartSoundAction) ValidateWithPath(path string) error {
	if err := m.Snd.ValidateWithPath(path + "/Snd"); err != nil {
		return err
	}
	return nil
}

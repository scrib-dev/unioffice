// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package wml

import (
	"encoding/xml"
	"fmt"

	"github.com/scrib-dev/unioffice"
)

type WdCT_PosH struct {
	RelativeFromAttr WdST_RelFromH
	Choice           *WdCT_PosHChoice
}

func NewWdCT_PosH() *WdCT_PosH {
	ret := &WdCT_PosH{}
	ret.RelativeFromAttr = WdST_RelFromH(1)
	ret.Choice = NewWdCT_PosHChoice()
	return ret
}

func (m *WdCT_PosH) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	attr, err := m.RelativeFromAttr.MarshalXMLAttr(xml.Name{Local: "relativeFrom"})
	if err != nil {
		return err
	}
	start.Attr = append(start.Attr, attr)
	e.EncodeToken(start)
	m.Choice.MarshalXML(e, xml.StartElement{})
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *WdCT_PosH) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.RelativeFromAttr = WdST_RelFromH(1)
	m.Choice = NewWdCT_PosHChoice()
	for _, attr := range start.Attr {
		if attr.Name.Local == "relativeFrom" {
			m.RelativeFromAttr.UnmarshalXMLAttr(attr)
			continue
		}
	}
lWdCT_PosH:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "align"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/wordprocessingDrawing", Local: "align"}:
				m.Choice = NewWdCT_PosHChoice()
				if err := d.DecodeElement(&m.Choice.Align, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "posOffset"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/wordprocessingDrawing", Local: "posOffset"}:
				m.Choice = NewWdCT_PosHChoice()
				if err := d.DecodeElement(&m.Choice.PosOffset, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on WdCT_PosH %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lWdCT_PosH
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the WdCT_PosH and its children
func (m *WdCT_PosH) Validate() error {
	return m.ValidateWithPath("WdCT_PosH")
}

// ValidateWithPath validates the WdCT_PosH and its children, prefixing error messages with path
func (m *WdCT_PosH) ValidateWithPath(path string) error {
	if m.RelativeFromAttr == WdST_RelFromHUnset {
		return fmt.Errorf("%s/RelativeFromAttr is a mandatory field", path)
	}
	if err := m.RelativeFromAttr.ValidateWithPath(path + "/RelativeFromAttr"); err != nil {
		return err
	}
	if err := m.Choice.ValidateWithPath(path + "/Choice"); err != nil {
		return err
	}
	return nil
}

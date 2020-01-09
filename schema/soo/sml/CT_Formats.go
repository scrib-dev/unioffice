// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package sml

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/scrib-dev/unioffice"
)

type CT_Formats struct {
	// Formats Count
	CountAttr *uint32
	// PivotTable Format
	Format []*CT_Format
}

func NewCT_Formats() *CT_Formats {
	ret := &CT_Formats{}
	return ret
}

func (m *CT_Formats) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.CountAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "count"},
			Value: fmt.Sprintf("%v", *m.CountAttr)})
	}
	e.EncodeToken(start)
	seformat := xml.StartElement{Name: xml.Name{Local: "ma:format"}}
	for _, c := range m.Format {
		e.EncodeElement(c, seformat)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Formats) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "count" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.CountAttr = &pt
			continue
		}
	}
lCT_Formats:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "format"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "format"}:
				tmp := NewCT_Format()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Format = append(m.Format, tmp)
			default:
				unioffice.Log("skipping unsupported element on CT_Formats %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Formats
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Formats and its children
func (m *CT_Formats) Validate() error {
	return m.ValidateWithPath("CT_Formats")
}

// ValidateWithPath validates the CT_Formats and its children, prefixing error messages with path
func (m *CT_Formats) ValidateWithPath(path string) error {
	for i, v := range m.Format {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Format[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}

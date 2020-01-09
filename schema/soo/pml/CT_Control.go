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
)

type CT_Control struct {
	ExtLst         *CT_ExtensionList
	Pic            *CT_Picture
	SpidAttr       *string
	NameAttr       *string
	ShowAsIconAttr *bool
	IdAttr         *string
	ImgWAttr       *int32
	ImgHAttr       *int32
}

func NewCT_Control() *CT_Control {
	ret := &CT_Control{}
	return ret
}

func (m *CT_Control) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.SpidAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "spid"},
			Value: fmt.Sprintf("%v", *m.SpidAttr)})
	}
	if m.NameAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "name"},
			Value: fmt.Sprintf("%v", *m.NameAttr)})
	}
	if m.ShowAsIconAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "showAsIcon"},
			Value: fmt.Sprintf("%d", b2i(*m.ShowAsIconAttr))})
	}
	if m.IdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "r:id"},
			Value: fmt.Sprintf("%v", *m.IdAttr)})
	}
	if m.ImgWAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "imgW"},
			Value: fmt.Sprintf("%v", *m.ImgWAttr)})
	}
	if m.ImgHAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "imgH"},
			Value: fmt.Sprintf("%v", *m.ImgHAttr)})
	}
	e.EncodeToken(start)
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "p:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	if m.Pic != nil {
		sepic := xml.StartElement{Name: xml.Name{Local: "p:pic"}}
		e.EncodeElement(m.Pic, sepic)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Control) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Space == "http://schemas.openxmlformats.org/officeDocument/2006/relationships" && attr.Name.Local == "id" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.IdAttr = &parsed
			continue
		}
		if attr.Name.Local == "spid" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.SpidAttr = &parsed
			continue
		}
		if attr.Name.Local == "name" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.NameAttr = &parsed
			continue
		}
		if attr.Name.Local == "showAsIcon" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowAsIconAttr = &parsed
			continue
		}
		if attr.Name.Local == "imgW" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := int32(parsed)
			m.ImgWAttr = &pt
			continue
		}
		if attr.Name.Local == "imgH" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := int32(parsed)
			m.ImgHAttr = &pt
			continue
		}
	}
lCT_Control:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "extLst"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "extLst"}:
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "pic"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "pic"}:
				m.Pic = NewCT_Picture()
				if err := d.DecodeElement(m.Pic, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_Control %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Control
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Control and its children
func (m *CT_Control) Validate() error {
	return m.ValidateWithPath("CT_Control")
}

// ValidateWithPath validates the CT_Control and its children, prefixing error messages with path
func (m *CT_Control) ValidateWithPath(path string) error {
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	if m.Pic != nil {
		if err := m.Pic.ValidateWithPath(path + "/Pic"); err != nil {
			return err
		}
	}
	if m.ImgWAttr != nil {
		if *m.ImgWAttr < 0 {
			return fmt.Errorf("%s/m.ImgWAttr must be >= 0 (have %v)", path, *m.ImgWAttr)
		}
	}
	if m.ImgHAttr != nil {
		if *m.ImgHAttr < 0 {
			return fmt.Errorf("%s/m.ImgHAttr must be >= 0 (have %v)", path, *m.ImgHAttr)
		}
	}
	return nil
}

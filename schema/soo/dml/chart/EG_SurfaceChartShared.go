// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package chart

import (
	"encoding/xml"
	"fmt"

	"github.com/scrib-dev/unioffice"
)

type EG_SurfaceChartShared struct {
	Wireframe *CT_Boolean
	Ser       []*CT_SurfaceSer
	BandFmts  *CT_BandFmts
}

func NewEG_SurfaceChartShared() *EG_SurfaceChartShared {
	ret := &EG_SurfaceChartShared{}
	return ret
}

func (m *EG_SurfaceChartShared) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.Wireframe != nil {
		sewireframe := xml.StartElement{Name: xml.Name{Local: "c:wireframe"}}
		e.EncodeElement(m.Wireframe, sewireframe)
	}
	if m.Ser != nil {
		seser := xml.StartElement{Name: xml.Name{Local: "c:ser"}}
		for _, c := range m.Ser {
			e.EncodeElement(c, seser)
		}
	}
	if m.BandFmts != nil {
		sebandFmts := xml.StartElement{Name: xml.Name{Local: "c:bandFmts"}}
		e.EncodeElement(m.BandFmts, sebandFmts)
	}
	return nil
}

func (m *EG_SurfaceChartShared) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lEG_SurfaceChartShared:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "wireframe"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "wireframe"}:
				m.Wireframe = NewCT_Boolean()
				if err := d.DecodeElement(m.Wireframe, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "ser"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "ser"}:
				tmp := NewCT_SurfaceSer()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Ser = append(m.Ser, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "bandFmts"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "bandFmts"}:
				m.BandFmts = NewCT_BandFmts()
				if err := d.DecodeElement(m.BandFmts, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on EG_SurfaceChartShared %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lEG_SurfaceChartShared
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the EG_SurfaceChartShared and its children
func (m *EG_SurfaceChartShared) Validate() error {
	return m.ValidateWithPath("EG_SurfaceChartShared")
}

// ValidateWithPath validates the EG_SurfaceChartShared and its children, prefixing error messages with path
func (m *EG_SurfaceChartShared) ValidateWithPath(path string) error {
	if m.Wireframe != nil {
		if err := m.Wireframe.ValidateWithPath(path + "/Wireframe"); err != nil {
			return err
		}
	}
	for i, v := range m.Ser {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Ser[%d]", path, i)); err != nil {
			return err
		}
	}
	if m.BandFmts != nil {
		if err := m.BandFmts.ValidateWithPath(path + "/BandFmts"); err != nil {
			return err
		}
	}
	return nil
}

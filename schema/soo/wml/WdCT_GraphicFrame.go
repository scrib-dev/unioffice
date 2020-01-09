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

	"github.com/scrib-dev/unioffice"
	"github.com/scrib-dev/unioffice/schema/soo/dml"
)

type WdCT_GraphicFrame struct {
	CNvPr   *dml.CT_NonVisualDrawingProps
	CNvFrPr *dml.CT_NonVisualGraphicFrameProperties
	Xfrm    *dml.CT_Transform2D
	Graphic *dml.Graphic
	ExtLst  *dml.CT_OfficeArtExtensionList
}

func NewWdCT_GraphicFrame() *WdCT_GraphicFrame {
	ret := &WdCT_GraphicFrame{}
	ret.CNvPr = dml.NewCT_NonVisualDrawingProps()
	ret.CNvFrPr = dml.NewCT_NonVisualGraphicFrameProperties()
	ret.Xfrm = dml.NewCT_Transform2D()
	ret.Graphic = dml.NewGraphic()
	return ret
}

func (m *WdCT_GraphicFrame) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	secNvPr := xml.StartElement{Name: xml.Name{Local: "wp:cNvPr"}}
	e.EncodeElement(m.CNvPr, secNvPr)
	secNvFrPr := xml.StartElement{Name: xml.Name{Local: "wp:cNvFrPr"}}
	e.EncodeElement(m.CNvFrPr, secNvFrPr)
	sexfrm := xml.StartElement{Name: xml.Name{Local: "wp:xfrm"}}
	e.EncodeElement(m.Xfrm, sexfrm)
	segraphic := xml.StartElement{Name: xml.Name{Local: "a:graphic"}}
	e.EncodeElement(m.Graphic, segraphic)
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "wp:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *WdCT_GraphicFrame) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CNvPr = dml.NewCT_NonVisualDrawingProps()
	m.CNvFrPr = dml.NewCT_NonVisualGraphicFrameProperties()
	m.Xfrm = dml.NewCT_Transform2D()
	m.Graphic = dml.NewGraphic()
lWdCT_GraphicFrame:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "cNvPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/wordprocessingDrawing", Local: "cNvPr"}:
				if err := d.DecodeElement(m.CNvPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "cNvFrPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/wordprocessingDrawing", Local: "cNvFrPr"}:
				if err := d.DecodeElement(m.CNvFrPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "xfrm"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/wordprocessingDrawing", Local: "xfrm"}:
				if err := d.DecodeElement(m.Xfrm, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "graphic"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "graphic"}:
				if err := d.DecodeElement(m.Graphic, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "extLst"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/wordprocessingDrawing", Local: "extLst"}:
				m.ExtLst = dml.NewCT_OfficeArtExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on WdCT_GraphicFrame %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lWdCT_GraphicFrame
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the WdCT_GraphicFrame and its children
func (m *WdCT_GraphicFrame) Validate() error {
	return m.ValidateWithPath("WdCT_GraphicFrame")
}

// ValidateWithPath validates the WdCT_GraphicFrame and its children, prefixing error messages with path
func (m *WdCT_GraphicFrame) ValidateWithPath(path string) error {
	if err := m.CNvPr.ValidateWithPath(path + "/CNvPr"); err != nil {
		return err
	}
	if err := m.CNvFrPr.ValidateWithPath(path + "/CNvFrPr"); err != nil {
		return err
	}
	if err := m.Xfrm.ValidateWithPath(path + "/Xfrm"); err != nil {
		return err
	}
	if err := m.Graphic.ValidateWithPath(path + "/Graphic"); err != nil {
		return err
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}

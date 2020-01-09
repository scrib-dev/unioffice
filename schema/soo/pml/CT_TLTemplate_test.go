// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package pml_test

import (
	"encoding/xml"
	"testing"

	"github.com/scrib-dev/unioffice/schema/soo/pml"
)

func TestCT_TLTemplateConstructor(t *testing.T) {
	v := pml.NewCT_TLTemplate()
	if v == nil {
		t.Errorf("pml.NewCT_TLTemplate must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_TLTemplate should validate: %s", err)
	}
}

func TestCT_TLTemplateMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_TLTemplate()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_TLTemplate()
	xml.Unmarshal(buf, v2)
}

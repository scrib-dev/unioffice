// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package wml_test

import (
	"encoding/xml"
	"testing"

	"github.com/scrib-dev/unioffice/schema/soo/wml"
)

func TestCT_FontsListConstructor(t *testing.T) {
	v := wml.NewCT_FontsList()
	if v == nil {
		t.Errorf("wml.NewCT_FontsList must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_FontsList should validate: %s", err)
	}
}

func TestCT_FontsListMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_FontsList()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_FontsList()
	xml.Unmarshal(buf, v2)
}

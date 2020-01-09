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

func TestEG_RangeMarkupElementsConstructor(t *testing.T) {
	v := wml.NewEG_RangeMarkupElements()
	if v == nil {
		t.Errorf("wml.NewEG_RangeMarkupElements must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.EG_RangeMarkupElements should validate: %s", err)
	}
}

func TestEG_RangeMarkupElementsMarshalUnmarshal(t *testing.T) {
	v := wml.NewEG_RangeMarkupElements()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewEG_RangeMarkupElements()
	xml.Unmarshal(buf, v2)
}

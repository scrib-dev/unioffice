// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package math_test

import (
	"encoding/xml"
	"testing"

	"github.com/scrib-dev/unioffice/schema/soo/ofc/math"
)

func TestCT_StyleConstructor(t *testing.T) {
	v := math.NewCT_Style()
	if v == nil {
		t.Errorf("math.NewCT_Style must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed math.CT_Style should validate: %s", err)
	}
}

func TestCT_StyleMarshalUnmarshal(t *testing.T) {
	v := math.NewCT_Style()
	buf, _ := xml.Marshal(v)
	v2 := math.NewCT_Style()
	xml.Unmarshal(buf, v2)
}

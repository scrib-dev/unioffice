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

func TestCT_TextEffectConstructor(t *testing.T) {
	v := wml.NewCT_TextEffect()
	if v == nil {
		t.Errorf("wml.NewCT_TextEffect must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_TextEffect should validate: %s", err)
	}
}

func TestCT_TextEffectMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_TextEffect()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_TextEffect()
	xml.Unmarshal(buf, v2)
}

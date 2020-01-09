// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package sml_test

import (
	"encoding/xml"
	"testing"

	"github.com/scrib-dev/unioffice/schema/soo/sml"
)

func TestPivotCacheRecordsConstructor(t *testing.T) {
	v := sml.NewPivotCacheRecords()
	if v == nil {
		t.Errorf("sml.NewPivotCacheRecords must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.PivotCacheRecords should validate: %s", err)
	}
}

func TestPivotCacheRecordsMarshalUnmarshal(t *testing.T) {
	v := sml.NewPivotCacheRecords()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewPivotCacheRecords()
	xml.Unmarshal(buf, v2)
}

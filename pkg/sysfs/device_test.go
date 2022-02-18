/*
Copyright 2020 The OpenEBS Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package sysfs

import (
	"reflect"
	"testing"
)

func TestNewSysFsDeviceFromDevPath(t *testing.T) {
	tests := map[string]struct {
		devPath string
		want    *Device
		wantErr bool
	}{
		"empty device path is used": {
			devPath: "",
			want:    nil,
			wantErr: true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := NewSysFsDeviceFromDevPath(tt.devPath)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewSysFsDeviceFromDevPath() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSysFsDeviceFromDevPath() got = %v, want %v", got, tt.want)
			}
		})
	}
}
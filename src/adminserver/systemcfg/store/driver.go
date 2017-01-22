/*
   Copyright (c) 2016 VMware, Inc. All Rights Reserved.
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

package store

import (
	"github.com/vmware/harbor/src/common/models"
)

// Driver defines methods that a configuration store driver must implement
type Driver interface {
	// Name returns a human-readable name of the driver
	Name() string
	// Read reads the configurations from store
	Read() (*models.SystemCfg, error)
	// Write writes the configurations to store
	Write(*models.SystemCfg) error
}
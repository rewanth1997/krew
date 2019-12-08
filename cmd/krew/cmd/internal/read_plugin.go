// Copyright 2019 The Kubernetes Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package internal

import (
	"net/http"

	"sigs.k8s.io/krew/pkg/index"
	"sigs.k8s.io/krew/pkg/index/indexscanner"
)

// readRemotePluginManifest function takes URL as input and returns plugin as output
func readRemotePluginManifest(url string) (index.Plugin, error) {
	resp, err := http.Get(url)
	if err != nil {
		return index.Plugin{}, err
	}
	defer resp.Body.Close()

	plugin, err := indexscanner.DecodePluginFile(resp.Body)
	if err != nil {
		return index.Plugin{}, err
	}

	return plugin, nil
}

// GetPlugin is an abstraction layer for manifest and manifest-url handlers and returns plugin object
func GetPlugin(manifest string, manifestURL string) (index.Plugin, error) {
	var plugin index.Plugin
	var err error

	if manifest != "" {
		plugin, err = indexscanner.ReadPluginFile(manifest)
		if err != nil {
			return index.Plugin{}, err
		}
	}

	if manifestURL != "" {
		plugin, err = readRemotePluginManifest(manifestURL)
		if err != nil {
			return index.Plugin{}, err
		}
	}

	return plugin, nil
}

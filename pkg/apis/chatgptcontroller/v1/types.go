/*
Copyright 2023 uucloud.
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

package v1

import (
	"fmt"
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ChatGPT is a specification for a ChatGPT resource
type ChatGPT struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ChatGPTSpec   `json:"spec"`
	Status ChatGPTStatus `json:"status"`
}

// DeploymentName returns the deployment name of the ChatGPT resource, which is a string composed of "gpt-" and the name of the ChatGPT resource.
func (c ChatGPT) DeploymentName() string {
	return "gpt-" + c.Name
}

// ChatGPTSpec is the spec for a ChatGPT resource
type ChatGPTSpec struct {
	Config      ChatGPTConfig `json:"config"`
	ServiceName string        `json:"service_name"`
	Image       string        `json:"image"`
}

// ChatGPTConfig is the config for ChatGPT
type ChatGPTConfig struct {
	Token        string  `json:"token"`
	Model        string  `json:"model"`
	Temperature  float32 `json:"temperature"`
	MaxToken     int     `json:"max_token"`
	TopP         float32 `json:"top_p"`
	N            int     `json:"n"`
	ContextLimit int     `json:"context_limit"`
}
type KV struct {
	Key   string
	Value string
}

// JsonTags returns a map that contains the JSON tag names and their corresponding
// values for each field in the ChatGPTConfig struct.
func (cfg ChatGPTConfig) JsonTags() []KV {
	ret := []KV{}

	t := reflect.TypeOf(cfg)
	v := reflect.New(t).Elem()
	v.Set(reflect.ValueOf(cfg))

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		jsonKey := field.Tag.Get("json")
		if jsonKey == "" { // Skip fields without JSON tags.
			continue
		}
		value := v.Field(i).Interface()
		ret = append(ret, KV{Key: jsonKey, Value: fmt.Sprintf("%v", value)})
	}

	return ret
}

// ChatGPTStatus is the status for a ChatGPT resource
type ChatGPTStatus struct {
	Status string `json:"status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ChatGPTList is a list of ChatGPT resources
type ChatGPTList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []ChatGPT `json:"items"`
}

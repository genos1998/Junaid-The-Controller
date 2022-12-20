/*
Copyright 2017 The Kubernetes Authors.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Foo is a specification for a Foo resource
type DSCRD struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DSCRDSpec   `json:"spec"`
	Status DSCRDStatus `json:"status"`
}

// FooSpec is the spec for a Foo resource
type DSCRDSpec struct {
	Args        ArgsSpec        `json:"args"`
	DataSources DataScourceSpec `json:"datasources"`
}

type ArgsSpec struct {
	Name      string       `json:"name"`
	Value     string       `json:"value"`
	ValueFrom ValueFromCRD `json:"valueFrom"`
}

type ValueFromCRD struct {
	SecretKeyRef SecretRef `json:"secretKeyRef"`
}

type SecretRef struct {
	Key  string `json:"key"`
	Name string `json:"name"`
}

type DataScourceSpec struct {
	Name     string             `json:"name"`
	Enabled  bool               `json:"enabled"`
	Provider DataSourceProvider `json:"provider"`
}

type DataSourceProvider struct {
	CloudWatch  cloudWatchProperties  `json:"cloudWatch,omitempty"`
	AppDynamics appDynamicProperties  `json:"appdynamics,omitempty"`
	DataDog     dataDogProperties     `json:"datadog,omitempty"`
	Elastic     elasticProperties     `json:"elastic,omitempty"`
	Graphite    graphiteProperties    `json:"graphite,omitempty"`
	GrayLog     graylogproperties     `json:"graylog,omitempty"`
	NewRelic    newRelicProperties    `json:"newRelic,omitempty"`
	Prometheus  prometheusProperties  `json:"prometheus,omitempty"`
	Splunk      splunkProperties      `json:"splunk,omitempty"`
	StackDriver stackDriverProperties `json:"stackDriver,omitempty"`
	SumoLogic   sumoLogicProperties   `json:"sumologic,omitempty"`
	VmwareTanzu vmwareTanzuProperties `json:"vmwareTanzu,omitempty"`
}

type cloudWatchProperties struct {
	AccessKeyId     string `json:"accessKeyId"`
	SecretAccessKey string `json:"secretAccessKey"`
}

type appDynamicProperties struct {
	ControllerHost       string `json:"controllerHost"`
	TemporaryAccessToken string `json:"temporaryAccessToken"`
}

type dataDogProperties struct {
	ApiKey         string `json:"apiKey"`
	ApplicationKey string `json:"applicationKey"`
}

type elasticProperties struct {
	ElasticEndpoint string `json:"elasticEndpoint"`
	ElasticUsername string `json:"elasticUsername"`
	ElasticPassword string `json:"elasticPassword"`
	KibanaEndpoint  string `json:"kibanaEndpoint"`
	KibanaUsername  string `json:"kibanaUsername"`
	KibanaPassword  string `json:"kibanaPassword"`
}

type graphiteProperties struct {
	Endpoint string `json:"endpoint"`
}

type graylogproperties struct {
	Endpoint string `json:"endpoint"`
	Token    string `json:"token"`
}

type newRelicProperties struct {
	AccountId string `json:"accountId"`
	ApiKey    string `json:"apiKey"`
}

type prometheusProperties struct {
	Endpoint string `json:"endpoint"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type splunkProperties struct {
	SplunkUrl          string `json:"splunkUrl"`
	Username           string `json:"username"`
	Password           string `json:"password"`
	SplunkDashboardUrl string `json:"splunkDashboardUrl"`
}

type stackDriverProperties struct {
	EncryptedKey string `json:"encryptedKey"`
}

type sumoLogicProperties struct {
	AccessId  string `json:"accessId"`
	AccessKey string `json:"accessKey"`
	Zone      string `json:"zone"`
}

type vmwareTanzuProperties struct {
	Endpoint string `json:"endpoint"`
	Email    string `json:"email"`
	ApiToken string `json:"apiToken"`
}

// FooStatus is the status for a Foo resource
type DSCRDStatus struct {
	AvailableReplicas int32 `json:"availableReplicas"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// FooList is a list of Foo resources
type DSCRDList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []DSCRD `json:"items"`
}

/*


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
	"database/sql/driver"
	"encoding/json"
	"errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ModelDeploymentSpec defines the desired state of ModelDeployment
type ModelDeploymentSpec struct {
	// Model Docker image
	Image string `json:"image"`
	// ID of Predictor to use
	Predictor string `json:"predictor"`
	// Resources for model deployment
	// The same format like k8s uses for pod resources.
	Resources *ResourceRequirements `json:"resources,omitempty"`
	// Annotations for model pods.
	Annotations map[string]string `json:"annotations,omitempty"`
	// Minimum number of pods for model. By default the min replicas parameter equals 0.
	MinReplicas *int32 `json:"minReplicas,omitempty"`
	// Maximum number of pods for model. By default the max replicas parameter equals 1.
	MaxReplicas *int32 `json:"maxReplicas,omitempty"`
	// Initial delay for liveness probe of model pod
	LivenessProbeInitialDelay *int32 `json:"livenessProbeInitialDelay,omitempty"`
	// Initial delay for readiness probe of model pod
	ReadinessProbeInitialDelay *int32 `json:"readinessProbeInitialDelay,omitempty"`
	// Initial delay for readiness probe of model pod
	RoleName *string `json:"roleName,omitempty"`
	// If pulling of your image requires authorization, then you should specify the connection id
	ImagePullConnectionID *string `json:"imagePullConnID,omitempty"`
	// Node selector for specifying a node pool
	NodeSelector map[string]string `json:"nodeSelector,omitempty"`
}

type ModelDeploymentState string

const (
	ModelDeploymentStateProcessing ModelDeploymentState = "Processing"
	ModelDeploymentStateReady      ModelDeploymentState = "Ready"
	ModelDeploymentStateFailed     ModelDeploymentState = "Failed"
	ModelDeploymentStateDeleting   ModelDeploymentState = "Deleting"
)

// ModelDeploymentStatus defines the observed state of ModelDeployment
type ModelDeploymentStatus struct {
	// The state of a model deployment.
	//   "Processing" - A model was not deployed. Because some parameters of the
	//                  custom resource are wrong. For example, there is not a model
	//                  image in a Docker registry.
	//   "Ready" - A model was deployed successfully.
	State ModelDeploymentState `json:"state,omitempty"`
	// The model k8s deployment name
	Deployment string `json:"deployment,omitempty"`
	// Host header value is a routing key for Istio Ingress
	// to forward a request to appropriate Knative Service
	HostHeader string `json:"hostHeader,omitempty"`
	// Number of available pods
	AvailableReplicas int32 `json:"availableReplicas"`
	// Expected number of pods under current load
	Replicas int32 `json:"replicas"`
	// Time when credentials was updated
	LastCredsUpdatedTime *metav1.Time `json:"lastUpdatedTime,omitempty"`
	// Model name discovered in ModelDeployment
	ModelName string `json:"modelName,omitempty"`
	// Model version discovered in ModelDeployment
	ModelVersion string `json:"modelVersion,omitempty"`
}

func (in ModelDeploymentSpec) Value() (driver.Value, error) {
	return json.Marshal(in)
}

func (in *ModelDeploymentSpec) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	res := json.Unmarshal(b, &in)
	return res
}

func (in ModelDeploymentStatus) Value() (driver.Value, error) {
	return json.Marshal(in)
}

func (in *ModelDeploymentStatus) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	res := json.Unmarshal(b, &in)
	return res
}

// +kubebuilder:object:root=true

// ModelDeployment is the Schema for the modeldeployments API
// +kubebuilder:printcolumn:name="State",type="string",JSONPath=".status.state"
// +kubebuilder:printcolumn:name="Model image",type="string",JSONPath=".spec.image"
// +kubebuilder:printcolumn:name="Service URL",type="string",JSONPath=".status.serviceURL"
// +kubebuilder:printcolumn:name="Available Replicas",type="string",JSONPath=".status.availableReplicas"
// +kubebuilder:resource:shortName=md
type ModelDeployment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ModelDeploymentSpec   `json:"spec,omitempty"`
	Status ModelDeploymentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ModelDeploymentList contains a list of ModelDeployment
type ModelDeploymentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ModelDeployment `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ModelDeployment{}, &ModelDeploymentList{})
}

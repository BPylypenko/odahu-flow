// Package predict_v2 provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package predict_v2

// InferenceErrorResponse defines model for inference_error_response.
type InferenceErrorResponse struct {
	Error *string `json:"error,omitempty"`
}

// InferenceRequest defines model for inference_request.
type InferenceRequest struct {
	Id         *string          `json:"id,omitempty"`
	Inputs     []RequestInput   `json:"inputs"`
	Outputs    *[]RequestOutput `json:"outputs,omitempty"`
	Parameters *Parameters      `json:"parameters,omitempty"`
}

// InferenceResponse defines model for inference_response.
type InferenceResponse struct {
	Id           *string          `json:"id,omitempty"`
	ModelName    string           `json:"model_name"`
	ModelVersion *string          `json:"model_version,omitempty"`
	Outputs      []ResponseOutput `json:"outputs"`
	Parameters   *Parameters      `json:"parameters,omitempty"`
}

// MetadataModelResponse defines model for metadata_model_response.
type MetadataModelResponse struct {
	Inputs   *[]MetadataTensor `json:"inputs,omitempty"`
	Name     string            `json:"name"`
	Outputs  *[]MetadataTensor `json:"outputs,omitempty"`
	Platform string            `json:"platform"`
	Versions *[]string         `json:"versions,omitempty"`
}

// MetadataServerErrorResponse defines model for metadata_server_error_response.
type MetadataServerErrorResponse struct {
	Error string `json:"error"`
}

// MetadataServerResponse defines model for metadata_server_response.
type MetadataServerResponse struct {
	Extensions []string `json:"extensions"`
	Name       string   `json:"name"`
	Version    string   `json:"version"`
}

// MetadataTensor defines model for metadata_tensor.
type MetadataTensor struct {
	Datatype string `json:"datatype"`
	Name     string `json:"name"`
	Shape    []int  `json:"shape"`
}

// Parameters defines model for parameters.
type Parameters map[string]interface{}

// RequestInput defines model for request_input.
type RequestInput struct {
	Data       TensorData  `json:"data"`
	Datatype   string      `json:"datatype"`
	Name       string      `json:"name"`
	Parameters *Parameters `json:"parameters,omitempty"`
	Shape      []int       `json:"shape"`
}

// RequestOutput defines model for request_output.
type RequestOutput struct {
	Name       string      `json:"name"`
	Parameters *Parameters `json:"parameters,omitempty"`
}

// ResponseOutput defines model for response_output.
type ResponseOutput struct {
	Data       TensorData  `json:"data"`
	Datatype   string      `json:"datatype"`
	Name       string      `json:"name"`
	Parameters *Parameters `json:"parameters,omitempty"`
	Shape      []int       `json:"shape"`
}

// TensorData defines model for tensor_data.
type TensorData []interface{}

// PostV2ModelsMODELNAMEVersionsMODELVERSIONInferJSONBody defines parameters for PostV2ModelsMODELNAMEVersionsMODELVERSIONInfer.
type PostV2ModelsMODELNAMEVersionsMODELVERSIONInferJSONBody InferenceRequest

// PostV2ModelsMODELNAMEVersionsMODELVERSIONInferJSONRequestBody defines body for PostV2ModelsMODELNAMEVersionsMODELVERSIONInfer for application/json ContentType.
type PostV2ModelsMODELNAMEVersionsMODELVERSIONInferJSONRequestBody PostV2ModelsMODELNAMEVersionsMODELVERSIONInferJSONBody

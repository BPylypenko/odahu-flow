//
//    Copyright 2019 EPAM Systems
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.
//

package training

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/odahu/odahu-flow/packages/operator/pkg/apis/training"
	"github.com/odahu/odahu-flow/packages/operator/pkg/apiserver/routes"
	"github.com/odahu/odahu-flow/packages/operator/pkg/errors"
	"github.com/odahu/odahu-flow/packages/operator/pkg/utils/filter"
	httputil "github.com/odahu/odahu-flow/packages/operator/pkg/utils/httputil"
	"net/http"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
)

var logTI = logf.Log.WithName("toolchain-integration-controller")

const (
	GetToolchainIntegrationURL    = "/toolchain/integration/:id"
	GetAllToolchainIntegrationURL = "/toolchain/integration"
	CreateToolchainIntegrationURL = "/toolchain/integration"
	UpdateToolchainIntegrationURL = "/toolchain/integration"
	DeleteToolchainIntegrationURL = "/toolchain/integration/:id"
	IDTiURLParam                  = "id"
)

var (
	emptyCache = map[string]int{}
)

type toolchainService interface {
	GetToolchainIntegration(name string) (*training.ToolchainIntegration, error)
	GetToolchainIntegrationList(options ...filter.ListOption) ([]training.ToolchainIntegration, error)
	CreateToolchainIntegration(md *training.ToolchainIntegration) error
	UpdateToolchainIntegration(md *training.ToolchainIntegration) error
	DeleteToolchainIntegration(name string) error
}

type ToolchainIntegrationController struct {
	service   toolchainService
	validator *TiValidator
}

// @Summary Get a ToolchainIntegration
// @Description Get a ToolchainIntegration by id
// @Tags Toolchain
// @Name id
// @Accept  json
// @Produce  json
// @Param id path string true "ToolchainIntegration id"
// @Success 200 {object} training.ToolchainIntegration
// @Failure 404 {object} httputil.HTTPResult
// @Failure 400 {object} httputil.HTTPResult
// @Router /api/v1/toolchain/integration/{id} [get]
func (tic *ToolchainIntegrationController) getToolchainIntegration(c *gin.Context) {
	tiID := c.Param(IDTiURLParam)

	ti, err := tic.service.GetToolchainIntegration(tiID)
	if err != nil {
		logTI.Error(err, fmt.Sprintf("Retrieving %s toolchain integration", tiID))
		c.AbortWithStatusJSON(errors.CalculateHTTPStatusCode(err), httputil.HTTPResult{Message: err.Error()})

		return
	}

	c.JSON(http.StatusOK, ti)
}

// @Summary Get list of ToolchainIntegrations
// @Description Get list of ToolchainIntegrations
// @Tags Toolchain
// @Accept  json
// @Produce  json
// @Param size path int false "Number of entities in a response"
// @Param page path int false "Number of a page"
// @Success 200 {array} training.ToolchainIntegration
// @Failure 400 {object} httputil.HTTPResult
// @Router /api/v1/toolchain/integration [get]
func (tic *ToolchainIntegrationController) getAllToolchainIntegrations(c *gin.Context) {
	size, page, err := routes.URLParamsToFilter(c, nil, emptyCache)
	if err != nil {
		logTI.Error(err, "Malformed url parameters of toolchain integration request")
		c.AbortWithStatusJSON(http.StatusBadRequest, httputil.HTTPResult{Message: err.Error()})

		return
	}

	tiList, err := tic.service.GetToolchainIntegrationList(
		filter.Size(size),
		filter.Page(page),
	)
	if err != nil {
		logTI.Error(err, "Retrieving list of toolchain integrations")
		c.AbortWithStatusJSON(errors.CalculateHTTPStatusCode(err), httputil.HTTPResult{Message: err.Error()})

		return
	}

	c.JSON(http.StatusOK, &tiList)
}

// @Summary Create a ToolchainIntegration
// @Description Create a ToolchainIntegration. Results is created ToolchainIntegration.
// @Param ti body training.ToolchainIntegration true "Create a ToolchainIntegration"
// @Tags Toolchain
// @Accept  json
// @Produce  json
// @Success 201 {object} training.ToolchainIntegration
// @Failure 400 {object} httputil.HTTPResult
// @Router /api/v1/toolchain/integration [post]
func (tic *ToolchainIntegrationController) createToolchainIntegration(c *gin.Context) {
	var ti training.ToolchainIntegration

	if err := c.ShouldBindJSON(&ti); err != nil {
		logTI.Error(err, "JSON binding of toolchain integration is failed")
		c.AbortWithStatusJSON(http.StatusBadRequest, httputil.HTTPResult{Message: err.Error()})

		return
	}

	if err := tic.validator.ValidatesAndSetDefaults(&ti); err != nil {
		logMT.Error(err, fmt.Sprintf("Validation of the toolchain integration is failed: %v", ti))
		c.AbortWithStatusJSON(http.StatusBadRequest, httputil.HTTPResult{Message: err.Error()})

		return
	}

	if err := tic.service.CreateToolchainIntegration(&ti); err != nil {
		logTI.Error(err, fmt.Sprintf("Creation of toolchain integration: %v", ti))
		c.AbortWithStatusJSON(errors.CalculateHTTPStatusCode(err), httputil.HTTPResult{Message: err.Error()})

		return
	}

	c.JSON(http.StatusCreated, ti)
}

// @Summary Update a ToolchainIntegration
// @Description Update a ToolchainIntegration. Results is updated ToolchainIntegration.
// @Param ti body training.ToolchainIntegration true "Update a ToolchainIntegration"
// @Tags Toolchain
// @Accept  json
// @Produce  json
// @Success 200 {object} training.ToolchainIntegration
// @Failure 404 {object} httputil.HTTPResult
// @Failure 400 {object} httputil.HTTPResult
// @Router /api/v1/toolchain/integration [put]
func (tic *ToolchainIntegrationController) updateToolchainIntegration(c *gin.Context) {
	var ti training.ToolchainIntegration

	if err := c.ShouldBindJSON(&ti); err != nil {
		logTI.Error(err, "JSON binding of toolchain integration is failed")
		c.AbortWithStatusJSON(http.StatusBadRequest, httputil.HTTPResult{Message: err.Error()})

		return
	}

	if err := tic.validator.ValidatesAndSetDefaults(&ti); err != nil {
		logMT.Error(err, fmt.Sprintf("Validation of the tollchain integration is failed: %v", ti))
		c.AbortWithStatusJSON(http.StatusBadRequest, httputil.HTTPResult{Message: err.Error()})

		return
	}

	if err := tic.service.UpdateToolchainIntegration(&ti); err != nil {
		logTI.Error(err, fmt.Sprintf("Update of toolchain integration: %v", ti))
		c.AbortWithStatusJSON(errors.CalculateHTTPStatusCode(err), httputil.HTTPResult{Message: err.Error()})

		return
	}

	c.JSON(http.StatusOK, ti)
}

// @Summary Delete a ToolchainIntegration
// @Description Delete a ToolchainIntegration by id
// @Tags Toolchain
// @Name id
// @Accept  json
// @Produce  json
// @Param id path string true "ToolchainIntegration id"
// @Success 200 {object} httputil.HTTPResult
// @Failure 404 {object} httputil.HTTPResult
// @Failure 400 {object} httputil.HTTPResult
// @Router /api/v1/toolchain/integration/{id} [delete]
func (tic *ToolchainIntegrationController) deleteToolchainIntegration(c *gin.Context) {
	tiID := c.Param(IDTiURLParam)

	if err := tic.service.DeleteToolchainIntegration(tiID); err != nil {
		logTI.Error(err, fmt.Sprintf("Deletion of %s toolchain integration is failed", tiID))
		c.AbortWithStatusJSON(errors.CalculateHTTPStatusCode(err), httputil.HTTPResult{Message: err.Error()})

		return
	}

	c.JSON(http.StatusOK, httputil.HTTPResult{Message: fmt.Sprintf("ToolchainIntegration %s was deleted", tiID)})
}

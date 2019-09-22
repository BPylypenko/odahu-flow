/**
 * EDI API
 * This is a EDI server.
 *
 * OpenAPI spec version: 1.0
 * 
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 * Do not edit the class manually.
 */

import * as models from './models';

export interface K8sTrainer {
    /**
     * Connection for training data
     */
    inputData?: Array<models.InputDataBindingDir>;

    /**
     * Model training
     */
    modelTraining?: models.ModelTraining;

    /**
     * Connection for trained model artifact
     */
    outputConn?: models.Connection;

    /**
     * Toolchain integration
     */
    toolchainIntegration?: models.ToolchainIntegration;

    /**
     * Connection for source code
     */
    vcs?: models.Connection;

}
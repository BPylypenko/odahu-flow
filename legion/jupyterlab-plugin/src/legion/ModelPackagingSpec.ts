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

export interface ModelPackagingSpec {
    /**
     * List of arguments. This parameter depends on the specific packaging integration
     */
    arguments?: any;

    /**
     * Training output artifact name
     */
    artifactName?: string;

    /**
     * Image name. Packaging integration image will be used if this parameters is missed
     */
    image?: string;

    /**
     * Packaging integration ID
     */
    integrationName?: string;

    /**
     * Resources for packager container The same format like k8s uses for pod resources.
     */
    resources?: models.ResourceRequirements;

    /**
     * List of targets. This parameter depends on the specific packaging integration
     */
    targets?: Array<models.Target>;

}
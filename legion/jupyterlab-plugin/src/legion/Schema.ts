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

export interface Schema {
    /**
     * Arguments schema
     */
    arguments?: models.JsonSchema;

    /**
     * Targets schema
     */
    targets?: Array<models.TargetSchema>;

}
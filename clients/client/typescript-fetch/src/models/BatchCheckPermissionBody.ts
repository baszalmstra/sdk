/* tslint:disable */
/* eslint-disable */
/**
 * Ory APIs
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v1.15.7
 * Contact: support@ory.sh
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { mapValues } from '../runtime';
import type { Relationship } from './Relationship';
import {
    RelationshipFromJSON,
    RelationshipFromJSONTyped,
    RelationshipToJSON,
} from './Relationship';

/**
 * Batch Check Permission Body
 * @export
 * @interface BatchCheckPermissionBody
 */
export interface BatchCheckPermissionBody {
    /**
     * 
     * @type {Array<Relationship>}
     * @memberof BatchCheckPermissionBody
     */
    tuples?: Array<Relationship>;
}

/**
 * Check if a given object implements the BatchCheckPermissionBody interface.
 */
export function instanceOfBatchCheckPermissionBody(value: object): value is BatchCheckPermissionBody {
    return true;
}

export function BatchCheckPermissionBodyFromJSON(json: any): BatchCheckPermissionBody {
    return BatchCheckPermissionBodyFromJSONTyped(json, false);
}

export function BatchCheckPermissionBodyFromJSONTyped(json: any, ignoreDiscriminator: boolean): BatchCheckPermissionBody {
    if (json == null) {
        return json;
    }
    return {
        
        'tuples': json['tuples'] == null ? undefined : ((json['tuples'] as Array<any>).map(RelationshipFromJSON)),
    };
}

export function BatchCheckPermissionBodyToJSON(value?: BatchCheckPermissionBody | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'tuples': value['tuples'] == null ? undefined : ((value['tuples'] as Array<any>).map(RelationshipToJSON)),
    };
}


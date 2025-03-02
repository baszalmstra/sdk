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
import type { CheckPermissionResultWithError } from './CheckPermissionResultWithError';
import {
    CheckPermissionResultWithErrorFromJSON,
    CheckPermissionResultWithErrorFromJSONTyped,
    CheckPermissionResultWithErrorToJSON,
} from './CheckPermissionResultWithError';

/**
 * Batch Check Permission Result
 * @export
 * @interface BatchCheckPermissionResult
 */
export interface BatchCheckPermissionResult {
    /**
     * An array of check results. The order aligns with the input order.
     * @type {Array<CheckPermissionResultWithError>}
     * @memberof BatchCheckPermissionResult
     */
    results: Array<CheckPermissionResultWithError>;
}

/**
 * Check if a given object implements the BatchCheckPermissionResult interface.
 */
export function instanceOfBatchCheckPermissionResult(value: object): value is BatchCheckPermissionResult {
    if (!('results' in value) || value['results'] === undefined) return false;
    return true;
}

export function BatchCheckPermissionResultFromJSON(json: any): BatchCheckPermissionResult {
    return BatchCheckPermissionResultFromJSONTyped(json, false);
}

export function BatchCheckPermissionResultFromJSONTyped(json: any, ignoreDiscriminator: boolean): BatchCheckPermissionResult {
    if (json == null) {
        return json;
    }
    return {
        
        'results': ((json['results'] as Array<any>).map(CheckPermissionResultWithErrorFromJSON)),
    };
}

export function BatchCheckPermissionResultToJSON(value?: BatchCheckPermissionResult | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'results': ((value['results'] as Array<any>).map(CheckPermissionResultWithErrorToJSON)),
    };
}


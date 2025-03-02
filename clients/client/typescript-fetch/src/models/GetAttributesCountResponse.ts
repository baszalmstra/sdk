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
import type { AttributesCountDatapoint } from './AttributesCountDatapoint';
import {
    AttributesCountDatapointFromJSON,
    AttributesCountDatapointFromJSONTyped,
    AttributesCountDatapointToJSON,
} from './AttributesCountDatapoint';

/**
 * Response of the getAttributesCount endpoint
 * @export
 * @interface GetAttributesCountResponse
 */
export interface GetAttributesCountResponse {
    /**
     * The list of data points.
     * @type {Array<AttributesCountDatapoint>}
     * @memberof GetAttributesCountResponse
     */
    readonly data: Array<AttributesCountDatapoint>;
}

/**
 * Check if a given object implements the GetAttributesCountResponse interface.
 */
export function instanceOfGetAttributesCountResponse(value: object): value is GetAttributesCountResponse {
    if (!('data' in value) || value['data'] === undefined) return false;
    return true;
}

export function GetAttributesCountResponseFromJSON(json: any): GetAttributesCountResponse {
    return GetAttributesCountResponseFromJSONTyped(json, false);
}

export function GetAttributesCountResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): GetAttributesCountResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'data': ((json['data'] as Array<any>).map(AttributesCountDatapointFromJSON)),
    };
}

export function GetAttributesCountResponseToJSON(value?: Omit<GetAttributesCountResponse, 'data'> | null): any {
    if (value == null) {
        return value;
    }
    return {
        
    };
}


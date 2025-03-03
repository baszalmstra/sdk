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
import type { SessionActivityDatapoint } from './SessionActivityDatapoint';
import {
    SessionActivityDatapointFromJSON,
    SessionActivityDatapointFromJSONTyped,
    SessionActivityDatapointToJSON,
} from './SessionActivityDatapoint';

/**
 * Response of the getSessionActivity endpoint
 * @export
 * @interface GetSessionActivityResponse
 */
export interface GetSessionActivityResponse {
    /**
     * The list of data points.
     * @type {Array<SessionActivityDatapoint>}
     * @memberof GetSessionActivityResponse
     */
    readonly data: Array<SessionActivityDatapoint>;
}

/**
 * Check if a given object implements the GetSessionActivityResponse interface.
 */
export function instanceOfGetSessionActivityResponse(value: object): value is GetSessionActivityResponse {
    if (!('data' in value) || value['data'] === undefined) return false;
    return true;
}

export function GetSessionActivityResponseFromJSON(json: any): GetSessionActivityResponse {
    return GetSessionActivityResponseFromJSONTyped(json, false);
}

export function GetSessionActivityResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): GetSessionActivityResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'data': ((json['data'] as Array<any>).map(SessionActivityDatapointFromJSON)),
    };
}

export function GetSessionActivityResponseToJSON(value?: Omit<GetSessionActivityResponse, 'data'> | null): any {
    if (value == null) {
        return value;
    }
    return {
        
    };
}


/* tslint:disable */
/* eslint-disable */
/**
 * Ory APIs
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v1.14.3
 * Contact: support@ory.sh
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { mapValues } from '../runtime';
/**
 * Update Login Flow with TOTP Method
 * @export
 * @interface UpdateLoginFlowWithTotpMethod
 */
export interface UpdateLoginFlowWithTotpMethod {
    /**
     * Sending the anti-csrf token is only required for browser login flows.
     * @type {string}
     * @memberof UpdateLoginFlowWithTotpMethod
     */
    csrf_token?: string;
    /**
     * Method should be set to "totp" when logging in using the TOTP strategy.
     * @type {string}
     * @memberof UpdateLoginFlowWithTotpMethod
     */
    method: string;
    /**
     * The TOTP code.
     * @type {string}
     * @memberof UpdateLoginFlowWithTotpMethod
     */
    totp_code: string;
    /**
     * Transient data to pass along to any webhooks
     * @type {object}
     * @memberof UpdateLoginFlowWithTotpMethod
     */
    transient_payload?: object;
}

/**
 * Check if a given object implements the UpdateLoginFlowWithTotpMethod interface.
 */
export function instanceOfUpdateLoginFlowWithTotpMethod(value: object): value is UpdateLoginFlowWithTotpMethod {
    if (!('method' in value) || value['method'] === undefined) return false;
    if (!('totp_code' in value) || value['totp_code'] === undefined) return false;
    return true;
}

export function UpdateLoginFlowWithTotpMethodFromJSON(json: any): UpdateLoginFlowWithTotpMethod {
    return UpdateLoginFlowWithTotpMethodFromJSONTyped(json, false);
}

export function UpdateLoginFlowWithTotpMethodFromJSONTyped(json: any, ignoreDiscriminator: boolean): UpdateLoginFlowWithTotpMethod {
    if (json == null) {
        return json;
    }
    return {
        
        'csrf_token': json['csrf_token'] == null ? undefined : json['csrf_token'],
        'method': json['method'],
        'totp_code': json['totp_code'],
        'transient_payload': json['transient_payload'] == null ? undefined : json['transient_payload'],
    };
}

export function UpdateLoginFlowWithTotpMethodToJSON(value?: UpdateLoginFlowWithTotpMethod | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'csrf_token': value['csrf_token'],
        'method': value['method'],
        'totp_code': value['totp_code'],
        'transient_payload': value['transient_payload'],
    };
}


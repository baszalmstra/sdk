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
/**
 * Indicates that a session was issued, and the application should use this token for authenticated requests
 * @export
 * @interface ContinueWithSetOrySessionToken
 */
export interface ContinueWithSetOrySessionToken {
    /**
     * Action will always be `set_ory_session_token`
     * set_ory_session_token ContinueWithActionSetOrySessionTokenString
     * @type {string}
     * @memberof ContinueWithSetOrySessionToken
     */
    action: ContinueWithSetOrySessionTokenActionEnum;
    /**
     * Token is the token of the session
     * @type {string}
     * @memberof ContinueWithSetOrySessionToken
     */
    ory_session_token: string;
}


/**
 * @export
 */
export const ContinueWithSetOrySessionTokenActionEnum = {
    SetOrySessionToken: 'set_ory_session_token'
} as const;
export type ContinueWithSetOrySessionTokenActionEnum = typeof ContinueWithSetOrySessionTokenActionEnum[keyof typeof ContinueWithSetOrySessionTokenActionEnum];


/**
 * Check if a given object implements the ContinueWithSetOrySessionToken interface.
 */
export function instanceOfContinueWithSetOrySessionToken(value: object): value is ContinueWithSetOrySessionToken {
    if (!('action' in value) || value['action'] === undefined) return false;
    if (!('ory_session_token' in value) || value['ory_session_token'] === undefined) return false;
    return true;
}

export function ContinueWithSetOrySessionTokenFromJSON(json: any): ContinueWithSetOrySessionToken {
    return ContinueWithSetOrySessionTokenFromJSONTyped(json, false);
}

export function ContinueWithSetOrySessionTokenFromJSONTyped(json: any, ignoreDiscriminator: boolean): ContinueWithSetOrySessionToken {
    if (json == null) {
        return json;
    }
    return {
        
        'action': json['action'],
        'ory_session_token': json['ory_session_token'],
    };
}

export function ContinueWithSetOrySessionTokenToJSON(value?: ContinueWithSetOrySessionToken | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'action': value['action'],
        'ory_session_token': value['ory_session_token'],
    };
}


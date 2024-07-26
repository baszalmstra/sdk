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
 * Used when an administrator creates a recovery code for an identity.
 * @export
 * @interface RecoveryCodeForIdentity
 */
export interface RecoveryCodeForIdentity {
    /**
     * Expires At is the timestamp of when the recovery flow expires
     * 
     * The timestamp when the recovery code expires.
     * @type {Date}
     * @memberof RecoveryCodeForIdentity
     */
    expires_at?: Date;
    /**
     * RecoveryCode is the code that can be used to recover the account
     * @type {string}
     * @memberof RecoveryCodeForIdentity
     */
    recovery_code: string;
    /**
     * RecoveryLink with flow
     * 
     * This link opens the recovery UI with an empty `code` field.
     * @type {string}
     * @memberof RecoveryCodeForIdentity
     */
    recovery_link: string;
}

/**
 * Check if a given object implements the RecoveryCodeForIdentity interface.
 */
export function instanceOfRecoveryCodeForIdentity(value: object): value is RecoveryCodeForIdentity {
    if (!('recovery_code' in value) || value['recovery_code'] === undefined) return false;
    if (!('recovery_link' in value) || value['recovery_link'] === undefined) return false;
    return true;
}

export function RecoveryCodeForIdentityFromJSON(json: any): RecoveryCodeForIdentity {
    return RecoveryCodeForIdentityFromJSONTyped(json, false);
}

export function RecoveryCodeForIdentityFromJSONTyped(json: any, ignoreDiscriminator: boolean): RecoveryCodeForIdentity {
    if (json == null) {
        return json;
    }
    return {
        
        'expires_at': json['expires_at'] == null ? undefined : (new Date(json['expires_at'])),
        'recovery_code': json['recovery_code'],
        'recovery_link': json['recovery_link'],
    };
}

export function RecoveryCodeForIdentityToJSON(value?: RecoveryCodeForIdentity | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'expires_at': value['expires_at'] == null ? undefined : ((value['expires_at']).toISOString()),
        'recovery_code': value['recovery_code'],
        'recovery_link': value['recovery_link'],
    };
}


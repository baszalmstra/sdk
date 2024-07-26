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
 * 
 * @export
 * @interface VerifiableCredentialProof
 */
export interface VerifiableCredentialProof {
    /**
     * 
     * @type {string}
     * @memberof VerifiableCredentialProof
     */
    jwt?: string;
    /**
     * 
     * @type {string}
     * @memberof VerifiableCredentialProof
     */
    proof_type?: string;
}

/**
 * Check if a given object implements the VerifiableCredentialProof interface.
 */
export function instanceOfVerifiableCredentialProof(value: object): value is VerifiableCredentialProof {
    return true;
}

export function VerifiableCredentialProofFromJSON(json: any): VerifiableCredentialProof {
    return VerifiableCredentialProofFromJSONTyped(json, false);
}

export function VerifiableCredentialProofFromJSONTyped(json: any, ignoreDiscriminator: boolean): VerifiableCredentialProof {
    if (json == null) {
        return json;
    }
    return {
        
        'jwt': json['jwt'] == null ? undefined : json['jwt'],
        'proof_type': json['proof_type'] == null ? undefined : json['proof_type'],
    };
}

export function VerifiableCredentialProofToJSON(value?: VerifiableCredentialProof | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'jwt': value['jwt'],
        'proof_type': value['proof_type'],
    };
}


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
 * @interface TaxLineItem
 */
export interface TaxLineItem {
    /**
     * 
     * @type {number}
     * @memberof TaxLineItem
     */
    amount_in_cent?: number;
    /**
     * 
     * @type {string}
     * @memberof TaxLineItem
     */
    title?: string;
}

/**
 * Check if a given object implements the TaxLineItem interface.
 */
export function instanceOfTaxLineItem(value: object): value is TaxLineItem {
    return true;
}

export function TaxLineItemFromJSON(json: any): TaxLineItem {
    return TaxLineItemFromJSONTyped(json, false);
}

export function TaxLineItemFromJSONTyped(json: any, ignoreDiscriminator: boolean): TaxLineItem {
    if (json == null) {
        return json;
    }
    return {
        
        'amount_in_cent': json['amount_in_cent'] == null ? undefined : json['amount_in_cent'],
        'title': json['title'] == null ? undefined : json['title'],
    };
}

export function TaxLineItemToJSON(value?: TaxLineItem | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'amount_in_cent': value['amount_in_cent'],
        'title': value['title'],
    };
}


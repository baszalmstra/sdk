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
 * @interface LineItemV1
 */
export interface LineItemV1 {
    /**
     * 
     * @type {number}
     * @memberof LineItemV1
     */
    amount_in_cent?: number;
    /**
     * 
     * @type {string}
     * @memberof LineItemV1
     */
    description?: string;
    /**
     * Each line item can have sub-items to create a hierarchy.
     * @type {Array<LineItemV1>}
     * @memberof LineItemV1
     */
    items?: Array<LineItemV1>;
    /**
     * 
     * @type {number}
     * @memberof LineItemV1
     */
    quantity?: number;
    /**
     * 
     * @type {string}
     * @memberof LineItemV1
     */
    title?: string;
    /**
     * 
     * @type {string}
     * @memberof LineItemV1
     */
    unit_price?: string;
}

/**
 * Check if a given object implements the LineItemV1 interface.
 */
export function instanceOfLineItemV1(value: object): value is LineItemV1 {
    return true;
}

export function LineItemV1FromJSON(json: any): LineItemV1 {
    return LineItemV1FromJSONTyped(json, false);
}

export function LineItemV1FromJSONTyped(json: any, ignoreDiscriminator: boolean): LineItemV1 {
    if (json == null) {
        return json;
    }
    return {
        
        'amount_in_cent': json['amount_in_cent'] == null ? undefined : json['amount_in_cent'],
        'description': json['description'] == null ? undefined : json['description'],
        'items': json['items'] == null ? undefined : ((json['items'] as Array<any>).map(LineItemV1FromJSON)),
        'quantity': json['quantity'] == null ? undefined : json['quantity'],
        'title': json['title'] == null ? undefined : json['title'],
        'unit_price': json['unit_price'] == null ? undefined : json['unit_price'],
    };
}

export function LineItemV1ToJSON(value?: LineItemV1 | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'amount_in_cent': value['amount_in_cent'],
        'description': value['description'],
        'items': value['items'] == null ? undefined : ((value['items'] as Array<any>).map(LineItemV1ToJSON)),
        'quantity': value['quantity'],
        'title': value['title'],
        'unit_price': value['unit_price'],
    };
}


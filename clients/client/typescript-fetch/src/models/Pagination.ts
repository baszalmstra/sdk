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
 * 
 * @export
 * @interface Pagination
 */
export interface Pagination {
    /**
     * Items per page
     * 
     * This is the number of items per page to return.
     * For details on pagination please head over to the [pagination documentation](https://www.ory.sh/docs/ecosystem/api-design#pagination).
     * @type {number}
     * @memberof Pagination
     */
    page_size?: number;
    /**
     * Next Page Token
     * 
     * The next page token.
     * For details on pagination please head over to the [pagination documentation](https://www.ory.sh/docs/ecosystem/api-design#pagination).
     * @type {string}
     * @memberof Pagination
     */
    page_token?: string;
}

/**
 * Check if a given object implements the Pagination interface.
 */
export function instanceOfPagination(value: object): value is Pagination {
    return true;
}

export function PaginationFromJSON(json: any): Pagination {
    return PaginationFromJSONTyped(json, false);
}

export function PaginationFromJSONTyped(json: any, ignoreDiscriminator: boolean): Pagination {
    if (json == null) {
        return json;
    }
    return {
        
        'page_size': json['page_size'] == null ? undefined : json['page_size'],
        'page_token': json['page_token'] == null ? undefined : json['page_token'],
    };
}

export function PaginationToJSON(value?: Pagination | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'page_size': value['page_size'],
        'page_token': value['page_token'],
    };
}


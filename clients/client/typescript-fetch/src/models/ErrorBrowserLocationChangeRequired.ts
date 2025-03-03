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
import type { ErrorGeneric } from './ErrorGeneric';
import {
    ErrorGenericFromJSON,
    ErrorGenericFromJSONTyped,
    ErrorGenericToJSON,
} from './ErrorGeneric';

/**
 * 
 * @export
 * @interface ErrorBrowserLocationChangeRequired
 */
export interface ErrorBrowserLocationChangeRequired {
    /**
     * 
     * @type {ErrorGeneric}
     * @memberof ErrorBrowserLocationChangeRequired
     */
    error?: ErrorGeneric;
    /**
     * Points to where to redirect the user to next.
     * @type {string}
     * @memberof ErrorBrowserLocationChangeRequired
     */
    redirect_browser_to?: string;
}

/**
 * Check if a given object implements the ErrorBrowserLocationChangeRequired interface.
 */
export function instanceOfErrorBrowserLocationChangeRequired(value: object): value is ErrorBrowserLocationChangeRequired {
    return true;
}

export function ErrorBrowserLocationChangeRequiredFromJSON(json: any): ErrorBrowserLocationChangeRequired {
    return ErrorBrowserLocationChangeRequiredFromJSONTyped(json, false);
}

export function ErrorBrowserLocationChangeRequiredFromJSONTyped(json: any, ignoreDiscriminator: boolean): ErrorBrowserLocationChangeRequired {
    if (json == null) {
        return json;
    }
    return {
        
        'error': json['error'] == null ? undefined : ErrorGenericFromJSON(json['error']),
        'redirect_browser_to': json['redirect_browser_to'] == null ? undefined : json['redirect_browser_to'],
    };
}

export function ErrorBrowserLocationChangeRequiredToJSON(value?: ErrorBrowserLocationChangeRequired | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'error': ErrorGenericToJSON(value['error']),
        'redirect_browser_to': value['redirect_browser_to'],
    };
}


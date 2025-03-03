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
import type { ContinueWithRecoveryUiFlow } from './ContinueWithRecoveryUiFlow';
import {
    ContinueWithRecoveryUiFlowFromJSON,
    ContinueWithRecoveryUiFlowFromJSONTyped,
    ContinueWithRecoveryUiFlowToJSON,
} from './ContinueWithRecoveryUiFlow';

/**
 * Indicates, that the UI flow could be continued by showing a recovery ui
 * @export
 * @interface ContinueWithRecoveryUi
 */
export interface ContinueWithRecoveryUi {
    /**
     * Action will always be `show_recovery_ui`
     * show_recovery_ui ContinueWithActionShowRecoveryUIString
     * @type {string}
     * @memberof ContinueWithRecoveryUi
     */
    action: ContinueWithRecoveryUiActionEnum;
    /**
     * 
     * @type {ContinueWithRecoveryUiFlow}
     * @memberof ContinueWithRecoveryUi
     */
    flow: ContinueWithRecoveryUiFlow;
}


/**
 * @export
 */
export const ContinueWithRecoveryUiActionEnum = {
    ShowRecoveryUi: 'show_recovery_ui'
} as const;
export type ContinueWithRecoveryUiActionEnum = typeof ContinueWithRecoveryUiActionEnum[keyof typeof ContinueWithRecoveryUiActionEnum];


/**
 * Check if a given object implements the ContinueWithRecoveryUi interface.
 */
export function instanceOfContinueWithRecoveryUi(value: object): value is ContinueWithRecoveryUi {
    if (!('action' in value) || value['action'] === undefined) return false;
    if (!('flow' in value) || value['flow'] === undefined) return false;
    return true;
}

export function ContinueWithRecoveryUiFromJSON(json: any): ContinueWithRecoveryUi {
    return ContinueWithRecoveryUiFromJSONTyped(json, false);
}

export function ContinueWithRecoveryUiFromJSONTyped(json: any, ignoreDiscriminator: boolean): ContinueWithRecoveryUi {
    if (json == null) {
        return json;
    }
    return {
        
        'action': json['action'],
        'flow': ContinueWithRecoveryUiFlowFromJSON(json['flow']),
    };
}

export function ContinueWithRecoveryUiToJSON(value?: ContinueWithRecoveryUi | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'action': value['action'],
        'flow': ContinueWithRecoveryUiFlowToJSON(value['flow']),
    };
}


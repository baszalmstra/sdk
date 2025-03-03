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


/**
 * choose_method: ask the user to choose a method (e.g. registration with email)
 * sent_email: the email has been sent to the user
 * passed_challenge: the request was successful and the registration challenge was passed.
 * @export
 */
export const RegistrationFlowState = {
    ChooseMethod: 'choose_method',
    SentEmail: 'sent_email',
    PassedChallenge: 'passed_challenge'
} as const;
export type RegistrationFlowState = typeof RegistrationFlowState[keyof typeof RegistrationFlowState];


export function instanceOfRegistrationFlowState(value: any): boolean {
    for (const key in RegistrationFlowState) {
        if (Object.prototype.hasOwnProperty.call(RegistrationFlowState, key)) {
            if ((RegistrationFlowState as Record<string, RegistrationFlowState>)[key] === value) {
                return true;
            }
        }
    }
    return false;
}

export function RegistrationFlowStateFromJSON(json: any): RegistrationFlowState {
    return RegistrationFlowStateFromJSONTyped(json, false);
}

export function RegistrationFlowStateFromJSONTyped(json: any, ignoreDiscriminator: boolean): RegistrationFlowState {
    return json as RegistrationFlowState;
}

export function RegistrationFlowStateToJSON(value?: RegistrationFlowState | null): any {
    return value as any;
}


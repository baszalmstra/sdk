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
 * @interface AddProjectToWorkspaceBody
 */
export interface AddProjectToWorkspaceBody {
    /**
     * The environment of the project in the workspace. Can be one of "prod" or
     * "dev". Note that the number of projects in the "prod" environment is limited
     * depending on the subscription.
     * prod Production
     * stage Staging
     * dev Development
     * @type {string}
     * @memberof AddProjectToWorkspaceBody
     */
    environment: AddProjectToWorkspaceBodyEnvironmentEnum;
    /**
     * The action to take with the project subscription. Can be one of "migrate", and
     * "ignore". "migrate" will migrate the project subscription to the workspace.
     * "ignore" will ignore the project subscription.
     * migrate ProjectSubscriptionActionMigrate  ProjectSubscriptionActionMigrate will migrate the project subscription to the  workspace.
     * ignore ProjectSubscriptionActionIgnore  ProjectSubscriptionActionIgnore will ignore the project subscription.
     * @type {string}
     * @memberof AddProjectToWorkspaceBody
     */
    project_subscription: AddProjectToWorkspaceBodyProjectSubscriptionEnum;
}


/**
 * @export
 */
export const AddProjectToWorkspaceBodyEnvironmentEnum = {
    Prod: 'prod',
    Stage: 'stage',
    Dev: 'dev'
} as const;
export type AddProjectToWorkspaceBodyEnvironmentEnum = typeof AddProjectToWorkspaceBodyEnvironmentEnum[keyof typeof AddProjectToWorkspaceBodyEnvironmentEnum];

/**
 * @export
 */
export const AddProjectToWorkspaceBodyProjectSubscriptionEnum = {
    Migrate: 'migrate',
    Ignore: 'ignore'
} as const;
export type AddProjectToWorkspaceBodyProjectSubscriptionEnum = typeof AddProjectToWorkspaceBodyProjectSubscriptionEnum[keyof typeof AddProjectToWorkspaceBodyProjectSubscriptionEnum];


/**
 * Check if a given object implements the AddProjectToWorkspaceBody interface.
 */
export function instanceOfAddProjectToWorkspaceBody(value: object): value is AddProjectToWorkspaceBody {
    if (!('environment' in value) || value['environment'] === undefined) return false;
    if (!('project_subscription' in value) || value['project_subscription'] === undefined) return false;
    return true;
}

export function AddProjectToWorkspaceBodyFromJSON(json: any): AddProjectToWorkspaceBody {
    return AddProjectToWorkspaceBodyFromJSONTyped(json, false);
}

export function AddProjectToWorkspaceBodyFromJSONTyped(json: any, ignoreDiscriminator: boolean): AddProjectToWorkspaceBody {
    if (json == null) {
        return json;
    }
    return {
        
        'environment': json['environment'],
        'project_subscription': json['project_subscription'],
    };
}

export function AddProjectToWorkspaceBodyToJSON(value?: AddProjectToWorkspaceBody | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'environment': value['environment'],
        'project_subscription': value['project_subscription'],
    };
}


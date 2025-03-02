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
 * @interface ProjectHost
 */
export interface ProjectHost {
    /**
     * The project's host.
     * @type {string}
     * @memberof ProjectHost
     */
    host: string;
    /**
     * The mapping's ID.
     * @type {string}
     * @memberof ProjectHost
     */
    readonly id: string;
    /**
     * The Revision's Project ID
     * @type {string}
     * @memberof ProjectHost
     */
    project_id: string;
}

/**
 * Check if a given object implements the ProjectHost interface.
 */
export function instanceOfProjectHost(value: object): value is ProjectHost {
    if (!('host' in value) || value['host'] === undefined) return false;
    if (!('id' in value) || value['id'] === undefined) return false;
    if (!('project_id' in value) || value['project_id'] === undefined) return false;
    return true;
}

export function ProjectHostFromJSON(json: any): ProjectHost {
    return ProjectHostFromJSONTyped(json, false);
}

export function ProjectHostFromJSONTyped(json: any, ignoreDiscriminator: boolean): ProjectHost {
    if (json == null) {
        return json;
    }
    return {
        
        'host': json['host'],
        'id': json['id'],
        'project_id': json['project_id'],
    };
}

export function ProjectHostToJSON(value?: Omit<ProjectHost, 'id'> | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'host': value['host'],
        'project_id': value['project_id'],
    };
}


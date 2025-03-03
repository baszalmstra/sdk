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
 * @interface ProjectMember
 */
export interface ProjectMember {
    /**
     * 
     * @type {string}
     * @memberof ProjectMember
     */
    email: string;
    /**
     * 
     * @type {boolean}
     * @memberof ProjectMember
     */
    email_verified: boolean;
    /**
     * 
     * @type {string}
     * @memberof ProjectMember
     */
    id: string;
    /**
     * 
     * @type {string}
     * @memberof ProjectMember
     */
    name: string;
    /**
     * 
     * @type {string}
     * @memberof ProjectMember
     */
    role: string;
}

/**
 * Check if a given object implements the ProjectMember interface.
 */
export function instanceOfProjectMember(value: object): value is ProjectMember {
    if (!('email' in value) || value['email'] === undefined) return false;
    if (!('email_verified' in value) || value['email_verified'] === undefined) return false;
    if (!('id' in value) || value['id'] === undefined) return false;
    if (!('name' in value) || value['name'] === undefined) return false;
    if (!('role' in value) || value['role'] === undefined) return false;
    return true;
}

export function ProjectMemberFromJSON(json: any): ProjectMember {
    return ProjectMemberFromJSONTyped(json, false);
}

export function ProjectMemberFromJSONTyped(json: any, ignoreDiscriminator: boolean): ProjectMember {
    if (json == null) {
        return json;
    }
    return {
        
        'email': json['email'],
        'email_verified': json['email_verified'],
        'id': json['id'],
        'name': json['name'],
        'role': json['role'],
    };
}

export function ProjectMemberToJSON(value?: ProjectMember | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'email': value['email'],
        'email_verified': value['email_verified'],
        'id': value['id'],
        'name': value['name'],
        'role': value['role'],
    };
}


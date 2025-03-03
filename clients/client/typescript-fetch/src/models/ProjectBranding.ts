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
import type { ProjectBrandingTheme } from './ProjectBrandingTheme';
import {
    ProjectBrandingThemeFromJSON,
    ProjectBrandingThemeFromJSONTyped,
    ProjectBrandingThemeToJSON,
} from './ProjectBrandingTheme';

/**
 * 
 * @export
 * @interface ProjectBranding
 */
export interface ProjectBranding {
    /**
     * The Customization Creation Date
     * @type {Date}
     * @memberof ProjectBranding
     */
    readonly created_at: Date;
    /**
     * 
     * @type {ProjectBrandingTheme}
     * @memberof ProjectBranding
     */
    default_theme: ProjectBrandingTheme;
    /**
     * The customization ID.
     * @type {string}
     * @memberof ProjectBranding
     */
    readonly id: string;
    /**
     * The Project's ID this customization is associated with
     * @type {string}
     * @memberof ProjectBranding
     */
    project_id: string;
    /**
     * 
     * @type {Array<ProjectBrandingTheme>}
     * @memberof ProjectBranding
     */
    themes: Array<ProjectBrandingTheme>;
    /**
     * Last Time Branding was Updated
     * @type {Date}
     * @memberof ProjectBranding
     */
    readonly updated_at: Date;
}

/**
 * Check if a given object implements the ProjectBranding interface.
 */
export function instanceOfProjectBranding(value: object): value is ProjectBranding {
    if (!('created_at' in value) || value['created_at'] === undefined) return false;
    if (!('default_theme' in value) || value['default_theme'] === undefined) return false;
    if (!('id' in value) || value['id'] === undefined) return false;
    if (!('project_id' in value) || value['project_id'] === undefined) return false;
    if (!('themes' in value) || value['themes'] === undefined) return false;
    if (!('updated_at' in value) || value['updated_at'] === undefined) return false;
    return true;
}

export function ProjectBrandingFromJSON(json: any): ProjectBranding {
    return ProjectBrandingFromJSONTyped(json, false);
}

export function ProjectBrandingFromJSONTyped(json: any, ignoreDiscriminator: boolean): ProjectBranding {
    if (json == null) {
        return json;
    }
    return {
        
        'created_at': (new Date(json['created_at'])),
        'default_theme': ProjectBrandingThemeFromJSON(json['default_theme']),
        'id': json['id'],
        'project_id': json['project_id'],
        'themes': ((json['themes'] as Array<any>).map(ProjectBrandingThemeFromJSON)),
        'updated_at': (new Date(json['updated_at'])),
    };
}

export function ProjectBrandingToJSON(value?: Omit<ProjectBranding, 'created_at'|'id'|'updated_at'> | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'default_theme': ProjectBrandingThemeToJSON(value['default_theme']),
        'project_id': value['project_id'],
        'themes': ((value['themes'] as Array<any>).map(ProjectBrandingThemeToJSON)),
    };
}


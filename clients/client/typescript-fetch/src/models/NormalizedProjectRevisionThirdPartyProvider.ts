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
 * @interface NormalizedProjectRevisionThirdPartyProvider
 */
export interface NormalizedProjectRevisionThirdPartyProvider {
    /**
     * 
     * @type {Array<string>}
     * @memberof NormalizedProjectRevisionThirdPartyProvider
     */
    additional_id_token_audiences?: Array<string>;
    /**
     * 
     * @type {string}
     * @memberof NormalizedProjectRevisionThirdPartyProvider
     */
    apple_private_key?: string | null;
    /**
     * Apple Private Key Identifier
     * 
     * Sign In with Apple Private Key Identifier needed for generating a JWT token for client secret
     * @type {string}
     * @memberof NormalizedProjectRevisionThirdPartyProvider
     */
    apple_private_key_id?: string;
    /**
     * Apple Developer Team ID
     * 
     * Apple Developer Team ID needed for generating a JWT token for client secret
     * @type {string}
     * @memberof NormalizedProjectRevisionThirdPartyProvider
     */
    apple_team_id?: string;
    /**
     * AuthURL is the authorize url, typically something like: https://example.org/oauth2/auth
     * Should only be used when the OAuth2 / OpenID Connect server is not supporting OpenID Connect Discovery and when
     * `provider` is set to `generic`.
     * @type {string}
     * @memberof NormalizedProjectRevisionThirdPartyProvider
     */
    auth_url?: string;
    /**
     * Tenant is the Azure AD Tenant to use for authentication, and must be set when `provider` is set to `microsoft`.
     * 
     * Can be either `common`, `organizations`, `consumers` for a multitenant application or a specific tenant like
     * `8eaef023-2b34-4da1-9baa-8bc8c9d6a490` or `contoso.onmicrosoft.com`.
     * @type {string}
     * @memberof NormalizedProjectRevisionThirdPartyProvider
     */
    azure_tenant?: string;
    /**
     * 
     * @type {string}
     * @memberof NormalizedProjectRevisionThirdPartyProvider
     */
    claims_source?: string | null;
    /**
     * ClientID is the application's Client ID.
     * @type {string}
     * @memberof NormalizedProjectRevisionThirdPartyProvider
     */
    client_id?: string;
    /**
     * 
     * @type {string}
     * @memberof NormalizedProjectRevisionThirdPartyProvider
     */
    client_secret?: string | null;
    /**
     * The Project's Revision Creation Date
     * @type {Date}
     * @memberof NormalizedProjectRevisionThirdPartyProvider
     */
    readonly created_at?: Date;
    /**
     * 
     * @type {string}
     * @memberof NormalizedProjectRevisionThirdPartyProvider
     */
    id?: string;
    /**
     * IssuerURL is the OpenID Connect Server URL. You can leave this empty if `provider` is not set to `generic`.
     * If set, neither `auth_url` nor `token_url` are required.
     * @type {string}
     * @memberof NormalizedProjectRevisionThirdPartyProvider
     */
    issuer_url?: string;
    /**
     * Label represents an optional label which can be used in the UI generation.
     * @type {string}
     * @memberof NormalizedProjectRevisionThirdPartyProvider
     */
    label?: string;
    /**
     * Mapper specifies the JSONNet code snippet which uses the OpenID Connect Provider's data (e.g. GitHub or Google
     * profile information) to hydrate the identity's data.
     * @type {string}
     * @memberof NormalizedProjectRevisionThirdPartyProvider
     */
    mapper_url?: string;
    /**
     * 
     * @type {string}
     * @memberof NormalizedProjectRevisionThirdPartyProvider
     */
    organization_id?: string | null;
    /**
     * The Revision's ID this schema belongs to
     * @type {string}
     * @memberof NormalizedProjectRevisionThirdPartyProvider
     */
    project_revision_id?: string;
    /**
     * Provider is either "generic" for a generic OAuth 2.0 / OpenID Connect Provider or one of:
     * generic
     * google
     * github
     * gitlab
     * microsoft
     * discord
     * slack
     * facebook
     * vk
     * yandex
     * apple
     * @type {string}
     * @memberof NormalizedProjectRevisionThirdPartyProvider
     */
    provider?: string;
    /**
     * ID is the provider's ID
     * @type {string}
     * @memberof NormalizedProjectRevisionThirdPartyProvider
     */
    provider_id?: string;
    /**
     * 
     * @type {object}
     * @memberof NormalizedProjectRevisionThirdPartyProvider
     */
    requested_claims?: object;
    /**
     * 
     * @type {Array<string>}
     * @memberof NormalizedProjectRevisionThirdPartyProvider
     */
    scope?: Array<string>;
    /**
     * State indicates the state of the provider
     * 
     * Only providers with state `enabled` will be used for authentication
     * enabled ThirdPartyProviderStateEnabled
     * disabled ThirdPartyProviderStateDisabled
     * @type {string}
     * @memberof NormalizedProjectRevisionThirdPartyProvider
     */
    state?: NormalizedProjectRevisionThirdPartyProviderStateEnum;
    /**
     * 
     * @type {string}
     * @memberof NormalizedProjectRevisionThirdPartyProvider
     */
    subject_source?: string | null;
    /**
     * TokenURL is the token url, typically something like: https://example.org/oauth2/token
     * 
     * Should only be used when the OAuth2 / OpenID Connect server is not supporting OpenID Connect Discovery and when
     * `provider` is set to `generic`.
     * @type {string}
     * @memberof NormalizedProjectRevisionThirdPartyProvider
     */
    token_url?: string;
    /**
     * Last Time Project's Revision was Updated
     * @type {Date}
     * @memberof NormalizedProjectRevisionThirdPartyProvider
     */
    readonly updated_at?: Date;
}


/**
 * @export
 */
export const NormalizedProjectRevisionThirdPartyProviderStateEnum = {
    Enabled: 'enabled',
    Disabled: 'disabled'
} as const;
export type NormalizedProjectRevisionThirdPartyProviderStateEnum = typeof NormalizedProjectRevisionThirdPartyProviderStateEnum[keyof typeof NormalizedProjectRevisionThirdPartyProviderStateEnum];


/**
 * Check if a given object implements the NormalizedProjectRevisionThirdPartyProvider interface.
 */
export function instanceOfNormalizedProjectRevisionThirdPartyProvider(value: object): value is NormalizedProjectRevisionThirdPartyProvider {
    return true;
}

export function NormalizedProjectRevisionThirdPartyProviderFromJSON(json: any): NormalizedProjectRevisionThirdPartyProvider {
    return NormalizedProjectRevisionThirdPartyProviderFromJSONTyped(json, false);
}

export function NormalizedProjectRevisionThirdPartyProviderFromJSONTyped(json: any, ignoreDiscriminator: boolean): NormalizedProjectRevisionThirdPartyProvider {
    if (json == null) {
        return json;
    }
    return {
        
        'additional_id_token_audiences': json['additional_id_token_audiences'] == null ? undefined : json['additional_id_token_audiences'],
        'apple_private_key': json['apple_private_key'] == null ? undefined : json['apple_private_key'],
        'apple_private_key_id': json['apple_private_key_id'] == null ? undefined : json['apple_private_key_id'],
        'apple_team_id': json['apple_team_id'] == null ? undefined : json['apple_team_id'],
        'auth_url': json['auth_url'] == null ? undefined : json['auth_url'],
        'azure_tenant': json['azure_tenant'] == null ? undefined : json['azure_tenant'],
        'claims_source': json['claims_source'] == null ? undefined : json['claims_source'],
        'client_id': json['client_id'] == null ? undefined : json['client_id'],
        'client_secret': json['client_secret'] == null ? undefined : json['client_secret'],
        'created_at': json['created_at'] == null ? undefined : (new Date(json['created_at'])),
        'id': json['id'] == null ? undefined : json['id'],
        'issuer_url': json['issuer_url'] == null ? undefined : json['issuer_url'],
        'label': json['label'] == null ? undefined : json['label'],
        'mapper_url': json['mapper_url'] == null ? undefined : json['mapper_url'],
        'organization_id': json['organization_id'] == null ? undefined : json['organization_id'],
        'project_revision_id': json['project_revision_id'] == null ? undefined : json['project_revision_id'],
        'provider': json['provider'] == null ? undefined : json['provider'],
        'provider_id': json['provider_id'] == null ? undefined : json['provider_id'],
        'requested_claims': json['requested_claims'] == null ? undefined : json['requested_claims'],
        'scope': json['scope'] == null ? undefined : json['scope'],
        'state': json['state'] == null ? undefined : json['state'],
        'subject_source': json['subject_source'] == null ? undefined : json['subject_source'],
        'token_url': json['token_url'] == null ? undefined : json['token_url'],
        'updated_at': json['updated_at'] == null ? undefined : (new Date(json['updated_at'])),
    };
}

export function NormalizedProjectRevisionThirdPartyProviderToJSON(value?: Omit<NormalizedProjectRevisionThirdPartyProvider, 'created_at'|'updated_at'> | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'additional_id_token_audiences': value['additional_id_token_audiences'],
        'apple_private_key': value['apple_private_key'],
        'apple_private_key_id': value['apple_private_key_id'],
        'apple_team_id': value['apple_team_id'],
        'auth_url': value['auth_url'],
        'azure_tenant': value['azure_tenant'],
        'claims_source': value['claims_source'],
        'client_id': value['client_id'],
        'client_secret': value['client_secret'],
        'id': value['id'],
        'issuer_url': value['issuer_url'],
        'label': value['label'],
        'mapper_url': value['mapper_url'],
        'organization_id': value['organization_id'],
        'project_revision_id': value['project_revision_id'],
        'provider': value['provider'],
        'provider_id': value['provider_id'],
        'requested_claims': value['requested_claims'],
        'scope': value['scope'],
        'state': value['state'],
        'subject_source': value['subject_source'],
        'token_url': value['token_url'],
    };
}


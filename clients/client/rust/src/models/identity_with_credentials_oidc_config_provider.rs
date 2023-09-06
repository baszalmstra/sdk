/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v1.2.1
 * Contact: support@ory.sh
 * Generated by: https://openapi-generator.tech
 */

/// IdentityWithCredentialsOidcConfigProvider : Create Identity and Import Social Sign In Credentials Configuration



#[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
pub struct IdentityWithCredentialsOidcConfigProvider {
    /// The OpenID Connect provider to link the subject to. Usually something like `google` or `github`.
    #[serde(rename = "provider")]
    pub provider: String,
    /// The subject (`sub`) of the OpenID Connect connection. Usually the `sub` field of the ID Token.
    #[serde(rename = "subject")]
    pub subject: String,
}


impl IdentityWithCredentialsOidcConfigProvider {
    /// Create Identity and Import Social Sign In Credentials Configuration
    pub fn new(provider: String, subject: String) -> IdentityWithCredentialsOidcConfigProvider {
        IdentityWithCredentialsOidcConfigProvider {
                provider,
                subject,
        }
    }
}



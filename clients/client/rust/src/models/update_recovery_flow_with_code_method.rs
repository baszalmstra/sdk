/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v1.1.3
 * Contact: support@ory.sh
 * Generated by: https://openapi-generator.tech
 */

/// UpdateRecoveryFlowWithCodeMethod : Update Recovery Flow with Code Method



#[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
pub struct UpdateRecoveryFlowWithCodeMethod {
    /// Code from recovery email  Sent to the user once a recovery has been initiated and is used to prove that the user is in possession of the email
    #[serde(rename = "code", skip_serializing_if = "Option::is_none")]
    pub code: Option<String>,
    /// Sending the anti-csrf token is only required for browser login flows.
    #[serde(rename = "csrf_token", skip_serializing_if = "Option::is_none")]
    pub csrf_token: Option<String>,
    /// Email to Recover  Needs to be set when initiating the flow. If the email is a registered recovery email, a recovery link will be sent. If the email is not known, a email with details on what happened will be sent instead.  format: email
    #[serde(rename = "email", skip_serializing_if = "Option::is_none")]
    pub email: Option<String>,
    /// Method supports `link` and `code` only right now.
    #[serde(rename = "method")]
    pub method: String,
}


impl UpdateRecoveryFlowWithCodeMethod {
    /// Update Recovery Flow with Code Method
    pub fn new(method: String) -> UpdateRecoveryFlowWithCodeMethod {
        UpdateRecoveryFlowWithCodeMethod {
                code: None,
                csrf_token: None,
                email: None,
                method,
        }
    }
}



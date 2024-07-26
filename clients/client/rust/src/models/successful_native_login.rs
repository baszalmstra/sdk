/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v1.14.3
 * Contact: support@ory.sh
 * Generated by: https://openapi-generator.tech
 */

use crate::models;
use serde::{Deserialize, Serialize};

/// SuccessfulNativeLogin : The Response for Login Flows via API
#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct SuccessfulNativeLogin {
    /// Contains a list of actions, that could follow this flow  It can, for example, this will contain a reference to the verification flow, created as part of the user's registration or the token of the session.
    #[serde(rename = "continue_with", skip_serializing_if = "Option::is_none")]
    pub continue_with: Option<Vec<models::ContinueWith>>,
    #[serde(rename = "session")]
    pub session: Box<models::Session>,
    /// The Session Token  A session token is equivalent to a session cookie, but it can be sent in the HTTP Authorization Header:  Authorization: bearer ${session-token}  The session token is only issued for API flows, not for Browser flows!
    #[serde(rename = "session_token", skip_serializing_if = "Option::is_none")]
    pub session_token: Option<String>,
}

impl SuccessfulNativeLogin {
    /// The Response for Login Flows via API
    pub fn new(session: models::Session) -> SuccessfulNativeLogin {
        SuccessfulNativeLogin {
            continue_with: None,
            session: Box::new(session),
            session_token: None,
        }
    }
}


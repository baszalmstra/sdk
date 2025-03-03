/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v1.15.7
 * Contact: support@ory.sh
 * Generated by: https://openapi-generator.tech
 */

use crate::models;
use serde::{Deserialize, Serialize};

/// ContinueWithRedirectBrowserTo : Indicates, that the UI flow could be continued by showing a recovery ui
#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct ContinueWithRedirectBrowserTo {
    /// Action will always be `redirect_browser_to` redirect_browser_to ContinueWithActionRedirectBrowserToString
    #[serde(rename = "action")]
    pub action: ActionEnum,
    /// The URL to redirect the browser to
    #[serde(rename = "redirect_browser_to")]
    pub redirect_browser_to: String,
}

impl ContinueWithRedirectBrowserTo {
    /// Indicates, that the UI flow could be continued by showing a recovery ui
    pub fn new(action: ActionEnum, redirect_browser_to: String) -> ContinueWithRedirectBrowserTo {
        ContinueWithRedirectBrowserTo {
            action,
            redirect_browser_to,
        }
    }
}
/// Action will always be `redirect_browser_to` redirect_browser_to ContinueWithActionRedirectBrowserToString
#[derive(Clone, Copy, Debug, Eq, PartialEq, Ord, PartialOrd, Hash, Serialize, Deserialize)]
pub enum ActionEnum {
    #[serde(rename = "redirect_browser_to")]
    RedirectBrowserTo,
}

impl Default for ActionEnum {
    fn default() -> ActionEnum {
        Self::RedirectBrowserTo
    }
}


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

/// ContinueWithSettingsUi : Indicates, that the UI flow could be continued by showing a settings ui
#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct ContinueWithSettingsUi {
    /// Action will always be `show_settings_ui` show_settings_ui ContinueWithActionShowSettingsUIString
    #[serde(rename = "action")]
    pub action: ActionEnum,
    #[serde(rename = "flow")]
    pub flow: Box<models::ContinueWithSettingsUiFlow>,
}

impl ContinueWithSettingsUi {
    /// Indicates, that the UI flow could be continued by showing a settings ui
    pub fn new(action: ActionEnum, flow: models::ContinueWithSettingsUiFlow) -> ContinueWithSettingsUi {
        ContinueWithSettingsUi {
            action,
            flow: Box::new(flow),
        }
    }
}
/// Action will always be `show_settings_ui` show_settings_ui ContinueWithActionShowSettingsUIString
#[derive(Clone, Copy, Debug, Eq, PartialEq, Ord, PartialOrd, Hash, Serialize, Deserialize)]
pub enum ActionEnum {
    #[serde(rename = "show_settings_ui")]
    ShowSettingsUi,
}

impl Default for ActionEnum {
    fn default() -> ActionEnum {
        Self::ShowSettingsUi
    }
}


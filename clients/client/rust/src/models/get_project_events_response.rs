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

/// GetProjectEventsResponse : Response of the getProjectEvents endpoint
#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct GetProjectEventsResponse {
    /// The list of data points.
    #[serde(rename = "events")]
    pub events: Vec<models::ProjectEventsDatapoint>,
    /// Pagination token to be included in next page request
    #[serde(rename = "page_token", skip_serializing_if = "Option::is_none")]
    pub page_token: Option<String>,
}

impl GetProjectEventsResponse {
    /// Response of the getProjectEvents endpoint
    pub fn new(events: Vec<models::ProjectEventsDatapoint>) -> GetProjectEventsResponse {
        GetProjectEventsResponse {
            events,
            page_token: None,
        }
    }
}


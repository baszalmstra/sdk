/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v1.1.37
 * Contact: support@ory.sh
 * Generated by: https://openapi-generator.tech
 */

/// CreateProjectInvitesResponse : Response to the create project invite request



#[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
pub struct CreateProjectInvitesResponse {
    #[serde(rename = "all_invites")]
    pub all_invites: Vec<crate::models::ProjectInvite>,
    #[serde(rename = "created_invites")]
    pub created_invites: Vec<crate::models::ProjectInvite>,
}


impl CreateProjectInvitesResponse {
    /// Response to the create project invite request
    pub fn new(all_invites: Vec<crate::models::ProjectInvite>, created_invites: Vec<crate::models::ProjectInvite>) -> CreateProjectInvitesResponse {
        CreateProjectInvitesResponse {
                all_invites,
                created_invites,
        }
    }
}



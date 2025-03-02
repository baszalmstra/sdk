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

/// CreateWorkspaceMemberInviteBody : Create Workspace Invite Request Body
#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct CreateWorkspaceMemberInviteBody {
    /// A email to invite
    #[serde(rename = "invitee_email")]
    pub invitee_email: String,
    /// The role the user will have in the workspace owner WorkspaceMemberRoleOwner developer WorkspaceMemberRoleDeveloper
    #[serde(rename = "role")]
    pub role: RoleEnum,
}

impl CreateWorkspaceMemberInviteBody {
    /// Create Workspace Invite Request Body
    pub fn new(invitee_email: String, role: RoleEnum) -> CreateWorkspaceMemberInviteBody {
        CreateWorkspaceMemberInviteBody {
            invitee_email,
            role,
        }
    }
}
/// The role the user will have in the workspace owner WorkspaceMemberRoleOwner developer WorkspaceMemberRoleDeveloper
#[derive(Clone, Copy, Debug, Eq, PartialEq, Ord, PartialOrd, Hash, Serialize, Deserialize)]
pub enum RoleEnum {
    #[serde(rename = "owner")]
    Owner,
    #[serde(rename = "developer")]
    Developer,
}

impl Default for RoleEnum {
    fn default() -> RoleEnum {
        Self::Owner
    }
}


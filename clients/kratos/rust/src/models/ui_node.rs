/*
 * Ory Kratos API
 *
 * Documentation for all public and administrative Ory Kratos APIs. Public and administrative APIs are exposed on different ports. Public APIs can face the public internet without any protection while administrative APIs should never be exposed without prior authorization. To protect the administative API port you should use something like Nginx, Ory Oathkeeper, or any other technology capable of authorizing incoming requests. 
 *
 * The version of the OpenAPI document: v0.7.3-alpha.8
 * Contact: hi@ory.sh
 * Generated by: https://openapi-generator.tech
 */

/// UiNode : Nodes are represented as HTML elements or their native UI equivalents. For example, a node can be an `<img>` tag, or an `<input element>` but also `some plain text`.



#[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
pub struct UiNode {
    #[serde(rename = "attributes")]
    pub attributes: Box<crate::models::UiNodeAttributes>,
    #[serde(rename = "group")]
    pub group: String,
    #[serde(rename = "messages")]
    pub messages: Vec<crate::models::UiText>,
    #[serde(rename = "meta")]
    pub meta: Box<crate::models::Meta>,
    #[serde(rename = "type")]
    pub _type: String,
}

impl UiNode {
    /// Nodes are represented as HTML elements or their native UI equivalents. For example, a node can be an `<img>` tag, or an `<input element>` but also `some plain text`.
    pub fn new(attributes: crate::models::UiNodeAttributes, group: String, messages: Vec<crate::models::UiText>, meta: crate::models::Meta, _type: String) -> UiNode {
        UiNode {
            attributes: Box::new(attributes),
            group,
            messages,
            meta: Box::new(meta),
            _type,
        }
    }
}



/*
 * Ory Kratos API
 *
 * Documentation for all public and administrative Ory Kratos APIs. Public and administrative APIs are exposed on different ports. Public APIs can face the public internet without any protection while administrative APIs should never be exposed without prior authorization. To protect the administative API port you should use something like Nginx, Ory Oathkeeper, or any other technology capable of authorizing incoming requests. 
 *
 * The version of the OpenAPI document: v0.7.3-alpha.8
 * Contact: hi@ory.sh
 * Generated by: https://openapi-generator.tech
 */

/// PluginInterfaceType : PluginInterfaceType plugin interface type



#[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
pub struct PluginInterfaceType {
    /// capability
    #[serde(rename = "Capability")]
    pub capability: String,
    /// prefix
    #[serde(rename = "Prefix")]
    pub prefix: String,
    /// version
    #[serde(rename = "Version")]
    pub version: String,
}

impl PluginInterfaceType {
    /// PluginInterfaceType plugin interface type
    pub fn new(capability: String, prefix: String, version: String) -> PluginInterfaceType {
        PluginInterfaceType {
            capability,
            prefix,
            version,
        }
    }
}



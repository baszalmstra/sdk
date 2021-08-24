/*
 * ORY Oathkeeper
 *
 * ORY Oathkeeper is a reverse proxy that checks the HTTP Authorization for validity against a set of rules. This service uses Hydra to validate access tokens and policies.
 *
 * The version of the OpenAPI document: v0.0.0-alpha.52
 * Contact: hi@ory.am
 * Generated by: https://openapi-generator.tech
 */

/// SwaggerJsonWebKey : SwaggerJSONWebKey swagger JSON web key



#[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
pub struct SwaggerJsonWebKey {
    /// The \"alg\" (algorithm) parameter identifies the algorithm intended for use with the key.  The values used should either be registered in the IANA \"JSON Web Signature and Encryption Algorithms\" registry established by [JWA] or be a value that contains a Collision- Resistant Name.
    #[serde(rename = "alg", skip_serializing_if = "Option::is_none")]
    pub alg: Option<String>,
    /// crv
    #[serde(rename = "crv", skip_serializing_if = "Option::is_none")]
    pub crv: Option<String>,
    /// d
    #[serde(rename = "d", skip_serializing_if = "Option::is_none")]
    pub d: Option<String>,
    /// dp
    #[serde(rename = "dp", skip_serializing_if = "Option::is_none")]
    pub dp: Option<String>,
    /// dq
    #[serde(rename = "dq", skip_serializing_if = "Option::is_none")]
    pub dq: Option<String>,
    /// e
    #[serde(rename = "e", skip_serializing_if = "Option::is_none")]
    pub e: Option<String>,
    /// k
    #[serde(rename = "k", skip_serializing_if = "Option::is_none")]
    pub k: Option<String>,
    /// The \"kid\" (key ID) parameter is used to match a specific key.  This is used, for instance, to choose among a set of keys within a JWK Set during key rollover.  The structure of the \"kid\" value is unspecified.  When \"kid\" values are used within a JWK Set, different keys within the JWK Set SHOULD use distinct \"kid\" values.  (One example in which different keys might use the same \"kid\" value is if they have different \"kty\" (key type) values but are considered to be equivalent alternatives by the application using them.)  The \"kid\" value is a case-sensitive string.
    #[serde(rename = "kid", skip_serializing_if = "Option::is_none")]
    pub kid: Option<String>,
    /// The \"kty\" (key type) parameter identifies the cryptographic algorithm family used with the key, such as \"RSA\" or \"EC\". \"kty\" values should either be registered in the IANA \"JSON Web Key Types\" registry established by [JWA] or be a value that contains a Collision- Resistant Name.  The \"kty\" value is a case-sensitive string.
    #[serde(rename = "kty", skip_serializing_if = "Option::is_none")]
    pub kty: Option<String>,
    /// n
    #[serde(rename = "n", skip_serializing_if = "Option::is_none")]
    pub n: Option<String>,
    /// p
    #[serde(rename = "p", skip_serializing_if = "Option::is_none")]
    pub p: Option<String>,
    /// q
    #[serde(rename = "q", skip_serializing_if = "Option::is_none")]
    pub q: Option<String>,
    /// qi
    #[serde(rename = "qi", skip_serializing_if = "Option::is_none")]
    pub qi: Option<String>,
    /// The \"use\" (public key use) parameter identifies the intended use of the public key. The \"use\" parameter is employed to indicate whether a public key is used for encrypting data or verifying the signature on data. Values are commonly \"sig\" (signature) or \"enc\" (encryption).
    #[serde(rename = "use", skip_serializing_if = "Option::is_none")]
    pub _use: Option<String>,
    /// x
    #[serde(rename = "x", skip_serializing_if = "Option::is_none")]
    pub x: Option<String>,
    /// The \"x5c\" (X.509 certificate chain) parameter contains a chain of one or more PKIX certificates [RFC5280].  The certificate chain is represented as a JSON array of certificate value strings.  Each string in the array is a base64-encoded (Section 4 of [RFC4648] -- not base64url-encoded) DER [ITU.X690.1994] PKIX certificate value. The PKIX certificate containing the key value MUST be the first certificate.
    #[serde(rename = "x5c", skip_serializing_if = "Option::is_none")]
    pub x5c: Option<Vec<String>>,
    /// y
    #[serde(rename = "y", skip_serializing_if = "Option::is_none")]
    pub y: Option<String>,
}

impl SwaggerJsonWebKey {
    /// SwaggerJSONWebKey swagger JSON web key
    pub fn new() -> SwaggerJsonWebKey {
        SwaggerJsonWebKey {
            alg: None,
            crv: None,
            d: None,
            dp: None,
            dq: None,
            e: None,
            k: None,
            kid: None,
            kty: None,
            n: None,
            p: None,
            q: None,
            qi: None,
            _use: None,
            x: None,
            x5c: None,
            y: None,
        }
    }
}



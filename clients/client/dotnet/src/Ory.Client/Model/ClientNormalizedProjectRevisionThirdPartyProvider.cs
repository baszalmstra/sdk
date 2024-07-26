/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v1.14.3
 * Contact: support@ory.sh
 * Generated by: https://github.com/openapitools/openapi-generator.git
 */


using System;
using System.Collections;
using System.Collections.Generic;
using System.Collections.ObjectModel;
using System.Linq;
using System.IO;
using System.Runtime.Serialization;
using System.Text;
using System.Text.RegularExpressions;
using Newtonsoft.Json;
using Newtonsoft.Json.Converters;
using Newtonsoft.Json.Linq;
using System.ComponentModel.DataAnnotations;
using OpenAPIDateConverter = Ory.Client.Client.OpenAPIDateConverter;

namespace Ory.Client.Model
{
    /// <summary>
    /// ClientNormalizedProjectRevisionThirdPartyProvider
    /// </summary>
    [DataContract(Name = "normalizedProjectRevisionThirdPartyProvider")]
    public partial class ClientNormalizedProjectRevisionThirdPartyProvider : IValidatableObject
    {
        /// <summary>
        /// State indicates the state of the provider  Only providers with state &#x60;enabled&#x60; will be used for authentication enabled ThirdPartyProviderStateEnabled disabled ThirdPartyProviderStateDisabled
        /// </summary>
        /// <value>State indicates the state of the provider  Only providers with state &#x60;enabled&#x60; will be used for authentication enabled ThirdPartyProviderStateEnabled disabled ThirdPartyProviderStateDisabled</value>
        [JsonConverter(typeof(StringEnumConverter))]
        public enum StateEnum
        {
            /// <summary>
            /// Enum Enabled for value: enabled
            /// </summary>
            [EnumMember(Value = "enabled")]
            Enabled = 1,

            /// <summary>
            /// Enum Disabled for value: disabled
            /// </summary>
            [EnumMember(Value = "disabled")]
            Disabled = 2
        }


        /// <summary>
        /// State indicates the state of the provider  Only providers with state &#x60;enabled&#x60; will be used for authentication enabled ThirdPartyProviderStateEnabled disabled ThirdPartyProviderStateDisabled
        /// </summary>
        /// <value>State indicates the state of the provider  Only providers with state &#x60;enabled&#x60; will be used for authentication enabled ThirdPartyProviderStateEnabled disabled ThirdPartyProviderStateDisabled</value>
        [DataMember(Name = "state", EmitDefaultValue = false)]
        public StateEnum? State { get; set; }
        /// <summary>
        /// Initializes a new instance of the <see cref="ClientNormalizedProjectRevisionThirdPartyProvider" /> class.
        /// </summary>
        /// <param name="additionalIdTokenAudiences">additionalIdTokenAudiences.</param>
        /// <param name="applePrivateKey">applePrivateKey.</param>
        /// <param name="applePrivateKeyId">Apple Private Key Identifier  Sign In with Apple Private Key Identifier needed for generating a JWT token for client secret.</param>
        /// <param name="appleTeamId">Apple Developer Team ID  Apple Developer Team ID needed for generating a JWT token for client secret.</param>
        /// <param name="authUrl">AuthURL is the authorize url, typically something like: https://example.org/oauth2/auth Should only be used when the OAuth2 / OpenID Connect server is not supporting OpenID Connect Discovery and when &#x60;provider&#x60; is set to &#x60;generic&#x60;..</param>
        /// <param name="azureTenant">Tenant is the Azure AD Tenant to use for authentication, and must be set when &#x60;provider&#x60; is set to &#x60;microsoft&#x60;.  Can be either &#x60;common&#x60;, &#x60;organizations&#x60;, &#x60;consumers&#x60; for a multitenant application or a specific tenant like &#x60;8eaef023-2b34-4da1-9baa-8bc8c9d6a490&#x60; or &#x60;contoso.onmicrosoft.com&#x60;..</param>
        /// <param name="claimsSource">claimsSource.</param>
        /// <param name="clientId">ClientID is the application&#39;s Client ID..</param>
        /// <param name="clientSecret">clientSecret.</param>
        /// <param name="id">id.</param>
        /// <param name="issuerUrl">IssuerURL is the OpenID Connect Server URL. You can leave this empty if &#x60;provider&#x60; is not set to &#x60;generic&#x60;. If set, neither &#x60;auth_url&#x60; nor &#x60;token_url&#x60; are required..</param>
        /// <param name="label">Label represents an optional label which can be used in the UI generation..</param>
        /// <param name="mapperUrl">Mapper specifies the JSONNet code snippet which uses the OpenID Connect Provider&#39;s data (e.g. GitHub or Google profile information) to hydrate the identity&#39;s data..</param>
        /// <param name="organizationId">organizationId.</param>
        /// <param name="projectRevisionId">The Revision&#39;s ID this schema belongs to.</param>
        /// <param name="provider">Provider is either \&quot;generic\&quot; for a generic OAuth 2.0 / OpenID Connect Provider or one of: generic google github gitlab microsoft discord slack facebook vk yandex apple.</param>
        /// <param name="providerId">ID is the provider&#39;s ID.</param>
        /// <param name="requestedClaims">requestedClaims.</param>
        /// <param name="scope">scope.</param>
        /// <param name="state">State indicates the state of the provider  Only providers with state &#x60;enabled&#x60; will be used for authentication enabled ThirdPartyProviderStateEnabled disabled ThirdPartyProviderStateDisabled.</param>
        /// <param name="subjectSource">subjectSource.</param>
        /// <param name="tokenUrl">TokenURL is the token url, typically something like: https://example.org/oauth2/token  Should only be used when the OAuth2 / OpenID Connect server is not supporting OpenID Connect Discovery and when &#x60;provider&#x60; is set to &#x60;generic&#x60;..</param>
        public ClientNormalizedProjectRevisionThirdPartyProvider(List<string> additionalIdTokenAudiences = default(List<string>), string applePrivateKey = default(string), string applePrivateKeyId = default(string), string appleTeamId = default(string), string authUrl = default(string), string azureTenant = default(string), string claimsSource = default(string), string clientId = default(string), string clientSecret = default(string), string id = default(string), string issuerUrl = default(string), string label = default(string), string mapperUrl = default(string), string organizationId = default(string), string projectRevisionId = default(string), string provider = default(string), string providerId = default(string), Object requestedClaims = default(Object), List<string> scope = default(List<string>), StateEnum? state = default(StateEnum?), string subjectSource = default(string), string tokenUrl = default(string))
        {
            this.AdditionalIdTokenAudiences = additionalIdTokenAudiences;
            this.ApplePrivateKey = applePrivateKey;
            this.ApplePrivateKeyId = applePrivateKeyId;
            this.AppleTeamId = appleTeamId;
            this.AuthUrl = authUrl;
            this.AzureTenant = azureTenant;
            this.ClaimsSource = claimsSource;
            this.ClientId = clientId;
            this.ClientSecret = clientSecret;
            this.Id = id;
            this.IssuerUrl = issuerUrl;
            this.Label = label;
            this.MapperUrl = mapperUrl;
            this.OrganizationId = organizationId;
            this.ProjectRevisionId = projectRevisionId;
            this.Provider = provider;
            this.ProviderId = providerId;
            this.RequestedClaims = requestedClaims;
            this.Scope = scope;
            this.State = state;
            this.SubjectSource = subjectSource;
            this.TokenUrl = tokenUrl;
            this.AdditionalProperties = new Dictionary<string, object>();
        }

        /// <summary>
        /// Gets or Sets AdditionalIdTokenAudiences
        /// </summary>
        [DataMember(Name = "additional_id_token_audiences", EmitDefaultValue = false)]
        public List<string> AdditionalIdTokenAudiences { get; set; }

        /// <summary>
        /// Gets or Sets ApplePrivateKey
        /// </summary>
        [DataMember(Name = "apple_private_key", EmitDefaultValue = true)]
        public string ApplePrivateKey { get; set; }

        /// <summary>
        /// Apple Private Key Identifier  Sign In with Apple Private Key Identifier needed for generating a JWT token for client secret
        /// </summary>
        /// <value>Apple Private Key Identifier  Sign In with Apple Private Key Identifier needed for generating a JWT token for client secret</value>
        /// <example>UX56C66723</example>
        [DataMember(Name = "apple_private_key_id", EmitDefaultValue = false)]
        public string ApplePrivateKeyId { get; set; }

        /// <summary>
        /// Apple Developer Team ID  Apple Developer Team ID needed for generating a JWT token for client secret
        /// </summary>
        /// <value>Apple Developer Team ID  Apple Developer Team ID needed for generating a JWT token for client secret</value>
        /// <example>KP76DQS54M</example>
        [DataMember(Name = "apple_team_id", EmitDefaultValue = false)]
        public string AppleTeamId { get; set; }

        /// <summary>
        /// AuthURL is the authorize url, typically something like: https://example.org/oauth2/auth Should only be used when the OAuth2 / OpenID Connect server is not supporting OpenID Connect Discovery and when &#x60;provider&#x60; is set to &#x60;generic&#x60;.
        /// </summary>
        /// <value>AuthURL is the authorize url, typically something like: https://example.org/oauth2/auth Should only be used when the OAuth2 / OpenID Connect server is not supporting OpenID Connect Discovery and when &#x60;provider&#x60; is set to &#x60;generic&#x60;.</value>
        /// <example>https://www.googleapis.com/oauth2/v2/auth</example>
        [DataMember(Name = "auth_url", EmitDefaultValue = false)]
        public string AuthUrl { get; set; }

        /// <summary>
        /// Tenant is the Azure AD Tenant to use for authentication, and must be set when &#x60;provider&#x60; is set to &#x60;microsoft&#x60;.  Can be either &#x60;common&#x60;, &#x60;organizations&#x60;, &#x60;consumers&#x60; for a multitenant application or a specific tenant like &#x60;8eaef023-2b34-4da1-9baa-8bc8c9d6a490&#x60; or &#x60;contoso.onmicrosoft.com&#x60;.
        /// </summary>
        /// <value>Tenant is the Azure AD Tenant to use for authentication, and must be set when &#x60;provider&#x60; is set to &#x60;microsoft&#x60;.  Can be either &#x60;common&#x60;, &#x60;organizations&#x60;, &#x60;consumers&#x60; for a multitenant application or a specific tenant like &#x60;8eaef023-2b34-4da1-9baa-8bc8c9d6a490&#x60; or &#x60;contoso.onmicrosoft.com&#x60;.</value>
        /// <example>contoso.onmicrosoft.com</example>
        [DataMember(Name = "azure_tenant", EmitDefaultValue = false)]
        public string AzureTenant { get; set; }

        /// <summary>
        /// Gets or Sets ClaimsSource
        /// </summary>
        [DataMember(Name = "claims_source", EmitDefaultValue = true)]
        public string ClaimsSource { get; set; }

        /// <summary>
        /// ClientID is the application&#39;s Client ID.
        /// </summary>
        /// <value>ClientID is the application&#39;s Client ID.</value>
        [DataMember(Name = "client_id", EmitDefaultValue = false)]
        public string ClientId { get; set; }

        /// <summary>
        /// Gets or Sets ClientSecret
        /// </summary>
        [DataMember(Name = "client_secret", EmitDefaultValue = true)]
        public string ClientSecret { get; set; }

        /// <summary>
        /// The Project&#39;s Revision Creation Date
        /// </summary>
        /// <value>The Project&#39;s Revision Creation Date</value>
        [DataMember(Name = "created_at", EmitDefaultValue = false)]
        public DateTime CreatedAt { get; private set; }

        /// <summary>
        /// Returns false as CreatedAt should not be serialized given that it's read-only.
        /// </summary>
        /// <returns>false (boolean)</returns>
        public bool ShouldSerializeCreatedAt()
        {
            return false;
        }
        /// <summary>
        /// Gets or Sets Id
        /// </summary>
        [DataMember(Name = "id", EmitDefaultValue = false)]
        public string Id { get; set; }

        /// <summary>
        /// IssuerURL is the OpenID Connect Server URL. You can leave this empty if &#x60;provider&#x60; is not set to &#x60;generic&#x60;. If set, neither &#x60;auth_url&#x60; nor &#x60;token_url&#x60; are required.
        /// </summary>
        /// <value>IssuerURL is the OpenID Connect Server URL. You can leave this empty if &#x60;provider&#x60; is not set to &#x60;generic&#x60;. If set, neither &#x60;auth_url&#x60; nor &#x60;token_url&#x60; are required.</value>
        /// <example>https://accounts.google.com</example>
        [DataMember(Name = "issuer_url", EmitDefaultValue = false)]
        public string IssuerUrl { get; set; }

        /// <summary>
        /// Label represents an optional label which can be used in the UI generation.
        /// </summary>
        /// <value>Label represents an optional label which can be used in the UI generation.</value>
        [DataMember(Name = "label", EmitDefaultValue = false)]
        public string Label { get; set; }

        /// <summary>
        /// Mapper specifies the JSONNet code snippet which uses the OpenID Connect Provider&#39;s data (e.g. GitHub or Google profile information) to hydrate the identity&#39;s data.
        /// </summary>
        /// <value>Mapper specifies the JSONNet code snippet which uses the OpenID Connect Provider&#39;s data (e.g. GitHub or Google profile information) to hydrate the identity&#39;s data.</value>
        [DataMember(Name = "mapper_url", EmitDefaultValue = false)]
        public string MapperUrl { get; set; }

        /// <summary>
        /// Gets or Sets OrganizationId
        /// </summary>
        [DataMember(Name = "organization_id", EmitDefaultValue = true)]
        public string OrganizationId { get; set; }

        /// <summary>
        /// The Revision&#39;s ID this schema belongs to
        /// </summary>
        /// <value>The Revision&#39;s ID this schema belongs to</value>
        [DataMember(Name = "project_revision_id", EmitDefaultValue = false)]
        public string ProjectRevisionId { get; set; }

        /// <summary>
        /// Provider is either \&quot;generic\&quot; for a generic OAuth 2.0 / OpenID Connect Provider or one of: generic google github gitlab microsoft discord slack facebook vk yandex apple
        /// </summary>
        /// <value>Provider is either \&quot;generic\&quot; for a generic OAuth 2.0 / OpenID Connect Provider or one of: generic google github gitlab microsoft discord slack facebook vk yandex apple</value>
        /// <example>google</example>
        [DataMember(Name = "provider", EmitDefaultValue = false)]
        public string Provider { get; set; }

        /// <summary>
        /// ID is the provider&#39;s ID
        /// </summary>
        /// <value>ID is the provider&#39;s ID</value>
        [DataMember(Name = "provider_id", EmitDefaultValue = false)]
        public string ProviderId { get; set; }

        /// <summary>
        /// Gets or Sets RequestedClaims
        /// </summary>
        [DataMember(Name = "requested_claims", EmitDefaultValue = false)]
        public Object RequestedClaims { get; set; }

        /// <summary>
        /// Gets or Sets Scope
        /// </summary>
        [DataMember(Name = "scope", EmitDefaultValue = false)]
        public List<string> Scope { get; set; }

        /// <summary>
        /// Gets or Sets SubjectSource
        /// </summary>
        [DataMember(Name = "subject_source", EmitDefaultValue = true)]
        public string SubjectSource { get; set; }

        /// <summary>
        /// TokenURL is the token url, typically something like: https://example.org/oauth2/token  Should only be used when the OAuth2 / OpenID Connect server is not supporting OpenID Connect Discovery and when &#x60;provider&#x60; is set to &#x60;generic&#x60;.
        /// </summary>
        /// <value>TokenURL is the token url, typically something like: https://example.org/oauth2/token  Should only be used when the OAuth2 / OpenID Connect server is not supporting OpenID Connect Discovery and when &#x60;provider&#x60; is set to &#x60;generic&#x60;.</value>
        /// <example>https://www.googleapis.com/oauth2/v4/token</example>
        [DataMember(Name = "token_url", EmitDefaultValue = false)]
        public string TokenUrl { get; set; }

        /// <summary>
        /// Last Time Project&#39;s Revision was Updated
        /// </summary>
        /// <value>Last Time Project&#39;s Revision was Updated</value>
        [DataMember(Name = "updated_at", EmitDefaultValue = false)]
        public DateTime UpdatedAt { get; private set; }

        /// <summary>
        /// Returns false as UpdatedAt should not be serialized given that it's read-only.
        /// </summary>
        /// <returns>false (boolean)</returns>
        public bool ShouldSerializeUpdatedAt()
        {
            return false;
        }
        /// <summary>
        /// Gets or Sets additional properties
        /// </summary>
        [JsonExtensionData]
        public IDictionary<string, object> AdditionalProperties { get; set; }

        /// <summary>
        /// Returns the string presentation of the object
        /// </summary>
        /// <returns>String presentation of the object</returns>
        public override string ToString()
        {
            StringBuilder sb = new StringBuilder();
            sb.Append("class ClientNormalizedProjectRevisionThirdPartyProvider {\n");
            sb.Append("  AdditionalIdTokenAudiences: ").Append(AdditionalIdTokenAudiences).Append("\n");
            sb.Append("  ApplePrivateKey: ").Append(ApplePrivateKey).Append("\n");
            sb.Append("  ApplePrivateKeyId: ").Append(ApplePrivateKeyId).Append("\n");
            sb.Append("  AppleTeamId: ").Append(AppleTeamId).Append("\n");
            sb.Append("  AuthUrl: ").Append(AuthUrl).Append("\n");
            sb.Append("  AzureTenant: ").Append(AzureTenant).Append("\n");
            sb.Append("  ClaimsSource: ").Append(ClaimsSource).Append("\n");
            sb.Append("  ClientId: ").Append(ClientId).Append("\n");
            sb.Append("  ClientSecret: ").Append(ClientSecret).Append("\n");
            sb.Append("  CreatedAt: ").Append(CreatedAt).Append("\n");
            sb.Append("  Id: ").Append(Id).Append("\n");
            sb.Append("  IssuerUrl: ").Append(IssuerUrl).Append("\n");
            sb.Append("  Label: ").Append(Label).Append("\n");
            sb.Append("  MapperUrl: ").Append(MapperUrl).Append("\n");
            sb.Append("  OrganizationId: ").Append(OrganizationId).Append("\n");
            sb.Append("  ProjectRevisionId: ").Append(ProjectRevisionId).Append("\n");
            sb.Append("  Provider: ").Append(Provider).Append("\n");
            sb.Append("  ProviderId: ").Append(ProviderId).Append("\n");
            sb.Append("  RequestedClaims: ").Append(RequestedClaims).Append("\n");
            sb.Append("  Scope: ").Append(Scope).Append("\n");
            sb.Append("  State: ").Append(State).Append("\n");
            sb.Append("  SubjectSource: ").Append(SubjectSource).Append("\n");
            sb.Append("  TokenUrl: ").Append(TokenUrl).Append("\n");
            sb.Append("  UpdatedAt: ").Append(UpdatedAt).Append("\n");
            sb.Append("  AdditionalProperties: ").Append(AdditionalProperties).Append("\n");
            sb.Append("}\n");
            return sb.ToString();
        }

        /// <summary>
        /// Returns the JSON string presentation of the object
        /// </summary>
        /// <returns>JSON string presentation of the object</returns>
        public virtual string ToJson()
        {
            return Newtonsoft.Json.JsonConvert.SerializeObject(this, Newtonsoft.Json.Formatting.Indented);
        }

        /// <summary>
        /// To validate all properties of the instance
        /// </summary>
        /// <param name="validationContext">Validation context</param>
        /// <returns>Validation Result</returns>
        IEnumerable<ValidationResult> IValidatableObject.Validate(ValidationContext validationContext)
        {
            yield break;
        }
    }

}

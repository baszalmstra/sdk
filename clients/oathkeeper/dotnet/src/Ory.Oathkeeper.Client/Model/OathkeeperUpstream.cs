/*
 * ORY Oathkeeper
 *
 * ORY Oathkeeper is a reverse proxy that checks the HTTP Authorization for validity against a set of rules. This service uses Hydra to validate access tokens and policies.
 *
 * The version of the OpenAPI document: v0.0.0-alpha.52
 * Contact: hi@ory.am
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
using OpenAPIDateConverter = Ory.Oathkeeper.Client.Client.OpenAPIDateConverter;

namespace Ory.Oathkeeper.Client.Model
{
    /// <summary>
    /// Upstream Upstream Upstream Upstream Upstream Upstream Upstream upstream
    /// </summary>
    [DataContract(Name = "Upstream")]
    public partial class OathkeeperUpstream : IEquatable<OathkeeperUpstream>, IValidatableObject
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="OathkeeperUpstream" /> class.
        /// </summary>
        /// <param name="preserveHost">PreserveHost, if false (the default), tells ORY Oathkeeper to set the upstream request&#39;s Host header to the hostname of the API&#39;s upstream&#39;s URL. Setting this flag to true instructs ORY Oathkeeper not to do so..</param>
        /// <param name="stripPath">StripPath if set, replaces the provided path prefix when forwarding the requested URL to the upstream URL..</param>
        /// <param name="url">URL is the URL the request will be proxied to..</param>
        public OathkeeperUpstream(bool preserveHost = default(bool), string stripPath = default(string), string url = default(string))
        {
            this.PreserveHost = preserveHost;
            this.StripPath = stripPath;
            this.Url = url;
            this.AdditionalProperties = new Dictionary<string, object>();
        }

        /// <summary>
        /// PreserveHost, if false (the default), tells ORY Oathkeeper to set the upstream request&#39;s Host header to the hostname of the API&#39;s upstream&#39;s URL. Setting this flag to true instructs ORY Oathkeeper not to do so.
        /// </summary>
        /// <value>PreserveHost, if false (the default), tells ORY Oathkeeper to set the upstream request&#39;s Host header to the hostname of the API&#39;s upstream&#39;s URL. Setting this flag to true instructs ORY Oathkeeper not to do so.</value>
        [DataMember(Name = "preserve_host", EmitDefaultValue = true)]
        public bool PreserveHost { get; set; }

        /// <summary>
        /// StripPath if set, replaces the provided path prefix when forwarding the requested URL to the upstream URL.
        /// </summary>
        /// <value>StripPath if set, replaces the provided path prefix when forwarding the requested URL to the upstream URL.</value>
        [DataMember(Name = "strip_path", EmitDefaultValue = false)]
        public string StripPath { get; set; }

        /// <summary>
        /// URL is the URL the request will be proxied to.
        /// </summary>
        /// <value>URL is the URL the request will be proxied to.</value>
        [DataMember(Name = "url", EmitDefaultValue = false)]
        public string Url { get; set; }

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
            var sb = new StringBuilder();
            sb.Append("class OathkeeperUpstream {\n");
            sb.Append("  PreserveHost: ").Append(PreserveHost).Append("\n");
            sb.Append("  StripPath: ").Append(StripPath).Append("\n");
            sb.Append("  Url: ").Append(Url).Append("\n");
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
        /// Returns true if objects are equal
        /// </summary>
        /// <param name="input">Object to be compared</param>
        /// <returns>Boolean</returns>
        public override bool Equals(object input)
        {
            return this.Equals(input as OathkeeperUpstream);
        }

        /// <summary>
        /// Returns true if OathkeeperUpstream instances are equal
        /// </summary>
        /// <param name="input">Instance of OathkeeperUpstream to be compared</param>
        /// <returns>Boolean</returns>
        public bool Equals(OathkeeperUpstream input)
        {
            if (input == null)
                return false;

            return 
                (
                    this.PreserveHost == input.PreserveHost ||
                    this.PreserveHost.Equals(input.PreserveHost)
                ) && 
                (
                    this.StripPath == input.StripPath ||
                    (this.StripPath != null &&
                    this.StripPath.Equals(input.StripPath))
                ) && 
                (
                    this.Url == input.Url ||
                    (this.Url != null &&
                    this.Url.Equals(input.Url))
                )
                && (this.AdditionalProperties.Count == input.AdditionalProperties.Count && !this.AdditionalProperties.Except(input.AdditionalProperties).Any());
        }

        /// <summary>
        /// Gets the hash code
        /// </summary>
        /// <returns>Hash code</returns>
        public override int GetHashCode()
        {
            unchecked // Overflow is fine, just wrap
            {
                int hashCode = 41;
                hashCode = hashCode * 59 + this.PreserveHost.GetHashCode();
                if (this.StripPath != null)
                    hashCode = hashCode * 59 + this.StripPath.GetHashCode();
                if (this.Url != null)
                    hashCode = hashCode * 59 + this.Url.GetHashCode();
                if (this.AdditionalProperties != null)
                    hashCode = hashCode * 59 + this.AdditionalProperties.GetHashCode();
                return hashCode;
            }
        }

        /// <summary>
        /// To validate all properties of the instance
        /// </summary>
        /// <param name="validationContext">Validation context</param>
        /// <returns>Validation Result</returns>
        IEnumerable<System.ComponentModel.DataAnnotations.ValidationResult> IValidatableObject.Validate(ValidationContext validationContext)
        {
            yield break;
        }
    }

}

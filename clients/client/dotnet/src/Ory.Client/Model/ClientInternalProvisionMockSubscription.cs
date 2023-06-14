/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v1.1.37
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
    /// Internal Provision Mock Subscription Request Body
    /// </summary>
    [DataContract(Name = "internalProvisionMockSubscription")]
    public partial class ClientInternalProvisionMockSubscription : IEquatable<ClientInternalProvisionMockSubscription>, IValidatableObject
    {
        /// <summary>
        /// Currency usd USD eur Euro
        /// </summary>
        /// <value>Currency usd USD eur Euro</value>
        [JsonConverter(typeof(StringEnumConverter))]
        public enum CurrencyEnum
        {
            /// <summary>
            /// Enum Usd for value: usd
            /// </summary>
            [EnumMember(Value = "usd")]
            Usd = 1,

            /// <summary>
            /// Enum Eur for value: eur
            /// </summary>
            [EnumMember(Value = "eur")]
            Eur = 2

        }


        /// <summary>
        /// Currency usd USD eur Euro
        /// </summary>
        /// <value>Currency usd USD eur Euro</value>
        [DataMember(Name = "currency", IsRequired = true, EmitDefaultValue = false)]
        public CurrencyEnum Currency { get; set; }
        /// <summary>
        /// Billing Interval monthly Monthly yearly Yearly
        /// </summary>
        /// <value>Billing Interval monthly Monthly yearly Yearly</value>
        [JsonConverter(typeof(StringEnumConverter))]
        public enum IntervalEnum
        {
            /// <summary>
            /// Enum Monthly for value: monthly
            /// </summary>
            [EnumMember(Value = "monthly")]
            Monthly = 1,

            /// <summary>
            /// Enum Yearly for value: yearly
            /// </summary>
            [EnumMember(Value = "yearly")]
            Yearly = 2

        }


        /// <summary>
        /// Billing Interval monthly Monthly yearly Yearly
        /// </summary>
        /// <value>Billing Interval monthly Monthly yearly Yearly</value>
        [DataMember(Name = "interval", IsRequired = true, EmitDefaultValue = false)]
        public IntervalEnum Interval { get; set; }
        /// <summary>
        /// Initializes a new instance of the <see cref="ClientInternalProvisionMockSubscription" /> class.
        /// </summary>
        [JsonConstructorAttribute]
        protected ClientInternalProvisionMockSubscription()
        {
            this.AdditionalProperties = new Dictionary<string, object>();
        }
        /// <summary>
        /// Initializes a new instance of the <see cref="ClientInternalProvisionMockSubscription" /> class.
        /// </summary>
        /// <param name="currency">Currency usd USD eur Euro (required).</param>
        /// <param name="identityId">Identity ID (required).</param>
        /// <param name="interval">Billing Interval monthly Monthly yearly Yearly (required).</param>
        /// <param name="plan">Plan ID (required).</param>
        public ClientInternalProvisionMockSubscription(CurrencyEnum currency = default(CurrencyEnum), string identityId = default(string), IntervalEnum interval = default(IntervalEnum), string plan = default(string))
        {
            this.Currency = currency;
            // to ensure "identityId" is required (not null)
            if (identityId == null) {
                throw new ArgumentNullException("identityId is a required property for ClientInternalProvisionMockSubscription and cannot be null");
            }
            this.IdentityId = identityId;
            this.Interval = interval;
            // to ensure "plan" is required (not null)
            if (plan == null) {
                throw new ArgumentNullException("plan is a required property for ClientInternalProvisionMockSubscription and cannot be null");
            }
            this.Plan = plan;
            this.AdditionalProperties = new Dictionary<string, object>();
        }

        /// <summary>
        /// Identity ID
        /// </summary>
        /// <value>Identity ID</value>
        [DataMember(Name = "identity_id", IsRequired = true, EmitDefaultValue = false)]
        public string IdentityId { get; set; }

        /// <summary>
        /// Plan ID
        /// </summary>
        /// <value>Plan ID</value>
        [DataMember(Name = "plan", IsRequired = true, EmitDefaultValue = false)]
        public string Plan { get; set; }

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
            sb.Append("class ClientInternalProvisionMockSubscription {\n");
            sb.Append("  Currency: ").Append(Currency).Append("\n");
            sb.Append("  IdentityId: ").Append(IdentityId).Append("\n");
            sb.Append("  Interval: ").Append(Interval).Append("\n");
            sb.Append("  Plan: ").Append(Plan).Append("\n");
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
            return this.Equals(input as ClientInternalProvisionMockSubscription);
        }

        /// <summary>
        /// Returns true if ClientInternalProvisionMockSubscription instances are equal
        /// </summary>
        /// <param name="input">Instance of ClientInternalProvisionMockSubscription to be compared</param>
        /// <returns>Boolean</returns>
        public bool Equals(ClientInternalProvisionMockSubscription input)
        {
            if (input == null)
            {
                return false;
            }
            return 
                (
                    this.Currency == input.Currency ||
                    this.Currency.Equals(input.Currency)
                ) && 
                (
                    this.IdentityId == input.IdentityId ||
                    (this.IdentityId != null &&
                    this.IdentityId.Equals(input.IdentityId))
                ) && 
                (
                    this.Interval == input.Interval ||
                    this.Interval.Equals(input.Interval)
                ) && 
                (
                    this.Plan == input.Plan ||
                    (this.Plan != null &&
                    this.Plan.Equals(input.Plan))
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
                hashCode = (hashCode * 59) + this.Currency.GetHashCode();
                if (this.IdentityId != null)
                {
                    hashCode = (hashCode * 59) + this.IdentityId.GetHashCode();
                }
                hashCode = (hashCode * 59) + this.Interval.GetHashCode();
                if (this.Plan != null)
                {
                    hashCode = (hashCode * 59) + this.Plan.GetHashCode();
                }
                if (this.AdditionalProperties != null)
                {
                    hashCode = (hashCode * 59) + this.AdditionalProperties.GetHashCode();
                }
                return hashCode;
            }
        }

        /// <summary>
        /// To validate all properties of the instance
        /// </summary>
        /// <param name="validationContext">Validation context</param>
        /// <returns>Validation Result</returns>
        public IEnumerable<System.ComponentModel.DataAnnotations.ValidationResult> Validate(ValidationContext validationContext)
        {
            yield break;
        }
    }

}

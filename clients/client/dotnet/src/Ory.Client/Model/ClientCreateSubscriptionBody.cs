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
    /// Create Subscription Request Body
    /// </summary>
    [DataContract(Name = "createSubscriptionBody")]
    public partial class ClientCreateSubscriptionBody : IEquatable<ClientCreateSubscriptionBody>, IValidatableObject
    {
        /// <summary>
        ///  usd USD eur Euro
        /// </summary>
        /// <value> usd USD eur Euro</value>
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
        ///  usd USD eur Euro
        /// </summary>
        /// <value> usd USD eur Euro</value>
        [DataMember(Name = "currency", EmitDefaultValue = false)]
        public CurrencyEnum? Currency { get; set; }
        /// <summary>
        ///  monthly Monthly yearly Yearly
        /// </summary>
        /// <value> monthly Monthly yearly Yearly</value>
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
        ///  monthly Monthly yearly Yearly
        /// </summary>
        /// <value> monthly Monthly yearly Yearly</value>
        [DataMember(Name = "interval", IsRequired = true, EmitDefaultValue = false)]
        public IntervalEnum Interval { get; set; }
        /// <summary>
        /// Initializes a new instance of the <see cref="ClientCreateSubscriptionBody" /> class.
        /// </summary>
        [JsonConstructorAttribute]
        protected ClientCreateSubscriptionBody()
        {
            this.AdditionalProperties = new Dictionary<string, object>();
        }
        /// <summary>
        /// Initializes a new instance of the <see cref="ClientCreateSubscriptionBody" /> class.
        /// </summary>
        /// <param name="currency"> usd USD eur Euro.</param>
        /// <param name="interval"> monthly Monthly yearly Yearly (required).</param>
        /// <param name="plan">plan (required).</param>
        /// <param name="provisionFirstProject">provisionFirstProject (required).</param>
        /// <param name="returnTo">returnTo.</param>
        public ClientCreateSubscriptionBody(CurrencyEnum? currency = default(CurrencyEnum?), IntervalEnum interval = default(IntervalEnum), string plan = default(string), string provisionFirstProject = default(string), string returnTo = default(string))
        {
            this.Interval = interval;
            // to ensure "plan" is required (not null)
            if (plan == null) {
                throw new ArgumentNullException("plan is a required property for ClientCreateSubscriptionBody and cannot be null");
            }
            this.Plan = plan;
            // to ensure "provisionFirstProject" is required (not null)
            if (provisionFirstProject == null) {
                throw new ArgumentNullException("provisionFirstProject is a required property for ClientCreateSubscriptionBody and cannot be null");
            }
            this.ProvisionFirstProject = provisionFirstProject;
            this.Currency = currency;
            this.ReturnTo = returnTo;
            this.AdditionalProperties = new Dictionary<string, object>();
        }

        /// <summary>
        /// Gets or Sets Plan
        /// </summary>
        [DataMember(Name = "plan", IsRequired = true, EmitDefaultValue = false)]
        public string Plan { get; set; }

        /// <summary>
        /// Gets or Sets ProvisionFirstProject
        /// </summary>
        [DataMember(Name = "provision_first_project", IsRequired = true, EmitDefaultValue = false)]
        public string ProvisionFirstProject { get; set; }

        /// <summary>
        /// Gets or Sets ReturnTo
        /// </summary>
        [DataMember(Name = "return_to", EmitDefaultValue = false)]
        public string ReturnTo { get; set; }

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
            sb.Append("class ClientCreateSubscriptionBody {\n");
            sb.Append("  Currency: ").Append(Currency).Append("\n");
            sb.Append("  Interval: ").Append(Interval).Append("\n");
            sb.Append("  Plan: ").Append(Plan).Append("\n");
            sb.Append("  ProvisionFirstProject: ").Append(ProvisionFirstProject).Append("\n");
            sb.Append("  ReturnTo: ").Append(ReturnTo).Append("\n");
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
            return this.Equals(input as ClientCreateSubscriptionBody);
        }

        /// <summary>
        /// Returns true if ClientCreateSubscriptionBody instances are equal
        /// </summary>
        /// <param name="input">Instance of ClientCreateSubscriptionBody to be compared</param>
        /// <returns>Boolean</returns>
        public bool Equals(ClientCreateSubscriptionBody input)
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
                    this.Interval == input.Interval ||
                    this.Interval.Equals(input.Interval)
                ) && 
                (
                    this.Plan == input.Plan ||
                    (this.Plan != null &&
                    this.Plan.Equals(input.Plan))
                ) && 
                (
                    this.ProvisionFirstProject == input.ProvisionFirstProject ||
                    (this.ProvisionFirstProject != null &&
                    this.ProvisionFirstProject.Equals(input.ProvisionFirstProject))
                ) && 
                (
                    this.ReturnTo == input.ReturnTo ||
                    (this.ReturnTo != null &&
                    this.ReturnTo.Equals(input.ReturnTo))
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
                hashCode = (hashCode * 59) + this.Interval.GetHashCode();
                if (this.Plan != null)
                {
                    hashCode = (hashCode * 59) + this.Plan.GetHashCode();
                }
                if (this.ProvisionFirstProject != null)
                {
                    hashCode = (hashCode * 59) + this.ProvisionFirstProject.GetHashCode();
                }
                if (this.ReturnTo != null)
                {
                    hashCode = (hashCode * 59) + this.ReturnTo.GetHashCode();
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

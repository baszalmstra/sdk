/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v1.4.5
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
    /// ClientMessage
    /// </summary>
    [DataContract(Name = "message")]
    public partial class ClientMessage : IEquatable<ClientMessage>, IValidatableObject
    {
        /// <summary>
        ///  recovery_invalid TypeRecoveryInvalid recovery_valid TypeRecoveryValid recovery_code_invalid TypeRecoveryCodeInvalid recovery_code_valid TypeRecoveryCodeValid verification_invalid TypeVerificationInvalid verification_valid TypeVerificationValid verification_code_invalid TypeVerificationCodeInvalid verification_code_valid TypeVerificationCodeValid otp TypeOTP stub TypeTestStub login_code_valid TypeLoginCodeValid registration_code_valid TypeRegistrationCodeValid
        /// </summary>
        /// <value> recovery_invalid TypeRecoveryInvalid recovery_valid TypeRecoveryValid recovery_code_invalid TypeRecoveryCodeInvalid recovery_code_valid TypeRecoveryCodeValid verification_invalid TypeVerificationInvalid verification_valid TypeVerificationValid verification_code_invalid TypeVerificationCodeInvalid verification_code_valid TypeVerificationCodeValid otp TypeOTP stub TypeTestStub login_code_valid TypeLoginCodeValid registration_code_valid TypeRegistrationCodeValid</value>
        [JsonConverter(typeof(StringEnumConverter))]
        public enum TemplateTypeEnum
        {
            /// <summary>
            /// Enum RecoveryInvalid for value: recovery_invalid
            /// </summary>
            [EnumMember(Value = "recovery_invalid")]
            RecoveryInvalid = 1,

            /// <summary>
            /// Enum RecoveryValid for value: recovery_valid
            /// </summary>
            [EnumMember(Value = "recovery_valid")]
            RecoveryValid = 2,

            /// <summary>
            /// Enum RecoveryCodeInvalid for value: recovery_code_invalid
            /// </summary>
            [EnumMember(Value = "recovery_code_invalid")]
            RecoveryCodeInvalid = 3,

            /// <summary>
            /// Enum RecoveryCodeValid for value: recovery_code_valid
            /// </summary>
            [EnumMember(Value = "recovery_code_valid")]
            RecoveryCodeValid = 4,

            /// <summary>
            /// Enum VerificationInvalid for value: verification_invalid
            /// </summary>
            [EnumMember(Value = "verification_invalid")]
            VerificationInvalid = 5,

            /// <summary>
            /// Enum VerificationValid for value: verification_valid
            /// </summary>
            [EnumMember(Value = "verification_valid")]
            VerificationValid = 6,

            /// <summary>
            /// Enum VerificationCodeInvalid for value: verification_code_invalid
            /// </summary>
            [EnumMember(Value = "verification_code_invalid")]
            VerificationCodeInvalid = 7,

            /// <summary>
            /// Enum VerificationCodeValid for value: verification_code_valid
            /// </summary>
            [EnumMember(Value = "verification_code_valid")]
            VerificationCodeValid = 8,

            /// <summary>
            /// Enum Otp for value: otp
            /// </summary>
            [EnumMember(Value = "otp")]
            Otp = 9,

            /// <summary>
            /// Enum Stub for value: stub
            /// </summary>
            [EnumMember(Value = "stub")]
            Stub = 10,

            /// <summary>
            /// Enum LoginCodeValid for value: login_code_valid
            /// </summary>
            [EnumMember(Value = "login_code_valid")]
            LoginCodeValid = 11,

            /// <summary>
            /// Enum RegistrationCodeValid for value: registration_code_valid
            /// </summary>
            [EnumMember(Value = "registration_code_valid")]
            RegistrationCodeValid = 12

        }


        /// <summary>
        ///  recovery_invalid TypeRecoveryInvalid recovery_valid TypeRecoveryValid recovery_code_invalid TypeRecoveryCodeInvalid recovery_code_valid TypeRecoveryCodeValid verification_invalid TypeVerificationInvalid verification_valid TypeVerificationValid verification_code_invalid TypeVerificationCodeInvalid verification_code_valid TypeVerificationCodeValid otp TypeOTP stub TypeTestStub login_code_valid TypeLoginCodeValid registration_code_valid TypeRegistrationCodeValid
        /// </summary>
        /// <value> recovery_invalid TypeRecoveryInvalid recovery_valid TypeRecoveryValid recovery_code_invalid TypeRecoveryCodeInvalid recovery_code_valid TypeRecoveryCodeValid verification_invalid TypeVerificationInvalid verification_valid TypeVerificationValid verification_code_invalid TypeVerificationCodeInvalid verification_code_valid TypeVerificationCodeValid otp TypeOTP stub TypeTestStub login_code_valid TypeLoginCodeValid registration_code_valid TypeRegistrationCodeValid</value>
        [DataMember(Name = "template_type", IsRequired = true, EmitDefaultValue = false)]
        public TemplateTypeEnum TemplateType { get; set; }
        /// <summary>
        /// Initializes a new instance of the <see cref="ClientMessage" /> class.
        /// </summary>
        [JsonConstructorAttribute]
        protected ClientMessage()
        {
            this.AdditionalProperties = new Dictionary<string, object>();
        }
        /// <summary>
        /// Initializes a new instance of the <see cref="ClientMessage" /> class.
        /// </summary>
        /// <param name="body">body (required).</param>
        /// <param name="createdAt">CreatedAt is a helper struct field for gobuffalo.pop. (required).</param>
        /// <param name="dispatches">Dispatches store information about the attempts of delivering a message May contain an error if any happened, or just the &#x60;success&#x60; state..</param>
        /// <param name="id">id (required).</param>
        /// <param name="recipient">recipient (required).</param>
        /// <param name="sendCount">sendCount (required).</param>
        /// <param name="status">status (required).</param>
        /// <param name="subject">subject (required).</param>
        /// <param name="templateType"> recovery_invalid TypeRecoveryInvalid recovery_valid TypeRecoveryValid recovery_code_invalid TypeRecoveryCodeInvalid recovery_code_valid TypeRecoveryCodeValid verification_invalid TypeVerificationInvalid verification_valid TypeVerificationValid verification_code_invalid TypeVerificationCodeInvalid verification_code_valid TypeVerificationCodeValid otp TypeOTP stub TypeTestStub login_code_valid TypeLoginCodeValid registration_code_valid TypeRegistrationCodeValid (required).</param>
        /// <param name="type">type (required).</param>
        /// <param name="updatedAt">UpdatedAt is a helper struct field for gobuffalo.pop. (required).</param>
        public ClientMessage(string body = default(string), DateTime createdAt = default(DateTime), List<ClientMessageDispatch> dispatches = default(List<ClientMessageDispatch>), string id = default(string), string recipient = default(string), long sendCount = default(long), ClientCourierMessageStatus status = default(ClientCourierMessageStatus), string subject = default(string), TemplateTypeEnum templateType = default(TemplateTypeEnum), ClientCourierMessageType type = default(ClientCourierMessageType), DateTime updatedAt = default(DateTime))
        {
            // to ensure "body" is required (not null)
            if (body == null) {
                throw new ArgumentNullException("body is a required property for ClientMessage and cannot be null");
            }
            this.Body = body;
            this.CreatedAt = createdAt;
            // to ensure "id" is required (not null)
            if (id == null) {
                throw new ArgumentNullException("id is a required property for ClientMessage and cannot be null");
            }
            this.Id = id;
            // to ensure "recipient" is required (not null)
            if (recipient == null) {
                throw new ArgumentNullException("recipient is a required property for ClientMessage and cannot be null");
            }
            this.Recipient = recipient;
            this.SendCount = sendCount;
            // to ensure "status" is required (not null)
            if (status == null) {
                throw new ArgumentNullException("status is a required property for ClientMessage and cannot be null");
            }
            this.Status = status;
            // to ensure "subject" is required (not null)
            if (subject == null) {
                throw new ArgumentNullException("subject is a required property for ClientMessage and cannot be null");
            }
            this.Subject = subject;
            this.TemplateType = templateType;
            // to ensure "type" is required (not null)
            if (type == null) {
                throw new ArgumentNullException("type is a required property for ClientMessage and cannot be null");
            }
            this.Type = type;
            this.UpdatedAt = updatedAt;
            this.Dispatches = dispatches;
            this.AdditionalProperties = new Dictionary<string, object>();
        }

        /// <summary>
        /// Gets or Sets Body
        /// </summary>
        [DataMember(Name = "body", IsRequired = true, EmitDefaultValue = false)]
        public string Body { get; set; }

        /// <summary>
        /// CreatedAt is a helper struct field for gobuffalo.pop.
        /// </summary>
        /// <value>CreatedAt is a helper struct field for gobuffalo.pop.</value>
        [DataMember(Name = "created_at", IsRequired = true, EmitDefaultValue = false)]
        public DateTime CreatedAt { get; set; }

        /// <summary>
        /// Dispatches store information about the attempts of delivering a message May contain an error if any happened, or just the &#x60;success&#x60; state.
        /// </summary>
        /// <value>Dispatches store information about the attempts of delivering a message May contain an error if any happened, or just the &#x60;success&#x60; state.</value>
        [DataMember(Name = "dispatches", EmitDefaultValue = false)]
        public List<ClientMessageDispatch> Dispatches { get; set; }

        /// <summary>
        /// Gets or Sets Id
        /// </summary>
        [DataMember(Name = "id", IsRequired = true, EmitDefaultValue = false)]
        public string Id { get; set; }

        /// <summary>
        /// Gets or Sets Recipient
        /// </summary>
        [DataMember(Name = "recipient", IsRequired = true, EmitDefaultValue = false)]
        public string Recipient { get; set; }

        /// <summary>
        /// Gets or Sets SendCount
        /// </summary>
        [DataMember(Name = "send_count", IsRequired = true, EmitDefaultValue = false)]
        public long SendCount { get; set; }

        /// <summary>
        /// Gets or Sets Status
        /// </summary>
        [DataMember(Name = "status", IsRequired = true, EmitDefaultValue = false)]
        public ClientCourierMessageStatus Status { get; set; }

        /// <summary>
        /// Gets or Sets Subject
        /// </summary>
        [DataMember(Name = "subject", IsRequired = true, EmitDefaultValue = false)]
        public string Subject { get; set; }

        /// <summary>
        /// Gets or Sets Type
        /// </summary>
        [DataMember(Name = "type", IsRequired = true, EmitDefaultValue = false)]
        public ClientCourierMessageType Type { get; set; }

        /// <summary>
        /// UpdatedAt is a helper struct field for gobuffalo.pop.
        /// </summary>
        /// <value>UpdatedAt is a helper struct field for gobuffalo.pop.</value>
        [DataMember(Name = "updated_at", IsRequired = true, EmitDefaultValue = false)]
        public DateTime UpdatedAt { get; set; }

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
            sb.Append("class ClientMessage {\n");
            sb.Append("  Body: ").Append(Body).Append("\n");
            sb.Append("  CreatedAt: ").Append(CreatedAt).Append("\n");
            sb.Append("  Dispatches: ").Append(Dispatches).Append("\n");
            sb.Append("  Id: ").Append(Id).Append("\n");
            sb.Append("  Recipient: ").Append(Recipient).Append("\n");
            sb.Append("  SendCount: ").Append(SendCount).Append("\n");
            sb.Append("  Status: ").Append(Status).Append("\n");
            sb.Append("  Subject: ").Append(Subject).Append("\n");
            sb.Append("  TemplateType: ").Append(TemplateType).Append("\n");
            sb.Append("  Type: ").Append(Type).Append("\n");
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
        /// Returns true if objects are equal
        /// </summary>
        /// <param name="input">Object to be compared</param>
        /// <returns>Boolean</returns>
        public override bool Equals(object input)
        {
            return this.Equals(input as ClientMessage);
        }

        /// <summary>
        /// Returns true if ClientMessage instances are equal
        /// </summary>
        /// <param name="input">Instance of ClientMessage to be compared</param>
        /// <returns>Boolean</returns>
        public bool Equals(ClientMessage input)
        {
            if (input == null)
            {
                return false;
            }
            return 
                (
                    this.Body == input.Body ||
                    (this.Body != null &&
                    this.Body.Equals(input.Body))
                ) && 
                (
                    this.CreatedAt == input.CreatedAt ||
                    (this.CreatedAt != null &&
                    this.CreatedAt.Equals(input.CreatedAt))
                ) && 
                (
                    this.Dispatches == input.Dispatches ||
                    this.Dispatches != null &&
                    input.Dispatches != null &&
                    this.Dispatches.SequenceEqual(input.Dispatches)
                ) && 
                (
                    this.Id == input.Id ||
                    (this.Id != null &&
                    this.Id.Equals(input.Id))
                ) && 
                (
                    this.Recipient == input.Recipient ||
                    (this.Recipient != null &&
                    this.Recipient.Equals(input.Recipient))
                ) && 
                (
                    this.SendCount == input.SendCount ||
                    this.SendCount.Equals(input.SendCount)
                ) && 
                (
                    this.Status == input.Status ||
                    (this.Status != null &&
                    this.Status.Equals(input.Status))
                ) && 
                (
                    this.Subject == input.Subject ||
                    (this.Subject != null &&
                    this.Subject.Equals(input.Subject))
                ) && 
                (
                    this.TemplateType == input.TemplateType ||
                    this.TemplateType.Equals(input.TemplateType)
                ) && 
                (
                    this.Type == input.Type ||
                    (this.Type != null &&
                    this.Type.Equals(input.Type))
                ) && 
                (
                    this.UpdatedAt == input.UpdatedAt ||
                    (this.UpdatedAt != null &&
                    this.UpdatedAt.Equals(input.UpdatedAt))
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
                if (this.Body != null)
                {
                    hashCode = (hashCode * 59) + this.Body.GetHashCode();
                }
                if (this.CreatedAt != null)
                {
                    hashCode = (hashCode * 59) + this.CreatedAt.GetHashCode();
                }
                if (this.Dispatches != null)
                {
                    hashCode = (hashCode * 59) + this.Dispatches.GetHashCode();
                }
                if (this.Id != null)
                {
                    hashCode = (hashCode * 59) + this.Id.GetHashCode();
                }
                if (this.Recipient != null)
                {
                    hashCode = (hashCode * 59) + this.Recipient.GetHashCode();
                }
                hashCode = (hashCode * 59) + this.SendCount.GetHashCode();
                if (this.Status != null)
                {
                    hashCode = (hashCode * 59) + this.Status.GetHashCode();
                }
                if (this.Subject != null)
                {
                    hashCode = (hashCode * 59) + this.Subject.GetHashCode();
                }
                hashCode = (hashCode * 59) + this.TemplateType.GetHashCode();
                if (this.Type != null)
                {
                    hashCode = (hashCode * 59) + this.Type.GetHashCode();
                }
                if (this.UpdatedAt != null)
                {
                    hashCode = (hashCode * 59) + this.UpdatedAt.GetHashCode();
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

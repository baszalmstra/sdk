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
using JsonSubTypes;
using System.ComponentModel.DataAnnotations;
using OpenAPIDateConverter = Ory.Client.Client.OpenAPIDateConverter;
using System.Reflection;

namespace Ory.Client.Model
{
    /// <summary>
    /// Update Verification Flow Request Body
    /// </summary>
    [JsonConverter(typeof(ClientUpdateVerificationFlowBodyJsonConverter))]
    [DataContract(Name = "updateVerificationFlowBody")]
    public partial class ClientUpdateVerificationFlowBody : AbstractOpenAPISchema, IValidatableObject
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="ClientUpdateVerificationFlowBody" /> class
        /// with the <see cref="ClientUpdateVerificationFlowWithLinkMethod" /> class
        /// </summary>
        /// <param name="actualInstance">An instance of ClientUpdateVerificationFlowWithLinkMethod.</param>
        public ClientUpdateVerificationFlowBody(ClientUpdateVerificationFlowWithLinkMethod actualInstance)
        {
            this.IsNullable = false;
            this.SchemaType= "oneOf";
            this.ActualInstance = actualInstance ?? throw new ArgumentException("Invalid instance found. Must not be null.");
        }

        /// <summary>
        /// Initializes a new instance of the <see cref="ClientUpdateVerificationFlowBody" /> class
        /// with the <see cref="ClientUpdateVerificationFlowWithCodeMethod" /> class
        /// </summary>
        /// <param name="actualInstance">An instance of ClientUpdateVerificationFlowWithCodeMethod.</param>
        public ClientUpdateVerificationFlowBody(ClientUpdateVerificationFlowWithCodeMethod actualInstance)
        {
            this.IsNullable = false;
            this.SchemaType= "oneOf";
            this.ActualInstance = actualInstance ?? throw new ArgumentException("Invalid instance found. Must not be null.");
        }


        private Object _actualInstance;

        /// <summary>
        /// Gets or Sets ActualInstance
        /// </summary>
        public override Object ActualInstance
        {
            get
            {
                return _actualInstance;
            }
            set
            {
                if (value.GetType() == typeof(ClientUpdateVerificationFlowWithCodeMethod) || value is ClientUpdateVerificationFlowWithCodeMethod)
                {
                    this._actualInstance = value;
                }
                else if (value.GetType() == typeof(ClientUpdateVerificationFlowWithLinkMethod) || value is ClientUpdateVerificationFlowWithLinkMethod)
                {
                    this._actualInstance = value;
                }
                else
                {
                    throw new ArgumentException("Invalid instance found. Must be the following types: ClientUpdateVerificationFlowWithCodeMethod, ClientUpdateVerificationFlowWithLinkMethod");
                }
            }
        }

        /// <summary>
        /// Get the actual instance of `ClientUpdateVerificationFlowWithLinkMethod`. If the actual instance is not `ClientUpdateVerificationFlowWithLinkMethod`,
        /// the InvalidClassException will be thrown
        /// </summary>
        /// <returns>An instance of ClientUpdateVerificationFlowWithLinkMethod</returns>
        public ClientUpdateVerificationFlowWithLinkMethod GetClientUpdateVerificationFlowWithLinkMethod()
        {
            return (ClientUpdateVerificationFlowWithLinkMethod)this.ActualInstance;
        }

        /// <summary>
        /// Get the actual instance of `ClientUpdateVerificationFlowWithCodeMethod`. If the actual instance is not `ClientUpdateVerificationFlowWithCodeMethod`,
        /// the InvalidClassException will be thrown
        /// </summary>
        /// <returns>An instance of ClientUpdateVerificationFlowWithCodeMethod</returns>
        public ClientUpdateVerificationFlowWithCodeMethod GetClientUpdateVerificationFlowWithCodeMethod()
        {
            return (ClientUpdateVerificationFlowWithCodeMethod)this.ActualInstance;
        }

        /// <summary>
        /// Returns the string presentation of the object
        /// </summary>
        /// <returns>String presentation of the object</returns>
        public override string ToString()
        {
            var sb = new StringBuilder();
            sb.Append("class ClientUpdateVerificationFlowBody {\n");
            sb.Append("  ActualInstance: ").Append(this.ActualInstance).Append("\n");
            sb.Append("}\n");
            return sb.ToString();
        }

        /// <summary>
        /// Returns the JSON string presentation of the object
        /// </summary>
        /// <returns>JSON string presentation of the object</returns>
        public override string ToJson()
        {
            return JsonConvert.SerializeObject(this.ActualInstance, ClientUpdateVerificationFlowBody.SerializerSettings);
        }

        /// <summary>
        /// Converts the JSON string into an instance of ClientUpdateVerificationFlowBody
        /// </summary>
        /// <param name="jsonString">JSON string</param>
        /// <returns>An instance of ClientUpdateVerificationFlowBody</returns>
        public static ClientUpdateVerificationFlowBody FromJson(string jsonString)
        {
            ClientUpdateVerificationFlowBody newClientUpdateVerificationFlowBody = null;

            if (string.IsNullOrEmpty(jsonString))
            {
                return newClientUpdateVerificationFlowBody;
            }
            int match = 0;
            List<string> matchedTypes = new List<string>();

            try
            {
                // if it does not contains "AdditionalProperties", use SerializerSettings to deserialize
                if (typeof(ClientUpdateVerificationFlowWithCodeMethod).GetProperty("AdditionalProperties") == null)
                {
                    newClientUpdateVerificationFlowBody = new ClientUpdateVerificationFlowBody(JsonConvert.DeserializeObject<ClientUpdateVerificationFlowWithCodeMethod>(jsonString, ClientUpdateVerificationFlowBody.SerializerSettings));
                }
                else
                {
                    newClientUpdateVerificationFlowBody = new ClientUpdateVerificationFlowBody(JsonConvert.DeserializeObject<ClientUpdateVerificationFlowWithCodeMethod>(jsonString, ClientUpdateVerificationFlowBody.AdditionalPropertiesSerializerSettings));
                }
                matchedTypes.Add("ClientUpdateVerificationFlowWithCodeMethod");
                match++;
            }
            catch (Exception exception)
            {
                // deserialization failed, try the next one
                System.Diagnostics.Debug.WriteLine(string.Format("Failed to deserialize `{0}` into ClientUpdateVerificationFlowWithCodeMethod: {1}", jsonString, exception.ToString()));
            }

            try
            {
                // if it does not contains "AdditionalProperties", use SerializerSettings to deserialize
                if (typeof(ClientUpdateVerificationFlowWithLinkMethod).GetProperty("AdditionalProperties") == null)
                {
                    newClientUpdateVerificationFlowBody = new ClientUpdateVerificationFlowBody(JsonConvert.DeserializeObject<ClientUpdateVerificationFlowWithLinkMethod>(jsonString, ClientUpdateVerificationFlowBody.SerializerSettings));
                }
                else
                {
                    newClientUpdateVerificationFlowBody = new ClientUpdateVerificationFlowBody(JsonConvert.DeserializeObject<ClientUpdateVerificationFlowWithLinkMethod>(jsonString, ClientUpdateVerificationFlowBody.AdditionalPropertiesSerializerSettings));
                }
                matchedTypes.Add("ClientUpdateVerificationFlowWithLinkMethod");
                match++;
            }
            catch (Exception exception)
            {
                // deserialization failed, try the next one
                System.Diagnostics.Debug.WriteLine(string.Format("Failed to deserialize `{0}` into ClientUpdateVerificationFlowWithLinkMethod: {1}", jsonString, exception.ToString()));
            }

            if (match == 0)
            {
                throw new InvalidDataException("The JSON string `" + jsonString + "` cannot be deserialized into any schema defined.");
            }
            else if (match > 1)
            {
                throw new InvalidDataException("The JSON string `" + jsonString + "` incorrectly matches more than one schema (should be exactly one match): " + String.Join(",", matchedTypes));
            }

            // deserialization is considered successful at this point if no exception has been thrown.
            return newClientUpdateVerificationFlowBody;
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

    /// <summary>
    /// Custom JSON converter for ClientUpdateVerificationFlowBody
    /// </summary>
    public class ClientUpdateVerificationFlowBodyJsonConverter : JsonConverter
    {
        /// <summary>
        /// To write the JSON string
        /// </summary>
        /// <param name="writer">JSON writer</param>
        /// <param name="value">Object to be converted into a JSON string</param>
        /// <param name="serializer">JSON Serializer</param>
        public override void WriteJson(JsonWriter writer, object value, JsonSerializer serializer)
        {
            writer.WriteRawValue((string)(typeof(ClientUpdateVerificationFlowBody).GetMethod("ToJson").Invoke(value, null)));
        }

        /// <summary>
        /// To convert a JSON string into an object
        /// </summary>
        /// <param name="reader">JSON reader</param>
        /// <param name="objectType">Object type</param>
        /// <param name="existingValue">Existing value</param>
        /// <param name="serializer">JSON Serializer</param>
        /// <returns>The object converted from the JSON string</returns>
        public override object ReadJson(JsonReader reader, Type objectType, object existingValue, JsonSerializer serializer)
        {
            switch(reader.TokenType) 
            {
                case JsonToken.StartObject:
                    return ClientUpdateVerificationFlowBody.FromJson(JObject.Load(reader).ToString(Formatting.None));
                case JsonToken.StartArray:
                    return ClientUpdateVerificationFlowBody.FromJson(JArray.Load(reader).ToString(Formatting.None));
                default:
                    return null;
            }
        }

        /// <summary>
        /// Check if the object can be converted
        /// </summary>
        /// <param name="objectType">Object type</param>
        /// <returns>True if the object can be converted</returns>
        public override bool CanConvert(Type objectType)
        {
            return false;
        }
    }

}

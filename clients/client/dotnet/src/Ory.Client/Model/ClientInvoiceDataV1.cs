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
    /// ClientInvoiceDataV1
    /// </summary>
    [DataContract(Name = "InvoiceDataV1")]
    public partial class ClientInvoiceDataV1 : IValidatableObject
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="ClientInvoiceDataV1" /> class.
        /// </summary>
        [JsonConstructorAttribute]
        protected ClientInvoiceDataV1()
        {
            this.AdditionalProperties = new Dictionary<string, object>();
        }
        /// <summary>
        /// Initializes a new instance of the <see cref="ClientInvoiceDataV1" /> class.
        /// </summary>
        /// <param name="billingPeriod">billingPeriod (required).</param>
        /// <param name="stripeInvoiceItem">stripeInvoiceItem.</param>
        /// <param name="stripeInvoiceStatus">The status of the invoice, one of &#x60;draft&#x60;, &#x60;open&#x60;, &#x60;paid&#x60;, &#x60;uncollectible&#x60;, or &#x60;void&#x60;. [Learn more](https://stripe.com/docs/billing/invoices/workflow#workflow-overview).</param>
        /// <param name="tax">tax.</param>
        /// <param name="totalInCent">totalInCent (required).</param>
        public ClientInvoiceDataV1(ClientTimeInterval billingPeriod = default(ClientTimeInterval), string stripeInvoiceItem = default(string), string stripeInvoiceStatus = default(string), ClientTaxLineItem tax = default(ClientTaxLineItem), long totalInCent = default(long))
        {
            // to ensure "billingPeriod" is required (not null)
            if (billingPeriod == null)
            {
                throw new ArgumentNullException("billingPeriod is a required property for ClientInvoiceDataV1 and cannot be null");
            }
            this.BillingPeriod = billingPeriod;
            this.TotalInCent = totalInCent;
            this.StripeInvoiceItem = stripeInvoiceItem;
            this.StripeInvoiceStatus = stripeInvoiceStatus;
            this.Tax = tax;
            this.AdditionalProperties = new Dictionary<string, object>();
        }

        /// <summary>
        /// Gets or Sets BillingPeriod
        /// </summary>
        [DataMember(Name = "billing_period", IsRequired = true, EmitDefaultValue = true)]
        public ClientTimeInterval BillingPeriod { get; set; }

        /// <summary>
        /// The currency of the invoice.
        /// </summary>
        /// <value>The currency of the invoice.</value>
        [DataMember(Name = "currency", IsRequired = true, EmitDefaultValue = true)]
        public string Currency { get; private set; }

        /// <summary>
        /// Returns false as Currency should not be serialized given that it's read-only.
        /// </summary>
        /// <returns>false (boolean)</returns>
        public bool ShouldSerializeCurrency()
        {
            return false;
        }
        /// <summary>
        /// The items that are part of this invoice.
        /// </summary>
        /// <value>The items that are part of this invoice.</value>
        [DataMember(Name = "items", IsRequired = true, EmitDefaultValue = true)]
        public List<ClientLineItemV1> Items { get; private set; }

        /// <summary>
        /// Returns false as Items should not be serialized given that it's read-only.
        /// </summary>
        /// <returns>false (boolean)</returns>
        public bool ShouldSerializeItems()
        {
            return false;
        }
        /// <summary>
        /// The plan that this invoice is based on.
        /// </summary>
        /// <value>The plan that this invoice is based on.</value>
        [DataMember(Name = "plan", EmitDefaultValue = false)]
        public string Plan { get; private set; }

        /// <summary>
        /// Returns false as Plan should not be serialized given that it's read-only.
        /// </summary>
        /// <returns>false (boolean)</returns>
        public bool ShouldSerializePlan()
        {
            return false;
        }
        /// <summary>
        /// Gets or Sets StripeInvoiceItem
        /// </summary>
        [DataMember(Name = "stripe_invoice_item", EmitDefaultValue = false)]
        public string StripeInvoiceItem { get; set; }

        /// <summary>
        /// The status of the invoice, one of &#x60;draft&#x60;, &#x60;open&#x60;, &#x60;paid&#x60;, &#x60;uncollectible&#x60;, or &#x60;void&#x60;. [Learn more](https://stripe.com/docs/billing/invoices/workflow#workflow-overview)
        /// </summary>
        /// <value>The status of the invoice, one of &#x60;draft&#x60;, &#x60;open&#x60;, &#x60;paid&#x60;, &#x60;uncollectible&#x60;, or &#x60;void&#x60;. [Learn more](https://stripe.com/docs/billing/invoices/workflow#workflow-overview)</value>
        [DataMember(Name = "stripe_invoice_status", EmitDefaultValue = false)]
        public string StripeInvoiceStatus { get; set; }

        /// <summary>
        /// An optional link to the invoice on Stripe.
        /// </summary>
        /// <value>An optional link to the invoice on Stripe.</value>
        [DataMember(Name = "stripe_link", EmitDefaultValue = false)]
        public string StripeLink { get; private set; }

        /// <summary>
        /// Returns false as StripeLink should not be serialized given that it's read-only.
        /// </summary>
        /// <returns>false (boolean)</returns>
        public bool ShouldSerializeStripeLink()
        {
            return false;
        }
        /// <summary>
        /// The subtitle of the invoice.
        /// </summary>
        /// <value>The subtitle of the invoice.</value>
        [DataMember(Name = "subtitle", EmitDefaultValue = false)]
        public string Subtitle { get; private set; }

        /// <summary>
        /// Returns false as Subtitle should not be serialized given that it's read-only.
        /// </summary>
        /// <returns>false (boolean)</returns>
        public bool ShouldSerializeSubtitle()
        {
            return false;
        }
        /// <summary>
        /// Gets or Sets Tax
        /// </summary>
        [DataMember(Name = "tax", EmitDefaultValue = false)]
        public ClientTaxLineItem Tax { get; set; }

        /// <summary>
        /// The title of the invoice.
        /// </summary>
        /// <value>The title of the invoice.</value>
        [DataMember(Name = "title", IsRequired = true, EmitDefaultValue = true)]
        public string Title { get; private set; }

        /// <summary>
        /// Returns false as Title should not be serialized given that it's read-only.
        /// </summary>
        /// <returns>false (boolean)</returns>
        public bool ShouldSerializeTitle()
        {
            return false;
        }
        /// <summary>
        /// Gets or Sets TotalInCent
        /// </summary>
        [DataMember(Name = "total_in_cent", IsRequired = true, EmitDefaultValue = true)]
        public long TotalInCent { get; set; }

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
            sb.Append("class ClientInvoiceDataV1 {\n");
            sb.Append("  BillingPeriod: ").Append(BillingPeriod).Append("\n");
            sb.Append("  Currency: ").Append(Currency).Append("\n");
            sb.Append("  Items: ").Append(Items).Append("\n");
            sb.Append("  Plan: ").Append(Plan).Append("\n");
            sb.Append("  StripeInvoiceItem: ").Append(StripeInvoiceItem).Append("\n");
            sb.Append("  StripeInvoiceStatus: ").Append(StripeInvoiceStatus).Append("\n");
            sb.Append("  StripeLink: ").Append(StripeLink).Append("\n");
            sb.Append("  Subtitle: ").Append(Subtitle).Append("\n");
            sb.Append("  Tax: ").Append(Tax).Append("\n");
            sb.Append("  Title: ").Append(Title).Append("\n");
            sb.Append("  TotalInCent: ").Append(TotalInCent).Append("\n");
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

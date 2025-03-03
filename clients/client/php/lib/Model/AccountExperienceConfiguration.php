<?php
/**
 * AccountExperienceConfiguration
 *
 * PHP version 7.4
 *
 * @category Class
 * @package  Ory\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 */

/**
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers.
 *
 * The version of the OpenAPI document: v1.15.7
 * Contact: support@ory.sh
 * Generated by: https://openapi-generator.tech
 * Generator version: 7.7.0
 */

/**
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

namespace Ory\Client\Model;

use \ArrayAccess;
use \Ory\Client\ObjectSerializer;

/**
 * AccountExperienceConfiguration Class Doc Comment
 *
 * @category Class
 * @package  Ory\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 * @implements \ArrayAccess<string, mixed>
 */
class AccountExperienceConfiguration implements ModelInterface, ArrayAccess, \JsonSerializable
{
    public const DISCRIMINATOR = null;

    /**
      * The original name of the model.
      *
      * @var string
      */
    protected static $openAPIModelName = 'accountExperienceConfiguration';

    /**
      * Array of property to type mappings. Used for (de)serialization
      *
      * @var string[]
      */
    protected static $openAPITypes = [
        'accountExperienceThemeStylesheet' => 'string',
        'faviconType' => 'string',
        'faviconUrl' => 'string',
        'kratosSelfserviceDefaultBrowserReturnUrl' => 'string',
        'kratosSelfserviceFlowsRecoveryEnabled' => 'bool',
        'kratosSelfserviceFlowsRegistrationEnabled' => 'bool',
        'kratosSelfserviceFlowsVerificationEnabled' => 'bool',
        'logoUrl' => 'string',
        'name' => 'string',
        'organizationMap' => 'array<string,string>'
    ];

    /**
      * Array of property to format mappings. Used for (de)serialization
      *
      * @var string[]
      * @phpstan-var array<string, string|null>
      * @psalm-var array<string, string|null>
      */
    protected static $openAPIFormats = [
        'accountExperienceThemeStylesheet' => null,
        'faviconType' => null,
        'faviconUrl' => null,
        'kratosSelfserviceDefaultBrowserReturnUrl' => null,
        'kratosSelfserviceFlowsRecoveryEnabled' => null,
        'kratosSelfserviceFlowsRegistrationEnabled' => null,
        'kratosSelfserviceFlowsVerificationEnabled' => null,
        'logoUrl' => null,
        'name' => null,
        'organizationMap' => null
    ];

    /**
      * Array of nullable properties. Used for (de)serialization
      *
      * @var boolean[]
      */
    protected static array $openAPINullables = [
        'accountExperienceThemeStylesheet' => false,
        'faviconType' => false,
        'faviconUrl' => false,
        'kratosSelfserviceDefaultBrowserReturnUrl' => false,
        'kratosSelfserviceFlowsRecoveryEnabled' => false,
        'kratosSelfserviceFlowsRegistrationEnabled' => false,
        'kratosSelfserviceFlowsVerificationEnabled' => false,
        'logoUrl' => false,
        'name' => false,
        'organizationMap' => false
    ];

    /**
      * If a nullable field gets set to null, insert it here
      *
      * @var boolean[]
      */
    protected array $openAPINullablesSetToNull = [];

    /**
     * Array of property to type mappings. Used for (de)serialization
     *
     * @return array
     */
    public static function openAPITypes()
    {
        return self::$openAPITypes;
    }

    /**
     * Array of property to format mappings. Used for (de)serialization
     *
     * @return array
     */
    public static function openAPIFormats()
    {
        return self::$openAPIFormats;
    }

    /**
     * Array of nullable properties
     *
     * @return array
     */
    protected static function openAPINullables(): array
    {
        return self::$openAPINullables;
    }

    /**
     * Array of nullable field names deliberately set to null
     *
     * @return boolean[]
     */
    private function getOpenAPINullablesSetToNull(): array
    {
        return $this->openAPINullablesSetToNull;
    }

    /**
     * Setter - Array of nullable field names deliberately set to null
     *
     * @param boolean[] $openAPINullablesSetToNull
     */
    private function setOpenAPINullablesSetToNull(array $openAPINullablesSetToNull): void
    {
        $this->openAPINullablesSetToNull = $openAPINullablesSetToNull;
    }

    /**
     * Checks if a property is nullable
     *
     * @param string $property
     * @return bool
     */
    public static function isNullable(string $property): bool
    {
        return self::openAPINullables()[$property] ?? false;
    }

    /**
     * Checks if a nullable property is set to null.
     *
     * @param string $property
     * @return bool
     */
    public function isNullableSetToNull(string $property): bool
    {
        return in_array($property, $this->getOpenAPINullablesSetToNull(), true);
    }

    /**
     * Array of attributes where the key is the local name,
     * and the value is the original name
     *
     * @var string[]
     */
    protected static $attributeMap = [
        'accountExperienceThemeStylesheet' => 'account_experience_theme_stylesheet',
        'faviconType' => 'favicon_type',
        'faviconUrl' => 'favicon_url',
        'kratosSelfserviceDefaultBrowserReturnUrl' => 'kratos_selfservice_default_browser_return_url',
        'kratosSelfserviceFlowsRecoveryEnabled' => 'kratos_selfservice_flows_recovery_enabled',
        'kratosSelfserviceFlowsRegistrationEnabled' => 'kratos_selfservice_flows_registration_enabled',
        'kratosSelfserviceFlowsVerificationEnabled' => 'kratos_selfservice_flows_verification_enabled',
        'logoUrl' => 'logo_url',
        'name' => 'name',
        'organizationMap' => 'organization_map'
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @var string[]
     */
    protected static $setters = [
        'accountExperienceThemeStylesheet' => 'setAccountExperienceThemeStylesheet',
        'faviconType' => 'setFaviconType',
        'faviconUrl' => 'setFaviconUrl',
        'kratosSelfserviceDefaultBrowserReturnUrl' => 'setKratosSelfserviceDefaultBrowserReturnUrl',
        'kratosSelfserviceFlowsRecoveryEnabled' => 'setKratosSelfserviceFlowsRecoveryEnabled',
        'kratosSelfserviceFlowsRegistrationEnabled' => 'setKratosSelfserviceFlowsRegistrationEnabled',
        'kratosSelfserviceFlowsVerificationEnabled' => 'setKratosSelfserviceFlowsVerificationEnabled',
        'logoUrl' => 'setLogoUrl',
        'name' => 'setName',
        'organizationMap' => 'setOrganizationMap'
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @var string[]
     */
    protected static $getters = [
        'accountExperienceThemeStylesheet' => 'getAccountExperienceThemeStylesheet',
        'faviconType' => 'getFaviconType',
        'faviconUrl' => 'getFaviconUrl',
        'kratosSelfserviceDefaultBrowserReturnUrl' => 'getKratosSelfserviceDefaultBrowserReturnUrl',
        'kratosSelfserviceFlowsRecoveryEnabled' => 'getKratosSelfserviceFlowsRecoveryEnabled',
        'kratosSelfserviceFlowsRegistrationEnabled' => 'getKratosSelfserviceFlowsRegistrationEnabled',
        'kratosSelfserviceFlowsVerificationEnabled' => 'getKratosSelfserviceFlowsVerificationEnabled',
        'logoUrl' => 'getLogoUrl',
        'name' => 'getName',
        'organizationMap' => 'getOrganizationMap'
    ];

    /**
     * Array of attributes where the key is the local name,
     * and the value is the original name
     *
     * @return array
     */
    public static function attributeMap()
    {
        return self::$attributeMap;
    }

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @return array
     */
    public static function setters()
    {
        return self::$setters;
    }

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @return array
     */
    public static function getters()
    {
        return self::$getters;
    }

    /**
     * The original name of the model.
     *
     * @return string
     */
    public function getModelName()
    {
        return self::$openAPIModelName;
    }


    /**
     * Associative array for storing property values
     *
     * @var mixed[]
     */
    protected $container = [];

    /**
     * Constructor
     *
     * @param mixed[] $data Associated array of property values
     *                      initializing the model
     */
    public function __construct(array $data = null)
    {
        $this->setIfExists('accountExperienceThemeStylesheet', $data ?? [], null);
        $this->setIfExists('faviconType', $data ?? [], null);
        $this->setIfExists('faviconUrl', $data ?? [], null);
        $this->setIfExists('kratosSelfserviceDefaultBrowserReturnUrl', $data ?? [], null);
        $this->setIfExists('kratosSelfserviceFlowsRecoveryEnabled', $data ?? [], null);
        $this->setIfExists('kratosSelfserviceFlowsRegistrationEnabled', $data ?? [], null);
        $this->setIfExists('kratosSelfserviceFlowsVerificationEnabled', $data ?? [], null);
        $this->setIfExists('logoUrl', $data ?? [], null);
        $this->setIfExists('name', $data ?? [], null);
        $this->setIfExists('organizationMap', $data ?? [], null);
    }

    /**
    * Sets $this->container[$variableName] to the given data or to the given default Value; if $variableName
    * is nullable and its value is set to null in the $fields array, then mark it as "set to null" in the
    * $this->openAPINullablesSetToNull array
    *
    * @param string $variableName
    * @param array  $fields
    * @param mixed  $defaultValue
    */
    private function setIfExists(string $variableName, array $fields, $defaultValue): void
    {
        if (self::isNullable($variableName) && array_key_exists($variableName, $fields) && is_null($fields[$variableName])) {
            $this->openAPINullablesSetToNull[] = $variableName;
        }

        $this->container[$variableName] = $fields[$variableName] ?? $defaultValue;
    }

    /**
     * Show all the invalid properties with reasons.
     *
     * @return array invalid properties with reasons
     */
    public function listInvalidProperties()
    {
        $invalidProperties = [];

        return $invalidProperties;
    }

    /**
     * Validate all the properties in the model
     * return true if all passed
     *
     * @return bool True if all properties are valid
     */
    public function valid()
    {
        return count($this->listInvalidProperties()) === 0;
    }


    /**
     * Gets accountExperienceThemeStylesheet
     *
     * @return string|null
     */
    public function getAccountExperienceThemeStylesheet()
    {
        return $this->container['accountExperienceThemeStylesheet'];
    }

    /**
     * Sets accountExperienceThemeStylesheet
     *
     * @param string|null $accountExperienceThemeStylesheet accountExperienceThemeStylesheet
     *
     * @return self
     */
    public function setAccountExperienceThemeStylesheet($accountExperienceThemeStylesheet)
    {
        if (is_null($accountExperienceThemeStylesheet)) {
            throw new \InvalidArgumentException('non-nullable accountExperienceThemeStylesheet cannot be null');
        }
        $this->container['accountExperienceThemeStylesheet'] = $accountExperienceThemeStylesheet;

        return $this;
    }

    /**
     * Gets faviconType
     *
     * @return string|null
     */
    public function getFaviconType()
    {
        return $this->container['faviconType'];
    }

    /**
     * Sets faviconType
     *
     * @param string|null $faviconType faviconType
     *
     * @return self
     */
    public function setFaviconType($faviconType)
    {
        if (is_null($faviconType)) {
            throw new \InvalidArgumentException('non-nullable faviconType cannot be null');
        }
        $this->container['faviconType'] = $faviconType;

        return $this;
    }

    /**
     * Gets faviconUrl
     *
     * @return string|null
     */
    public function getFaviconUrl()
    {
        return $this->container['faviconUrl'];
    }

    /**
     * Sets faviconUrl
     *
     * @param string|null $faviconUrl faviconUrl
     *
     * @return self
     */
    public function setFaviconUrl($faviconUrl)
    {
        if (is_null($faviconUrl)) {
            throw new \InvalidArgumentException('non-nullable faviconUrl cannot be null');
        }
        $this->container['faviconUrl'] = $faviconUrl;

        return $this;
    }

    /**
     * Gets kratosSelfserviceDefaultBrowserReturnUrl
     *
     * @return string|null
     */
    public function getKratosSelfserviceDefaultBrowserReturnUrl()
    {
        return $this->container['kratosSelfserviceDefaultBrowserReturnUrl'];
    }

    /**
     * Sets kratosSelfserviceDefaultBrowserReturnUrl
     *
     * @param string|null $kratosSelfserviceDefaultBrowserReturnUrl kratosSelfserviceDefaultBrowserReturnUrl
     *
     * @return self
     */
    public function setKratosSelfserviceDefaultBrowserReturnUrl($kratosSelfserviceDefaultBrowserReturnUrl)
    {
        if (is_null($kratosSelfserviceDefaultBrowserReturnUrl)) {
            throw new \InvalidArgumentException('non-nullable kratosSelfserviceDefaultBrowserReturnUrl cannot be null');
        }
        $this->container['kratosSelfserviceDefaultBrowserReturnUrl'] = $kratosSelfserviceDefaultBrowserReturnUrl;

        return $this;
    }

    /**
     * Gets kratosSelfserviceFlowsRecoveryEnabled
     *
     * @return bool|null
     */
    public function getKratosSelfserviceFlowsRecoveryEnabled()
    {
        return $this->container['kratosSelfserviceFlowsRecoveryEnabled'];
    }

    /**
     * Sets kratosSelfserviceFlowsRecoveryEnabled
     *
     * @param bool|null $kratosSelfserviceFlowsRecoveryEnabled kratosSelfserviceFlowsRecoveryEnabled
     *
     * @return self
     */
    public function setKratosSelfserviceFlowsRecoveryEnabled($kratosSelfserviceFlowsRecoveryEnabled)
    {
        if (is_null($kratosSelfserviceFlowsRecoveryEnabled)) {
            throw new \InvalidArgumentException('non-nullable kratosSelfserviceFlowsRecoveryEnabled cannot be null');
        }
        $this->container['kratosSelfserviceFlowsRecoveryEnabled'] = $kratosSelfserviceFlowsRecoveryEnabled;

        return $this;
    }

    /**
     * Gets kratosSelfserviceFlowsRegistrationEnabled
     *
     * @return bool|null
     */
    public function getKratosSelfserviceFlowsRegistrationEnabled()
    {
        return $this->container['kratosSelfserviceFlowsRegistrationEnabled'];
    }

    /**
     * Sets kratosSelfserviceFlowsRegistrationEnabled
     *
     * @param bool|null $kratosSelfserviceFlowsRegistrationEnabled kratosSelfserviceFlowsRegistrationEnabled
     *
     * @return self
     */
    public function setKratosSelfserviceFlowsRegistrationEnabled($kratosSelfserviceFlowsRegistrationEnabled)
    {
        if (is_null($kratosSelfserviceFlowsRegistrationEnabled)) {
            throw new \InvalidArgumentException('non-nullable kratosSelfserviceFlowsRegistrationEnabled cannot be null');
        }
        $this->container['kratosSelfserviceFlowsRegistrationEnabled'] = $kratosSelfserviceFlowsRegistrationEnabled;

        return $this;
    }

    /**
     * Gets kratosSelfserviceFlowsVerificationEnabled
     *
     * @return bool|null
     */
    public function getKratosSelfserviceFlowsVerificationEnabled()
    {
        return $this->container['kratosSelfserviceFlowsVerificationEnabled'];
    }

    /**
     * Sets kratosSelfserviceFlowsVerificationEnabled
     *
     * @param bool|null $kratosSelfserviceFlowsVerificationEnabled kratosSelfserviceFlowsVerificationEnabled
     *
     * @return self
     */
    public function setKratosSelfserviceFlowsVerificationEnabled($kratosSelfserviceFlowsVerificationEnabled)
    {
        if (is_null($kratosSelfserviceFlowsVerificationEnabled)) {
            throw new \InvalidArgumentException('non-nullable kratosSelfserviceFlowsVerificationEnabled cannot be null');
        }
        $this->container['kratosSelfserviceFlowsVerificationEnabled'] = $kratosSelfserviceFlowsVerificationEnabled;

        return $this;
    }

    /**
     * Gets logoUrl
     *
     * @return string|null
     */
    public function getLogoUrl()
    {
        return $this->container['logoUrl'];
    }

    /**
     * Sets logoUrl
     *
     * @param string|null $logoUrl logoUrl
     *
     * @return self
     */
    public function setLogoUrl($logoUrl)
    {
        if (is_null($logoUrl)) {
            throw new \InvalidArgumentException('non-nullable logoUrl cannot be null');
        }
        $this->container['logoUrl'] = $logoUrl;

        return $this;
    }

    /**
     * Gets name
     *
     * @return string|null
     */
    public function getName()
    {
        return $this->container['name'];
    }

    /**
     * Sets name
     *
     * @param string|null $name name
     *
     * @return self
     */
    public function setName($name)
    {
        if (is_null($name)) {
            throw new \InvalidArgumentException('non-nullable name cannot be null');
        }
        $this->container['name'] = $name;

        return $this;
    }

    /**
     * Gets organizationMap
     *
     * @return array<string,string>|null
     */
    public function getOrganizationMap()
    {
        return $this->container['organizationMap'];
    }

    /**
     * Sets organizationMap
     *
     * @param array<string,string>|null $organizationMap organizationMap
     *
     * @return self
     */
    public function setOrganizationMap($organizationMap)
    {
        if (is_null($organizationMap)) {
            throw new \InvalidArgumentException('non-nullable organizationMap cannot be null');
        }
        $this->container['organizationMap'] = $organizationMap;

        return $this;
    }
    /**
     * Returns true if offset exists. False otherwise.
     *
     * @param integer $offset Offset
     *
     * @return boolean
     */
    public function offsetExists($offset): bool
    {
        return isset($this->container[$offset]);
    }

    /**
     * Gets offset.
     *
     * @param integer $offset Offset
     *
     * @return mixed|null
     */
    #[\ReturnTypeWillChange]
    public function offsetGet($offset)
    {
        return $this->container[$offset] ?? null;
    }

    /**
     * Sets value based on offset.
     *
     * @param int|null $offset Offset
     * @param mixed    $value  Value to be set
     *
     * @return void
     */
    public function offsetSet($offset, $value): void
    {
        if (is_null($offset)) {
            $this->container[] = $value;
        } else {
            $this->container[$offset] = $value;
        }
    }

    /**
     * Unsets offset.
     *
     * @param integer $offset Offset
     *
     * @return void
     */
    public function offsetUnset($offset): void
    {
        unset($this->container[$offset]);
    }

    /**
     * Serializes the object to a value that can be serialized natively by json_encode().
     * @link https://www.php.net/manual/en/jsonserializable.jsonserialize.php
     *
     * @return mixed Returns data which can be serialized by json_encode(), which is a value
     * of any type other than a resource.
     */
    #[\ReturnTypeWillChange]
    public function jsonSerialize()
    {
       return ObjectSerializer::sanitizeForSerialization($this);
    }

    /**
     * Gets the string presentation of the object
     *
     * @return string
     */
    public function __toString()
    {
        return json_encode(
            ObjectSerializer::sanitizeForSerialization($this),
            JSON_PRETTY_PRINT
        );
    }

    /**
     * Gets a header-safe presentation of the object
     *
     * @return string
     */
    public function toHeaderValue()
    {
        return json_encode(ObjectSerializer::sanitizeForSerialization($this));
    }
}



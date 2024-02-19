<?php
/**
 * NormalizedProjectRevisionCourierChannel
 *
 * PHP version 7.3
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
 * The version of the OpenAPI document: v1.6.2
 * Contact: support@ory.sh
 * Generated by: https://openapi-generator.tech
 * OpenAPI Generator version: 5.4.0
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
 * NormalizedProjectRevisionCourierChannel Class Doc Comment
 *
 * @category Class
 * @package  Ory\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 * @implements \ArrayAccess<TKey, TValue>
 * @template TKey int|null
 * @template TValue mixed|null
 */
class NormalizedProjectRevisionCourierChannel implements ModelInterface, ArrayAccess, \JsonSerializable
{
    public const DISCRIMINATOR = null;

    /**
      * The original name of the model.
      *
      * @var string
      */
    protected static $openAPIModelName = 'NormalizedProjectRevisionCourierChannel';

    /**
      * Array of property to type mappings. Used for (de)serialization
      *
      * @var string[]
      */
    protected static $openAPITypes = [
        'channelId' => 'string',
        'createdAt' => '\DateTime',
        'requestConfigAuthConfigApiKeyIn' => 'string',
        'requestConfigAuthConfigApiKeyName' => 'string',
        'requestConfigAuthConfigApiKeyValue' => 'string',
        'requestConfigAuthConfigBasicAuthPassword' => 'string',
        'requestConfigAuthConfigBasicAuthUser' => 'string',
        'requestConfigAuthType' => 'string',
        'requestConfigBody' => 'string',
        'requestConfigHeaders' => 'object',
        'requestConfigMethod' => 'string',
        'requestConfigUrl' => 'string',
        'updatedAt' => '\DateTime'
    ];

    /**
      * Array of property to format mappings. Used for (de)serialization
      *
      * @var string[]
      * @phpstan-var array<string, string|null>
      * @psalm-var array<string, string|null>
      */
    protected static $openAPIFormats = [
        'channelId' => null,
        'createdAt' => 'date-time',
        'requestConfigAuthConfigApiKeyIn' => null,
        'requestConfigAuthConfigApiKeyName' => null,
        'requestConfigAuthConfigApiKeyValue' => null,
        'requestConfigAuthConfigBasicAuthPassword' => null,
        'requestConfigAuthConfigBasicAuthUser' => null,
        'requestConfigAuthType' => null,
        'requestConfigBody' => null,
        'requestConfigHeaders' => null,
        'requestConfigMethod' => null,
        'requestConfigUrl' => null,
        'updatedAt' => 'date-time'
    ];

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
     * Array of attributes where the key is the local name,
     * and the value is the original name
     *
     * @var string[]
     */
    protected static $attributeMap = [
        'channelId' => 'channel_id',
        'createdAt' => 'created_at',
        'requestConfigAuthConfigApiKeyIn' => 'request_config_auth_config_api_key_in',
        'requestConfigAuthConfigApiKeyName' => 'request_config_auth_config_api_key_name',
        'requestConfigAuthConfigApiKeyValue' => 'request_config_auth_config_api_key_value',
        'requestConfigAuthConfigBasicAuthPassword' => 'request_config_auth_config_basic_auth_password',
        'requestConfigAuthConfigBasicAuthUser' => 'request_config_auth_config_basic_auth_user',
        'requestConfigAuthType' => 'request_config_auth_type',
        'requestConfigBody' => 'request_config_body',
        'requestConfigHeaders' => 'request_config_headers',
        'requestConfigMethod' => 'request_config_method',
        'requestConfigUrl' => 'request_config_url',
        'updatedAt' => 'updated_at'
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @var string[]
     */
    protected static $setters = [
        'channelId' => 'setChannelId',
        'createdAt' => 'setCreatedAt',
        'requestConfigAuthConfigApiKeyIn' => 'setRequestConfigAuthConfigApiKeyIn',
        'requestConfigAuthConfigApiKeyName' => 'setRequestConfigAuthConfigApiKeyName',
        'requestConfigAuthConfigApiKeyValue' => 'setRequestConfigAuthConfigApiKeyValue',
        'requestConfigAuthConfigBasicAuthPassword' => 'setRequestConfigAuthConfigBasicAuthPassword',
        'requestConfigAuthConfigBasicAuthUser' => 'setRequestConfigAuthConfigBasicAuthUser',
        'requestConfigAuthType' => 'setRequestConfigAuthType',
        'requestConfigBody' => 'setRequestConfigBody',
        'requestConfigHeaders' => 'setRequestConfigHeaders',
        'requestConfigMethod' => 'setRequestConfigMethod',
        'requestConfigUrl' => 'setRequestConfigUrl',
        'updatedAt' => 'setUpdatedAt'
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @var string[]
     */
    protected static $getters = [
        'channelId' => 'getChannelId',
        'createdAt' => 'getCreatedAt',
        'requestConfigAuthConfigApiKeyIn' => 'getRequestConfigAuthConfigApiKeyIn',
        'requestConfigAuthConfigApiKeyName' => 'getRequestConfigAuthConfigApiKeyName',
        'requestConfigAuthConfigApiKeyValue' => 'getRequestConfigAuthConfigApiKeyValue',
        'requestConfigAuthConfigBasicAuthPassword' => 'getRequestConfigAuthConfigBasicAuthPassword',
        'requestConfigAuthConfigBasicAuthUser' => 'getRequestConfigAuthConfigBasicAuthUser',
        'requestConfigAuthType' => 'getRequestConfigAuthType',
        'requestConfigBody' => 'getRequestConfigBody',
        'requestConfigHeaders' => 'getRequestConfigHeaders',
        'requestConfigMethod' => 'getRequestConfigMethod',
        'requestConfigUrl' => 'getRequestConfigUrl',
        'updatedAt' => 'getUpdatedAt'
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

    const REQUEST_CONFIG_AUTH_TYPE_BASIC_AUTH = 'basic_auth';
    const REQUEST_CONFIG_AUTH_TYPE_API_KEY = 'api_key';

    /**
     * Gets allowable values of the enum
     *
     * @return string[]
     */
    public function getRequestConfigAuthTypeAllowableValues()
    {
        return [
            self::REQUEST_CONFIG_AUTH_TYPE_BASIC_AUTH,
            self::REQUEST_CONFIG_AUTH_TYPE_API_KEY,
        ];
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
        $this->container['channelId'] = $data['channelId'] ?? null;
        $this->container['createdAt'] = $data['createdAt'] ?? null;
        $this->container['requestConfigAuthConfigApiKeyIn'] = $data['requestConfigAuthConfigApiKeyIn'] ?? null;
        $this->container['requestConfigAuthConfigApiKeyName'] = $data['requestConfigAuthConfigApiKeyName'] ?? null;
        $this->container['requestConfigAuthConfigApiKeyValue'] = $data['requestConfigAuthConfigApiKeyValue'] ?? null;
        $this->container['requestConfigAuthConfigBasicAuthPassword'] = $data['requestConfigAuthConfigBasicAuthPassword'] ?? null;
        $this->container['requestConfigAuthConfigBasicAuthUser'] = $data['requestConfigAuthConfigBasicAuthUser'] ?? null;
        $this->container['requestConfigAuthType'] = $data['requestConfigAuthType'] ?? null;
        $this->container['requestConfigBody'] = $data['requestConfigBody'] ?? null;
        $this->container['requestConfigHeaders'] = $data['requestConfigHeaders'] ?? null;
        $this->container['requestConfigMethod'] = $data['requestConfigMethod'] ?? null;
        $this->container['requestConfigUrl'] = $data['requestConfigUrl'] ?? null;
        $this->container['updatedAt'] = $data['updatedAt'] ?? null;
    }

    /**
     * Show all the invalid properties with reasons.
     *
     * @return array invalid properties with reasons
     */
    public function listInvalidProperties()
    {
        $invalidProperties = [];

        if ($this->container['channelId'] === null) {
            $invalidProperties[] = "'channelId' can't be null";
        }
        $allowedValues = $this->getRequestConfigAuthTypeAllowableValues();
        if (!is_null($this->container['requestConfigAuthType']) && !in_array($this->container['requestConfigAuthType'], $allowedValues, true)) {
            $invalidProperties[] = sprintf(
                "invalid value '%s' for 'requestConfigAuthType', must be one of '%s'",
                $this->container['requestConfigAuthType'],
                implode("', '", $allowedValues)
            );
        }

        if ($this->container['requestConfigBody'] === null) {
            $invalidProperties[] = "'requestConfigBody' can't be null";
        }
        if ($this->container['requestConfigMethod'] === null) {
            $invalidProperties[] = "'requestConfigMethod' can't be null";
        }
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
     * Gets channelId
     *
     * @return string
     */
    public function getChannelId()
    {
        return $this->container['channelId'];
    }

    /**
     * Sets channelId
     *
     * @param string $channelId The Channel's public ID
     *
     * @return self
     */
    public function setChannelId($channelId)
    {
        $this->container['channelId'] = $channelId;

        return $this;
    }

    /**
     * Gets createdAt
     *
     * @return \DateTime|null
     */
    public function getCreatedAt()
    {
        return $this->container['createdAt'];
    }

    /**
     * Sets createdAt
     *
     * @param \DateTime|null $createdAt The creation date
     *
     * @return self
     */
    public function setCreatedAt($createdAt)
    {
        $this->container['createdAt'] = $createdAt;

        return $this;
    }

    /**
     * Gets requestConfigAuthConfigApiKeyIn
     *
     * @return string|null
     */
    public function getRequestConfigAuthConfigApiKeyIn()
    {
        return $this->container['requestConfigAuthConfigApiKeyIn'];
    }

    /**
     * Sets requestConfigAuthConfigApiKeyIn
     *
     * @param string|null $requestConfigAuthConfigApiKeyIn API key location  Can either be \"header\" or \"query\"
     *
     * @return self
     */
    public function setRequestConfigAuthConfigApiKeyIn($requestConfigAuthConfigApiKeyIn)
    {
        $this->container['requestConfigAuthConfigApiKeyIn'] = $requestConfigAuthConfigApiKeyIn;

        return $this;
    }

    /**
     * Gets requestConfigAuthConfigApiKeyName
     *
     * @return string|null
     */
    public function getRequestConfigAuthConfigApiKeyName()
    {
        return $this->container['requestConfigAuthConfigApiKeyName'];
    }

    /**
     * Sets requestConfigAuthConfigApiKeyName
     *
     * @param string|null $requestConfigAuthConfigApiKeyName API key name  Only used if the auth type is api_key
     *
     * @return self
     */
    public function setRequestConfigAuthConfigApiKeyName($requestConfigAuthConfigApiKeyName)
    {
        $this->container['requestConfigAuthConfigApiKeyName'] = $requestConfigAuthConfigApiKeyName;

        return $this;
    }

    /**
     * Gets requestConfigAuthConfigApiKeyValue
     *
     * @return string|null
     */
    public function getRequestConfigAuthConfigApiKeyValue()
    {
        return $this->container['requestConfigAuthConfigApiKeyValue'];
    }

    /**
     * Sets requestConfigAuthConfigApiKeyValue
     *
     * @param string|null $requestConfigAuthConfigApiKeyValue API key value  Only used if the auth type is api_key
     *
     * @return self
     */
    public function setRequestConfigAuthConfigApiKeyValue($requestConfigAuthConfigApiKeyValue)
    {
        $this->container['requestConfigAuthConfigApiKeyValue'] = $requestConfigAuthConfigApiKeyValue;

        return $this;
    }

    /**
     * Gets requestConfigAuthConfigBasicAuthPassword
     *
     * @return string|null
     */
    public function getRequestConfigAuthConfigBasicAuthPassword()
    {
        return $this->container['requestConfigAuthConfigBasicAuthPassword'];
    }

    /**
     * Sets requestConfigAuthConfigBasicAuthPassword
     *
     * @param string|null $requestConfigAuthConfigBasicAuthPassword Basic Auth Password  Only used if the auth type is basic_auth
     *
     * @return self
     */
    public function setRequestConfigAuthConfigBasicAuthPassword($requestConfigAuthConfigBasicAuthPassword)
    {
        $this->container['requestConfigAuthConfigBasicAuthPassword'] = $requestConfigAuthConfigBasicAuthPassword;

        return $this;
    }

    /**
     * Gets requestConfigAuthConfigBasicAuthUser
     *
     * @return string|null
     */
    public function getRequestConfigAuthConfigBasicAuthUser()
    {
        return $this->container['requestConfigAuthConfigBasicAuthUser'];
    }

    /**
     * Sets requestConfigAuthConfigBasicAuthUser
     *
     * @param string|null $requestConfigAuthConfigBasicAuthUser Basic Auth Username  Only used if the auth type is basic_auth
     *
     * @return self
     */
    public function setRequestConfigAuthConfigBasicAuthUser($requestConfigAuthConfigBasicAuthUser)
    {
        $this->container['requestConfigAuthConfigBasicAuthUser'] = $requestConfigAuthConfigBasicAuthUser;

        return $this;
    }

    /**
     * Gets requestConfigAuthType
     *
     * @return string|null
     */
    public function getRequestConfigAuthType()
    {
        return $this->container['requestConfigAuthType'];
    }

    /**
     * Sets requestConfigAuthType
     *
     * @param string|null $requestConfigAuthType HTTP Auth Method to use for the HTTP call  Can either be basic_auth or api_key basic_auth CourierChannelAuthTypeBasicAuth api_key CourierChannelAuthTypeApiKey
     *
     * @return self
     */
    public function setRequestConfigAuthType($requestConfigAuthType)
    {
        $allowedValues = $this->getRequestConfigAuthTypeAllowableValues();
        if (!is_null($requestConfigAuthType) && !in_array($requestConfigAuthType, $allowedValues, true)) {
            throw new \InvalidArgumentException(
                sprintf(
                    "Invalid value '%s' for 'requestConfigAuthType', must be one of '%s'",
                    $requestConfigAuthType,
                    implode("', '", $allowedValues)
                )
            );
        }
        $this->container['requestConfigAuthType'] = $requestConfigAuthType;

        return $this;
    }

    /**
     * Gets requestConfigBody
     *
     * @return string
     */
    public function getRequestConfigBody()
    {
        return $this->container['requestConfigBody'];
    }

    /**
     * Sets requestConfigBody
     *
     * @param string $requestConfigBody URI pointing to the JsonNet template used for HTTP body payload generation.
     *
     * @return self
     */
    public function setRequestConfigBody($requestConfigBody)
    {
        $this->container['requestConfigBody'] = $requestConfigBody;

        return $this;
    }

    /**
     * Gets requestConfigHeaders
     *
     * @return object|null
     */
    public function getRequestConfigHeaders()
    {
        return $this->container['requestConfigHeaders'];
    }

    /**
     * Sets requestConfigHeaders
     *
     * @param object|null $requestConfigHeaders NullJSONRawMessage represents a json.RawMessage that works well with JSON, SQL, and Swagger and is NULLable-
     *
     * @return self
     */
    public function setRequestConfigHeaders($requestConfigHeaders)
    {
        $this->container['requestConfigHeaders'] = $requestConfigHeaders;

        return $this;
    }

    /**
     * Gets requestConfigMethod
     *
     * @return string
     */
    public function getRequestConfigMethod()
    {
        return $this->container['requestConfigMethod'];
    }

    /**
     * Sets requestConfigMethod
     *
     * @param string $requestConfigMethod The HTTP method to use (GET, POST, etc) for the HTTP call
     *
     * @return self
     */
    public function setRequestConfigMethod($requestConfigMethod)
    {
        $this->container['requestConfigMethod'] = $requestConfigMethod;

        return $this;
    }

    /**
     * Gets requestConfigUrl
     *
     * @return string|null
     */
    public function getRequestConfigUrl()
    {
        return $this->container['requestConfigUrl'];
    }

    /**
     * Sets requestConfigUrl
     *
     * @param string|null $requestConfigUrl requestConfigUrl
     *
     * @return self
     */
    public function setRequestConfigUrl($requestConfigUrl)
    {
        $this->container['requestConfigUrl'] = $requestConfigUrl;

        return $this;
    }

    /**
     * Gets updatedAt
     *
     * @return \DateTime|null
     */
    public function getUpdatedAt()
    {
        return $this->container['updatedAt'];
    }

    /**
     * Sets updatedAt
     *
     * @param \DateTime|null $updatedAt Last upate time
     *
     * @return self
     */
    public function setUpdatedAt($updatedAt)
    {
        $this->container['updatedAt'] = $updatedAt;

        return $this;
    }
    /**
     * Returns true if offset exists. False otherwise.
     *
     * @param integer $offset Offset
     *
     * @return boolean
     */
    public function offsetExists($offset)
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
    public function offsetSet($offset, $value)
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
    public function offsetUnset($offset)
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



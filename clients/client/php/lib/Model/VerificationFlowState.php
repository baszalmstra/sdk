<?php
/**
 * VerificationFlowState
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
use \Ory\Client\ObjectSerializer;

/**
 * VerificationFlowState Class Doc Comment
 *
 * @category Class
 * @description The state represents the state of the verification flow.  choose_method: ask the user to choose a method (e.g. recover account via email) sent_email: the email has been sent to the user passed_challenge: the request was successful and the recovery challenge was passed.
 * @package  Ory\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 */
class VerificationFlowState
{
    /**
     * Possible values of this enum
     */
    public const CHOOSE_METHOD = 'choose_method';

    public const SENT_EMAIL = 'sent_email';

    public const PASSED_CHALLENGE = 'passed_challenge';

    /**
     * Gets allowable values of the enum
     * @return string[]
     */
    public static function getAllowableEnumValues()
    {
        return [
            self::CHOOSE_METHOD,
            self::SENT_EMAIL,
            self::PASSED_CHALLENGE
        ];
    }
}



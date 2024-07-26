# coding: utf-8

"""
    Ory APIs

    Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

    The version of the OpenAPI document: v1.14.3
    Contact: support@ory.sh
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""  # noqa: E501


from __future__ import annotations
import pprint
import re  # noqa: F401
import json

from datetime import datetime
from pydantic import BaseModel, ConfigDict, Field, StrictBool, StrictInt, StrictStr, field_validator
from typing import Any, ClassVar, Dict, List, Optional
from typing_extensions import Annotated
from typing import Optional, Set
from typing_extensions import Self

class OAuth2Client(BaseModel):
    """
    OAuth 2.0 Clients are used to perform OAuth 2.0 and OpenID Connect flows. Usually, OAuth 2.0 clients are generated for applications which want to consume your OAuth 2.0 or OpenID Connect capabilities.
    """ # noqa: E501
    access_token_strategy: Optional[StrictStr] = Field(default=None, description="OAuth 2.0 Access Token Strategy  AccessTokenStrategy is the strategy used to generate access tokens. Valid options are `jwt` and `opaque`. `jwt` is a bad idea, see https://www.ory.sh/docs/hydra/advanced#json-web-tokens Setting the stragegy here overrides the global setting in `strategies.access_token`.")
    allowed_cors_origins: Optional[List[StrictStr]] = None
    audience: Optional[List[StrictStr]] = None
    authorization_code_grant_access_token_lifespan: Optional[Annotated[str, Field(strict=True)]] = None
    authorization_code_grant_id_token_lifespan: Optional[Annotated[str, Field(strict=True)]] = None
    authorization_code_grant_refresh_token_lifespan: Optional[Annotated[str, Field(strict=True)]] = None
    backchannel_logout_session_required: Optional[StrictBool] = Field(default=None, description="OpenID Connect Back-Channel Logout Session Required  Boolean value specifying whether the RP requires that a sid (session ID) Claim be included in the Logout Token to identify the RP session with the OP when the backchannel_logout_uri is used. If omitted, the default value is false.")
    backchannel_logout_uri: Optional[StrictStr] = Field(default=None, description="OpenID Connect Back-Channel Logout URI  RP URL that will cause the RP to log itself out when sent a Logout Token by the OP.")
    client_credentials_grant_access_token_lifespan: Optional[Annotated[str, Field(strict=True)]] = None
    client_id: Optional[StrictStr] = Field(default=None, description="OAuth 2.0 Client ID  The ID is immutable. If no ID is provided, a UUID4 will be generated.")
    client_name: Optional[StrictStr] = Field(default=None, description="OAuth 2.0 Client Name  The human-readable name of the client to be presented to the end-user during authorization.")
    client_secret: Optional[StrictStr] = Field(default=None, description="OAuth 2.0 Client Secret  The secret will be included in the create request as cleartext, and then never again. The secret is kept in hashed format and is not recoverable once lost.")
    client_secret_expires_at: Optional[StrictInt] = Field(default=None, description="OAuth 2.0 Client Secret Expires At  The field is currently not supported and its value is always 0.")
    client_uri: Optional[StrictStr] = Field(default=None, description="OAuth 2.0 Client URI  ClientURI is a URL string of a web page providing information about the client. If present, the server SHOULD display this URL to the end-user in a clickable fashion.")
    contacts: Optional[List[StrictStr]] = None
    created_at: Optional[datetime] = Field(default=None, description="OAuth 2.0 Client Creation Date  CreatedAt returns the timestamp of the client's creation.")
    frontchannel_logout_session_required: Optional[StrictBool] = Field(default=None, description="OpenID Connect Front-Channel Logout Session Required  Boolean value specifying whether the RP requires that iss (issuer) and sid (session ID) query parameters be included to identify the RP session with the OP when the frontchannel_logout_uri is used. If omitted, the default value is false.")
    frontchannel_logout_uri: Optional[StrictStr] = Field(default=None, description="OpenID Connect Front-Channel Logout URI  RP URL that will cause the RP to log itself out when rendered in an iframe by the OP. An iss (issuer) query parameter and a sid (session ID) query parameter MAY be included by the OP to enable the RP to validate the request and to determine which of the potentially multiple sessions is to be logged out; if either is included, both MUST be.")
    grant_types: Optional[List[StrictStr]] = None
    implicit_grant_access_token_lifespan: Optional[Annotated[str, Field(strict=True)]] = None
    implicit_grant_id_token_lifespan: Optional[Annotated[str, Field(strict=True)]] = None
    jwks: Optional[Any] = Field(default=None, description="OAuth 2.0 Client JSON Web Key Set  Client's JSON Web Key Set [JWK] document, passed by value. The semantics of the jwks parameter are the same as the jwks_uri parameter, other than that the JWK Set is passed by value, rather than by reference. This parameter is intended only to be used by Clients that, for some reason, are unable to use the jwks_uri parameter, for instance, by native applications that might not have a location to host the contents of the JWK Set. If a Client can use jwks_uri, it MUST NOT use jwks. One significant downside of jwks is that it does not enable key rotation (which jwks_uri does, as described in Section 10 of OpenID Connect Core 1.0 [OpenID.Core]). The jwks_uri and jwks parameters MUST NOT be used together.")
    jwks_uri: Optional[StrictStr] = Field(default=None, description="OAuth 2.0 Client JSON Web Key Set URL  URL for the Client's JSON Web Key Set [JWK] document. If the Client signs requests to the Server, it contains the signing key(s) the Server uses to validate signatures from the Client. The JWK Set MAY also contain the Client's encryption keys(s), which are used by the Server to encrypt responses to the Client. When both signing and encryption keys are made available, a use (Key Use) parameter value is REQUIRED for all keys in the referenced JWK Set to indicate each key's intended usage. Although some algorithms allow the same key to be used for both signatures and encryption, doing so is NOT RECOMMENDED, as it is less secure. The JWK x5c parameter MAY be used to provide X.509 representations of keys provided. When used, the bare key values MUST still be present and MUST match those in the certificate.")
    jwt_bearer_grant_access_token_lifespan: Optional[Annotated[str, Field(strict=True)]] = None
    logo_uri: Optional[StrictStr] = Field(default=None, description="OAuth 2.0 Client Logo URI  A URL string referencing the client's logo.")
    metadata: Optional[Dict[str, Any]] = None
    owner: Optional[StrictStr] = Field(default=None, description="OAuth 2.0 Client Owner  Owner is a string identifying the owner of the OAuth 2.0 Client.")
    policy_uri: Optional[StrictStr] = Field(default=None, description="OAuth 2.0 Client Policy URI  PolicyURI is a URL string that points to a human-readable privacy policy document that describes how the deployment organization collects, uses, retains, and discloses personal data.")
    post_logout_redirect_uris: Optional[List[StrictStr]] = None
    redirect_uris: Optional[List[StrictStr]] = None
    refresh_token_grant_access_token_lifespan: Optional[Annotated[str, Field(strict=True)]] = None
    refresh_token_grant_id_token_lifespan: Optional[Annotated[str, Field(strict=True)]] = None
    refresh_token_grant_refresh_token_lifespan: Optional[Annotated[str, Field(strict=True)]] = None
    registration_access_token: Optional[StrictStr] = Field(default=None, description="OpenID Connect Dynamic Client Registration Access Token  RegistrationAccessToken can be used to update, get, or delete the OAuth2 Client. It is sent when creating a client using Dynamic Client Registration.")
    registration_client_uri: Optional[StrictStr] = Field(default=None, description="OpenID Connect Dynamic Client Registration URL  RegistrationClientURI is the URL used to update, get, or delete the OAuth2 Client.")
    request_object_signing_alg: Optional[StrictStr] = Field(default=None, description="OpenID Connect Request Object Signing Algorithm  JWS [JWS] alg algorithm [JWA] that MUST be used for signing Request Objects sent to the OP. All Request Objects from this Client MUST be rejected, if not signed with this algorithm.")
    request_uris: Optional[List[StrictStr]] = None
    response_types: Optional[List[StrictStr]] = None
    scope: Optional[StrictStr] = Field(default=None, description="OAuth 2.0 Client Scope  Scope is a string containing a space-separated list of scope values (as described in Section 3.3 of OAuth 2.0 [RFC6749]) that the client can use when requesting access tokens.")
    sector_identifier_uri: Optional[StrictStr] = Field(default=None, description="OpenID Connect Sector Identifier URI  URL using the https scheme to be used in calculating Pseudonymous Identifiers by the OP. The URL references a file with a single JSON array of redirect_uri values.")
    skip_consent: Optional[StrictBool] = Field(default=None, description="SkipConsent skips the consent screen for this client. This field can only be set from the admin API.")
    skip_logout_consent: Optional[StrictBool] = Field(default=None, description="SkipLogoutConsent skips the logout consent screen for this client. This field can only be set from the admin API.")
    subject_type: Optional[StrictStr] = Field(default=None, description="OpenID Connect Subject Type  The `subject_types_supported` Discovery parameter contains a list of the supported subject_type values for this server. Valid types include `pairwise` and `public`.")
    token_endpoint_auth_method: Optional[StrictStr] = Field(default='client_secret_basic', description="OAuth 2.0 Token Endpoint Authentication Method  Requested Client Authentication method for the Token Endpoint. The options are:  `client_secret_basic`: (default) Send `client_id` and `client_secret` as `application/x-www-form-urlencoded` encoded in the HTTP Authorization header. `client_secret_post`: Send `client_id` and `client_secret` as `application/x-www-form-urlencoded` in the HTTP body. `private_key_jwt`: Use JSON Web Tokens to authenticate the client. `none`: Used for public clients (native apps, mobile apps) which can not have secrets.")
    token_endpoint_auth_signing_alg: Optional[StrictStr] = Field(default=None, description="OAuth 2.0 Token Endpoint Signing Algorithm  Requested Client Authentication signing algorithm for the Token Endpoint.")
    tos_uri: Optional[StrictStr] = Field(default=None, description="OAuth 2.0 Client Terms of Service URI  A URL string pointing to a human-readable terms of service document for the client that describes a contractual relationship between the end-user and the client that the end-user accepts when authorizing the client.")
    updated_at: Optional[datetime] = Field(default=None, description="OAuth 2.0 Client Last Update Date  UpdatedAt returns the timestamp of the last update.")
    userinfo_signed_response_alg: Optional[StrictStr] = Field(default=None, description="OpenID Connect Request Userinfo Signed Response Algorithm  JWS alg algorithm [JWA] REQUIRED for signing UserInfo Responses. If this is specified, the response will be JWT [JWT] serialized, and signed using JWS. The default, if omitted, is for the UserInfo Response to return the Claims as a UTF-8 encoded JSON object using the application/json content-type.")
    additional_properties: Dict[str, Any] = {}
    __properties: ClassVar[List[str]] = ["access_token_strategy", "allowed_cors_origins", "audience", "authorization_code_grant_access_token_lifespan", "authorization_code_grant_id_token_lifespan", "authorization_code_grant_refresh_token_lifespan", "backchannel_logout_session_required", "backchannel_logout_uri", "client_credentials_grant_access_token_lifespan", "client_id", "client_name", "client_secret", "client_secret_expires_at", "client_uri", "contacts", "created_at", "frontchannel_logout_session_required", "frontchannel_logout_uri", "grant_types", "implicit_grant_access_token_lifespan", "implicit_grant_id_token_lifespan", "jwks", "jwks_uri", "jwt_bearer_grant_access_token_lifespan", "logo_uri", "metadata", "owner", "policy_uri", "post_logout_redirect_uris", "redirect_uris", "refresh_token_grant_access_token_lifespan", "refresh_token_grant_id_token_lifespan", "refresh_token_grant_refresh_token_lifespan", "registration_access_token", "registration_client_uri", "request_object_signing_alg", "request_uris", "response_types", "scope", "sector_identifier_uri", "skip_consent", "skip_logout_consent", "subject_type", "token_endpoint_auth_method", "token_endpoint_auth_signing_alg", "tos_uri", "updated_at", "userinfo_signed_response_alg"]

    @field_validator('authorization_code_grant_access_token_lifespan')
    def authorization_code_grant_access_token_lifespan_validate_regular_expression(cls, value):
        """Validates the regular expression"""
        if value is None:
            return value

        if not re.match(r"^[0-9]+(ns|us|ms|s|m|h)$", value):
            raise ValueError(r"must validate the regular expression /^[0-9]+(ns|us|ms|s|m|h)$/")
        return value

    @field_validator('authorization_code_grant_id_token_lifespan')
    def authorization_code_grant_id_token_lifespan_validate_regular_expression(cls, value):
        """Validates the regular expression"""
        if value is None:
            return value

        if not re.match(r"^[0-9]+(ns|us|ms|s|m|h)$", value):
            raise ValueError(r"must validate the regular expression /^[0-9]+(ns|us|ms|s|m|h)$/")
        return value

    @field_validator('authorization_code_grant_refresh_token_lifespan')
    def authorization_code_grant_refresh_token_lifespan_validate_regular_expression(cls, value):
        """Validates the regular expression"""
        if value is None:
            return value

        if not re.match(r"^[0-9]+(ns|us|ms|s|m|h)$", value):
            raise ValueError(r"must validate the regular expression /^[0-9]+(ns|us|ms|s|m|h)$/")
        return value

    @field_validator('client_credentials_grant_access_token_lifespan')
    def client_credentials_grant_access_token_lifespan_validate_regular_expression(cls, value):
        """Validates the regular expression"""
        if value is None:
            return value

        if not re.match(r"^[0-9]+(ns|us|ms|s|m|h)$", value):
            raise ValueError(r"must validate the regular expression /^[0-9]+(ns|us|ms|s|m|h)$/")
        return value

    @field_validator('implicit_grant_access_token_lifespan')
    def implicit_grant_access_token_lifespan_validate_regular_expression(cls, value):
        """Validates the regular expression"""
        if value is None:
            return value

        if not re.match(r"^[0-9]+(ns|us|ms|s|m|h)$", value):
            raise ValueError(r"must validate the regular expression /^[0-9]+(ns|us|ms|s|m|h)$/")
        return value

    @field_validator('implicit_grant_id_token_lifespan')
    def implicit_grant_id_token_lifespan_validate_regular_expression(cls, value):
        """Validates the regular expression"""
        if value is None:
            return value

        if not re.match(r"^[0-9]+(ns|us|ms|s|m|h)$", value):
            raise ValueError(r"must validate the regular expression /^[0-9]+(ns|us|ms|s|m|h)$/")
        return value

    @field_validator('jwt_bearer_grant_access_token_lifespan')
    def jwt_bearer_grant_access_token_lifespan_validate_regular_expression(cls, value):
        """Validates the regular expression"""
        if value is None:
            return value

        if not re.match(r"^[0-9]+(ns|us|ms|s|m|h)$", value):
            raise ValueError(r"must validate the regular expression /^[0-9]+(ns|us|ms|s|m|h)$/")
        return value

    @field_validator('refresh_token_grant_access_token_lifespan')
    def refresh_token_grant_access_token_lifespan_validate_regular_expression(cls, value):
        """Validates the regular expression"""
        if value is None:
            return value

        if not re.match(r"^[0-9]+(ns|us|ms|s|m|h)$", value):
            raise ValueError(r"must validate the regular expression /^[0-9]+(ns|us|ms|s|m|h)$/")
        return value

    @field_validator('refresh_token_grant_id_token_lifespan')
    def refresh_token_grant_id_token_lifespan_validate_regular_expression(cls, value):
        """Validates the regular expression"""
        if value is None:
            return value

        if not re.match(r"^[0-9]+(ns|us|ms|s|m|h)$", value):
            raise ValueError(r"must validate the regular expression /^[0-9]+(ns|us|ms|s|m|h)$/")
        return value

    @field_validator('refresh_token_grant_refresh_token_lifespan')
    def refresh_token_grant_refresh_token_lifespan_validate_regular_expression(cls, value):
        """Validates the regular expression"""
        if value is None:
            return value

        if not re.match(r"^[0-9]+(ns|us|ms|s|m|h)$", value):
            raise ValueError(r"must validate the regular expression /^[0-9]+(ns|us|ms|s|m|h)$/")
        return value

    model_config = ConfigDict(
        populate_by_name=True,
        validate_assignment=True,
        protected_namespaces=(),
    )


    def to_str(self) -> str:
        """Returns the string representation of the model using alias"""
        return pprint.pformat(self.model_dump(by_alias=True))

    def to_json(self) -> str:
        """Returns the JSON representation of the model using alias"""
        # TODO: pydantic v2: use .model_dump_json(by_alias=True, exclude_unset=True) instead
        return json.dumps(self.to_dict())

    @classmethod
    def from_json(cls, json_str: str) -> Optional[Self]:
        """Create an instance of OAuth2Client from a JSON string"""
        return cls.from_dict(json.loads(json_str))

    def to_dict(self) -> Dict[str, Any]:
        """Return the dictionary representation of the model using alias.

        This has the following differences from calling pydantic's
        `self.model_dump(by_alias=True)`:

        * `None` is only added to the output dict for nullable fields that
          were set at model initialization. Other fields with value `None`
          are ignored.
        * Fields in `self.additional_properties` are added to the output dict.
        """
        excluded_fields: Set[str] = set([
            "additional_properties",
        ])

        _dict = self.model_dump(
            by_alias=True,
            exclude=excluded_fields,
            exclude_none=True,
        )
        # puts key-value pairs in additional_properties in the top level
        if self.additional_properties is not None:
            for _key, _value in self.additional_properties.items():
                _dict[_key] = _value

        # set to None if authorization_code_grant_access_token_lifespan (nullable) is None
        # and model_fields_set contains the field
        if self.authorization_code_grant_access_token_lifespan is None and "authorization_code_grant_access_token_lifespan" in self.model_fields_set:
            _dict['authorization_code_grant_access_token_lifespan'] = None

        # set to None if authorization_code_grant_id_token_lifespan (nullable) is None
        # and model_fields_set contains the field
        if self.authorization_code_grant_id_token_lifespan is None and "authorization_code_grant_id_token_lifespan" in self.model_fields_set:
            _dict['authorization_code_grant_id_token_lifespan'] = None

        # set to None if authorization_code_grant_refresh_token_lifespan (nullable) is None
        # and model_fields_set contains the field
        if self.authorization_code_grant_refresh_token_lifespan is None and "authorization_code_grant_refresh_token_lifespan" in self.model_fields_set:
            _dict['authorization_code_grant_refresh_token_lifespan'] = None

        # set to None if client_credentials_grant_access_token_lifespan (nullable) is None
        # and model_fields_set contains the field
        if self.client_credentials_grant_access_token_lifespan is None and "client_credentials_grant_access_token_lifespan" in self.model_fields_set:
            _dict['client_credentials_grant_access_token_lifespan'] = None

        # set to None if implicit_grant_access_token_lifespan (nullable) is None
        # and model_fields_set contains the field
        if self.implicit_grant_access_token_lifespan is None and "implicit_grant_access_token_lifespan" in self.model_fields_set:
            _dict['implicit_grant_access_token_lifespan'] = None

        # set to None if implicit_grant_id_token_lifespan (nullable) is None
        # and model_fields_set contains the field
        if self.implicit_grant_id_token_lifespan is None and "implicit_grant_id_token_lifespan" in self.model_fields_set:
            _dict['implicit_grant_id_token_lifespan'] = None

        # set to None if jwks (nullable) is None
        # and model_fields_set contains the field
        if self.jwks is None and "jwks" in self.model_fields_set:
            _dict['jwks'] = None

        # set to None if jwt_bearer_grant_access_token_lifespan (nullable) is None
        # and model_fields_set contains the field
        if self.jwt_bearer_grant_access_token_lifespan is None and "jwt_bearer_grant_access_token_lifespan" in self.model_fields_set:
            _dict['jwt_bearer_grant_access_token_lifespan'] = None

        # set to None if refresh_token_grant_access_token_lifespan (nullable) is None
        # and model_fields_set contains the field
        if self.refresh_token_grant_access_token_lifespan is None and "refresh_token_grant_access_token_lifespan" in self.model_fields_set:
            _dict['refresh_token_grant_access_token_lifespan'] = None

        # set to None if refresh_token_grant_id_token_lifespan (nullable) is None
        # and model_fields_set contains the field
        if self.refresh_token_grant_id_token_lifespan is None and "refresh_token_grant_id_token_lifespan" in self.model_fields_set:
            _dict['refresh_token_grant_id_token_lifespan'] = None

        # set to None if refresh_token_grant_refresh_token_lifespan (nullable) is None
        # and model_fields_set contains the field
        if self.refresh_token_grant_refresh_token_lifespan is None and "refresh_token_grant_refresh_token_lifespan" in self.model_fields_set:
            _dict['refresh_token_grant_refresh_token_lifespan'] = None

        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of OAuth2Client from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "access_token_strategy": obj.get("access_token_strategy"),
            "allowed_cors_origins": obj.get("allowed_cors_origins"),
            "audience": obj.get("audience"),
            "authorization_code_grant_access_token_lifespan": obj.get("authorization_code_grant_access_token_lifespan"),
            "authorization_code_grant_id_token_lifespan": obj.get("authorization_code_grant_id_token_lifespan"),
            "authorization_code_grant_refresh_token_lifespan": obj.get("authorization_code_grant_refresh_token_lifespan"),
            "backchannel_logout_session_required": obj.get("backchannel_logout_session_required"),
            "backchannel_logout_uri": obj.get("backchannel_logout_uri"),
            "client_credentials_grant_access_token_lifespan": obj.get("client_credentials_grant_access_token_lifespan"),
            "client_id": obj.get("client_id"),
            "client_name": obj.get("client_name"),
            "client_secret": obj.get("client_secret"),
            "client_secret_expires_at": obj.get("client_secret_expires_at"),
            "client_uri": obj.get("client_uri"),
            "contacts": obj.get("contacts"),
            "created_at": obj.get("created_at"),
            "frontchannel_logout_session_required": obj.get("frontchannel_logout_session_required"),
            "frontchannel_logout_uri": obj.get("frontchannel_logout_uri"),
            "grant_types": obj.get("grant_types"),
            "implicit_grant_access_token_lifespan": obj.get("implicit_grant_access_token_lifespan"),
            "implicit_grant_id_token_lifespan": obj.get("implicit_grant_id_token_lifespan"),
            "jwks": obj.get("jwks"),
            "jwks_uri": obj.get("jwks_uri"),
            "jwt_bearer_grant_access_token_lifespan": obj.get("jwt_bearer_grant_access_token_lifespan"),
            "logo_uri": obj.get("logo_uri"),
            "metadata": obj.get("metadata"),
            "owner": obj.get("owner"),
            "policy_uri": obj.get("policy_uri"),
            "post_logout_redirect_uris": obj.get("post_logout_redirect_uris"),
            "redirect_uris": obj.get("redirect_uris"),
            "refresh_token_grant_access_token_lifespan": obj.get("refresh_token_grant_access_token_lifespan"),
            "refresh_token_grant_id_token_lifespan": obj.get("refresh_token_grant_id_token_lifespan"),
            "refresh_token_grant_refresh_token_lifespan": obj.get("refresh_token_grant_refresh_token_lifespan"),
            "registration_access_token": obj.get("registration_access_token"),
            "registration_client_uri": obj.get("registration_client_uri"),
            "request_object_signing_alg": obj.get("request_object_signing_alg"),
            "request_uris": obj.get("request_uris"),
            "response_types": obj.get("response_types"),
            "scope": obj.get("scope"),
            "sector_identifier_uri": obj.get("sector_identifier_uri"),
            "skip_consent": obj.get("skip_consent"),
            "skip_logout_consent": obj.get("skip_logout_consent"),
            "subject_type": obj.get("subject_type"),
            "token_endpoint_auth_method": obj.get("token_endpoint_auth_method") if obj.get("token_endpoint_auth_method") is not None else 'client_secret_basic',
            "token_endpoint_auth_signing_alg": obj.get("token_endpoint_auth_signing_alg"),
            "tos_uri": obj.get("tos_uri"),
            "updated_at": obj.get("updated_at"),
            "userinfo_signed_response_alg": obj.get("userinfo_signed_response_alg")
        })
        # store additional fields in additional_properties
        for _key in obj.keys():
            if _key not in cls.__properties:
                _obj.additional_properties[_key] = obj.get(_key)

        return _obj



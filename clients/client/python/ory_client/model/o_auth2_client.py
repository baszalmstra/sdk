"""
    Ory APIs

    Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers.   # noqa: E501

    The version of the OpenAPI document: v1.1.37
    Contact: support@ory.sh
    Generated by: https://openapi-generator.tech
"""


import re  # noqa: F401
import sys  # noqa: F401

from ory_client.model_utils import (  # noqa: F401
    ApiTypeError,
    ModelComposed,
    ModelNormal,
    ModelSimple,
    cached_property,
    change_keys_js_to_python,
    convert_js_args_to_python_args,
    date,
    datetime,
    file_type,
    none_type,
    validate_get_composed_info,
    OpenApiModel
)
from ory_client.exceptions import ApiAttributeError


def lazy_import():
    from ory_client.model.null_duration import NullDuration
    from ory_client.model.string_slice_json_format import StringSliceJSONFormat
    globals()['NullDuration'] = NullDuration
    globals()['StringSliceJSONFormat'] = StringSliceJSONFormat


class OAuth2Client(ModelNormal):
    """NOTE: This class is auto generated by OpenAPI Generator.
    Ref: https://openapi-generator.tech

    Do not edit the class manually.

    Attributes:
      allowed_values (dict): The key is the tuple path to the attribute
          and the for var_name this is (var_name,). The value is a dict
          with a capitalized key describing the allowed value and an allowed
          value. These dicts store the allowed enum values.
      attribute_map (dict): The key is attribute name
          and the value is json key in definition.
      discriminator_value_class_map (dict): A dict to go from the discriminator
          variable value to the discriminator class name.
      validations (dict): The key is the tuple path to the attribute
          and the for var_name this is (var_name,). The value is a dict
          that stores validations for max_length, min_length, max_items,
          min_items, exclusive_maximum, inclusive_maximum, exclusive_minimum,
          inclusive_minimum, and regex.
      additional_properties_type (tuple): A tuple of classes accepted
          as additional properties values.
    """

    allowed_values = {
    }

    validations = {
    }

    @cached_property
    def additional_properties_type():
        """
        This must be a method because a model may have properties that are
        of type self, this must run after the class is loaded
        """
        lazy_import()
        return (bool, date, datetime, dict, float, int, list, str, none_type,)  # noqa: E501

    _nullable = False

    @cached_property
    def openapi_types():
        """
        This must be a method because a model may have properties that are
        of type self, this must run after the class is loaded

        Returns
            openapi_types (dict): The key is attribute name
                and the value is attribute type.
        """
        lazy_import()
        return {
            'access_token_strategy': (str,),  # noqa: E501
            'allowed_cors_origins': (StringSliceJSONFormat,),  # noqa: E501
            'audience': (StringSliceJSONFormat,),  # noqa: E501
            'authorization_code_grant_access_token_lifespan': (NullDuration,),  # noqa: E501
            'authorization_code_grant_id_token_lifespan': (NullDuration,),  # noqa: E501
            'authorization_code_grant_refresh_token_lifespan': (NullDuration,),  # noqa: E501
            'backchannel_logout_session_required': (bool,),  # noqa: E501
            'backchannel_logout_uri': (str,),  # noqa: E501
            'client_credentials_grant_access_token_lifespan': (NullDuration,),  # noqa: E501
            'client_id': (str,),  # noqa: E501
            'client_name': (str,),  # noqa: E501
            'client_secret': (str,),  # noqa: E501
            'client_secret_expires_at': (int,),  # noqa: E501
            'client_uri': (str,),  # noqa: E501
            'contacts': (StringSliceJSONFormat,),  # noqa: E501
            'created_at': (datetime,),  # noqa: E501
            'frontchannel_logout_session_required': (bool,),  # noqa: E501
            'frontchannel_logout_uri': (str,),  # noqa: E501
            'grant_types': (StringSliceJSONFormat,),  # noqa: E501
            'implicit_grant_access_token_lifespan': (NullDuration,),  # noqa: E501
            'implicit_grant_id_token_lifespan': (NullDuration,),  # noqa: E501
            'jwks': (bool, date, datetime, dict, float, int, list, str, none_type,),  # noqa: E501
            'jwks_uri': (str,),  # noqa: E501
            'jwt_bearer_grant_access_token_lifespan': (NullDuration,),  # noqa: E501
            'logo_uri': (str,),  # noqa: E501
            'metadata': ({str: (bool, date, datetime, dict, float, int, list, str, none_type)},),  # noqa: E501
            'owner': (str,),  # noqa: E501
            'policy_uri': (str,),  # noqa: E501
            'post_logout_redirect_uris': (StringSliceJSONFormat,),  # noqa: E501
            'redirect_uris': (StringSliceJSONFormat,),  # noqa: E501
            'refresh_token_grant_access_token_lifespan': (NullDuration,),  # noqa: E501
            'refresh_token_grant_id_token_lifespan': (NullDuration,),  # noqa: E501
            'refresh_token_grant_refresh_token_lifespan': (NullDuration,),  # noqa: E501
            'registration_access_token': (str,),  # noqa: E501
            'registration_client_uri': (str,),  # noqa: E501
            'request_object_signing_alg': (str,),  # noqa: E501
            'request_uris': (StringSliceJSONFormat,),  # noqa: E501
            'response_types': (StringSliceJSONFormat,),  # noqa: E501
            'scope': (str,),  # noqa: E501
            'sector_identifier_uri': (str,),  # noqa: E501
            'skip_consent': (bool,),  # noqa: E501
            'subject_type': (str,),  # noqa: E501
            'token_endpoint_auth_method': (str,),  # noqa: E501
            'token_endpoint_auth_signing_alg': (str,),  # noqa: E501
            'tos_uri': (str,),  # noqa: E501
            'updated_at': (datetime,),  # noqa: E501
            'userinfo_signed_response_alg': (str,),  # noqa: E501
        }

    @cached_property
    def discriminator():
        return None


    attribute_map = {
        'access_token_strategy': 'access_token_strategy',  # noqa: E501
        'allowed_cors_origins': 'allowed_cors_origins',  # noqa: E501
        'audience': 'audience',  # noqa: E501
        'authorization_code_grant_access_token_lifespan': 'authorization_code_grant_access_token_lifespan',  # noqa: E501
        'authorization_code_grant_id_token_lifespan': 'authorization_code_grant_id_token_lifespan',  # noqa: E501
        'authorization_code_grant_refresh_token_lifespan': 'authorization_code_grant_refresh_token_lifespan',  # noqa: E501
        'backchannel_logout_session_required': 'backchannel_logout_session_required',  # noqa: E501
        'backchannel_logout_uri': 'backchannel_logout_uri',  # noqa: E501
        'client_credentials_grant_access_token_lifespan': 'client_credentials_grant_access_token_lifespan',  # noqa: E501
        'client_id': 'client_id',  # noqa: E501
        'client_name': 'client_name',  # noqa: E501
        'client_secret': 'client_secret',  # noqa: E501
        'client_secret_expires_at': 'client_secret_expires_at',  # noqa: E501
        'client_uri': 'client_uri',  # noqa: E501
        'contacts': 'contacts',  # noqa: E501
        'created_at': 'created_at',  # noqa: E501
        'frontchannel_logout_session_required': 'frontchannel_logout_session_required',  # noqa: E501
        'frontchannel_logout_uri': 'frontchannel_logout_uri',  # noqa: E501
        'grant_types': 'grant_types',  # noqa: E501
        'implicit_grant_access_token_lifespan': 'implicit_grant_access_token_lifespan',  # noqa: E501
        'implicit_grant_id_token_lifespan': 'implicit_grant_id_token_lifespan',  # noqa: E501
        'jwks': 'jwks',  # noqa: E501
        'jwks_uri': 'jwks_uri',  # noqa: E501
        'jwt_bearer_grant_access_token_lifespan': 'jwt_bearer_grant_access_token_lifespan',  # noqa: E501
        'logo_uri': 'logo_uri',  # noqa: E501
        'metadata': 'metadata',  # noqa: E501
        'owner': 'owner',  # noqa: E501
        'policy_uri': 'policy_uri',  # noqa: E501
        'post_logout_redirect_uris': 'post_logout_redirect_uris',  # noqa: E501
        'redirect_uris': 'redirect_uris',  # noqa: E501
        'refresh_token_grant_access_token_lifespan': 'refresh_token_grant_access_token_lifespan',  # noqa: E501
        'refresh_token_grant_id_token_lifespan': 'refresh_token_grant_id_token_lifespan',  # noqa: E501
        'refresh_token_grant_refresh_token_lifespan': 'refresh_token_grant_refresh_token_lifespan',  # noqa: E501
        'registration_access_token': 'registration_access_token',  # noqa: E501
        'registration_client_uri': 'registration_client_uri',  # noqa: E501
        'request_object_signing_alg': 'request_object_signing_alg',  # noqa: E501
        'request_uris': 'request_uris',  # noqa: E501
        'response_types': 'response_types',  # noqa: E501
        'scope': 'scope',  # noqa: E501
        'sector_identifier_uri': 'sector_identifier_uri',  # noqa: E501
        'skip_consent': 'skip_consent',  # noqa: E501
        'subject_type': 'subject_type',  # noqa: E501
        'token_endpoint_auth_method': 'token_endpoint_auth_method',  # noqa: E501
        'token_endpoint_auth_signing_alg': 'token_endpoint_auth_signing_alg',  # noqa: E501
        'tos_uri': 'tos_uri',  # noqa: E501
        'updated_at': 'updated_at',  # noqa: E501
        'userinfo_signed_response_alg': 'userinfo_signed_response_alg',  # noqa: E501
    }

    read_only_vars = {
    }

    _composed_schemas = {}

    @classmethod
    @convert_js_args_to_python_args
    def _from_openapi_data(cls, *args, **kwargs):  # noqa: E501
        """OAuth2Client - a model defined in OpenAPI

        Keyword Args:
            _check_type (bool): if True, values for parameters in openapi_types
                                will be type checked and a TypeError will be
                                raised if the wrong type is input.
                                Defaults to True
            _path_to_item (tuple/list): This is a list of keys or values to
                                drill down to the model in received_data
                                when deserializing a response
            _spec_property_naming (bool): True if the variable names in the input data
                                are serialized names, as specified in the OpenAPI document.
                                False if the variable names in the input data
                                are pythonic names, e.g. snake case (default)
            _configuration (Configuration): the instance to use when
                                deserializing a file_type parameter.
                                If passed, type conversion is attempted
                                If omitted no type conversion is done.
            _visited_composed_classes (tuple): This stores a tuple of
                                classes that we have traveled through so that
                                if we see that class again we will not use its
                                discriminator again.
                                When traveling through a discriminator, the
                                composed schema that is
                                is traveled through is added to this set.
                                For example if Animal has a discriminator
                                petType and we pass in "Dog", and the class Dog
                                allOf includes Animal, we move through Animal
                                once using the discriminator, and pick Dog.
                                Then in Dog, we will make an instance of the
                                Animal class but this time we won't travel
                                through its discriminator because we passed in
                                _visited_composed_classes = (Animal,)
            access_token_strategy (str): OAuth 2.0 Access Token Strategy  AccessTokenStrategy is the strategy used to generate access tokens. Valid options are `jwt` and `opaque`. `jwt` is a bad idea, see https://www.ory.sh/docs/hydra/advanced#json-web-tokens Setting the stragegy here overrides the global setting in `strategies.access_token`.. [optional]  # noqa: E501
            allowed_cors_origins (StringSliceJSONFormat): [optional]  # noqa: E501
            audience (StringSliceJSONFormat): [optional]  # noqa: E501
            authorization_code_grant_access_token_lifespan (NullDuration): [optional]  # noqa: E501
            authorization_code_grant_id_token_lifespan (NullDuration): [optional]  # noqa: E501
            authorization_code_grant_refresh_token_lifespan (NullDuration): [optional]  # noqa: E501
            backchannel_logout_session_required (bool): OpenID Connect Back-Channel Logout Session Required  Boolean value specifying whether the RP requires that a sid (session ID) Claim be included in the Logout Token to identify the RP session with the OP when the backchannel_logout_uri is used. If omitted, the default value is false.. [optional]  # noqa: E501
            backchannel_logout_uri (str): OpenID Connect Back-Channel Logout URI  RP URL that will cause the RP to log itself out when sent a Logout Token by the OP.. [optional]  # noqa: E501
            client_credentials_grant_access_token_lifespan (NullDuration): [optional]  # noqa: E501
            client_id (str): OAuth 2.0 Client ID  The ID is autogenerated and immutable.. [optional]  # noqa: E501
            client_name (str): OAuth 2.0 Client Name  The human-readable name of the client to be presented to the end-user during authorization.. [optional]  # noqa: E501
            client_secret (str): OAuth 2.0 Client Secret  The secret will be included in the create request as cleartext, and then never again. The secret is kept in hashed format and is not recoverable once lost.. [optional]  # noqa: E501
            client_secret_expires_at (int): OAuth 2.0 Client Secret Expires At  The field is currently not supported and its value is always 0.. [optional]  # noqa: E501
            client_uri (str): OAuth 2.0 Client URI  ClientURI is a URL string of a web page providing information about the client. If present, the server SHOULD display this URL to the end-user in a clickable fashion.. [optional]  # noqa: E501
            contacts (StringSliceJSONFormat): [optional]  # noqa: E501
            created_at (datetime): OAuth 2.0 Client Creation Date  CreatedAt returns the timestamp of the client's creation.. [optional]  # noqa: E501
            frontchannel_logout_session_required (bool): OpenID Connect Front-Channel Logout Session Required  Boolean value specifying whether the RP requires that iss (issuer) and sid (session ID) query parameters be included to identify the RP session with the OP when the frontchannel_logout_uri is used. If omitted, the default value is false.. [optional]  # noqa: E501
            frontchannel_logout_uri (str): OpenID Connect Front-Channel Logout URI  RP URL that will cause the RP to log itself out when rendered in an iframe by the OP. An iss (issuer) query parameter and a sid (session ID) query parameter MAY be included by the OP to enable the RP to validate the request and to determine which of the potentially multiple sessions is to be logged out; if either is included, both MUST be.. [optional]  # noqa: E501
            grant_types (StringSliceJSONFormat): [optional]  # noqa: E501
            implicit_grant_access_token_lifespan (NullDuration): [optional]  # noqa: E501
            implicit_grant_id_token_lifespan (NullDuration): [optional]  # noqa: E501
            jwks (bool, date, datetime, dict, float, int, list, str, none_type): OAuth 2.0 Client JSON Web Key Set  Client's JSON Web Key Set [JWK] document, passed by value. The semantics of the jwks parameter are the same as the jwks_uri parameter, other than that the JWK Set is passed by value, rather than by reference. This parameter is intended only to be used by Clients that, for some reason, are unable to use the jwks_uri parameter, for instance, by native applications that might not have a location to host the contents of the JWK Set. If a Client can use jwks_uri, it MUST NOT use jwks. One significant downside of jwks is that it does not enable key rotation (which jwks_uri does, as described in Section 10 of OpenID Connect Core 1.0 [OpenID.Core]). The jwks_uri and jwks parameters MUST NOT be used together.. [optional]  # noqa: E501
            jwks_uri (str): OAuth 2.0 Client JSON Web Key Set URL  URL for the Client's JSON Web Key Set [JWK] document. If the Client signs requests to the Server, it contains the signing key(s) the Server uses to validate signatures from the Client. The JWK Set MAY also contain the Client's encryption keys(s), which are used by the Server to encrypt responses to the Client. When both signing and encryption keys are made available, a use (Key Use) parameter value is REQUIRED for all keys in the referenced JWK Set to indicate each key's intended usage. Although some algorithms allow the same key to be used for both signatures and encryption, doing so is NOT RECOMMENDED, as it is less secure. The JWK x5c parameter MAY be used to provide X.509 representations of keys provided. When used, the bare key values MUST still be present and MUST match those in the certificate.. [optional]  # noqa: E501
            jwt_bearer_grant_access_token_lifespan (NullDuration): [optional]  # noqa: E501
            logo_uri (str): OAuth 2.0 Client Logo URI  A URL string referencing the client's logo.. [optional]  # noqa: E501
            metadata ({str: (bool, date, datetime, dict, float, int, list, str, none_type)}): [optional]  # noqa: E501
            owner (str): OAuth 2.0 Client Owner  Owner is a string identifying the owner of the OAuth 2.0 Client.. [optional]  # noqa: E501
            policy_uri (str): OAuth 2.0 Client Policy URI  PolicyURI is a URL string that points to a human-readable privacy policy document that describes how the deployment organization collects, uses, retains, and discloses personal data.. [optional]  # noqa: E501
            post_logout_redirect_uris (StringSliceJSONFormat): [optional]  # noqa: E501
            redirect_uris (StringSliceJSONFormat): [optional]  # noqa: E501
            refresh_token_grant_access_token_lifespan (NullDuration): [optional]  # noqa: E501
            refresh_token_grant_id_token_lifespan (NullDuration): [optional]  # noqa: E501
            refresh_token_grant_refresh_token_lifespan (NullDuration): [optional]  # noqa: E501
            registration_access_token (str): OpenID Connect Dynamic Client Registration Access Token  RegistrationAccessToken can be used to update, get, or delete the OAuth2 Client. It is sent when creating a client using Dynamic Client Registration.. [optional]  # noqa: E501
            registration_client_uri (str): OpenID Connect Dynamic Client Registration URL  RegistrationClientURI is the URL used to update, get, or delete the OAuth2 Client.. [optional]  # noqa: E501
            request_object_signing_alg (str): OpenID Connect Request Object Signing Algorithm  JWS [JWS] alg algorithm [JWA] that MUST be used for signing Request Objects sent to the OP. All Request Objects from this Client MUST be rejected, if not signed with this algorithm.. [optional]  # noqa: E501
            request_uris (StringSliceJSONFormat): [optional]  # noqa: E501
            response_types (StringSliceJSONFormat): [optional]  # noqa: E501
            scope (str): OAuth 2.0 Client Scope  Scope is a string containing a space-separated list of scope values (as described in Section 3.3 of OAuth 2.0 [RFC6749]) that the client can use when requesting access tokens.. [optional]  # noqa: E501
            sector_identifier_uri (str): OpenID Connect Sector Identifier URI  URL using the https scheme to be used in calculating Pseudonymous Identifiers by the OP. The URL references a file with a single JSON array of redirect_uri values.. [optional]  # noqa: E501
            skip_consent (bool): SkipConsent skips the consent screen for this client. This field can only be set from the admin API.. [optional]  # noqa: E501
            subject_type (str): OpenID Connect Subject Type  The `subject_types_supported` Discovery parameter contains a list of the supported subject_type values for this server. Valid types include `pairwise` and `public`.. [optional]  # noqa: E501
            token_endpoint_auth_method (str): OAuth 2.0 Token Endpoint Authentication Method  Requested Client Authentication method for the Token Endpoint. The options are:  `client_secret_basic`: (default) Send `client_id` and `client_secret` as `application/x-www-form-urlencoded` encoded in the HTTP Authorization header. `client_secret_post`: Send `client_id` and `client_secret` as `application/x-www-form-urlencoded` in the HTTP body. `private_key_jwt`: Use JSON Web Tokens to authenticate the client. `none`: Used for public clients (native apps, mobile apps) which can not have secrets.. [optional] if omitted the server will use the default value of "client_secret_basic"  # noqa: E501
            token_endpoint_auth_signing_alg (str): OAuth 2.0 Token Endpoint Signing Algorithm  Requested Client Authentication signing algorithm for the Token Endpoint.. [optional]  # noqa: E501
            tos_uri (str): OAuth 2.0 Client Terms of Service URI  A URL string pointing to a human-readable terms of service document for the client that describes a contractual relationship between the end-user and the client that the end-user accepts when authorizing the client.. [optional]  # noqa: E501
            updated_at (datetime): OAuth 2.0 Client Last Update Date  UpdatedAt returns the timestamp of the last update.. [optional]  # noqa: E501
            userinfo_signed_response_alg (str): OpenID Connect Request Userinfo Signed Response Algorithm  JWS alg algorithm [JWA] REQUIRED for signing UserInfo Responses. If this is specified, the response will be JWT [JWT] serialized, and signed using JWS. The default, if omitted, is for the UserInfo Response to return the Claims as a UTF-8 encoded JSON object using the application/json content-type.. [optional]  # noqa: E501
        """

        _check_type = kwargs.pop('_check_type', True)
        _spec_property_naming = kwargs.pop('_spec_property_naming', True)
        _path_to_item = kwargs.pop('_path_to_item', ())
        _configuration = kwargs.pop('_configuration', None)
        _visited_composed_classes = kwargs.pop('_visited_composed_classes', ())

        self = super(OpenApiModel, cls).__new__(cls)

        if args:
            for arg in args:
                if isinstance(arg, dict):
                    kwargs.update(arg)
                else:
                    raise ApiTypeError(
                        "Invalid positional arguments=%s passed to %s. Remove those invalid positional arguments." % (
                            args,
                            self.__class__.__name__,
                        ),
                        path_to_item=_path_to_item,
                        valid_classes=(self.__class__,),
                    )

        self._data_store = {}
        self._check_type = _check_type
        self._spec_property_naming = _spec_property_naming
        self._path_to_item = _path_to_item
        self._configuration = _configuration
        self._visited_composed_classes = _visited_composed_classes + (self.__class__,)

        for var_name, var_value in kwargs.items():
            if var_name not in self.attribute_map and \
                        self._configuration is not None and \
                        self._configuration.discard_unknown_keys and \
                        self.additional_properties_type is None:
                # discard variable.
                continue
            setattr(self, var_name, var_value)
        return self

    required_properties = set([
        '_data_store',
        '_check_type',
        '_spec_property_naming',
        '_path_to_item',
        '_configuration',
        '_visited_composed_classes',
    ])

    @convert_js_args_to_python_args
    def __init__(self, *args, **kwargs):  # noqa: E501
        """OAuth2Client - a model defined in OpenAPI

        Keyword Args:
            _check_type (bool): if True, values for parameters in openapi_types
                                will be type checked and a TypeError will be
                                raised if the wrong type is input.
                                Defaults to True
            _path_to_item (tuple/list): This is a list of keys or values to
                                drill down to the model in received_data
                                when deserializing a response
            _spec_property_naming (bool): True if the variable names in the input data
                                are serialized names, as specified in the OpenAPI document.
                                False if the variable names in the input data
                                are pythonic names, e.g. snake case (default)
            _configuration (Configuration): the instance to use when
                                deserializing a file_type parameter.
                                If passed, type conversion is attempted
                                If omitted no type conversion is done.
            _visited_composed_classes (tuple): This stores a tuple of
                                classes that we have traveled through so that
                                if we see that class again we will not use its
                                discriminator again.
                                When traveling through a discriminator, the
                                composed schema that is
                                is traveled through is added to this set.
                                For example if Animal has a discriminator
                                petType and we pass in "Dog", and the class Dog
                                allOf includes Animal, we move through Animal
                                once using the discriminator, and pick Dog.
                                Then in Dog, we will make an instance of the
                                Animal class but this time we won't travel
                                through its discriminator because we passed in
                                _visited_composed_classes = (Animal,)
            access_token_strategy (str): OAuth 2.0 Access Token Strategy  AccessTokenStrategy is the strategy used to generate access tokens. Valid options are `jwt` and `opaque`. `jwt` is a bad idea, see https://www.ory.sh/docs/hydra/advanced#json-web-tokens Setting the stragegy here overrides the global setting in `strategies.access_token`.. [optional]  # noqa: E501
            allowed_cors_origins (StringSliceJSONFormat): [optional]  # noqa: E501
            audience (StringSliceJSONFormat): [optional]  # noqa: E501
            authorization_code_grant_access_token_lifespan (NullDuration): [optional]  # noqa: E501
            authorization_code_grant_id_token_lifespan (NullDuration): [optional]  # noqa: E501
            authorization_code_grant_refresh_token_lifespan (NullDuration): [optional]  # noqa: E501
            backchannel_logout_session_required (bool): OpenID Connect Back-Channel Logout Session Required  Boolean value specifying whether the RP requires that a sid (session ID) Claim be included in the Logout Token to identify the RP session with the OP when the backchannel_logout_uri is used. If omitted, the default value is false.. [optional]  # noqa: E501
            backchannel_logout_uri (str): OpenID Connect Back-Channel Logout URI  RP URL that will cause the RP to log itself out when sent a Logout Token by the OP.. [optional]  # noqa: E501
            client_credentials_grant_access_token_lifespan (NullDuration): [optional]  # noqa: E501
            client_id (str): OAuth 2.0 Client ID  The ID is autogenerated and immutable.. [optional]  # noqa: E501
            client_name (str): OAuth 2.0 Client Name  The human-readable name of the client to be presented to the end-user during authorization.. [optional]  # noqa: E501
            client_secret (str): OAuth 2.0 Client Secret  The secret will be included in the create request as cleartext, and then never again. The secret is kept in hashed format and is not recoverable once lost.. [optional]  # noqa: E501
            client_secret_expires_at (int): OAuth 2.0 Client Secret Expires At  The field is currently not supported and its value is always 0.. [optional]  # noqa: E501
            client_uri (str): OAuth 2.0 Client URI  ClientURI is a URL string of a web page providing information about the client. If present, the server SHOULD display this URL to the end-user in a clickable fashion.. [optional]  # noqa: E501
            contacts (StringSliceJSONFormat): [optional]  # noqa: E501
            created_at (datetime): OAuth 2.0 Client Creation Date  CreatedAt returns the timestamp of the client's creation.. [optional]  # noqa: E501
            frontchannel_logout_session_required (bool): OpenID Connect Front-Channel Logout Session Required  Boolean value specifying whether the RP requires that iss (issuer) and sid (session ID) query parameters be included to identify the RP session with the OP when the frontchannel_logout_uri is used. If omitted, the default value is false.. [optional]  # noqa: E501
            frontchannel_logout_uri (str): OpenID Connect Front-Channel Logout URI  RP URL that will cause the RP to log itself out when rendered in an iframe by the OP. An iss (issuer) query parameter and a sid (session ID) query parameter MAY be included by the OP to enable the RP to validate the request and to determine which of the potentially multiple sessions is to be logged out; if either is included, both MUST be.. [optional]  # noqa: E501
            grant_types (StringSliceJSONFormat): [optional]  # noqa: E501
            implicit_grant_access_token_lifespan (NullDuration): [optional]  # noqa: E501
            implicit_grant_id_token_lifespan (NullDuration): [optional]  # noqa: E501
            jwks (bool, date, datetime, dict, float, int, list, str, none_type): OAuth 2.0 Client JSON Web Key Set  Client's JSON Web Key Set [JWK] document, passed by value. The semantics of the jwks parameter are the same as the jwks_uri parameter, other than that the JWK Set is passed by value, rather than by reference. This parameter is intended only to be used by Clients that, for some reason, are unable to use the jwks_uri parameter, for instance, by native applications that might not have a location to host the contents of the JWK Set. If a Client can use jwks_uri, it MUST NOT use jwks. One significant downside of jwks is that it does not enable key rotation (which jwks_uri does, as described in Section 10 of OpenID Connect Core 1.0 [OpenID.Core]). The jwks_uri and jwks parameters MUST NOT be used together.. [optional]  # noqa: E501
            jwks_uri (str): OAuth 2.0 Client JSON Web Key Set URL  URL for the Client's JSON Web Key Set [JWK] document. If the Client signs requests to the Server, it contains the signing key(s) the Server uses to validate signatures from the Client. The JWK Set MAY also contain the Client's encryption keys(s), which are used by the Server to encrypt responses to the Client. When both signing and encryption keys are made available, a use (Key Use) parameter value is REQUIRED for all keys in the referenced JWK Set to indicate each key's intended usage. Although some algorithms allow the same key to be used for both signatures and encryption, doing so is NOT RECOMMENDED, as it is less secure. The JWK x5c parameter MAY be used to provide X.509 representations of keys provided. When used, the bare key values MUST still be present and MUST match those in the certificate.. [optional]  # noqa: E501
            jwt_bearer_grant_access_token_lifespan (NullDuration): [optional]  # noqa: E501
            logo_uri (str): OAuth 2.0 Client Logo URI  A URL string referencing the client's logo.. [optional]  # noqa: E501
            metadata ({str: (bool, date, datetime, dict, float, int, list, str, none_type)}): [optional]  # noqa: E501
            owner (str): OAuth 2.0 Client Owner  Owner is a string identifying the owner of the OAuth 2.0 Client.. [optional]  # noqa: E501
            policy_uri (str): OAuth 2.0 Client Policy URI  PolicyURI is a URL string that points to a human-readable privacy policy document that describes how the deployment organization collects, uses, retains, and discloses personal data.. [optional]  # noqa: E501
            post_logout_redirect_uris (StringSliceJSONFormat): [optional]  # noqa: E501
            redirect_uris (StringSliceJSONFormat): [optional]  # noqa: E501
            refresh_token_grant_access_token_lifespan (NullDuration): [optional]  # noqa: E501
            refresh_token_grant_id_token_lifespan (NullDuration): [optional]  # noqa: E501
            refresh_token_grant_refresh_token_lifespan (NullDuration): [optional]  # noqa: E501
            registration_access_token (str): OpenID Connect Dynamic Client Registration Access Token  RegistrationAccessToken can be used to update, get, or delete the OAuth2 Client. It is sent when creating a client using Dynamic Client Registration.. [optional]  # noqa: E501
            registration_client_uri (str): OpenID Connect Dynamic Client Registration URL  RegistrationClientURI is the URL used to update, get, or delete the OAuth2 Client.. [optional]  # noqa: E501
            request_object_signing_alg (str): OpenID Connect Request Object Signing Algorithm  JWS [JWS] alg algorithm [JWA] that MUST be used for signing Request Objects sent to the OP. All Request Objects from this Client MUST be rejected, if not signed with this algorithm.. [optional]  # noqa: E501
            request_uris (StringSliceJSONFormat): [optional]  # noqa: E501
            response_types (StringSliceJSONFormat): [optional]  # noqa: E501
            scope (str): OAuth 2.0 Client Scope  Scope is a string containing a space-separated list of scope values (as described in Section 3.3 of OAuth 2.0 [RFC6749]) that the client can use when requesting access tokens.. [optional]  # noqa: E501
            sector_identifier_uri (str): OpenID Connect Sector Identifier URI  URL using the https scheme to be used in calculating Pseudonymous Identifiers by the OP. The URL references a file with a single JSON array of redirect_uri values.. [optional]  # noqa: E501
            skip_consent (bool): SkipConsent skips the consent screen for this client. This field can only be set from the admin API.. [optional]  # noqa: E501
            subject_type (str): OpenID Connect Subject Type  The `subject_types_supported` Discovery parameter contains a list of the supported subject_type values for this server. Valid types include `pairwise` and `public`.. [optional]  # noqa: E501
            token_endpoint_auth_method (str): OAuth 2.0 Token Endpoint Authentication Method  Requested Client Authentication method for the Token Endpoint. The options are:  `client_secret_basic`: (default) Send `client_id` and `client_secret` as `application/x-www-form-urlencoded` encoded in the HTTP Authorization header. `client_secret_post`: Send `client_id` and `client_secret` as `application/x-www-form-urlencoded` in the HTTP body. `private_key_jwt`: Use JSON Web Tokens to authenticate the client. `none`: Used for public clients (native apps, mobile apps) which can not have secrets.. [optional] if omitted the server will use the default value of "client_secret_basic"  # noqa: E501
            token_endpoint_auth_signing_alg (str): OAuth 2.0 Token Endpoint Signing Algorithm  Requested Client Authentication signing algorithm for the Token Endpoint.. [optional]  # noqa: E501
            tos_uri (str): OAuth 2.0 Client Terms of Service URI  A URL string pointing to a human-readable terms of service document for the client that describes a contractual relationship between the end-user and the client that the end-user accepts when authorizing the client.. [optional]  # noqa: E501
            updated_at (datetime): OAuth 2.0 Client Last Update Date  UpdatedAt returns the timestamp of the last update.. [optional]  # noqa: E501
            userinfo_signed_response_alg (str): OpenID Connect Request Userinfo Signed Response Algorithm  JWS alg algorithm [JWA] REQUIRED for signing UserInfo Responses. If this is specified, the response will be JWT [JWT] serialized, and signed using JWS. The default, if omitted, is for the UserInfo Response to return the Claims as a UTF-8 encoded JSON object using the application/json content-type.. [optional]  # noqa: E501
        """

        _check_type = kwargs.pop('_check_type', True)
        _spec_property_naming = kwargs.pop('_spec_property_naming', False)
        _path_to_item = kwargs.pop('_path_to_item', ())
        _configuration = kwargs.pop('_configuration', None)
        _visited_composed_classes = kwargs.pop('_visited_composed_classes', ())

        if args:
            for arg in args:
                if isinstance(arg, dict):
                    kwargs.update(arg)
                else:
                    raise ApiTypeError(
                        "Invalid positional arguments=%s passed to %s. Remove those invalid positional arguments." % (
                            args,
                            self.__class__.__name__,
                        ),
                        path_to_item=_path_to_item,
                        valid_classes=(self.__class__,),
                    )

        self._data_store = {}
        self._check_type = _check_type
        self._spec_property_naming = _spec_property_naming
        self._path_to_item = _path_to_item
        self._configuration = _configuration
        self._visited_composed_classes = _visited_composed_classes + (self.__class__,)

        for var_name, var_value in kwargs.items():
            if var_name not in self.attribute_map and \
                        self._configuration is not None and \
                        self._configuration.discard_unknown_keys and \
                        self.additional_properties_type is None:
                # discard variable.
                continue
            setattr(self, var_name, var_value)
            if var_name in self.read_only_vars:
                raise ApiAttributeError(f"`{var_name}` is a read-only attribute. Use `from_openapi_data` to instantiate "
                                     f"class with read only attributes.")

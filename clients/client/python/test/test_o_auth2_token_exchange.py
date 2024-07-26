# coding: utf-8

"""
    Ory APIs

    Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

    The version of the OpenAPI document: v1.14.3
    Contact: support@ory.sh
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""  # noqa: E501


import unittest

from ory_client.models.o_auth2_token_exchange import OAuth2TokenExchange

class TestOAuth2TokenExchange(unittest.TestCase):
    """OAuth2TokenExchange unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> OAuth2TokenExchange:
        """Test OAuth2TokenExchange
            include_optional is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `OAuth2TokenExchange`
        """
        model = OAuth2TokenExchange()
        if include_optional:
            return OAuth2TokenExchange(
                access_token = '',
                expires_in = 56,
                id_token = '',
                refresh_token = '',
                scope = '',
                token_type = ''
            )
        else:
            return OAuth2TokenExchange(
        )
        """

    def testOAuth2TokenExchange(self):
        """Test OAuth2TokenExchange"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()

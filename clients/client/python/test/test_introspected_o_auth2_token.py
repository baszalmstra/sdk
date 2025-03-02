# coding: utf-8

"""
    Ory APIs

    Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

    The version of the OpenAPI document: v1.15.7
    Contact: support@ory.sh
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""  # noqa: E501


import unittest

from ory_client.models.introspected_o_auth2_token import IntrospectedOAuth2Token

class TestIntrospectedOAuth2Token(unittest.TestCase):
    """IntrospectedOAuth2Token unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> IntrospectedOAuth2Token:
        """Test IntrospectedOAuth2Token
            include_optional is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `IntrospectedOAuth2Token`
        """
        model = IntrospectedOAuth2Token()
        if include_optional:
            return IntrospectedOAuth2Token(
                active = True,
                aud = [
                    ''
                    ],
                client_id = '',
                exp = 56,
                ext = {
                    'key' : null
                    },
                iat = 56,
                iss = '',
                nbf = 56,
                obfuscated_subject = '',
                scope = '',
                sub = '',
                token_type = '',
                token_use = '',
                username = ''
            )
        else:
            return IntrospectedOAuth2Token(
                active = True,
        )
        """

    def testIntrospectedOAuth2Token(self):
        """Test IntrospectedOAuth2Token"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()

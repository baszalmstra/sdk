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

from ory_client.models.custom_domain import CustomDomain

class TestCustomDomain(unittest.TestCase):
    """CustomDomain unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> CustomDomain:
        """Test CustomDomain
            include_optional is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `CustomDomain`
        """
        model = CustomDomain()
        if include_optional:
            return CustomDomain(
                cookie_domain = '',
                cors_allowed_origins = [
                    ''
                    ],
                cors_enabled = True,
                created_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'),
                custom_ui_base_url = '',
                hostname = '',
                id = '',
                ssl_status = 'initializing',
                updated_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'),
                verification_errors = [
                    ''
                    ],
                verification_status = ''
            )
        else:
            return CustomDomain(
        )
        """

    def testCustomDomain(self):
        """Test CustomDomain"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()

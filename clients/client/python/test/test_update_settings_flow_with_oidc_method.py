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

from ory_client.models.update_settings_flow_with_oidc_method import UpdateSettingsFlowWithOidcMethod

class TestUpdateSettingsFlowWithOidcMethod(unittest.TestCase):
    """UpdateSettingsFlowWithOidcMethod unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> UpdateSettingsFlowWithOidcMethod:
        """Test UpdateSettingsFlowWithOidcMethod
            include_optional is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `UpdateSettingsFlowWithOidcMethod`
        """
        model = UpdateSettingsFlowWithOidcMethod()
        if include_optional:
            return UpdateSettingsFlowWithOidcMethod(
                flow = '',
                link = '',
                method = '',
                traits = ory_client.models.traits.traits(),
                transient_payload = ory_client.models.transient_payload.transient_payload(),
                unlink = '',
                upstream_parameters = ory_client.models.upstream_parameters.upstream_parameters()
            )
        else:
            return UpdateSettingsFlowWithOidcMethod(
                method = '',
        )
        """

    def testUpdateSettingsFlowWithOidcMethod(self):
        """Test UpdateSettingsFlowWithOidcMethod"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()

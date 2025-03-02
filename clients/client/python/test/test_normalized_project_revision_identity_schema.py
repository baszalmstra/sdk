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

from ory_client.models.normalized_project_revision_identity_schema import NormalizedProjectRevisionIdentitySchema

class TestNormalizedProjectRevisionIdentitySchema(unittest.TestCase):
    """NormalizedProjectRevisionIdentitySchema unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> NormalizedProjectRevisionIdentitySchema:
        """Test NormalizedProjectRevisionIdentitySchema
            include_optional is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `NormalizedProjectRevisionIdentitySchema`
        """
        model = NormalizedProjectRevisionIdentitySchema()
        if include_optional:
            return NormalizedProjectRevisionIdentitySchema(
                created_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'),
                id = '',
                identity_schema = ory_client.models.schema_represents_an_ory_kratos_identity_schema.Schema represents an Ory Kratos Identity Schema(
                    blob_name = '', 
                    blob_url = '', 
                    content_hash = '', 
                    created_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), 
                    id = '', 
                    name = 'CustomerIdentity', 
                    updated_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), ),
                identity_schema_id = '',
                import_id = '',
                import_url = 'base64://ey...',
                is_default = True,
                preset = '',
                project_revision_id = '',
                updated_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f')
            )
        else:
            return NormalizedProjectRevisionIdentitySchema(
        )
        """

    def testNormalizedProjectRevisionIdentitySchema(self):
        """Test NormalizedProjectRevisionIdentitySchema"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()

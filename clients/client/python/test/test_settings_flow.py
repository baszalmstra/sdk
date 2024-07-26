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

from ory_client.models.settings_flow import SettingsFlow

class TestSettingsFlow(unittest.TestCase):
    """SettingsFlow unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> SettingsFlow:
        """Test SettingsFlow
            include_optional is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `SettingsFlow`
        """
        model = SettingsFlow()
        if include_optional:
            return SettingsFlow(
                active = '',
                continue_with = [
                    null
                    ],
                expires_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'),
                id = '',
                identity = ory_client.models.identity_represents_an_ory_kratos_identity.Identity represents an Ory Kratos identity(
                    created_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), 
                    credentials = {
                        'key' : ory_client.models.identity_credentials.identityCredentials(
                            config = ory_client.models.json_raw_message_represents_a_json/raw_message_that_works_well_with_json,_sql,_and_swagger/.JSONRawMessage represents a json.RawMessage that works well with JSON, SQL, and Swagger.(), 
                            created_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), 
                            identifiers = [
                                ''
                                ], 
                            type = 'password', 
                            updated_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), 
                            version = 56, )
                        }, 
                    id = '', 
                    metadata_admin = ory_client.models.null_json_raw_message.nullJsonRawMessage(), 
                    metadata_public = ory_client.models.null_json_raw_message.nullJsonRawMessage(), 
                    organization_id = '', 
                    recovery_addresses = [
                        ory_client.models.recovery_identity_address.recoveryIdentityAddress(
                            created_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), 
                            id = '', 
                            updated_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), 
                            value = '', 
                            via = '', )
                        ], 
                    schema_id = '', 
                    schema_url = '', 
                    state = 'active', 
                    state_changed_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), 
                    traits = null, 
                    updated_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), 
                    verifiable_addresses = [
                        ory_client.models.verifiable_identity_address.verifiableIdentityAddress(
                            created_at = '2014-01-01T23:28:56.782Z', 
                            id = '', 
                            status = '', 
                            updated_at = '2014-01-01T23:28:56.782Z', 
                            value = '', 
                            verified = True, 
                            verified_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), 
                            via = 'email', )
                        ], ),
                issued_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'),
                request_url = '',
                return_to = '',
                state = None,
                transient_payload = None,
                type = '',
                ui = ory_client.models.ui_container.uiContainer(
                    action = '', 
                    messages = [
                        ory_client.models.ui_text.uiText(
                            context = ory_client.models.context.context(), 
                            id = 56, 
                            text = '', 
                            type = 'info', )
                        ], 
                    method = '', 
                    nodes = [
                        ory_client.models.node_represents_a_flow's_nodes.Node represents a flow's nodes(
                            attributes = null, 
                            group = 'default', 
                            messages = [
                                ory_client.models.ui_text.uiText(
                                    context = ory_client.models.context.context(), 
                                    id = 56, 
                                    text = '', 
                                    type = 'info', )
                                ], 
                            meta = ory_client.models.a_node's_meta_information.A Node's Meta Information(
                                label = , ), 
                            type = 'text', )
                        ], )
            )
        else:
            return SettingsFlow(
                expires_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'),
                id = '',
                identity = ory_client.models.identity_represents_an_ory_kratos_identity.Identity represents an Ory Kratos identity(
                    created_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), 
                    credentials = {
                        'key' : ory_client.models.identity_credentials.identityCredentials(
                            config = ory_client.models.json_raw_message_represents_a_json/raw_message_that_works_well_with_json,_sql,_and_swagger/.JSONRawMessage represents a json.RawMessage that works well with JSON, SQL, and Swagger.(), 
                            created_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), 
                            identifiers = [
                                ''
                                ], 
                            type = 'password', 
                            updated_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), 
                            version = 56, )
                        }, 
                    id = '', 
                    metadata_admin = ory_client.models.null_json_raw_message.nullJsonRawMessage(), 
                    metadata_public = ory_client.models.null_json_raw_message.nullJsonRawMessage(), 
                    organization_id = '', 
                    recovery_addresses = [
                        ory_client.models.recovery_identity_address.recoveryIdentityAddress(
                            created_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), 
                            id = '', 
                            updated_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), 
                            value = '', 
                            via = '', )
                        ], 
                    schema_id = '', 
                    schema_url = '', 
                    state = 'active', 
                    state_changed_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), 
                    traits = null, 
                    updated_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), 
                    verifiable_addresses = [
                        ory_client.models.verifiable_identity_address.verifiableIdentityAddress(
                            created_at = '2014-01-01T23:28:56.782Z', 
                            id = '', 
                            status = '', 
                            updated_at = '2014-01-01T23:28:56.782Z', 
                            value = '', 
                            verified = True, 
                            verified_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'), 
                            via = 'email', )
                        ], ),
                issued_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'),
                request_url = '',
                state = None,
                type = '',
                ui = ory_client.models.ui_container.uiContainer(
                    action = '', 
                    messages = [
                        ory_client.models.ui_text.uiText(
                            context = ory_client.models.context.context(), 
                            id = 56, 
                            text = '', 
                            type = 'info', )
                        ], 
                    method = '', 
                    nodes = [
                        ory_client.models.node_represents_a_flow's_nodes.Node represents a flow's nodes(
                            attributes = null, 
                            group = 'default', 
                            messages = [
                                ory_client.models.ui_text.uiText(
                                    context = ory_client.models.context.context(), 
                                    id = 56, 
                                    text = '', 
                                    type = 'info', )
                                ], 
                            meta = ory_client.models.a_node's_meta_information.A Node's Meta Information(
                                label = , ), 
                            type = 'text', )
                        ], ),
        )
        """

    def testSettingsFlow(self):
        """Test SettingsFlow"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()

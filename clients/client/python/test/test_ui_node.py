"""
    Ory APIs

    Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers.   # noqa: E501

    The version of the OpenAPI document: v1.1.37
    Contact: support@ory.sh
    Generated by: https://openapi-generator.tech
"""


import sys
import unittest

import ory_client
from ory_client.model.ui_node_attributes import UiNodeAttributes
from ory_client.model.ui_node_meta import UiNodeMeta
from ory_client.model.ui_texts import UiTexts
globals()['UiNodeAttributes'] = UiNodeAttributes
globals()['UiNodeMeta'] = UiNodeMeta
globals()['UiTexts'] = UiTexts
from ory_client.model.ui_node import UiNode


class TestUiNode(unittest.TestCase):
    """UiNode unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def testUiNode(self):
        """Test UiNode"""
        # FIXME: construct object with mandatory attributes with example values
        # model = UiNode()  # noqa: E501
        pass


if __name__ == '__main__':
    unittest.main()

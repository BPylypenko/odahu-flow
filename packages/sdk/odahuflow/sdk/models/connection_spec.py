# coding: utf-8

from __future__ import absolute_import
from datetime import date, datetime  # noqa: F401

from typing import List, Dict  # noqa: F401

from odahuflow.sdk.models.base_model_ import Model
from odahuflow.sdk.models import util


class ConnectionSpec(Model):
    """NOTE: This class is auto generated by the swagger code generator program.

    Do not edit the class manually.
    """

    def __init__(self, description: str=None, key_id: str=None, key_secret: str=None, password: str=None, public_key: str=None, reference: str=None, region: str=None, role: str=None, type: str=None, uri: str=None, username: str=None, web_ui_link: str=None):  # noqa: E501
        """ConnectionSpec - a model defined in Swagger

        :param description: The description of this ConnectionSpec.  # noqa: E501
        :type description: str
        :param key_id: The key_id of this ConnectionSpec.  # noqa: E501
        :type key_id: str
        :param key_secret: The key_secret of this ConnectionSpec.  # noqa: E501
        :type key_secret: str
        :param password: The password of this ConnectionSpec.  # noqa: E501
        :type password: str
        :param public_key: The public_key of this ConnectionSpec.  # noqa: E501
        :type public_key: str
        :param reference: The reference of this ConnectionSpec.  # noqa: E501
        :type reference: str
        :param region: The region of this ConnectionSpec.  # noqa: E501
        :type region: str
        :param role: The role of this ConnectionSpec.  # noqa: E501
        :type role: str
        :param type: The type of this ConnectionSpec.  # noqa: E501
        :type type: str
        :param uri: The uri of this ConnectionSpec.  # noqa: E501
        :type uri: str
        :param username: The username of this ConnectionSpec.  # noqa: E501
        :type username: str
        :param web_ui_link: The web_ui_link of this ConnectionSpec.  # noqa: E501
        :type web_ui_link: str
        """
        self.swagger_types = {
            'description': str,
            'key_id': str,
            'key_secret': str,
            'password': str,
            'public_key': str,
            'reference': str,
            'region': str,
            'role': str,
            'type': str,
            'uri': str,
            'username': str,
            'web_ui_link': str
        }

        self.attribute_map = {
            'description': 'description',
            'key_id': 'keyID',
            'key_secret': 'keySecret',
            'password': 'password',
            'public_key': 'publicKey',
            'reference': 'reference',
            'region': 'region',
            'role': 'role',
            'type': 'type',
            'uri': 'uri',
            'username': 'username',
            'web_ui_link': 'webUILink'
        }

        self._description = description
        self._key_id = key_id
        self._key_secret = key_secret
        self._password = password
        self._public_key = public_key
        self._reference = reference
        self._region = region
        self._role = role
        self._type = type
        self._uri = uri
        self._username = username
        self._web_ui_link = web_ui_link

    @classmethod
    def from_dict(cls, dikt) -> 'ConnectionSpec':
        """Returns the dict as a model

        :param dikt: A dict.
        :type: dict
        :return: The ConnectionSpec of this ConnectionSpec.  # noqa: E501
        :rtype: ConnectionSpec
        """
        return util.deserialize_model(dikt, cls)

    @property
    def description(self) -> str:
        """Gets the description of this ConnectionSpec.

        Custom description  # noqa: E501

        :return: The description of this ConnectionSpec.
        :rtype: str
        """
        return self._description

    @description.setter
    def description(self, description: str):
        """Sets the description of this ConnectionSpec.

        Custom description  # noqa: E501

        :param description: The description of this ConnectionSpec.
        :type description: str
        """

        self._description = description

    @property
    def key_id(self) -> str:
        """Gets the key_id of this ConnectionSpec.

        Key ID  # noqa: E501

        :return: The key_id of this ConnectionSpec.
        :rtype: str
        """
        return self._key_id

    @key_id.setter
    def key_id(self, key_id: str):
        """Sets the key_id of this ConnectionSpec.

        Key ID  # noqa: E501

        :param key_id: The key_id of this ConnectionSpec.
        :type key_id: str
        """

        self._key_id = key_id

    @property
    def key_secret(self) -> str:
        """Gets the key_secret of this ConnectionSpec.

        SSH or service account secret  # noqa: E501

        :return: The key_secret of this ConnectionSpec.
        :rtype: str
        """
        return self._key_secret

    @key_secret.setter
    def key_secret(self, key_secret: str):
        """Sets the key_secret of this ConnectionSpec.

        SSH or service account secret  # noqa: E501

        :param key_secret: The key_secret of this ConnectionSpec.
        :type key_secret: str
        """

        self._key_secret = key_secret

    @property
    def password(self) -> str:
        """Gets the password of this ConnectionSpec.

        Password  # noqa: E501

        :return: The password of this ConnectionSpec.
        :rtype: str
        """
        return self._password

    @password.setter
    def password(self, password: str):
        """Sets the password of this ConnectionSpec.

        Password  # noqa: E501

        :param password: The password of this ConnectionSpec.
        :type password: str
        """

        self._password = password

    @property
    def public_key(self) -> str:
        """Gets the public_key of this ConnectionSpec.

        SSH public key  # noqa: E501

        :return: The public_key of this ConnectionSpec.
        :rtype: str
        """
        return self._public_key

    @public_key.setter
    def public_key(self, public_key: str):
        """Sets the public_key of this ConnectionSpec.

        SSH public key  # noqa: E501

        :param public_key: The public_key of this ConnectionSpec.
        :type public_key: str
        """

        self._public_key = public_key

    @property
    def reference(self) -> str:
        """Gets the reference of this ConnectionSpec.

        VCS reference  # noqa: E501

        :return: The reference of this ConnectionSpec.
        :rtype: str
        """
        return self._reference

    @reference.setter
    def reference(self, reference: str):
        """Sets the reference of this ConnectionSpec.

        VCS reference  # noqa: E501

        :param reference: The reference of this ConnectionSpec.
        :type reference: str
        """

        self._reference = reference

    @property
    def region(self) -> str:
        """Gets the region of this ConnectionSpec.

        AWS region or GCP project  # noqa: E501

        :return: The region of this ConnectionSpec.
        :rtype: str
        """
        return self._region

    @region.setter
    def region(self, region: str):
        """Sets the region of this ConnectionSpec.

        AWS region or GCP project  # noqa: E501

        :param region: The region of this ConnectionSpec.
        :type region: str
        """

        self._region = region

    @property
    def role(self) -> str:
        """Gets the role of this ConnectionSpec.

        Service account role  # noqa: E501

        :return: The role of this ConnectionSpec.
        :rtype: str
        """
        return self._role

    @role.setter
    def role(self, role: str):
        """Sets the role of this ConnectionSpec.

        Service account role  # noqa: E501

        :param role: The role of this ConnectionSpec.
        :type role: str
        """

        self._role = role

    @property
    def type(self) -> str:
        """Gets the type of this ConnectionSpec.

        Required value. Available values:   * s3   * gcs   * azureblob   * git   * docker  # noqa: E501

        :return: The type of this ConnectionSpec.
        :rtype: str
        """
        return self._type

    @type.setter
    def type(self, type: str):
        """Sets the type of this ConnectionSpec.

        Required value. Available values:   * s3   * gcs   * azureblob   * git   * docker  # noqa: E501

        :param type: The type of this ConnectionSpec.
        :type type: str
        """

        self._type = type

    @property
    def uri(self) -> str:
        """Gets the uri of this ConnectionSpec.

        URI. It is required value  # noqa: E501

        :return: The uri of this ConnectionSpec.
        :rtype: str
        """
        return self._uri

    @uri.setter
    def uri(self, uri: str):
        """Sets the uri of this ConnectionSpec.

        URI. It is required value  # noqa: E501

        :param uri: The uri of this ConnectionSpec.
        :type uri: str
        """

        self._uri = uri

    @property
    def username(self) -> str:
        """Gets the username of this ConnectionSpec.

        Username  # noqa: E501

        :return: The username of this ConnectionSpec.
        :rtype: str
        """
        return self._username

    @username.setter
    def username(self, username: str):
        """Sets the username of this ConnectionSpec.

        Username  # noqa: E501

        :param username: The username of this ConnectionSpec.
        :type username: str
        """

        self._username = username

    @property
    def web_ui_link(self) -> str:
        """Gets the web_ui_link of this ConnectionSpec.

        Custom web UI link  # noqa: E501

        :return: The web_ui_link of this ConnectionSpec.
        :rtype: str
        """
        return self._web_ui_link

    @web_ui_link.setter
    def web_ui_link(self, web_ui_link: str):
        """Sets the web_ui_link of this ConnectionSpec.

        Custom web UI link  # noqa: E501

        :param web_ui_link: The web_ui_link of this ConnectionSpec.
        :type web_ui_link: str
        """

        self._web_ui_link = web_ui_link
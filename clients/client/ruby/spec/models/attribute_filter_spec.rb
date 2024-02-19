=begin
#Ory APIs

#Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

The version of the OpenAPI document: v1.6.2
Contact: support@ory.sh
Generated by: https://openapi-generator.tech
OpenAPI Generator version: 6.0.1

=end

require 'spec_helper'
require 'json'
require 'date'

# Unit tests for OryClient::AttributeFilter
# Automatically generated by openapi-generator (https://openapi-generator.tech)
# Please update as you see appropriate
describe OryClient::AttributeFilter do
  let(:instance) { OryClient::AttributeFilter.new }

  describe 'test an instance of AttributeFilter' do
    it 'should create an instance of AttributeFilter' do
      expect(instance).to be_instance_of(OryClient::AttributeFilter)
    end
  end
  describe 'test attribute "attribute"' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  describe 'test attribute "condition"' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
      # validator = Petstore::EnumTest::EnumAttributeValidator.new('String', ["equals", "not_equals", "contains", "not_contains", "regex", "not_regex", "set", "not_set"])
      # validator.allowable_values.each do |value|
      #   expect { instance.condition = value }.not_to raise_error
      # end
    end
  end

  describe 'test attribute "value"' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

end

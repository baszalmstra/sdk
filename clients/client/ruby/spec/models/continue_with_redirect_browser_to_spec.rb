=begin
#Ory APIs

#Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

The version of the OpenAPI document: v1.14.3
Contact: support@ory.sh
Generated by: https://openapi-generator.tech
Generator version: 7.7.0

=end

require 'spec_helper'
require 'json'
require 'date'

# Unit tests for OryClient::ContinueWithRedirectBrowserTo
# Automatically generated by openapi-generator (https://openapi-generator.tech)
# Please update as you see appropriate
describe OryClient::ContinueWithRedirectBrowserTo do
  let(:instance) { OryClient::ContinueWithRedirectBrowserTo.new }

  describe 'test an instance of ContinueWithRedirectBrowserTo' do
    it 'should create an instance of ContinueWithRedirectBrowserTo' do
      # uncomment below to test the instance creation
      #expect(instance).to be_instance_of(OryClient::ContinueWithRedirectBrowserTo)
    end
  end

  describe 'test attribute "action"' do
    it 'should work' do
      # assertion here. ref: https://rspec.info/features/3-12/rspec-expectations/built-in-matchers/
      # validator = Petstore::EnumTest::EnumAttributeValidator.new('String', ["redirect_browser_to"])
      # validator.allowable_values.each do |value|
      #   expect { instance.action = value }.not_to raise_error
      # end
    end
  end

  describe 'test attribute "redirect_browser_to"' do
    it 'should work' do
      # assertion here. ref: https://rspec.info/features/3-12/rspec-expectations/built-in-matchers/
    end
  end

end

=begin
#Ory APIs

#Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

The version of the OpenAPI document: v1.15.7
Contact: support@ory.sh
Generated by: https://openapi-generator.tech
Generator version: 7.7.0

=end

require 'spec_helper'
require 'json'

# Unit tests for OryClient::RelationshipApi
# Automatically generated by openapi-generator (https://openapi-generator.tech)
# Please update as you see appropriate
describe 'RelationshipApi' do
  before do
    # run before each test
    @api_instance = OryClient::RelationshipApi.new
  end

  after do
    # run after each test
  end

  describe 'test an instance of RelationshipApi' do
    it 'should create an instance of RelationshipApi' do
      expect(@api_instance).to be_instance_of(OryClient::RelationshipApi)
    end
  end

  # unit tests for check_opl_syntax
  # Check the syntax of an OPL file
  # The OPL file is expected in the body of the request.
  # @param [Hash] opts the optional parameters
  # @option opts [String] :body 
  # @return [CheckOplSyntaxResult]
  describe 'check_opl_syntax test' do
    it 'should work' do
      # assertion here. ref: https://rspec.info/features/3-12/rspec-expectations/built-in-matchers/
    end
  end

  # unit tests for create_relationship
  # Create a Relationship
  # Use this endpoint to create a relationship.
  # @param [Hash] opts the optional parameters
  # @option opts [CreateRelationshipBody] :create_relationship_body 
  # @return [Relationship]
  describe 'create_relationship test' do
    it 'should work' do
      # assertion here. ref: https://rspec.info/features/3-12/rspec-expectations/built-in-matchers/
    end
  end

  # unit tests for delete_relationships
  # Delete Relationships
  # Use this endpoint to delete relationships
  # @param [Hash] opts the optional parameters
  # @option opts [String] :namespace Namespace of the Relationship
  # @option opts [String] :object Object of the Relationship
  # @option opts [String] :relation Relation of the Relationship
  # @option opts [String] :subject_id SubjectID of the Relationship
  # @option opts [String] :subject_set_namespace Namespace of the Subject Set
  # @option opts [String] :subject_set_object Object of the Subject Set
  # @option opts [String] :subject_set_relation Relation of the Subject Set
  # @return [nil]
  describe 'delete_relationships test' do
    it 'should work' do
      # assertion here. ref: https://rspec.info/features/3-12/rspec-expectations/built-in-matchers/
    end
  end

  # unit tests for get_relationships
  # Query relationships
  # Get all relationships that match the query. Only the namespace field is required.
  # @param [Hash] opts the optional parameters
  # @option opts [String] :page_token 
  # @option opts [Integer] :page_size 
  # @option opts [String] :namespace Namespace of the Relationship
  # @option opts [String] :object Object of the Relationship
  # @option opts [String] :relation Relation of the Relationship
  # @option opts [String] :subject_id SubjectID of the Relationship
  # @option opts [String] :subject_set_namespace Namespace of the Subject Set
  # @option opts [String] :subject_set_object Object of the Subject Set
  # @option opts [String] :subject_set_relation Relation of the Subject Set
  # @return [Relationships]
  describe 'get_relationships test' do
    it 'should work' do
      # assertion here. ref: https://rspec.info/features/3-12/rspec-expectations/built-in-matchers/
    end
  end

  # unit tests for list_relationship_namespaces
  # Query namespaces
  # Get all namespaces
  # @param [Hash] opts the optional parameters
  # @return [RelationshipNamespaces]
  describe 'list_relationship_namespaces test' do
    it 'should work' do
      # assertion here. ref: https://rspec.info/features/3-12/rspec-expectations/built-in-matchers/
    end
  end

  # unit tests for patch_relationships
  # Patch Multiple Relationships
  # Use this endpoint to patch one or more relationships.
  # @param [Hash] opts the optional parameters
  # @option opts [Array<RelationshipPatch>] :relationship_patch 
  # @return [nil]
  describe 'patch_relationships test' do
    it 'should work' do
      # assertion here. ref: https://rspec.info/features/3-12/rspec-expectations/built-in-matchers/
    end
  end

end

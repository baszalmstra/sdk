=begin
#Ory APIs

#Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

The version of the OpenAPI document: v1.2.1
Contact: support@ory.sh
Generated by: https://openapi-generator.tech
OpenAPI Generator version: 6.0.1

=end

require 'spec_helper'
require 'json'

# Unit tests for OryClient::ProjectApi
# Automatically generated by openapi-generator (https://openapi-generator.tech)
# Please update as you see appropriate
describe 'ProjectApi' do
  before do
    # run before each test
    @api_instance = OryClient::ProjectApi.new
  end

  after do
    # run after each test
  end

  describe 'test an instance of ProjectApi' do
    it 'should create an instance of ProjectApi' do
      expect(@api_instance).to be_instance_of(OryClient::ProjectApi)
    end
  end

  # unit tests for create_project
  # Create a Project
  # Creates a new project.
  # @param [Hash] opts the optional parameters
  # @option opts [CreateProjectBody] :create_project_body 
  # @return [Project]
  describe 'create_project test' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  # unit tests for create_project_api_key
  # Create project API token
  # Create an API token for a project.
  # @param project The Project ID or Project slug
  # @param [Hash] opts the optional parameters
  # @option opts [CreateProjectApiKeyRequest] :create_project_api_key_request 
  # @return [ProjectApiKey]
  describe 'create_project_api_key test' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  # unit tests for delete_project_api_key
  # Delete project API token
  # Deletes an API token and immediately removes it.
  # @param project The Project ID or Project slug
  # @param token_id The Token ID
  # @param [Hash] opts the optional parameters
  # @return [nil]
  describe 'delete_project_api_key test' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  # unit tests for get_active_project_in_console
  # Returns the Ory Network Project selected in the Ory Network Console
  # Use this API to get your active project in the Ory Network Console UI.
  # @param [Hash] opts the optional parameters
  # @return [ActiveProjectInConsole]
  describe 'get_active_project_in_console test' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  # unit tests for get_project
  # Get a Project
  # Get a projects you have access to by its ID.
  # @param project_id Project ID  The project&#39;s ID.
  # @param [Hash] opts the optional parameters
  # @return [Project]
  describe 'get_project test' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  # unit tests for get_project_members
  # Get all members associated with this project
  # This endpoint requires the user to be a member of the project with the role &#x60;OWNER&#x60; or &#x60;DEVELOPER&#x60;.
  # @param project 
  # @param [Hash] opts the optional parameters
  # @return [Array<CloudAccount>]
  describe 'get_project_members test' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  # unit tests for get_project_metrics
  # Retrieves project metrics for the specified event type and time range
  # @param project_id Project ID
  # @param event_type The event type to query for
  # @param resolution The resolution of the buckets  The minimum resolution is 1 hour.
  # @param from The start time of the time window
  # @param to The end time of the time window
  # @param [Hash] opts the optional parameters
  # @return [GetProjectMetricsResponse]
  describe 'get_project_metrics test' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  # unit tests for list_project_api_keys
  # List a project&#39;s API Tokens
  # A list of all the project&#39;s API tokens.
  # @param project The Project ID or Project slug
  # @param [Hash] opts the optional parameters
  # @return [Array<ProjectApiKey>]
  describe 'list_project_api_keys test' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  # unit tests for list_projects
  # List All Projects
  # Lists all projects you have access to.
  # @param [Hash] opts the optional parameters
  # @return [Array<ProjectMetadata>]
  describe 'list_projects test' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  # unit tests for patch_project
  # Patch an Ory Network Project Configuration
  # Deprecated: Use the &#x60;patchProjectWithRevision&#x60; endpoint instead to specify the exact revision the patch was generated for.  This endpoints allows you to patch individual Ory Network project configuration keys for Ory&#39;s services (identity, permission, ...). The configuration format is fully compatible with the open source projects for the respective services (e.g. Ory Kratos for Identity, Ory Keto for Permissions).  This endpoint expects the &#x60;version&#x60; key to be set in the payload. If it is unset, it will try to import the config as if it is from the most recent version.  If you have an older version of a configuration, you should set the version key in the payload!  While this endpoint is able to process all configuration items related to features (e.g. password reset), it does not support operational configuration items (e.g. port, tracing, logging) otherwise available in the open source.  For configuration items that can not be translated to the Ory Network, this endpoint will return a list of warnings to help you understand which parts of your config could not be processed.
  # @param project_id Project ID  The project&#39;s ID.
  # @param [Hash] opts the optional parameters
  # @option opts [Array<JsonPatch>] :json_patch 
  # @return [SuccessfulProjectUpdate]
  describe 'patch_project test' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  # unit tests for purge_project
  # Irrecoverably purge a project
  # !! Use with extreme caution !!  Using this API endpoint you can purge (completely delete) a project and its data. This action can not be undone and will delete ALL your data.  !! Use with extreme caution !!
  # @param project_id Project ID  The project&#39;s ID.
  # @param [Hash] opts the optional parameters
  # @return [nil]
  describe 'purge_project test' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  # unit tests for remove_project_member
  # Remove a member associated with this project
  # This also sets their invite status to &#x60;REMOVED&#x60;. This endpoint requires the user to be a member of the project with the role &#x60;OWNER&#x60;.
  # @param project 
  # @param member 
  # @param [Hash] opts the optional parameters
  # @return [nil]
  describe 'remove_project_member test' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  # unit tests for set_active_project_in_console
  # Sets the Ory Network Project active in the Ory Network Console
  # Use this API to set your active project in the Ory Network Console UI.
  # @param [Hash] opts the optional parameters
  # @option opts [SetActiveProjectInConsoleBody] :set_active_project_in_console_body 
  # @return [nil]
  describe 'set_active_project_in_console test' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  # unit tests for set_project
  # Update an Ory Network Project Configuration
  # This endpoints allows you to update the Ory Network project configuration for individual services (identity, permission, ...). The configuration is fully compatible with the open source projects for the respective services (e.g. Ory Kratos for Identity, Ory Keto for Permissions).  This endpoint expects the &#x60;version&#x60; key to be set in the payload. If it is unset, it will try to import the config as if it is from the most recent version.  If you have an older version of a configuration, you should set the version key in the payload!  While this endpoint is able to process all configuration items related to features (e.g. password reset), it does not support operational configuration items (e.g. port, tracing, logging) otherwise available in the open source.  For configuration items that can not be translated to the Ory Network, this endpoint will return a list of warnings to help you understand which parts of your config could not be processed.  Be aware that updating any service&#39;s configuration will completely override your current configuration for that service!
  # @param project_id Project ID  The project&#39;s ID.
  # @param [Hash] opts the optional parameters
  # @option opts [SetProject] :set_project 
  # @return [SuccessfulProjectUpdate]
  describe 'set_project test' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

end

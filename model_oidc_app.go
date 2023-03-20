/*
OneLogin API

OpenAPI Specification for OneLogin

API version: 3.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onelogin

import (
	"encoding/json"
)

// checks if the OidcApp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OidcApp{}

// OidcApp struct for OidcApp
type OidcApp struct {
	// Apps unique ID in OneLogin.
	Id *int32 `json:"id,omitempty"`
	// The name of the app.
	Name string `json:"name"`
	// Indicates if the app is visible in the OneLogin portal.
	Visible bool `json:"visible"`
	// Freeform description of the app.
	Description string `json:"description"`
	// Freeform notes about the app.
	Notes *string `json:"notes,omitempty"`
	// A link to the apps icon url
	IconUrl *string `json:"icon_url,omitempty"`
	AuthMethod *AuthMethod `json:"auth_method,omitempty"`
	// The security policy assigned to the app.
	PolicyId int32 `json:"policy_id"`
	// Indicates whether or not administrators can access the app as a user that they have assumed control over.
	AllowAssumedSignin *bool `json:"allow_assumed_signin,omitempty"`
	// ID of the OneLogin portal tab that the app is assigned to.
	TabId *int32 `json:"tab_id,omitempty"`
	// ID of the connector to base the app from.
	ConnectorId int32 `json:"connector_id"`
	// the date the app was created
	CreatedAt *string `json:"created_at,omitempty"`
	// the date the app was last updated
	UpdatedAt *string `json:"updated_at,omitempty"`
	// List of Role IDs that are assigned to the app. On App Create or Update the entire array is replaced with the values provided.
	RoleIds []int32 `json:"role_ids,omitempty"`
	Provisioning *GenericAppProvisioning `json:"provisioning,omitempty"`
	Parameters *AppParameters `json:"parameters,omitempty"`
	EnforcementPoint *EnforcementPoint `json:"enforcement_point,omitempty"`
	Configuration ConfigurationOidc `json:"configuration"`
	Sso *SsoOidc `json:"sso,omitempty"`
}

// NewOidcApp instantiates a new OidcApp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOidcApp(name string, visible bool, description string, policyId int32, connectorId int32, configuration ConfigurationOidc) *OidcApp {
	this := OidcApp{}
	this.Name = name
	this.Visible = visible
	this.Description = description
	this.PolicyId = policyId
	this.ConnectorId = connectorId
	this.Configuration = configuration
	return &this
}

// NewOidcAppWithDefaults instantiates a new OidcApp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOidcAppWithDefaults() *OidcApp {
	this := OidcApp{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OidcApp) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OidcApp) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OidcApp) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *OidcApp) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *OidcApp) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OidcApp) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *OidcApp) SetName(v string) {
	o.Name = v
}

// GetVisible returns the Visible field value
func (o *OidcApp) GetVisible() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Visible
}

// GetVisibleOk returns a tuple with the Visible field value
// and a boolean to check if the value has been set.
func (o *OidcApp) GetVisibleOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Visible, true
}

// SetVisible sets field value
func (o *OidcApp) SetVisible(v bool) {
	o.Visible = v
}

// GetDescription returns the Description field value
func (o *OidcApp) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *OidcApp) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *OidcApp) SetDescription(v string) {
	o.Description = v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *OidcApp) GetNotes() string {
	if o == nil || isNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OidcApp) GetNotesOk() (*string, bool) {
	if o == nil || isNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *OidcApp) HasNotes() bool {
	if o != nil && !isNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *OidcApp) SetNotes(v string) {
	o.Notes = &v
}

// GetIconUrl returns the IconUrl field value if set, zero value otherwise.
func (o *OidcApp) GetIconUrl() string {
	if o == nil || isNil(o.IconUrl) {
		var ret string
		return ret
	}
	return *o.IconUrl
}

// GetIconUrlOk returns a tuple with the IconUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OidcApp) GetIconUrlOk() (*string, bool) {
	if o == nil || isNil(o.IconUrl) {
		return nil, false
	}
	return o.IconUrl, true
}

// HasIconUrl returns a boolean if a field has been set.
func (o *OidcApp) HasIconUrl() bool {
	if o != nil && !isNil(o.IconUrl) {
		return true
	}

	return false
}

// SetIconUrl gets a reference to the given string and assigns it to the IconUrl field.
func (o *OidcApp) SetIconUrl(v string) {
	o.IconUrl = &v
}

// GetAuthMethod returns the AuthMethod field value if set, zero value otherwise.
func (o *OidcApp) GetAuthMethod() AuthMethod {
	if o == nil || isNil(o.AuthMethod) {
		var ret AuthMethod
		return ret
	}
	return *o.AuthMethod
}

// GetAuthMethodOk returns a tuple with the AuthMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OidcApp) GetAuthMethodOk() (*AuthMethod, bool) {
	if o == nil || isNil(o.AuthMethod) {
		return nil, false
	}
	return o.AuthMethod, true
}

// HasAuthMethod returns a boolean if a field has been set.
func (o *OidcApp) HasAuthMethod() bool {
	if o != nil && !isNil(o.AuthMethod) {
		return true
	}

	return false
}

// SetAuthMethod gets a reference to the given AuthMethod and assigns it to the AuthMethod field.
func (o *OidcApp) SetAuthMethod(v AuthMethod) {
	o.AuthMethod = &v
}

// GetPolicyId returns the PolicyId field value
func (o *OidcApp) GetPolicyId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PolicyId
}

// GetPolicyIdOk returns a tuple with the PolicyId field value
// and a boolean to check if the value has been set.
func (o *OidcApp) GetPolicyIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PolicyId, true
}

// SetPolicyId sets field value
func (o *OidcApp) SetPolicyId(v int32) {
	o.PolicyId = v
}

// GetAllowAssumedSignin returns the AllowAssumedSignin field value if set, zero value otherwise.
func (o *OidcApp) GetAllowAssumedSignin() bool {
	if o == nil || isNil(o.AllowAssumedSignin) {
		var ret bool
		return ret
	}
	return *o.AllowAssumedSignin
}

// GetAllowAssumedSigninOk returns a tuple with the AllowAssumedSignin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OidcApp) GetAllowAssumedSigninOk() (*bool, bool) {
	if o == nil || isNil(o.AllowAssumedSignin) {
		return nil, false
	}
	return o.AllowAssumedSignin, true
}

// HasAllowAssumedSignin returns a boolean if a field has been set.
func (o *OidcApp) HasAllowAssumedSignin() bool {
	if o != nil && !isNil(o.AllowAssumedSignin) {
		return true
	}

	return false
}

// SetAllowAssumedSignin gets a reference to the given bool and assigns it to the AllowAssumedSignin field.
func (o *OidcApp) SetAllowAssumedSignin(v bool) {
	o.AllowAssumedSignin = &v
}

// GetTabId returns the TabId field value if set, zero value otherwise.
func (o *OidcApp) GetTabId() int32 {
	if o == nil || isNil(o.TabId) {
		var ret int32
		return ret
	}
	return *o.TabId
}

// GetTabIdOk returns a tuple with the TabId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OidcApp) GetTabIdOk() (*int32, bool) {
	if o == nil || isNil(o.TabId) {
		return nil, false
	}
	return o.TabId, true
}

// HasTabId returns a boolean if a field has been set.
func (o *OidcApp) HasTabId() bool {
	if o != nil && !isNil(o.TabId) {
		return true
	}

	return false
}

// SetTabId gets a reference to the given int32 and assigns it to the TabId field.
func (o *OidcApp) SetTabId(v int32) {
	o.TabId = &v
}

// GetConnectorId returns the ConnectorId field value
func (o *OidcApp) GetConnectorId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ConnectorId
}

// GetConnectorIdOk returns a tuple with the ConnectorId field value
// and a boolean to check if the value has been set.
func (o *OidcApp) GetConnectorIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectorId, true
}

// SetConnectorId sets field value
func (o *OidcApp) SetConnectorId(v int32) {
	o.ConnectorId = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *OidcApp) GetCreatedAt() string {
	if o == nil || isNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OidcApp) GetCreatedAtOk() (*string, bool) {
	if o == nil || isNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *OidcApp) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *OidcApp) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *OidcApp) GetUpdatedAt() string {
	if o == nil || isNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OidcApp) GetUpdatedAtOk() (*string, bool) {
	if o == nil || isNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *OidcApp) HasUpdatedAt() bool {
	if o != nil && !isNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *OidcApp) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetRoleIds returns the RoleIds field value if set, zero value otherwise.
func (o *OidcApp) GetRoleIds() []int32 {
	if o == nil || isNil(o.RoleIds) {
		var ret []int32
		return ret
	}
	return o.RoleIds
}

// GetRoleIdsOk returns a tuple with the RoleIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OidcApp) GetRoleIdsOk() ([]int32, bool) {
	if o == nil || isNil(o.RoleIds) {
		return nil, false
	}
	return o.RoleIds, true
}

// HasRoleIds returns a boolean if a field has been set.
func (o *OidcApp) HasRoleIds() bool {
	if o != nil && !isNil(o.RoleIds) {
		return true
	}

	return false
}

// SetRoleIds gets a reference to the given []int32 and assigns it to the RoleIds field.
func (o *OidcApp) SetRoleIds(v []int32) {
	o.RoleIds = v
}

// GetProvisioning returns the Provisioning field value if set, zero value otherwise.
func (o *OidcApp) GetProvisioning() GenericAppProvisioning {
	if o == nil || isNil(o.Provisioning) {
		var ret GenericAppProvisioning
		return ret
	}
	return *o.Provisioning
}

// GetProvisioningOk returns a tuple with the Provisioning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OidcApp) GetProvisioningOk() (*GenericAppProvisioning, bool) {
	if o == nil || isNil(o.Provisioning) {
		return nil, false
	}
	return o.Provisioning, true
}

// HasProvisioning returns a boolean if a field has been set.
func (o *OidcApp) HasProvisioning() bool {
	if o != nil && !isNil(o.Provisioning) {
		return true
	}

	return false
}

// SetProvisioning gets a reference to the given GenericAppProvisioning and assigns it to the Provisioning field.
func (o *OidcApp) SetProvisioning(v GenericAppProvisioning) {
	o.Provisioning = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *OidcApp) GetParameters() AppParameters {
	if o == nil || isNil(o.Parameters) {
		var ret AppParameters
		return ret
	}
	return *o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OidcApp) GetParametersOk() (*AppParameters, bool) {
	if o == nil || isNil(o.Parameters) {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *OidcApp) HasParameters() bool {
	if o != nil && !isNil(o.Parameters) {
		return true
	}

	return false
}

// SetParameters gets a reference to the given AppParameters and assigns it to the Parameters field.
func (o *OidcApp) SetParameters(v AppParameters) {
	o.Parameters = &v
}

// GetEnforcementPoint returns the EnforcementPoint field value if set, zero value otherwise.
func (o *OidcApp) GetEnforcementPoint() EnforcementPoint {
	if o == nil || isNil(o.EnforcementPoint) {
		var ret EnforcementPoint
		return ret
	}
	return *o.EnforcementPoint
}

// GetEnforcementPointOk returns a tuple with the EnforcementPoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OidcApp) GetEnforcementPointOk() (*EnforcementPoint, bool) {
	if o == nil || isNil(o.EnforcementPoint) {
		return nil, false
	}
	return o.EnforcementPoint, true
}

// HasEnforcementPoint returns a boolean if a field has been set.
func (o *OidcApp) HasEnforcementPoint() bool {
	if o != nil && !isNil(o.EnforcementPoint) {
		return true
	}

	return false
}

// SetEnforcementPoint gets a reference to the given EnforcementPoint and assigns it to the EnforcementPoint field.
func (o *OidcApp) SetEnforcementPoint(v EnforcementPoint) {
	o.EnforcementPoint = &v
}

// GetConfiguration returns the Configuration field value
func (o *OidcApp) GetConfiguration() ConfigurationOidc {
	if o == nil {
		var ret ConfigurationOidc
		return ret
	}

	return o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value
// and a boolean to check if the value has been set.
func (o *OidcApp) GetConfigurationOk() (*ConfigurationOidc, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Configuration, true
}

// SetConfiguration sets field value
func (o *OidcApp) SetConfiguration(v ConfigurationOidc) {
	o.Configuration = v
}

// GetSso returns the Sso field value if set, zero value otherwise.
func (o *OidcApp) GetSso() SsoOidc {
	if o == nil || isNil(o.Sso) {
		var ret SsoOidc
		return ret
	}
	return *o.Sso
}

// GetSsoOk returns a tuple with the Sso field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OidcApp) GetSsoOk() (*SsoOidc, bool) {
	if o == nil || isNil(o.Sso) {
		return nil, false
	}
	return o.Sso, true
}

// HasSso returns a boolean if a field has been set.
func (o *OidcApp) HasSso() bool {
	if o != nil && !isNil(o.Sso) {
		return true
	}

	return false
}

// SetSso gets a reference to the given SsoOidc and assigns it to the Sso field.
func (o *OidcApp) SetSso(v SsoOidc) {
	o.Sso = &v
}

func (o OidcApp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OidcApp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: id is readOnly
	toSerialize["name"] = o.Name
	toSerialize["visible"] = o.Visible
	toSerialize["description"] = o.Description
	if !isNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	if !isNil(o.IconUrl) {
		toSerialize["icon_url"] = o.IconUrl
	}
	if !isNil(o.AuthMethod) {
		toSerialize["auth_method"] = o.AuthMethod
	}
	toSerialize["policy_id"] = o.PolicyId
	if !isNil(o.AllowAssumedSignin) {
		toSerialize["allow_assumed_signin"] = o.AllowAssumedSignin
	}
	if !isNil(o.TabId) {
		toSerialize["tab_id"] = o.TabId
	}
	toSerialize["connector_id"] = o.ConnectorId
	if !isNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !isNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !isNil(o.RoleIds) {
		toSerialize["role_ids"] = o.RoleIds
	}
	if !isNil(o.Provisioning) {
		toSerialize["provisioning"] = o.Provisioning
	}
	if !isNil(o.Parameters) {
		toSerialize["parameters"] = o.Parameters
	}
	if !isNil(o.EnforcementPoint) {
		toSerialize["enforcement_point"] = o.EnforcementPoint
	}
	toSerialize["configuration"] = o.Configuration
	if !isNil(o.Sso) {
		toSerialize["sso"] = o.Sso
	}
	return toSerialize, nil
}

type NullableOidcApp struct {
	value *OidcApp
	isSet bool
}

func (v NullableOidcApp) Get() *OidcApp {
	return v.value
}

func (v *NullableOidcApp) Set(val *OidcApp) {
	v.value = val
	v.isSet = true
}

func (v NullableOidcApp) IsSet() bool {
	return v.isSet
}

func (v *NullableOidcApp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOidcApp(val *OidcApp) *NullableOidcApp {
	return &NullableOidcApp{value: val, isSet: true}
}

func (v NullableOidcApp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOidcApp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


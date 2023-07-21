// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package orgpolicy

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgresource"
	transport_tpg "github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/transport"
)

const OrgPolicyCustomConstraintAssetType string = "orgpolicy.googleapis.com/CustomConstraint"

func ResourceConverterOrgPolicyCustomConstraint() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType: OrgPolicyCustomConstraintAssetType,
		Convert:   GetOrgPolicyCustomConstraintCaiObject,
	}
}

func GetOrgPolicyCustomConstraintCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	name, err := tpgresource.AssetName(d, config, "//orgpolicy.googleapis.com/{{parent}}/customConstraints/{{name}}")
	if err != nil {
		return []tpgresource.Asset{}, err
	}
	if obj, err := GetOrgPolicyCustomConstraintApiObject(d, config); err == nil {
		return []tpgresource.Asset{{
			Name: name,
			Type: OrgPolicyCustomConstraintAssetType,
			Resource: &tpgresource.AssetResource{
				Version:              "v2",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/orgpolicy/v2/rest",
				DiscoveryName:        "CustomConstraint",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []tpgresource.Asset{}, err
	}
}

func GetOrgPolicyCustomConstraintApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	nameProp, err := expandOrgPolicyCustomConstraintName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	displayNameProp, err := expandOrgPolicyCustomConstraintDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	descriptionProp, err := expandOrgPolicyCustomConstraintDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	conditionProp, err := expandOrgPolicyCustomConstraintCondition(d.Get("condition"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("condition"); !tpgresource.IsEmptyValue(reflect.ValueOf(conditionProp)) && (ok || !reflect.DeepEqual(v, conditionProp)) {
		obj["condition"] = conditionProp
	}
	actionTypeProp, err := expandOrgPolicyCustomConstraintActionType(d.Get("action_type"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("action_type"); !tpgresource.IsEmptyValue(reflect.ValueOf(actionTypeProp)) && (ok || !reflect.DeepEqual(v, actionTypeProp)) {
		obj["actionType"] = actionTypeProp
	}
	methodTypesProp, err := expandOrgPolicyCustomConstraintMethodTypes(d.Get("method_types"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("method_types"); !tpgresource.IsEmptyValue(reflect.ValueOf(methodTypesProp)) && (ok || !reflect.DeepEqual(v, methodTypesProp)) {
		obj["methodTypes"] = methodTypesProp
	}
	resourceTypesProp, err := expandOrgPolicyCustomConstraintResourceTypes(d.Get("resource_types"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("resource_types"); !tpgresource.IsEmptyValue(reflect.ValueOf(resourceTypesProp)) && (ok || !reflect.DeepEqual(v, resourceTypesProp)) {
		obj["resourceTypes"] = resourceTypesProp
	}

	return obj, nil
}

func expandOrgPolicyCustomConstraintName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return tpgresource.ReplaceVars(d, config, "{{parent}}/customConstraints/{{name}}")
}

func expandOrgPolicyCustomConstraintDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOrgPolicyCustomConstraintDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOrgPolicyCustomConstraintCondition(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOrgPolicyCustomConstraintActionType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOrgPolicyCustomConstraintMethodTypes(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOrgPolicyCustomConstraintResourceTypes(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
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

package cloudrunv2

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgiamresource"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgresource"
	transport_tpg "github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/transport"
)

// Provide a separate asset type constant so we don't have to worry about name conflicts between IAM and non-IAM converter files
const CloudRunV2ServiceIAMAssetType string = "run.googleapis.com/Service"

func ResourceConverterCloudRunV2ServiceIamPolicy() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType:         CloudRunV2ServiceIAMAssetType,
		Convert:           GetCloudRunV2ServiceIamPolicyCaiObject,
		MergeCreateUpdate: MergeCloudRunV2ServiceIamPolicy,
	}
}

func ResourceConverterCloudRunV2ServiceIamBinding() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType:         CloudRunV2ServiceIAMAssetType,
		Convert:           GetCloudRunV2ServiceIamBindingCaiObject,
		FetchFullResource: FetchCloudRunV2ServiceIamPolicy,
		MergeCreateUpdate: MergeCloudRunV2ServiceIamBinding,
		MergeDelete:       MergeCloudRunV2ServiceIamBindingDelete,
	}
}

func ResourceConverterCloudRunV2ServiceIamMember() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType:         CloudRunV2ServiceIAMAssetType,
		Convert:           GetCloudRunV2ServiceIamMemberCaiObject,
		FetchFullResource: FetchCloudRunV2ServiceIamPolicy,
		MergeCreateUpdate: MergeCloudRunV2ServiceIamMember,
		MergeDelete:       MergeCloudRunV2ServiceIamMemberDelete,
	}
}

func GetCloudRunV2ServiceIamPolicyCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	return newCloudRunV2ServiceIamAsset(d, config, tpgiamresource.ExpandIamPolicyBindings)
}

func GetCloudRunV2ServiceIamBindingCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	return newCloudRunV2ServiceIamAsset(d, config, tpgiamresource.ExpandIamRoleBindings)
}

func GetCloudRunV2ServiceIamMemberCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	return newCloudRunV2ServiceIamAsset(d, config, tpgiamresource.ExpandIamMemberBindings)
}

func MergeCloudRunV2ServiceIamPolicy(existing, incoming tpgresource.Asset) tpgresource.Asset {
	existing.IAMPolicy = incoming.IAMPolicy
	return existing
}

func MergeCloudRunV2ServiceIamBinding(existing, incoming tpgresource.Asset) tpgresource.Asset {
	return tpgiamresource.MergeIamAssets(existing, incoming, tpgiamresource.MergeAuthoritativeBindings)
}

func MergeCloudRunV2ServiceIamBindingDelete(existing, incoming tpgresource.Asset) tpgresource.Asset {
	return tpgiamresource.MergeDeleteIamAssets(existing, incoming, tpgiamresource.MergeDeleteAuthoritativeBindings)
}

func MergeCloudRunV2ServiceIamMember(existing, incoming tpgresource.Asset) tpgresource.Asset {
	return tpgiamresource.MergeIamAssets(existing, incoming, tpgiamresource.MergeAdditiveBindings)
}

func MergeCloudRunV2ServiceIamMemberDelete(existing, incoming tpgresource.Asset) tpgresource.Asset {
	return tpgiamresource.MergeDeleteIamAssets(existing, incoming, tpgiamresource.MergeDeleteAdditiveBindings)
}

func newCloudRunV2ServiceIamAsset(
	d tpgresource.TerraformResourceData,
	config *transport_tpg.Config,
	expandBindings func(d tpgresource.TerraformResourceData) ([]tpgresource.IAMBinding, error),
) ([]tpgresource.Asset, error) {
	bindings, err := expandBindings(d)
	if err != nil {
		return []tpgresource.Asset{}, fmt.Errorf("expanding bindings: %v", err)
	}

	name, err := tpgresource.AssetName(d, config, "//run.googleapis.com/projects/{{project}}/locations/{{location}}/services/{{name}}")
	if err != nil {
		return []tpgresource.Asset{}, err
	}

	return []tpgresource.Asset{{
		Name: name,
		Type: CloudRunV2ServiceIAMAssetType,
		IAMPolicy: &tpgresource.IAMPolicy{
			Bindings: bindings,
		},
	}}, nil
}

func FetchCloudRunV2ServiceIamPolicy(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (tpgresource.Asset, error) {
	// Check if the identity field returns a value
	if _, ok := d.GetOk("location"); !ok {
		return tpgresource.Asset{}, tpgresource.ErrEmptyIdentityField
	}
	if _, ok := d.GetOk("name"); !ok {
		return tpgresource.Asset{}, tpgresource.ErrEmptyIdentityField
	}

	return tpgiamresource.FetchIamPolicy(
		CloudRunV2ServiceIamUpdaterProducer,
		d,
		config,
		"//run.googleapis.com/projects/{{project}}/locations/{{location}}/services/{{name}}",
		CloudRunV2ServiceIAMAssetType,
	)
}
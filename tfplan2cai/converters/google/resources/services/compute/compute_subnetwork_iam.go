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

package compute

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgiamresource"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgresource"
	transport_tpg "github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/transport"
)

// Provide a separate asset type constant so we don't have to worry about name conflicts between IAM and non-IAM converter files
const ComputeSubnetworkIAMAssetType string = "compute.googleapis.com/Subnetwork"

func ResourceConverterComputeSubnetworkIamPolicy() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType:         ComputeSubnetworkIAMAssetType,
		Convert:           GetComputeSubnetworkIamPolicyCaiObject,
		MergeCreateUpdate: MergeComputeSubnetworkIamPolicy,
	}
}

func ResourceConverterComputeSubnetworkIamBinding() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType:         ComputeSubnetworkIAMAssetType,
		Convert:           GetComputeSubnetworkIamBindingCaiObject,
		FetchFullResource: FetchComputeSubnetworkIamPolicy,
		MergeCreateUpdate: MergeComputeSubnetworkIamBinding,
		MergeDelete:       MergeComputeSubnetworkIamBindingDelete,
	}
}

func ResourceConverterComputeSubnetworkIamMember() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType:         ComputeSubnetworkIAMAssetType,
		Convert:           GetComputeSubnetworkIamMemberCaiObject,
		FetchFullResource: FetchComputeSubnetworkIamPolicy,
		MergeCreateUpdate: MergeComputeSubnetworkIamMember,
		MergeDelete:       MergeComputeSubnetworkIamMemberDelete,
	}
}

func GetComputeSubnetworkIamPolicyCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	return newComputeSubnetworkIamAsset(d, config, tpgiamresource.ExpandIamPolicyBindings)
}

func GetComputeSubnetworkIamBindingCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	return newComputeSubnetworkIamAsset(d, config, tpgiamresource.ExpandIamRoleBindings)
}

func GetComputeSubnetworkIamMemberCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	return newComputeSubnetworkIamAsset(d, config, tpgiamresource.ExpandIamMemberBindings)
}

func MergeComputeSubnetworkIamPolicy(existing, incoming tpgresource.Asset) tpgresource.Asset {
	existing.IAMPolicy = incoming.IAMPolicy
	return existing
}

func MergeComputeSubnetworkIamBinding(existing, incoming tpgresource.Asset) tpgresource.Asset {
	return tpgiamresource.MergeIamAssets(existing, incoming, tpgiamresource.MergeAuthoritativeBindings)
}

func MergeComputeSubnetworkIamBindingDelete(existing, incoming tpgresource.Asset) tpgresource.Asset {
	return tpgiamresource.MergeDeleteIamAssets(existing, incoming, tpgiamresource.MergeDeleteAuthoritativeBindings)
}

func MergeComputeSubnetworkIamMember(existing, incoming tpgresource.Asset) tpgresource.Asset {
	return tpgiamresource.MergeIamAssets(existing, incoming, tpgiamresource.MergeAdditiveBindings)
}

func MergeComputeSubnetworkIamMemberDelete(existing, incoming tpgresource.Asset) tpgresource.Asset {
	return tpgiamresource.MergeDeleteIamAssets(existing, incoming, tpgiamresource.MergeDeleteAdditiveBindings)
}

func newComputeSubnetworkIamAsset(
	d tpgresource.TerraformResourceData,
	config *transport_tpg.Config,
	expandBindings func(d tpgresource.TerraformResourceData) ([]tpgresource.IAMBinding, error),
) ([]tpgresource.Asset, error) {
	bindings, err := expandBindings(d)
	if err != nil {
		return []tpgresource.Asset{}, fmt.Errorf("expanding bindings: %v", err)
	}

	name, err := tpgresource.AssetName(d, config, "//compute.googleapis.com/projects/{{project}}/regions/{{region}}/subnetworks/{{subnetwork}}")
	if err != nil {
		return []tpgresource.Asset{}, err
	}

	return []tpgresource.Asset{{
		Name: name,
		Type: ComputeSubnetworkIAMAssetType,
		IAMPolicy: &tpgresource.IAMPolicy{
			Bindings: bindings,
		},
	}}, nil
}

func FetchComputeSubnetworkIamPolicy(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (tpgresource.Asset, error) {
	// Check if the identity field returns a value
	if _, ok := d.GetOk("region"); !ok {
		return tpgresource.Asset{}, tpgresource.ErrEmptyIdentityField
	}
	if _, ok := d.GetOk("subnetwork"); !ok {
		return tpgresource.Asset{}, tpgresource.ErrEmptyIdentityField
	}

	return tpgiamresource.FetchIamPolicy(
		ComputeSubnetworkIamUpdaterProducer,
		d,
		config,
		"//compute.googleapis.com/projects/{{project}}/regions/{{region}}/subnetworks/{{subnetwork}}",
		ComputeSubnetworkIAMAssetType,
	)
}
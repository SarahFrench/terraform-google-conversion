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

package dataplex

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgiamresource"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgresource"
	transport_tpg "github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/transport"
)

// Provide a separate asset type constant so we don't have to worry about name conflicts between IAM and non-IAM converter files
const DataplexZoneIAMAssetType string = "dataplex.googleapis.com/Zone"

func ResourceConverterDataplexZoneIamPolicy() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType:         DataplexZoneIAMAssetType,
		Convert:           GetDataplexZoneIamPolicyCaiObject,
		MergeCreateUpdate: MergeDataplexZoneIamPolicy,
	}
}

func ResourceConverterDataplexZoneIamBinding() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType:         DataplexZoneIAMAssetType,
		Convert:           GetDataplexZoneIamBindingCaiObject,
		FetchFullResource: FetchDataplexZoneIamPolicy,
		MergeCreateUpdate: MergeDataplexZoneIamBinding,
		MergeDelete:       MergeDataplexZoneIamBindingDelete,
	}
}

func ResourceConverterDataplexZoneIamMember() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType:         DataplexZoneIAMAssetType,
		Convert:           GetDataplexZoneIamMemberCaiObject,
		FetchFullResource: FetchDataplexZoneIamPolicy,
		MergeCreateUpdate: MergeDataplexZoneIamMember,
		MergeDelete:       MergeDataplexZoneIamMemberDelete,
	}
}

func GetDataplexZoneIamPolicyCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	return newDataplexZoneIamAsset(d, config, tpgiamresource.ExpandIamPolicyBindings)
}

func GetDataplexZoneIamBindingCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	return newDataplexZoneIamAsset(d, config, tpgiamresource.ExpandIamRoleBindings)
}

func GetDataplexZoneIamMemberCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	return newDataplexZoneIamAsset(d, config, tpgiamresource.ExpandIamMemberBindings)
}

func MergeDataplexZoneIamPolicy(existing, incoming tpgresource.Asset) tpgresource.Asset {
	existing.IAMPolicy = incoming.IAMPolicy
	return existing
}

func MergeDataplexZoneIamBinding(existing, incoming tpgresource.Asset) tpgresource.Asset {
	return tpgiamresource.MergeIamAssets(existing, incoming, tpgiamresource.MergeAuthoritativeBindings)
}

func MergeDataplexZoneIamBindingDelete(existing, incoming tpgresource.Asset) tpgresource.Asset {
	return tpgiamresource.MergeDeleteIamAssets(existing, incoming, tpgiamresource.MergeDeleteAuthoritativeBindings)
}

func MergeDataplexZoneIamMember(existing, incoming tpgresource.Asset) tpgresource.Asset {
	return tpgiamresource.MergeIamAssets(existing, incoming, tpgiamresource.MergeAdditiveBindings)
}

func MergeDataplexZoneIamMemberDelete(existing, incoming tpgresource.Asset) tpgresource.Asset {
	return tpgiamresource.MergeDeleteIamAssets(existing, incoming, tpgiamresource.MergeDeleteAdditiveBindings)
}

func newDataplexZoneIamAsset(
	d tpgresource.TerraformResourceData,
	config *transport_tpg.Config,
	expandBindings func(d tpgresource.TerraformResourceData) ([]tpgresource.IAMBinding, error),
) ([]tpgresource.Asset, error) {
	bindings, err := expandBindings(d)
	if err != nil {
		return []tpgresource.Asset{}, fmt.Errorf("expanding bindings: %v", err)
	}

	name, err := tpgresource.AssetName(d, config, "//dataplex.googleapis.com/projects/{{project}}/locations/{{location}}/lakes/{{lake}}/zones/{{dataplex_zone}}")
	if err != nil {
		return []tpgresource.Asset{}, err
	}

	return []tpgresource.Asset{{
		Name: name,
		Type: DataplexZoneIAMAssetType,
		IAMPolicy: &tpgresource.IAMPolicy{
			Bindings: bindings,
		},
	}}, nil
}

func FetchDataplexZoneIamPolicy(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (tpgresource.Asset, error) {
	// Check if the identity field returns a value
	if _, ok := d.GetOk("location"); !ok {
		return tpgresource.Asset{}, tpgresource.ErrEmptyIdentityField
	}
	if _, ok := d.GetOk("lake"); !ok {
		return tpgresource.Asset{}, tpgresource.ErrEmptyIdentityField
	}
	if _, ok := d.GetOk("dataplex_zone"); !ok {
		return tpgresource.Asset{}, tpgresource.ErrEmptyIdentityField
	}

	return tpgiamresource.FetchIamPolicy(
		DataplexZoneIamUpdaterProducer,
		d,
		config,
		"//dataplex.googleapis.com/projects/{{project}}/locations/{{location}}/lakes/{{lake}}/zones/{{dataplex_zone}}",
		DataplexZoneIAMAssetType,
	)
}
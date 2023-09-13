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

package storageinsights

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const StorageInsightsReportConfigAssetType string = "storageinsights.googleapis.com/ReportConfig"

func ResourceConverterStorageInsightsReportConfig() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: StorageInsightsReportConfigAssetType,
		Convert:   GetStorageInsightsReportConfigCaiObject,
	}
}

func GetStorageInsightsReportConfigCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//storageinsights.googleapis.com/projects/{{project}}/locations/{{location}}/reportConfigs/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetStorageInsightsReportConfigApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: StorageInsightsReportConfigAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/storageinsights/v1/rest",
				DiscoveryName:        "ReportConfig",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetStorageInsightsReportConfigApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	frequencyOptionsProp, err := expandStorageInsightsReportConfigFrequencyOptions(d.Get("frequency_options"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("frequency_options"); !tpgresource.IsEmptyValue(reflect.ValueOf(frequencyOptionsProp)) && (ok || !reflect.DeepEqual(v, frequencyOptionsProp)) {
		obj["frequencyOptions"] = frequencyOptionsProp
	}
	csvOptionsProp, err := expandStorageInsightsReportConfigCsvOptions(d.Get("csv_options"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("csv_options"); !tpgresource.IsEmptyValue(reflect.ValueOf(csvOptionsProp)) && (ok || !reflect.DeepEqual(v, csvOptionsProp)) {
		obj["csvOptions"] = csvOptionsProp
	}
	objectMetadataReportOptionsProp, err := expandStorageInsightsReportConfigObjectMetadataReportOptions(d.Get("object_metadata_report_options"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("object_metadata_report_options"); !tpgresource.IsEmptyValue(reflect.ValueOf(objectMetadataReportOptionsProp)) && (ok || !reflect.DeepEqual(v, objectMetadataReportOptionsProp)) {
		obj["objectMetadataReportOptions"] = objectMetadataReportOptionsProp
	}
	displayNameProp, err := expandStorageInsightsReportConfigDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}

	return obj, nil
}

func expandStorageInsightsReportConfigFrequencyOptions(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedFrequency, err := expandStorageInsightsReportConfigFrequencyOptionsFrequency(original["frequency"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedFrequency); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["frequency"] = transformedFrequency
	}

	transformedStartDate, err := expandStorageInsightsReportConfigFrequencyOptionsStartDate(original["start_date"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedStartDate); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["startDate"] = transformedStartDate
	}

	transformedEndDate, err := expandStorageInsightsReportConfigFrequencyOptionsEndDate(original["end_date"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEndDate); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["endDate"] = transformedEndDate
	}

	return transformed, nil
}

func expandStorageInsightsReportConfigFrequencyOptionsFrequency(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandStorageInsightsReportConfigFrequencyOptionsStartDate(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedDay, err := expandStorageInsightsReportConfigFrequencyOptionsStartDateDay(original["day"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDay); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["day"] = transformedDay
	}

	transformedMonth, err := expandStorageInsightsReportConfigFrequencyOptionsStartDateMonth(original["month"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMonth); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["month"] = transformedMonth
	}

	transformedYear, err := expandStorageInsightsReportConfigFrequencyOptionsStartDateYear(original["year"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedYear); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["year"] = transformedYear
	}

	return transformed, nil
}

func expandStorageInsightsReportConfigFrequencyOptionsStartDateDay(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandStorageInsightsReportConfigFrequencyOptionsStartDateMonth(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandStorageInsightsReportConfigFrequencyOptionsStartDateYear(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandStorageInsightsReportConfigFrequencyOptionsEndDate(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedDay, err := expandStorageInsightsReportConfigFrequencyOptionsEndDateDay(original["day"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDay); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["day"] = transformedDay
	}

	transformedMonth, err := expandStorageInsightsReportConfigFrequencyOptionsEndDateMonth(original["month"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMonth); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["month"] = transformedMonth
	}

	transformedYear, err := expandStorageInsightsReportConfigFrequencyOptionsEndDateYear(original["year"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedYear); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["year"] = transformedYear
	}

	return transformed, nil
}

func expandStorageInsightsReportConfigFrequencyOptionsEndDateDay(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandStorageInsightsReportConfigFrequencyOptionsEndDateMonth(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandStorageInsightsReportConfigFrequencyOptionsEndDateYear(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandStorageInsightsReportConfigCsvOptions(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedRecordSeparator, err := expandStorageInsightsReportConfigCsvOptionsRecordSeparator(original["record_separator"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRecordSeparator); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["recordSeparator"] = transformedRecordSeparator
	}

	transformedDelimiter, err := expandStorageInsightsReportConfigCsvOptionsDelimiter(original["delimiter"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDelimiter); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["delimiter"] = transformedDelimiter
	}

	transformedHeaderRequired, err := expandStorageInsightsReportConfigCsvOptionsHeaderRequired(original["header_required"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedHeaderRequired); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["headerRequired"] = transformedHeaderRequired
	}

	return transformed, nil
}

func expandStorageInsightsReportConfigCsvOptionsRecordSeparator(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandStorageInsightsReportConfigCsvOptionsDelimiter(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandStorageInsightsReportConfigCsvOptionsHeaderRequired(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandStorageInsightsReportConfigObjectMetadataReportOptions(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedMetadataFields, err := expandStorageInsightsReportConfigObjectMetadataReportOptionsMetadataFields(original["metadata_fields"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMetadataFields); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["metadataFields"] = transformedMetadataFields
	}

	transformedStorageFilters, err := expandStorageInsightsReportConfigObjectMetadataReportOptionsStorageFilters(original["storage_filters"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedStorageFilters); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["storageFilters"] = transformedStorageFilters
	}

	transformedStorageDestinationOptions, err := expandStorageInsightsReportConfigObjectMetadataReportOptionsStorageDestinationOptions(original["storage_destination_options"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedStorageDestinationOptions); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["storageDestinationOptions"] = transformedStorageDestinationOptions
	}

	return transformed, nil
}

func expandStorageInsightsReportConfigObjectMetadataReportOptionsMetadataFields(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandStorageInsightsReportConfigObjectMetadataReportOptionsStorageFilters(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedBucket, err := expandStorageInsightsReportConfigObjectMetadataReportOptionsStorageFiltersBucket(original["bucket"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedBucket); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["bucket"] = transformedBucket
	}

	return transformed, nil
}

func expandStorageInsightsReportConfigObjectMetadataReportOptionsStorageFiltersBucket(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandStorageInsightsReportConfigObjectMetadataReportOptionsStorageDestinationOptions(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedBucket, err := expandStorageInsightsReportConfigObjectMetadataReportOptionsStorageDestinationOptionsBucket(original["bucket"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedBucket); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["bucket"] = transformedBucket
	}

	transformedDestinationPath, err := expandStorageInsightsReportConfigObjectMetadataReportOptionsStorageDestinationOptionsDestinationPath(original["destination_path"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDestinationPath); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["destinationPath"] = transformedDestinationPath
	}

	return transformed, nil
}

func expandStorageInsightsReportConfigObjectMetadataReportOptionsStorageDestinationOptionsBucket(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandStorageInsightsReportConfigObjectMetadataReportOptionsStorageDestinationOptionsDestinationPath(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandStorageInsightsReportConfigDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
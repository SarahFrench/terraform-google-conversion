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

package google

import (
	"context"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Whether the port should be set or not
func validatePortSpec(diff *schema.ResourceDiff, blockName string) error {
	block := diff.Get(blockName + ".0").(map[string]interface{})
	portSpec := block["port_specification"]
	portName := block["port_name"]
	port := block["port"]

	hasPort := (port != nil && port != 0)
	noName := (portName == nil || portName == "")

	if portSpec == "USE_NAMED_PORT" && hasPort {
		return fmt.Errorf("Error in %s: port cannot be specified when using port_specification USE_NAMED_PORT.", blockName)
	}
	if portSpec == "USE_NAMED_PORT" && noName {
		return fmt.Errorf("Error in %s: Must specify port_name when using USE_NAMED_PORT as port_specification.", blockName)
	}

	if portSpec == "USE_SERVING_PORT" && hasPort {
		return fmt.Errorf("Error in %s: port cannot be specified when using port_specification USE_SERVING_PORT.", blockName)
	}
	if portSpec == "USE_SERVING_PORT" && !noName {
		return fmt.Errorf("Error in %s: port_name cannot be specified when using port_specification USE_SERVING_PORT.", blockName)
	}

	return nil
}

func healthCheckCustomizeDiff(_ context.Context, diff *schema.ResourceDiff, v interface{}) error {
	if diff.Get("http_health_check") != nil {
		return validatePortSpec(diff, "http_health_check")
	}
	if diff.Get("https_health_check") != nil {
		return validatePortSpec(diff, "https_health_check")
	}
	if diff.Get("http2_health_check") != nil {
		return validatePortSpec(diff, "http2_health_check")
	}
	if diff.Get("tcp_health_check") != nil {
		return validatePortSpec(diff, "tcp_health_check")
	}
	if diff.Get("ssl_health_check") != nil {
		return validatePortSpec(diff, "ssl_health_check")
	}

	return nil
}

func portDiffSuppress(k, old, new string, d *schema.ResourceData) bool {
	b := strings.Split(k, ".")
	if len(b) > 2 {
		attr := b[2]

		if attr == "port" {
			var defaultPort int64

			blockType := b[0]

			switch blockType {
			case "http_health_check":
				defaultPort = 80
			case "https_health_check":
				defaultPort = 443
			case "http2_health_check":
				defaultPort = 443
			case "tcp_health_check":
				defaultPort = 80
			case "ssl_health_check":
				defaultPort = 443
			}

			oldPort, _ := strconv.Atoi(old)
			newPort, _ := strconv.Atoi(new)

			portSpec := d.Get(b[0] + ".0.port_specification")
			if int64(oldPort) == defaultPort && newPort == 0 && (portSpec == "USE_FIXED_PORT" || portSpec == "") {
				return true
			}
		}
	}

	return false
}

const ComputeHealthCheckAssetType string = "compute.googleapis.com/HealthCheck"

func resourceConverterComputeHealthCheck() ResourceConverter {
	return ResourceConverter{
		AssetType: ComputeHealthCheckAssetType,
		Convert:   GetComputeHealthCheckCaiObject,
	}
}

func GetComputeHealthCheckCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	name, err := assetName(d, config, "//compute.googleapis.com/projects/{{project}}/global/healthChecks/{{name}}")
	if err != nil {
		return []Asset{}, err
	}
	if obj, err := GetComputeHealthCheckApiObject(d, config); err == nil {
		return []Asset{{
			Name: name,
			Type: ComputeHealthCheckAssetType,
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/compute/v1/rest",
				DiscoveryName:        "HealthCheck",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []Asset{}, err
	}
}

func GetComputeHealthCheckApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	checkIntervalSecProp, err := expandComputeHealthCheckCheckIntervalSec(d.Get("check_interval_sec"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("check_interval_sec"); !isEmptyValue(reflect.ValueOf(checkIntervalSecProp)) && (ok || !reflect.DeepEqual(v, checkIntervalSecProp)) {
		obj["checkIntervalSec"] = checkIntervalSecProp
	}
	descriptionProp, err := expandComputeHealthCheckDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); ok || !reflect.DeepEqual(v, descriptionProp) {
		obj["description"] = descriptionProp
	}
	healthyThresholdProp, err := expandComputeHealthCheckHealthyThreshold(d.Get("healthy_threshold"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("healthy_threshold"); !isEmptyValue(reflect.ValueOf(healthyThresholdProp)) && (ok || !reflect.DeepEqual(v, healthyThresholdProp)) {
		obj["healthyThreshold"] = healthyThresholdProp
	}
	nameProp, err := expandComputeHealthCheckName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	timeoutSecProp, err := expandComputeHealthCheckTimeoutSec(d.Get("timeout_sec"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("timeout_sec"); !isEmptyValue(reflect.ValueOf(timeoutSecProp)) && (ok || !reflect.DeepEqual(v, timeoutSecProp)) {
		obj["timeoutSec"] = timeoutSecProp
	}
	unhealthyThresholdProp, err := expandComputeHealthCheckUnhealthyThreshold(d.Get("unhealthy_threshold"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("unhealthy_threshold"); !isEmptyValue(reflect.ValueOf(unhealthyThresholdProp)) && (ok || !reflect.DeepEqual(v, unhealthyThresholdProp)) {
		obj["unhealthyThreshold"] = unhealthyThresholdProp
	}
	httpHealthCheckProp, err := expandComputeHealthCheckHttpHealthCheck(d.Get("http_health_check"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("http_health_check"); !isEmptyValue(reflect.ValueOf(httpHealthCheckProp)) && (ok || !reflect.DeepEqual(v, httpHealthCheckProp)) {
		obj["httpHealthCheck"] = httpHealthCheckProp
	}
	httpsHealthCheckProp, err := expandComputeHealthCheckHttpsHealthCheck(d.Get("https_health_check"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("https_health_check"); !isEmptyValue(reflect.ValueOf(httpsHealthCheckProp)) && (ok || !reflect.DeepEqual(v, httpsHealthCheckProp)) {
		obj["httpsHealthCheck"] = httpsHealthCheckProp
	}
	tcpHealthCheckProp, err := expandComputeHealthCheckTcpHealthCheck(d.Get("tcp_health_check"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("tcp_health_check"); !isEmptyValue(reflect.ValueOf(tcpHealthCheckProp)) && (ok || !reflect.DeepEqual(v, tcpHealthCheckProp)) {
		obj["tcpHealthCheck"] = tcpHealthCheckProp
	}
	sslHealthCheckProp, err := expandComputeHealthCheckSslHealthCheck(d.Get("ssl_health_check"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("ssl_health_check"); !isEmptyValue(reflect.ValueOf(sslHealthCheckProp)) && (ok || !reflect.DeepEqual(v, sslHealthCheckProp)) {
		obj["sslHealthCheck"] = sslHealthCheckProp
	}
	http2HealthCheckProp, err := expandComputeHealthCheckHttp2HealthCheck(d.Get("http2_health_check"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("http2_health_check"); !isEmptyValue(reflect.ValueOf(http2HealthCheckProp)) && (ok || !reflect.DeepEqual(v, http2HealthCheckProp)) {
		obj["http2HealthCheck"] = http2HealthCheckProp
	}
	grpcHealthCheckProp, err := expandComputeHealthCheckGrpcHealthCheck(d.Get("grpc_health_check"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("grpc_health_check"); !isEmptyValue(reflect.ValueOf(grpcHealthCheckProp)) && (ok || !reflect.DeepEqual(v, grpcHealthCheckProp)) {
		obj["grpcHealthCheck"] = grpcHealthCheckProp
	}
	logConfigProp, err := expandComputeHealthCheckLogConfig(d.Get("log_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("log_config"); !isEmptyValue(reflect.ValueOf(logConfigProp)) && (ok || !reflect.DeepEqual(v, logConfigProp)) {
		obj["logConfig"] = logConfigProp
	}

	return resourceComputeHealthCheckEncoder(d, config, obj)
}

func resourceComputeHealthCheckEncoder(d TerraformResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {

	if _, ok := d.GetOk("http_health_check"); ok {
		hc := d.Get("http_health_check").([]interface{})[0]
		ps := hc.(map[string]interface{})["port_specification"]
		pn := hc.(map[string]interface{})["port_name"]

		if ps == "USE_FIXED_PORT" || (ps == "" && pn == "") {
			m := obj["httpHealthCheck"].(map[string]interface{})
			if m["port"] == nil {
				m["port"] = 80
			}
		}
		obj["type"] = "HTTP"
		return obj, nil
	}
	if _, ok := d.GetOk("https_health_check"); ok {
		hc := d.Get("https_health_check").([]interface{})[0]
		ps := hc.(map[string]interface{})["port_specification"]
		pn := hc.(map[string]interface{})["port_name"]

		if ps == "USE_FIXED_PORT" || (ps == "" && pn == "") {
			m := obj["httpsHealthCheck"].(map[string]interface{})
			if m["port"] == nil {
				m["port"] = 443
			}
		}
		obj["type"] = "HTTPS"
		return obj, nil
	}
	if _, ok := d.GetOk("http2_health_check"); ok {
		hc := d.Get("http2_health_check").([]interface{})[0]
		ps := hc.(map[string]interface{})["port_specification"]
		pn := hc.(map[string]interface{})["port_name"]

		if ps == "USE_FIXED_PORT" || (ps == "" && pn == "") {
			m := obj["http2HealthCheck"].(map[string]interface{})
			if m["port"] == nil {
				m["port"] = 443
			}
		}
		obj["type"] = "HTTP2"
		return obj, nil
	}
	if _, ok := d.GetOk("tcp_health_check"); ok {
		hc := d.Get("tcp_health_check").([]interface{})[0]
		ps := hc.(map[string]interface{})["port_specification"]
		pn := hc.(map[string]interface{})["port_name"]

		if ps == "USE_FIXED_PORT" || (ps == "" && pn == "") {
			m := obj["tcpHealthCheck"].(map[string]interface{})
			if m["port"] == nil {
				m["port"] = 80
			}
		}
		obj["type"] = "TCP"
		return obj, nil
	}
	if _, ok := d.GetOk("ssl_health_check"); ok {
		hc := d.Get("ssl_health_check").([]interface{})[0]
		ps := hc.(map[string]interface{})["port_specification"]
		pn := hc.(map[string]interface{})["port_name"]

		if ps == "USE_FIXED_PORT" || (ps == "" && pn == "") {
			m := obj["sslHealthCheck"].(map[string]interface{})
			if m["port"] == nil {
				m["port"] = 443
			}
		}
		obj["type"] = "SSL"
		return obj, nil
	}

	if _, ok := d.GetOk("grpc_health_check"); ok {
		hc := d.Get("grpc_health_check").([]interface{})[0]
		ps := hc.(map[string]interface{})["port_specification"]
		pn := hc.(map[string]interface{})["port_name"]

		if ps == "USE_FIXED_PORT" || (ps == "" && pn == "") {
			m := obj["grpcHealthCheck"].(map[string]interface{})
			if m["port"] == nil {
				return nil, fmt.Errorf("error in HealthCheck %s: `port` must be set for GRPC health checks`.", d.Get("name").(string))
			}
		}
		obj["type"] = "GRPC"
		return obj, nil
	}

	return nil, fmt.Errorf("error in HealthCheck %s: No health check block specified.", d.Get("name").(string))
}

func expandComputeHealthCheckCheckIntervalSec(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHealthCheckDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHealthCheckHealthyThreshold(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHealthCheckName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHealthCheckTimeoutSec(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHealthCheckUnhealthyThreshold(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHealthCheckHttpHealthCheck(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedHost, err := expandComputeHealthCheckHttpHealthCheckHost(original["host"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedHost); val.IsValid() && !isEmptyValue(val) {
		transformed["host"] = transformedHost
	}

	transformedRequestPath, err := expandComputeHealthCheckHttpHealthCheckRequestPath(original["request_path"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRequestPath); val.IsValid() && !isEmptyValue(val) {
		transformed["requestPath"] = transformedRequestPath
	}

	transformedResponse, err := expandComputeHealthCheckHttpHealthCheckResponse(original["response"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedResponse); val.IsValid() && !isEmptyValue(val) {
		transformed["response"] = transformedResponse
	}

	transformedPort, err := expandComputeHealthCheckHttpHealthCheckPort(original["port"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPort); val.IsValid() && !isEmptyValue(val) {
		transformed["port"] = transformedPort
	}

	transformedPortName, err := expandComputeHealthCheckHttpHealthCheckPortName(original["port_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPortName); val.IsValid() && !isEmptyValue(val) {
		transformed["portName"] = transformedPortName
	}

	transformedProxyHeader, err := expandComputeHealthCheckHttpHealthCheckProxyHeader(original["proxy_header"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedProxyHeader); val.IsValid() && !isEmptyValue(val) {
		transformed["proxyHeader"] = transformedProxyHeader
	}

	transformedPortSpecification, err := expandComputeHealthCheckHttpHealthCheckPortSpecification(original["port_specification"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPortSpecification); val.IsValid() && !isEmptyValue(val) {
		transformed["portSpecification"] = transformedPortSpecification
	}

	return transformed, nil
}

func expandComputeHealthCheckHttpHealthCheckHost(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHealthCheckHttpHealthCheckRequestPath(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHealthCheckHttpHealthCheckResponse(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHealthCheckHttpHealthCheckPort(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHealthCheckHttpHealthCheckPortName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHealthCheckHttpHealthCheckProxyHeader(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHealthCheckHttpHealthCheckPortSpecification(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHealthCheckHttpsHealthCheck(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedHost, err := expandComputeHealthCheckHttpsHealthCheckHost(original["host"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedHost); val.IsValid() && !isEmptyValue(val) {
		transformed["host"] = transformedHost
	}

	transformedRequestPath, err := expandComputeHealthCheckHttpsHealthCheckRequestPath(original["request_path"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRequestPath); val.IsValid() && !isEmptyValue(val) {
		transformed["requestPath"] = transformedRequestPath
	}

	transformedResponse, err := expandComputeHealthCheckHttpsHealthCheckResponse(original["response"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedResponse); val.IsValid() && !isEmptyValue(val) {
		transformed["response"] = transformedResponse
	}

	transformedPort, err := expandComputeHealthCheckHttpsHealthCheckPort(original["port"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPort); val.IsValid() && !isEmptyValue(val) {
		transformed["port"] = transformedPort
	}

	transformedPortName, err := expandComputeHealthCheckHttpsHealthCheckPortName(original["port_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPortName); val.IsValid() && !isEmptyValue(val) {
		transformed["portName"] = transformedPortName
	}

	transformedProxyHeader, err := expandComputeHealthCheckHttpsHealthCheckProxyHeader(original["proxy_header"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedProxyHeader); val.IsValid() && !isEmptyValue(val) {
		transformed["proxyHeader"] = transformedProxyHeader
	}

	transformedPortSpecification, err := expandComputeHealthCheckHttpsHealthCheckPortSpecification(original["port_specification"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPortSpecification); val.IsValid() && !isEmptyValue(val) {
		transformed["portSpecification"] = transformedPortSpecification
	}

	return transformed, nil
}

func expandComputeHealthCheckHttpsHealthCheckHost(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHealthCheckHttpsHealthCheckRequestPath(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHealthCheckHttpsHealthCheckResponse(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHealthCheckHttpsHealthCheckPort(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHealthCheckHttpsHealthCheckPortName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHealthCheckHttpsHealthCheckProxyHeader(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHealthCheckHttpsHealthCheckPortSpecification(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHealthCheckTcpHealthCheck(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedRequest, err := expandComputeHealthCheckTcpHealthCheckRequest(original["request"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRequest); val.IsValid() && !isEmptyValue(val) {
		transformed["request"] = transformedRequest
	}

	transformedResponse, err := expandComputeHealthCheckTcpHealthCheckResponse(original["response"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedResponse); val.IsValid() && !isEmptyValue(val) {
		transformed["response"] = transformedResponse
	}

	transformedPort, err := expandComputeHealthCheckTcpHealthCheckPort(original["port"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPort); val.IsValid() && !isEmptyValue(val) {
		transformed["port"] = transformedPort
	}

	transformedPortName, err := expandComputeHealthCheckTcpHealthCheckPortName(original["port_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPortName); val.IsValid() && !isEmptyValue(val) {
		transformed["portName"] = transformedPortName
	}

	transformedProxyHeader, err := expandComputeHealthCheckTcpHealthCheckProxyHeader(original["proxy_header"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedProxyHeader); val.IsValid() && !isEmptyValue(val) {
		transformed["proxyHeader"] = transformedProxyHeader
	}

	transformedPortSpecification, err := expandComputeHealthCheckTcpHealthCheckPortSpecification(original["port_specification"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPortSpecification); val.IsValid() && !isEmptyValue(val) {
		transformed["portSpecification"] = transformedPortSpecification
	}

	return transformed, nil
}

func expandComputeHealthCheckTcpHealthCheckRequest(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHealthCheckTcpHealthCheckResponse(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHealthCheckTcpHealthCheckPort(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHealthCheckTcpHealthCheckPortName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHealthCheckTcpHealthCheckProxyHeader(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHealthCheckTcpHealthCheckPortSpecification(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHealthCheckSslHealthCheck(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedRequest, err := expandComputeHealthCheckSslHealthCheckRequest(original["request"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRequest); val.IsValid() && !isEmptyValue(val) {
		transformed["request"] = transformedRequest
	}

	transformedResponse, err := expandComputeHealthCheckSslHealthCheckResponse(original["response"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedResponse); val.IsValid() && !isEmptyValue(val) {
		transformed["response"] = transformedResponse
	}

	transformedPort, err := expandComputeHealthCheckSslHealthCheckPort(original["port"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPort); val.IsValid() && !isEmptyValue(val) {
		transformed["port"] = transformedPort
	}

	transformedPortName, err := expandComputeHealthCheckSslHealthCheckPortName(original["port_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPortName); val.IsValid() && !isEmptyValue(val) {
		transformed["portName"] = transformedPortName
	}

	transformedProxyHeader, err := expandComputeHealthCheckSslHealthCheckProxyHeader(original["proxy_header"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedProxyHeader); val.IsValid() && !isEmptyValue(val) {
		transformed["proxyHeader"] = transformedProxyHeader
	}

	transformedPortSpecification, err := expandComputeHealthCheckSslHealthCheckPortSpecification(original["port_specification"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPortSpecification); val.IsValid() && !isEmptyValue(val) {
		transformed["portSpecification"] = transformedPortSpecification
	}

	return transformed, nil
}

func expandComputeHealthCheckSslHealthCheckRequest(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHealthCheckSslHealthCheckResponse(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHealthCheckSslHealthCheckPort(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHealthCheckSslHealthCheckPortName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHealthCheckSslHealthCheckProxyHeader(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHealthCheckSslHealthCheckPortSpecification(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHealthCheckHttp2HealthCheck(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedHost, err := expandComputeHealthCheckHttp2HealthCheckHost(original["host"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedHost); val.IsValid() && !isEmptyValue(val) {
		transformed["host"] = transformedHost
	}

	transformedRequestPath, err := expandComputeHealthCheckHttp2HealthCheckRequestPath(original["request_path"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRequestPath); val.IsValid() && !isEmptyValue(val) {
		transformed["requestPath"] = transformedRequestPath
	}

	transformedResponse, err := expandComputeHealthCheckHttp2HealthCheckResponse(original["response"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedResponse); val.IsValid() && !isEmptyValue(val) {
		transformed["response"] = transformedResponse
	}

	transformedPort, err := expandComputeHealthCheckHttp2HealthCheckPort(original["port"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPort); val.IsValid() && !isEmptyValue(val) {
		transformed["port"] = transformedPort
	}

	transformedPortName, err := expandComputeHealthCheckHttp2HealthCheckPortName(original["port_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPortName); val.IsValid() && !isEmptyValue(val) {
		transformed["portName"] = transformedPortName
	}

	transformedProxyHeader, err := expandComputeHealthCheckHttp2HealthCheckProxyHeader(original["proxy_header"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedProxyHeader); val.IsValid() && !isEmptyValue(val) {
		transformed["proxyHeader"] = transformedProxyHeader
	}

	transformedPortSpecification, err := expandComputeHealthCheckHttp2HealthCheckPortSpecification(original["port_specification"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPortSpecification); val.IsValid() && !isEmptyValue(val) {
		transformed["portSpecification"] = transformedPortSpecification
	}

	return transformed, nil
}

func expandComputeHealthCheckHttp2HealthCheckHost(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHealthCheckHttp2HealthCheckRequestPath(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHealthCheckHttp2HealthCheckResponse(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHealthCheckHttp2HealthCheckPort(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHealthCheckHttp2HealthCheckPortName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHealthCheckHttp2HealthCheckProxyHeader(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHealthCheckHttp2HealthCheckPortSpecification(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHealthCheckGrpcHealthCheck(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedPort, err := expandComputeHealthCheckGrpcHealthCheckPort(original["port"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPort); val.IsValid() && !isEmptyValue(val) {
		transformed["port"] = transformedPort
	}

	transformedPortName, err := expandComputeHealthCheckGrpcHealthCheckPortName(original["port_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPortName); val.IsValid() && !isEmptyValue(val) {
		transformed["portName"] = transformedPortName
	}

	transformedPortSpecification, err := expandComputeHealthCheckGrpcHealthCheckPortSpecification(original["port_specification"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPortSpecification); val.IsValid() && !isEmptyValue(val) {
		transformed["portSpecification"] = transformedPortSpecification
	}

	transformedGrpcServiceName, err := expandComputeHealthCheckGrpcHealthCheckGrpcServiceName(original["grpc_service_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedGrpcServiceName); val.IsValid() && !isEmptyValue(val) {
		transformed["grpcServiceName"] = transformedGrpcServiceName
	}

	return transformed, nil
}

func expandComputeHealthCheckGrpcHealthCheckPort(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHealthCheckGrpcHealthCheckPortName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHealthCheckGrpcHealthCheckPortSpecification(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHealthCheckGrpcHealthCheckGrpcServiceName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeHealthCheckLogConfig(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedEnable, err := expandComputeHealthCheckLogConfigEnable(original["enable"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEnable); val.IsValid() && !isEmptyValue(val) {
		transformed["enable"] = transformedEnable
	}

	return transformed, nil
}

func expandComputeHealthCheckLogConfigEnable(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
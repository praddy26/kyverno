package generate

import (
	"testing"

	"github.com/go-logr/logr"
	"gotest.tools/assert"
)

func TestValidatePass(t *testing.T) {
	resource := map[string]interface{}{
		"spec": map[string]interface{}{
			"egress": map[string]interface{}{
				"port": []interface{}{
					map[string]interface{}{
						"port":     5353,
						"protocol": "UDP",
					},
					map[string]interface{}{
						"port":     5353,
						"protocol": "TCP",
					},
				},
			},
		},
	}
	pattern := map[string]interface{}{
		"spec": map[string]interface{}{
			"egress": map[string]interface{}{
				"port": []interface{}{
					map[string]interface{}{
						"port":     5353,
						"protocol": "UDP",
					},
					map[string]interface{}{
						"port":     5353,
						"protocol": "TCP",
					},
				},
			},
		},
	}

	var log logr.Logger
	_, err := ValidateResourceWithPattern(log, resource, pattern)
	assert.NilError(t, err)
}

func TestValidateFail(t *testing.T) {
	resource := map[string]interface{}{
		"spec": map[string]interface{}{
			"egress": map[string]interface{}{
				"port": []interface{}{
					map[string]interface{}{
						"port":     5353,
						"protocol": "TCP",
					},
				},
			},
		},
	}
	pattern := map[string]interface{}{
		"spec": map[string]interface{}{
			"egress": map[string]interface{}{
				"port": []interface{}{
					map[string]interface{}{
						"port":     5353,
						"protocol": "UDP",
					},
					map[string]interface{}{
						"port":     5353,
						"protocol": "TCP",
					},
				},
			},
		},
	}

	var log logr.Logger
	_, err := ValidateResourceWithPattern(log, resource, pattern)
	assert.Assert(t, err != nil)
}

func TestValidateResourceExtraMetadataFail(t *testing.T) {
	resource := map[string]interface{}{
		"metadata": map[string]interface{}{
			"labels": map[string]interface{}{
				"xyz": "abc",
				"uvw": "iot",
			},
		},
	}
	pattern := map[string]interface{}{
		"metadata": map[string]interface{}{
			"labels": map[string]interface{}{
				"uvw": "iot",
			},
		},
	}

	var log logr.Logger
	_, err := ValidateResourceWithPattern(log, resource, pattern)
	assert.Assert(t, err != nil)
}

func TestValidateResourceMetadataFail(t *testing.T) {
	resource := map[string]interface{}{
		"metadata": map[string]interface{}{
			"labels": map[string]interface{}{
				"xyz": "abcde",
			},
			"annotations": map[string]interface{}{
				"uvw": "iotsa",
			},
		},
	}
	pattern := map[string]interface{}{
		"metadata": map[string]interface{}{
			"labels": map[string]interface{}{
				"xyz": "abc",
			},
			"annotations": map[string]interface{}{
				"uvw": "iot",
			},
		},
	}

	var log logr.Logger
	_, err := ValidateResourceWithPattern(log, resource, pattern)
	assert.Assert(t, err != nil)
}

func TestValidateResourceExtraDataFail(t *testing.T) {
	resource := map[string]interface{}{
		"data": map[string]interface{}{
			"amma":   "bHV2",
			"praddy": "cHJhZGR5Cg==",
		},
	}
	pattern := map[string]interface{}{
		"data": map[string]interface{}{
			"praddy": "cHJhZGR5Cg==",
		},
	}

	var log logr.Logger
	_, err := ValidateResourceWithPattern(log, resource, pattern)
	assert.Assert(t, err != nil)
}

func TestValidateResourceDataFail(t *testing.T) {
	resource := map[string]interface{}{
		"data": map[string]interface{}{
			"praddy": "cHJhZGR5Cg",
		},
	}
	pattern := map[string]interface{}{
		"data": map[string]interface{}{
			"praddy": "cHJhZGR5Cg==",
		},
	}

	var log logr.Logger
	_, err := ValidateResourceWithPattern(log, resource, pattern)
	assert.Assert(t, err != nil)
}

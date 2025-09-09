package constants

var AWS_REGIONS = map[string]struct{}{
	"us-east-1":      {},
	"us-east-2":      {},
	"us-west-1":      {},
	"us-west-2":      {},
	"af-south-1":     {}, // Cape Town
	"ap-east-1":      {}, // Hong Kong
	"ap-south-1":     {}, // Mumbai
	"ap-northeast-1": {}, // Tokyo
	"ap-northeast-2": {}, // Seoul
	"ap-northeast-3": {}, // Osaka
	"ap-southeast-1": {}, // Singapore
	"ap-southeast-2": {}, // Sydney
	"ca-central-1":   {}, // Canada
	"eu-central-1":   {}, // Frankfurt
	"eu-west-1":      {}, // Ireland
	"eu-west-2":      {}, // London
	"eu-west-3":      {}, // Paris
	"eu-north-1":     {}, // Stockholm
	"eu-south-1":     {}, // Milan
	"me-south-1":     {}, // Bahrain
	"sa-east-1":      {}, // São Paulo
	"us-gov-east-1":  {}, // AWS GovCloud (US-East)
	"us-gov-west-1":  {}, // AWS GovCloud (US-West)
	"cn-north-1":     {}, // Beijing
	"cn-northwest-1": {}, // Ningxia
}

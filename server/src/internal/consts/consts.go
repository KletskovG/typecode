package consts

// GitHubStartDate is the date when GitHub was created
const GitHubStartDate string = "GITHUB"

// AppliedOrgs return array of urls of applied orgs
func AppliedOrgs () []string {
	AppliedOrgs := []string {
		"google", // Google
		"apple", // Apple
		"rambler-digital-solutions", // Rambler Group 
		"oracle", // Oracle
		"microsoft", // Microsoft
		"ibm", //IBM
		"yandex", // Yandex
		"mailru", // Mail Ru
		"adobe", // Adobe
		"netflix", 
		"docker",
		"stripe",
		"alibaba",
		"cloudflare",
		"sap",
		"shopify",
		"yelp",
		"twitter",
		"aws",
	};
	return AppliedOrgs
}

// MinCodeSize size of github file content
const MinCodeSize = 1000

// DB_SERVICE_URL url of TypeCode db service
const DB_SERVICE_URL = "http://db.typecode.kletskovg.tech/"


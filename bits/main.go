package main

import "fmt"

func main() {
	tests := []struct {
		name string
		a, b int64
	}{
		{"TenantListRead", TenantListRead, iTenantListRead},
		{"TenantListWrite", TenantListWrite, iTenantListWrite},
		{"TenantListNone", TenantListNone, iTenantListNone},
		{"ApplicationListRead", ApplicationListRead, iApplicationListRead},
		{"ApplicationListWrite", ApplicationListWrite, iApplicationListWrite},
		{"ApplicationListNone", ApplicationListNone, iApplicationListNone},
		{"ApplicationDetailRead", ApplicationDetailRead, iApplicationDetailRead},
		{"ApplicationDetailWrite", ApplicationDetailWrite, iApplicationDetailWrite},
		{"ApplicationDetailNone", ApplicationDetailNone, iApplicationDetailNone},
		{"ProductListRead", ProductListRead, iProductListRead},
		{"ProductListWrite", ProductListWrite, iProductListWrite},
		{"ProductListNone", ProductListNone, iProductListNone},
		{"ProductDetailRead", ProductDetailRead, iProductDetailRead},
		{"ProductDetailWrite", ProductDetailWrite, iProductDetailWrite},
		{"ProductDetailNone", ProductDetailNone, iProductDetailNone},
		{"PortfolioAlertDefinitionsListRead", PortfolioAlertDefinitionsListRead, iPortfolioAlertDefinitionsListRead},
		{"PortfolioAlertDefinitionsListWrite", PortfolioAlertDefinitionsListWrite, iPortfolioAlertDefinitionsListWrite},
		{"PortfolioAlertDefinitionsListNone", PortfolioAlertDefinitionsListNone, iPortfolioAlertDefinitionsListNone},
		{"PortfolioAlertDefinitionsDetailRead", PortfolioAlertDefinitionsDetailRead, iPortfolioAlertDefinitionsDetailRead},
		{"PortfolioAlertDefinitionsDetailWrite", PortfolioAlertDefinitionsDetailWrite, iPortfolioAlertDefinitionsDetailWrite},
		{"PortfolioAlertDefinitionsDetailNone", PortfolioAlertDefinitionsDetailNone, iPortfolioAlertDefinitionsDetailNone},
		{"TemplatesListRead", TemplatesListRead, iTemplatesListRead},
		{"TemplatesListWrite", TemplatesListWrite, iTemplatesListWrite},
		{"TemplatesListNone", TemplatesListNone, iTemplatesListNone},
		{"TemplatesDetailRead", TemplatesDetailRead, iTemplatesDetailRead},
		{"TemplatesDetailWrite", TemplatesDetailWrite, iTemplatesDetailWrite},
		{"TemplatesDetailNone", TemplatesDetailNone, iTemplatesDetailNone},
		{"SysadminListRead", SysadminListRead, iSysadminListRead},
		{"SysadminListWrite", SysadminListWrite, iSysadminListWrite},
		{"SysadminListNone", SysadminListNone, iSysadminListNone},
		{"SysadminDetailRead", SysadminDetailRead, iSysadminDetailRead},
		{"SysadminDetailWrite", SysadminDetailWrite, iSysadminDetailWrite},
		{"SysadminDetailNone", SysadminDetailNone, iSysadminDetailNone},
		{"EventsListRead", EventsListRead, iEventsListRead},
		{"EventsListNone", EventsListNone, iEventsListNone},
		{"SitesListRead", SitesListRead, iSitesListRead},
		{"SitesListNone", SitesListNone, iSitesListNone},
		{"ReportsBillingRead", ReportsBillingRead, iReportsBillingRead},
		{"ReportsBillingNone", ReportsBillingNone, iReportsBillingNone},
		{"ReportsDownloadsRead", ReportsDownloadsRead, iReportsDownloadsRead},
		{"ReportsDownloadsNone", ReportsDownloadsNone, iReportsDownloadsNone},
		{"ChangePasswordWrite", ChangePasswordWrite, iChangePasswordWrite},
		{"TenantDetailRead", TenantDetailRead, iTenantDetailRead},
		{"TenantDetailWrite", TenantDetailWrite, iTenantDetailWrite},
		{"TenantDetailNone", TenantDetailNone, iTenantDetailNone},
		{"TenantConsumerListRead", TenantConsumerListRead, iTenantConsumerListRead},
		{"TenantConsumerListWrite", TenantConsumerListWrite, iTenantConsumerListWrite},
		{"TenantConsumerListNone", TenantConsumerListNone, iTenantConsumerListNone},
		{"TenantConsumerDetailRead", TenantConsumerDetailRead, iTenantConsumerDetailRead},
		{"TenantConsumerDetailWrite", TenantConsumerDetailWrite, iTenantConsumerDetailWrite},
		{"TenantConsumerDetailNone", TenantConsumerDetailNone, iTenantConsumerDetailNone},
		{"TenantAlertDefinitionsListRead", TenantAlertDefinitionsListRead, iTenantAlertDefinitionsListRead},
		{"TenantAlertDefinitionsListWrite", TenantAlertDefinitionsListWrite, iTenantAlertDefinitionsListWrite},
		{"TenantAlertDefinitionsListNone", TenantAlertDefinitionsListNone, iTenantAlertDefinitionsListNone},
		{"TenantAlertDefinitionsDetailRead", TenantAlertDefinitionsDetailRead, iTenantAlertDefinitionsDetailRead},
		{"TenantAlertDefinitionsDetailWrite", TenantAlertDefinitionsDetailWrite, iTenantAlertDefinitionsDetailWrite},
		{"TenantAlertDefinitionsDetailNone", TenantAlertDefinitionsDetailNone, iTenantAlertDefinitionsDetailNone},
		{"TenantPortfoliosListRead", TenantPortfoliosListRead, iTenantPortfoliosListRead},
		{"TenantPortfoliosListWrite", TenantPortfoliosListWrite, iTenantPortfoliosListWrite},
		{"TenantPortfoliosListNone", TenantPortfoliosListNone, iTenantPortfoliosListNone},
		{"TenantPortfoliosReportsListRead", TenantPortfoliosReportsListRead, iTenantPortfoliosReportsListRead},
		{"TenantPortfoliosReportsListNone", TenantPortfoliosReportsListNone, iTenantPortfoliosReportsListNone},
		/*{"TenantPortfoliosDetailRead", TenantPortfoliosDetailRead, iTenantPortfoliosDetailRead},
		{"TenantPortfoliosDetailWrite", TenantPortfoliosDetailWrite, iTenantPortfoliosDetailWrite},
		{"TenantPortfoliosDetailNone", TenantPortfoliosDetailNone, iTenantPortfoliosDetailNone},
		{"TenantPortfoliosDetailOccurrencesRead", TenantPortfoliosDetailOccurrencesRead, iTenantPortfoliosDetailOccurrencesRead},
		{"TenantPortfoliosDetailOccurrencesNone", TenantPortfoliosDetailOccurrencesNone, iTenantPortfoliosDetailOccurrencesNone},
		{"TenantPortfoliosReportsRead", TenantPortfoliosReportsRead, iTenantPortfoliosReportsRead},
		{"TenantPortfoliosReportsNone", TenantPortfoliosReportsNone, iTenantPortfoliosReportsNone},
		{"TenantPortfoliosDetailAlertsRead", TenantPortfoliosDetailAlertsRead, iTenantPortfoliosDetailAlertsRead},
		{"TenantPortfoliosDetailAlertWrite", TenantPortfoliosDetailAlertWrite, iTenantPortfoliosDetailAlertWrite},
		{"TenantPortfoliosDetailAlertsNone", TenantPortfoliosDetailAlertsNone, iTenantPortfoliosDetailAlertsNone},
		{"TenantPortfoliosDetailConsumersRead", TenantPortfoliosDetailConsumersRead, iTenantPortfoliosDetailConsumersRead},
		{"TenantPortfoliosDetailConsumersWrite", TenantPortfoliosDetailConsumersWrite, iTenantPortfoliosDetailConsumersWrite},
		{"TenantPortfoliosDetailConsumersNone", TenantPortfoliosDetailConsumersNone, iTenantPortfoliosDetailConsumersNone},
		{"TenantAlertOccurrencesListRead", TenantAlertOccurrencesListRead, iTenantAlertOccurrencesListRead},
		{"TenantAlertOccurrencesListNone", TenantAlertOccurrencesListNone, iTenantAlertOccurrencesListNone},
		{"TenantWorkflowsListRead", TenantWorkflowsListRead, iTenantWorkflowsListRead},
		{"TenantWorkflowsListNone", TenantWorkflowsListNone, iTenantWorkflowsListNone},
		{"TenantWorkflowsDetailRead", TenantWorkflowsDetailRead, iTenantWorkflowsDetailRead},
		{"TenantWorkflowsDetailNone", TenantWorkflowsDetailNone, iTenantWorkflowsDetailNone},
		{"TenantOperatorsListRead", TenantOperatorsListRead, iTenantOperatorsListRead},
		{"TenantOperatorsListWrite", TenantOperatorsListWrite, iTenantOperatorsListWrite},
		{"TenantOperatorsListNone", TenantOperatorsListNone, iTenantOperatorsListNone},
		{"TenantOperatorsDetailRead", TenantOperatorsDetailRead, iTenantOperatorsDetailRead},
		{"TenantOperatorsDetailWrite", TenantOperatorsDetailWrite, iTenantOperatorsDetailWrite},
		{"TenantOperatorsDetailNone", TenantOperatorsDetailNone, iTenantOperatorsDetailNone},
		{"TenantOrdersListRead", TenantOrdersListRead, iTenantOrdersListRead},
		{"TenantOrdersListWrite", TenantOrdersListWrite, iTenantOrdersListWrite},
		{"TenantOrdersListNone", TenantOrdersListNone, iTenantOrdersListNone},
		{"TenantOrdersDetailRead", TenantOrdersDetailRead, iTenantOrdersDetailRead},
		{"TenantOrdersDetailNone", TenantOrdersDetailNone, iTenantOrdersDetailNone},
		{"TenantEventsListRead", TenantEventsListRead, iTenantEventsListRead},
		{"TenantEventsListNone", TenantEventsListNone, iTenantEventsListNone},
		{"TenantProductsRead", TenantProductsRead, iTenantProductsRead},
		{"TenantProductsWrite", TenantProductsWrite, iTenantProductsWrite},
		{"TenantProductsNone", TenantProductsNone, iTenantProductsNone},*/
	}

	var different int
	for _, t := range tests {
		fmt.Printf("%40s | %064b | %d | %t\n", t.name, t.b, t.b, t.a == t.b)
		if t.a != t.b {
			different++
		}
	}

	fmt.Println("Different:", different)
}

// ...
const (
	TenantListRead                       = 2
	TenantListWrite                      = 4
	TenantListNone                       = 8
	ApplicationListRead                  = 16
	ApplicationListWrite                 = 32
	ApplicationListNone                  = 64
	ApplicationDetailRead                = 128
	ApplicationDetailWrite               = 256
	ApplicationDetailNone                = 512
	ProductListRead                      = 1024
	ProductListWrite                     = 2048
	ProductListNone                      = 4096
	ProductDetailRead                    = 8192
	ProductDetailWrite                   = 16384
	ProductDetailNone                    = 32768
	PortfolioAlertDefinitionsListRead    = 65536
	PortfolioAlertDefinitionsListWrite   = 131072
	PortfolioAlertDefinitionsListNone    = 262144
	PortfolioAlertDefinitionsDetailRead  = 524288
	PortfolioAlertDefinitionsDetailWrite = 1048576
	PortfolioAlertDefinitionsDetailNone  = 2097152
	TemplatesListRead                    = 4194304
	TemplatesListWrite                   = 8388608
	TemplatesListNone                    = 16777216
	TemplatesDetailRead                  = 33554432
	TemplatesDetailWrite                 = 67108864
	TemplatesDetailNone                  = 134217728
	SysadminListRead                     = 268435456
	SysadminListWrite                    = 536870912
	SysadminListNone                     = 1073741824
	SysadminDetailRead                   = 2147483648
	SysadminDetailWrite                  = 4294967296
	SysadminDetailNone                   = 8589934592
	EventsListRead                       = 17179869184
	EventsListNone                       = 34359738368
	SitesListRead                        = 68719476736
	SitesListNone                        = 137438953472
	ReportsBillingRead                   = 274877906944
	ReportsBillingNone                   = 549755813888
	ReportsDownloadsRead                 = 1099511627776
	ReportsDownloadsNone                 = 2199023255552
	ChangePasswordWrite                  = 4398046511104
	TenantDetailRead                     = 8796093022208
	TenantDetailWrite                    = 17592186044416
	TenantDetailNone                     = 35184372088832
	TenantConsumerListRead               = 70368744177664
	TenantConsumerListWrite              = 140737488355328
	TenantConsumerListNone               = 281474976710656
	TenantConsumerDetailRead             = 562949953421312
	TenantConsumerDetailWrite            = 1125899906842624
	TenantConsumerDetailNone             = 2251799813685248
	TenantAlertDefinitionsListRead       = 4503599627370496
	TenantAlertDefinitionsListWrite      = 9007199254740990
	TenantAlertDefinitionsListNone       = 18014398509482000
	TenantAlertDefinitionsDetailRead     = 36028797018964000
	TenantAlertDefinitionsDetailWrite    = 72057594037927900
	TenantAlertDefinitionsDetailNone     = 144115188075856000
	TenantPortfoliosListRead             = 288230376151712000
	TenantPortfoliosListWrite            = 576460752303423000
	TenantPortfoliosListNone             = 1152921504606850000
	TenantPortfoliosReportsListRead      = 2305843009213690000
	TenantPortfoliosReportsListNone      = 4611686018427390000
	/*TenantPortfoliosDetailRead            = 9223372036854780000
	TenantPortfoliosDetailWrite           = 18446744073709600000
	TenantPortfoliosDetailNone            = 36893488147419100000
	TenantPortfoliosDetailOccurrencesRead = 73786976294838200000
	TenantPortfoliosDetailOccurrencesNone = 147573952589676000000
	TenantPortfoliosReportsRead           = 295147905179353000000
	TenantPortfoliosReportsNone           = 590295810358706000000
	TenantPortfoliosDetailAlertsRead      = 1180591620717410000000
	TenantPortfoliosDetailAlertWrite      = 2361183241434820000000
	TenantPortfoliosDetailAlertsNone      = 4722366482869650000000
	TenantPortfoliosDetailConsumersRead   = 9444732965739290000000
	TenantPortfoliosDetailConsumersWrite  = 18889465931478600000000
	TenantPortfoliosDetailConsumersNone   = 37778931862957200000000
	TenantAlertOccurrencesListRead        = 75557863725914300000000
	TenantAlertOccurrencesListNone        = 151115727451829000000000
	TenantWorkflowsListRead               = 302231454903657000000000
	TenantWorkflowsListNone               = 604462909807315000000000
	TenantWorkflowsDetailRead             = 1208925819614630000000000
	TenantWorkflowsDetailNone             = 2417851639229260000000000
	TenantOperatorsListRead               = 4835703278458520000000000
	TenantOperatorsListWrite              = 9671406556917030000000000
	TenantOperatorsListNone               = 19342813113834100000000000
	TenantOperatorsDetailRead             = 38685626227668100000000000
	TenantOperatorsDetailWrite            = 77371252455336300000000000
	TenantOperatorsDetailNone             = 154742504910673000000000000
	TenantOrdersListRead                  = 309485009821345000000000000
	TenantOrdersListWrite                 = 618970019642690000000000000
	TenantOrdersListNone                  = 1237940039285380000000000000
	TenantOrdersDetailRead                = 2475880078570760000000000000
	TenantOrdersDetailNone                = 4951760157141520000000000000
	TenantEventsListRead                  = 9903520314283040000000000000
	TenantEventsListNone                  = 19807040628566100000000000000
	TenantProductsRead                    = 39614081257132200000000000000
	TenantProductsWrite                   = 79228162514264300000000000000
	TenantProductsNone                    = 158456325028529000000000000000*/
)

// ...
const (
	iTenantListRead = 1 << iota
	iTenantListWrite
	iTenantListNone
	iApplicationListRead
	iApplicationListWrite
	iApplicationListNone
	iApplicationDetailRead
	iApplicationDetailWrite
	iApplicationDetailNone
	iProductListRead
	iProductListWrite
	iProductListNone
	iProductDetailRead
	iProductDetailWrite
	iProductDetailNone
	iPortfolioAlertDefinitionsListRead
	iPortfolioAlertDefinitionsListWrite
	iPortfolioAlertDefinitionsListNone
	iPortfolioAlertDefinitionsDetailRead
	iPortfolioAlertDefinitionsDetailWrite
	iPortfolioAlertDefinitionsDetailNone
	iTemplatesListRead
	iTemplatesListWrite
	iTemplatesListNone
	iTemplatesDetailRead
	iTemplatesDetailWrite
	iTemplatesDetailNone
	iSysadminListRead
	iSysadminListWrite
	iSysadminListNone
	iSysadminDetailRead
	iSysadminDetailWrite
	iSysadminDetailNone
	iEventsListRead
	iEventsListNone
	iSitesListRead
	iSitesListNone
	iReportsBillingRead
	iReportsBillingNone
	iReportsDownloadsRead
	iReportsDownloadsNone
	iChangePasswordWrite
	iTenantDetailRead
	iTenantDetailWrite
	iTenantDetailNone
	iTenantConsumerListRead
	iTenantConsumerListWrite
	iTenantConsumerListNone
	iTenantConsumerDetailRead
	iTenantConsumerDetailWrite
	iTenantConsumerDetailNone
	iTenantAlertDefinitionsListRead
	iTenantAlertDefinitionsListWrite
	iTenantAlertDefinitionsListNone
	iTenantAlertDefinitionsDetailRead
	iTenantAlertDefinitionsDetailWrite
	iTenantAlertDefinitionsDetailNone
	iTenantPortfoliosListRead
	iTenantPortfoliosListWrite
	iTenantPortfoliosListNone
	iTenantPortfoliosReportsListRead
	iTenantPortfoliosReportsListNone
	/*iTenantPortfoliosDetailRead
	iTenantPortfoliosDetailWrite
	iTenantPortfoliosDetailNone
	iTenantPortfoliosDetailOccurrencesRead
	iTenantPortfoliosDetailOccurrencesNone
	iTenantPortfoliosReportsRead
	iTenantPortfoliosReportsNone
	iTenantPortfoliosDetailAlertsRead
	iTenantPortfoliosDetailAlertWrite
	iTenantPortfoliosDetailAlertsNone
	iTenantPortfoliosDetailConsumersRead
	iTenantPortfoliosDetailConsumersWrite
	iTenantPortfoliosDetailConsumersNone
	iTenantAlertOccurrencesListRead
	iTenantAlertOccurrencesListNone
	iTenantWorkflowsListRead
	iTenantWorkflowsListNone
	iTenantWorkflowsDetailRead
	iTenantWorkflowsDetailNone
	iTenantOperatorsListRead
	iTenantOperatorsListWrite
	iTenantOperatorsListNone
	iTenantOperatorsDetailRead
	iTenantOperatorsDetailWrite
	iTenantOperatorsDetailNone
	iTenantOrdersListRead
	iTenantOrdersListWrite
	iTenantOrdersListNone
	iTenantOrdersDetailRead
	iTenantOrdersDetailNone
	iTenantEventsListRead
	iTenantEventsListNone
	iTenantProductsRead
	iTenantProductsWrite
	iTenantProductsNone*/
)

package ratsit

// Person is a struct that describes the detailed person information returned by Ratsit
type Person struct {
	ResponseCode        string `json:"responseCode"`
	ResponseMessage     string `json:"responseMessage"`
	SsnStatus           string `json:"ssnStatus"`
	RegisteredAtAddress struct {
		Persons   []Person `json:"persons"`
		Companies []struct {
			OrgNr       string `json:"orgNr"`
			CompanyName string `json:"companyName"`
			Street      string `json:"street"`
			City        string `json:"city"`
			CompanyCode string `json:"companyCode"`
		} `json:"companies"`
		CompanyPhones []struct {
			OrgNr       string `json:"orgNr"`
			CompanyName string `json:"companyName"`
			City        string `json:"city"`
			PhoneNumber string `json:"phoneNumber"`
		} `json:"companyPhones"`
		PersonPhones []struct {
			Ssn         string `json:"ssn"`
			FirstName   string `json:"firstName"`
			LastName    string `json:"lastName"`
			City        string `json:"city"`
			PhoneNumber string `json:"phoneNumber"`
		} `json:"personPhones"`
	} `json:"registeredAtAddress"`
	Name struct {
		FirstName  string `json:"firstName"`
		GivenName  string `json:"givenName"`
		SurName    string `json:"surName"`
		MiddleName string `json:"middleName"`
		LastName   string `json:"lastName"`
		Age        int    `json:"age"`
		Gender     string `json:"gender"`
		IsMarried  bool   `json:"isMarried"`
		City       string `json:"city"`
	} `json:"name"`
	Basic struct {
		FirstName  string `json:"firstName"`
		GivenName  string `json:"givenName"`
		SurName    string `json:"surName"`
		MiddleName string `json:"middleName"`
		LastName   string `json:"lastName"`
		Co         string `json:"co"`
		Street     string `json:"street"`
		ZipCode    string `json:"zipCode"`
		City       string `json:"city"`
	} `json:"basic"`
	Extended struct {
		Municipality                string `json:"municipality"`
		Parish                      string `json:"parish"`
		District                    string `json:"district"`
		County                      string `json:"county"`
		NextBirthDay                string `json:"nextBirthDay"`
		DateOfBirth                 string `json:"dateOfBirth"`
		Age                         int    `json:"age"`
		Gender                      string `json:"gender"`
		IsMarried                   bool   `json:"isMarried"`
		SocialStatusWith            string `json:"socialStatusWith"`
		SocialStatusWithSSN         string `json:"socialStatusWithSSN"`
		NumberOfCompanyInvolvements int    `json:"numberOfCompanyInvolvements"`
	} `json:"extended"`
	CompanyInvolvements struct {
		NumberOfCompanyInvolvements []struct {
			OrgNr         string `json:"orgNr"`
			CompanyName   string `json:"companyName"`
			CompanyCode   string `json:"companyCode"`
			FunctionTexts string `json:"functionTexts"`
			Func1         string `json:"func1"`
			Func2         string `json:"func2"`
			Func3         string `json:"func3"`
			Func4         string `json:"func4"`
			Func1Text     string `json:"func1Text"`
			Func2Text     string `json:"func2Text"`
			Func3Text     string `json:"func3Text"`
			Func4Text     string `json:"func4Text"`
		} `json:"numberOfCompanyInvolvements"`
	} `json:"companyInvolvements"`
	PhoneNumbers struct {
		PhoneNumbers []string `json:"phoneNumbers"`
	} `json:"phoneNumbers"`
	Ssn struct {
		Ssn string `json:"ssn"`
	} `json:"ssn"`
	SpecialAddress struct {
		SpecialAddress struct {
			Co      string `json:"co"`
			Street  string `json:"street"`
			ZipCode string `json:"zipCode"`
			City    string `json:"city"`
		} `json:"specialAddress"`
	} `json:"specialAddress"`
	Remark struct {
		Status         string `json:"status"`
		HasRemark      bool   `json:"hasRemark"`
		LastRemarkDate string `json:"lastRemarkDate"`
	} `json:"remark"`
	RemarkExtended struct {
		Status                         string `json:"status"`
		HasDebtRelief                  bool   `json:"hasDebtRelief"`
		HasSeizure                     bool   `json:"hasSeizure"`
		HasRealEstateHolding           bool   `json:"hasRealEstateHolding"`
		PaymentRemarkNumber            int    `json:"paymentRemarkNumber"`
		PaymentRemarkSum               int    `json:"paymentRemarkSum"`
		PublicRemarkNumber             int    `json:"publicRemarkNumber"`
		PublicRemarkSum                int    `json:"publicRemarkSum"`
		PublicRemarkTVlicenseNumber    int    `json:"publicRemarkTVlicenseNumber"`
		PublicRemarkTVlicenseSum       int    `json:"publicRemarkTVlicenseSum"`
		PublicRemarkParkingFeeNumber   int    `json:"publicRemarkParkingFeeNumber"`
		PublicRemarkParkingFeeSum      int    `json:"publicRemarkParkingFeeSum"`
		PublicRemarkAlimonyNumber      int    `json:"publicRemarkAlimonyNumber"`
		PublicRemarkAlimonySum         int    `json:"publicRemarkAlimonySum"`
		PublicRemarkTaxArrearsNumber   int    `json:"publicRemarkTaxArrearsNumber"`
		PublicRemarkTaxArrearsSum      int    `json:"publicRemarkTaxArrearsSum"`
		PublicRemarkStudentLoansNumber int    `json:"publicRemarkStudentLoansNumber"`
		PublicRemarkStudentLoansSum    int    `json:"publicRemarkStudentLoansSum"`
		PrivateRemarkNumber            int    `json:"privateRemarkNumber"`
		PrivateRemarkAmount            int    `json:"privateRemarkAmount"`
		LastRemarkDate                 string `json:"lastRemarkDate"`
		DebtNumber                     int    `json:"debtNumber"`
		DebtSum                        int    `json:"debtSum"`
		DebtPublicRemarkNumber         int    `json:"debtPublicRemarkNumber"`
		DebtPublicRemarkSum            int    `json:"debtPublicRemarkSum"`
		DebtPrivateRemarkNumber        int    `json:"debtPrivateRemarkNumber"`
		DebtPrivateRemarkSum           int    `json:"debtPrivateRemarkSum"`
		LastDebtDate                   string `json:"lastDebtDate"`
	} `json:"remarkExtended"`
	RemarksFiveLatest struct {
		Remarks []struct {
			Date         string `json:"date"`
			Amount       int    `json:"amount"`
			Remarks      string `json:"remarks"`
			Creditor     string `json:"creditor"`
			Government   string `json:"government"`
			TypeOfRemark string `json:"typeOfRemark"`
			StatusDate   string `json:"statusDate"`
		} `json:"remarks"`
	} `json:"remarksFiveLatest"`
	TaxInformation struct {
		TaxYear                  int    `json:"taxYear"`
		SumOfIncome              int    `json:"sumOfIncome"`
		TaxatedAcquisitiedIncome int    `json:"taxatedAcquisitiedIncome"`
		IncomeThroughService     int    `json:"incomeThroughService"`
		IncomeThroughCapital     int    `json:"incomeThroughCapital"`
		DeficitOfCapital         int    `json:"deficitOfCapital"`
		HasRealEstateHolding     bool   `json:"hasRealEstateHolding"`
		Status                   string `json:"status"`
		Events                   []struct {
			Date      string `json:"date"`
			EventText string `json:"eventText"`
		} `json:"events"`
	} `json:"taxInformation"`
	TaxInformationPreviouslyYear struct {
		SumOfIncome              int `json:"sumOfIncome"`
		TaxatedAcquisitiedIncome int `json:"taxatedAcquisitiedIncome"`
		IncomeThroughService     int `json:"incomeThroughService"`
		IncomeThroughCapital     int `json:"incomeThroughCapital"`
		DeficitOfCapital         int `json:"deficitOfCapital"`
	} `json:"taxInformationPreviouslyYear"`
	RealEstateHolding struct {
		HasRealEstateHolding bool `json:"hasRealEstateHolding"`
	} `json:"realEstateHolding"`
}

type SearchResults struct {
	ExtendedResult struct {
		RecordsReturned   int           `json:"recordsReturned"`
		TotalRecordsFound int           `json:"totalRecordsFound"`
		Records           []PersonBasic `json:"records"`
	} `json:"extendedResult"`
}

type PersonBasic struct {
	Ssn          string `json:"ssn"`
	FirstName    string `json:"firstName"`
	GivenName    string `json:"givenName"`
	LastName     string `json:"lastName"`
	Street       string `json:"street"`
	ZipCode      string `json:"zipCode"`
	City         string `json:"city"`
	Municipality string `json:"municipality"`
	Age          int    `json:"age"`
}

type CompanySearchResults struct {
	BasicResult struct {
		RecordsReturned   int            `json:"recordsReturned"`
		TotalRecordsFound int            `json:"totalRecordsFound"`
		Records           []CompanyBasic `json:"records"`
	} `json:"basicResult"`
}

type CompanyBasic struct {
	OrgNr       string `json:"orgNr"`
	CompanyName string `json:"companyName"`
	CompanyCode string `json:"companyCode"`
	Co          string `json:"co"`
	Street      string `json:"street"`
	ZipCode     string `json:"zipCode"`
	City        string `json:"city"`
}

// CompanyInformationResponse defines the response from GET /api/v1/companyinformation
type CompanyInformationResponse struct {
	ResponseCode    string `json:"responseCode"`
	ResponseMessage string `json:"responseMessage"`
	Name            struct {
		CompanyName string `json:"companyName"`
		City        string `json:"city"`
	} `json:"name"`
	Basic struct {
		CompanyName        string `json:"companyName"`
		OrganizationNumber string `json:"organizationNumber"`
		Co                 string `json:"co"`
		Street             string `json:"street"`
		ZipCode            string `json:"zipCode"`
		City               string `json:"city"`
		CompanyCode        string `json:"companyCode"`
		CompanyCodeText    string `json:"companyCodeText"`
		RegistrationDate   string `json:"registrationDate"`
	} `json:"basic"`
	Extended struct {
		CorporationFoundedDate    string `json:"corporationFoundedDate"`
		CorporationRegisteredDate string `json:"corporationRegisteredDate"`
		RegistrationDate          string `json:"registrationDate"`
		CompanyRegulationDate     string `json:"companyRegulationDate"`
		CorporateTax              string `json:"corporateTax"`
		EmployerStatus            string `json:"employerStatus"`
		VatStatus                 string `json:"vatStatus"`
		RegisteredSNI             []struct {
			Code string `json:"code"`
			Text string `json:"text"`
		} `json:"registeredSNI"`
		RegisteredCorporateDescription string `json:"registeredCorporateDescription"`
		Status                         string `json:"status"`
	} `json:"extended"`
	PhoneNumbers struct {
		PhoneNumbers []string `json:"phoneNumbers"`
	} `json:"phoneNumbers"`
	EconomicExtended struct {
		Slutkod                                           string `json:"slutkod"`
		FinancialPeriod                                   string `json:"financialPeriod"`
		FinancialNetLength                                string `json:"financialNetLength"`
		RedType                                           string `json:"redType"`
		AccountantComment                                 string `json:"accountantComment"`
		AccountantCommentText                             string `json:"accountantCommentText"`
		Bristkod                                          string `json:"bristkod"`
		BristkodText                                      string `json:"bristkodText"`
		Utdelningskod                                     string `json:"utdelningskod"`
		Utdelning                                         int    `json:"utdelning"`
		NetSalesAndProfitSettlementInvoices               int    `json:"netSalesAndProfitSettlementInvoices"`
		CostOfGoodsSold                                   int    `json:"costOfGoodsSold"`
		GrossProfitOrGrossLoss                            int    `json:"grossProfitOrGrossLoss"`
		CostOfSales                                       int    `json:"costOfSales"`
		AdministrativeCosts                               int    `json:"administrativeCosts"`
		CostOfRD                                          int    `json:"costOfRD"`
		ItemsAffectingComparabilityPostRr2                int    `json:"itemsAffectingComparabilityPostRr2"`
		OtherOperatingIncomeRr2                           int    `json:"otherOperatingIncomeRr2"`
		OtherOperatingExpenses                            int    `json:"otherOperatingExpenses"`
		ChangeInInventory                                 int    `json:"changeInInventory"`
		WorkPerformedOwnAccount                           int    `json:"workPerformedOwnAccount"`
		OtherOperatingIncomeRr1                           int    `json:"otherOperatingIncomeRr1"`
		RawMaterialsAndConsumables                        int    `json:"rawMaterialsAndConsumables"`
		Merchandise                                       int    `json:"merchandise"`
		OtherExternalExpenses                             int    `json:"otherExternalExpenses"`
		PersonnelCosts                                    int    `json:"personnelCosts"`
		Depreciation                                      int    `json:"depreciation"`
		ItemsAffectingComparabilitypostRr1                int    `json:"itemsAffectingComparabilitypostRr1"`
		OtherOperatingExpensesRr1                         int    `json:"otherOperatingExpensesRr1"`
		OperatingProfit                                   int    `json:"operatingProfit"`
		ResultatFranAndelarIKoncernOchIntresseforetag     int    `json:"resultatFranAndelarIKoncernOchIntresseforetag"`
		InterestIncomeFromGroupCompanies                  int    `json:"interestIncomeFromGroupCompanies"`
		ExternalInterestIncome                            int    `json:"externalInterestIncome"`
		OtherFinancialInterestIncome                      int    `json:"otherFinancialInterestIncome"`
		InterestExpensesToGroupCompanies                  int    `json:"interestExpensesToGroupCompanies"`
		ExternalInterestExpenses                          int    `json:"externalInterestExpenses"`
		OtherFinancialExpenses                            int    `json:"otherFinancialExpenses"`
		NonRecurringFinancialItems                        int    `json:"nonRecurringFinancialItems"`
		ProfitAfterFinancialIncomeAndExpenses             int    `json:"profitAfterFinancialIncomeAndExpenses"`
		ExtraordinaryIncome                               int    `json:"extraordinaryIncome"`
		ExtraordinaryExpenses                             int    `json:"extraordinaryExpenses"`
		GroupContribution                                 int    `json:"groupContribution"`
		ShareholdersContribution                          int    `json:"shareholdersContribution"`
		Appropriations                                    int    `json:"appropriations"`
		Tax                                               int    `json:"tax"`
		MinoritetsintrSamtVinstForlIDotterbolForForvarv   int    `json:"minoritetsintrSamtVinstForlIDotterbolForForvarv"`
		NetIncome                                         int    `json:"netIncome"`
		SubscribedUnpaidCapital                           int    `json:"subscribedUnpaidCapital"`
		CapitalisedExpenditureRD                          int    `json:"capitalisedExpenditureRD"`
		PatentsAndLicenses                                int    `json:"patentsAndLicenses"`
		Goodwill                                          int    `json:"goodwill"`
		OtherIntangibleAssets                             int    `json:"otherIntangibleAssets"`
		TotalIntangibleAssets                             int    `json:"totalIntangibleAssets"`
		LandAndBuildings                                  int    `json:"landAndBuildings"`
		Machines                                          int    `json:"machines"`
		Inventory                                         int    `json:"inventory"`
		MachinesAndInventory                              int    `json:"machinesAndInventory"`
		OtherTangibleFixedAssetsNoneDepreciable           int    `json:"otherTangibleFixedAssetsNoneDepreciable"`
		OtherTangibleFixedAssetsDepreciable               int    `json:"otherTangibleFixedAssetsDepreciable"`
		SumTangibleAssets                                 int    `json:"sumTangibleAssets"`
		InvestmentsInSubsidiariesAndAssociates            int    `json:"investmentsInSubsidiariesAndAssociates"`
		CurrentReceivablesFromGroupAndAssociatedCompanies int    `json:"currentReceivablesFromGroupAndAssociatedCompanies"`
		LoansToShareholderAndRelatedParties               int    `json:"loansToShareholderAndRelatedParties"`
		OtherFinancialFixedAssets                         int    `json:"otherFinancialFixedAssets"`
		SumFinancialFixedAssets                           int    `json:"sumFinancialFixedAssets"`
		SumFixedAssets                                    int    `json:"sumFixedAssets"`
		WorkInProgressOnBehalfOfOthers                    int    `json:"workInProgressOnBehalfOfOthers"`
		OtherInventories                                  int    `json:"otherInventories"`
		TotalInventories                                  int    `json:"totalInventories"`
		AccountsReceivable                                int    `json:"accountsReceivable"`
		ReceivablesAtGroupAndAssociatedCompanies          int    `json:"receivablesAtGroupAndAssociatedCompanies"`
		OtherCurrentReceivables                           int    `json:"otherCurrentReceivables"`
		SumCurrentReceivables                             int    `json:"sumCurrentReceivables"`
		TotalShortTermInvestments                         int    `json:"totalShortTermInvestments"`
		TotalCashAndCashEquivalents                       int    `json:"totalCashAndCashEquivalents"`
		TotalCurrentAssets                                int    `json:"totalCurrentAssets"`
		TotalAssets                                       int    `json:"totalAssets"`
		ShareCapital                                      int    `json:"shareCapital"`
		SharePremiumReserve                               int    `json:"sharePremiumReserve"`
		RevaluationReserve                                int    `json:"revaluationReserve"`
		OtherRestrictedEquity                             int    `json:"otherRestrictedEquity"`
		RetainedEarnings                                  int    `json:"retainedEarnings"`
		ReceivedPaidGroupContributions                    int    `json:"receivedPaidGroupContributions"`
		ReceivedShareholderContributions                  int    `json:"receivedShareholderContributions"`
		NetIncome2                                        int    `json:"netIncome2"`
		TotalEquity                                       int    `json:"totalEquity"`
		TotalUntaxedReserves                              int    `json:"totalUntaxedReserves"`
		MinorityInterests                                 int    `json:"minorityInterests"`
		TotalProvisions                                   int    `json:"totalProvisions"`
		Bond                                              int    `json:"bond"`
		LiabilitiesToCreditInstitutionsLongTerm           int    `json:"liabilitiesToCreditInstitutionsLongTerm"`
		LiabilitiesToGroupCompaniesAndAssociatesLongTerm  int    `json:"liabilitiesToGroupCompaniesAndAssociatesLongTerm"`
		OtherLongTermLiabilities                          int    `json:"otherLongTermLiabilities"`
		TotalLongTermLiabilities                          int    `json:"totalLongTermLiabilities"`
		LiabilitiesToCreditInstitutionsShortTerm          int    `json:"liabilitiesToCreditInstitutionsShortTerm"`
		AccountsPayable                                   int    `json:"accountsPayable"`
		LiabilitiesToGroupCompaniesAndAssociatesShortTerm int    `json:"liabilitiesToGroupCompaniesAndAssociatesShortTerm"`
		OtherShortTermLiabilities                         int    `json:"otherShortTermLiabilities"`
		TotalShortTermLiabilities                         int    `json:"totalShortTermLiabilities"`
		TotalEquityAndLiabilities                         int    `json:"totalEquityAndLiabilities"`
		CompanyAssets                                     int    `json:"companyAssets"`
		RealEstateAssets                                  int    `json:"realEstateAssets"`
		OtherCollateral                                   int    `json:"otherCollateral"`
		TotalCollateral                                   int    `json:"totalCollateral"`
		ConditionalShareholderContribution                int    `json:"conditionalShareholderContribution"`
		OtherLiabilities                                  int    `json:"otherLiabilities"`
		SumLiabilities                                    int    `json:"sumLiabilities"`
		NumberOfEmployees                                 int    `json:"numberOfEmployees"`
		SalariesOfTheBoardAndCEO                          int    `json:"salariesOfTheBoardAndCEO"`
		BonusOfTheBoardAndCEO                             int    `json:"bonusOfTheBoardAndCEO"`
		SalariesToOtherEmployees                          int    `json:"salariesToOtherEmployees"`
		IncludingPerformanceBonusToOtherEmployees         int    `json:"includingPerformanceBonusToOtherEmployees"`
		SocialExpenses                                    int    `json:"socialExpenses"`
		DepreciationAndCostOfGoodsSold                    int    `json:"depreciationAndCostOfGoodsSold"`
		DepreciationInSellingExpenses                     int    `json:"depreciationInSellingExpenses"`
		DepreciationInAdminstrationExpenses               int    `json:"depreciationInAdminstrationExpenses"`
		DepreciationInRDExpenses                          int    `json:"depreciationInRDExpenses"`
		DepreciationOtherExpenses                         int    `json:"depreciationOtherExpenses"`
		UnspecifiedDepreciation                           int    `json:"unspecifiedDepreciation"`
		OverdraftFacilities                               int    `json:"overdraftFacilities"`
		BankOverdrafts                                    int    `json:"bankOverdrafts"`
		NetInterestFinance                                int    `json:"netInterestFinance"`
	} `json:"economicExtended"`
}

// HTTPErrorBody holds the body of an error returned by ratsit in the event of an HTTP Error
type HTTPErrorBody struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

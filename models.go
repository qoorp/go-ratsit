package ratsit

// Person is a struct that describes the detailed person information returned by Ratsit
type Person struct {
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

package mydata

import (
	"github.com/invopop/gobl/cbc"
	"github.com/invopop/gobl/i18n"
)

// Regime extension codes.
const (
	ExtKeyVATRate      = "gr-mydata-vat-rate"
	ExtKeyInvoiceType  = "gr-mydata-invoice-type"
	ExtKeyExemption    = "gr-mydata-exemption"
	ExtKeyIncomeCat    = "gr-mydata-income-cat"
	ExtKeyIncomeType   = "gr-mydata-income-type"
	ExtKeyPaymentMeans = "gr-mydata-payment-means"

	InvoiceTypeRetailPrefix = "11."
)

var extensions = []*cbc.KeyDefinition{
	{
		Key: ExtKeyVATRate,
		Name: i18n.String{
			i18n.EN: "VAT rate",
			i18n.EL: "Κατηγορία ΦΠΑ",
		},
		Values: []*cbc.ValueDefinition{
			{
				Value: "1",
				Name: i18n.String{
					i18n.EN: "Standard rate",
					i18n.EL: "Κανονικός συντελεστής",
				},
			},
			{
				Value: "2",
				Name: i18n.String{
					i18n.EN: "Reduced rate",
					i18n.EL: "Μειωμένος συντελεστής",
				},
			},
			{
				Value: "3",
				Name: i18n.String{
					i18n.EN: "Super-Reduced Rate",
					i18n.EL: "Υπερμειωμένος συντελεστής",
				},
			},
			{
				Value: "4",
				Name: i18n.String{
					i18n.EN: "Standard rate (Island)",
					i18n.EL: "Κανονικός συντελεστής (Νησί)",
				},
			},
			{
				Value: "5",
				Name: i18n.String{
					i18n.EN: "Reduced rate (Island)",
					i18n.EL: "Μειωμένος συντελεστής (Νησί)",
				},
			},
			{
				Value: "6",
				Name: i18n.String{
					i18n.EN: "Super-reduced rate (Island)",
					i18n.EL: "Υπερμειωμένος συντελεστής (Νησί)",
				},
			},
			{
				Value: "7",
				Name: i18n.String{
					i18n.EN: "Without VAT",
					i18n.EL: "Άνευ ΦΠΑ",
				},
			},
			{
				Value: "8",
				Name: i18n.String{
					i18n.EN: "Records without VAT (e.g. Payroll, Amortisations)",
					i18n.EL: "Εγγραφές χωρίς ΦΠΑ (πχ Μισθοδοσία, Αποσβέσεις)",
				},
			},
		},
	},
	{
		Key: ExtKeyInvoiceType,
		Name: i18n.String{
			i18n.EN: "Invoice type",
			i18n.EL: "Είδος παραστατικού",
		},
		Values: []*cbc.ValueDefinition{
			{
				Value: "1.1",
				Name: i18n.String{
					i18n.EN: "Sales Invoice",
					i18n.EL: "Τιμολόγιο Πώλησης",
				},
			},
			{
				Value: "1.2",
				Name: i18n.String{
					i18n.EN: "Sales Invoice/Intra-community Supplies",
					i18n.EL: "Τιμολόγιο Πώλησης/Ενδοκοινοτικές Παραδόσεις",
				},
			},
			{
				Value: "1.3",
				Name: i18n.String{
					i18n.EN: "Sales Invoice/Third Country Supplies",
					i18n.EL: "Τιμολόγιο Πώλησης/Παραδόσεις Τρίτων Χωρών",
				},
			},
			{
				Value: "1.4",
				Name: i18n.String{
					i18n.EN: "Sales Invoice/Sale on Behalf of Third Parties",
					i18n.EL: "Τιμολόγιο Πώλησης/Πώληση για Λογαριασμό Τρίτων",
				},
			},
			{
				Value: "1.5",
				Name: i18n.String{
					i18n.EN: "Sales Invoice/Clearance of Sales on Behalf of Third Parties – Fees from Sales on Behalf of Third Parties",
					i18n.EL: "Τιμολόγιο Πώλησης/Εκκαθάριση Πωλήσεων Τρίτων - Αμοιβή από Πωλήσεις Τρίτων",
				},
			},
			{
				Value: "1.6",
				Name: i18n.String{
					i18n.EN: "Sales Invoice/Supplemental Accounting Source Document",
					i18n.EL: "Τιμολόγιο Πώλησης/Συμπληρωματικό Παραστατικό",
				},
			},
			{
				Value: "2.1",
				Name: i18n.String{
					i18n.EN: "Service Rendered Invoice",
					i18n.EL: "Τιμολόγιο Παροχής Υπηρεσιών",
				},
			},
			{
				Value: "2.2",
				Name: i18n.String{
					i18n.EN: "Intra-community Service Rendered Invoice",
					i18n.EL: "Τιμολόγιο Παροχής/Ενδοκοινοτική Παροχή Υπηρεσιών",
				},
			},
			{
				Value: "2.3",
				Name: i18n.String{
					i18n.EN: "Third Country Service Rendered Invoice",
					i18n.EL: "Τιμολόγιο Παροχής/Παροχή Υπηρεσιών σε λήπτη Τρίτης Χώρας",
				},
			},
			{
				Value: "2.4",
				Name: i18n.String{
					i18n.EN: "Service Rendered Invoice/Supplemental Accounting Source Document",
					i18n.EL: "Τιμολόγιο Παροχής/Συμπληρωματικό Παραστατικό",
				},
			},
			{
				Value: "3.1",
				Name: i18n.String{
					i18n.EN: "Proof of Expenditure (non-liable Issuer)",
					i18n.EL: "Τίτλος Κτήσης (μη υπόχρεος Εκδότης)",
				},
			},
			{
				Value: "3.2",
				Name: i18n.String{
					i18n.EN: "Proof of Expenditure (denial of issuance by liable Issuer)",
					i18n.EL: "Τίτλος Κτήσης (άρνηση έκδοσης από υπόχρεο Εκδότη)",
				},
			},
			{
				Value: "5.1",
				Name: i18n.String{
					i18n.EN: "Credit Invoice/Associated",
					i18n.EL: "Πιστωτικό Τιμολόγιο/Συσχετιζόμενο",
				},
			},
			{
				Value: "5.2",
				Name: i18n.String{
					i18n.EN: "Credit Invoice/Non-Associated",
					i18n.EL: "Πιστωτικό Τιμολόγιο/Μη Συσχετιζόμενο",
				},
			},
			{
				Value: "6.1",
				Name: i18n.String{
					i18n.EN: "Self-Delivery Record",
					i18n.EL: "Στοιχείο Αυτοπαράδοσης",
				},
			},
			{
				Value: "6.2",
				Name: i18n.String{
					i18n.EN: "Self-Supply Record",
					i18n.EL: "Στοιχείο Ιδιοχρησιμοποίησης",
				},
			},
			{
				Value: "7.1",
				Name: i18n.String{
					i18n.EN: "Contract – Income",
					i18n.EL: "Συμβόλαιο - Έσοδο",
				},
			},
			{
				Value: "8.1",
				Name: i18n.String{
					i18n.EN: "Rents – Income",
					i18n.EL: "Ενοίκια - Έσοδο",
				},
			},
			{
				Value: "8.2",
				Name: i18n.String{
					i18n.EN: "Special Record – Accommodation Tax Collection/Payment Receipt",
					i18n.EL: "Ειδικό Στοιχείο – Απόδειξης Είσπραξης Φόρου Διαμονής",
				},
			},
			{
				Value: "11.1",
				Name: i18n.String{
					i18n.EN: "Retail Sales Receipt",
					i18n.EL: "ΑΛΠ",
				},
			},
			{
				Value: "11.2",
				Name: i18n.String{
					i18n.EN: "Service Rendered Receipt",
					i18n.EL: "ΑΠΥ",
				},
			},
			{
				Value: "11.3",
				Name: i18n.String{
					i18n.EN: "Simplified Invoice",
					i18n.EL: "Απλοποιημένο Τιμολόγιο",
				},
			},
			{
				Value: "11.4",
				Name: i18n.String{
					i18n.EN: "Retail Sales Credit Note",
					i18n.EL: "Πιστωτικό Στοιχ. Λιανικής",
				},
			},
			{
				Value: "11.5",
				Name: i18n.String{
					i18n.EN: "Retail Sales Receipt on Behalf of Third Parties",
					i18n.EL: "Απόδειξη Λιανικής Πώλησης για Λογ/σμό Τρίτων",
				},
			},
			{
				Value: "13.1",
				Name: i18n.String{
					i18n.EN: "Expenses – Domestic/Foreign Retail Transaction Purchases",
					i18n.EL: "Έξοδα - Αγορές Λιανικών Συναλλαγών ημεδαπής / αλλοδαπής",
				},
			},
			{
				Value: "13.2",
				Name: i18n.String{
					i18n.EN: "Domestic/Foreign Retail Transaction Provision",
					i18n.EL: "Παροχή Λιανικών Συναλλαγών ημεδαπής / αλλοδαπής",
				},
			},
			{
				Value: "13.3",
				Name: i18n.String{
					i18n.EN: "Shared Utility Bills",
					i18n.EL: "Κοινόχρηστα",
				},
			},
			{
				Value: "13.4",
				Name: i18n.String{
					i18n.EN: "Subscriptions",
					i18n.EL: "Συνδρομές",
				},
			},
			{
				Value: "13.30",
				Name: i18n.String{
					i18n.EN: "Self-Declared Entity Accounting Source Documents (Dynamic)",
					i18n.EL: "Παραστατικά Οντότητας ως Αναγράφονται από την ίδια (Δυναμικό)",
				},
			},
			{
				Value: "13.31",
				Name: i18n.String{
					i18n.EN: "Domestic/Foreign Retail Sales Credit Note",
					i18n.EL: "Πιστωτικό Στοιχ. Λιανικής ημεδαπής / αλλοδαπής",
				},
			},
			{
				Value: "14.1",
				Name: i18n.String{
					i18n.EN: "Invoice/Intra-community Acquisitions",
					i18n.EL: "Τιμολόγιο / Ενδοκοινοτικές Αποκτήσεις",
				},
			},
			{
				Value: "14.2",
				Name: i18n.String{
					i18n.EN: "Invoice/Third Country Acquisitions",
					i18n.EL: "Τιμολόγιο / Αποκτήσεις Τρίτων Χωρών",
				},
			},
			{
				Value: "14.3",
				Name: i18n.String{
					i18n.EN: "Invoice/Intra-community Services Receipt",
					i18n.EL: "Τιμολόγιο / Ενδοκοινοτική Λήψη Υπηρεσιών",
				},
			},
			{
				Value: "14.4",
				Name: i18n.String{
					i18n.EN: "Invoice/Third Country Services Receipt",
					i18n.EL: "Τιμολόγιο / Λήψη Υπηρεσιών Τρίτων Χωρών",
				},
			},
			{
				Value: "14.5",
				Name: i18n.String{
					i18n.EN: "EFKA",
					i18n.EL: "ΕΦΚΑ και λοιποί Ασφαλιστικοί Οργανισμοί",
				},
			},
			{
				Value: "14.30",
				Name: i18n.String{
					i18n.EN: "Self-Declared Entity Accounting Source Documents (Dynamic)",
					i18n.EL: "Παραστατικά Οντότητας ως Αναγράφονται από την ίδια (Δυναμικό)",
				},
			},
			{
				Value: "14.31",
				Name: i18n.String{
					i18n.EN: "Domestic/Foreign Credit Note",
					i18n.EL: "Πιστωτικό ημεδαπής / αλλοδαπής",
				},
			},
			{
				Value: "15.1",
				Name: i18n.String{
					i18n.EN: "Contract-Expense",
					i18n.EL: "Συμβόλαιο - Έξοδο",
				},
			},
			{
				Value: "16.1",
				Name: i18n.String{
					i18n.EN: "Rent-Expense",
					i18n.EL: "Ενοίκιο Έξοδο",
				},
			},
			{
				Value: "17.1",
				Name: i18n.String{
					i18n.EN: "Payroll",
					i18n.EL: "Μισθοδοσία",
				},
			},
			{
				Value: "17.2",
				Name: i18n.String{
					i18n.EN: "Amortisations",
					i18n.EL: "Αποσβέσεις",
				},
			},
			{
				Value: "17.3",
				Name: i18n.String{
					i18n.EN: "Other Income Adjustment/Regularisation Entries – Accounting Base",
					i18n.EL: "Λοιπές Εγγραφές Τακτοποίησης Εσόδων - Λογιστική Βάση",
				},
			},
			{
				Value: "17.4",
				Name: i18n.String{
					i18n.EN: "Other Income Adjustment/Regularisation Entries – Tax Base",
					i18n.EL: "Λοιπές Εγγραφές Τακτοποίησης Εσόδων - Φορολογική Βάση",
				},
			},
			{
				Value: "17.5",
				Name: i18n.String{
					i18n.EN: "Other Expense Adjustment/Regularisation Entries – Accounting Base",
					i18n.EL: "Λοιπές Εγγραφές Τακτοποίησης Εξόδων - Λογιστική Βάση",
				},
			},
			{
				Value: "17.6",
				Name: i18n.String{
					i18n.EN: "Other Expense Adjustment/Regularisation Entries – Tax Base",
					i18n.EL: "Λοιπές Εγγραφές Τακτοποίησης Εξόδων - Φορολογική Βάση",
				},
			},
		},
	},
	{
		Key: ExtKeyPaymentMeans,
		Name: i18n.String{
			i18n.EN: "Payment means",
			i18n.EL: "Τρόπος Πληρωμής",
		},
		Values: []*cbc.ValueDefinition{
			{
				Value: "1",
				Name: i18n.String{
					i18n.EN: "Domestic Payments Account Number",
					i18n.EL: "Επαγ. Λογαριασμός Πληρωμών Ημεδαπής",
				},
			},
			{
				Value: "2",
				Name: i18n.String{
					i18n.EN: "Foreign Payments Account Number",
					i18n.EL: "Επαγ. Λογαριασμός Πληρωμών Αλλοδαπής",
				},
			},
			{
				Value: "3",
				Name: i18n.String{
					i18n.EN: "Cash",
					i18n.EL: "Μετρητά",
				},
			},
			{
				Value: "4",
				Name: i18n.String{
					i18n.EN: "Check",
					i18n.EL: "Επιταγή",
				},
			},
			{
				Value: "5",
				Name: i18n.String{
					i18n.EN: "On credit",
					i18n.EL: "Επί Πιστώσει",
				},
			},
			{
				Value: "6",
				Name: i18n.String{
					i18n.EN: "Web Banking",
					i18n.EL: "Web Banking",
				},
			},
			{
				Value: "7",
				Name: i18n.String{
					i18n.EN: "POS / e-POS",
					i18n.EL: "POS / e-POS",
				},
			},
		},
	},
	{
		Key: ExtKeyExemption,
		Name: i18n.String{
			i18n.EN: "VAT exemption cause",
			i18n.EL: "Κατηγορία Αιτίας Εξαίρεσης ΦΠΑ",
		},
		Values: []*cbc.ValueDefinition{
			{
				Value: "1",
				Name: i18n.String{
					i18n.EN: "Without VAT - article 3 of the VAT code",
					i18n.EL: "Χωρίς ΦΠΑ - άρθρο 3 του Κώδικα ΦΠΑ",
				},
			},
			{
				Value: "2",
				Name: i18n.String{
					i18n.EN: "Without VAT - article 5 of the VAT code",
					i18n.EL: "Χωρίς ΦΠΑ - άρθρο 5 του Κώδικα ΦΠΑ",
				},
			},
			{
				Value: "3",
				Name: i18n.String{
					i18n.EN: "Without VAT - article 13 of the VAT code",
					i18n.EL: "Χωρίς ΦΠΑ - άρθρο 13 του Κώδικα ΦΠΑ",
				},
			},
			{
				Value: "4",
				Name: i18n.String{
					i18n.EN: "Without VAT - article 14 of the VAT code",
					i18n.EL: "Χωρίς ΦΠΑ - άρθρο 14 του Κώδικα ΦΠΑ",
				},
			},
			{
				Value: "5",
				Name: i18n.String{
					i18n.EN: "Without VAT - article 16 of the VAT code",
					i18n.EL: "Χωρίς ΦΠΑ - άρθρο 16 του Κώδικα ΦΠΑ",
				},
			},
			{
				Value: "6",
				Name: i18n.String{
					i18n.EN: "Without VAT - article 19 of the VAT code",
					i18n.EL: "Χωρίς ΦΠΑ - άρθρο 19 του Κώδικα ΦΠΑ",
				},
			},
			{
				Value: "7",
				Name: i18n.String{
					i18n.EN: "Without VAT - article 22 of the VAT code",
					i18n.EL: "Χωρίς ΦΠΑ - άρθρο 22 του Κώδικα ΦΠΑ",
				},
			},
			{
				Value: "8",
				Name: i18n.String{
					i18n.EN: "Without VAT - article 24 of the VAT code",
					i18n.EL: "Χωρίς ΦΠΑ - άρθρο 24 του Κώδικα ΦΠΑ",
				},
			},
			{
				Value: "9",
				Name: i18n.String{
					i18n.EN: "Without VAT - article 25 of the VAT code",
					i18n.EL: "Χωρίς ΦΠΑ - άρθρο 25 του Κώδικα ΦΠΑ",
				},
			},
			{
				Value: "10",
				Name: i18n.String{
					i18n.EN: "Without VAT - article 26 of the VAT code",
					i18n.EL: "Χωρίς ΦΠΑ - άρθρο 26 του Κώδικα ΦΠΑ",
				},
			},
			{
				Value: "11",
				Name: i18n.String{
					i18n.EN: "Without VAT - article 27 of the VAT code",
					i18n.EL: "Χωρίς ΦΠΑ - άρθρο 27 του Κώδικα ΦΠΑ",
				},
			},
			{
				Value: "12",
				Name: i18n.String{
					i18n.EN: "Without VAT - article 27 - Seagoing Vessels of the VAT code",
					i18n.EL: "Χωρίς ΦΠΑ - άρθρο 27 - Πλοία Ανοικτής Θαλάσσης του Κώδικα ΦΠΑ",
				},
			},
			{
				Value: "13",
				Name: i18n.String{
					i18n.EN: "Without VAT - article 27.1.γ - Seagoing Vessels of the VAT code",
					i18n.EL: "Χωρίς ΦΠΑ - άρθρο 27.1.γ - Πλοία Ανοικτής Θαλάσσης του Κώδικα ΦΠΑ",
				},
			},
			{
				Value: "14",
				Name: i18n.String{
					i18n.EN: "Without VAT - article 28 of the VAT code",
					i18n.EL: "Χωρίς ΦΠΑ - άρθρο 28 του Κώδικα ΦΠΑ",
				},
			},
			{
				Value: "15",
				Name: i18n.String{
					i18n.EN: "Without VAT - article 39 of the VAT code",
					i18n.EL: "Χωρίς ΦΠΑ - άρθρο 39 του Κώδικα ΦΠΑ",
				},
			},
			{
				Value: "16",
				Name: i18n.String{
					i18n.EN: "Without VAT - article 39a of the VAT code",
					i18n.EL: "Χωρίς ΦΠΑ - άρθρο 39α του Κώδικα ΦΠΑ",
				},
			},
			{
				Value: "17",
				Name: i18n.String{
					i18n.EN: "Without VAT - article 40 of the VAT code",
					i18n.EL: "Χωρίς ΦΠΑ - άρθρο 40 του Κώδικα ΦΠΑ",
				},
			},
			{
				Value: "18",
				Name: i18n.String{
					i18n.EN: "Without VAT - article 41 of the VAT code",
					i18n.EL: "Χωρίς ΦΠΑ - άρθρο 41 του Κώδικα ΦΠΑ",
				},
			},
			{
				Value: "19",
				Name: i18n.String{
					i18n.EN: "Without VAT - article 47 of the VAT code",
					i18n.EL: "Χωρίς ΦΠΑ - άρθρο 47 του Κώδικα ΦΠΑ",
				},
			},
			{
				Value: "20",
				Name: i18n.String{
					i18n.EN: "VAT included - article 43 of the VAT code",
					i18n.EL: "ΦΠΑ εμπεριεχόμενος - άρθρο 43 του Κώδικα ΦΠΑ",
				},
			},
			{
				Value: "21",
				Name: i18n.String{
					i18n.EN: "VAT included - article 44 of the VAT code",
					i18n.EL: "ΦΠΑ εμπεριεχόμενος - άρθρο 44 του Κώδικα ΦΠΑ",
				},
			},
			{
				Value: "22",
				Name: i18n.String{
					i18n.EN: "VAT included - article 45 of the VAT code",
					i18n.EL: "ΦΠΑ εμπεριεχόμενος - άρθρο 45 του Κώδικα ΦΠΑ",
				},
			},
			{
				Value: "23",
				Name: i18n.String{
					i18n.EN: "VAT included - article 46 of the VAT code",
					i18n.EL: "ΦΠΑ εμπεριεχόμενος - άρθρο 46 του Κώδικα ΦΠΑ",
				},
			},
			{
				Value: "24",
				Name: i18n.String{
					i18n.EN: "Without VAT - article 6 of the VAT code",
					i18n.EL: "Χωρίς ΦΠΑ - άρθρο 6 του Κώδικα ΦΠΑ",
				},
			},
			{
				Value: "25",
				Name: i18n.String{
					i18n.EN: "Without VAT - ΠΟΛ.1029/1995",
					i18n.EL: "Χωρίς ΦΠΑ - ΠΟΛ.1029/1995",
				},
			},
			{
				Value: "26",
				Name: i18n.String{
					i18n.EN: "Without VAT - ΠΟΛ.1167/2015",
					i18n.EL: "Χωρίς ΦΠΑ - ΠΟΛ.1167/2015",
				},
			},
			{
				Value: "27",
				Name: i18n.String{
					i18n.EN: "Without VAT - Other VAT exceptions",
					i18n.EL: "Λοιπές Εξαιρέσεις ΦΠΑ",
				},
			},
			{
				Value: "28",
				Name: i18n.String{
					i18n.EN: "Without VAT - Article 24 (b) (1) of the VAT Code (Tax Free)",
					i18n.EL: "Χωρίς ΦΠΑ – άρθρο 24 περ. β' παρ.1 του Κώδικα ΦΠΑ, (Tax Free)",
				},
			},
			{
				Value: "29",
				Name: i18n.String{
					i18n.EN: "Without VAT - Article 47b of the VAT Code (OSS non-EU scheme)",
					i18n.EL: "Χωρίς ΦΠΑ – άρθρο 47β, του Κώδικα ΦΠΑ (OSS μη ενωσιακό καθεστώς)",
				},
			},
			{
				Value: "30",
				Name: i18n.String{
					i18n.EN: "Without VAT - Article 47c of the VAT Code (OSS EU scheme)",
					i18n.EL: "Χωρίς ΦΠΑ – άρθρο 47γ, του Κώδικα ΦΠΑ (OSS ενωσιακό καθεστώς)",
				},
			},
			{
				Value: "31",
				Name: i18n.String{
					i18n.EN: "Excluding VAT - Article 47d of the VAT Code (IOSS)",
					i18n.EL: "Χωρίς ΦΠΑ – άρθρο 47δ του Κώδικα ΦΠΑ (IOSS)",
				},
			},
		},
	},
	{
		Key: ExtKeyIncomeCat,
		Name: i18n.String{
			i18n.EN: "Income Classification Category",
			i18n.EL: "Κωδικός Κατηγορίας Χαρακτηρισμού Εσόδων",
		},
		Values: []*cbc.ValueDefinition{
			{
				Value: "category1_1",
				Name: i18n.String{
					i18n.EN: "Commodity Sale Income (+)/(-)",
					i18n.EL: "Έσοδα από Πώληση Εμπορευμάτων (+)/(-)",
				},
			},
			{
				Value: "category1_2",
				Name: i18n.String{
					i18n.EN: "Product Sale Income (+)/(-)",
					i18n.EL: "Έσοδα από Πώληση Προϊόντων (+)/(-)",
				},
			},
			{
				Value: "category1_3",
				Name: i18n.String{
					i18n.EN: "Provision of Services Income (+)/(-)",
					i18n.EL: "Έσοδα από Παροχή Υπηρεσιών (+)/(-)",
				},
			},
			{
				Value: "category1_4",
				Name: i18n.String{
					i18n.EN: "Sale of Fixed Assets Income (+)/(-)",
					i18n.EL: "Έσοδα από Πώληση Παγίων (+)/(-)",
				},
			},
			{
				Value: "category1_5",
				Name: i18n.String{
					i18n.EN: "Other Income/Profits (+)/(-)",
					i18n.EL: "Λοιπά Έσοδα/ Κέρδη (+)/(-)",
				},
			},
			{
				Value: "category1_6",
				Name: i18n.String{
					i18n.EN: "Self-Deliveries/Self-Supplies (+)/(-)",
					i18n.EL: "Αυτοπαραδόσεις / Ιδιοχρησιμοποιήσεις (+)/(-)",
				},
			},
			{
				Value: "category1_7",
				Name: i18n.String{
					i18n.EN: "Income on behalf of Third Parties (+)/(-)",
					i18n.EL: "Έσοδα για λ/σμο τρίτων (+)/(-)",
				},
			},
			{
				Value: "category1_8",
				Name: i18n.String{
					i18n.EN: "Past fiscal years income (+)/(-)",
					i18n.EL: "Έσοδα προηγούμενων χρήσεων (+)/ (-)",
				},
			},
			{
				Value: "category1_9",
				Name: i18n.String{
					i18n.EN: "Future fiscal years income (+)/(-)",
					i18n.EL: "Έσοδα επομένων χρήσεων (+)/(-)",
				},
			},
			{
				Value: "category1_10",
				Name: i18n.String{
					i18n.EN: "Other Income Adjustment/Regularisation Entries (+)/(-)",
					i18n.EL: "Λοιπές Εγγραφές Τακτοποίησης Εσόδων (+)/(-)",
				},
			},
			{
				Value: "category1_95",
				Name: i18n.String{
					i18n.EN: "Other Income-related Information (+)/(-)",
					i18n.EL: "Λοιπά Πληροφοριακά Στοιχεία Εσόδων (+)/(-)",
				},
			},
		},
	},
	{
		Key: ExtKeyIncomeType,
		Name: i18n.String{
			i18n.EN: "Income Classification Type",
			i18n.EL: "Κωδικός Τύπου Χαρακτηρισμού Εσόδων",
		},
		Values: []*cbc.ValueDefinition{
			{
				Value: "E3_106",
				Name: i18n.String{
					i18n.EN: "Self-Production of Fixed Assets – Self-Deliveries – Destroying inventory/Commodities",
					i18n.EL: "Ιδιοπαραγωγή παγίων - Αυτοπαραδόσεις - Καταστροφές αποθεμάτων/Εμπορεύματα",
				},
			},
			{
				Value: "E3_205",
				Name: i18n.String{
					i18n.EN: "Self-Production of Fixed Assets – Self-Deliveries – Destroying inventory/Raw and other materials",
					i18n.EL: "Ιδιοπαραγωγή παγίων - Αυτοπαραδόσεις - Καταστροφές αποθεμάτων/Πρώτες ύλες και λοιπά υλικά",
				},
			},
			{
				Value: "E3_210",
				Name: i18n.String{
					i18n.EN: "Self-Production of Fixed Assets – Self-Deliveries – Destroying inventory/Products and production in progress",
					i18n.EL: "Ιδιοπαραγωγή παγίων - Αυτοπαραδόσεις - Καταστροφές αποθεμάτων/Προϊόντα και παραγωγή σε εξέλιξη",
				},
			},
			{
				Value: "E3_305",
				Name: i18n.String{
					i18n.EN: "Self-Production of Fixed Assets – Self-Deliveries – Destroying inventory/Raw and other materials",
					i18n.EL: "Ιδιοπαραγωγή παγίων - Αυτοπαραδόσεις – Καταστροφές αποθεμάτων/Πρώτες ύλες και λοιπά υλικά",
				},
			},
			{
				Value: "E3_310",
				Name: i18n.String{
					i18n.EN: "Self-Production of Fixed Assets – Self-Deliveries – Destroying inventory/Products and production in progress",
					i18n.EL: "Ιδιοπαραγωγή παγίων - Αυτοπαραδόσεις - Καταστροφές αποθεμάτων/Προϊόντα και παραγωγή σε εξέλιξη",
				},
			},
			{
				Value: "E3_318",
				Name: i18n.String{
					i18n.EN: "Self-Production of Fixed Assets – Self-Deliveries – Destroying inventory/Production expenses",
					i18n.EL: "Ιδιοπαραγωγή παγίων - Αυτοπαραδόσεις - Καταστροφές αποθεμάτων/Έξοδα παραγωγής",
				},
			},
			{
				Value: "E3_561_001",
				Name: i18n.String{
					i18n.EN: "Wholesale Sales of Goods and Services – for Traders",
					i18n.EL: "Πωλήσεις αγαθών και υπηρεσιών Χονδρικές - Επιτηδευματιών",
				},
			},
			{
				Value: "E3_561_002",
				Name: i18n.String{
					i18n.EN: "Wholesale Sales of Goods and Services pursuant to article 39a paragraph 5 of the VAT Code (Law 2859/2000)",
					i18n.EL: "Πωλήσεις αγαθών και υπηρεσιών Χονδρικές βάσει άρθρου 39α παρ 5 του Κώδικα Φ.Π.Α. (Ν.2859/2000)",
				},
			},
			{
				Value: "E3_561_003",
				Name: i18n.String{
					i18n.EN: "Retail Sales of Goods and Services – Private Clientele",
					i18n.EL: "Πωλήσεις αγαθών και υπηρεσιών Λιανικές - Ιδιωτική Πελατεία",
				},
			},
			{
				Value: "E3_561_004",
				Name: i18n.String{
					i18n.EN: "Retail Sales of Goods and Services pursuant to article 39a paragraph 5 of the VAT Code (Law 2859/2000)",
					i18n.EL: "Πωλήσεις αγαθών και υπηρεσιών Λιανικές βάσει άρθρου 39α παρ 5 του Κώδικα Φ.Π.Α. (Ν.2859/2000)",
				},
			},
			{
				Value: "E3_561_005",
				Name: i18n.String{
					i18n.EN: "Intra-Community Foreign Sales of Goods and Services",
					i18n.EL: "Πωλήσεις αγαθών και υπηρεσιών Εξωτερικού Ενδοκοινοτικές",
				},
			},
			{
				Value: "E3_561_006",
				Name: i18n.String{
					i18n.EN: "Third Country Foreign Sales of Goods and Services",
					i18n.EL: "Πωλήσεις αγαθών και υπηρεσιών Εξωτερικού Τρίτες Χώρες",
				},
			},
			{
				Value: "E3_561_007",
				Name: i18n.String{
					i18n.EN: "Other Sales of Goods and Services",
					i18n.EL: "Πωλήσεις αγαθών και υπηρεσιών Λοιπά",
				},
			},
			{
				Value: "E3_562",
				Name: i18n.String{
					i18n.EN: "Other Ordinary Income",
					i18n.EL: "Λοιπά συνήθη έσοδα",
				},
			},
			{
				Value: "E3_563",
				Name: i18n.String{
					i18n.EN: "Credit Interest and Related Income",
					i18n.EL: "Πιστωτικοί τόκοι και συναφή έσοδα",
				},
			},
			{
				Value: "E3_564",
				Name: i18n.String{
					i18n.EN: "Credit Exchange Differences",
					i18n.EL: "Πιστωτικές συναλλαγματικές διαφορές",
				},
			},
			{
				Value: "E3_565",
				Name: i18n.String{
					i18n.EN: "Income from Participations",
					i18n.EL: "Έσοδα συμμετοχών",
				},
			},
			{
				Value: "E3_566",
				Name: i18n.String{
					i18n.EN: "Profits from Disposing Non-Current Assets",
					i18n.EL: "Κέρδη από διάθεση μη κυκλοφορούντων περιουσιακών στοιχείων",
				},
			},
			{
				Value: "E3_567",
				Name: i18n.String{
					i18n.EN: "Profits from the Reversal of Provisions and Impairments",
					i18n.EL: "Κέρδη από αναστροφή προβλέψεων και απομειώσεων",
				},
			},
			{
				Value: "E3_568",
				Name: i18n.String{
					i18n.EN: "Profits from Measurement at Fair Value",
					i18n.EL: "Κέρδη από επιμέτρηση στην εύλογη αξία",
				},
			},
			{
				Value: "E3_570",
				Name: i18n.String{
					i18n.EN: "Extraordinary income and profits",
					i18n.EL: "Ασυνήθη έσοδα και κέρδη",
				},
			},
			{
				Value: "E3_595",
				Name: i18n.String{
					i18n.EN: "Self-Production Expenses",
					i18n.EL: "Έξοδα σε ιδιοπαραγωγή",
				},
			},
			{
				Value: "E3_596",
				Name: i18n.String{
					i18n.EN: "Subsidies - Grants",
					i18n.EL: "Επιδοτήσεις - Επιχορηγήσεις",
				},
			},
			{
				Value: "E3_597",
				Name: i18n.String{
					i18n.EN: "Subsidies – Grants for Investment Purposes – Expense Coverage",
					i18n.EL: "Επιδοτήσεις - Επιχορηγήσεις για επενδυτικούς σκοπούς - κάλυψη δαπανών",
				},
			},
			{
				Value: "E3_880_001",
				Name: i18n.String{
					i18n.EN: "Wholesale Sales of Fixed Assets",
					i18n.EL: "Πωλήσεις Παγίων Χονδρικές",
				},
			},
			{
				Value: "E3_880_002",
				Name: i18n.String{
					i18n.EN: "Retail Sales of Fixed Assets",
					i18n.EL: "Πωλήσεις Παγίων Λιανικές",
				},
			},
			{
				Value: "E3_880_003",
				Name: i18n.String{
					i18n.EN: "Intra-Community Foreign Sales of Fixed Assets",
					i18n.EL: "Πωλήσεις Παγίων Εξωτερικού Ενδοκοινοτικές",
				},
			},
			{
				Value: "E3_880_004",
				Name: i18n.String{
					i18n.EN: "Third Country Foreign Sales of Fixed Assets",
					i18n.EL: "Πωλήσεις Παγίων Εξωτερικού Τρίτες Χώρες",
				},
			},
			{
				Value: "E3_881_001",
				Name: i18n.String{
					i18n.EN: "Wholesale Sales on behalf of Third Parties",
					i18n.EL: "Πωλήσεις για λογ/σμο Τρίτων Χονδρικές",
				},
			},
			{
				Value: "E3_881_002",
				Name: i18n.String{
					i18n.EN: "Retail Sales on behalf of Third Parties",
					i18n.EL: "Πωλήσεις για λογ/σμο Τρίτων Λιανικές",
				},
			},
			{
				Value: "E3_881_003",
				Name: i18n.String{
					i18n.EN: "Intra-Community Foreign Sales on behalf of Third Parties",
					i18n.EL: "Πωλήσεις για λογ/σμο Τρίτων Εξωτερικού Ενδοκοινοτικές",
				},
			},
			{
				Value: "E3_881_004",
				Name: i18n.String{
					i18n.EN: "Third Country Foreign Sales on behalf of Third Parties",
					i18n.EL: "Πωλήσεις για λογ/σμο Τρίτων Εξωτερικού Τρίτες Χώρες",
				},
			},
			{
				Value: "E3_598_001",
				Name: i18n.String{
					i18n.EN: "Sales of goods belonging to excise duty",
					i18n.EL: "Πωλήσεις αγαθών που υπάγονται σε ΕΦΚ",
				},
			},
			{
				Value: "E3_598_003",
				Name: i18n.String{
					i18n.EN: "Sales on behalf of farmers through an agricultural cooperative e.t.c.",
					i18n.EL: "Πωλήσεις για λογαριασμό αγροτών μέσω αγροτικού συνεταιρισμού κ.λ.π.",
				},
			},
		},
	},
}

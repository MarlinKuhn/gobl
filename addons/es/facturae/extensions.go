package facturae

import (
	"github.com/invopop/gobl/cbc"
	"github.com/invopop/gobl/i18n"
)

// Extension keys for use with FacturaE
const (
	ExtKeyCorrection   cbc.Key = "es-facturae-correction"
	ExtKeyDocType      cbc.Key = "es-facturae-doc-type"
	ExtKeyInvoiceClass cbc.Key = "es-facturae-invoice-class"
)

var extensions = []*cbc.KeyDefinition{
	{
		Key: ExtKeyDocType,
		Name: i18n.String{
			i18n.EN: "FacturaE: Document Type",
			i18n.ES: "FacturaE: Tipo de Documento",
		},
		Values: []*cbc.ValueDefinition{
			{
				Value: "FC",
				Name: i18n.String{
					i18n.EN: "Commercial Invoice",
					i18n.ES: "Factura Comercial",
				},
			},
			{
				Value: "FA",
				Name: i18n.String{
					i18n.EN: "Simplified Invoice",
					i18n.ES: "Factura Simplificada",
				},
			},
			{
				Value: "AF",
				Name: i18n.String{
					i18n.EN: "Self-billed Invoice",
					i18n.ES: "Auto-Factura",
				},
			},
		},
	},
	{
		Key: ExtKeyInvoiceClass,
		Name: i18n.String{
			i18n.EN: "FacturaE: Invoice Class",
			i18n.ES: "FacturaE: Clase de Factura",
		},
		Values: []*cbc.ValueDefinition{
			{
				Value: "OO",
				Name: i18n.String{
					i18n.EN: "Original",
					i18n.ES: "Original",
				},
			},
			{
				Value: "OR",
				Name: i18n.String{
					i18n.EN: "Corrective Original",
					i18n.ES: "Original Rectificativa",
				},
			},
			{
				Value: "OC",
				Name: i18n.String{
					i18n.EN: "Summary Original",
					i18n.ES: "Original Recapitulativa",
				},
			},
			{
				Value: "CO",
				Name: i18n.String{
					i18n.EN: "Copy of the Original",
					i18n.ES: "Duplicado Original",
				},
			},
			{
				Value: "CR",
				Name: i18n.String{
					i18n.EN: "Copy of the Corrective",
					i18n.ES: "Duplicado Rectificativa",
				},
			},
			{
				Value: "CC",
				Name: i18n.String{
					i18n.EN: "Copy of the Summary",
					i18n.ES: "Duplicado Recapitulativa",
				},
			},
		},
	},
	{
		Key: ExtKeyCorrection,
		Name: i18n.String{
			i18n.EN: "FacturaE Change",
			i18n.ES: "Cambio de FacturaE",
		},
		Desc: i18n.String{
			i18n.EN: "FacturaE requires a specific and single code that explains why the previous invoice is being corrected.",
			i18n.ES: "FacturaE requiere un código específico y único que explique por qué se está corrigiendo la factura anterior.",
		},
		// Codes take from FacturaE XSD
		Values: []*cbc.ValueDefinition{
			{
				Value: "01",
				Name: i18n.String{
					i18n.EN: "Invoice code",
					i18n.ES: "Número de la factura",
				},
			},
			{
				Value: "02",
				Name: i18n.String{
					i18n.EN: "Invoice series",
					i18n.ES: "Serie de la factura",
				},
			},
			{
				Value: "03",
				Name: i18n.String{
					i18n.EN: "Issue date",
					i18n.ES: "Fecha expedición",
				},
			},
			{
				Value: "04",
				Name: i18n.String{
					i18n.EN: "Name and surnames/Corporate name - Issuer (Sender)",
					i18n.ES: "Nombre y apellidos/Razón Social-Emisor",
				},
			},
			{
				Value: "05",
				Name: i18n.String{
					i18n.EN: "Name and surnames/Corporate name - Receiver",
					i18n.ES: "Nombre y apellidos/Razón Social-Receptor",
				},
			},
			{
				Value: "06",
				Name: i18n.String{
					i18n.EN: "Issuer's Tax Identification Number",
					i18n.ES: "Identificación fiscal Emisor/obligado",
				},
			},
			{
				Value: "07",
				Name: i18n.String{
					i18n.EN: "Receiver's Tax Identification Number",
					i18n.ES: "Identificación fiscal Receptor",
				},
			},
			{
				Value: "08",
				Name: i18n.String{
					i18n.EN: "Supplier's address",
					i18n.ES: "Domicilio Emisor/Obligado",
				},
			},
			{
				Value: "09",
				Name: i18n.String{
					i18n.EN: "Customer's address",
					i18n.ES: "Domicilio Receptor",
				},
			},
			{
				Value: "10",
				Name: i18n.String{
					i18n.EN: "Item line",
					i18n.ES: "Detalle Operación",
				},
			},
			{
				Value: "11",
				Name: i18n.String{
					i18n.EN: "Applicable Tax Rate",
					i18n.ES: "Porcentaje impositivo a aplicar",
				},
			},
			{
				Value: "12",
				Name: i18n.String{
					i18n.EN: "Applicable Tax Amount",
					i18n.ES: "Cuota tributaria a aplicar",
				},
			},
			{
				Value: "13",
				Name: i18n.String{
					i18n.EN: "Applicable Date/Period",
					i18n.ES: "Fecha/Periodo a aplicar",
				},
			},
			{
				Value: "14",
				Name: i18n.String{
					i18n.EN: "Invoice Class",
					i18n.ES: "Clase de factura",
				},
			},
			{
				Value: "15",
				Name: i18n.String{
					i18n.EN: "Legal literals",
					i18n.ES: "Literales legales",
				},
			},
			{
				Value: "16",
				Name: i18n.String{
					i18n.EN: "Taxable Base",
					i18n.ES: "Base imponible",
				},
			},
			{
				Value: "80",
				Name: i18n.String{
					i18n.EN: "Calculation of tax outputs",
					i18n.ES: "Cálculo de cuotas repercutidas",
				},
			},
			{
				Value: "81",
				Name: i18n.String{
					i18n.EN: "Calculation of tax inputs",
					i18n.ES: "Cálculo de cuotas retenidas",
				},
			},
			{
				Value: "82",
				Name: i18n.String{
					i18n.EN: "Taxable Base modified due to return of packages and packaging materials",
					i18n.ES: "Base imponible modificada por devolución de envases / embalajes",
				},
			},
			{
				Value: "83",
				Name: i18n.String{
					i18n.EN: "Taxable Base modified due to discounts and rebates",
					i18n.ES: "Base imponible modificada por descuentos y bonificaciones",
				},
			},
			{
				Value: "84",
				Name: i18n.String{
					i18n.EN: "Taxable Base modified due to firm court ruling or administrative decision",
					i18n.ES: "Base imponible modificada por resolución firme, judicial o administrativa",
				},
			},
			{
				Value: "85",
				Name: i18n.String{
					i18n.EN: "Taxable Base modified due to unpaid outputs where there is a judgement opening insolvency proceedings",
					i18n.ES: "Base imponible modificada cuotas repercutidas no satisfechas. Auto de declaración de concurso",
				},
			},
		},
	},
}

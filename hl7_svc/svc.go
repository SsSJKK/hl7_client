package hl7svc

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/kardianos/hl7"
	"github.com/kardianos/hl7/h231"
)

type Object struct {
	HL7 struct {
	} `json:"HL7"`
	MSH struct {
		HL7 struct {
		} `json:"HL7"`
		FieldSeparator       string      `json:"FieldSeparator"`
		EncodingCharacters   string      `json:"EncodingCharacters"`
		SendingApplication   interface{} `json:"SendingApplication"`
		SendingFacility      interface{} `json:"SendingFacility"`
		ReceivingApplication interface{} `json:"ReceivingApplication"`
		ReceivingFacility    interface{} `json:"ReceivingFacility"`
		DateTimeOfMessage    time.Time   `json:"DateTimeOfMessage"`
		Security             string      `json:"Security"`
		MessageType          struct {
			HL7 struct {
			} `json:"HL7"`
			MessageType      string `json:"MessageType"`
			TriggerEvent     string `json:"TriggerEvent"`
			MessageStructure string `json:"MessageStructure"`
		} `json:"MessageType"`
		MessageControlID string `json:"MessageControlID"`
		ProcessingID     struct {
			HL7 struct {
			} `json:"HL7"`
			ProcessingID   string `json:"ProcessingID"`
			ProcessingMode string `json:"ProcessingMode"`
		} `json:"ProcessingID"`
		VersionID struct {
			HL7 struct {
			} `json:"HL7"`
			VersionID                string      `json:"VersionID"`
			InternationalizationCode interface{} `json:"InternationalizationCode"`
			InternationalVersionID   interface{} `json:"InternationalVersionID"`
		} `json:"VersionID"`
		SequenceNumber                      string      `json:"SequenceNumber"`
		ContinuationPointer                 string      `json:"ContinuationPointer"`
		AcceptAcknowledgmentType            string      `json:"AcceptAcknowledgmentType"`
		ApplicationAcknowledgmentType       string      `json:"ApplicationAcknowledgmentType"`
		CountryCode                         string      `json:"CountryCode"`
		CharacterSet                        []string    `json:"CharacterSet"`
		PrincipalLanguageOfMessage          interface{} `json:"PrincipalLanguageOfMessage"`
		AlternateCharacterSetHandlingScheme string      `json:"AlternateCharacterSetHandlingScheme"`
	} `json:"MSH"`
	PatientResult []struct {
		HL7 struct {
		} `json:"HL7"`
		Patient struct {
			HL7 struct {
			} `json:"HL7"`
			PID struct {
				HL7 struct {
				} `json:"HL7"`
				SetID     string `json:"SetID"`
				PatientID struct {
					HL7 struct {
					} `json:"HL7"`
					ID                                         string      `json:"ID"`
					CheckDigit                                 string      `json:"CheckDigit"`
					CodeIdentifyingTheCheckDigitSchemeEmployed string      `json:"CodeIdentifyingTheCheckDigitSchemeEmployed"`
					AssigningAuthority                         interface{} `json:"AssigningAuthority"`
					IdentifierTypeCode                         string      `json:"IdentifierTypeCode"`
					AssigningFacility                          interface{} `json:"AssigningFacility"`
				} `json:"PatientID"`
				PatientIdentifierList []struct {
					HL7 struct {
					} `json:"HL7"`
					ID                                         string `json:"ID"`
					CheckDigit                                 string `json:"CheckDigit"`
					CodeIdentifyingTheCheckDigitSchemeEmployed string `json:"CodeIdentifyingTheCheckDigitSchemeEmployed"`
					AssigningAuthority                         struct {
						HL7 struct {
						} `json:"HL7"`
						NamespaceID     string `json:"NamespaceID"`
						UniversalID     string `json:"UniversalID"`
						UniversalIDType string `json:"UniversalIDType"`
					} `json:"AssigningAuthority"`
					IdentifierTypeCode string      `json:"IdentifierTypeCode"`
					AssigningFacility  interface{} `json:"AssigningFacility"`
				} `json:"PatientIdentifierList"`
				AlternatePatientID interface{} `json:"AlternatePatientID"`
				PatientName        []struct {
					HL7 struct {
					} `json:"HL7"`
					FamilyNameLastNamePrefix string `json:"FamilyNameLastNamePrefix"`
					GivenName                string `json:"GivenName"`
					MiddleInitialOrName      string `json:"MiddleInitialOrName"`
					Suffix                   string `json:"Suffix"`
					Prefix                   string `json:"Prefix"`
					Degree                   string `json:"Degree"`
					NameTypeCode             string `json:"NameTypeCode"`
					NameRepresentationCode   string `json:"NameRepresentationCode"`
				} `json:"PatientName"`
				MotherSMaidenName           interface{} `json:"MotherSMaidenName"`
				DateTimeOfBirth             time.Time   `json:"DateTimeOfBirth"`
				Sex                         string      `json:"Sex"`
				PatientAlias                interface{} `json:"PatientAlias"`
				Race                        interface{} `json:"Race"`
				PatientAddress              interface{} `json:"PatientAddress"`
				CountyCode                  string      `json:"CountyCode"`
				PhoneNumberHome             interface{} `json:"PhoneNumberHome"`
				PhoneNumberBusiness         interface{} `json:"PhoneNumberBusiness"`
				PrimaryLanguage             interface{} `json:"PrimaryLanguage"`
				MaritalStatus               interface{} `json:"MaritalStatus"`
				Religion                    interface{} `json:"Religion"`
				PatientAccountNumber        interface{} `json:"PatientAccountNumber"`
				SSNNumberPatient            string      `json:"SSNNumberPatient"`
				DriversLicenseNumberPatient interface{} `json:"DriversLicenseNumberPatient"`
				MothersIdentifier           interface{} `json:"MothersIdentifier"`
				EthnicGroup                 interface{} `json:"EthnicGroup"`
				BirthPlace                  string      `json:"BirthPlace"`
				MultipleBirthIndicator      string      `json:"MultipleBirthIndicator"`
				BirthOrder                  string      `json:"BirthOrder"`
				Citizenship                 interface{} `json:"Citizenship"`
				VeteransMilitaryStatus      interface{} `json:"VeteransMilitaryStatus"`
				Nationality                 interface{} `json:"Nationality"`
				PatientDeathDateAndTime     time.Time   `json:"PatientDeathDateAndTime"`
				PatientDeathIndicator       string      `json:"PatientDeathIndicator"`
			} `json:"PID"`
			PD1   interface{} `json:"PD1"`
			NK1   interface{} `json:"NK1"`
			NTE   interface{} `json:"NTE"`
			Visit struct {
				HL7 struct {
				} `json:"HL7"`
				PV1 struct {
					HL7 struct {
					} `json:"HL7"`
					SetID                   string      `json:"SetID"`
					PatientClass            string      `json:"PatientClass"`
					AssignedPatientLocation interface{} `json:"AssignedPatientLocation"`
					AdmissionType           string      `json:"AdmissionType"`
					PreadmitNumber          interface{} `json:"PreadmitNumber"`
					PriorPatientLocation    interface{} `json:"PriorPatientLocation"`
					AttendingDoctor         interface{} `json:"AttendingDoctor"`
					ReferringDoctor         interface{} `json:"ReferringDoctor"`
					ConsultingDoctor        interface{} `json:"ConsultingDoctor"`
					HospitalService         string      `json:"HospitalService"`
					TemporaryLocation       interface{} `json:"TemporaryLocation"`
					PreadmitTestIndicator   string      `json:"PreadmitTestIndicator"`
					ReAdmissionIndicator    string      `json:"ReAdmissionIndicator"`
					AdmitSource             string      `json:"AdmitSource"`
					AmbulatoryStatus        interface{} `json:"AmbulatoryStatus"`
					VIPIndicator            string      `json:"VIPIndicator"`
					AdmittingDoctor         interface{} `json:"AdmittingDoctor"`
					PatientType             string      `json:"PatientType"`
					VisitNumber             interface{} `json:"VisitNumber"`
					FinancialClass          interface{} `json:"FinancialClass"`
					ChargePriceIndicator    string      `json:"ChargePriceIndicator"`
					CourtesyCode            string      `json:"CourtesyCode"`
					CreditRating            string      `json:"CreditRating"`
					ContractCode            interface{} `json:"ContractCode"`
					ContractEffectiveDate   interface{} `json:"ContractEffectiveDate"`
					ContractAmount          interface{} `json:"ContractAmount"`
					ContractPeriod          interface{} `json:"ContractPeriod"`
					InterestCode            string      `json:"InterestCode"`
					TransferToBadDebtCode   string      `json:"TransferToBadDebtCode"`
					TransferToBadDebtDate   time.Time   `json:"TransferToBadDebtDate"`
					BadDebtAgencyCode       string      `json:"BadDebtAgencyCode"`
					BadDebtTransferAmount   string      `json:"BadDebtTransferAmount"`
					BadDebtRecoveryAmount   string      `json:"BadDebtRecoveryAmount"`
					DeleteAccountIndicator  string      `json:"DeleteAccountIndicator"`
					DeleteAccountDate       time.Time   `json:"DeleteAccountDate"`
					DischargeDisposition    string      `json:"DischargeDisposition"`
					DischargedToLocation    interface{} `json:"DischargedToLocation"`
					DietType                interface{} `json:"DietType"`
					ServicingFacility       string      `json:"ServicingFacility"`
					BedStatus               string      `json:"BedStatus"`
					AccountStatus           string      `json:"AccountStatus"`
					PendingLocation         interface{} `json:"PendingLocation"`
					PriorTemporaryLocation  interface{} `json:"PriorTemporaryLocation"`
					AdmitDateTime           time.Time   `json:"AdmitDateTime"`
					DischargeDateTime       time.Time   `json:"DischargeDateTime"`
					CurrentPatientBalance   string      `json:"CurrentPatientBalance"`
					TotalCharges            string      `json:"TotalCharges"`
					TotalAdjustments        string      `json:"TotalAdjustments"`
					TotalPayments           string      `json:"TotalPayments"`
					AlternateVisitID        interface{} `json:"AlternateVisitID"`
					VisitIndicator          string      `json:"VisitIndicator"`
					OtherHealthcareProvider interface{} `json:"OtherHealthcareProvider"`
				} `json:"PV1"`
				PV2 interface{} `json:"PV2"`
			} `json:"Visit"`
		} `json:"Patient"`
		OrderObservation []struct {
			HL7 struct {
			} `json:"HL7"`
			ORC interface{} `json:"ORC"`
			OBR struct {
				HL7 struct {
				} `json:"HL7"`
				SetID             string      `json:"SetID"`
				PlacerOrderNumber interface{} `json:"PlacerOrderNumber"`
				FillerOrderNumber struct {
					HL7 struct {
					} `json:"HL7"`
					EntityIdentifier string `json:"EntityIdentifier"`
					NamespaceID      string `json:"NamespaceID"`
					UniversalID      string `json:"UniversalID"`
					UniversalIDType  string `json:"UniversalIDType"`
				} `json:"FillerOrderNumber"`
				UniversalServiceID struct {
					HL7 struct {
					} `json:"HL7"`
					Identifier                  string `json:"Identifier"`
					Text                        string `json:"Text"`
					NameOfCodingSystem          string `json:"NameOfCodingSystem"`
					AlternateComponents         string `json:"AlternateComponents"`
					AlternateText               string `json:"AlternateText"`
					NameOfAlternateCodingSystem string `json:"NameOfAlternateCodingSystem"`
				} `json:"UniversalServiceID"`
				Priority                     string      `json:"Priority"`
				RequestedDateTime            time.Time   `json:"RequestedDateTime"`
				ObservationDateTime          time.Time   `json:"ObservationDateTime"`
				ObservationEndDateTime       time.Time   `json:"ObservationEndDateTime"`
				CollectionVolume             interface{} `json:"CollectionVolume"`
				CollectorIdentifier          interface{} `json:"CollectorIdentifier"`
				SpecimenActionCode           string      `json:"SpecimenActionCode"`
				DangerCode                   interface{} `json:"DangerCode"`
				RelevantClinicalInfo         string      `json:"RelevantClinicalInfo"`
				SpecimenReceivedDateTime     time.Time   `json:"SpecimenReceivedDateTime"`
				SpecimenSource               interface{} `json:"SpecimenSource"`
				OrderingProvider             interface{} `json:"OrderingProvider"`
				OrderCallbackPhoneNumber     interface{} `json:"OrderCallbackPhoneNumber"`
				PlacerField1                 string      `json:"PlacerField1"`
				PlacerField2                 string      `json:"PlacerField2"`
				FillerField1                 string      `json:"FillerField1"`
				FillerField2                 string      `json:"FillerField2"`
				ResultsRptStatusChngDateTime time.Time   `json:"ResultsRptStatusChngDateTime"`
				ChargeToPractice             interface{} `json:"ChargeToPractice"`
				DiagnosticServSectID         string      `json:"DiagnosticServSectID"`
				ResultStatus                 string      `json:"ResultStatus"`
				ParentResult                 interface{} `json:"ParentResult"`
				QuantityTiming               interface{} `json:"QuantityTiming"`
				ResultCopiesTo               interface{} `json:"ResultCopiesTo"`
				ParentNumber                 interface{} `json:"ParentNumber"`
				TransportationMode           string      `json:"TransportationMode"`
				ReasonForStudy               interface{} `json:"ReasonForStudy"`
				PrincipalResultInterpreter   struct {
					HL7 struct {
					} `json:"HL7"`
					OPName struct {
						HL7 struct {
						} `json:"HL7"`
						IDNumber            string      `json:"IDNumber"`
						FamilyName          string      `json:"FamilyName"`
						GivenName           string      `json:"GivenName"`
						MiddleInitialOrName string      `json:"MiddleInitialOrName"`
						Suffix              string      `json:"Suffix"`
						Prefix              string      `json:"Prefix"`
						Degree              string      `json:"Degree"`
						SourceTable         string      `json:"SourceTable"`
						AssigningAuthority  interface{} `json:"AssigningAuthority"`
					} `json:"OPName"`
					StartDateTime      time.Time   `json:"StartDateTime"`
					EndDateTime        time.Time   `json:"EndDateTime"`
					PointOfCare        string      `json:"PointOfCare"`
					Room               string      `json:"Room"`
					Bed                string      `json:"Bed"`
					Facility           interface{} `json:"Facility"`
					LocationStatus     string      `json:"LocationStatus"`
					PersonLocationType string      `json:"PersonLocationType"`
					Building           string      `json:"Building"`
					Floor              string      `json:"Floor"`
				} `json:"PrincipalResultInterpreter"`
				AssistantResultInterpreter          interface{} `json:"AssistantResultInterpreter"`
				Technician                          interface{} `json:"Technician"`
				Transcriptionist                    interface{} `json:"Transcriptionist"`
				ScheduledDateTime                   time.Time   `json:"ScheduledDateTime"`
				NumberOfSampleContainers            string      `json:"NumberOfSampleContainers"`
				TransportLogisticsOfCollectedSample interface{} `json:"TransportLogisticsOfCollectedSample"`
				CollectorSComment                   interface{} `json:"CollectorSComment"`
				TransportArrangementResponsibility  interface{} `json:"TransportArrangementResponsibility"`
				TransportArranged                   string      `json:"TransportArranged"`
				EscortRequired                      string      `json:"EscortRequired"`
				PlannedPatientTransportComment      interface{} `json:"PlannedPatientTransportComment"`
				ProcedureCode                       interface{} `json:"ProcedureCode"`
				ProcedureCodeModifier               interface{} `json:"ProcedureCodeModifier"`
			} `json:"OBR"`
			NTE         interface{} `json:"NTE"`
			Observation []struct {
				HL7 struct {
				} `json:"HL7"`
				OBX struct {
					HL7 struct {
					} `json:"HL7"`
					SetID                 string `json:"SetID"`
					ValueType             string `json:"ValueType"`
					ObservationIdentifier struct {
						HL7 struct {
						} `json:"HL7"`
						Identifier                  string `json:"Identifier"`
						Text                        string `json:"Text"`
						NameOfCodingSystem          string `json:"NameOfCodingSystem"`
						AlternateComponents         string `json:"AlternateComponents"`
						AlternateText               string `json:"AlternateText"`
						NameOfAlternateCodingSystem string `json:"NameOfAlternateCodingSystem"`
					} `json:"ObservationIdentifier"`
					ObservationSubID         string      `json:"ObservationSubID"`
					ObservationValue         []string    `json:"ObservationValue"`
					Units                    interface{} `json:"Units"`
					ReferencesRange          string      `json:"ReferencesRange"`
					AbnormalFlags            interface{} `json:"AbnormalFlags"`
					Probability              interface{} `json:"Probability"`
					NatureOfAbnormalTest     interface{} `json:"NatureOfAbnormalTest"`
					ObservationResultStatus  string      `json:"ObservationResultStatus"`
					DateLastObsNormalValues  time.Time   `json:"DateLastObsNormalValues"`
					UserDefinedAccessChecks  string      `json:"UserDefinedAccessChecks"`
					DateTimeOfTheObservation time.Time   `json:"DateTimeOfTheObservation"`
					ProducersID              interface{} `json:"ProducersID"`
					ResponsibleObserver      interface{} `json:"ResponsibleObserver"`
					ObservationMethod        interface{} `json:"ObservationMethod"`
				} `json:"OBX"`
				NTE interface{} `json:"NTE"`
			} `json:"Observation"`
			CTI interface{} `json:"CTI"`
		} `json:"OrderObservation"`
	} `json:"PatientResult"`
	DSC interface{} `json:"DSC"`
}

func Decoder(data []byte) {
	hl7Decoder := hl7.NewDecoder(h231.Registry, nil)
	fmt.Printf("data: %s\n", data) //
	// parceData, err := hl7Decoder.Decode(data)
	parceData, err := hl7Decoder.Decode(data)
	// fmt.Printf("err: %+v\n", err)/
	jData, err := json.Marshal(parceData)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Printf("parceData: %+v\n", parceData)
	fmt.Printf("parceData: %s\n", jData)
}

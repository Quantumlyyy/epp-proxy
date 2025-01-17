use chrono::prelude::*;

/// domain-ext-2.5

#[derive(Debug, Deserialize)]
pub struct EURIDDomainCheckData {
    #[serde(rename = "{http://www.eurid.eu/xml/epp/domain-ext-2.5}domain")]
    pub domains: Vec<EURIDDomainCheckDatum>,
}

#[derive(Debug, Deserialize, Clone)]
pub struct EURIDDomainCheckDatum {
    #[serde(rename = "{http://www.eurid.eu/xml/epp/domain-ext-2.5}name")]
    pub name: String,
    #[serde(
        rename = "{http://www.eurid.eu/xml/epp/domain-ext-2.5}availableDate",
        deserialize_with = "super::deserialize_datetime_opt",
        default
    )]
    pub available_date: Option<DateTime<Utc>>,
    #[serde(rename = "{http://www.eurid.eu/xml/epp/domain-ext-2.5}status", default)]
    pub status: Vec<super::domain::EPPDomainStatus>,
}

#[derive(Debug, Serialize)]
pub struct EURIDDomainCreate {
    #[serde(
        rename = "{http://www.eurid.eu/xml/epp/domain-ext-2.5}domain-ext:contact",
        skip_serializing_if = "Vec::is_empty"
    )]
    pub contacts: Vec<EURIDDomainContact>,
    #[serde(
        rename = "{http://www.eurid.eu/xml/epp/domain-ext-2.5}domain-ext:nsgroup",
        skip_serializing_if = "Vec::is_empty"
    )]
    pub nsgroups: Vec<String>,
    #[serde(
        rename = "{http://www.eurid.eu/xml/epp/domain-ext-2.5}domain-ext:keygroup",
        skip_serializing_if = "Option::is_none"
    )]
    pub keygroup: Option<String>,
    #[serde(
        rename = "{http://www.eurid.eu/xml/epp/domain-ext-2.5}domain-ext:registrarReference",
        skip_serializing_if = "Option::is_none"
    )]
    pub registrar_reference: Option<String>,
}

#[derive(Debug, Deserialize)]
pub struct EURIDDomainInfo {
    #[serde(rename = "{http://www.eurid.eu/xml/epp/domain-ext-2.5}onHold")]
    pub on_hold: bool,
    #[serde(rename = "{http://www.eurid.eu/xml/epp/domain-ext-2.5}reserved")]
    pub reserved: bool,
    #[serde(rename = "{http://www.eurid.eu/xml/epp/domain-ext-2.5}quarantined")]
    pub quarantined: bool,
    #[serde(rename = "{http://www.eurid.eu/xml/epp/domain-ext-2.5}suspended")]
    pub suspended: bool,
    #[serde(rename = "{http://www.eurid.eu/xml/epp/domain-ext-2.5}delayed")]
    pub delayed: bool,
    #[serde(rename = "{http://www.eurid.eu/xml/epp/domain-ext-2.5}seized")]
    pub seized: bool,
    #[serde(
        rename = "{http://www.eurid.eu/xml/epp/domain-ext-2.5}deletionDate",
        deserialize_with = "super::deserialize_datetime_opt",
        default
    )]
    pub deletion_date: Option<DateTime<Utc>>,
    #[serde(
        rename = "{http://www.eurid.eu/xml/epp/domain-ext-2.5}contact",
        default
    )]
    pub contacts: Vec<EURIDDomainContact>,
    #[serde(
        rename = "{http://www.eurid.eu/xml/epp/domain-ext-2.5}nsgroup",
        default
    )]
    pub nsgroups: Vec<String>,
    #[serde(
        rename = "{http://www.eurid.eu/xml/epp/domain-ext-2.5}keygroup",
        default
    )]
    pub keygroup: Option<String>,
    #[serde(rename = "{http://www.eurid.eu/xml/epp/domain-ext-2.5}maxExtensionPeriod")]
    pub max_extension_period: u32,
    #[serde(rename = "{http://www.eurid.eu/xml/epp/domain-ext-2.5}registrantCountry")]
    pub registrant_country: String,
    #[serde(
        rename = "{http://www.eurid.eu/xml/epp/domain-ext-2.5}registrantCountryOfCitizenship",
        default
    )]
    pub registrant_country_of_citizenship: Option<String>,
    #[serde(
        rename = "{http://www.eurid.eu/xml/epp/domain-ext-2.5}registrarReference",
        default
    )]
    pub registrar_reference: Option<String>,
}

#[derive(Debug, Serialize)]
pub struct EURIDDomainUpdate {
    #[serde(
        rename = "{http://www.eurid.eu/xml/epp/domain-ext-2.5}domain-ext:add",
        skip_serializing_if = "Option::is_none"
    )]
    pub add: Option<EURIDDomainUpdateAddRemove>,
    #[serde(
        rename = "{http://www.eurid.eu/xml/epp/domain-ext-2.5}domain-ext:rem",
        skip_serializing_if = "Option::is_none"
    )]
    pub remove: Option<EURIDDomainUpdateAddRemove>,
    #[serde(
        rename = "{http://www.eurid.eu/xml/epp/domain-ext-2.5}domain-ext:chg",
        skip_serializing_if = "Option::is_none"
    )]
    pub change: Option<EURIDDomainUpdateChange>,
}

#[derive(Debug, Serialize)]
pub struct EURIDDomainUpdateAddRemove {
    #[serde(
        rename = "{http://www.eurid.eu/xml/epp/domain-ext-2.5}domain-ext:contact",
        skip_serializing_if = "Vec::is_empty"
    )]
    pub contacts: Vec<EURIDDomainContact>,
    #[serde(
        rename = "{http://www.eurid.eu/xml/epp/domain-ext-2.5}domain-ext:nsgroup",
        skip_serializing_if = "Vec::is_empty"
    )]
    pub nsgroups: Vec<String>,
    #[serde(
        rename = "{http://www.eurid.eu/xml/epp/domain-ext-2.5}domain-ext:keygroup",
        skip_serializing_if = "Option::is_none"
    )]
    pub keygroup: Option<String>,
}

#[derive(Debug, Serialize)]
pub struct EURIDDomainUpdateChange {
    #[serde(
        rename = "{http://www.eurid.eu/xml/epp/domain-ext-2.5}domain-ext:registrarReference",
        skip_serializing_if = "Option::is_none"
    )]
    pub registrar_reference: Option<String>,
}

#[derive(Debug, Deserialize, Serialize)]
pub struct EURIDDomainContact {
    #[serde(rename = "$attr:type")]
    pub contact_type: EURIDContactType,
    #[serde(rename = "$value")]
    pub contact_id: String,
}

#[derive(Debug, Deserialize)]
pub struct EURIDDomainRenewData {
    #[serde(rename = "$value", default)]
    pub data: Vec<EURIDDomainRenewDataType>,
}

#[derive(Debug, Deserialize, PartialEq, Eq)]
pub enum EURIDDomainRenewDataType {
    #[serde(rename = "{http://www.eurid.eu/xml/epp/domain-ext-2.5}removedDeletionDate")]
    RemovedDeletionDate,
}

#[derive(Debug, Serialize)]
pub struct EURIDDomainTransfer {
    #[serde(
        rename = "{http://www.eurid.eu/xml/epp/domain-ext-2.5}domain-ext:request",
        skip_serializing_if = "Option::is_none"
    )]
    pub transfer_request: Option<EURIDDomainTransferRequest>,
}

#[derive(Debug, Serialize)]
pub struct EURIDDomainTransferRequest {
    #[serde(rename = "{http://www.eurid.eu/xml/epp/domain-ext-2.5}domain-ext:registrant")]
    pub registrant: String,
    #[serde(
        rename = "{http://www.eurid.eu/xml/epp/domain-ext-2.5}domain-ext:contact",
        skip_serializing_if = "Vec::is_empty"
    )]
    pub contacts: Vec<EURIDDomainContact>,
    #[serde(
        rename = "{http://www.eurid.eu/xml/epp/domain-ext-2.5}domain-ext:ns",
        skip_serializing_if = "Option::is_none"
    )]
    pub nameservers: Option<super::domain::EPPDomainInfoNameservers>,
    #[serde(
        rename = "{http://www.eurid.eu/xml/epp/domain-ext-2.5}domain-ext:nsgroup",
        skip_serializing_if = "Vec::is_empty"
    )]
    pub nsgroups: Vec<String>,
    #[serde(
        rename = "{http://www.eurid.eu/xml/epp/domain-ext-2.5}domain-ext:keygroup",
        skip_serializing_if = "Option::is_none"
    )]
    pub keygroup: Option<String>,
    #[serde(
        rename = "{http://www.eurid.eu/xml/epp/domain-ext-2.5}domain-ext:registrarReference",
        skip_serializing_if = "Option::is_none"
    )]
    pub registrar_reference: Option<String>,
}

#[derive(Debug, Deserialize)]
pub struct EURIDDomainTransferData {
    #[serde(rename = "{http://www.eurid.eu/xml/epp/domain-ext-2.5}onHold")]
    pub on_hold: bool,
    #[serde(rename = "{http://www.eurid.eu/xml/epp/domain-ext-2.5}reserved")]
    pub reserved: bool,
    #[serde(rename = "{http://www.eurid.eu/xml/epp/domain-ext-2.5}quarantined")]
    pub quarantined: bool,
    #[serde(rename = "{http://www.eurid.eu/xml/epp/domain-ext-2.5}delayed")]
    pub delayed: bool,
    #[serde(rename = "{http://www.eurid.eu/xml/epp/domain-ext-2.5}reason")]
    pub reason: String,
    #[serde(rename = "{http://www.eurid.eu/xml/epp/domain-ext-2.5}registrant")]
    pub registrant: String,
    #[serde(
        rename = "{http://www.eurid.eu/xml/epp/domain-ext-2.5}contact",
        default
    )]
    pub contacts: Vec<EURIDDomainContact>,
    #[serde(rename = "{http://www.eurid.eu/xml/epp/domain-ext-2.5}ns", default)]
    pub nameservers: Option<super::domain::EPPDomainInfoNameservers>,
    #[serde(
        rename = "{http://www.eurid.eu/xml/epp/domain-ext-2.5}nsgroup",
        default
    )]
    pub nsgroups: Vec<String>,
    #[serde(
        rename = "{http://www.eurid.eu/xml/epp/domain-ext-2.5}keygroup",
        default
    )]
    pub keygroup: Option<String>,
    #[serde(
        rename = "{http://www.eurid.eu/xml/epp/domain-ext-2.5}registrarReference",
        default
    )]
    pub registrar_reference: Option<String>,
}

#[derive(Debug, Serialize)]
pub enum EURIDDomainDelete {
    #[serde(rename = "{http://www.eurid.eu/xml/epp/domain-ext-2.5}domain-ext:delete")]
    Schedule(EURIDDomainDeleteSchedule),
    #[serde(rename = "{http://www.eurid.eu/xml/epp/domain-ext-2.5}domain-ext:cancel")]
    Cancel {},
}

#[derive(Debug, Serialize)]
pub struct EURIDDomainDeleteSchedule {
    #[serde(rename = "{http://www.eurid.eu/xml/epp/domain-ext-2.5}domain-ext:delDate")]
    pub delete_date: DateTime<Utc>,
}

/// contact-ext-1.4

#[derive(Debug, Serialize, Deserialize)]
pub struct EURIDContactInfo {
    #[serde(rename = "{http://www.eurid.eu/xml/epp/contact-ext-1.4}contact-ext:type")]
    pub contact_type: EURIDContactType,
    #[serde(
        rename = "{http://www.eurid.eu/xml/epp/contact-ext-1.4}contact-ext:vat",
        skip_serializing_if = "Option::is_none",
        default
    )]
    pub vat: Option<String>,
    #[serde(rename = "{http://www.eurid.eu/xml/epp/contact-ext-1.4}contact-ext:lang")]
    pub language: String,
    #[serde(
        rename = "{http://www.eurid.eu/xml/epp/contact-ext-1.4}contact-ext:whoisEmail",
        skip_serializing_if = "Option::is_none",
        default
    )]
    pub whois_email: Option<String>,
    #[serde(rename = "{http://www.eurid.eu/xml/epp/contact-ext-1.4}contact-ext:naturalPerson")]
    pub natural_person: bool,
    #[serde(
        rename = "{http://www.eurid.eu/xml/epp/contact-ext-1.4}contact-ext:countryOfCitizenship",
        skip_serializing_if = "Option::is_none",
        default
    )]
    pub country_of_citizenship: Option<String>,
    #[serde(
        rename = "{http://www.eurid.eu/xml/epp/contact-ext-1.4}contact-ext:hasReservedDomain",
        skip_serializing,
        default
    )]
    pub has_reserved_domain: bool,
    #[serde(
        rename = "{http://www.eurid.eu/xml/epp/contact-ext-1.4}contact-ext:hasDomainOnHold",
        skip_serializing,
        default
    )]
    pub has_domain_on_hold: bool,
    #[serde(
        rename = "{http://www.eurid.eu/xml/epp/contact-ext-1.4}contact-ext:hasDomainRegistryLock",
        skip_serializing,
        default
    )]
    pub has_domain_registry_lock: bool,
    #[serde(
        rename = "{http://www.eurid.eu/xml/epp/contact-ext-1.4}contact-ext:hasDomainRegistrarLock",
        skip_serializing,
        default
    )]
    pub has_domain_registrar_lock: bool,
}

#[derive(Debug, Serialize)]
pub struct EURIDContactUpdate {
    #[serde(rename = "{http://www.eurid.eu/xml/epp/contact-ext-1.4}contact-ext:chg")]
    pub change: EURIDContactUpdateInfo,
}

#[derive(Debug, Serialize)]
pub struct EURIDContactUpdateInfo {
    #[serde(
        rename = "{http://www.eurid.eu/xml/epp/contact-ext-1.4}contact-ext:whoisEmial",
        skip_serializing_if = "Option::is_none"
    )]
    pub whois_email: Option<String>,
    #[serde(
        rename = "{http://www.eurid.eu/xml/epp/contact-ext-1.4}contact-ext:vat",
        skip_serializing_if = "Option::is_none"
    )]
    pub vat: Option<String>,
    #[serde(
        rename = "{http://www.eurid.eu/xml/epp/contact-ext-1.4}contact-ext:lang",
        skip_serializing_if = "Option::is_none"
    )]
    pub language: Option<String>,
    #[serde(
        rename = "{http://www.eurid.eu/xml/epp/contact-ext-1.4}contact-ext:naturalPerson",
        skip_serializing_if = "Option::is_none"
    )]
    pub natural_person: Option<bool>,
    #[serde(
        rename = "{http://www.eurid.eu/xml/epp/contact-ext-1.4}contact-ext:countryOfCitizenship",
        skip_serializing_if = "Option::is_none"
    )]
    pub country_of_citizenship: Option<String>,
}

#[derive(Debug, Serialize, Deserialize)]
pub enum EURIDContactType {
    #[serde(rename = "billing")]
    Billing,
    #[serde(rename = "tech")]
    Tech,
    #[serde(rename = "registrant")]
    Registrant,
    #[serde(rename = "onsite")]
    OnSite,
    #[serde(rename = "reseller")]
    Reseller,
}

/// dnsQuality-2.0

#[derive(Debug, Serialize)]
pub struct EURIDDNSQualityInfo {
    #[serde(rename = "{http://www.eurid.eu/xml/epp/dnsQuality-2.0}dnsQuality:name")]
    pub name: String,
}

#[derive(Debug, Deserialize)]
pub struct EURIDDNSQualityInfoData {
    #[serde(rename = "{http://www.eurid.eu/xml/epp/dnsQuality-2.0}name")]
    pub name: String,
    #[serde(
        rename = "{http://www.eurid.eu/xml/epp/dnsQuality-2.0}checkTime",
        deserialize_with = "super::deserialize_datetime_opt",
        default
    )]
    pub check_time: Option<DateTime<Utc>>,
    #[serde(rename = "{http://www.eurid.eu/xml/epp/dnsQuality-2.0}score")]
    pub score: String,
}

/// dnssecEligibility-1.0

#[derive(Debug, Serialize)]
pub struct EURIDDNSSECEligibilityInfo {
    #[serde(rename = "{http://www.eurid.eu/xml/epp/dnssecEligibility-1.0}dnssecEligibility:name")]
    pub name: String,
}

#[derive(Debug, Deserialize)]
pub struct EURIDDNSSECEligibilityInfoData {
    #[serde(rename = "{http://www.eurid.eu/xml/epp/dnssecEligibility-1.0}name")]
    pub name: String,
    #[serde(rename = "{http://www.eurid.eu/xml/epp/dnssecEligibility-1.0}eligible")]
    pub eligible: bool,
    #[serde(rename = "{http://www.eurid.eu/xml/epp/dnssecEligibility-1.0}msg")]
    pub msg: String,
    #[serde(rename = "{http://www.eurid.eu/xml/epp/dnssecEligibility-1.0}code")]
    pub code: u32,
}

/// homoglyph-1.0

#[derive(Debug, Deserialize)]
pub struct EURIDHomoglyphData {
    #[serde(rename = "{http://www.eurid.eu/xml/epp/homoglyph-1.0}domain")]
    pub domains: Vec<EURIDHomoglyphDomainData>,
}

#[derive(Debug, Deserialize)]
pub struct EURIDHomoglyphDomainData {
    #[serde(rename = "{http://www.eurid.eu/xml/epp/homoglyph-1.0}name")]
    pub unicode: String,
    #[serde(rename = "{http://www.eurid.eu/xml/epp/homoglyph-1.0}blockedBy")]
    pub blocked_by: Vec<String>,
}

/// authInfo-1.1

#[derive(Debug, Serialize)]
pub enum EURIDAuthInfo {
    #[serde(rename = "{http://www.eurid.eu/xml/epp/authInfo-1.1}authInfo:request")]
    Request {},
    #[serde(rename = "{http://www.eurid.eu/xml/epp/authInfo-1.1}authInfo:cancel")]
    Cancel {},
}

#[derive(Debug, Deserialize)]
pub struct EURIDAuthInfoData {
    #[serde(
        rename = "{http://www.eurid.eu/xml/epp/authInfo-1.1}validUntil",
        deserialize_with = "super::deserialize_datetime"
    )]
    pub valid_until: DateTime<Utc>,
}

/// idn-1.0

#[derive(Debug, Deserialize)]
pub struct EURIDIDNMapping {
    #[serde(rename = "{http://www.eurid.eu/xml/epp/idn-1.0}name")]
    pub names: Vec<EURIDIDNNameMapping>,
}

#[derive(Debug, Deserialize)]
pub struct EURIDIDNNameMapping {
    #[serde(rename = "{http://www.eurid.eu/xml/epp/idn-1.0}ace")]
    pub ace: String,
    #[serde(rename = "{http://www.eurid.eu/xml/epp/idn-1.0}unicode")]
    pub unicode: String,
}

/// registrarFinance-1.0

#[derive(Debug, Deserialize)]
pub struct EURIDRegistrarFinanceInfoData {
    #[serde(rename = "{http://www.eurid.eu/xml/epp/registrarFinance-1.0}paymentMode")]
    pub payment_mode: EURIDRegistrarFinancePaymentMode,
    #[serde(rename = "{http://www.eurid.eu/xml/epp/registrarFinance-1.0}accountBalance")]
    pub account_balance: String,
    #[serde(rename = "{http://www.eurid.eu/xml/epp/registrarFinance-1.0}availableAmount")]
    pub available_amount: Option<String>,
    #[serde(rename = "{http://www.eurid.eu/xml/epp/registrarFinance-1.0}dueAmount")]
    pub due_amount: Option<String>,
    #[serde(rename = "{http://www.eurid.eu/xml/epp/registrarFinance-1.0}overdueAmount")]
    pub overdue_amount: Option<String>,
}

#[derive(Debug, Deserialize)]
pub enum EURIDRegistrarFinancePaymentMode {
    #[serde(rename = "PRE_PAYMENT")]
    PrePayment,
    #[serde(rename = "POST_PAYMENT")]
    PostPayment,
}

/// registrarHitPoints-1.0

#[derive(Debug, Deserialize)]
pub struct EURIDRegistrarHitPointsInfoData {
    #[serde(rename = "{http://www.eurid.eu/xml/epp/registrarHitPoints-1.0}nbrHitPoints")]
    pub hit_points: u64,
    #[serde(rename = "{http://www.eurid.eu/xml/epp/registrarHitPoints-1.0}maxNbrHitPoints")]
    pub max_hit_points: u64,
    #[serde(
        rename = "{http://www.eurid.eu/xml/epp/registrarHitPoints-1.0}blockedUntil",
        deserialize_with = "super::deserialize_datetime_opt",
        default
    )]
    pub blocked_until: Option<DateTime<Utc>>,
}

/// registrationLimit-1.1

#[derive(Debug, Deserialize)]
pub struct EURIDRegistrationLimitInfoData {
    #[serde(rename = "{http://www.eurid.eu/xml/epp/registrationLimit-1.1}monthlyRegistrations")]
    pub monthly_registrations: u64,
    #[serde(
        rename = "{http://www.eurid.eu/xml/epp/registrationLimit-1.1}maxMonthlyRegistrations",
        default
    )]
    pub max_monthly_registrations: Option<u64>,
    #[serde(
        rename = "{http://www.eurid.eu/xml/epp/registrationLimit-1.1}limitedUntil",
        deserialize_with = "super::deserialize_datetime_opt",
        default
    )]
    pub limited_until: Option<DateTime<Utc>>,
}

/// poll-1.2

#[derive(Debug, Deserialize)]
pub struct EURIDPollData {
    #[serde(rename = "{http://www.eurid.eu/xml/epp/poll-1.2}context")]
    pub context: String,
    #[serde(rename = "{http://www.eurid.eu/xml/epp/poll-1.2}objectType")]
    pub object_type: String,
    #[serde(rename = "{http://www.eurid.eu/xml/epp/poll-1.2}object")]
    pub object: String,
    #[serde(
        rename = "{http://www.eurid.eu/xml/epp/poll-1.2}objectUnicode",
        default
    )]
    pub object_unicode: Option<String>,
    #[serde(rename = "{http://www.eurid.eu/xml/epp/poll-1.2}action")]
    pub action: String,
    #[serde(rename = "{http://www.eurid.eu/xml/epp/poll-1.2}code")]
    pub code: u32,
    #[serde(rename = "{http://www.eurid.eu/xml/epp/poll-1.2}detail", default)]
    pub detail: Option<String>,
    #[serde(rename = "{http://www.eurid.eu/xml/epp/poll-1.2}registrar", default)]
    pub registrar: Option<String>,
}

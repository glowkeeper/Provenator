import { keccak256 } from 'js-sha3';

class PremisRightsHandler {

  static actOptions = ['Replicate','Migrate','Modify','Use','Disseminate','Delete']
  static restrictionsOptions = ['Allow','Disallow', 'Conditional']
  static basisOptions = ['Copyright','Statute','License','Donor','Policy','Other','Acts granted or restricted']

  static jurisdictionCountryCodes = ['AD',
  'AE',
  'AF',
  'AG',
  'AI',
  'AL',
  'AM',
  'AO',
  'AQ',
  'AR',
  'AS',
  'AT',
  'AU',
  'AW',
  'AX',
  'AZ',
  'BA',
  'BB',
  'BD',
  'BE',
  'BF',
  'BG',
  'BH',
  'BI',
  'BJ',
  'BL',
  'BM',
  'BN',
  'BO',
  'BQ',
  'BR',
  'BS',
  'BT',
  'BV',
  'BW',
  'BY',
  'BZ',
  'CA',
  'CC',
  'CD',
  'CF',
  'CG',
  'CH',
  'CI',
  'CK',
  'CL',
  'CM',
  'CN',
  'CO',
  'CR',
  'CU',
  'CV',
  'CW',
  'CX',
  'CY',
  'CZ',
  'DE',
  'DJ',
  'DK',
  'DM',
  'DO',
  'DZ',
  'EC',
  'EE',
  'EG',
  'EH',
  'ER',
  'ES',
  'ET',
  'FI',
  'FJ',
  'FK',
  'FM',
  'FO',
  'FR',
  'GA',
  'GB',
  'GD',
  'GE',
  'GF',
  'GG',
  'GH',
  'GI',
  'GL',
  'GM',
  'GN',
  'GP',
  'GQ',
  'GR',
  'GS',
  'GT',
  'GU',
  'GW',
  'GY',
  'HK',
  'HM',
  'HN',
  'HR',
  'HT',
  'HU',
  'ID',
  'IE',
  'IL',
  'IM',
  'IN',
  'IO',
  'IQ',
  'IR',
  'IS',
  'IT',
  'JE',
  'JM',
  'JO',
  'JP',
  'KE',
  'KG',
  'KH',
  'KI',
  'KM',
  'KN',
  'KP',
  'KR',
  'KW',
  'KY',
  'KZ',
  'LA',
  'LB',
  'LC',
  'LI',
  'LK',
  'LR',
  'LS',
  'LT',
  'LU',
  'LV',
  'LY',
  'MA',
  'MC',
  'MD',
  'ME',
  'MF',
  'MG',
  'MH',
  'MK',
  'ML',
  'MM',
  'MN',
  'MO',
  'MP',
  'MQ',
  'MR',
  'MS',
  'MT',
  'MU',
  'MV',
  'MW',
  'MX',
  'MY',
  'MZ',
  'NA',
  'NC',
  'NE',
  'NF',
  'NG',
  'NI',
  'NL',
  'NO',
  'NP',
  'NR',
  'NU',
  'NZ',
  'OM',
  'PA',
  'PE',
  'PF',
  'PG',
  'PH',
  'PK',
  'PL',
  'PM',
  'PN',
  'PR',
  'PS',
  'PT',
  'PW',
  'PY',
  'QA',
  'RE',
  'RO',
  'RS',
  'RU',
  'RW',
  'SA',
  'SB',
  'SC',
  'SD',
  'SE',
  'SG',
  'SH',
  'SI',
  'SJ',
  'SK',
  'SL',
  'SM',
  'SN',
  'SO',
  'SR',
  'SS',
  'ST',
  'SV',
  'SX',
  'SY',
  'SZ',
  'TC',
  'TD',
  'TF',
  'TG',
  'TH',
  'TJ',
  'TK',
  'TL',
  'TM',
  'TN',
  'TO',
  'TR',
  'TT',
  'TV',
  'TW',
  'TZ',
  'UA',
  'UG',
  'UM',
  'US',
  'UY',
  'UZ',
  'VA',
  'VC',
  'VE',
  'VG',
  'VI',
  'VN',
  'VU',
  'WF',
  'WS',
  'YE',
  'YT',
  'ZA',
  'ZM',
  'ZW']

  constructor () {

    this.rights = {
      id: undefined,
      basis: undefined,
      status: undefined,
      countryCode: undefined,
      determinationDate: undefined,
      note: undefined,
      act: undefined,
      restriction: undefined
    }
  }

  reset () {
    this.rights.id = undefined
    this.rights.basis = undefined
    this.rights.status = undefined
    this.rights.countryCode = undefined
    this.rights.determinationDate = undefined
    this.rights.note = undefined
    this.rights.act = undefined
    this.rights.restriction = undefined
  }

  checkSet () {
    if ((this.rights.id === undefined) ||
        (this.rights.basis === undefined) ||
        (this.rights.status === undefined) ||
        (this.rights.countryCode === undefined) ||
        (this.rights.determinationDate === undefined) ||
        (this.rights.note === undefined) ||
        (this.rights.act === undefined) ||
        (this.rights.restriction === undefined)) {
      return false
    } else {
      return true
    }
  }

  setId () {
    const keyString = this.rights.id +
                      this.rights.basis +
                      this.rights.status +
                      this.rights.countryCode +
                      this.rights.determinationDate +
                      this.rights.act + this.rights.restriction
    const key = keccak256(keyString)
    this.rights.id = key;
  }

  setBasis (_basis) {
    this.rights.basis = _basis
  }

  setStatus (_status) {
    this.rights.status = _status
  }

  setCountryCode (_code) {
    this.rights.countryCode = _code
  }

  setDeterminationDate (_date) {
    this.rights.determinationDate = _date
  }

  setNote (_note) {
    this.rights.note = _note
  }

  setAct (_act) {
    this.rights.act = _act
  }

  setRestriction (_restriction) {
    this.rights.restriction = _restriction
  }

  getId () {
    return this.rights.id;
  }

  getBasis () {
    return this.rights.basis
  }

  getStatus () {
    return this.rights.status
  }

  getCountryCode () {
    return this.rights.countryCode
  }

  getDeterminationdate () {
    return this.rights.determinationDate
  }

  getNote () {
    return this.rights.note
  }

  getAct () {
    return this.rights.act
  }

  getRestriction () {
    return this.rights.restriction
  }

}

export default PremisRightsHandler

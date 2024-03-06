const HL7 = require("hl7-standard");

// Hardcoded HL7 message
const hl7Message = `MSH|^~\\&|HIS|XYZHOSPITAL|LABADT|XYZHOSPITAL|198808181126|SECURITY|ADT^A01|MSG00001|P|2.2
EVN|A01|198808181123
PID|||PATID1234^5^M11||JONES^WILLIAM^A^III||19610615|M-||C|1200 N ELM STREET^^GREENSBORO^NC^27401-1020|GL|(919)379-1212|(919)271-3434||S||PATID12345001^2^M10|123456789|9-87654^NC
NK1|1|JONES^BARBARA^K|WIFE||||||NK^NEXT OF KIN
PV1|1|I|2000^2012^01||||004777^LEBAUER^SIDNEY^J.|||SUR||||ADM|A0|`;

let hl7 = new HL7(hl7Message);
hl7.transform((err) => {
  if (err) throw err;
  // code here

  let familyName = hl7.get("PID.5.1");

  let patientLanguage = hl7.get("PID.15");
  let address = hl7.get("PID.11");

  console.log(familyName, patientLanguage, address);
});

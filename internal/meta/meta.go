// Copyright 2021 Adam Chalkley
//
// https://github.com/atc0005/query-meta
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package meta

import (
	"fmt"

	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
)

// String implements the Stringer interface in order to expose all fields
// (other than SSN) in their unmodified form.
func (r LibraryPatronRecord) String() string {

	return fmt.Sprintf(
		"{ "+
			"GIDStatus: %v, "+
			"GID: %v, "+
			"IDNum: %v, "+
			"SprIdenID: %v, "+
			"SSN: %v, "+
			"GivenName: %v, "+
			"NameMiddle: %v, "+
			"Surname: %v, "+
			"LibIssueNumber: %v, "+
			"EmpStudentFlag: %v, "+
			"StudentStatus: %v, "+
			"BannerLastTermEnrolled: %v, "+
			"StudentCurrentlyEnrolled: %v, "+
			"StudentLevelCode: %v, "+
			"EmployeeStatus: %v, "+
			"EmployeeType: %v, "+
			"SalaryTable: %v, "+
			"EmployeeGroupCode: %v, "+
			"TESFlag: %v, "+
			"CampusCode: %v, "+
			"EmployeeTerminationDate: %v, "+
			"OU: %v, "+
			"Title: %v, "+
			"Location: %v, "+
			"AddressDeptCity: %v, "+
			"AddressDeptState: %v, "+
			"AddressDeptZip: %v, "+
			"TelephoneNumber: %v, "+
			"AddressStudentPermStreet1: %v, "+
			"AddressStudentPermStreet2: %v, "+
			"AddressStudentPermStreet3: %v, "+
			"AddressStudentPermCity: %v, "+
			"AddressStudentPermState: %v, "+
			"AddressStudentPermZip: %v, "+
			"AddressNationCode: %v, "+
			"PhoneStudentPerm: %v, "+
			"StudentType: %v, "+
			"OtherStatus: %v, "+
			"OtherTypeCode: %v, "+
			"StudentDegreeEarned: %v"+
			" }",
		r.GIDStatus.String,
		r.GID.String,
		r.IDNum.String,
		r.SprIdenID.String,
		"REDACTED", // r.SSN.String,
		r.GivenName.String,
		r.NameMiddle.String,
		r.Surname.String,
		r.LibIssueNumber.String,
		r.EmpStudentFlag.String,
		r.StudentStatus.String,
		r.BannerLastTermEnrolled.String,
		r.StudentCurrentlyEnrolled.String,
		r.StudentLevelCode.String,
		r.EmployeeStatus.String,
		r.EmployeeType.String,
		r.SalaryTable.String,
		r.EmployeeGroupCode.String,
		r.TESFlag.String,
		r.CampusCode.String,
		r.EmployeeTerminationDate.String,
		r.OU.String,
		r.Title.String,
		r.Location.String,
		r.AddressDeptCity.String,
		r.AddressDeptState.String,
		r.AddressDeptZip.String,
		r.TelephoneNumber.String,
		r.AddressStudentPermStreet1.String,
		r.AddressStudentPermStreet2.String,
		r.AddressStudentPermStreet3.String,
		r.AddressStudentPermCity.String,
		r.AddressStudentPermState.String,
		r.AddressStudentPermZip.String,
		r.AddressNationCode.String,
		r.PhoneStudentPerm.String,
		r.StudentType.String,
		r.OtherStatus.String,
		r.OtherTypeCode.String,
		r.StudentDegreeEarned.String,
	)

}

// padEmpty replaces any empty, non-NULL field values with a single space to
// emulate the behavior of the original query_meta.php script.
func (r *LibraryPatronRecord) padEmpty() {

	if r.GIDStatus.Valid && r.GIDStatus.String == "" {
		r.GIDStatus.String = " "
	}

	if r.GID.Valid && r.GID.String == "" {
		r.GID.String = " "
	}

	if r.IDNum.Valid && r.IDNum.String == "" {
		r.IDNum.String = " "
	}

	if r.SprIdenID.Valid && r.SprIdenID.String == "" {
		r.SprIdenID.String = " "
	}

	if r.SSN.Valid && r.SSN.String == "" {
		r.SSN.String = " "
	}

	if r.GivenName.Valid && r.GivenName.String == "" {
		r.GivenName.String = " "
	}

	if r.NameMiddle.Valid && r.NameMiddle.String == "" {
		r.NameMiddle.String = " "
	}

	if r.Surname.Valid && r.Surname.String == "" {
		r.Surname.String = " "
	}

	if r.LibIssueNumber.Valid && r.LibIssueNumber.String == "" {
		r.LibIssueNumber.String = " "
	}

	if r.EmpStudentFlag.Valid && r.EmpStudentFlag.String == "" {
		r.EmpStudentFlag.String = " "
	}

	if r.StudentStatus.Valid && r.StudentStatus.String == "" {
		r.StudentStatus.String = " "
	}

	if r.BannerLastTermEnrolled.Valid && r.BannerLastTermEnrolled.String == "" {
		r.BannerLastTermEnrolled.String = " "
	}

	if r.StudentCurrentlyEnrolled.Valid && r.StudentCurrentlyEnrolled.String == "" {
		r.StudentCurrentlyEnrolled.String = " "
	}

	if r.StudentLevelCode.Valid && r.StudentLevelCode.String == "" {
		r.StudentLevelCode.String = " "
	}

	if r.EmployeeStatus.Valid && r.EmployeeStatus.String == "" {
		r.EmployeeStatus.String = " "
	}

	if r.EmployeeType.Valid && r.EmployeeType.String == "" {
		r.EmployeeType.String = " "
	}

	if r.SalaryTable.Valid && r.SalaryTable.String == "" {
		r.SalaryTable.String = " "
	}

	if r.EmployeeGroupCode.Valid && r.EmployeeGroupCode.String == "" {
		r.EmployeeGroupCode.String = " "
	}

	if r.TESFlag.Valid && r.TESFlag.String == "" {
		r.TESFlag.String = " "
	}

	if r.CampusCode.Valid && r.CampusCode.String == "" {
		r.CampusCode.String = " "
	}

	if r.EmployeeTerminationDate.Valid && r.EmployeeTerminationDate.String == "" {
		r.EmployeeTerminationDate.String = " "
	}

	if r.OU.Valid && r.OU.String == "" {
		r.OU.String = " "
	}

	if r.Title.Valid && r.Title.String == "" {
		r.Title.String = " "
	}

	if r.Location.Valid && r.Location.String == "" {
		r.Location.String = " "
	}

	if r.AddressDeptCity.Valid && r.AddressDeptCity.String == "" {
		r.AddressDeptCity.String = " "
	}

	if r.AddressDeptState.Valid && r.AddressDeptState.String == "" {
		r.AddressDeptState.String = " "
	}

	if r.AddressDeptZip.Valid && r.AddressDeptZip.String == "" {
		r.AddressDeptZip.String = " "
	}

	if r.TelephoneNumber.Valid && r.TelephoneNumber.String == "" {
		r.TelephoneNumber.String = " "
	}

	if r.AddressStudentPermStreet1.Valid && r.AddressStudentPermStreet1.String == "" {
		r.AddressStudentPermStreet1.String = " "
	}

	if r.AddressStudentPermStreet2.Valid && r.AddressStudentPermStreet2.String == "" {
		r.AddressStudentPermStreet2.String = " "
	}

	if r.AddressStudentPermStreet3.Valid && r.AddressStudentPermStreet3.String == "" {
		r.AddressStudentPermStreet3.String = " "
	}

	if r.AddressStudentPermCity.Valid && r.AddressStudentPermCity.String == "" {
		r.AddressStudentPermCity.String = " "
	}

	if r.AddressStudentPermState.Valid && r.AddressStudentPermState.String == "" {
		r.AddressStudentPermState.String = " "
	}

	if r.AddressStudentPermZip.Valid && r.AddressStudentPermZip.String == "" {
		r.AddressStudentPermZip.String = " "
	}

	if r.AddressNationCode.Valid && r.AddressNationCode.String == "" {
		r.AddressNationCode.String = " "
	}

	if r.PhoneStudentPerm.Valid && r.PhoneStudentPerm.String == "" {
		r.PhoneStudentPerm.String = " "
	}

	if r.StudentType.Valid && r.StudentType.String == "" {
		r.StudentType.String = " "
	}

	if r.OtherStatus.Valid && r.OtherStatus.String == "" {
		r.OtherStatus.String = " "
	}

	if r.OtherTypeCode.Valid && r.OtherTypeCode.String == "" {
		r.OtherTypeCode.String = " "
	}

	if r.StudentDegreeEarned.Valid && r.StudentDegreeEarned.String == "" {
		r.StudentDegreeEarned.String = " "
	}

}

// CSV produces a patron record in a pipe delimited, view_extract.txt file
// format. Further transformation and or encoding is needed before writing
// this entry to an output file for further processing by later stage tooling.
func (r LibraryPatronRecord) CSV() string {
	// mimic functionality from original query_meta.php script by
	// hand-crafting pipe-delimited output.
	//
	// TODO: After sufficient testing has been performed, switch to stdlib CSV
	// module for generating output?

	// Original script:
	//
	// NULL values in field seem to map to empty value
	// empty string in field seem to map to single space

	// We emulate the original script behavior by replacing empty (not NULL)
	// field values with a single space.
	r.padEmpty()

	outputRecordTmpl := "%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s\n"

	return fmt.Sprintf(
		outputRecordTmpl,
		r.GIDStatus.String,
		r.GID.String,
		r.IDNum.String,
		r.SprIdenID.String,
		r.SSN.String,
		r.GivenName.String,
		r.NameMiddle.String,
		r.Surname.String,
		r.LibIssueNumber.String,
		r.EmpStudentFlag.String,
		r.StudentStatus.String,
		r.BannerLastTermEnrolled.String,
		r.StudentCurrentlyEnrolled.String,
		r.StudentLevelCode.String,
		r.EmployeeStatus.String,
		r.EmployeeType.String,
		r.SalaryTable.String,
		r.EmployeeGroupCode.String,
		r.TESFlag.String,
		r.CampusCode.String,
		r.EmployeeTerminationDate.String,
		r.OU.String,
		r.Title.String,
		r.Location.String,
		r.AddressDeptCity.String,
		r.AddressDeptState.String,
		r.AddressDeptZip.String,
		r.TelephoneNumber.String,
		r.AddressStudentPermStreet1.String,
		r.AddressStudentPermStreet2.String,
		r.AddressStudentPermStreet3.String,
		r.AddressStudentPermCity.String,
		r.AddressStudentPermState.String,
		r.AddressStudentPermZip.String,
		r.AddressNationCode.String,
		r.PhoneStudentPerm.String,
		r.StudentType.String,
		r.OtherStatus.String,
		r.OtherTypeCode.String,
		r.StudentDegreeEarned.String,
	)
}

// Extract produces a patron record in the expected format for generating a
// (pipe delimited) view_extract.txt file entry. As part of this process,
// UTF-8 to ISO 8859-1 (Latin-1) transformation is performed for required
// compatibility with the Voyager SIF input format. This transformed patron
// record is suitable for use in the output file ingested by other tools in
// the patron records processing pipeline.
func (r LibraryPatronRecord) Extract() (string, int, error) {

	// Encode as ISO 8859-1 (Latin-1), replacing any unsupported characters
	// with the encoding.ASCIISub substitute character.
	latin1Transformer := encoding.ReplaceUnsupported(charmap.ISO8859_1.NewEncoder())

	res, n, err := transform.String(latin1Transformer, r.CSV())

	return res, n, err

}

// Copyright 2021 Adam Chalkley
//
// https://github.com/atc0005/query-meta
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package meta

import "database/sql"

// LibraryPatronRecord represents fields from a view from the `Meta` database.
type LibraryPatronRecord struct {
	GIDStatus                 sql.NullString `arg:"-"`
	GID                       sql.NullString `arg:"-"`
	IDNum                     sql.NullString `arg:"-"`
	SprIdenID                 sql.NullString `arg:"-"`
	SSN                       sql.NullString `arg:"-"`
	GivenName                 sql.NullString `arg:"-"`
	NameMiddle                sql.NullString `arg:"-"`
	Surname                   sql.NullString `arg:"-"`
	LibIssueNumber            sql.NullString `arg:"-"`
	EmpStudentFlag            sql.NullString `arg:"-"`
	StudentStatus             sql.NullString `arg:"-"`
	BannerLastTermEnrolled    sql.NullString `arg:"-"`
	StudentCurrentlyEnrolled  sql.NullString `arg:"-"`
	StudentLevelCode          sql.NullString `arg:"-"`
	EmployeeStatus            sql.NullString `arg:"-"`
	EmployeeType              sql.NullString `arg:"-"`
	SalaryTable               sql.NullString `arg:"-"`
	EmployeeGroupCode         sql.NullString `arg:"-"`
	TESFlag                   sql.NullString `arg:"-"`
	CampusCode                sql.NullString `arg:"-"`
	EmployeeTerminationDate   sql.NullString `arg:"-"`
	OU                        sql.NullString `arg:"-"`
	Title                     sql.NullString `arg:"-"`
	Location                  sql.NullString `arg:"-"`
	AddressDeptCity           sql.NullString `arg:"-"`
	AddressDeptState          sql.NullString `arg:"-"`
	AddressDeptZip            sql.NullString `arg:"-"`
	TelephoneNumber           sql.NullString `arg:"-"`
	AddressStudentPermStreet1 sql.NullString `arg:"-"`
	AddressStudentPermStreet2 sql.NullString `arg:"-"`
	AddressStudentPermStreet3 sql.NullString `arg:"-"`
	AddressStudentPermCity    sql.NullString `arg:"-"`
	AddressStudentPermState   sql.NullString `arg:"-"`
	AddressStudentPermZip     sql.NullString `arg:"-"`
	AddressNationCode         sql.NullString `arg:"-"`
	PhoneStudentPerm          sql.NullString `arg:"-"`
	StudentType               sql.NullString `arg:"-"`
	OtherStatus               sql.NullString `arg:"-"`
	OtherTypeCode             sql.NullString `arg:"-"`
	StudentDegreeEarned       sql.NullString `arg:"-"`
}

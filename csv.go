package gotable

import (
	"fmt"

	"github.com/dustin/go-humanize"
)

// cellSep used to separate csv cells
var cellSep = ","

// SprintTableCSV return the table header in csv layout
func (t *Table) SprintTableCSV(f int) (string, error) {
	// get headers first
	s, err := t.SprintColHdrsCSV()
	if err != nil {
		return "", err
	}

	// then append table body
	rs, err := t.SprintRows(f)
	if err != nil {
		return "", err
	}
	s += rs

	// finally return CSS table layout
	// fmt.Println(strings.Replace(s, "\\\"", "'", -1))
	return s, nil
}

// SprintColHdrsCSV return the table header in csv layout
func (t *Table) SprintColHdrsCSV() (string, error) {
	tHeader := ""
	for i := 0; i < len(t.ColDefs); i++ {
		// quote string with "%q"
		tHeader += fmt.Sprintf("%q", t.ColDefs[i].ColTitle) + cellSep
	}
	// remove last cellSep characters
	tHeader = tHeader[0:len(tHeader)-len(cellSep)] + "\n"
	return tHeader, nil
}

// SprintRowsCSV returns the table rows in csv layout
func (t *Table) SprintRowsCSV(f int) (string, error) {
	rowsStr := ""
	for i := 0; i < t.Rows(); i++ {
		s, err := t.SprintRow(i, f)
		if err != nil {
			return "", err
		}
		rowsStr += s
	}
	return rowsStr, nil
}

// SprintRowCSV return a table row in csv layout
func (t *Table) SprintRowCSV(row int) (string, error) {
	tRow := ""

	// fill the content in rowTextList for the first line
	for i := 0; i < len(t.Row[row].Col); i++ {
		var cellStr string
		// append content in TD
		switch t.Row[row].Col[i].Type {
		case CELLFLOAT:
			cellStr = fmt.Sprintf(t.ColDefs[i].Pfmt, humanize.FormatFloat("#,###.##", t.Row[row].Col[i].Fval))
		case CELLINT:
			cellStr = fmt.Sprintf(t.ColDefs[i].Pfmt, t.Row[row].Col[i].Ival)
		case CELLSTRING:
			// FOR CSV, APPEND FULL STRING, THERE ARE NO MULTILINE STRING IN THIS
			cellStr = fmt.Sprintf("%q", t.Row[row].Col[i].Sval)
		case CELLDATE:
			cellStr = fmt.Sprintf("%*.*s", t.ColDefs[i].Width, t.ColDefs[i].Width, t.Row[row].Col[i].Dval.Format(t.DateFmt))
		case CELLDATETIME:
			cellStr = fmt.Sprintf("%*.*s", t.ColDefs[i].Width, t.ColDefs[i].Width, t.Row[row].Col[i].Dval.Format(t.DateTimeFmt))
		default:
			cellStr = mkstr(t.ColDefs[i].Width, ' ')
		}
		tRow += cellStr + cellSep
	}
	// remove last cellSep characters
	tRow = tRow[0:len(tRow)-len(cellSep)] + "\n"
	return tRow, nil
}

package gotable

// SprintTablePDF return the table header in PDF layout
func (t *Table) SprintTablePDF(f int) (string, error) {
	return "", ErrPDF
}

// SprintColHdrsPDF return the table header in PDF layout
func (t *Table) SprintColHdrsPDF() (string, error) {
	return "", ErrPDF
}

// SprintRowsPDF returns the table rows in PDF layout
func (t *Table) SprintRowsPDF(f int) (string, error) {
	return "", ErrPDF
}

// SprintRowPDF return a table row in PDF layout
func (t *Table) SprintRowPDF(row int) (string, error) {
	return "", ErrPDF
}

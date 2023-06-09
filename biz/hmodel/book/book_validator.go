// Code generated by Validator v0.1.4. DO NOT EDIT.

package book

import (
	"bytes"
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"time"
)

// unused protection
var (
	_ = fmt.Formatter(nil)
	_ = (*bytes.Buffer)(nil)
	_ = (*strings.Builder)(nil)
	_ = reflect.Type(nil)
	_ = (*regexp.Regexp)(nil)
	_ = time.Nanosecond
)

func (p *Book) IsValid() error {
	return nil
}
func (p *ListBooksRequest) IsValid() error {
	return nil
}
func (p *ListBooksResponse) IsValid() error {
	return nil
}
func (p *GetBookRequest) IsValid() error {
	return nil
}
func (p *CreateBookRequest) IsValid() error {
	if p.Book != nil {
		if err := p.Book.IsValid(); err != nil {
			return fmt.Errorf("field Book not valid, %w", err)
		}
	}
	return nil
}
func (p *UpdateBookRequest) IsValid() error {
	if p.Book != nil {
		if err := p.Book.IsValid(); err != nil {
			return fmt.Errorf("field Book not valid, %w", err)
		}
	}
	return nil
}
func (p *RenameBookRequest) IsValid() error {
	return nil
}
func (p *RenameBookResponse) IsValid() error {
	if p.Book != nil {
		if err := p.Book.IsValid(); err != nil {
			return fmt.Errorf("field Book not valid, %w", err)
		}
	}
	return nil
}
func (p *DeleteBookRequest) IsValid() error {
	return nil
}

package main

import "strings"

type StringValue []string

func (i *StringValue) String() string {
	return strings.Join(*i,",")
}

func (i *StringValue) Set(value string) error {
	*i = append(*i, value)
	return nil
}
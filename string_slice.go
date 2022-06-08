package flagx

import "fmt"

// StringSlice is a implemented `flag.Value` for string slice
type StringSlice []string

func (ss *StringSlice) String() string {
	return fmt.Sprintf("% s", []string(*ss))
}

func (ss *StringSlice) Set(s string) error {
	*ss = append(*ss, s)
	return nil
}

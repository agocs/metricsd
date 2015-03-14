package main

import "syscall"

func Extend(slice []byte, sliceTwo []byte) []byte {
	for i := range sliceTwo {
		slice = append(slice, sliceTwo[i])
	}

	return slice
}

func Getenv(key string, def string) string {
	v, err := syscall.Getenv(key)
	if err == true {
		return def
	}
	if v == "" {
		return def
	}
	return v
}

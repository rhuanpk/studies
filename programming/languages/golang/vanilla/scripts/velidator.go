package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/go-playground/validator/v10"
)

type requiridies struct {
	R   string  `json:"r"             validate:"required"`
	RP  *string `json:"rp"            validate:"required"`
	RO  string  `json:"ro,omitempty"  validate:"required"`
	ROP *string `json:"rop,omitempty" validate:"required"`
}

type optionals struct {
	O   string  `json:"o"             validate:"omitempty,numeric"`
	OP  *string `json:"op"            validate:"omitempty,numeric"`
	OO  string  `json:"oo,omitempty"  validate:"omitempty,numeric"`
	OOP *string `json:"oop,omitempty" validate:"omitempty,numeric"`
}

var validate = validator.New(validator.WithRequiredStructEnabled())

func check(obj any) {
	fmt.Printf("%#v\n", obj)
	bytes, err := json.MarshalIndent(obj, "", "\t")
	if err != nil {
		log.Println("error in marshal indent:", err)
	}
	println(string(bytes))
	if err := validate.Struct(obj); err != nil {
		fmt.Println(err)
	}
}

func main() {
	var r1 requiridies
	check(r1)
	println()

	var r2 requiridies
	if err := json.NewDecoder(strings.NewReader(`{}`)).Decode(&r2); err != nil {
		log.Println("error in json decode:", err)
	}
	check(r2)
	println()

	var r3 requiridies
	if err := json.NewDecoder(strings.NewReader(`{"r":"r"}`)).Decode(&r3); err != nil {
		log.Println("error in json decode:", err)
	}
	check(r3)
	println()

	var r4 requiridies
	if err := json.NewDecoder(strings.NewReader(`{"rp":"rp"}`)).Decode(&r4); err != nil {
		log.Println("error in json decode:", err)
	}
	check(r4)
	println()

	var r5 requiridies
	if err := json.NewDecoder(strings.NewReader(`{"ro":"ro"}`)).Decode(&r5); err != nil {
		log.Println("error in json decode:", err)
	}
	check(r5)
	println()

	var r6 requiridies
	if err := json.NewDecoder(strings.NewReader(`{"rop":"rop"}`)).Decode(&r6); err != nil {
		log.Println("error in json decode:", err)
	}
	check(r6)

	println("\n--------------------------------------------------\n")

	var o1 optionals
	check(o1)
	println()

	var o2 optionals
	if err := json.NewDecoder(strings.NewReader(`{}`)).Decode(&o2); err != nil {
		log.Println("error in json decode:", err)
	}
	check(o2)
	println()

	var o3 optionals
	if err := json.NewDecoder(strings.NewReader(`{"o":"o"}`)).Decode(&o3); err != nil {
		log.Println("error in json decode:", err)
	}
	check(o3)
	println()

	var o4 optionals
	if err := json.NewDecoder(strings.NewReader(`{"op":"op"}`)).Decode(&o4); err != nil {
		log.Println("error in json decode:", err)
	}
	check(o4)
	println()

	var o5 optionals
	if err := json.NewDecoder(strings.NewReader(`{"oo":"oo"}`)).Decode(&o5); err != nil {
		log.Println("error in json decode:", err)
	}
	check(o5)
	println()

	var o6 optionals
	if err := json.NewDecoder(strings.NewReader(`{"oop":"oop"}`)).Decode(&o6); err != nil {
		log.Println("error in json decode:", err)
	}
	check(o6)
}

package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func Test_NegativeNameTest(t *testing.T) {
	g := NewGomegaWithT(t)

	emp := Employee{
		Name:       "",
		Email:      "wasuta@gmail.com",
		EmployeeID: "J12345678",
	}

	ok, err := govalidator.ValidateStruct(emp)

	g.Expect(ok).ToNot(BeTrue())

	g.Expect(err).ToNot(BeNil())

	g.Expect(err.Error()).To(Equal("ต้องไม่เป็นค่าว่าง"))
}

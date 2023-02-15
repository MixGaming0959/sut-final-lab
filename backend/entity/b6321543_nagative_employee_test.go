package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func Test_NegativeEmployeeTest(t *testing.T) {
	g := NewGomegaWithT(t)

	emp := Employee{
		Name:       "นายวสุธา เศรษฐบุตร",
		Email:      "wasuta@gmail.com",
		EmployeeID: "B12345678",
	}

	ok, err := govalidator.ValidateStruct(emp)

	g.Expect(ok).ToNot(BeTrue())

	g.Expect(err).ToNot(BeNil())

	g.Expect(err.Error()).To(Equal("รหัสพนักงานขึนต้นด้วย J หรือ M หรือ S แล้วตามด้วยตัวเลขจํานวน 8 ตัว"))
}

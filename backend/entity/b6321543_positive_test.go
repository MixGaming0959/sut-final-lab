package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestPass(t *testing.T) {
	g := NewGomegaWithT(t)

	emp := Employee{
		Name:       "นายวสุธา เศรษฐบุตร",
		Email:      "wasuta@gmail.com",
		EmployeeID: "M12345678",
	}

	ok, err := govalidator.ValidateStruct(emp)

	g.Expect(ok).To(BeTrue())

	g.Expect(err).To(BeNil())
}

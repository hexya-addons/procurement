// Copyright 2019 NDP Systèmes. All Rights Reserved.
// See LICENSE file for full licensing details.

package procurement

import (
	"testing"

	"github.com/hexya-addons/base"
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/hexya/src/models/security"
	"github.com/hexya-erp/hexya/src/tests"
	"github.com/hexya-erp/pool/h"
	"github.com/hexya-erp/pool/q"
	. "github.com/smartystreets/goconvey/convey"
)

func TestMain(m *testing.M) {
	tests.RunTests(m, "procurement")
}

func createProcurement(env models.Environment, uid int64, values *h.ProcurementOrderData) h.ProcurementOrderSet {
	proc := h.ProcurementOrder().NewSet(env).Sudo(uid).Create(values)
	changedVals := proc.OnchangeProduct()
	proc.Write(changedVals)
	return proc
}

func TestBase(t *testing.T) {
	Convey("Testing base procurements", t, func() {
		So(models.SimulateInNewEnvironment(security.SuperUserID, func(env models.Environment) {
			userEmployee := h.User().NewSet(env).
				WithContext("no_reset_password", true).
				WithContext("mail_create_nosubscribe", true).
				Create(h.User().NewData().
					SetName("Fabricette Manivelle").
					SetLogin("fabricette").
					SetEmail("f.f@exmaple.com").
					SetGroups(h.Group().Search(env, q.Group().GroupID().Equals(base.GroupUser.ID))))
			h.Group().NewSet(env).ReloadGroups()
			uomUnit := h.ProductUom().Search(env, q.ProductUom().HexyaExternalID().Equals("product_product_uom_unit"))
			uomDunit := h.ProductUom().Create(env, h.ProductUom().NewData().
				SetName("DeciUnit").
				SetCategory(uomUnit.Category()).
				SetFactorInv(0.1).
				SetFactor(10).
				SetUomType("smaller").
				SetRounding(0.001))
			product1 := h.ProductProduct().Create(env, h.ProductProduct().NewData().
				SetName("Courage").
				SetType("consu").
				SetUom(uomUnit).
				SetUomPo(uomDunit))

			Convey("Procurement Order should be in exception as there is no suitable rule", func() {
				procurement := createProcurement(env, userEmployee.ID(),
					h.ProcurementOrder().NewData().
						SetProduct(product1).
						SetName("Procurement Test").
						SetProductQty(15))
				So(procurement.State(), ShouldEqual, "exception")
			})
		}), ShouldBeNil)
	})
}

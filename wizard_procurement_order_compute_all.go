// Copyright 2017 NDP Systèmes. All Rights Reserved.
// See LICENSE file for full licensing details.

package procurement

import (
	"github.com/hexya-erp/hexya/src/actions"
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/pool/h"
	"github.com/hexya-erp/pool/m"
)

func init() {

	h.ProcurementOrderComputeAll().DeclareTransientModel()

	h.ProcurementOrderComputeAll().Methods().ProcureCalculationAll().DeclareMethod(
		`ProcureCalculationAll`,
		func(rs m.ProcurementOrderComputeAllSet) {
			models.ExecuteInNewEnvironment(rs.Env().Uid(), func(env models.Environment) {
				// TODO Avoid to run the scheduler multiple times in the same time
				companies := h.User().NewSet(env).CurrentUser().Companies()
				for _, company := range companies.Records() {
					h.ProcurementOrder().NewSet(env).RunScheduler(true, company)
				}
			})
		})

	h.ProcurementOrderComputeAll().Methods().ProcureCalculation().DeclareMethod(
		`ProcureCalculation`,
		func(rs m.ProcurementOrderComputeAllSet) *actions.Action {
			go rs.ProcureCalculationAll()
			return &actions.Action{
				Type: actions.ActionCloseWindow,
			}
		})

}

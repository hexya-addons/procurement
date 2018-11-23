// Copyright 2017 NDP Systèmes. All Rights Reserved.
// See LICENSE file for full licensing details.

package procurement

import (
	"github.com/hexya-addons/base"
	"github.com/hexya-erp/pool/h"
)

func init() {

	h.ProcurementOrder().Methods().AllowAllToGroup(base.GroupUser)
	h.ProcurementGroup().Methods().AllowAllToGroup(base.GroupUser)
	h.ProcurementRule().Methods().AllowAllToGroup(base.GroupUser)

}

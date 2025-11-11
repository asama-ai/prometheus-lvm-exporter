package main

import "github.com/hansmi/prometheus-lvm-exporter/lvmreport"

var pvsegGroup = &group{
	name:           lvmreport.PVSEG,
	infoMetricName: "pvseg_info",

	keyFields: []*textField{
		{
			fieldName:  "pv_uuid",
			metricName: "pvseg_pv_uuid",
			desc:       "Physical volume UUID",
		},
	},

	textFields: []*textField{
		{
			fieldName:  "lv_uuid",
			metricName: "pvseg_lv_uuid",
			flags:      asInfoLabel,
			desc:       "Logical volume UUID associated with the PV segment",
		},
	},
}

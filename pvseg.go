package main

import "github.com/hansmi/prometheus-lvm-exporter/lvmreport"

var pvsegGroup = &group{
	name:           lvmreport.PVSEG,
	infoMetricName: "pvseg_info",

	keyFields: []*textField{
		{
			fieldName:  "pv_uuid",
			metricName: "pv_uuid",
			flags:      asInfoLabel,
			desc:       "Physical volume UUID",
		},
	},

	textFields: []*textField{
		{
			fieldName:  "lv_uuid",
			metricName: "lv_uuid",
			flags:      asInfoLabel,
			desc:       "Logical volume UUID associated with the PV segment",
		},
	},
}

package utils

import "kub/dashboardES/internal/models"

func FilterChargeDefsByReferenceType(chargedefs []models.ChargeDef, refType string) []models.ChargeDef {
	var filteredChargedefs []models.ChargeDef
	for _, chargedef := range chargedefs {
		if chargedef.ReferenceType == refType {
			filteredChargedefs = append(filteredChargedefs, chargedef)
		}
	}
	return filteredChargedefs
}

package main

import "github.com/sedwards2009/nuview"

// Center returns a new primitive which shows the provided primitive in its
// center, given the provided primitive's size.
func Center(width, height int, p nuview.Primitive) nuview.Primitive {
	subFlex := nuview.NewFlex()
	subFlex.SetDirection(nuview.FlexRow)
	subFlex.AddItem(nuview.NewBox(), 0, 1, false)
	subFlex.AddItem(p, height, 1, true)
	subFlex.AddItem(nuview.NewBox(), 0, 1, false)

	flex := nuview.NewFlex()
	flex.AddItem(nuview.NewBox(), 0, 1, false)
	flex.AddItem(subFlex, width, 1, true)
	flex.AddItem(nuview.NewBox(), 0, 1, false)

	return flex
}

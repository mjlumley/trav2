package main

import "math"

// wbhWorld details a Digest Group Publications (DGP) World Builder's Handbook
// detailed world. These are based on the MegaTraveller world, and if a new
// world is needed, then a MegaTraveller world must be first generated, and
// the WBH detailing is done.

// TODO: Another option is to allow detailing of an existing world.

type wbhWorld struct {
	world // The underlying world to be detailed.

	diameter    int     // World diameter in kilometres
	densityType string  // The planet body density type
	density     float64 // The planet density in earth standard densities
	mass        float64 // The world's mass in standard (earth) masses
	gravity     float64 // The worlds gravity in standard (earth) gees (=9.8m/s/s)

}

// generateWBHWorld generates a detailed MegaTraveller world with the given basic information.
// It returns the world generated.
func generateWBHWorld(name, hexLoc, sector, allegiance, traffic string) (w wbhWorld) {

	// First let's generate the world as a MegaTraveller world
	w.world = generateMTWorld(name, hexLoc, sector, allegiance, traffic)
	w.world.genType = WgtMtWBH

	return
}

// determineDensity determines the planet density type and density in standard (earth) densities. It returns both the density description (like "molten core" or similar)
// as a string and the density as a floating point number.
func (w *wbhWorld) determineDensity() (dt string, dens float64) {
	dm := 0

	roll := D6() + D6()
	if w.world.uwp.sizeInt <= 4 {
		dm++
	}
	if w.world.uwp.sizeInt >= 6 {
		dm = dm - 2
	}
	if w.world.uwp.atmInt <= 3 {
		dm++
	}
	if w.world.uwp.atmInt >= 6 {
		dm = dm - 2
	}
	roll = roll + dm
	roll2 := D6() + D6() + D6()

	switch {
	case roll < 1:
		dt = pdTypeHeavyCore
		dens = 0.95 + 0.05*float64(roll2)
		if roll2 > 13 {
			dens = dens + (float64(roll2)-13.0)*0.05
		}
		if roll2 == 18 {
			dens = 2.25
		}
	case roll >= 15:
		dt = pdTypeIcyBody
		dens = 0.12 + 0.02*float64(roll2)
	case roll >= 11 && roll <= 14:
		dt = pdTypeRockyBody
		dens = 0.44 + 0.02*float64(roll2)
	default:
		dt = pdTypeMoltenCore
		dens = 0.76 + 0.02*float64(roll2)
	}

	return
}

// getDiameterKm determines the diameter for the world in kilometres, from the UWP digit and a variance. This from MT World Builders Handbook.
// Value returned is world diameter is kilometres as an integer.
func (w *wbhWorld) determineDiameterKm() (d int) {
	variance := Flux(0) * 100

	if w.world.uwp.sizeInt == 0 {
		d = variance + 600
	} else {
		d = variance + w.world.uwp.sizeInt*1000
	}
	d = d * 8 / 5 // Convert from miles to kilometres. Integer rounding is OK.
	return
}

// getMass provides the mass of a world in standard (earth) masses. Value returned is the mass as a float.
func (w *wbhWorld) getMass() float64 {
	r := float64(w.world.uwp.sizeInt)
	if r == 0.0 {
		r = 0.6
	}

	return w.density * math.Pow(r/8.0, 3.0)
}

// getGravity gets the gravity of a world in gees. It is based on mass and size. Value returned is the gravity as a float.
func (w *wbhWorld) getGravity() float64 {
	r := float64(w.world.uwp.sizeInt)
	if r == 0.0 {
		r = 0.6
	}
	return w.mass * 64.0 / (r * r)
}

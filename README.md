# Go-Volvo
*A simple Go API client for [Volvo's APIs](https://developer.volvocars.com/apis/)*

Very much work in progress. Hoping to create a small, simple API client for Go.

## Roadmap

- [ ] Initiate auth code flow
- [ ] Exchange auth code for access token
- [ ] Auto refresh access token
- [ ] Full API Support
  - [ ] Connected Vehicle API (in progress)
  - [ ] Energy API
  - [ ] Extended Vehicle API
  - [ ] Location API

## Spec vs. Actual Responses

The [official Volvo docs](https://developer.volvocars.com/apis/) are great, and have been invaluable for this project. I have noted a few what appear to be errors in the documentation, however.

* `/connected-vehicle/v2/vehicles/{vin}/doors` documents `carLocked` named object being returned, but actual object name is `centralLock`
* `/connected-vehicle/v2/vehicles/{vin}/statistics` documents `averageFuelConsumptionAutomatic` and `averageEnergyConsumptionAutomatic` named objects, but these are seemingly never returned
* `/connected-vehicle/v2/vehicles/{vin}/vehicle` documents `interiorImageURL` named object being returned, but actual object name is `internalImageUrl`
* `/connected-vehicle/v2/vehicles/{vin}/diagnostics` documents `serviceTrigger` named object being returned, but this seemingly never returns. Additionally, `washerFluidLevelWarning` returns but is not documented
* `/connected-vehicle/v2/vehicles/{vin}/windows` returns a `sunroof` objects, which is not documented

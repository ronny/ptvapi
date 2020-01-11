# ptvapi

This is a client library to interact with the Public Transport Victoria (PTV)
Timetable API v3.

This library is not an official library from PTV.

More information about the PTV Timetable API:
<a href="https://www.ptv.vic.gov.au/footer/data-and-reporting/datasets/ptv-timetable-api/">https://www.ptv.vic.gov.au/footer/data-and-reporting/datasets/ptv-timetable-api/</a>

This library is mostly auto generated from the published swagger file:
<a href="http://timetableapi.ptv.vic.gov.au/swagger/ui/index">http://timetableapi.ptv.vic.gov.au/swagger/ui/index</a>
using the go-swagger tool:
<a href="https://github.com/go-swagger/go-swagger">https://github.com/go-swagger/go-swagger</a>

See the [transport](transport) package for information on how to use a client that
automatically includes the `devid` and `signature` params as required by the
PTV Timetable API v3.

See [examples](examples) to see examples on how to use this library.

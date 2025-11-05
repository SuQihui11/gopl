package meterconv

func MToKm(m Meter) Kilometer { return Kilometer(m * 0.001) }

func KmToM(km Kilometer) Meter { return Meter(km * 1000) }

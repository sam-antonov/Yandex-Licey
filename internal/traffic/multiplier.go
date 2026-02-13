package traffic

func GetTrafficMultiplier(trafficLevel int) float64 {
    return 1.0 + float64(trafficLevel-1)*0.1
}
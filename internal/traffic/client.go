package traffic

type TrafficClient interface {
    GetTrafficLevel(lat, lng float64) int // 1–5
}

type RealTrafficClient struct{}

func (c *RealTrafficClient) GetTrafficLevel(lat, lng float64) int {
    return 3 // Константное значение в нашем примере, в реальности оно будет вычисляться сервисом Яндекс Карты
}
# json 구조체 예제

data := []byte(`{"a": [["12.58861425",10.52046452],["12.58861426",4.1073]],"b": [["12.58861425",10.52046452],["12.58861426",4.1073]],"c": "true","d": 1234}`)

type OrderBook struct {
    A [][2]interface{} `json:"a"`
    B [][2]interface{} `json:"b"`
    C string           `json:"c"`
    D uint32           `json:"d"`
}
orders := new(OrderBook)
err := json.Unmarshal(data, &orders)
if err != nil {
    fmt.Printf(err.Error())
} else {
    fmt.Printf("%s\n", orders.A[0][0].(string))     
    fmt.Printf("%f\n", orders.A[0][1].(float64))
}

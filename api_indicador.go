package main

import (
  "fmt"
  "net/http"
  "io/ioutil"
  "encoding/json"
  "reflect"
)

type Information struct {
  Version string `json:"version"`
  Autor string `json:"autor"`
  Fecha string  `json:"fecha"`
  Uf IndicadorData `json:"uf"`
  IVP IndicadorData `json:"ivp"`
  Dolar IndicadorData `json:"dolar"`
  DolarInterCambio IndicadorData `json:"dolar_intercambio"`
  Euro IndicadorData `json:"euro"`
  IPC IndicadorData `json:"ipc"`
  UTM IndicadorData `json:"utm"`
  IMACEC IndicadorData `json:"imacec"`
  TPM IndicadorData `json:"tpm"`
  LibraCobre IndicadorData `json:"libra_cobre"`
  TasaDesempleo IndicadorData `json:"tasa_desempleo"`
  Bitcoin IndicadorData `json:"bitcoin"`
}

type IndicadorData struct {
  Codigo string `json:"codigo"`
  Nombre string `json:"nombre"`
  UnidadMedida string `json:"unidad_medida"`
  Fecha string `json:"fecha"`
  Valor float64 `json:"valor"`
}

var indicadores = []string{"Uf", "IVP", "Dolar", "DolarInterCambio", "Euro", "IPC", "UTM", "IMACEC", "TPM", "LibraCobre", "TasaDesempleo", "Bitcoin"}
var attributes = []string{"Codigo", "Nombre", "UnidadMedida", "Fecha", "Valor"}

const ( 
  URL = "https://mindicador.cl/api"
  METHOD = "GET"
)

func log(err error) {
  if err != nil {
    fmt.Println(err)
  }
}

func main() {
  client := &http.Client{}
  req, err := http.NewRequest(METHOD, URL, nil)
  log(err)
  res, err := client.Do(req)
  defer res.Body.Close()
  body, err := ioutil.ReadAll(res.Body)
  info := Information{}
	jsonErr := json.Unmarshal(body, &info)
  log(jsonErr)
  fmt.Println("API DESCRIPTION")
  info.apiDesc()
  fmt.Println("\nDATA")
  info.reflectData()
}

func (info Information) apiDesc(){
  fmt.Println("Version:",info.Version, "Autor:",info.Autor, "Fecha:",info.Fecha)
}

func (info Information) reflectData() {
  t := reflect.ValueOf(info)
  for _, value := range indicadores{
    reflect_struct := t.FieldByName(value)
    for _, value := range attributes {
      fmt.Println(value,":", reflect_struct.FieldByName(value))
    }
    fmt.Println()
  }
}

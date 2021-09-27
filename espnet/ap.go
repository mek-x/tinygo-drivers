package espnet

/*
#cgo CFLAGS: -DCONFIG_IDF_TARGET_ESP32C3
#cgo CFLAGS: -Iinclude
#cgo CFLAGS: -Iesp-idf/components/esp_common/include
#cgo CFLAGS: -Iesp-idf/components/esp_event/include
#cgo CFLAGS: -Iesp-idf/components/esp_netif/include
#cgo CFLAGS: -Iesp-idf/components/esp_wifi/include

#cgo LDFLAGS: -Lesp-idf/components/esp_wifi/lib/esp32c3 -lnet80211 -lpp -lphy -lmesh -lcore
#cgo LDFLAGS: -Tesp-idf/components/esp_rom/esp32c3/ld/esp32c3.rom.ld

#include "esp_private/wifi.h"
#include "esp_wifi_types.h"
#include "espnet.h"
*/
import "C"

type ESPWiFi struct {
}

var WiFi = &ESPWiFi{}

type Config struct {
}

var internalConfig = C.wifi_init_config_t{
	osi_funcs: &C.g_wifi_osi_funcs,
}

func (wifi ESPWiFi) Configure(config Config) error {
	return makeError(C.esp_wifi_init_internal(&internalConfig))
}

func (wifi ESPWiFi) AccessPointMAC() ([6]byte, error) {
	var mac [6]byte
	errCode := C.esp_wifi_get_mac(C.ESP_IF_WIFI_AP, &mac[0])
	return mac, makeError(errCode)
}
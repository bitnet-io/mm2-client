package main

import (
	"fmt"
	"github.com/kpango/glg"
	"mm2_client/config"
	"mm2_client/constants"
	"mm2_client/services"
	"net/url"
	"syscall/js"
)

func startPriceService() js.Func {
	jsfunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if constants.GPricesServicesRunning {
			_ = glg.Warn("Price service already running - skipping")
			return nil
		}
		if len(config.GCFGRegistry) == 0 {
			_ = glg.Warn("Desktop cfg need to be loaded first before running the price service")
			return nil
		}
		services.LaunchServices()
		return nil
	})
	return jsfunc
}

func loadDesktopCfgFromUrl() js.Func {
	jsfunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		_ = glg.Info("load_desktop_cfg_from_url called")
		if len(args) != 1 {
			usage := "invalid nb args - usage: load_desktop_cfg_from_url(\"my_url_to_desktop_cfg\")"
			_ = glg.Error(usage)
			result := map[string]interface{}{
				"error": usage,
			}
			return result
		}
		inputUrl := args[0].String()
		_, err := url.ParseRequestURI(inputUrl)
		if err != nil {
			errStr := fmt.Sprintf("invalid url: %v\n", err)
			_ = glg.Errorf("%s", errStr)
			result := map[string]interface{}{
				"error": errStr,
			}
			return result
		}
		_ = glg.Infof("url is: %s", inputUrl)
		go func() {
			err = config.ParseDesktopRegistryFromUrl(inputUrl)
			if err != nil {
				errStr := fmt.Sprintf("error when parsing cfg: %v\n", err)
				_ = glg.Errorf("%s", errStr)
			}
			_ = glg.Infof("cfg successfully parsed: %d", len(config.GCFGRegistry))
		}()
		return true
	})
	return jsfunc
}

func main() {
	glg.Get().SetMode(glg.STD)
	_ = glg.Info("Hello from webassembly")
	js.Global().Set("load_desktop_cfg_from_url", loadDesktopCfgFromUrl())
	//js.Global().Set("load_desktop_cfg_from_string", startPriceService())
	//js.Global().Set("load_desktop_cfg_from_file", startPriceService())
	js.Global().Set("start_price_service", startPriceService())
	<-make(chan bool)
}
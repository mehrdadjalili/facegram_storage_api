package requestlimiter

import (
	"errors"
	"facegram_file_server/config"
	"facegram_file_server/pkg/cache"
	"time"
)

//CheckAndUpdate connection limiter
func CheckAndUpdate(id string, cfg *config.RequestLimiterItemDetail) (result string, t int, tt string, e error) {
	count, e0 := cache.GetClientConnection(cfg.Key + id)
	if e0 != nil {
		return "", 0, "", errors.New("error")
	} else {
		if count == 0 || count < cfg.CountRequest {
			cnt := count + 1
			result = "open"
			t = 0
			tt = ""
			var TimeD time.Duration
			if count == 0 {
				TimeD = buildTimeNumber(cfg.LimitTime, cfg.TimeType)
			} else {
				ttl, e1 := cache.GetKeyTTl(cfg.Key + id)
				if e1 != nil {
					return "", 0, "", e1
				}
				TimeD = ttl
			}
			e2 := cache.SetClientConnection(cfg.Key+id, cnt, TimeD)
			if e2 != nil {
				return "", 0, "", e2
			}
			return
		} else {
			result = "close"
			t = cfg.LimitTime
			tt = cfg.TimeType
			return result, t, tt, nil
		}
	}
}

//CheckOnly connection limiter
func CheckOnly(id string, cfg *config.RequestLimiterItemDetail) (result string, t int, tt string, e error) {
	count, e0 := cache.GetClientConnection(cfg.Key + id)
	if e0 != nil {
		return "", 0, "", errors.New("error")
	} else {
		if count == 0 || count < cfg.CountRequest {
			result = "open"
			t = 0
			tt = ""
			return
		} else {
			result = "close"
			t = cfg.LimitTime
			tt = cfg.TimeType
			return result, t, tt, nil
		}
	}
}

//Set connection limiter
func Set(id string, cfg *config.RequestLimiterItemDetail) (bool, error) {
	count, e0 := cache.GetClientConnection(cfg.Key + id)
	if e0 != nil {
		return false, e0
	} else {
		cnt := count + 1
		var TimeD time.Duration
		if count == 0 {
			TimeD = buildTimeNumber(cfg.LimitTime, cfg.TimeType)
		} else {
			ttl, e1 := cache.GetKeyTTl(cfg.Key + id)
			if e1 != nil {
				return false, e1
			}
			TimeD = ttl
		}
		e1 := cache.SetClientConnection(cfg.Key+id, cnt, TimeD)
		if e1 != nil {
			return false, e1
		}
		return true, nil
	}
}

func buildTimeNumber(count int, t string) time.Duration {
	if t == "m" {
		return time.Duration(count * 60 * 1000 * 1000 * 1000)
	} else if t == "h" {
		return time.Duration(count * 60 * 60 * 1000 * 1000 * 1000)
	} else if t == "d" {
		return time.Duration(count * 24 * 60 * 60 * 1000 * 1000 * 1000)
	} else {
		return time.Duration(count * 60 * 1000 * 1000 * 1000)
	}
}

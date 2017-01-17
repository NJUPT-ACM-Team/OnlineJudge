package config

import (
	"errors"
	"sync"
	"sync/atomic"
)

type RuntimeConfig struct {
}

type AtomicRuntimeConfig struct {
	v  *atomic.Value
	mu sync.Mutex
}

type ARCfg AtomicRuntimeConfig
type RCfg RuntimeConfig

func (this *AtomicRuntimeConfig) Load() *RuntimeConfig {
	return this.v.Load().(*RuntimeConfig)
}

func (this *AtomicRuntimeConfig) Store(cfg RuntimeConfig) {
	this.mu.Lock()
	defer this.mu.Unlock()
	this.v.Store(&cfg)
}

func LoadRCfgFromARCfg(arcfg *AtomicRuntimeConfig) *RuntimeConfig {
	if arcfg == nil {
		return nil
	}
	return arcfg.Load()
}

func GetRCfg(arcfg *AtomicRuntimeConfig,
	getter func(rcfg *RuntimeConfig) interface{}) (interface{}, error) {

	rcfg := LoadRCfgFromARCfg(arcfg)
	if rcfg == nil {
		return nil, errors.New("RuntimeConfig is nil, maybe due to not store yet")
	}
	return getter(rcfg), nil
}

func SetRCfg(arcfg *AtomicRuntimeConfig, v interface{},
	setter func(rcfg *RuntimeConfig, v interface{})) error {

	rcfg := LoadRCfgFromARCfg(arcfg)
	if rcfg == nil {
		return errors.New("RuntimeConfig is nil, maybe due to not store yet")
	}
	setter(rcfg, v)
	arcfg.Store(*rcfg)
	return nil
}

package funcs

import (
	log "github.com/cihub/seelog"

	"github.com/anchnet/service-monitor/windows-agent/tools/mem"
	"github.com/open-falcon/common/model"
)

func MemMetrics() []*model.MetricValue {
	m, err := mem.VirtualMemory()
	if err != nil {
		log.Info("Get memory fail: ", err)
		return nil
	}

	return []*model.MetricValue{
		GaugeValue("mem.memtotal", m.Total),
		GaugeValue("mem.memused", m.Used),
		GaugeValue("mem.memfree", m.Available),
		GaugeValue("mem.memfree.percent", 100.0-m.UsedPercent),
		GaugeValue("mem.memused.percent", m.UsedPercent),
	}

}

package clients

import (
	"code.cloudfoundry.org/cli/util/configv3"
	"fmt"
	noaaconsumer "github.com/cloudfoundry/noaa/consumer"
	"github.com/cloudfoundry/sonde-go/events"
	"strings"
	"time"
)

const LogTimestampFormat = "2006-01-02T15:04:05.00-0700"

type NOAAHelper struct {
	consumer    *noaaconsumer.Consumer
	configStore *configv3.Config
}

func NewNOAAHelper(consumer *noaaconsumer.Consumer, configStore *configv3.Config) *NOAAHelper {
	return &NOAAHelper{
		consumer:    consumer,
		configStore: configStore,
	}
}

func (c NOAAHelper) RecentLogs(appGUID string, maxMessages int) (string, error) {
	logMsgs, err := c.consumer.RecentLogs(appGUID, c.configStore.AccessToken())
	if err != nil {
		return "", err
	}
	maxLen := maxMessages
	if maxLen < 0 || len(logMsgs) < maxLen {
		maxLen = len(logMsgs)
	}
	if maxLen-1 < 0 {
		return "", nil
	}
	logs := ""
	for i := maxLen - 1; i >= 0; i-- {
		logMsg := logMsgs[i]
		t := time.Unix(0, logMsg.GetTimestamp()).In(time.Local).Format(LogTimestampFormat)
		typeMessage := "OUT"
		if logMsg.GetMessageType() != events.LogMessage_OUT {
			typeMessage = "ERR"
		}
		header := fmt.Sprintf("%s [%s/%s] %s ",
			t,
			logMsg.GetSourceType(),
			logMsg.GetSourceInstance(),
			typeMessage,
		)
		message := string(logMsg.GetMessage())
		for _, line := range strings.Split(message, "\n") {
			logs += fmt.Sprintf("\t%s%s\n", header, strings.TrimRight(line, "\r\n"))
		}
	}
	return logs, nil
}

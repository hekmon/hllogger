# High Level Logger

[![GoDoc](https://pkg.go.dev/badge/github.com/hekmon/hllogger/v2?status.svg)](https://pkg.go.dev/github.com/hekmon/hllogger/v2)


HlLogger is an opinionated logger on top of the Golang standard library logger. It adds:

* syslog inspired logging facilities such as:
  * Debug
  * Info
  * Notice
  * Warning
  * Error
  * Critical
  * Alert
  * Emergency
* automatic systemd-journald integration if systemd is detected as executor and output is `os.Stdout` or `os.Stderr`
* automatic AWS CloudWatch Logs integration if AWS Lambda execution environment is detected

## Installation

```bash
go get github.com/hekmon/hllogger/v2
```

## Usage

```golang
package main

import (
	"os"

	"github.com/hekmon/hllogger/v2"
)

func main() {
	logger := hllogger.New(os.Stdout, hllogger.Debug)
	logger.Debug("test")
	logger.Info("test")
	logger.Notice("test")
	logger.Warning("test")
	logger.Error("test")
	logger.Critical("test")
	logger.Alert("test")
	logger.Emergency("test")
}
```

If you init the logger with the `hllogger.Info` level, all calls to `logger.Debug()` will no-op (won't print anything). If you init the logger with the `hllogger.Notice` level, all calls to to `logger.Debug()` and `logger.Info()` will no-op (won't print anything). And so on...

### Regular output

Regular (default) ouput on your terminal or on a file.

```raw
2022/02/14 20:37:31     DEBUG: test
2022/02/14 20:37:31      INFO: test
2022/02/14 20:37:31    NOTICE: test
2022/02/14 20:37:31   WARNING: test
2022/02/14 20:37:31     ERROR: test
2022/02/14 20:37:31  CRITICAL: test
2022/02/14 20:37:31     ALERT: test
2022/02/14 20:37:31 EMERGENCY: test
```

### systemd-journald integration

Automatically enabled if started by systemd and output is `os.Stdout` or `os.Stderr`. By integrating with systemd-journald, the lib will indicate the log level of each log to journald. It will allows log manipulation directly with journald (printing only info level while debug have been printed) but also log level highlighting (colors).

#### log level highlighting

![systemd-journald output with color highlighting](journald.png "systemd-journald output")

#### emergency broadcast

```bash
hekmon@testserver1:~$ sudo systemctl start logleveltest.service

Broadcast message from systemd-journald@testserver1 (Mon 2022-02-14 20:47:52 UTC):

hlloggertest[1031126]: EMERGENCY: test

hekmon@testserver1:~$
```

### AWS Lambda integration

If a AWS lambda execution environment is detected and therefor logs are being handled by AWS Cloudwatch logs, the logger won't use any time based flags as CloudWatch Logs will be taking care of it. Only log level and actual log message will be output:

```raw
    DEBUG: test
     INFO: test
   NOTICE: test
  WARNING: test
    ERROR: test
 CRITICAL: test
    ALERT: test
EMERGENCY: test
```
